package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"bufio"
	"io"
)

type person struct {
	Name string `xml:"nanme,attr"`
	Age  int    `xml:"age"`
}

func main() {
	//ch2_2()
	//ch3_1()
	//ch4_1()
	//ch4_2()
	//ch5_2()
	//ch5_3()
	//ch5_4()
	ch6_3()
}

//字符串基本操作
func ch2_1() {
	s := "hello world"
	//是否包含
	fmt.Println(strings.Contains(s, "hello"), strings.Contains(s, "?"))
	//索引, base()
	fmt.Println(strings.Index(s, "o"))
	ss := "1#2#345"
	// 分割字符串
	splitedss := strings.Split(ss, "#")
	fmt.Println(splitedss)
	//合并字符串
	fmt.Println(strings.Join(splitedss, "#"))
	fmt.Println(strings.HasPrefix(s, "he"), strings.HasSuffix(s, "ld"))
	fmt.Println(strings.Count(ss, "#"))
	fmt.Println(len(ss))
}

//字符串转换
func ch2_2() {
	// int -> string
	fmt.Println(strconv.Itoa(10))
	// int <- string
	fmt.Println(strconv.Atoi("711"))

	fmt.Println(strconv.ParseBool("false"))
	fmt.Println(strconv.ParseFloat("3.14", 32))

	fmt.Println(strconv.FormatBool(true))
	fmt.Println(strconv.FormatInt(123, 8))
	pai := 3.141519
	fmt.Println(strconv.FormatFloat(pai, 'f', 5, 32))
	fmt.Println(strconv.FormatFloat(pai, 'e', 5, 32))
	fmt.Println(strconv.FormatFloat(pai, 'g', 5, 32))

	fmt.Println(strconv.FormatFloat(pai, 'f', -1, 32))
	fmt.Println(strconv.FormatFloat(pai, 'e', -1, 32))
	fmt.Println(strconv.FormatFloat(pai, 'g', -1, 32))

}

//struct对象的序列号和反序列化
func ch3_1() {
	p := person{"davy", 18}
	var data []byte
	var err error
	if data, err = xml.MarshalIndent(p, "", " "); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	p2 := new(person)

	if err = xml.Unmarshal(data, p2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(p2)

}

//命令行解析
func ch4_1() {
	fmt.Println("-------")
	fmt.Println(os.Args)
}

//4-2 使用flag获取复杂参数
func ch4_2() {
	//fmt.Println("----style1")
	//style1()
	fmt.Println("----style2")
	style2()
	fmt.Println("----defaults")
	flag.PrintDefaults()
}
func style1() {
	methodPtr := flag.String("method", "default", "method of sample")
	valuePtr := flag.Int("value", -1, "value of sample")
	flag.Parse()
	fmt.Println(*methodPtr, *valuePtr)
}

func style2() {
	var method string
	var value int
	flag.StringVar(&method, "method", "default", "method of sample")
	flag.IntVar(&value, "value", -1, "value of sample")
	flag.Parse()

	fmt.Println(method, value)
}

// 获取所有节点名称
func ch5_2() {
	content, err := ioutil.ReadFile("WebApplication.csproj")
	if err != nil {
		log.Panic(err)
	}
	decoder := xml.NewDecoder(bytes.NewBuffer(content))

	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		//token 是每个xml的节点
		//根据token的类型做不同的处理
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			fmt.Println(name)
		case xml.EndElement:

		}
	}
}

//获取指定节点
func ch5_3() {
	content, err := ioutil.ReadFile("WebApplication.csproj")
	if err != nil {
		log.Panic(err)
	}
	decoder := xml.NewDecoder(bytes.NewBuffer(content))

	var t xml.Token
	//使用状态机的概念
	//ItemGroup - Compile
	var inItemGroup bool
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		//token 是每个xml的节点
		//根据token的类型做不同的处理
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			if inItemGroup {
				if name == "Compile" {
					fmt.Println(name)
				}
			} else {
				if name == "ItemGroup" {
					inItemGroup = true
				}
			}
		case xml.EndElement:
			if inItemGroup {
				if token.Name.Local == "ItemGroup" {
					inItemGroup = false
				}
			}
		}
	}
}

// 获取节点属性值&思路整理
func ch5_4() {
	content, err := ioutil.ReadFile("WebApplication.csproj")
	if err != nil {
		log.Panic(err)
	}
	decoder := xml.NewDecoder(bytes.NewBuffer(content))

	var t xml.Token
	//使用状态机的概念
	//ItemGroup - Compile
	var inItemGroup bool
	for t, err = decoder.Token(); err == nil; t, err = decoder.Token() {
		//token 是每个xml的节点
		//根据token的类型做不同的处理
		switch token := t.(type) {
		case xml.StartElement:
			name := token.Name.Local
			if inItemGroup {
				if name == "Compile" {
					//fmt.Println(name)
					fmt.Println(getAttributeValue(token.Attr, "Include"))
				}
			} else {
				if name == "ItemGroup" {
					inItemGroup = true
				}
			}
		case xml.EndElement:
			if inItemGroup {
				if token.Name.Local == "ItemGroup" {
					inItemGroup = false
				}
			}
		}
	}
}

func getAttributeValue(attr []xml.Attr, name string) string {
	for _, a := range attr {
		if a.Name.Local == name {
			return a.Value
		}
	}
	return ""
}

//6.1-3 模拟命令行拷贝文件
func ch6_3() {
	var showProgress, force bool
	flag.BoolVar(&force, "f", false, "force copy when existing")
	flag.BoolVar(&showProgress, "v", false, "explain what is being done")

	flag.Parse()

	if flag.NArg() < 2 {
		flag.Usage()
		return
	}

	copyFileAction(flag.Arg(0), flag.Arg(1), showProgress, force)

}

func copyFileAction(src, dst string, showProgress, force bool) {
	if !force {
		if fileExists(dst) {
			fmt.Printf("%s exists override?y/n\n", dst)
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine()

			if strings.TrimSpace(string(data)) != "y" {
				return
			}

		}
	}

	copyFile(src, dst)
	if showProgress {
		fmt.Printf("%s->%s\n", src, dst)
	}
}

func copyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()
	dstFile, err := os.Create(dst)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, srcFile)
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
