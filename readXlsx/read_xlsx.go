package readXlsx

import (
	"log"
	"github.com/tealeg/xlsx"
)

func ReadXlsx(fileName string) (deviceIDCard map[string] (map[string] byte)) {//key是device，value的map的string是IDCard
	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		log.Println("打开文件错误！")
		return nil
	}
	deviceIDCard = make(map[string] (map[string] byte))
	
	for _, sheet := range xlFile.Sheets {
		i := 0
		for _, row := range sheet.Rows {
			i += 1
			if i == 1 {  //第一行是表头，不读取
				continue
			}
			
			j := 0
			device,idCard := "",""
			for _, cell := range row.Cells {
				j += 1
				if j == 4{  //第四列是device
					device = cell.String()
				}
				if j == 15 { //第15列是ID_CARD
					idCard = cell.String()
				}
			}
			
			//过滤device和ID_CARD
			if id,ok := deviceIDCard[device]; ok { //如果已经存在device,则直接加入map
				id[idCard]=byte(1)
				deviceIDCard[device]=id
			} else {                              //如果未存在device,则创建一个value的map再加入
				idc := make(map[string] byte)
				idc[idCard]= byte(1)
				deviceIDCard[device]=idc
			}
//			log.Printf("device: %s   IDcard: %s\n",device,IDCard)
		}
	}
	return deviceIDCard
}