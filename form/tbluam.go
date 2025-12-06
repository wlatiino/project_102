package SO_Form

import (
	SO_Class "SOApp_GO/class"

	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type tbluam struct{}

func (saya tbluam) LoadGridTBLUSR(c *gin.Context) (hasil SO_Class.Hasil) {
	SO_Class.Log.Println(true, "Masuk TBLUAM-LoadGridTBLUSR()")
	var kondisi string
	kondisi = ""
	if c.Param("sqlCondition") == "" {
		kondisi = " and 1 = 0 "
	}

	if c.Param("sort") == "" {
		param := gin.Param{Key: "sort", Value: `[{"property": "tmnomr","direction": "asc"}]`}
		c.Params = append(c.Params, param)
	}
	sqlstm :=
		SO_Class.Fmt.Sprint(`
			select
				tanomriy, tmnomr, taaces 				
				, concat(repeat('.',length(tmnomr)*2),
						case when tmaces = '' then tmmenu 
						else concat('[', rtrim(tmscut), '] ', tmmenu) end
					) as "tamenu"
				, tausct, talsdt::bpchar talsdt
				, taremk, tausrm 
				`, Form.GetDefaultField("ta"), ` 
			from tbluam			
			left join tblmnu on tmmenuiy = tamenuiy
			where tadlfg = '0'`, kondisi, `
		`)
	SO_Class.Log.Println(false, sqlstm)
	hasil = Form.LoadGrid(ParamLoadGrid{
		c:      c,
		sqlstm: sqlstm,
		key:    "tanomriy",
		columns: []Kolum{
			{"tanomriy", KolumProperty{name: "ID", hidden: true}},
			{"tmnomr", KolumProperty{name: "No"}},
			{"tamenu", KolumProperty{name: "Menu"}},
			{"taaces", KolumProperty{name: "Access"}},
			{"tausct", KolumProperty{name: "Use Count"}},
			{"talsdt", KolumProperty{name: "Last Use Date"}},
			{"taremk", KolumProperty{name: "Remark"}},
			{"tausrm", KolumProperty{name: "User Remark"}},
		},
	})
	// hasil = Form.GetRs(c, sqlstm)
	// hasil = Form.GetRecordSet(c, sqlstm)
	return hasil
}

func init() {
	Form.Add("TBLUAM", TBLUAM)
	if Form.logPrintInitFlag {
		SO_Class.Log.Println(true, "Masuk form-tbluam-init()")
	}
}

// Exported instance
var TBLUAM tbluam
