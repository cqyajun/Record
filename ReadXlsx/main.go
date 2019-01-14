package main

import (
	"ReadXlsx/control"
	"ReadXlsx/db"
	"ReadXlsx/tools"
	"runtime"

	//_ "ReadXlsx/db"
	"fmt"
	"os"
	"time"
)

//数据库分页查询长度
var databasePageLimt = 5000

func main() {

	maxProcs := runtime.NumCPU()
	runtime.GOMAXPROCS(maxProcs)
	t1 := time.Now().UnixNano() / 1e6
	fmt.Println("===启动程序==")

	has, _ := tools.PathExists(control.TargetDir)
	if has {
		os.RemoveAll(control.TargetDir)
	}
	tools.CreareDirFile(control.TargetDir)
	//xlfile, err := xlsx.OpenFile("conf/demo.xlsx")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//for _, sheet := range xlfile.Sheets {
	//	for _, row := range sheet.Rows {
	//		s_User := new(control.UserDataInfo)
	//
	//		for _index, cell := range row.Cells {
	//			text := cell.String()
	//			if _index == 0 {
	//				s_User.Name = text
	//			} else if _index == 1 {
	//				s_User.Telphone = text
	//			} else if _index == 2 {
	//				s_User.IdCard = text
	//			}
	//		}
	//		//fmt.Println(s_User)
	//		s_User.InitData()
	//		control.AddUserInfo(s_User)
	//	}
	//}
	//control.SaveLessAll()

	totleDataNum := db.PageGetPages(db.DB_Write)

	pages := totleDataNum/int64(databasePageLimt) + 1

	fmt.Println("总数据：", totleDataNum, "总页数:", pages)

	//return

	var totleitems = 0

	for i := 0; i < int(pages); i++ {

		infoList, _ := db.PageGetAll(db.DB_Write, databasePageLimt, i)
		//fmt.Println( "=======>",infoList)
		for _index, info := range *infoList {
			s_User := new(control.UserDataInfo)
			s_User.Name = info.Name
			s_User.Telphone = info.Mobile
			s_User.IdCard = info.Ctfid
			s_User.Email = info.Email
			s_User.Nation = info.Nation
			s_User.Sex = info.Gender
			s_User.Address = info.Address
			isSess := s_User.InitData()
			if isSess {
				control.AddUserInfo(s_User)
				totleitems++
				if _index+1 == len(*infoList) {
					fmt.Println(i, "/", pages, " : ", totleitems)
				}

			}

		}
	}

	control.SaveLessAll()

	control.G_WG.Wait()

	t2 := time.Now().UnixNano() / 1e6
	fmt.Println("over cost ", t2-t1, totleitems)
}

func Gos() {
	for {
		select {
		case <-time.After(time.Second * 1):
			fmt.Println("==============")
		}
	}
}
