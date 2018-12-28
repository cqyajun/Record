package control

import (
	//"github.com/Luxurioust/excelize"

	"fmt"
	"github.com/tealeg/xlsx"
	_ "github.com/tealeg/xlsx"
)

type NewXlsx struct {
	file   *xlsx.File
	sheet  *xlsx.Sheet
	path   string
	NIndex int
	Key    string
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
	if len(f.sheet.Rows) < FILE_MAX_ROW {
		row := f.sheet.AddRow()
		for i := 0; i < info.CellLen; i++ {
			cell := row.AddCell()
			cell.SetValue(info.GetAttriVale(i))
		}
		if len(f.sheet.Rows) == FILE_MAX_ROW { //保存文件

			err := f.file.Save(fmt.Sprintf("%s-%d.xlsx", info.SaveFilePath, f.NIndex))
			if err != nil {
				fmt.Printf(err.Error())
			}
			createNew(f)
		}

	}
}

func (f *NewXlsx) SaveLessAll() {
	if len(f.sheet.Rows) > 0 {
		fmt.Println("SaveLessAll", f.path, f.NIndex, len(f.sheet.Rows))
		err := f.file.Save(fmt.Sprintf("%s-%d.xlsx", f.path, f.NIndex))
		if err != nil {
			fmt.Printf(err.Error())
		}
	}

}

func createNew(f *NewXlsx) {

	newxlsx := new(NewXlsx) // NewXlsx{}
	newxlsx.Init(f.Key)
	newxlsx.path = f.path
	newxlsx.NIndex = f.NIndex + 1
	xlsxMap[newxlsx.Key] = newxlsx

}
