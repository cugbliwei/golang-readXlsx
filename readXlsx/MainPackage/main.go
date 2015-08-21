package main

import (
	"readXlsx"
)

func main() {
	//读取表格数据比较费时，所以需静待1分钟左右方可运行完毕
	readXlsx.QueryAllID() //位于result.go
}