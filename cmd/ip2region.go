package main

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
	"path/filepath"
)

func main() {
	var ip string
	fmt.Println("请输入需要识别IP：")
	fmt.Scanln(&ip)
	fmt.Println("IP：" + ip + "识别中。。。。")
	path, _ := filepath.Abs("storage/files/ip2region.db")
	region, err := ip2region.New(path)
	defer region.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	ipInfo, err := region.MemorySearch(ip)
	fmt.Println("memory算法", ipInfo, err)
	ipInfo, err = region.BinarySearch(ip)
	fmt.Println("binary算法二分查找", ipInfo, err)
	ipInfo, err = region.BtreeSearch(ip)
	fmt.Println("b-tree算法", ipInfo, err)
}
