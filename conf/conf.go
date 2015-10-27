package conf

import (
//"fmt"
)

type Config struct {
	P1        int8
	P2        int8
	BoardSize int
	Depth     int
}

var Conf chan Config
var SetConf chan Config

func init() {
	Conf = make(chan Config)
	SetConf = make(chan Config)
	conf := Config{
		Depth: 2,
		P1:    1,
		P2:    2,
	}
	go func() {
		for {
			select {
			case conf = <-SetConf:
				//				fmt.Println("recieved conf", conf)
			case Conf <- conf:
				//				fmt.Println("sedning conf ")
			}
		}
	}()
}
