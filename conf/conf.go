package conf

import ()

type Config struct {
	P1        int8
	P2        int8
	BoardSize int
}

var Conf chan Config
var SetConf chan Config

func init() {
	go func() {
		Conf = make(chan Config)
		SetConf = make(chan Config)
		conf := <-SetConf
		for {
			select {
			case conf = <-SetConf:
			case Conf <- conf:
			}
		}
	}()
}
