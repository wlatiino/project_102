package SO_Form

import (
	SO_Class "SOApp_GO/class"

	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type tblmnu struct {
	table struct {
		TMMENUIY string `tipe:"1" json:"ID" key:"true" required:"true"`
		TMNOMR   string `tipe:"2" json:"No Urut" required:"true"`
		TMGRUP   string `tipe:"2" json:"Group"`
		TMSCUT   string `tipe:"2" json:"Short Cut"`
		TMMENU   string `tipe:"2" json:"Menu" required:"true"`
		TMACES   string `tipe:"2" json:"ACES"`
		TMSYFG   string `tipe:"2" json:"System Flag"`
		TMURLW   string `tipe:"2" json:"URL"`
		TMMNTP   string `tipe:"2" json:"Menu Type"`
		TMBCDT   string `tipe:"2" json:"Back Date"`
		TMFWDT   string `tipe:"2" json:"Forward Date"`
		TMUSCT   string `tipe:"X" json:"User Count"`
		TMLSBY   string `tipe:"X" json:"Last Used By"`
		TMLSDT   string `tipe:"X" json:"Last Used Date"`
		TMRLDT   string `tipe:"X" json:"Release Date"`
		TMDPFG   string `tipe:"2" json:"Display Flag"`
		TMUSRM   string `tipe:"2" json:"User Remark"`
		TMREMK   string `tipe:"2" json:"Remark"`
		TMJSON   string `tipe:"2" json:"JSon"`
		TMDESC   string `tipe:"2" json:"Menu Description"`
		TMGRID   string `tipe:"2" json:"Grid"`
		TMCSDT   string `tipe:"X" `
	}
}

func (saya tblmnu) Save(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLMNU-Save()")
	hasil.Sukses = false
	hasil.Pesan = "Masuk TBLMNU - Save"
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

			// _, errBFCS := Form.CheckRecord_BFCS(ParamBFCS{
			// 	Tx:       tx,
			// 	Mode:     c.Param("Mode"),
			// 	Table:    "tblmnu",
			// 	KeyField: "tmmenuiy",
			// 	KeyValue: c.Param("tmmenuiy"),
			// 	CSDT:     c.Param("tmcsdt"),
			// })
			// if errBFCS != nil {
			// 	tr.Error = errBFCS
			// 	return tr
			// }

			// if c.Param("Mode") == "A" {
			// 	saya.table.TMMENUIY = Form.GetTBLNOR(tx, c.Param("username"), "tblmnu")
			// }

			// sqlstm := Form.GetSintaxSQL_IUD(ParamIUD{
			// 	TableName:  "tblmnu",
			// 	UserName:   c.Param("username"),
			// 	Source:     c.Param("source"),
			// 	Mode:       c.Param("Mode"),
			// 	StructAnda: saya.table,
			// })
			// SO_Class.Log.Println(false, " tblmnu save ", sqlstm)

			// _, err := Form.Execute(tx, &tr, userName, sqlstm)
			// if err != nil {
			// 	tr.Error = err
			// 	return tr
			// }

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

func (saya tblmnu) StpSave(tx Transaction,
	mode string, userName string, source string, csdt string, table tblmnu,
) (tr TransactionResult) {

	data := table.table
	tr.Sintax = " -- tblmnu.StpSave() "

	_, errBFCS := Form.CheckRecord_BFCS(ParamBFCS{
		Tx:       tx,
		Mode:     mode,
		Table:    "tblmnu",
		KeyField: "tmmenuiy",
		KeyValue: data.TMMENUIY,
		CSDT:     csdt,
	})
	if errBFCS != nil {
		tr.Error = errBFCS
		return tr
	}

	if mode == "A" {
		data.TMMENUIY = Form.GetTBLNOR(tx, userName, "tblmnu")
	}

	if data.TMJSON == "" {
		data.TMJSON = "{}"
	}

	sqlstm := Form.GetSintaxSQL_IUD(ParamIUD{
		TableName:  "tblmnu",
		UserName:   userName,
		Source:     source,
		Mode:       mode,
		StructAnda: data,
	})
	SO_Class.Log.Println(false, " tblmnu save ", sqlstm)

	_, err := Form.Execute(tx, &tr, userName, sqlstm)
	if err != nil {
		tr.Error = err
		return tr
	}

	return tr
}

func (saya tblmnu) LoadGrid(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLMNU-LoadGrid()")
	if c.Param("sort") == "" {
		param := gin.Param{Key: "sort", Value: `[{"property": "tmnomr","direction": "asc"}]`}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select
				tmmenuiy
				, tmnomr, tmurlw, tmmenu, tmaces, tmscut 
				, tmbcdt, tmfwdt 
				, tmusct, tmlsby, tmlsdt::bpchar tmlsdt
				, tmremk, tmusrm
				`, Form.GetDefaultField("tm"), ` 
			from tblmnu
			where 1 = 1
		`)
	SO_Class.Log.Println(false, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "tmmenuiy",
		columns: []Kolum{
			{"tmnomr", KolumProperty{name: "No", noHideable: true}},
			{"tmurlw", KolumProperty{name: "Url"}},
			{"tmmenu", KolumProperty{name: "Nama Menu"}},
			{"tmaces", KolumProperty{name: "Hak Akses"}},
			{"tmscut", KolumProperty{name: "Shortcut"}},
			{"tmbcdt", KolumProperty{name: "Back Date"}},
			{"tmfwdt", KolumProperty{name: "Forward Date"}},
			{"tmusct", KolumProperty{name: "Use Count"}},
			{"tmlsby", KolumProperty{name: "Last Use By"}},
			{"tmlsdt", KolumProperty{name: "Last Use Date"}},
			{"tmremk", KolumProperty{name: "Remark"}},
			{"tmusrm", KolumProperty{name: "User Remark"}},
		},
		defaultField: true,
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}
func (saya tblmnu) FillForm(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLMNU-FillForm()")
	key := c.Param("tmmenuiy")
	sqlstm := SO_Class.Fmt.Sprint("select * from tblmnu where tmmenuiy = '", key, "'")
	SO_Class.Log.Println(true, sqlstm)
	hasil = Form.GetRs(c, sqlstm)
	return hasil
}

func (saya tblmnu) LoadFormObject(c *gin.Context) (hasil SO_Class.Hasil) {
	hasil.Sukses = false
	hasil.Pesan = ""
	hasil.Data = nil

	frmId := c.Param("FrmId")
	menuId := SO_Class.Strings.ToUpper("tblmnu")

	result := FormObject{
		Prefix: "tm",
		Frame1: Form.CrtObj(ObjGrd{
			Id: "Grid1", FrmId: frmId, MenuId: menuId,
			Title:      "Table Menu",
			Controller: SO_Class.Strings.ToUpper("tblmnu"),
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
								Id: "tmmenuiy", Name: "IY", Hidden: true,
							}),
							Form.CrtObj(ObjTxt{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "tmnomr", Name: "Code Menu", AllowBlank: false,
							}),
							Form.CrtObj(ObjTxt{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "tmgrup", Name: "File Group",
							}),
							Form.CrtObj(ObjTxt{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "tmscut", Name: "Shortcut",
							}),
							Form.CrtObj(ObjTxt{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "tmmenu", Name: "Menu Name", AllowBlank: false, Width: 400,
							}),
							Form.CrtObj(ObjChg{Mode: "1", FrmId: frmId, MenuId: menuId, C: c,
								Id: "tmaces", Name: "Menu Access", Table: "MODE", AllowBlank: true, Columns: 3, Width: 580,
							}),
							Form.CrtObj(ObjTxt{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "tmsyfg", Name: "System Flag", AllowBlank: false,
							}),
							Form.CrtObj(ObjTxt{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "tmurlw", Name: "URL",
							}),
							Form.CrtObj(ObjRad{Mode: "0", FrmId: frmId, MenuId: menuId, C: c,
								Id: "tmmntp", Name: "Menu Type", Table: "MNTP",
							}),
							Form.CrtObj(ObjNum{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "tmbcdt", Name: "Back Date", Width: 100, Min: 0, Max: 9999,
							}),
							Form.CrtObj(ObjNum{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "tmfwdt", Name: "Forward Date", Width: 100, Min: 0, Max: 9999,
							}),
						},
					}),
					Form.CrtObj(ObjPnl{
						FrmId: frmId, MenuId: menuId,
						Id: "PanelB", LabelWidth: 200,
						Items: []interface{}{
							Form.CrtObj(ObjNum{Mode: "3", FrmId: frmId, MenuId: menuId,
								Id: "tmusct", Name: "Use Count",
							}),
							Form.CrtObj(ObjTxt{Mode: "3", FrmId: frmId, MenuId: menuId,
								Id: "tmlsby", Name: "Last Use By",
							}),
							Form.CrtObj(ObjDtp{Mode: "3", FrmId: frmId, MenuId: menuId,
								Id: "tmlsdt", Name: "Last Use Date", AllowBlank: true,
							}),
							Form.CrtObj(ObjDtp{Mode: "3", FrmId: frmId, MenuId: menuId,
								Id: "tmrldt", Name: "Release Date", AllowBlank: true,
							}),
							Form.CrtObj(ObjCmb{Mode: "0", FrmId: frmId, MenuId: menuId, C: c,
								Id: "tmdpfg", Name: "Display Flag", Table: "YESNO", AllowBlank: false, Width: 100,
							}),
							Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "tmusrm", Name: "Internal Use Remark",
							}),
							Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "tmremk", Name: "Remark",
							}),
							Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "tmjson", Name: "JSON",
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
							Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "tmdesc", Name: "Description", Height: 300,
							}),
						},
					}),
					Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
						Id: "PanelD", LabelWidth: 200,
						Items: []interface{}{
							Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmId, MenuId: menuId,
								Id: "tmgrid", Name: "Query Loadgrid", Height: 300,
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

func (saya tblmnu) LoadMenu(c *gin.Context) (hasil SO_Class.Hasil) {

	hasil.Sukses = false
	hasil.Pesan = ""
	hasil.Data = nil

	// kondisi := c.Param("kondisi")
	tuuser := c.Param("username")
	var syfg string
	syfg = "W"

	if c.Param("syfg") != "" {
		syfg = c.Param("syfg")
	}

	sqlstm :=
		SO_Class.Fmt.Sprint(`
			Select 
				tmmenuiy, 
				rtrim(tmnomr) tmnomr, rtrim(tmurlw) tmurlw, rtrim(tmmenu) tmmenu, 
				replace(concat('[', rtrim(tmscut), '] ', rtrim(tmmenu)),'[] ','') tmtext, 
				rtrim(tmaces) tmaces, 
				rtrim(coalesce(taaces,'')) taaces, 
				rtrim(tumntp) tumntp, 
				coalesce(tmjson::varchar, '') tmjson,
				rtrim(tmscut) tmscut, 
				tmremk 
			From tblmnu 
			left join tblusr on tuuser = '` + tuuser + `'
			left join tbluam on tauseriy = tuuseriy and tamenuiy = tmmenuiy 
			Where 1 = 1 
				And tmdpfg = '1'
				And tmsyfg like '%` + syfg + `%'
				And rtrim(tumntp) like concat('%',rtrim(tmmntp),'%')
			Order By tmnomr
		`)
	SO_Class.Log.Println(false, sqlstm)
	tblmnu := Form.GetRs(c, sqlstm)

	if tblmnu.Data == nil {
		SO_Class.Log.Println(false, "func TBLMNU_LoadMenu --> record menu tidak ada!!! (no rows effected) ")
		hasil.Pesan = "Tidak ada record Menu"
		hasil.Data = nil
		return hasil
	}

	menu := tblmnu.Data.([]map[string]interface{})
	//Begin........................................
	var results []map[string]interface{}

	SO_Class.Log.Println(false, "================================")
	SO_Class.Log.Println(false, menu)
	SO_Class.Log.Println(false, "================================")

	var iyMenu = make(map[string]int64)

	var Pid int64
	Pid = 0
	for key, element := range menu {
		SO_Class.Log.Println(false, "Key:", key, "=>", "Element:", element)
		SO_Class.Log.Println(false, element["tmmenuiy"])
		SO_Class.Log.Println(false, element["tmnomr"].(string))

		iyMenu[element["tmnomr"].(string)] = element["tmmenuiy"].(int64)

		n := ""
		Nomr := element["tmnomr"].(string)
		rows := make(map[string]interface{})
		if len(Nomr) == 2 {
			Pid = 0
		} else {

			vNomr := []rune(Nomr)
			if element["tmurlw"].(string) == "" {
				if len(Nomr) > 2 {
					n = string(vNomr[0:(len(Nomr) - 2)])
					Pid = iyMenu[n]
				} else {
					Pid = 0
					n = ""
				}
			} else {
				if n != string(vNomr[0:(len(Nomr)-2)]) {
					n = string(vNomr[0:(len(Nomr) - 2)])
					Pid = iyMenu[n]
				}
			}
		}

		// iId := UbahStoI(element["tmmenuiy"].(string))
		// iPid := UbahStoI(Pid)
		iId := element["tmmenuiy"].(int64)
		// iId := strconv.FormatInt(element["tmmenuiy"].(int64), 10)
		iPid := Pid

		rows["id"] = iId
		rows["name"] = element["tmmenu"].(string)
		rows["pId"] = iPid
		rows["mAkses"] = element["tmaces"].(string)
		rows["uAkses"] = element["taaces"].(string)
		rows["idMenu"] = element["tmurlw"].(string)
		rows["scutMenu"] = element["tmscut"].(string)
		rows["jsonMenu"] = element["tmjson"].(string)
		rows["nomor"] = element["tmnomr"].(string)
		rows["tipeMenu"] = element["tumntp"].(string)
		// rows["idMenu"] = StrAcak(element["tmurlw"].(string), 2)
		rows["judul"] = element["tmtext"].(string)
		rows["expanded"] = false
		rows["disabled"] = true
		rows["disabledCls"] = "x-item-disabled"

		rows["icon"] = ""
		if element["tmaces"].(string) != "" {
			index := SO_Class.Strings.Index(element["taaces"].(string), "V")
			if index != -1 {
				rows["icon"] = "resources/asset/menu_tree_add.png"
				rows["disabled"] = false
			} else {
				rows["icon"] = "resources/asset/menu_tree_delete.png"
				rows["disabled"] = true
			}
		}

		// SO_Class.Log.Println(false, rows)
		results = append(results, rows)

	}

	SO_Class.Log.Println(false, results)
	SO_Class.Log.Println(false, "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")

	// hasil.Data = buildMenuTrees(results)
	hasil.Data = buildMenuTrees(results)
	//End........................................

	hasil.Sukses = true
	hasil.Pesan = "Load Menu Tree Berhasil"
	return hasil
}
func buildMenuTrees(dataset interface{}) map[string]interface{} {
	currentTime := Form.GetCurrentTime()
	results := make([]interface{}, 0)
	c := map[string]interface{}{
		"id":          0,
		"name":        "root",
		"mAkses":      "",
		"uAkses":      "",
		"idMenu":      "",
		"nomor":       "",
		"jsonMenu":    "",
		"scutMenu":    "",
		"tipeMenu":    "",
		"text":        "",
		"tgl":         currentTime.Format("2006-01-02 3:4:5"),
		"pId":         0,
		"expanded":    false,
		"disabled":    true,
		"disabledCls": "x-item-disabled",
		"children":    buildMenuTree(dataset, 0),
	}
	results = append(results, c)
	return results[0].(map[string]interface{})
}
func buildMenuTree(dataset interface{}, p int64) []interface{} {
	results := make([]interface{}, 0)
	currentTime := Form.GetCurrentTime()
	vDataSet := dataset.([]map[string]interface{})

	for _, element := range vDataSet {
		// fmt.Println("mana........ ", i, " - ", element)
		pID := element["pId"].(int64)
		if p == pID {
			c := map[string]interface{}{
				"id":          element["id"].(int64),
				"name":        element["name"].(string),
				"mAkses":      element["mAkses"].(string),
				"uAkses":      element["uAkses"].(string),
				"idMenu":      element["idMenu"].(string),
				"nomor":       element["nomor"].(string),
				"jsonMenu":    element["jsonMenu"].(string),
				"scutMenu":    element["scutMenu"].(string),
				"tipeMenu":    element["tipeMenu"].(string),
				"text":        element["judul"].(string),
				"expanded":    element["expanded"].(bool),
				"tgl":         currentTime.Format("2006-01-02 3:4:5"),
				"pId":         pID,
				"disabled":    element["disabled"].(bool),
				"disabledCls": element["disabledCls"].(string),
				"icon":        element["icon"].(string),
				"children":    buildMenuTree(dataset, element["id"].(int64)),
			}
			results = append(results, c)
		}
	}
	// log.Println(" MenuTree : ", results)
	return results
}

func (saya tblmnu) ClickMenu(c *gin.Context) (hasil SO_Class.Hasil) {

	hasil.Sukses = true
	hasil.Pesan = "sukses"
	hasil.Data = nil

	exec := Form.ExecQueryMultiple(c, c.Param("username"), func(tx Transaction) TransactionResult {
		var tr TransactionResult
		tr.Error = nil
		tr.Sintax = ""

		currentTime := Form.GetCurrentTime()
		sqlstm := SO_Class.Fmt.Sprint(`
				update tblmnu set
					tmlsby = '` + c.Param("username") + `',
					tmlsdt = '` + currentTime.Format("2006-01-02 3:4:5") + `',
					tmusct = coalesce(tmusct,0)+1,
					tmsrce = '` + c.Param("source") + `'
				where tmmenuiy = '` + c.Param("menuiy") + `';		

				update tbluam set
					talsdt = '` + currentTime.Format("2006-01-02 3:4:5") + `',
					tausct = coalesce(tausct,0)+1,
					tasrce = '` + c.Param("source") + `'
				where tamenuiy = '` + c.Param("menuiy") + `'
				and tauseriy in (
					select 
						tuuseriy 
					from tblusr 
					where tuuser = '` + c.Param("username") + `'				
				);
		`)
		SO_Class.Log.Println(false, " tblmnu clikmenu ", sqlstm)

		_, err := Form.Execute(tx, &tr, c.Param("username"), sqlstm)
		if err != nil {
			tr.Error = err
			return tr
		}

		return tr
	})

	if exec != nil {
		hasil.Sukses = false
		hasil.Pesan = SO_Class.Fmt.Sprint("Gagal", exec.Error())
	} else {
		hasil.Sukses = true
		hasil.Pesan = SO_Class.Fmt.Sprint("Sukses", "")
	}

	return hasil
}
func init() {
	Form.Add("TBLMNU", TBLMNU)
	if Form.logPrintInitFlag {
		SO_Class.Log.Println(true, "Masuk form-tblmnu-init()")
	}
}

// Exported instance
var TBLMNU tblmnu
