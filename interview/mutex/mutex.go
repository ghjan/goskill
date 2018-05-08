package main

import (
"sync"
"fmt"
)

type UserAges struct {
	ages map[string]int
	sync.RWMutex  //不是Mutex，而是读写锁
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	ua.RLock()
	defer ua.RUnlock()
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}

//Set 加锁了，Get 也得加锁。这里最好使用 sync.RWMutex
func main() {
	ages := make(map[string]int)
	ua := UserAges{ages: ages}
	fmt.Println(ua.Get("dav"))
	ua.Add("david", 30)
	ua.Add("alice", 26)
	fmt.Println(ua.Get("dav"))
	fmt.Println(ua.Get("david"))
	fmt.Println(ua.Get("alice"))
	fmt.Println(ua.Get("ali"))
}
