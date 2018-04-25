package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name   string "user name" //这小米点里面的就是tag
	Passwd string "user passsword"
}

type User2 struct {
	Name   string `json:"name,omitempty"`
	Passwd string `json:"passwd,omitempty"`
}

func test_struct_tag() {
	user := &User{"chronos", "pass"}
	s := reflect.TypeOf(user).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag) //将tag输出出来
	}    
	user2 := &User2{"chronos", "pass"}
	s2 := reflect.TypeOf(user2).Elem() //通过反射获取type定义
	for i := 0; i < s2.NumField(); i++ {
		fmt.Println(s2.Field(i).Tag) //将tag输出出来
	}

    u_json, _ := json.Marshal(user)
    fmt.Println(string(u_json))
    
    u2_json, _ := json.Marshal(user2)
    fmt.Println(string(u2_json))

	u := &User{Name: "tony", Passwd: "tony"}
	j, _ := json.Marshal(u)
    fmt.Println(string(j))

}
