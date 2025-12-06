package SO_Form

import (
	SO_Class "SOApp_GO/class"

	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type tbluah struct {
	table struct {
		// TBNOMRIY string `tipe:"X" json:"ID" key:"true" required:"true"`
		TBUSERIY string `tipe:"2" json:"User ID" required:"true"`
		TBMENUIY string `tipe:"2" json:"Menu ID" required:"true"`
		TBACES   string `tipe:"2" json:"Access Menu User" required:"true"`
		TBACEC   string `tipe:"2" json:"Access Menu Creator" required:"true"`
		TBUSRM   string `tipe:"2" json:"User Remark"`
		TBREMK   string `tipe:"2" json:"Remark"`
	}
}

func (saya tbluah) Save(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLUAH-Save()")
	hasil.Sukses = false
	hasil.Pesan = "Masuk TBLUAH - Save"
	hasil.Data = nil
	hasil.Kode = ""

	// err := SO_Class.Fungsi.BindParamsToStruct(c, &saya.table)
	// if err != nil {
	// 	hasil.Sukses = false
	// 	hasil.Pesan = err.Error()
	// 	return
	// }

	AllGrid := SO_Class.Fungsi.ConvertToJSON(c.Param("GRID"))
	SO_Class.Log.Println(false, "", AllGrid)
	SO_Class.Log.Println(false, "", AllGrid.Data.(map[string]interface{})["Grid2"])
	// return hasil

	exec := Form.ExecQueryMultiple(c, c.Param("username"),
		func(tx Transaction) TransactionResult {
			var tr TransactionResult
			tr.Error = nil
			tr.Sintax = ""

			grid2 := AllGrid.Data.(map[string]interface{})["Grid2"].([]interface{})

			for key, value := range grid2 {
				SO_Class.Log.Println(true, "", key, value)
				row := value.(map[string]interface{})
				saya.table.TBUSERIY = row["tbuseriy"].(string)
				saya.table.TBMENUIY = row["tbmenuiy"].(string)
				saya.table.TBACES = row["tbvalu"].(string) + row["tbnewv"].(string)
				saya.table.TBACEC = row["tbacec"].(string)
				// saya.table.TBUSRM = row["tbusrm"]
				// saya.table.TBREMK = row["tbremk"]

				tr = saya.StpSave(
					tx,
					c.Param("Mode"),
					c.Param("username"),
					c.Param("source"),
					c.Param("tbcsdt"),
					saya)

			}

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

func (saya tbluah) StpSave(tx Transaction,
	mode string, userName string, source string, csdt string, table tbluah,
) (tr TransactionResult) {

	data := table.table
	tr.Sintax = " -- tbluah.StpSave() "

	sqlstm := Form.GetSintaxSQL_IUD(ParamIUD{
		TableName:  "tbluah",
		UserName:   userName,
		Source:     source,
		Mode:       mode,
		StructAnda: data,
	})
	SO_Class.Log.Println(false, " tbluah save ", sqlstm)

	_mode := data.TBACES[:1]
	_value := data.TBACES[1:]

	sqlstm = SO_Class.Fmt.Sprint(
		sqlstm,
		` update tbluam  
			set taaces = case when '+' = '`, _value, `' 
						then concat(replace(taaces,'`, _mode, `',''),'`, _mode, `') 
						else replace(taaces,'`, _mode, `','') end
			, tachid = '`, userName, `'
			, tachdt = '`, Form.GetCurrentTime().Format("2006-01-02 3:4:5"), `'
			, tachno = coalesce(tachno,0)+1
			, tasrce = '`, source, `'
			, tacsid = '`, userName, `'
			, tacsdt = '`, Form.GetCurrentTime().Format("2006-01-02 3:4:5"), `'
			where tamenuiy = '`, data.TBMENUIY, `'
			and tauseriy = '`, data.TBUSERIY, `';`)

	SO_Class.Log.Println(true, " tbluam save ", sqlstm)

	_, err := Form.Execute(tx, &tr, userName, sqlstm)
	if err != nil {
		tr.Error = err
		return tr
	}

	return tr
}

func (saya tbluah) LoadGrid(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLUAH-LoadGrid()")
	if c.Param("sort") == "" {
		param := gin.Param{Key: "sort", Value: `[{"property": "tbcsdt","direction": "desc"}]`}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select
				tbnomriy, tuuser, tmscut, tmmenu, tbaces 
				, tbremk, tbusrm
				`, Form.GetDefaultField("tb"), ` 
			from tbluah
			left join tblmnu on tmmenuiy = tbmenuiy
			left join tblusr on tuuseriy = tbuseriy
			where 1 = 1
		`)
	SO_Class.Log.Println(false, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "tbnomriy",
		columns: []Kolum{
			{"tbnomriy", KolumProperty{name: "No", hidden: true}},
			{"tuuser", KolumProperty{name: "User"}},
			{"tmscut", KolumProperty{name: "Menu Code"}},
			{"tmmenu", KolumProperty{name: "Menu"}},
			{"tbaces", KolumProperty{name: "Akses"}},
			{"tbremk", KolumProperty{name: "Remark"}},
			{"tbusrm", KolumProperty{name: "User Remark"}},
		},
		defaultField: true,
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}
func (saya tbluah) FillForm(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLUAH-FillForm()")
	key := c.Param("tmmenuiy")
	sqlstm := SO_Class.Fmt.Sprint("select * from tbluah where tmmenuiy = '", key, "'")
	SO_Class.Log.Println(true, sqlstm)
	hasil = Form.GetRs(c, sqlstm)
	return hasil
}

func (saya tbluah) LoadFormObject(c *gin.Context) (hasil SO_Class.Hasil) {
	hasil.Sukses = false
	hasil.Pesan = ""
	hasil.Data = nil

	frmId := c.Param("FrmId")
	menuId := SO_Class.Strings.ToUpper("tbluah")

	result := FormObject{
		Prefix: "tm",
		Frame1: Form.CrtObj(ObjGrd{
			Id: "Grid1", FrmId: frmId, MenuId: menuId,
			Title:      "Table User Hak Akses Menu",
			Controller: SO_Class.Strings.ToUpper("tbluah"),
			Method:     "LoadGrid",
		}),
		Frame2: []interface{}{
			Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
				Id: "PanelX", LabelWidth: 150,
				Items: []interface{}{
					Form.CrtObj(ObjPop{
						FrmId: frmId, MenuId: menuId, Mode: "2", Id: "tbuseriy", PopCode: "tuuser", PopDesc: "tuname",
						Name: "Username", Controller: "TBLUSR", Method: "LoadGridUserLevel",
						SqlCondition: ` and tudpfg = '1' `,
					}),
					Form.CrtObj(ObjPnl{FrmId: frmId, MenuId: menuId,
						Id: "PanelGrid", Margin: "0 0 0 150",
						Items: []interface{}{
							Form.CrtObj(ObjGrd{FrmId: frmId, MenuId: menuId,
								Controller: SO_Class.Strings.ToUpper("tbluah"),
								Method:     "LoadGrid2", NoPaging: true,
								Id: "Grid2", Width: 800, Height: 350,
								Title:  "Table User Hak Akses Menu",
								Action: "AD", ShowGridAdvSearch: false,
								WindowWidth: 920, WindowHeight: 550,
								OnActionGridSubmit: saya.onActionGrid2Submit(frmId),
								DetailItems: []interface{}{
									Form.CrtObj(ObjPnl{FrmId: "Win" + frmId, MenuId: menuId,
										Id: "PanelA", LabelWidth: 150,
										Items: []interface{}{
											Form.CrtObj(ObjCnt{FrmId: "Win" + frmId, MenuId: menuId,
												Id: "ContainerU",
												Items: []interface{}{
													Form.CrtObj(ObjCmb{C: c, Mode: "1", FrmId: "Win" + frmId, MenuId: menuId,
														Id: "valu", Name: "Akses", Table: "AUTH_FG", Value: "-",
													}),
													Form.CrtObj(ObjBtn{Mode: "1", FrmId: "Win" + frmId, MenuId: menuId,
														Id: "BtnValu", Text: "Go",
														Width: 100, Scale: "small", Margin: "0 0 0 5",
														OnKlik: saya.onKlikBtnValu(frmId),
													}),
												},
											}),
											Form.CrtObj(ObjGrd{
												Id: "Grid3", FrmId: "Win" + frmId, MenuId: menuId,
												Controller: SO_Class.Strings.ToUpper("tbluah"), Method: "LoadGrid3",
												Height: 400, Width: 895, ShowGridAdvSearch: false,
												OnBeforeLoadGrid: saya.onBeforeLoadGrid3(frmId),
												OnAfterLoadGrid:  saya.onAfterLoadGrid3(frmId),
											}),
										},
									}),
								},
							}),
						},
					}),

					Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmId, MenuId: menuId,
						Id: "tbusrm", Name: "Internal Use Remark",
					}),
					Form.CrtObj(ObjRmk{Mode: "0", FrmId: frmId, MenuId: menuId,
						Id: "tbremk", Name: "Remark",
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

func (saya tbluah) onKlikBtnValu(frmId string) string {
	return `
		var akses = Ext.getCmp('FrmWin` + frmId + `Obj` + `valu').getValue();

		Ext.MessageBox.confirm(
			'Reset Password', 
			'Are you sure you want to set all access menu to ' + akses + '?', 
			function (btn) {
				if (btn === 'yes') {

					var G3 = Ext.getCmp('FrmWin` + frmId + `GrdGrid3')
									
					var data = G3.store.query()
					data.each(
						function(record){
							if (record.data.taflag === "Y") {	
								record.set('tanewv', akses);
							}
						}
					)

					Ext.MessageBox.alert('Success', 'Done');
				}
			}
		);	
	`
}

func (saya tbluah) onBeforeLoadGrid3(frmId string) string {
	return `
		var inArray = []
		var G2 = Ext.getCmp('Frm` + frmId + `GrdGrid2')
		var data = G2.store.query()
		data.each(
			function(record){
				var hasil = record.data.tbmenuiy + '' + record.data.tbvalu
				inArray.push(hasil)
			}
		)
		// console.log('inArray', inArray)
		// console.log('inArray', "'"+inArray.join("','")+"'")

		var user = Ext.getCmp('Frm` + frmId + `Objtbuseriy_tuuser').value
		saya.soObj_sqlCondition  = " and tuuser = '" + user + "' "
		saya.soObj_sqlCondition += " and concat(tamenuiy,tmvalu) not in ( '" + inArray.join("','") + "') "
	`
}

func (saya tbluah) onActionGrid2Submit(frmId string) string {
	return `

		// console.log('eBtn', eBtn)
		// console.log('eBtn.soMode', eBtn.soMode)

		switch (eBtn.soMode) {
			case "A":

				var G2 = Ext.getCmp('Frm` + frmId + `GrdGrid2')
				var G3 = Ext.getCmp('FrmWin` + frmId + `GrdGrid3')
				
				var data = G3.store.query()
				data.each(
					function(record){
						
						var hasil = {}
						if(record.data.taoldv != record.data.tanewv) {
																	
							hasil = {
								tbuseriy: record.data.tauseriy,
								tuuser: record.data.tuuser,
								tbmenuiy: record.data.tamenuiy,
								tmscut: record.data.tmscut,
								tmmenu: record.data.tmmenu,
								tbacec: record.data.tmaces,
								tbvalu: record.data.tmvalu,
								tboldv: record.data.taoldv,
								tbnewv: record.data.tanewv,
								tbremk: "",
							}							

							G2.getStore().add(hasil);
							
						}

					}
				)

				var w = this.up('window')
				w.close()
				
				break
			case "D":
				break
		}
	
	`
}

func (saya tbluah) onAfterLoadGrid3(frmId string) string {
	return `																					
		var tanewv = saya.query('[dataIndex=tanewv]')[0];
		// console.log('ssssssssssssssssssssss', tanewv);		
		var cmbData = Ext.getCmp('FrmWin` + frmId + `Obj` + `valu');									
		// console.log('cmbData', cmbData)												
		// console.log('cmbData.getStore().getData().items', cmbData.getStore().getData().items);	
		
		var records = cmbData.getStore().getData().items;
		var dataArray = []
		Ext.Array.each(records, function(record) {
			// Access fields or properties of the record as needed
			dataArray.push({
				'Cod':record.get('kode'), 
				'Dsc':record.get('nama')
			})
		});									
		// console.log('dataArray', dataArray)	

		tanewv.setEditor({
			xtype: 'combo',  
			autoSelect: true,
			forceSelection: true, 
			triggerAction: 'all',
			allowBlank: true,
			autoSelect: true,
			// value: myStoreYESNO.getAt('1').get('Cod'),
			queryMode: 'local',
			valueField: 'Cod',
			displayField:'Dsc',
			store: {
				fields: ['Cod', 'Dsc'],
				data: dataArray
			}
		});											


		Ext.onReady(function() {
			Ext.suspendLayouts();										
				saya.on({					
					beforeedit: function (editor, context) {
						// console.log("editor", editor, "context", context)
						if (context.field == "tanewv") {
							if (context.record.data.taflag == "N") {
								return false
							}
						}
						return true
					},		
				});		
			Ext.resumeLayouts(true);
		});

		saya.getView().refresh();
	`
}

func (saya tbluah) LoadGrid2(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLUAH-LoadGrid2()")
	tbuseriy := c.Param("tbuseriy")
	if tbuseriy == "" {
		tbuseriy = "0"
	}
	if c.Param("sort") == "" {
		param := gin.Param{Key: "sort", Value: `[{"property": "tbcsdt","direction": "desc"}]`}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select
				tbnomriy, tbuseriy, tbmenuiy, 
				tuuser, tmscut, tmmenu, tbacec, 
				left(rtrim(tbaces),1)  tbvalu, 
				right(rtrim(tbaces),1) tboldv, 
				right(rtrim(tbaces),1) tbnewv, 
				tbremk, tbusrm
				`, Form.GetDefaultField("tb"), ` 
			from tbluah
			left join tblmnu on tmmenuiy = tbmenuiy
			left join tblusr on tuuseriy = tbuseriy
			where tbuseriy = '`, tbuseriy, `'
		`)
	SO_Class.Log.Println(false, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "tbnomriy",
		columns: []Kolum{
			{"tbnomriy", KolumProperty{name: "No", hidden: true}},
			{"tbuseriy", KolumProperty{name: "UserID", hidden: true}},
			{"tbmenuiy", KolumProperty{name: "MenuID", hidden: true}},
			{"tuuser", KolumProperty{name: "User", hidden: true}},
			{"tmscut", KolumProperty{name: "Menu Code"}},
			{"tmmenu", KolumProperty{name: "Menu"}},
			{"tbacec", KolumProperty{name: "Menu Akses", hidden: true}},
			{"tbvalu", KolumProperty{name: "Akses"}},
			{"tboldv", KolumProperty{name: "Akses Lama"}},
			{"tbnewv", KolumProperty{name: "Akses Baru"}},
			{"tbremk", KolumProperty{name: "Remark", hidden: true}},
			{"tbusrm", KolumProperty{name: "User Remark", hidden: true}},
		},
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}

func (saya tbluah) LoadGrid3(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLUAH-LoadGrid3()")

	if c.Param("sort") == "" {
		param := gin.Param{Key: "sort", Value: `[{"property": "tmscut_tmmenu","direction": "desc"}]`}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select 
				tauseriy, tamenuiy, tmscut_tmmenu, 
				tuuser, tmnomr, tmscut, tmmenu, tmaces, tmvalu, tssynm, 
				taoldv, tanewv, taflag 
			from (
				select 
					tmmenuiy, tmscut, tmnomr, tmmenu, tmaces, tmvalu,						
					concat('[', rtrim(tmscut), '] - ', tmmenu) as "tmscut_tmmenu"
				from tblmnu 
				CROSS  JOIN LATERAL unnest(regexp_split_to_array(rtrim(tmaces), '')) AS s(tmvalu)
				where tmaces <> ''
			) m
			left join (
				select 
					tanomriy, tauseriy, tamenuiy, taaces, tavalu, 
					case when strpos(taaces, tavalu) = 0 then '-' else '+' end taoldv,
					case when strpos(taaces, tavalu) = 0 then '-' else '+' end tanewv,
					case when strpos(l_aces, tavalu) = 0 then 'N' else 'Y' end taflag  
				from tbluam 
				left join tblmnu on tmmenuiy = tamenuiy					
				left join (
					select 
						tamenuiy as l_menuiy, taaces as l_aces 
					from tbluam 
					left join tblusr on tuuseriy = tauseriy 
					where tuuser = '` + c.Param("username") + `'
				) u on l_menuiy = tmmenuiy	
				CROSS  JOIN LATERAL unnest(regexp_split_to_array(rtrim(tmaces), '')) AS s(tavalu)
			) a on tamenuiy = tmmenuiy and tmvalu = tavalu 
			left join tblusr on tuuseriy = tauseriy
			left join tblsys on tsdscd = 'MODE' and tssycd = tmvalu
			where 1 = 1
		`)
	SO_Class.Log.Println(false, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "tmscut_tmmenu",
		columns: []Kolum{
			{"tauseriy", KolumProperty{name: "UserID", hidden: true}},
			{"tamenuiy", KolumProperty{name: "MenuID", hidden: true}},
			{"tuuser", KolumProperty{name: "Username"}},
			{"tmscut_tmmenu", KolumProperty{name: "Menu Code"}},
			{"tmnomr", KolumProperty{name: "No Urut", hidden: true}},
			{"tmscut", KolumProperty{name: "Short Cut"}},
			{"tmmenu", KolumProperty{name: "Menu", hidden: true}},
			{"tmaces", KolumProperty{name: "Menu Akses", hidden: true}},
			{"tmvalu", KolumProperty{name: "Value", hidden: true}},
			{"tssynm", KolumProperty{name: "Akses"}},
			{"taoldv", KolumProperty{name: "Akses Lama"}},
			{"tanewv", KolumProperty{name: "Akses Baru"}},
			{"taflag", KolumProperty{name: "Flag", hidden: true}},
		},
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}

func init() {
	Form.Add("TBLUAH", TBLUAH)
	if Form.logPrintInitFlag {
		SO_Class.Log.Println(true, "Masuk form-tbluah-init()")
	}
}

// Exported instance
var TBLUAH tbluah
