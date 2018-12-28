package main

import (
	"ReadXlsx/control"
	"ReadXlsx/tools"

	//_ "ReadXlsx/db"
	"fmt"
	"github.com/tealeg/xlsx"
	"os"
	"time"
)

func main() {

	t1 := time.Now().UnixNano() / 1e6
	fmt.Println("===启动程序==")

	has, _ := tools.PathExists(control.TargetDir)
	if has {
		os.RemoveAll(control.TargetDir)
	}
	tools.CreareDirFile(control.TargetDir)
	xlfile, err := xlsx.OpenFile("conf/demo.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, sheet := range xlfile.Sheets {
		for _, row := range sheet.Rows {
			s_User := new(control.UserDataInfo)

			for _index, cell := range row.Cells {
				text := cell.String()
				if _index == 0 {
					s_User.Name = text
				} else if _index == 1 {
					s_User.Telphone = text
				} else if _index == 2 {
					s_User.IdCard = text
				}
			}
			//fmt.Println(s_User)
			s_User.InitData()
			control.AddUserInfo(s_User)
		}
	}
	control.SaveLessAll()
	t2 := time.Now().UnixNano() / 1e6
	fmt.Println("over cost ", t2-t1)

}
