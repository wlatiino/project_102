package SO_Form

import (
	SO_Class "SOApp_GO/class"

	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type tblprm struct {
	table struct {
		TRPRCD string `tipe:"1" json:"Code" key:"true" required:"true" max:"20"` //required:"true" max:"20"
		TRPRNM string `tipe:"2" json:"Name" required:"true"`
		TRSYT1 string `tipe:"2" json:"Text 1"`
		TRSYT2 string `tipe:"2" json:"Text 2"`
		TRSYT3 string `tipe:"2" json:"Text 3"`
		TRSYV1 string `tipe:"2" json:"Value 1"`
		TRSYV2 string `tipe:"2" json:"Value 2"`
		TRSYV3 string `tipe:"2" json:"Value 3"`
		TRDPFG string `tipe:"2" json:"Display Flag"`
		TRREMK string `tipe:"2" json:"Remark"`
		TRUSRM string `tipe:"2" json:"User Remark"`
		TRCSDT string `tipe:"X" `
	}
}

func (saya tblprm) Save(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.CetakKunci(true, c, "Masuk TBLPRM-Save()")
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

			tr = saya.StpSave(
				tx,
				c.Param("Mode"),
				c.Param("username"),
				c.Param("source"),
				c.Param("trcsdt"),
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

func (saya tblprm) StpSave(tx Transaction,
	mode string, userName string, source string, csdt string, table tblprm,
) (tr TransactionResult) {

	data := table.table
	tr.Sintax = " -- tblprm.StpSave() "

	_, errBFCS := Form.CheckRecord_BFCS(ParamBFCS{
		Tx:       tx,
		Mode:     mode,
		Table:    "tblprm",
		KeyField: "trprcd",
		KeyValue: data.TRPRCD,
		CSDT:     csdt,
	})
	if errBFCS != nil {
		tr.Error = errBFCS
		return tr
	}

	sqlstm := Form.GetSintaxSQL_IUD(ParamIUD{
		TableName:  "tblprm",
		UserName:   userName,
		Source:     source,
		Mode:       mode,
		StructAnda: data,
	})
	SO_Class.Log.Println(false, " tblprm save ", sqlstm)

	_, err := Form.Execute(tx, &tr, userName, sqlstm)
	if err != nil {
		tr.Error = err
		return tr
	}

	return tr
}

func (saya tblprm) LoadGrid(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.CetakKunci(false, c, "Masuk TBLPRM-LoadGrid()")
	if c.Param("sort") == "" {
		param := gin.Param{
			Key: "sort",
			Value: SO_Class.Fmt.Sprint(
				`[{"property": "trprcd","direction": "asc"}]`,
			),
		}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select 
				trprcd, trprnm, 
				trsyv1::dec(24,0) trsyv1, 
				trsyv2::dec(24,2) trsyv2,
				trsyv3::dec(24,4) trsyv3, 
				synm trdpfg_desc,
				trremk, trusrm
				`, Form.GetDefaultField("tr"), ` 
			from tblprm 			
			left join (
				select tssycd sycd, tssynm synm from tblsys 
				where tsdscd = 'YESNO'
			) YN on sycd = trdpfg
			where 1 = 1
		`)
	SO_Class.Log.CetakKunci(false, c, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "trprcd",
		columns: []Kolum{
			{"trprcd", KolumProperty{name: "Code"}},
			{"trprnm", KolumProperty{name: "Description"}},
			{"trsyv1", KolumProperty{name: "Value 1"}},
			{"trsyv2", KolumProperty{name: "Value 2"}},
			{"trsyv3", KolumProperty{name: "Value 3"}},
			{"trdpfg_desc", KolumProperty{name: "Display Flag"}},
			{"trremk", KolumProperty{name: "Remark"}},
			{"trusrm", KolumProperty{name: "User Remark"}},
		},
		defaultField: true,
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}
func (saya tblprm) FillForm(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.CetakKunci(false, c, "Masuk TBLPRM-FillForm()")
	key := c.Param("trprcd")
	sqlstm := SO_Class.Fmt.Sprint(`
		select * from tblprm 
		where trprcd = '`, key, `'
	`)
	SO_Class.Log.CetakKunci(true, c, sqlstm)
	hasil = Form.GetRs(c, sqlstm)
	return hasil
}

func (saya tblprm) LoadFormObject(c *gin.Context) (hasil SO_Class.Hasil) {
	hasil.Sukses = false
	hasil.Pesan = ""
	hasil.Data = nil

	frmId := c.Param("FrmId")
	menuId := SO_Class.Strings.ToUpper("tblprm")

	result := FormObject{
		Prefix: "tr",
		Frame1: Form.CrtObj(ObjGrd{
			Id: "Grid1", FrmId: frmId, MenuId: menuId,
			Title:      "Master Table Parameter",
			Controller: SO_Class.Strings.ToUpper("tblprm"),
			Method:     "LoadGrid",
		}),
		Frame2: Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
			Id: "Panel1",
			Items: []interface{}{
				Form.CrtObj(ObjTxt{Mode: "2", FrmId: frmId, MenuId: menuId,
					Id: "trprcd", Name: "Code", AllowBlank: false, Width: 100,
				}),
				Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "trprnm", Name: "Description", AllowBlank: false, Width: 300,
				}),
				Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "trsyt1", Name: "Text 1", Width: 300,
				}),
				Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "trsyt2", Name: "Text 2", Width: 300,
				}),
				Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "trsyt3", Name: "Text 3", Width: 300,
				}),
				Form.CrtObj(ObjNum{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "trsyv1", Name: "Value 1", Width: 100,
				}),
				Form.CrtObj(ObjNum{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "trsyv2", Name: "Value 2", Width: 100,
				}),
				Form.CrtObj(ObjNum{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "trsyv3", Name: "Value 3", Width: 100,
				}),
				Form.CrtObj(ObjCmb{C: c, Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "trdpfg", Name: "Display Flag", AllowBlank: false,
					Table: "YESNO", Value: "1", Width: 100,
				}),
				Form.CrtObj(ObjRmk{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "trremk", Name: "Remark",
				}),
				Form.CrtObj(ObjRmk{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "trusrm", Name: "Internal Use Remark",
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
	Form.Add("TBLPRM", TBLPRM)
	if Form.logPrintInitFlag {
		SO_Class.Log.Println(true, "Masuk form-tblprm-init()")
	}
}

// Exported instance
var TBLPRM tblprm
