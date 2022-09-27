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
	unlockCh  chan struct{} //the unlock chan used to notice watdog co to exit
	holdTime  time.Duration //lock duration
	tryCount  int32         //max trying times
	tryGap    time.Duration //the interval time of tring to get lock
}

type RedisLockerWorker interface {
	Lock() error

	Unlock() error
}

func (rl *RedisLocker) Lock() error {
	var resp *redis.BoolCmd
	for i := 0; i < int(rl.tryCount); i++ {

		resp = rl.client.SetNX(rl.lockKey, rl.lockValue, rl.holdTime)

		lockSuccess, err := resp.Result()

		if err == nil && lockSuccess {

			//lock sucessfully, run watchDog

			go watchDog(rl)

			return nil

		} else {
			//sleep && spin
			time.Sleep(time.Duration(rl.tryGap))
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

		return fmt.Errorf("unlock failed:%s", err.Error())

	} else {

		//notice watchDog to exit

		rl.unlockCh <- struct{}{}
	}
	return nil
}

// watch dog prolinging period of lock validity
func watchDog(rl *RedisLocker) {

	expTicker := time.NewTicker(time.Second * time.Duration(rl.holdTime/5*4))

	//using lua script to make sure atomic process

	script := redis.NewScript(`
		if redis.call('get', KEYS[1]) == ARGV[1] then 
			return redis.call('expire', KEYS[1], ARGV[2]) 
		else 
			return 0 
	  	end
	`)

	for {

		select {

		case <-expTicker.C:

			resp := script.Run(rl.client, []string{rl.lockKey}, rl.lockValue, rl.holdTime)

			if result, err := resp.Result(); err != nil || result == int64(0) {

				log.Println("expire lock failed", err)

			}

		case <-rl.unlockCh: //receive exit ch

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
