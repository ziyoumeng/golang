package main

import (
	"fmt"

	"github.com/micro/go-config"
)

/*
测试watch功能.
TODO:
发现vim修改监听项时，只有:wq退出，监听才会被触发.比如先w然后q，竟然不行。没有找到原因。猜测和刷入到硬盘的时机有关
*/
func main() {
	//conf := config.NewConfig()
	err := config.LoadFile("config.json")
	if err != nil {
		panic(err)
	}

	for {
		w, err := config.DefaultConfig.Watch("hosts", "database")
		if err != nil {
			panic(err)
		}

		fmt.Printf("wait next..")
		v, err := w.Next()
		if err != nil {
			panic(err)
		}

		var host Host
		err = v.Scan(&host)
		if err != nil {
			panic(err)
		}

		fmt.Printf("host = %+v\n", host)
	}
}

type Host struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
}

type Config struct {
	Hosts map[string]Host `json:"hosts"`
}
