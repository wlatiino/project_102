package SO_Form

import (
	SO_Class "SOApp_GO/class"

	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type mitmas struct {
	table struct {
		MMITNOIY string `tipe:"1" json:"ID" key:"true" required:"true"`
		MMITNO   string `tipe:"2" json:"Item No" required:"true"`
		MMITDS   string `tipe:"2" json:"Item Description" required:"true"`
		MMITCL   string `tipe:"2" json:"Item Category"`
		MMITGR   string `tipe:"2" json:"Item Group"`
		MMITTY   string `tipe:"2" json:"Item Type"`
		MMITBR   string `tipe:"2" json:"Item Brand"`
		MMUNMS   string `tipe:"2" json:"Unit Measurement" required:"true"`
		MMRUMS   string `tipe:"2" json:"Rumus" required:"true"`
		MMUSRM   string `tipe:"2" json:"Internal Use Remark"`
		MMREMK   string `tipe:"2" json:"Remark"`
		MMCSDT   string `tipe:"X" `
	}
}

func (saya mitmas) Save(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.CetakKunci(true, c, "Masuk MITMAS-Save()")
	hasil.Sukses = false
	hasil.Pesan = "Masuk MITMAS - Save"
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
				tx, c,
				c.Param("Mode"),
				c.Param("username"),
				c.Param("source"),
				c.Param("mmcsdt"),
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

func (saya mitmas) StpSave(tx Transaction, c *gin.Context,
	mode string, userName string, source string, csdt string, table mitmas,
) (tr TransactionResult) {

	data := table.table
	tr.Sintax = " -- mitmas.StpSave() "

	_, errBFCS := Form.CheckRecord_BFCS(ParamBFCS{
		Tx:       tx,
		Mode:     mode,
		Table:    "mitmas",
		KeyField: "mmitnoiy",
		KeyValue: data.MMITNOIY,
		CSDT:     csdt,
	})
	if errBFCS != nil {
		tr.Error = errBFCS
		return tr
	}

	if mode == "A" {
		data.MMITNOIY = Form.GetTBLNOR(tx, userName, "mitmas")
	}

	sqlstm := Form.GetSintaxSQL_IUD(ParamIUD{
		TableName:  "mitmas",
		UserName:   userName,
		Source:     source,
		Mode:       mode,
		StructAnda: data,
	})
	SO_Class.Log.CetakKunci(true, c, " mitmas.stpSave ", sqlstm)

	_, err := Form.Execute(tx, &tr, userName, sqlstm)
	if err != nil {
		tr.Error = err
		return tr
	}

	return tr
}

func (saya mitmas) LoadGrid(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.CetakKunci(false, c, "Masuk MITMAS-LoadGrid()")
	if c.Param("sort") == "" {
		param := gin.Param{Key: "sort", Value: `[{"property": "mmitno","direction": "asc"}]`}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select
				mmitnoiy
				, mmitno, mmitds, mmitcl, mmitgr 
				, mmitty, mmitbr 
				, mmunms, mmrums
				, mmremk, mmusrm
				`, Form.GetDefaultField("mm"), ` 
			from mitmas
			where 1 = 1
		`)
	SO_Class.Log.CetakKunci(false, c, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "mmitnoiy",
		columns: []Kolum{
			{"mmitno", KolumProperty{name: "Item No", noHideable: true}},
			{"mmitds", KolumProperty{name: "Item Description"}},
			{"mmitcl", KolumProperty{name: "Item Category"}},
			{"mmitgr", KolumProperty{name: "Item Group"}},
			{"mmitty", KolumProperty{name: "Item Type"}},
			{"mmitbr", KolumProperty{name: "Item Brand"}},
			{"mmunms", KolumProperty{name: "Unit Measurement"}},
			{"mmrums", KolumProperty{name: "Rumus"}},
			{"mmremk", KolumProperty{name: "Remark"}},
			{"mmusrm", KolumProperty{name: "User Remark"}},
		},
		defaultField: true,
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}
func (saya mitmas) FillForm(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.CetakKunci(false, c, "Masuk MITMAS-FillForm()")
	key := c.Param("mmitnoiy")
	sqlstm := SO_Class.Fmt.Sprint(`
		select mitmas.*
			, itcl.tssycd as mmitcl_tssycd, itcl.tssynm as mmitcl_tssynm 
			, itgr.tssycd as mmitgr_tssycd, itgr.tssynm as mmitgr_tssynm 
			, itty.tssycd as mmitty_tssycd, itty.tssynm as mmitty_tssynm 
			, itbr.tssycd as mmitbr_tssycd, itbr.tssynm as mmitbr_tssynm 
			, unms.tssycd as mmunms_tssycd, unms.tssynm as mmunms_tssynm 
		from mitmas 
		left join tblsys itcl on itcl.tsdscd = 'ITCL' and itcl.tssycd = mmitcl
		left join tblsys itgr on itgr.tsdscd = 'ITGR' and itgr.tssycd = mmitgr
		left join tblsys itty on itty.tsdscd = 'ITTY' and itty.tssycd = mmitty
		left join tblsys itbr on itbr.tsdscd = 'ITBR' and itbr.tssycd = mmitbr
		left join tblsys unms on unms.tsdscd = 'UNMS' and unms.tssycd = mmunms
		where mmitnoiy = '`, key, `'
	`)
	SO_Class.Log.CetakKunci(true, c, sqlstm)
	hasil = Form.GetRs(c, sqlstm)
	return hasil
}

func (saya mitmas) LoadFormObject(c *gin.Context) (hasil SO_Class.Hasil) {
	hasil.Sukses = false
	hasil.Pesan = ""
	hasil.Data = nil

	frmId := c.Param("FrmId")
	menuId := SO_Class.Strings.ToUpper("mitmas")

	result := FormObject{
		Prefix: "mm",
		Frame1: Form.CrtObj(ObjGrd{
			Id: "Grid1", FrmId: frmId, MenuId: menuId,
			Title:      "Master Item",
			Controller: SO_Class.Strings.ToUpper("mitmas"),
			Method:     "LoadGrid",
		}),
		Frame2: []interface{}{
			Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
				Id: "PanelX", TipeVH: "H",
				Items: []interface{}{
					Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
						Id: "PanelA", LabelWidth: 180,
						Items: []interface{}{
							Form.CrtObj(ObjTxt{Mode: "3", FrmId: frmId, MenuId: menuId,
								Id: "mmitnoiy", Name: "IY", Hidden: true,
							}),
							Form.CrtObj(ObjTxt{Mode: "2", FrmId: frmId, MenuId: menuId,
								Id: "mmitno", Name: "Item No", AllowBlank: false,
							}),
							Form.CrtObj(ObjTxt{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "mmitds", Name: "Item Description", AllowBlank: false, Width: 400,
							}),
							Form.CrtObj(ObjPop{Mode: "2", FrmId: frmId, MenuId: menuId,
								Controller: SO_Class.Strings.ToUpper("tblsys"), Method: "LoadGridPopUp",
								Id: "mmitcl", Name: "Item Category", AllowBlank: false,
								PopCode: "tssycd", PopDesc: "tssynm",
								PopCodeFf: "mmitcl_tssycd", PopDescFf: "mmitcl_tssynm",
								SqlCondition: " and tsdscd = 'ITCL' ",
							}),
							Form.CrtObj(ObjPop{Mode: "2", FrmId: frmId, MenuId: menuId,
								Controller: SO_Class.Strings.ToUpper("tblsys"), Method: "LoadGridPopUp",
								Id: "mmitgr", Name: "Item Group", AllowBlank: false,
								PopCode: "tssycd", PopDesc: "tssynm",
								PopCodeFf: "mmitgr_tssycd", PopDescFf: "mmitgr_tssynm",
								SqlCondition: " and tsdscd = 'ITGR' ",
							}),
							Form.CrtObj(ObjPop{Mode: "2", FrmId: frmId, MenuId: menuId,
								Controller: SO_Class.Strings.ToUpper("tblsys"), Method: "LoadGridPopUp",
								Id: "mmitty", Name: "Item Type", AllowBlank: false,
								PopCode: "tssycd", PopDesc: "tssynm",
								PopCodeFf: "mmitty_tssycd", PopDescFf: "mmitty_tssynm",
								SqlCondition: " and tsdscd = 'ITTY' ",
							}),
							Form.CrtObj(ObjPop{Mode: "2", FrmId: frmId, MenuId: menuId,
								Controller: SO_Class.Strings.ToUpper("tblsys"), Method: "LoadGridPopUp",
								Id: "mmitbr", Name: "Item Brand", AllowBlank: false,
								PopCode: "tssycd", PopDesc: "tssynm",
								PopCodeFf: "mmitbr_tssycd", PopDescFf: "mmitbr_tssynm",
								SqlCondition: " and tsdscd = 'ITBR' ",
							}),
							Form.CrtObj(ObjPop{Mode: "2", FrmId: frmId, MenuId: menuId,
								Controller: SO_Class.Strings.ToUpper("tblsys"), Method: "LoadGridPopUp",
								Id: "mmunms", Name: "Unit Measurement", AllowBlank: false,
								PopCode: "tssycd", PopDesc: "tssynm",
								PopCodeFf: "mmunms_tssycd", PopDescFf: "mmunms_tssynm",
								SqlCondition: " and tsdscd = 'UNMS' ",
							}),
							Form.CrtObj(ObjCmb{Mode: "0", FrmId: frmId, MenuId: menuId, C: c,
								Id: "mmrums", Name: "Stock Formula", Table: "MITMAS_RUMS", AllowBlank: false, Width: 150,
							}),
							Form.CrtObj(ObjCmb{Mode: "0", FrmId: frmId, MenuId: menuId, C: c,
								Id: "mmdpfg", Name: "Display Flag", Table: "YESNO", AllowBlank: false, Width: 100,
							}),
							Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "mmremk", Name: "Remark",
							}),
							Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "mmusrm", Name: "Internal Use Remark",
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
	Form.Add("MITMAS", MITMAS)
	if Form.logPrintInitFlag {
		SO_Class.Log.Println(true, "Masuk form-mitmas-init()")
	}
}

// Exported instance
var MITMAS mitmas
