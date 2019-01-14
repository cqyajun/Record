package control

import (
	"fmt"
	"strconv"
	"sync"
)

var TargetDir = "out/"

var G_WG sync.WaitGroup

var FILE_MAX_ROW int = 1000

var xlsxMap map[string]*NewXlsx

func init() {
	xlsxMap = make(map[string]*NewXlsx)
}

type UserDataInfo struct {
	Name         string
	Telphone     string
	IdCard       string
	ProviID      string //省份id
	DisID        string //  区域id
	CellLen      int
	SaveFilePath string
	FileDicPath  string
	Sex          string
	Address      string
	Email        string
	Nation       string
}

func (user *UserDataInfo) InitData() bool {
	//if(len(user.IdCard) != 18 && len(user.IdCard) != 15){
	//	fmt.Println(user.IdCard)
	//	return false
	//}
	user.ProviID = string([]byte(user.IdCard)[:3])
	nProviID, err := strconv.Atoi(user.ProviID)
	if err != nil || nProviID < 110 || nProviID > 820 {
		//fmt.Println("err=====>> ",user.ProviID)
		return false
	}

	user.DisID = string([]byte(user.IdCard)[:6])
	_, err = strconv.Atoi(user.DisID)
	if err != nil {
		//fmt.Println("err=====>> ",user.ProviID)
		return false
	}
	user.CellLen = 7

	//tools.CreareDirFile(fmt.Sprintf("%s/%s", TargetDir, user.ProviID))
	user.FileDicPath = fmt.Sprintf("%s/%s", TargetDir, user.ProviID)
	user.SaveFilePath = fmt.Sprintf("%s/%s/%s", TargetDir, user.ProviID, user.DisID)

	return true
}

func (user *UserDataInfo) GetAttriVale(_index int) string {
	switch _index {
	case 0:
		return user.Name
	case 1:
		if user.Sex == "F" {
			return "女"
		} else if user.Sex == "M" {
			return "男"
		}
	case 2:
		return user.IdCard
	case 3:
		return user.Telphone
	case 4:
		return user.Address
	case 5:
		return user.Email
	case 6:
		return user.Nation

	}
	return ""
}

func AddUserInfo(info *UserDataInfo) {

	_xlsx, ok := xlsxMap[info.DisID]
	if !ok {
		newXlsx := new(NewXlsx) //new(NewXlsx)  //  NewXlsx{}
		newXlsx.Init(info.DisID)
		xlsxMap[info.DisID] = newXlsx
		newXlsx.AddToUserInfo(info)
	} else {
		_xlsx.AddToUserInfo(info)
	}

}

func SaveLessAll() {
	for _, xlsx := range xlsxMap {
		xlsx.SaveLessAll()
	}
}
