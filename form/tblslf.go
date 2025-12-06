package SO_Form

import (
	SO_Class "SOApp_GO/class"

	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type tblslf struct {
	// table struct {
	// 	TQNOMRIY string `tipe:"1" json:"No" key:"true" required:"true"` //required:"true" max:"20"
	// 	TQSTMT   string `tipe:"2" json:"Sintax" required:"true"`
	// 	TQREMK   string `tipe:"2" json:"Remark"`
	// 	TQUSRM   string `tipe:"2" json:"User Remark"`
	// 	TQCSDT   string `tipe:"X" `
	// }
}

func (saya tblslf) Save(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLSLF-Save()")
	hasil.Sukses = false
	hasil.Pesan = "No Action for this Method!"
	hasil.Data = ""
	hasil.Kode = ""
	return hasil
}

func (saya tblslf) LoadGrid(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLSLF-LoadGrid()")
	if c.Param("sort") == "" {
		param := gin.Param{
			Key: "sort",
			Value: SO_Class.Fmt.Sprint(
				`[{"property": "tqnomriy","direction": "desc"}]`,
			),
		}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select 
				tqnomriy, tqstmt,
				tqremk, tqusrm
				`, Form.GetDefaultField("tq"), ` 
			from tblslf 
			where 1 = 1
		`)
	SO_Class.Log.Println(false, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "tqnomriy",
		columns: []Kolum{
			{"tqnomriy", KolumProperty{name: "No"}},
			{"tqstmt", KolumProperty{name: "Sintax"}},
			{"tqremk", KolumProperty{name: "Remark"}},
			{"tqusrm", KolumProperty{name: "User Remark"}},
		},
		defaultField: true,
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}
func (saya tblslf) FillForm(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLSLF-FillForm()")
	key := c.Param("tqnomriy")
	sqlstm := SO_Class.Fmt.Sprint("select * from tblslf where tqnomriy = '", key, "'")
	SO_Class.Log.Println(true, sqlstm)
	hasil = Form.GetRs(c, sqlstm)
	return hasil
}

func (saya tblslf) LoadFormObject(c *gin.Context) (hasil SO_Class.Hasil) {
	hasil.Sukses = false
	hasil.Pesan = ""
	hasil.Data = nil

	frmId := c.Param("FrmId")
	menuId := SO_Class.Strings.ToUpper("tblslf")

	result := FormObject{
		Prefix: "tq",
		Frame1: Form.CrtObj(ObjGrd{
			Id: "Grid1", FrmId: frmId, MenuId: menuId,
			Title:      "Table Syntax Logfile",
			Controller: SO_Class.Strings.ToUpper("tblslf"),
			Method:     "LoadGrid",
		}),
		Frame2: Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
			Id: "Panel1",
			Items: []interface{}{
				Form.CrtObj(ObjTxt{Mode: "2", FrmId: frmId, MenuId: menuId,
					Id: "tqnomriy", Name: "No", AllowBlank: false, Width: 100,
				}),
				Form.CrtObj(ObjRmk{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "tqstmt", Name: "Sintax", AllowBlank: false, Width: 600, Height: 400,
				}),
				Form.CrtObj(ObjRmk{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "tqremk", Name: "Remark",
				}),
				Form.CrtObj(ObjRmk{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "tqusrm", Name: "User Remark",
				}),
			},
		}),
	}

	hasil.Sukses = true
	hasil.Pesan = "sukses"
	hasil.Data = result
	return hasil
}

func init() {
	Form.Add("TBLSLF", TBLSLF)
	if Form.logPrintInitFlag {
		SO_Class.Log.Println(true, "Masuk form-tblslf-init()")
	}
}

// Exported instance
var TBLSLF tblslf
