package readXlsx

import (
	"fmt"
)

var Matrix [31005][]int  //构建邻接表
var IDToInt map[string] int //离散化身份证为数值，即身份证对应一个数值
var IntToID map[int] string //用数值重新取回身份证号
var AllID int  //身份证号总数

func ConstructMatrix() {
	deviceIDCard := ReadXlsx("/usr/local/connect.xlsx")
	IDToInt = make(map[string] int)
	IntToID = make(map[int] string)
	AllID = 0
	
	for _, idCard := range deviceIDCard {
		if len(idCard) == 1 {  //如果考虑一个点的情况，则注释掉
			continue
		}
		idSlice := make([]int,0)
		for id, _ := range idCard {
			if _,ok := IDToInt[id]; ok == false {
				IDToInt[id] = AllID
				IntToID[AllID] = id
				AllID += 1   //测试得出AllID = 4925
			}
			idSlice = append(idSlice,IDToInt[id])
		}
//		fmt.Println("len(idSlice): ",len(idSlice))
		for j := 0; j < len(idSlice); j++ {
			for k := j + 1; k <len(idSlice); k++ {
				Matrix[idSlice[j]] = append(Matrix[idSlice[j]],idSlice[k])//构建邻接表
				Matrix[idSlice[k]] = append(Matrix[idSlice[k]],idSlice[j])
			}
		}
//		if len(idSlice) == 1 {  //如果考虑一个点的情况则把注释去掉
//			Matrix[idSlice[0]] = append(Matrix[idSlice[0]],idSlice[0])//构建邻接表
//		}
	}
	fmt.Println("AllID: ",AllID)
	
}