package main

func main() {
	initConfig()
}

// 我用的 mq 的 sdk 线下好像读的配置选项不太对，我想看看他到底读了啥
type Config struct {
	Address string
}

var correctConfig = Config {
	Address : "10.100.1.131",
}

var wrongConfig = Config {
	Address : "127.0.0.1",
}

func initConfig() {
	// a lot of read code
	var conf = wrongConfig
	useConf(conf)
}

func useConf(c Config) {}
