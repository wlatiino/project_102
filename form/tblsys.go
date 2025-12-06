package SO_Form

import (
	SO_Class "SOApp_GO/class"

	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type tblsys struct {
	table struct {
		TSNOMRIY string `tipe:"X" json:"ID"`
		TSDSCD   string `tipe:"1" json:"Table Code" key:"true" required:"true" max:"20"` //required:"true" max:"20"
		TSSYCD   string `tipe:"1" json:"Code" key:"true" required:"true" max:"20"`       //required:"true" max:"20"
		TSSYNM   string `tipe:"2" json:"Name" required:"true"`
		TSSYT1   string `tipe:"2" json:"Text 1"`
		TSSYT2   string `tipe:"2" json:"Text 2"`
		TSSYT3   string `tipe:"2" json:"Text 3"`
		TSSYV1   string `tipe:"2" json:"Value 1"`
		TSSYV2   string `tipe:"2" json:"Value 2"`
		TSSYV3   string `tipe:"2" json:"Value 3"`
		TSLST1   string `tipe:"2" json:"Label Text 1"`
		TSLST2   string `tipe:"2" json:"Label Text 2"`
		TSLST3   string `tipe:"2" json:"Label Text 3"`
		TSLSV1   string `tipe:"2" json:"Label Value 1"`
		TSLSV2   string `tipe:"2" json:"Label Value 2"`
		TSLSV3   string `tipe:"2" json:"Label Value 3"`
		TSDPFG   string `tipe:"2" json:"Display Flag"`
		TSREMK   string `tipe:"2" json:"Remark"`
		TSUSRM   string `tipe:"2" json:"User Remark"`
		TSCSDT   string `tipe:"X" `
	}
}

func (saya tblsys) Save(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLSYS-Save()")
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
				c.Param("tscsdt"),
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

func (saya tblsys) StpSave(tx Transaction,
	mode string, userName string, source string, csdt string, table tblsys,
) (tr TransactionResult) {

	data := table.table
	tr.Sintax = " -- tblsys.StpSave() "

	_, errBFCS := Form.CheckRecord_BFCS(ParamBFCS{
		Tx:       tx,
		Mode:     mode,
		Table:    "tblsys",
		KeyField: "tsnomriy",
		KeyValue: data.TSNOMRIY,
		CSDT:     csdt,
	})
	if errBFCS != nil {
		tr.Error = errBFCS
		return tr
	}

	sqlstm := Form.GetSintaxSQL_IUD(ParamIUD{
		TableName:  "tblsys",
		UserName:   userName,
		Source:     source,
		Mode:       mode,
		StructAnda: data,
	})
	SO_Class.Log.Println(false, " tblsys save ", sqlstm)

	_, err := Form.Execute(tx, &tr, userName, sqlstm)
	if err != nil {
		tr.Error = err
		return tr
	}

	return tr
}

func (saya tblsys) LoadData(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLSYS-LoadGrid()")

	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select 
				tsnomriy, tsdscd, tssycd, tssynm, 
				tssyv1::dec(24,0) tssyv1, 
				tssyv2::dec(24,2) tssyv2,
				tssyv3::dec(24,4) tssyv3, 
				tsremk  
			from tblsys 
			where 1 = 1
		`)
	SO_Class.Log.Println(true, sqlstm)
	hasil = Form.GetRecordSet(c, sqlstm)
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}

func (saya tblsys) LoadGrid(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLSYS-LoadGrid()")
	if c.Param("sort") == "" {
		param := gin.Param{
			Key: "sort",
			Value: SO_Class.Fmt.Sprint(
				`[{"property": "tsdscd","direction": "asc"}`,
				`,{"property": "tssycd","direction": "asc"}]`,
			),
		}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select 
				tsnomriy, tsdscd, tssycd, tssynm, 
				tssyv1::dec(24,0) tssyv1, 
				tssyv2::dec(24,2) tssyv2,
				tssyv3::dec(24,4) tssyv3, 
				synm tsdpfg_desc,
				tsremk, tsusrm
				`, Form.GetDefaultField("ts"), ` 
			from tblsys 			
			left join (
				select tssycd sycd, tssynm synm from tblsys 
				where tsdscd = 'YESNO'
			) YN on sycd = tsdpfg
			where 1 = 1
		`)
	SO_Class.Log.Println(false, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "tsnomriy",
		columns: []Kolum{
			{"tsdscd", KolumProperty{name: "Table Code"}},
			{"tssycd", KolumProperty{name: "Code"}},
			{"tssynm", KolumProperty{name: "Description"}},
			{"tssyv1", KolumProperty{name: "Value 1"}},
			{"tssyv2", KolumProperty{name: "Value 2"}},
			{"tssyv3", KolumProperty{name: "Value 3"}},
			{"tsdpfg_desc", KolumProperty{name: "Display Flag"}},
			{"tsremk", KolumProperty{name: "Remark"}},
			{"tsusrm", KolumProperty{name: "User Remark"}},
		},
		defaultField: true,
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}
func (saya tblsys) FillForm(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLSYS-FillForm()")
	key := c.Param("tsnomriy")
	sqlstm := SO_Class.Fmt.Sprint(`
		select * from tblsys 
		left join tbldsc on tddscd = tsdscd
		where tsnomriy = '`, key, `'`,
	)
	SO_Class.Log.Println(true, sqlstm)
	hasil = Form.GetRs(c, sqlstm)
	return hasil
}

func (saya tblsys) LoadFormObject(c *gin.Context) (hasil SO_Class.Hasil) {
	hasil.Sukses = false
	hasil.Pesan = ""
	hasil.Data = nil

	frmId := c.Param("FrmId")
	menuId := SO_Class.Strings.ToUpper("tblsys")

	result := FormObject{
		Prefix: "ts",
		Frame1: Form.CrtObj(ObjGrd{
			Id: "Grid1", FrmId: frmId, MenuId: menuId,
			Title:      "Master Table System",
			Controller: SO_Class.Strings.ToUpper("tblsys"),
			Method:     "LoadGrid",
		}),
		Frame2: Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
			Id: "Panel1",
			Items: []interface{}{
				Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "tsnomriy", Name: "ID", Hidden: true,
				}),
				// Form.CrtObj(ObjTxt{
				// 	FrmId: frmId, MenuId: menuId, Mode: "2", Id: "tsdscd", Name: "Table Code",
				// }),
				Form.CrtObj(ObjPop{Mode: "2", FrmId: frmId, MenuId: menuId,
					Controller: SO_Class.Strings.ToUpper("tbldsc"), Method: "LoadGrid",
					Id: "tsdscd", Name: "Table Code", AllowBlank: false,
					PopCode: "tddscd", PopDesc: "tddsnm",
				}),
				Form.CrtObj(ObjTxt{Mode: "2", FrmId: frmId, MenuId: menuId,
					Id: "tssycd", Name: "Code", AllowBlank: false, Width: 100,
				}),
				Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "tssynm", Name: "Description", AllowBlank: false, Width: 300,
				}),
				Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
					Id: "PanelH", TipeVH: "H",
					Items: []interface{}{

						Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
							Id: "PanelH1", LabelWidth: 80,
							Items: []interface{}{

								Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
									Id: "tssyt1", Name: "Text 1", Width: 100,
								}),
								Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
									Id: "tssyt2", Name: "Text 2", Width: 100,
								}),
								Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
									Id: "tssyt3", Name: "Text 3", Width: 100,
								}),
								Form.CrtObj(ObjNum{Mode: "1", FrmId: frmId, MenuId: menuId,
									Id: "tssyv1", Name: "Value 1", Width: 100,
								}),
								Form.CrtObj(ObjNum{Mode: "1", FrmId: frmId, MenuId: menuId,
									Id: "tssyv2", Name: "Value 2", Width: 100,
								}),
								Form.CrtObj(ObjNum{Mode: "1", FrmId: frmId, MenuId: menuId,
									Id: "tssyv3", Name: "Value 3", Width: 100,
								}),
							},
						}),

						Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
							Id: "PanelH2",
							Items: []interface{}{

								Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
									Id: "tslst1", Name: "Label Text 1", Width: 200,
								}),
								Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
									Id: "tslst2", Name: "Label Text 2", Width: 200,
								}),
								Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
									Id: "tslst3", Name: "Label Text 3", Width: 200,
								}),

								Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
									Id: "tslsv1", Name: "Label Value 1", Width: 200,
								}),
								Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
									Id: "tslsv2", Name: "Label Value 2", Width: 200,
								}),
								Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
									Id: "tslsv3", Name: "Label Value 3", Width: 200,
								}),
							},
						}),
					},
				}),
				Form.CrtObj(ObjCmb{C: c, Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "tsdpfg", Name: "Display Flag", AllowBlank: false,
					Table: "YESNO", Value: "1", Width: 100,
				}),
				Form.CrtObj(ObjRmk{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "tsremk", Name: "Remark",
				}),
				Form.CrtObj(ObjRmk{Mode: "1", FrmId: frmId, MenuId: menuId,
					Id: "tsusrm", Name: "User Remark",
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
	Form.Add("TBLSYS", TBLSYS)
	if Form.logPrintInitFlag {
		SO_Class.Log.Println(true, "Masuk form-tblsys-init()")
	}
}

// Exported instance
var TBLSYS tblsys
