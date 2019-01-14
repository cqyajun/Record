package db

type TbInfo struct {
	Name      string `xorm:"VARCHAR(10)"`
	Cardno    string `xorm:"VARCHAR(20)"`
	Descriot  string `xorm:"VARCHAR(20)"`
	Ctftp     string `xorm:"VARCHAR(20)"`
	Ctfid     string `xorm:"VARCHAR(50)"`
	Gender    string `xorm:"VARCHAR(2)"`
	Birthday  string `xorm:"VARCHAR(20)"`
	Address   string `xorm:"VARCHAR(200)"`
	Zip       string `xorm:"VARCHAR(10)"`
	Dirty     string `xorm:"VARCHAR(50)"`
	District1 string `xorm:"VARCHAR(10)"`
	District2 string `xorm:"VARCHAR(10)"`
	District3 string `xorm:"VARCHAR(10)"`
	District4 string `xorm:"VARCHAR(10)"`
	District5 string `xorm:"VARCHAR(10)"`
	District6 string `xorm:"VARCHAR(10)"`
	Firstnm   string `xorm:"VARCHAR(20)"`
	Lastnm    string `xorm:"VARCHAR(30)"`
	Duty      string `xorm:"VARCHAR(20)"`
	Mobile    string `xorm:"VARCHAR(30)"`
	Tel       string `xorm:"VARCHAR(30)"`
	Fax       string `xorm:"VARCHAR(30)"`
	Email     string `xorm:"VARCHAR(100)"`
	Nation    string `xorm:"VARCHAR(50)"`
	Taste     string `xorm:"VARCHAR(50)"`
	Education string `xorm:"VARCHAR(50)"`
	Company   string `xorm:"VARCHAR(50)"`
	Ctel      string `xorm:"VARCHAR(100)"`
	Caddress  string `xorm:"VARCHAR(200)"`
	Czip      string `xorm:"VARCHAR(20)"`
	Family    string `xorm:"VARCHAR(50)"`
	Version   string `xorm:"VARCHAR(100)"`
	Id        int64  `xorm:"pk BIGINT(20)"`
}
