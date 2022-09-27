package main

import (
	"fmt"
	"goutil/csvparse"
	"goutil/redislocker"

	"math/rand"
	"os"
	"strings"
	"time"

	//csvdata "goutil/static/conf"

	"github.com/go-redis/redis"
)

func testredislocker() {
	for i := 1; i < 20; i++ {
		go func(idx int) {
			tick := time.Now()
			defer func() {
				fmt.Printf("%d,cost time %d \n", idx, time.Now().Sub(tick).Milliseconds())
			}()
			client := redis.NewClient(&redis.Options{

				Addr: "127.0.0.1:6379",

				Password: "zhengjc3225",

				DB: 0, //redis默认拥有16个db（0~15），且默认连接db0
			})
			locker := redislocker.NewLocker(client, "testkey", int64(idx))

			err := locker.Lock()
			if err != nil {
				fmt.Printf("%d,lock fail,%s\n", idx, err.Error())
				return
			}
			rand.Seed(time.Now().UnixNano())
			tm := time.Second / 100 * time.Duration(rand.Intn(10)+10)
			fmt.Printf("%d,do some job,%d\n", idx, tm.Milliseconds())
			time.Sleep(tm)
			err = locker.Unlock()
			if err != nil {
				fmt.Printf("%d,unlock fail,%s\n", idx, err.Error())
				return
			} else {
				fmt.Printf("%d,finish job\n", idx)
			}
		}(i)
	}
}

// func testcsvparseread() {
// 	data := csvdata.CSVData
// 	data2 := data.UserTable[11]
// 	fmt.Printf("%+v\n", data)
// 	fmt.Printf("%+v\n", data2)
// }

func testcsvparse() {
	pwd, _ := os.Getwd()
	fileName := "config"
	var bt strings.Builder
	bt.WriteString(pwd)
	bt.WriteString("\\")
	bt.WriteString(fileName)
	outDir := pwd + "\\static\\conf"
	if err := os.RemoveAll(outDir); err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	csvparse.ParseDir(bt.String(), outDir, "csvdata")
}

func main() {
	// 	testredislocker()
	// 	time.Sleep(time.Minute)
	testcsvparse()
	//testcsvparseread()
}
