package SO_Form

import (
	SO_Class "SOApp_GO/class"
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type ivhead struct {
	table struct {
		IHIHNOIY string `tipe:"1" json:"ID" key:"true" required:"true"`
		IHTRNO   string `tipe:"1" json:"Inventory No" required:"true"`
		IHTRDT   string `tipe:"1" json:"Inventory Date" required:"true"`
		IHTYPE   string `tipe:"1" json:"Inventory Type"`
		IHUSRM   string `tipe:"2" json:"Internal Use Remark"`
		IHREMK   string `tipe:"2" json:"Remark"`
		IHCSDT   string `tipe:"X" `
	}
}

func (saya ivhead) Save(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.CetakKunci(true, c, "Masuk IVHEAD-Save()")
	hasil.Sukses = false
	hasil.Pesan = "Masuk IVHEAD - Save"
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
				c.Param("ihcsdt"),
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

func (saya ivhead) StpSave(tx Transaction, c *gin.Context,
	mode string, userName string, source string, csdt string, table ivhead,
) (tr TransactionResult) {

	data := table.table
	tr.Sintax = " -- ivhead.StpSave() "

	_, errBFCS := Form.CheckRecord_BFCS(ParamBFCS{
		Tx:       tx,
		Mode:     mode,
		Table:    "ivhead",
		KeyField: "ihihnoiy",
		KeyValue: data.IHIHNOIY,
		CSDT:     csdt,
	})
	if errBFCS != nil {
		tr.Error = errBFCS
		return tr
	}

	var err error

	if mode == "A" {
		data.IHIHNOIY = Form.GetTBLNOR(tx, userName, "ivhead")
		data.IHTRNO, err = Form.GenerateAutoNo(tx, userName, "INV", "INV",
			SO_Class.Fungsi.Left(data.IHTRDT, 4),
			SO_Class.Fungsi.Mid(data.IHTRDT, 4, 2),
			"A", 6)
		if err != nil {
			tr.Error = err
			return tr
		}
		if data.IHTRNO == "" {
			tr.Error = SO_Class.Fmt.Errorf("Generate Auto No Failed")
			return tr
		}
	}

	sqlstm := Form.GetSintaxSQL_IUD(ParamIUD{
		TableName:  "ivhead",
		UserName:   userName,
		Source:     source,
		Mode:       mode,
		StructAnda: data,
	})
	SO_Class.Log.CetakKunci(true, c, " ivhead.stpSave ", sqlstm)

	_, err = Form.Execute(tx, &tr, userName, sqlstm)
	if err != nil {
		tr.Error = err
		return tr
	}

	var dataGrid map[string][]map[string]interface{}

	errDataGrid := json.Unmarshal([]byte(c.Param("GRID")), &dataGrid)
	if errDataGrid != nil {
		tr.Error = errDataGrid
		return tr
	}

	// begin insert detail
	var ivline ivline
	for no, item := range dataGrid["Grid2"] {
		err_ivline := SO_Class.Fungsi.BindDataToStruct(item, &ivline.table)
		if err_ivline != nil {
			tr.Error = err_ivline
			return
		}
		ivline.table.ILIHNOIY = data.IHIHNOIY
		ivline.table.ILILNO = SO_Class.Fmt.Sprint(no + 1)
		tr = ivline.StpSave(
			tx, c,
			c.Param("Mode"),
			c.Param("username"),
			c.Param("source"),
			c.Param("ihcsdt"),
			ivline)
		if tr.Error != nil {
			return
		}
	}
	// end insert detail

	// tr.Error = SO_Class.Fmt.Errorf("CobaCoba")

	return tr
}

func (saya ivhead) LoadGrid(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.CetakKunci(false, c, "Masuk IVHEAD-LoadGrid()")
	if c.Param("sort") == "" {
		param := gin.Param{Key: "sort", Value: `[{"property": "ihtrno","direction": "asc"}]`}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select
				ihihnoiy
				, ihtrno, ihtrdt::date ihtrdt, ihtype
				, ihremk, ihusrm
				`, Form.GetDefaultField("ih"), ` 
			from ivhead
			where 1 = 1
		`)
	SO_Class.Log.CetakKunci(false, c, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "ihihnoiy",
		columns: []Kolum{
			{"ihtrno", KolumProperty{name: "Transaction No", noHideable: true}},
			{"ihtrdt", KolumProperty{name: "Inventory Date"}},
			{"ihtype", KolumProperty{name: "Inventory Type"}},
			{"ihremk", KolumProperty{name: "Remark"}},
			{"ihusrm", KolumProperty{name: "User Remark"}},
		},
		defaultField: true,
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}
func (saya ivhead) FillForm(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.CetakKunci(false, c, "Masuk IVHEAD-FillForm()")
	key := c.Param("ihihnoiy")
	sqlstm := SO_Class.Fmt.Sprint(`
		select ivhead.*
			, ivtp.tssycd as ihtype_tssycd, ivtp.tssynm as ihtype_tssynm 
		from ivhead 
		left join tblsys ivtp on ivtp.tsdscd = 'IVHEAD_TYPE' and ivtp.tssycd = ihtype
		where ihihnoiy = '`, key, `'
	`)
	SO_Class.Log.CetakKunci(true, c, sqlstm)
	hasil = Form.GetRs(c, sqlstm)
	return hasil
}

func (saya ivhead) LoadFormObject(c *gin.Context) (hasil SO_Class.Hasil) {
	hasil.Sukses = false
	hasil.Pesan = ""
	hasil.Data = nil

	frmId := c.Param("FrmId")
	menuId := SO_Class.Strings.ToUpper("ivhead")
	frmDetail := "Win" + frmId

	result := FormObject{
		Prefix: "ih",
		Frame1: Form.CrtObj(ObjGrd{
			Id: "Grid1", FrmId: frmId, MenuId: menuId,
			Title:      "Inventory",
			Controller: SO_Class.Strings.ToUpper("ivhead"),
			Method:     "LoadGrid",
		}),
		Frame2: []interface{}{
			Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
				Id: "PanelX", TipeVH: "H",
				Items: []interface{}{
					Form.CrtObj(ObjTab{
						FrmId: frmId, MenuId: menuId, Id: "TabPanelA",
						Items: []interface{}{
							saya.LoadObjectHeader(c, frmId, menuId),
							saya.LoadObjectDetail(c, frmId, menuId, frmDetail),
						},
					}),
				},
			}),
		},
		SOFn: SO_Class.Fmt.Sprint(`{
			formAction: {
				add: function ({form}) {			
					A7App.SO.SetGridBtnDisableEnable('`, frmId, `', 'Grid2', 'AED');
				},
				edit: function ({form}) {			
					A7App.SO.SetGridBtnDisableEnable('`, frmId, `', 'Grid2', 'AED');
					form.SOFn.fljLoadGrid2();				
				},
				view: function ({form}) {
					A7App.SO.SetGridBtnDisableEnable('`, frmId, `', 'Grid2', 'V');
					form.SOFn.fljLoadGrid2();				
				},
			},			
			fljLoadGrid2: function () {				
				var ihnoiy = Ext.getCmp('Frm`+frmId+`Obj`+`ihihnoiy').getValue();
				var grid2 = Ext.getCmp('Frm`+frmId+`Grd'+'Grid2');
					grid2.soObj_sqlCondition = " and ilihnoiy = '"+ihnoiy+"' ";
					grid2.loadGrid();		
			},
			fljHitungTotal: function () {
				let qtys = Ext.getCmp('Frm`+frmDetail+`Obj`+`ilqtys').getValue();
				let harg = Ext.getCmp('Frm`+frmDetail+`Obj`+`ilharg').getValue();
				let totl = 0;
					totl = qtys * harg;
				Ext.getCmp('Frm`+frmDetail+`Obj`+`iltotl').setNilai(totl);
			},
		}`),
	}

	hasil.Sukses = true
	hasil.Pesan = "sukses"
	hasil.Data = result
	return hasil
}

func (saya ivhead) LoadObjectHeader(c *gin.Context, frmId string, menuId string) map[string]interface{} {
	return Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
		Id: "PanelHeader", LabelWidth: 180, Title: "Data",
		Items: []interface{}{
			Form.CrtObj(ObjTxt{Mode: "3", FrmId: frmId, MenuId: menuId,
				Id: "ihihnoiy", Name: "IY", Hidden: true,
			}),
			Form.CrtObj(ObjTxt{Mode: "3", FrmId: frmId, MenuId: menuId,
				Id: "ihtrno", Name: "Inventory No",
			}),
			Form.CrtObj(ObjDtp{Mode: "2", FrmId: frmId, MenuId: menuId,
				Id: "ihtrdt", Name: "Inventory Date", AllowBlank: false,
			}),
			Form.CrtObj(ObjPop{Mode: "2", FrmId: frmId, MenuId: menuId,
				Controller: SO_Class.Strings.ToUpper("tblsys"), Method: "LoadGridPopUp",
				Id: "ihtype", Name: "Inventory Type", AllowBlank: false,
				PopCode: "tssycd", PopDesc: "tssynm",
				PopCodeFf: "ihtype_tssycd", PopDescFf: "ihtype_tssynm",
				SqlCondition: " and tsdscd = 'IVHEAD_TYPE' ",
			}),
			Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmId, MenuId: menuId,
				Id: "ihremk", Name: "Remark",
			}),
			Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmId, MenuId: menuId,
				Id: "ihusrm", Name: "Internal Use Remark",
			}),
		},
	})
}

func (saya ivhead) LoadObjectDetail(c *gin.Context, frmId string, menuId string, frmDetail string) map[string]interface{} {
	return Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
		Id: "PanelDetail", LabelWidth: 150, Title: "Detail",
		Items: []interface{}{
			Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
				Id: "PanelGrid", //Margin: "0 0 0 0",
				Items: []interface{}{
					Form.CrtObj(ObjGrd{FrmId: frmId, MenuId: menuId,
						Controller: SO_Class.Strings.ToUpper("ivhead"),
						Method:     "LoadGrid2", NoPaging: true,
						Id: "Grid2", Width: 800, Height: 350,
						Title:  "Detail Inventory",
						Action: "VAED", ShowGridAdvSearch: false,
						SqlCondition: ` and 1 = 0 `,
						WindowWidth:  650, WindowHeight: 450,
						// OnActionGridSubmit: saya.onActionGrid2Submit(frmId),
						DetailItems: []interface{}{
							Form.CrtObj(ObjPnl{FrmId: frmDetail, MenuId: menuId,
								Id: "PanelA", LabelWidth: 150,
								Items: []interface{}{
									Form.CrtObj(ObjPop{Mode: "2", FrmId: frmDetail, MenuId: menuId,
										Id: "ilitnoiy", PopCode: "mmitno", PopDesc: "mmitds",
										Name: "Item No", Controller: SO_Class.Strings.ToUpper("mitmas"), Method: "LoadGrid",
										SqlCondition: ` and mmdpfg = '1' `,
										OnAjax: `
											unms = Ext.getCmp('Frm` + frmDetail + `Obj` + `mmunms'); 
											unms.setNilai(data["mmunms"]);
										`,
									}),
									Form.CrtObj(ObjCnt{FrmId: "Win" + frmDetail, MenuId: menuId,
										Id: "ContainerU",
										Items: []interface{}{
											Form.CrtObj(ObjNum{Mode: "1", FrmId: frmDetail, MenuId: menuId,
												Id: "ilqtys", Name: "Qty", Width: 100,
											}),
											Form.CrtObj(ObjTxt{Mode: "3", FrmId: frmDetail, MenuId: menuId,
												Id: "mmunms", Name: "UM", LabelWidth: 50, Width: 100,
											}),
										},
									}),
									Form.CrtObj(ObjNum{Mode: "1", FrmId: frmDetail, MenuId: menuId,
										Id: "ilharg", Name: "Price", Width: 100,
										OnAjax: `
											form = Ext.getCmp('Frm` + frmId + `'); 
											form.SOFn.fljHitungTotal();
										`,
									}),
									Form.CrtObj(ObjNum{Mode: "3", FrmId: frmDetail, MenuId: menuId,
										Id: "iltotl", Name: "Total", Width: 100,
									}),
									Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmDetail, MenuId: menuId,
										Id: "ilremk", Name: "Remark",
									}),
									Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmDetail, MenuId: menuId,
										Id: "ilusrm", Name: "Internal Use Remark",
									}),
								},
							}),
						},
						OnActionGridSubmit: `
							switch (eBtn.soMode) {
								case "A":
									A7App.SO.insertGridDetail({frmID: '` + frmId + `', gridID: 'Grid2'});
									break;
								case "E":
									A7App.SO.updateGridDetail({frmID: '` + frmId + `', gridID: 'Grid2'});
									break;
							}
						`,
					}),
				},
			}),
		},
	})
}

func (saya ivhead) LoadGrid2(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.CetakKunci(false, c, "Masuk IVHEAD-LoadGrid2()")
	if c.Param("sort") == "" {
		param := gin.Param{Key: "sort", Value: `[{"property": "ililno","direction": "asc"}]`}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select
				ililnoiy, ilihnoiy
				, ililno 
				, ilitnoiy, mmitno, mmitds, mmunms
				, ilqtys::dec(24,0) ilqtys
				, ilharg::dec(24,0) ilharg
				, iltotl::dec(24,0) iltotl
				, ilremk, ilusrm
				`, Form.GetDefaultField("il"), ` 
			from ivline
			left join mitmas on mmitnoiy = ilitnoiy
			where 1 = 1
		`)
	SO_Class.Log.CetakKunci(false, c, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "ililnoiy",
		columns: []Kolum{
			// {"ihtrno", KolumProperty{name: "Transaction No", noHideable: true}},
			{"ilihnoiy", KolumProperty{name: "IY Head", hidden: true}},
			{"ililno", KolumProperty{name: "Line No", hidden: true}},
			{"ilitnoiy", KolumProperty{name: "Item Iy", hidden: true}},
			{"mmitno", KolumProperty{name: "Item No"}},
			{"mmitds", KolumProperty{name: "Item Description"}},
			{"mmunms", KolumProperty{name: "UM"}},
			{"ilqtys", KolumProperty{name: "Qty"}},
			{"ilharg", KolumProperty{name: "Harga"}},
			{"iltotl", KolumProperty{name: "Total"}},
			{"ilremk", KolumProperty{name: "Remark"}},
			{"ilusrm", KolumProperty{name: "User Remark"}},
		},
		defaultField: false,
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}

func init() {
	Form.Add("IVHEAD", IVHEAD)
	if Form.logPrintInitFlag {
		SO_Class.Log.Println(true, "Masuk form-ivhead-init()")
	}
}

// Exported instance
var IVHEAD ivhead
