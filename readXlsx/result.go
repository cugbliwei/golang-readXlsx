package readXlsx

import (
	"log"
	"strconv"
	"os"
)
var visited [31005]bool  //标记当前结点是否已经访问过

func dfs(idCart int,result *os.File) {
	result.WriteString(IntToID[idCart] + " ")
	for i := 0; i < len(Matrix[idCart]); i++ {
		if visited[Matrix[idCart][i]] == false {
			visited[Matrix[idCart][i]] = true
			dfs(Matrix[idCart][i],result)
		}
	}
}

func QueryAllID() {
	ConstructMatrix() //得到邻接表
	count := 1  //结果组数
	
	result,err := os.Create("/users/cugbliwei/desktop/result1.dat") //不同机器可修改不同路径
	defer result.Close()
	if err != nil {
		log.Println("创建文件失败！")
		return
	}
	
	result.WriteString("下面的数据为共用device的分组数据，数据为身份证号，对应表格中的 S_ID_CARD  \n\n\n")
	
	for i := 0; i < AllID; i++ {
		if visited[i] == false {
			result.WriteString("第 " + strconv.Itoa(count) + " 组数据：\n")
			visited[i] = true
			dfs(i,result)
			result.WriteString("\n\n")
			count += 1
		}
	}
}