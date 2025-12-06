package SO_Form

import (
	SO_Class "SOApp_GO/class"

	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type tblelf struct {
	// table struct {
	// 	TENOMRIY string `tipe:"1" json:"No" key:"true" required:"true"` //required:"true" max:"20"
	// 	TEERMS   string `tipe:"2" json:"Error Message" required:"true"`
	// 	TESTMT   string `tipe:"2" json:"Error Sintax" required:"true"`
	// 	TEREMK   string `tipe:"2" json:"Remark"`
	// 	TEUSRM   string `tipe:"2" json:"User Remark"`
	// 	TECSDT   string `tipe:"X" `
	// }
}

func (saya tblelf) Save(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLELF-Save()")
	hasil.Sukses = false
	hasil.Pesan = "No Action for this Method!"
	hasil.Data = ""
	hasil.Kode = ""
	return hasil
}

func (saya tblelf) LoadGrid(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLELF-LoadGrid()")
	if c.Param("sort") == "" {
		param := gin.Param{
			Key: "sort",
			Value: SO_Class.Fmt.Sprint(
				`[{"property": "tenomriy","direction": "desc"}]`,
			),
		}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select 
				tenomriy, teerms, testmt,
				teremk, teusrm
				`, Form.GetDefaultField("te"), ` 
			from tblelf 
			where 1 = 1
		`)
	SO_Class.Log.Println(false, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "tenomriy",
		columns: []Kolum{
			{"tenomriy", KolumProperty{name: "No"}},
			{"teerms", KolumProperty{name: "Error Message"}},
			{"testmt", KolumProperty{name: "Error Sintax"}},
			{"teremk", KolumProperty{name: "Remark"}},
			{"teusrm", KolumProperty{name: "User Remark"}},
		},
		defaultField: true,
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}
func (saya tblelf) FillForm(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLELF-FillForm()")
	key := c.Param("tenomriy")
	sqlstm := SO_Class.Fmt.Sprint("select * from tblelf where tenomriy = '", key, "'")
	SO_Class.Log.Println(true, sqlstm)
	hasil = Form.GetRs(c, sqlstm)
	return hasil
}

func (saya tblelf) LoadFormObject(c *gin.Context) (hasil SO_Class.Hasil) {
	hasil.Sukses = false
	hasil.Pesan = ""
	hasil.Data = nil

	frmId := c.Param("FrmId")
	menuId := SO_Class.Strings.ToUpper("tblelf")

	result := FormObject{
		Prefix: "te",
		Frame1: Form.CrtObj(ObjGrd{
			Id: "Grid1", FrmId: frmId, MenuId: menuId,
			Title:      "Table Error Logfile",
			Controller: SO_Class.Strings.ToUpper("tblelf"),
			Method:     "LoadGrid",
		}),
		Frame2: Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
			Id: "Panel1",
			Items: []interface{}{
				Form.CrtObj(ObjTxt{Mode: "2", FrmId: frmId, MenuId: menuId,
					Id: "tenomriy", Name: "No", AllowBlank: false, Width: 100,
				}),
				Form.CrtObj(ObjTxt{Mode: "2", FrmId: frmId, MenuId: menuId,
					Id: "teerms", Name: "Error Message", AllowBlank: false, Width: 600,
				}),
				Form.CrtObj(ObjRmk{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "testmt", Name: "Error Sintax", AllowBlank: false, Width: 600, Height: 400,
				}),
				Form.CrtObj(ObjRmk{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "teremk", Name: "Remark",
				}),
				Form.CrtObj(ObjRmk{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "teusrm", Name: "User Remark",
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
	Form.Add("TBLELF", TBLELF)
	if Form.logPrintInitFlag {
		SO_Class.Log.Println(true, "Masuk form-tblelf-init()")
	}
}

// Exported instance
var TBLELF tblelf
