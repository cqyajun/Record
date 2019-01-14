package control

import (
	//"github.com/Luxurioust/excelize"

	"ReadXlsx/tools"
	"fmt"
	"github.com/tealeg/xlsx"
	_ "github.com/tealeg/xlsx"
)

type NewXlsx struct {
	file       *xlsx.File
	sheet      *xlsx.Sheet
	path       string
	fatherPath string
	NIndex     int
	Key        string
}

func (f *NewXlsx) Init(key string) {
	var err error
	f.file = xlsx.NewFile()
	f.NIndex = 1
	f.Key = key
	f.path = ""
	f.sheet, err = f.file.AddSheet("Sheet1")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (f *NewXlsx) AddToUserInfo(info *UserDataInfo) {

	f.path = info.SaveFilePath
	f.fatherPath = info.FileDicPath
	if len(f.sheet.Rows) < FILE_MAX_ROW {
		row := f.sheet.AddRow()
		for i := 0; i < info.CellLen; i++ {
			cell := row.AddCell()
			cell.SetValue(info.GetAttriVale(i))
		}
		if len(f.sheet.Rows) == FILE_MAX_ROW { //保存文件

			createNew(f)
			G_WG.Add(1)
			go gofuncSave(f)
		}

	}
}

func gofuncSave(f *NewXlsx) {
	defer G_WG.Done()
	tools.CreareDirFile(f.fatherPath)
	err := f.file.Save(fmt.Sprintf("%s-%d.xlsx", f.path, f.NIndex))
	if err != nil {
		fmt.Printf(err.Error())
	}
}
func (f *NewXlsx) SaveLessAll() {
	if len(f.sheet.Rows) > 100 {

		G_WG.Add(1)
		go gofuncSave(f)
	}

}

func createNew(f *NewXlsx) {

	newxlsx := new(NewXlsx) // NewXlsx{}
	newxlsx.Init(f.Key)
	newxlsx.path = f.path
	newxlsx.fatherPath = f.fatherPath
	newxlsx.NIndex = f.NIndex + 1
	xlsxMap[newxlsx.Key] = newxlsx

}
