package snowflake

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/golang/glog"
)

const (
	epoch          = int64(1664553600000)              // 设置起始时间(时间戳/毫秒)：2022-10-01 00:00:00，有效期34年
	timestampBits  = uint(40)                          // 时间戳占用位数
	workeridBits   = uint(11)                          // 服务id所占位数
	sequenceBits   = uint(12)                          // 序列所占的位数
	timestampMax   = int64(-1 ^ (-1 << timestampBits)) // 时间戳最大值
	workeridMax    = int64(-1 ^ (-1 << workeridBits))  // 支持的最大服务id数量
	sequenceMax    = int64(-1 ^ (-1 << sequenceBits))  // 支持的最大序列id数量
	workeridShift  = sequenceBits                      // 服务id左移位数
	timestampShift = sequenceBits + workeridBits       // 时间戳左移位数
)

type Snowflake struct {
	mutex     sync.Mutex // 锁
	timestamp int64      // 时间戳 ，毫秒
	workerid  int64      // 服务id
	sequence  int64      // 序列号
}

var instanse *Snowflake

func NewSnow(wid int64) error {
	if wid < 0 || wid > workeridMax {
		return fmt.Errorf("workeridMax limit = %d", workeridMax)
	}
	instanse = &Snowflake{
		mutex:     sync.Mutex{},
		timestamp: 0,
		workerid:  wid,
		sequence:  0,
	}
	return nil
}

func Instanse() (*Snowflake, error) {
	if instanse == nil {
		return nil, errors.New("snowflake not init")
	}
	return instanse, nil
}

type InterfaceSnowFlake interface {
	NextVal() int64
}

func (s *Snowflake) NextVal() int64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	now := time.Now().UnixMilli() //毫秒
	if s.timestamp == now {
		// 当同一时间戳（精度：毫秒）下多次生成id会增加序列号
		s.sequence = (s.sequence + 1) & sequenceMax
		if s.sequence == 0 {
			// 如果当前序列超出12bit长度，则需要等待下一毫秒
			// 下一毫秒将使用sequence:0
			for now <= s.timestamp {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		// 不同时间戳（精度：毫秒）下直接使用序列号：0
		s.sequence = 0
	}
	t := now - epoch
	if t > timestampMax {
		glog.Errorf("epoch must be between 0 and %d", timestampMax-1)
		return 0
	}
	s.timestamp = now
	r := int64((t)<<timestampShift | (s.workerid << workeridShift) | (s.sequence))
	return r
}
