package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/miekg/dns"
	"github.com/spf13/viper"
)

var (
	configFile string
	configer   *viper.Viper
)

func rootdig() {
	// 读取当前运行环境的 /etc/resolv.conf，获得 name server 的配置
	// config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")

	// 构造发起 DNS 请求的客户端
	dnsClient := new(dns.Client)

	// 构造 DNS 报文
	m := new(dns.Msg)

	// 设置问题字段，即查询命令行参数第一个参数的 A 记录
	m.SetQuestion(dns.Fqdn("5251255"), dns.TypeA)
	m.RecursionDesired = true

	// fmt.Println(m)

	// client 发起 DNS 请求，其中 c 为上文创建的 client，m 为构造的 DNS 报文
	// config 为从 /etc/resolv.conf 构造出来的配置
	hostPort := net.JoinHostPort("8.8.8.58", "53")
	r, rtt, err := dnsClient.Exchange(m, hostPort)
	if r == nil {
		log.Fatalf("*** 错误: %s\n", err.Error())
	}
	fmt.Println(r.Rcode)
	fmt.Println(r)
	nanoseconds := float32(rtt.Nanoseconds())
	fmt.Printf("%.2fus\n", nanoseconds/1000)          // us
	fmt.Printf("%.2fms\n", nanoseconds/1000/1000)     //ms
	fmt.Printf("%.2fs\n", nanoseconds/1000/1000/1000) //s
	if r.Rcode != dns.RcodeSuccess {
		log.Fatalf("*** invalid answer name after MX query\n")
	}

	// 如果 DNS 查询成功
	for _, a := range r.Answer {
		fmt.Printf("%v\n", a)
	}

}

func init() {
	flag.StringVar(&configFile, "config", "", "config file")
}

func main() {
	flag.Parse()
	configer = Configer(configFile)
	fmt.Println(configer.Get("table"))
	rootdig()
}
