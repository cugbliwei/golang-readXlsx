package readXlsx

import (
	"testing"
	"github.com/tealeg/xlsx"
	"connected/connectedComponent"
)

func ReadXlsx_test(t *testing.T) {
	xlFile, err = xlsx.OpenFile("/usr/local/connect.xlsx")
	if err != nil {
		t.Error(err)
	}
	i := 0
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			j := 0
			for _, cell := range row.Cells {
				j += 1
				if j == 4{
					fmt.Printf("device: %s  ", cell.String())
				}
				if j == 15 {
					fmt.Printf("ID_CART: %s\n", cell.String())
				}
			}
			i += 1
			if i > 10 {
				break
			}
		}
	}
}

func deviceIDCard_test(t *testing.T) {
	deviceIDCard := connectedComponent.ReadXlsx("/usr/local/connect.xlsx")
	t.Log()
	for device, IDCard := range deviceIDCard {
		fmt.Printf("device: %s  ",device)
		for ID,_ := range IDCard {
			fmt.Printf("IDCard: %s  ",ID)
		}
		fmt.Println("++++")
	}
}

func ReadXlsx(fileName string) (deviceIDCard map[string] (map[string] byte)) {//key是device，value的map的string是IDCard
	xlFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		log.Println("打开文件错误！")
		return nil
	}
	deviceIDCard = make(map[string] (map[string] byte))
	
	outputFile,errOut := os.Create("/users/cugbliwei/desktop/original_deviceIDCard.dat")  //把device和IDCard数据写入文件
	defer outputFile.Close()
	if errOut != nil {
		log.Println("创建文件失败！")
		return nil
	}
	
	for _, sheet := range xlFile.Sheets {
		i := 0
		for _, row := range sheet.Rows {
			i += 1
			if i == 1 {  //第一行是表头，不读取
				continue
			}
			
			j := 0
			device,IDCard := "",""
			for _, cell := range row.Cells {
				j += 1
				if j == 4{
					device = cell.String()
				}
				if j == 15 {
					IDCard = cell.String()
				}
			}
			outputFile.WriteString("device: " + device + "   IDCard: " + IDCard + "\n")
			
			if ID,ok := deviceIDCard[device]; ok {
				ID[IDCard]=byte(1)
				deviceIDCard[device]=ID
			} else {
				IDC := make(map[string] byte)
				IDC[IDCard]= byte(1)
				deviceIDCard[device]=IDC
			}
//			log.Printf("device: %s   IDcard: %d\n",device,IDCard)
		}
	}
	return deviceIDCard
}