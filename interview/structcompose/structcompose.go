package main

import "fmt"

func main() {
	fmt.Println(stringutil.Reverse("!selpmaxe oG ,olleH"))

	fmt.Println("------test1-------------")
	test1()
	fmt.Println("------test2-------------")
	test2()
	fmt.Println("------anonymouseTest-------------")
	anonymouseTest()
}

//=================test1==================================
func test1() {
	t := Teacher{}
	//teacher虽然没有ShowA方法，但是他的成员有这个方法
	t.ShowA()
	t.ShowB()
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

//=================test2==================================
func test2() {
	a := person{
		Name: "cpwl",
		Age:  21,
		human: human{ //默认用嵌入结构类型名作为字段名
			Sex: 1,
		},
	}
	fmt.Println(a) //{{1} cpwl 21}
	//默认把嵌入结构的字段都给了外层结构，可以直接访问与赋值
	fmt.Println(a.Sex) //1
	//保留用嵌入结构类型名作为字段名的方式，
	//是为了防止多个嵌入结构有相同字段名，发生冲突
	fmt.Println(a.human.Sex) //1
}

type human struct {
	Sex int
}
type person struct {
	human //嵌入结构，只需要写出结构类型
	Name  string
	Age   int
}

//=====================anonymouseTest==============================
func anonymouseTest() {
	chk := new(Sub)
	chk.Flying()
	chk2 := &Sub{Base{"Bob", "Steven", 2.0}, "China"}
	fmt.Println(chk2.Area)
}

type Base struct {
	FirstName, LastName string
	Age                 float32
}

func (base *Base) HasFeet() {
	fmt.Println(base.FirstName + base.LastName + "has feet! Base")
}

func (base *Base) Flying() {
	fmt.Println("Base Can flying!")
}

type Sub struct {
	Base
	Area string
}

func (sub *Sub) Flying() {
	sub.Base.Flying()
	fmt.Println("Sub flying")
}
