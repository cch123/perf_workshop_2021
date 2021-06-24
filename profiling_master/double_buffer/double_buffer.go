package main

import (
	"sync"
	"sync/atomic"
)

var config atomic.Value

func init() {
	var initialConfig = &MyConfig{
		WhiteList: make(map[int]struct{}),
	}
	config.Store(initialConfig)
}

type MyConfig struct {
	WhiteList map[int]struct{}
}

func getConfig() *MyConfig {
	return config.Load().(*MyConfig)
}

func updateConfig1() {
	var newConfig = &MyConfig{
		WhiteList: make(map[int]struct{}),
	}

	// do a lot of computation
	for i := 0; i < 10000; i++ {
		newConfig.WhiteList[i] = struct{}{}
	}

	config.Store(newConfig)
}

// partial update
func updateConfig2() {
	var oldConfig = getConfig()
	var newConfig = &MyConfig{
		WhiteList: make(map[int]struct{}),
	}

	// copy from old
	for k, v := range oldConfig.WhiteList {
		newConfig.WhiteList[k] = v
	}

	// add some new keys
	newConfig.WhiteList[121212] = struct{}{}
	newConfig.WhiteList[23333] = struct{}{}

	config.Store(newConfig)
}

var updateLock sync.Mutex

// 如果 update 本身可能出现并发
func updateConfig3() {
	// lock update
	updateLock.Lock()
	defer updateLock.Unlock()

	var oldConfig = getConfig()
	var newConfig = &MyConfig{
		WhiteList: make(map[int]struct{}),
	}

	// copy from old
	for k, v := range oldConfig.WhiteList {
		newConfig.WhiteList[k] = v
	}

	// add some new keys
	newConfig.WhiteList[121212] = struct{}{}
	newConfig.WhiteList[23333] = struct{}{}

	config.Store(newConfig)

}

func main() {}
