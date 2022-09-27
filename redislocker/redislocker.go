package redislocker

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type RedisLocker struct {
	client    *redis.Client
	lockKey   string
	lockValue int64
	unlockCh  chan struct{} //用户解锁通知通道
	holdTime  time.Duration //持有时间  ns
	tryCount  int32         //lock时尝试次数
	tryGap    time.Duration //lock时尝试间隔 ns
}

type RedisLockerWorker interface {
	Lock() error

	Unlock() error
}

func (rl *RedisLocker) Lock() error {
	var resp *redis.BoolCmd
	for i := 0; i < int(rl.tryCount); i++ {

		resp = rl.client.SetNX(rl.lockKey, rl.lockValue, rl.holdTime) //返回执行结果

		lockSuccess, err := resp.Result()

		if err == nil && lockSuccess {

			//抢锁成功，开启看门狗 并跳出，否则失败继续自旋

			go watchDog(rl)

			return nil

		} else {
			//fmt.Printf("%v spin %d\n", rl.lockValu e,i)
			time.Sleep(time.Duration(rl.tryGap)) //休眠
		}
	}
	return errors.New("Lock Time Out")
}

func (rl *RedisLocker) Unlock() error {

	script := redis.NewScript(`
	  	if redis.call('get', KEYS[1]) == ARGV[1]	then 
			return redis.call('del', KEYS[1]) 
		else 
			return 0  
	   	end
	`)

	resp := script.Run(rl.client, []string{rl.lockKey}, rl.lockValue)

	if result, err := resp.Result(); err != nil || result == 0 {

		return errors.New(fmt.Sprintf("unlock failed:", err))

	} else {

		//删锁成功后，通知看门狗退出

		rl.unlockCh <- struct{}{}
	}
	return nil
}

//自动续期看门狗

func watchDog(rl *RedisLocker) {

	// 创建一个定时器NewTicker, 每隔2秒触发一次,类似于闹钟

	expTicker := time.NewTicker(time.Second * time.Duration(rl.holdTime/5*4))

	//确认锁与锁续期打包原子化

	script := redis.NewScript(`
		if redis.call('get', KEYS[1]) == ARGV[1] then 
			return redis.call('expire', KEYS[1], ARGV[2]) 
		else 
			return 0 
	  	end
	`)

	for {

		select {

		case <-expTicker.C: //定时器，所以每隔80%*holdTime都会触发

			resp := script.Run(rl.client, []string{rl.lockKey}, rl.lockValue, rl.holdTime)

			if result, err := resp.Result(); err != nil || result == int64(0) {

				//续期失败

				log.Println("expire lock failed", err)

			}

		case <-rl.unlockCh: //任务完成后用户解锁通知看门狗退出

			return

		}

	}

}

func NewLocker(rclient *redis.Client, key string, val int64) *RedisLocker {
	locker := RedisLocker{}
	locker.client = rclient
	locker.lockKey = key
	locker.lockValue = val
	locker.holdTime = 10 * time.Second
	locker.tryCount = 100
	locker.tryGap = time.Second / 100
	locker.unlockCh = make(chan struct{}, 0)
	return &locker
}
