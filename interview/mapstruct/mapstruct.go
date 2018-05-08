package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("--------paseStudentBad------")
	paseStudentBad()
	fmt.Println("--------paseStudentBad------")
	paseStudentGood()
}

func paseStudentBad() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	//在for循环过程中，stu是一个变量（不是每轮为一个新的变量）
	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	for k, v := range m {
		fmt.Printf("k=%v,v=%v\n", k, v)
	}
}

func paseStudentGood() {
	m := make(map[string]student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		m[stu.Name] = stu
	}
	for k, v := range m {
		fmt.Printf("k=%v,v=%v\n", k, v)
	}
}

//TODO:map不是并发安全的?
