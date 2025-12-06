package SO_Form

import (
	SO_Class "SOApp_GO/class"

	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type tbldsc struct {
	table struct {
		TDDSCD string `tipe:"1" json:"Code" key:"true" required:"true" max:"20"` //required:"true" max:"20"
		TDDSNM string `tipe:"2" json:"Name" required:"true"`
		TDLGTH string `tipe:"2" json:"Char Length" required:"true"`
		TDDPFG string `tipe:"2" json:"Display Flag"`
		TDSYFG string `tipe:"2" json:"System Flag"`
		TDREMK string `tipe:"2" json:"Remark"`
		TDUSRM string `tipe:"2" json:"User Remark"`
		TDCSDT string `tipe:"X" `
	}
}

func (saya tbldsc) Save(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLDSC-Save()")
	hasil.Sukses = false
	hasil.Pesan = "No Action for this Method!"
	hasil.Data = ""
	hasil.Kode = ""

	err := SO_Class.Fungsi.BindParamsToStruct(c, &saya.table)
	if err != nil {
		hasil.Sukses = false
		hasil.Pesan = err.Error()
		return
	}

	exec := Form.ExecQueryMultiple(c, c.Param("username"),
		func(tx Transaction) TransactionResult {
			var tr TransactionResult
			tr.Error = nil
			tr.Sintax = ""

			// var iud ParamIUD
			// iud.TableName = "tbldsc"
			// iud.UserName = c.Param("username")
			// iud.Source = c.Param("source")
			// iud.Mode = c.Param("Mode")
			// iud.StructAnda = saya.table

			// _, errBFCS := Form.CheckRecord_BFCS(ParamBFCS{
			// 	Tx:       tx,
			// 	Mode:     c.Param("Mode"),
			// 	Table:    "tbldsc",
			// 	KeyField: "tddscd",
			// 	KeyValue: c.Param("tddscd"),
			// 	CSDT:     c.Param("tdcsdt"),
			// })
			// if errBFCS != nil {
			// 	tr.Error = errBFCS
			// 	return tr
			// }

			// sqlstm := Form.GetSintaxSQL_IUD(ParamIUD{
			// 	TableName:  "tbldsc",
			// 	UserName:   c.Param("username"),
			// 	Source:     c.Param("source"),
			// 	Mode:       c.Param("Mode"),
			// 	StructAnda: saya.table,
			// })
			// SO_Class.Log.Println(false, " tbldsc save ", sqlstm)

			// _, err := Form.Execute(tx, &tr, userName, sqlstm)
			// if err != nil {
			// 	tr.Error = err
			// 	return tr
			// }

			tr = saya.StpSave(
				tx,
				c.Param("Mode"),
				c.Param("username"),
				c.Param("source"),
				c.Param("tdcsdt"),
				saya)

			return tr
		},
	)

	if exec != nil {
		hasil.Sukses = false
		hasil.Pesan = SO_Class.Fmt.Sprint("", exec.Error())
	} else {
		hasil.Sukses = true
		hasil.Pesan = SO_Class.Fmt.Sprint("Sukses ", "")
	}

	hasil.Data = ""

	return hasil
}

func (saya tbldsc) StpSave(tx Transaction,
	mode string, userName string, source string, csdt string, table tbldsc,
) (tr TransactionResult) {

	data := table.table
	tr.Sintax = " -- tbldsc.StpSave() "

	_, errBFCS := Form.CheckRecord_BFCS(ParamBFCS{
		Tx:       tx,
		Mode:     mode,
		Table:    "tbldsc",
		KeyField: "tddscd",
		KeyValue: data.TDDSCD,
		CSDT:     csdt,
	})
	if errBFCS != nil {
		tr.Error = errBFCS
		return tr
	}

	sqlstm := Form.GetSintaxSQL_IUD(ParamIUD{
		TableName:  "tbldsc",
		UserName:   userName,
		Source:     source,
		Mode:       mode,
		StructAnda: data,
	})
	SO_Class.Log.Println(false, " tbldsc save ", sqlstm)

	_, err := Form.Execute(tx, &tr, userName, sqlstm)
	if err != nil {
		tr.Error = err
		return tr
	}

	return tr
}

func (saya tbldsc) LoadGrid(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLDSC-LoadGrid()")
	if c.Param("sort") == "" {
		param := gin.Param{
			Key: "sort",
			Value: SO_Class.Fmt.Sprint(
				`[{"property": "tddscd","direction": "asc"}]`,
			),
		}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select 
				tddscd, tddsnm, tdlgth, 
				tdsyfg, s.tssynm tdsyfg_desc, 
				tddpfg, d.tssynm tddpfg_desc,
				tdremk, tdusrm
				`, Form.GetDefaultField("td"), ` 
			from tbldsc 
			left join tblsys s on s.tsdscd = 'YESNO' and s.tssycd = tdsyfg::varchar
			left join tblsys d on d.tsdscd = 'YESNO' and d.tssycd = tddpfg
			where 1 = 1
		`)
	SO_Class.Log.Println(false, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "tddscd",
		columns: []Kolum{
			{"tddscd", KolumProperty{name: "Code"}},
			{"tddsnm", KolumProperty{name: "Description"}},
			{"tdlgth", KolumProperty{name: "Length"}},
			{"tdsyfg_desc", KolumProperty{name: "System Flag"}},
			{"tddpfg_desc", KolumProperty{name: "Display Flag"}},
			{"tdremk", KolumProperty{name: "Remark"}},
			{"tdusrm", KolumProperty{name: "User Remark"}},
		},
		defaultField: true,
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}
func (saya tbldsc) FillForm(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLDSC-FillForm()")
	key := c.Param("tddscd")
	sqlstm := SO_Class.Fmt.Sprint("select * from tbldsc where tddscd = '", key, "'")
	SO_Class.Log.Println(true, sqlstm)
	hasil = Form.GetRs(c, sqlstm)
	return hasil
}

func (saya tbldsc) LoadFormObject(c *gin.Context) (hasil SO_Class.Hasil) {
	hasil.Sukses = false
	hasil.Pesan = ""
	hasil.Data = nil

	frmId := c.Param("FrmId")
	menuId := SO_Class.Strings.ToUpper("tbldsc")

	result := FormObject{
		Prefix: "td",
		Frame1: Form.CrtObj(ObjGrd{
			Id: "Grid1", FrmId: frmId, MenuId: menuId,
			Title:      "Master Table Description",
			Controller: SO_Class.Strings.ToUpper("tbldsc"),
			Method:     "LoadGrid",
		}),
		Frame2: Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
			Id: "Panel1",
			Items: []interface{}{
				Form.CrtObj(ObjTxt{Mode: "2", FrmId: frmId, MenuId: menuId,
					Id: "tddscd", Name: "Code", AllowBlank: false, Width: 100,
				}),
				Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "tddsnm", Name: "Description", AllowBlank: false, Width: 300,
				}),
				Form.CrtObj(ObjNum{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "tdlgth", Name: "Length", Width: 50,
				}),
				Form.CrtObj(ObjCmb{C: c, Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "tdsyfg", Name: "System Flag",
					Table: "YESNO", Value: "1", Width: 100,
				}),
				Form.CrtObj(ObjCmb{C: c, Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "tddpfg", Name: "Display Flag",
					Table: "YESNO", Value: "1", Width: 100,
				}),
				Form.CrtObj(ObjRmk{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "tdremk", Name: "Remark",
				}),
				Form.CrtObj(ObjRmk{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "tdusrm", Name: "User Remark",
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
	Form.Add("TBLDSC", TBLDSC)
	if Form.logPrintInitFlag {
		SO_Class.Log.Println(true, "Masuk form-tbldsc-init()")
	}
}

// Exported instance
var TBLDSC tbldsc
