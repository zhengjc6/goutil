package main

import (
	"fmt"
	"goutil/csvparse"
	"goutil/redislocker"
	"goutil/snowflake"

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

func testsnowflake() {
	if err := snowflake.NewSnow(1005); err != nil {
		fmt.Println(err)
		return
	}
	snowins, err := snowflake.Instanse()
	if err != nil {
		fmt.Println(err)
		return
	}
	mp := map[int64]bool{}
	maxN := 10000
	ch := make(chan int64, maxN)
	for i := 0; i < maxN; i++ {
		go func() {
			uid := snowins.NextVal()
			fmt.Println(uid)
			ch <- uid
		}()
	}
	for {
		select {
		case data := <-ch:
			if _, ok := mp[data]; ok {
				fmt.Println("same uid!")
			} else {
				mp[data] = true
			}
		case <-time.Tick(time.Second * 5):
			fmt.Println("running,mapsize = ", len(mp))
		}
	}
}

func main() {
	// 	testredislocker()
	// 	time.Sleep(time.Minute)
	//testcsvparse()
	//testcsvparseread()
	// fmt.Println("return:", testdefer())
	// fmt.Println("return:", testdefer2())
	testsnowflake()
}
