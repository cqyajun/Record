package control

import (
	"ReadXlsx/tools"
	"fmt"
)

var TargetDir = "out/"

var FILE_MAX_ROW = 3

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
}

func (user *UserDataInfo) InitData() {

	user.ProviID = string([]byte(user.IdCard)[:3])
	user.DisID = string([]byte(user.IdCard)[:6])
	user.CellLen = 3

	tools.CreareDirFile(fmt.Sprintf("%s/%s", TargetDir, user.ProviID))
	user.SaveFilePath = fmt.Sprintf("%s/%s/%s", TargetDir, user.ProviID, user.DisID)

}

func (user *UserDataInfo) GetAttriVale(_index int) string {
	switch _index {
	case 0:
		return user.Name
	case 1:
		return user.Telphone
	case 2:
		return user.IdCard

	}
	return ""
}

func AddUserInfo(info *UserDataInfo) {

	_xlsx, ok := xlsxMap[info.DisID]
	if !ok {
		newXlsx := new(NewXlsx)
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
