package SO_Form

import (
	SO_Class "SOApp_GO/class"

	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type tblusr struct {
	table struct {
		TUUSERIY string `tipe:"1" json:"User IY" key:"true"`
		TUUSER   string `tipe:"1" json:"User Login"  max:"50"` // max:"20"
		TUPSWD   string `tipe:"1" json:"Password"`
		TUNAME   string `tipe:"2" json:"User Name"`
		TUUSLV   string `tipe:"2" json:"User Level"`
		TUEMID   string `tipe:"2" json:"Employee ID"`
		TUDEPT   string `tipe:"2" json:"Departuent"`
		TUMAIL   string `tipe:"2" json:"Mail"`
		TUWELC   string `tipe:"2" json:"Welcome Text"`
		TUEXPP   string `tipe:"2" json:"Expired"`
		TUEXPD   string `tipe:"2" json:"Expired Date"`
		TUEXPV   string `tipe:"2" json:"Expired Value"`
		TUMNTP   string `tipe:"2" json:"Menu Type"`
		TULGCT   string `tipe:"X" json:"Login Counter"`
		TULSLI   string `tipe:"X" json:"Last Login"`
		TULSLO   string `tipe:"X" json:"Last Logoff"`
		TUTOKN   string `tipe:"X" json:"Token"`
		TUDPFG   string `tipe:"2" json:"Display Flag"`
		TUREMK   string `tipe:"2" json:"Keterangan"`
		TUUSRM   string `tipe:"2" json:"Internal User Remark"`
		TUCSDT   string `tipe:"X" `
	}
}

func (saya tblusr) Save(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLUSR-Save()")
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
				c.Param("tucsdt"),
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

func (saya tblusr) StpSave(tx Transaction,
	mode string, userName string, source string, csdt string, table tblusr,
) (tr TransactionResult) {

	data := table.table
	tr.Sintax = " -- tblusr.StpSave() "

	_, errBFCS := Form.CheckRecord_BFCS(ParamBFCS{
		Tx:       tx,
		Mode:     mode,
		Table:    "tblusr",
		KeyField: "tuuseriy",
		KeyValue: data.TUUSERIY,
		CSDT:     csdt,
	})
	if errBFCS != nil {
		tr.Error = errBFCS
		return tr
	}

	if mode == "A" {
		data.TUUSERIY = Form.GetTBLNOR(tx, userName, "tblusr")
		data.TUPSWD, _ = Form.HashPassword(data.TUPSWD)
	}

	uslv := ""
	errUSLV := tx.QueryRow(SO_Class.Fmt.Sprint("select tuuslv from tblusr where tuuser = '", userName, "'")).Scan(&uslv)
	SO_Class.Log.Println(true, " get user level ", errUSLV)
	if uslv >= data.TUUSLV {
		tr.Error = SO_Class.Fmt.Errorf("Your user level '%s' is not qualified!", userName)
		return tr
	}

	sqlstm := Form.GetSintaxSQL_IUD(ParamIUD{
		TableName:  "tblusr",
		UserName:   userName,
		Source:     source,
		Mode:       mode,
		StructAnda: data,
	})
	SO_Class.Log.Println(false, " tblusr save ", sqlstm)

	_, err := Form.Execute(tx, &tr, userName, sqlstm)
	if err != nil {
		tr.Error = err
		return tr
	}

	return tr
}

func (saya tblusr) LoadGrid(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLUSR-LoadGrid()")
	if c.Param("sort") == "" {
		param := gin.Param{Key: "sort", Value: `[{"property": "tuuser","direction": "asc"}]`}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select
				tuuseriy, tuuser , tuname 
				, tuemid, tudept, tumail
				, tuuslv 
				, tulsli, tulslo             
				, nullif(tuexpd, '')::date tuexpd          
				, tuwelc 
				, turemk, tuusrm
				`, Form.GetDefaultField("tu"), ` 
			from tblusr
			where 1 = 1
		`)
	SO_Class.Log.Println(false, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "tuuseriy",
		columns: []Kolum{
			{"tuuseriy", KolumProperty{name: "ID", hidden: true}},
			{"tuuser", KolumProperty{name: "Login Name"}},
			{"tuname", KolumProperty{name: "Full Name"}},
			{"tuuslv", KolumProperty{name: "User Level"}},
			{"tuemid", KolumProperty{name: "Employee ID"}},
			{"tudept", KolumProperty{name: "Departemen"}},
			{"tumail", KolumProperty{name: "E-Mail"}},
			{"tulsli", KolumProperty{name: "Last Login"}},
			{"tulslo", KolumProperty{name: "Last Logout"}},
			{"tuexpd", KolumProperty{name: "Expired Date"}},
			{"tuwelc", KolumProperty{name: "Welcome Text"}},
			{"turemk", KolumProperty{name: "Remark"}},
			{"tuusrm", KolumProperty{name: "User Remark"}},
		},
		defaultField: true,
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}

func (saya tblusr) FillForm(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLUSR-FillForm()")
	key := c.Param("tuuseriy")
	sqlstm := SO_Class.Fmt.Sprint("select * from tblusr where tuuseriy = '", key, "'")
	SO_Class.Log.Println(true, sqlstm)
	hasil = Form.GetRs(c, sqlstm)
	return hasil
}

func (saya tblusr) LoadFormObject(c *gin.Context) (hasil SO_Class.Hasil) {
	hasil.Sukses = false
	hasil.Pesan = ""
	hasil.Data = nil

	frmId := c.Param("FrmId")
	menuId := SO_Class.Strings.ToUpper("tblusr")

	result := FormObject{
		Prefix: "tu",
		Frame1: Form.CrtObj(ObjGrd{
			Id: "Grid1", FrmId: frmId, MenuId: menuId,
			Title:      "Table User",
			Controller: SO_Class.Strings.ToUpper("tblusr"),
			Method:     "LoadGrid",
		}),
		Frame2: []interface{}{
			Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
				Id: "PanelA", LabelWidth: 150,
				Items: []interface{}{
					Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
						Id: "PanelAH", TipeVH: "H", Margin: "0", Padding: "0",
						Items: []interface{}{
							Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
								Id: "PanelA1", LabelWidth: 150, Margin: "0", Padding: "0",
								Items: []interface{}{
									Form.CrtObj(ObjTxt{Mode: "3", FrmId: frmId, MenuId: menuId,
										Id: "tuuseriy", Name: "IY", Hidden: true,
									}),
									Form.CrtObj(ObjTxt{Mode: "2", FrmId: frmId, MenuId: menuId,
										Id: "tuuser", Name: "Login Username", AllowBlank: false,
									}),
									Form.CrtObj(ObjTxt{Mode: "2", FrmId: frmId, MenuId: menuId,
										Id: "tuname", Name: "Full Name", AllowBlank: false, Width: 500,
									}),
									Form.CrtObj(ObjTxt{Mode: "2", FrmId: frmId, MenuId: menuId,
										Id: "tupswd", Name: "Password", AllowBlank: false, InputType: "password", Width: 250,
									}),
									Form.CrtObj(ObjCnt{FrmId: frmId, MenuId: menuId,
										Id: "ContainerA",
										Items: []interface{}{
											Form.CrtObj(ObjCmb{C: c, Mode: "3", FrmId: frmId, MenuId: menuId,
												Id: "tudsfg", Name: "Disable User", Table: "YESNO", AllowBlank: false, Value: "0",
											}),
											Form.CrtObj(ObjCmb{C: c, Mode: "0", FrmId: frmId, MenuId: menuId,
												Id: "tudpfg", Name: "Display User", Table: "DSPLY", AllowBlank: false, Value: "1", LabelWidth: 120,
											}),
										},
									}),
									Form.CrtObj(ObjRad{C: c, Mode: "0", FrmId: frmId, MenuId: menuId,
										Id: "tuuslv", Name: "User Level", Table: "USLV", AllowBlank: false,
									}),
									Form.CrtObj(ObjChg{C: c, Mode: "0", FrmId: frmId, MenuId: menuId,
										Id: "tumntp", Name: "Menu Type", Table: "MNTP", AllowBlank: false, Columns: 2,
									}),
									Form.CrtObj(ObjRad{C: c, Mode: "0", FrmId: frmId, MenuId: menuId,
										Id: "tuexpp", Name: "Expired Password", Table: "YESNO", AllowBlank: false, Value: "1",
									}),
									Form.CrtObj(ObjCnt{FrmId: frmId, MenuId: menuId,
										Id: "ContainerB",
										Items: []interface{}{
											Form.CrtObj(ObjDtp{Mode: "0", FrmId: frmId, MenuId: menuId,
												Id: "tuexpd", Name: "Expired Date",
											}),
											Form.CrtObj(ObjNum{Mode: "0", FrmId: frmId, MenuId: menuId,
												Id: "tuexpv", Name: "Expired Value", AllowBlank: false, Width: 80, LabelWidth: 120,
											}),
										},
									}),
									// Form.CrtObj(ObjCnt{FrmId: frmId, MenuId: menuId,
									// 	Id: "ContainerC",
									// 	Items: []interface{}{
									// 		Form.CrtObj(ObjBtn{Mode: "0", FrmId: frmId, MenuId: menuId,
									// 			Id: "btnresetpass", text: "Reset Password", Margin: "5 0 0 155", onKlik: onBtnResetPass, hidden: hiddenasup,
									// 		}),
									// 		Form.CrtObj(ObjBtn{Mode: "0", FrmId: frmId, MenuId: menuId,
									// 			Id: "btnresetpin", text: "Reset PIN", Margin: "5 0 0 5", onKlik: onBtnResetPIN, hidden: hiddenasup,
									// 		}),
									// 	},
									// }),
								},
							}),
							Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
								Id: "PanelA2", LabelWidth: 150, Margin: "0", Padding: "0",
								Items: []interface{}{
									Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
										Id: "tuemid", Name: "Employee ID", AllowBlank: true,
									}),
									Form.CrtObj(ObjTxt{Mode: "1", FrmId: frmId, MenuId: menuId,
										Id: "tudept", Name: "Department", AllowBlank: true,
									}),
									Form.CrtObj(ObjTxt{Mode: "2", FrmId: frmId, MenuId: menuId,
										Id: "tumail", Name: "Email", AllowBlank: true,
									}),
									Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmId, MenuId: menuId,
										Id: "tuwelc", Name: "Welcome Text",
									}),
									Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmId, MenuId: menuId,
										Id: "turemk", Name: "Remark",
									}),
									Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmId, MenuId: menuId,
										Id: "tuusrm", Name: "Internal Use Remark",
									}),
								},
							}),
						},
					}),
					Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
						Id: "PanelAB", Margin: "0", Padding: "0",
						Items: []interface{}{
							Form.CrtObj(ObjGrd{FrmId: frmId, MenuId: menuId,
								Controller: SO_Class.Strings.ToUpper("tbluam"), Method: "LoadGridTBLUSR",
								Id: "GridX", Width: 1000, Height: 400, ShowGridAdvSearch: false,
								Title: "Table Menu Access",
							}),
						},
					}),
				},
			}),
		},
		SOFn: `{
			formAction: {
				add: function ({form}) {
					// Ext.getCmp('Frm` + frmId + `Obj` + `btnresetpass').hide();
					// Ext.getCmp('Frm` + frmId + `Obj` + `btnresetpin').hide();
				},
				edit: function ({form}) {
					// Ext.getCmp('Frm` + frmId + `Obj` + `btnresetpass').hide();
					// Ext.getCmp('Frm` + frmId + `Obj` + `btnresetpin').hide();
				},
				cancel: function ({form}) {
					Ext.getCmp('Frm` + frmId + `Grd` + `GridX').hide();
				},
				view: function ({form}) {
					// Ext.getCmp('Frm` + frmId + `Obj` + `btnresetpass').hide();
					// Ext.getCmp('Frm` + frmId + `Obj` + `btnresetpin').hide();

					var useriy = Ext.getCmp('Frm` + frmId + `Obj` + `tuuseriy').getValue();
					var gridX = Ext.getCmp('Frm` + frmId + `Grd` + `GridX');

					gridX.soObj_sqlCondition = " And tauseriy = '"+useriy+"' ";
					gridX.loadGrid();

					gridX.show();
				},
			},
		}`,
	}

	hasil.Sukses = true
	hasil.Pesan = "sukses"
	hasil.Data = result
	return hasil
}

func (saya tblusr) LoadGridUserLevel(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLUSR-LoadGridUserLevel()")
	if c.Param("sort") == "" {
		param := gin.Param{Key: "sort", Value: `[{"property": "tuuser","direction": "asc"}]`}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select
				tuuseriy, tuuser , tuname 
				, tuemid, tudept, tumail
				, tuuslv 
				, tulsli, tulslo             
				, nullif(tuexpd, '')::date tuexpd          
				, tuwelc 
				, turemk, tuusrm
				`, Form.GetDefaultField("tu"), ` 
			from tblusr
			where tuuslv > (
				select tuuslv from tblusr where tuuser = '`, c.Param(("username")), `'
			)
		`)
	SO_Class.Log.Println(true, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "tuuseriy",
		columns: []Kolum{
			{"tuuseriy", KolumProperty{name: "ID", hidden: true}},
			{"tuuser", KolumProperty{name: "Login Name"}},
			{"tuname", KolumProperty{name: "Full Name"}},
			{"tuuslv", KolumProperty{name: "User Level"}},
			{"tuemid", KolumProperty{name: "Employee ID"}},
			{"tudept", KolumProperty{name: "Departemen"}},
			{"tumail", KolumProperty{name: "E-Mail"}},
			{"tulsli", KolumProperty{name: "Last Login"}},
			{"tulslo", KolumProperty{name: "Last Logout"}},
			{"tuexpd", KolumProperty{name: "Expired Date"}},
			{"tuwelc", KolumProperty{name: "Welcome Text"}},
			{"turemk", KolumProperty{name: "Remark"}},
			{"tuusrm", KolumProperty{name: "User Remark"}},
		},
		defaultField: true,
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}

func init() {
	Form.Add("TBLUSR", TBLUSR)
	if Form.logPrintInitFlag {
		SO_Class.Log.Println(true, "Masuk form-tblusr-init()")
	}
}

// Exported instance
var TBLUSR tblusr
