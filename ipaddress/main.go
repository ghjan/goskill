package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

// go (golang) DNS域名解析实现
// https://blog.csdn.net/mumumuwudi/article/details/48200505

//Go语言获取外网和本地IP
//https://studygolang.com/articles/4465
var get_ip = flag.String("get_ip", "", "external|internal|all")
var domain = flag.String("domain", "", "www.davidzhang.xin")

func main() {
	ip, err := getIP()
	fmt.Println("getIP:")
	processResult([]string{ip}, err)

	good := false
	flag.Parse()
	if *get_ip == "external" {
		fmt.Println("get_external ips:")
		ns, err := get_external()
		processResult([]string{ns}, err)
		good = true
	} else if *get_ip == "internal" {
		fmt.Println("get_internal ips:")
		ips, err := get_internal()
		processResult(ips, err)
		good = true
	} else if *get_ip == "all" {
		fmt.Println("get_external ips:")
		ns, err := get_external()
		processResult([]string{ns}, err)

		fmt.Println("get_internal ips:")
		ips, err := get_internal()
		processResult(ips, err)
		good = true
	}
	if strings.Count(*domain, "") > 1 {
		fmt.Printf("lookupHostIP ips for %s\n", *domain)
		ips, err := lookupHostIP(*domain)
		processResult(ips, err)
		good = true
	}
	if !good {
		fmt.Println("Usage of ./ipaddress --get_ip=(external|internal|all) --domain=www.davidzhang.xin")
		os.Exit(1)
	}

}
func processResult(results []string, err error) {
	if err != nil {
		//fmt.Println(err)
		log.Panic(err)
	}
	for _, content := range results {
		fmt.Fprintf(os.Stdout, "--%s\n", content)
	}
}

func getIP() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
	}
	ip := ""
	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}

		}
	}
	return ip, err
}

//获取外部ip地址
func get_external() (string, error) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		return "", err
	}
	defer resp.Body.Close()
	//io.Copy(os.Stdout, resp.Body)
	bb, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		return string(bb[:]), nil
	} else {
		return "", err
	}
}

//获取内部ip地址
func get_internal() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		return ips, err
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}
	return ips, err
}

func lookupHostIP(domainName string) ([]string, error) {
	ns, err := net.LookupHost(domainName)
	return ns, err
}
