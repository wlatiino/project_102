package SO_Form

import (
	SO_Class "SOApp_GO/class"

	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type tblplf struct {
	table struct {
		TZNOMRIY string `tipe:"1" json:"IY" key:"true" required:"true"`
		TZNOTR   string `tipe:"2" json:"No Transaksi" required:"true"`
		TZMENUIY string `tipe:"2" json:"TBLMNU IY"`
		TZKEYC   string `tipe:"2" json:"Key Code"`
		TZKEYD   string `tipe:"2" json:"Key Duplication"`
		TZSTMT   string `tipe:"2" json:"SqlStatement"`
		TZSTAT   string `tipe:"2" json:"Status"`
		TZSTTM   string `tipe:"2" json:"Start Date"`
		TZENTM   string `tipe:"2" json:"End Date"`
		TZUSRM   string `tipe:"2" json:"User Remark"`
		TZREMK   string `tipe:"2" json:"Remark"`
		TZCSDT   string `tipe:"X" `
	}
}

func (saya tblplf) Save(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.CetakKunci(true, c, "Masuk TBLPLF-Save()")
	hasil.Sukses = false
	hasil.Pesan = "Masuk TBLPLF - Save"
	hasil.Data = nil
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
				c.Param("tmcsdt"),
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

func (saya tblplf) StpSave(tx Transaction,
	mode string, userName string, source string, csdt string, table tblplf,
) (tr TransactionResult) {

	data := table.table
	tr.Sintax = " -- tblplf.StpSave() "

	_, errBFCS := Form.CheckRecord_BFCS(ParamBFCS{
		Tx:       tx,
		Mode:     mode,
		Table:    "tblplf",
		KeyField: "tznomriy",
		KeyValue: data.TZNOMRIY,
		CSDT:     csdt,
	})
	if errBFCS != nil {
		tr.Error = errBFCS
		return tr
	}

	if mode == "A" {
		data.TZNOMRIY = Form.GetTBLNOR(tx, userName, "tblplf")
	}

	sqlstm := Form.GetSintaxSQL_IUD(ParamIUD{
		TableName:  "tblplf",
		UserName:   userName,
		Source:     source,
		Mode:       mode,
		StructAnda: data,
	})
	SO_Class.Log.Println(false, " tblplf save ", sqlstm)

	_, err := Form.Execute(tx, &tr, userName, sqlstm)
	if err != nil {
		tr.Error = err
		return tr
	}

	return tr
}

func (saya tblplf) LoadGrid(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.CetakKunci(false, c, "Masuk TBLPLF-LoadGrid()")
	if c.Param("sort") == "" {
		param := gin.Param{Key: "sort", Value: `[{"property": "tznotr","direction": "asc"}]`}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select
				tznomriy
				, tznotr, tzkeyc, tzkeyd
				, tzstmt, tzstat 
				, tzsttm, tzentm 
				, tzremk, tzusrm
				`, Form.GetDefaultField("tz"), ` 
			from tblplf
			where 1 = 1
		`)
	SO_Class.Log.CetakKunci(false, c, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "tznomriy",
		columns: []Kolum{
			{"tznotr", KolumProperty{name: "No", noHideable: true}},
			{"tzkeyc", KolumProperty{name: "Code"}},
			{"tzkeyd", KolumProperty{name: "Description"}},
			{"tzstmt", KolumProperty{name: "Statement"}},
			{"tzstat", KolumProperty{name: "Status"}},
			{"tzsttm", KolumProperty{name: "Start Date"}},
			{"tzentm", KolumProperty{name: "End Date"}},
			{"tzremk", KolumProperty{name: "Remark"}},
			{"tzusrm", KolumProperty{name: "User Remark"}},
		},
		defaultField: true,
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}
func (saya tblplf) FillForm(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.CetakKunci(false, c, "Masuk TBLPLF-FillForm()")
	key := c.Param("tznomriy")
	sqlstm := SO_Class.Fmt.Sprint("select * from tblplf where tznomriy = '", key, "'")
	SO_Class.Log.CetakKunci(true, c, sqlstm)
	hasil = Form.GetRs(c, sqlstm)
	return hasil
}

func (saya tblplf) LoadFormObject(c *gin.Context) (hasil SO_Class.Hasil) {
	hasil.Sukses = false
	hasil.Pesan = ""
	hasil.Data = nil

	frmId := c.Param("FrmId")
	menuId := SO_Class.Strings.ToUpper("tblplf")

	result := FormObject{
		Prefix: "tz",
		Frame1: Form.CrtObj(ObjGrd{
			Id: "Grid1", FrmId: frmId, MenuId: menuId,
			Title:      "Table Process Batch Log File",
			Controller: SO_Class.Strings.ToUpper("tblplf"),
			Method:     "LoadGrid",
		}),
		Frame2: []interface{}{
			Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
				Id: "PanelX", TipeVH: "H",
				Items: []interface{}{
					Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
						Id: "PanelA",
						Items: []interface{}{
							Form.CrtObj(ObjTxt{Mode: "3", FrmId: frmId, MenuId: menuId,
								Id: "tznomriy", Name: "IY", Hidden: true,
							}),
							Form.CrtObj(ObjTxt{Mode: "3", FrmId: frmId, MenuId: menuId,
								Id: "tznotr", Name: "No Transaction", AllowBlank: false,
							}),
							Form.CrtObj(ObjCmb{Mode: "3", FrmId: frmId, MenuId: menuId, C: c,
								Id: "tzstat", Name: "Status", Table: "YESNO", AllowBlank: false, Width: 100,
							}),
							Form.CrtObj(ObjTxt{Mode: "3", FrmId: frmId, MenuId: menuId,
								Id: "tzkeyc", Name: "Code", AllowBlank: false, Width: 400,
							}),
							Form.CrtObj(ObjTxt{Mode: "3", FrmId: frmId, MenuId: menuId,
								Id: "tzsttm", Name: "Start Date", AllowBlank: false,
							}),
							Form.CrtObj(ObjTxt{Mode: "3", FrmId: frmId, MenuId: menuId,
								Id: "tzentm", Name: "End Date", AllowBlank: false,
							}),
							Form.CrtObj(ObjRmk{Mode: "3", FrmId: frmId, MenuId: menuId,
								Id: "tzusrm", Name: "Internal Use Remark",
							}),
							Form.CrtObj(ObjRmk{Mode: "3", FrmId: frmId, MenuId: menuId,
								Id: "tzremk", Name: "Remark",
							}),
						},
					}),
				},
			}),
			Form.CrtObj(ObjPnl{
				FrmId: frmId, MenuId: menuId,
				Id: "PanelY", TipeVH: "H",
				Items: []interface{}{
					Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
						Id: "PanelC",
						Items: []interface{}{
							Form.CrtObj(ObjRmk{Mode: "3", FrmId: frmId, MenuId: menuId,
								Id: "tzkeyd", Name: "Description", Height: 300,
							}),
						},
					}),
					Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
						Id: "PanelD", LabelWidth: 200,
						Items: []interface{}{
							Form.CrtObj(ObjRmk{Mode: "3", FrmId: frmId, MenuId: menuId,
								Id: "tzstmt", Name: "Statement", Height: 300,
							}),
						},
					}),
				},
			}),
		},
	}

	hasil.Sukses = true
	hasil.Pesan = "sukses"
	hasil.Data = result
	return hasil
}

func init() {
	Form.Add("TBLPLF", TBLPLF)
	if Form.logPrintInitFlag {
		SO_Class.Log.Println(true, "Masuk form-tblplf-init()")
	}
}

// Exported instance
var TBLPLF tblplf
