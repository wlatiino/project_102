package SO_Form

import (
	SO_Class "SOApp_GO/class"
	SO_Module "SOApp_GO/module"
	SO_Object "SOApp_GO/soobj"
	"database/sql"
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// type identitasForm struct {
// 	rd SO_Module.AuthInterface
// 	tk SO_Module.TokenInterface
// }

type FormType struct {
	forms            map[string]interface{}
	logPrintInitFlag bool
	rd               SO_Module.AuthInterface
	tk               SO_Module.TokenInterface
}

type FormObject struct {
	Prefix  any
	Frame1  any
	Frame2  any
	Initial any
	SOFn    any
	Api     any
}

func init() {
	SO_Class.Log.Println(true, "Masuk form-form-init()")
}

func (f *FormType) BuatIdentitasForm(rd SO_Module.AuthInterface, tk SO_Module.TokenInterface) {
	f.rd = rd
	f.tk = tk
}

func NewFormType() *FormType {
	return &FormType{
		forms: make(map[string]interface{}),
	}
}

func (f *FormType) Add(name string, form interface{}) {

	if len(f.forms) == 0 {
		if SO_Class.Strings.ToUpper(os.Getenv("GIN_MODE")) == "DEV" {
			Form.SetLogPrintInitFlag(true)
		}
	}

	f.forms[name] = form

}

func (f *FormType) GetDbConnList(c *gin.Context) []map[string]interface{} {
	return SO_Module.Database.GetDbConnList()
}

func (f *FormType) HashPassword(password string) (string, error) {
	return SO_Module.Database.HashPassword(password)
}
func (f *FormType) VerifyPassword(password, hash string) bool {
	return SO_Module.Database.VerifyPassword(password, hash)
}

func (f *FormType) GetRs(c *gin.Context, sqlstm string) (hasil SO_Class.Hasil) {

	// t, _ := Form.tk.ExtractTokenMetadata(c)
	// SO_Class.Log.Println(false, t)
	// rs, jmlhRec, err := SO_Module.Database.GetRs(t.Database, sqlstm)
	db, _ := c.Get("globalDB")
	SO_Class.Log.Println(false, db)
	rs, jmlhRec, err := SO_Module.Database.GetRs(db.(string), sqlstm)
	if err != nil {
		hasil.Sukses = false
		hasil.Pesan = err.Error()
		hasil.Data = nil
		hasil.Kode = ""
		SO_Class.Log.Println(true, "Form GetRs Error", sqlstm, err.Error())
	} else {
		hasil.Sukses = true
		if jmlhRec == 0 {
			hasil.Pesan = "Record Kosong!!!"
			hasil.Data = nil
		} else {
			hasil.Data = rs
		}
	}

	return hasil

}

func (f *FormType) GetRecordSet(c *gin.Context, sqlstm string) (hasil SO_Class.Hasil) {

	// t, _ := Form.tk.ExtractTokenMetadata(c)
	// SO_Class.Log.Println(false, t)
	// rs, err := SO_Module.Database.GetRecordSet(t.Database, sqlstm)
	db, _ := c.Get("globalDB")
	SO_Class.Log.Println(false, db)
	rs, err := SO_Module.Database.GetRecordSet(db.(string), sqlstm)
	if err != nil {
		hasil.Sukses = false
		hasil.Pesan = err.Error()
		hasil.Data = nil
		hasil.Kode = ""
	} else {
		hasil.Sukses = true
		hasil.Data = rs
	}
	/*
		result :
			{
				"Sukses": true,
				"Pesan": "",
				"Kode": "",
				"Data": {
					"columns": [],
					"rows": []
				}
			}
	*/
	return hasil

}

//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//

type KolumProperty struct {
	name       string
	hidden     bool
	noHideable bool
}

type Kolum struct {
	Code  string
	Kolum KolumProperty
}

func (f *FormType) SetQueryPaging(c *gin.Context, queryPaging string) string {

	query := queryPaging
	if c.Param("tipeGrid") == "WithPaging" {
		start, err := strconv.Atoi(c.Param("start"))
		if err != nil {
			start = 0
		}

		limit, err := strconv.Atoi(c.Param("limit"))
		if err != nil {
			limit = 0
		}

		// offSet := start * limit
		// queryFinal = fmt.Sprint(query, queryWhere, querySort, " LIMIT ", limit, " OFFSET ", start)
		query = SO_Class.Fmt.Sprint(query, " limit ", limit, " offset ", start)
	}
	return query
}

func (f *FormType) SetQueryFilter(filter string) (string, []map[string]interface{}) {

	flagPrint := false

	filterParam := filter
	var filterMap []map[string]interface{}
	var queryFilter string
	queryFilter = ""
	if filterParam != "" {
		err := json.Unmarshal([]byte(filterParam), &filterMap)
		if err != nil {
			SO_Class.Log.Println(flagPrint, " form SetQueryFilter ", err)
			return "", nil
		}

		for key, values := range filterMap {
			SO_Class.Log.Println(flagPrint, " form SetQueryFilter ", "key : ", key, "values : ", values, "fMap : ", values["operator"])

			var nilai string
			var operator string
			// log.Println("type", reflect.TypeOf(values["value"]))
			_, okFloat := values["value"].(float64)
			if !okFloat {
				// 	nilai = strconv.FormatFloat(floatValue, 'f', 10, 64)
				// } else {
				switch values["operator"] {
				case "in", "ni":
					nilai = values["value"].(string)
				default:
					nilai = SO_Class.Strings.Replace(values["value"].(string), "'", "''", -1)
				}
			}
			lower := "lower"
			switch values["operator"] {
			case "eq":
				operator = "="
				// Parse the input date string into a time.Time object
				if !okFloat {
					parsedTime, err := time.Parse("01/02/2006", nilai)
					if err == nil {
						nilai = parsedTime.Format("2006-01-02")
					}
				}
				nilai = SO_Class.Fmt.Sprint("", nilai, "")
				lower = ""
			case "lt":
				operator = "<"
				// Parse the input date string into a time.Time object
				if !okFloat {
					parsedTime, err := time.Parse("01/02/2006", nilai)
					if err == nil {
						nilai = parsedTime.Format("2006-01-02")
					}
				}
				nilai = SO_Class.Fmt.Sprint("", nilai, "")
				lower = ""
			case "gt":
				operator = ">"
				// Parse the input date string into a time.Time object
				if !okFloat {
					parsedTime, err := time.Parse("01/02/2006", nilai)
					if err == nil {
						nilai = parsedTime.Format("2006-01-02")
					}
				}
				nilai = SO_Class.Fmt.Sprint("", nilai, "")
				lower = ""
			case "likeright":
				operator = "like"
				nilai = SO_Class.Fmt.Sprint("", nilai, "%")
			case "like":
				operator = "like"
				nilai = SO_Class.Fmt.Sprint("", nilai, "%") // Permintaan Pak Indra WA Group RndAppSys 2025 04 09
			case "in":
				operator = "in"
				nilai = SO_Class.Fmt.Sprint("(", nilai, ")")
			// case "inlike":
			// 	operator = "like"
			// 	nilai = SO_Class.Fmt.Sprint("%", nilai, "%")
			case "=":
				operator = "="
				nilai = SO_Class.Fmt.Sprint("", nilai, "")
			case "le":
				operator = "<="
				if !okFloat {
					parsedTime, err := time.Parse("01/02/2006", nilai)
					if err == nil {
						nilai = parsedTime.Format("2006-01-02")
					}
				}
				nilai = SO_Class.Fmt.Sprint("", nilai, "")
				lower = ""
			case "ge":
				operator = ">="
				if !okFloat {
					parsedTime, err := time.Parse("01/02/2006", nilai)
					if err == nil {
						nilai = parsedTime.Format("2006-01-02")
					}
				}
				nilai = SO_Class.Fmt.Sprint("", nilai, "")
				lower = ""
			case "ni":
				operator = "not in"
				nilai = SO_Class.Fmt.Sprint("(", nilai, ")")
			case "ne":
				operator = "<>"
				nilai = SO_Class.Fmt.Sprint("", nilai, "")
			case "nn":
				operator = "is not null"
				nilai = SO_Class.Fmt.Sprint("")
			case "nu":
				operator = "is null"
				nilai = SO_Class.Fmt.Sprint("")
			}

			if okFloat {
				queryFilter += SO_Class.Fmt.Sprint(
					" And ", lower, "(", values["property"].(string), ")",
					" ", operator,
					" '", values["value"], "'",
					" ",
				)
			} else {
				switch values["operator"] {
				case "in", "ni":
					queryFilter += SO_Class.Fmt.Sprint(
						" And ", lower, "(", values["property"].(string), "::varchar)",
						" ", operator,
						" ", SO_Class.Strings.ToLower(nilai), " ",
						" ",
					)
				default:
					queryFilter += SO_Class.Fmt.Sprint(
						" And ", lower, "(", values["property"].(string), "::varchar)",
						" ", operator,
						" '", SO_Class.Strings.ToLower(nilai), "'",
						" ",
					)
				}

			}

		}

	}
	return queryFilter, filterMap
}

func (f *FormType) SetQuerySort(sort string) (string, []map[string]interface{}) {

	flagPrint := false
	SO_Class.Log.Println(flagPrint, " form SetQuerySort ", sort)

	var sortMap []map[string]interface{}
	var querySort string
	querySort = ""
	if sort != "" {
		err := json.Unmarshal([]byte(sort), &sortMap)
		if err != nil {
			SO_Class.Log.Println(true, " form SetQuerySort ", err)
			return "", nil
		}
		var koma string
		koma = ""
		for key, values := range sortMap {
			SO_Class.Log.Println(flagPrint, " form SetQuerySort ", "key : ", key, "values : ", values, "sMap : ", values["property"])
			if querySort != "" {
				koma = ","
			}
			querySort += SO_Class.Fmt.Sprint(
				" ", koma,
				" ", values["property"],
				" ", values["direction"],
				" ",
			)
		}

		if querySort != "" {
			querySort = SO_Class.Fmt.Sprint(" Order By ", querySort)
		}
	}

	return querySort, sortMap
}

func (f *FormType) SetQueryCondition(c *gin.Context) string {

	var queryKondisi string
	queryKondisi = c.Param("sqlCondition")
	if c.Param("searchAll") != "" {
		var searchAll map[string]interface{}

		err := json.Unmarshal([]byte(c.Param("searchAll")), &searchAll)
		if err != nil {
			SO_Class.Log.Println(false, " form SetQueryCondition ", err)
		}

		searchValue := searchAll["searchValue"].(string)
		dataIndex := searchAll["dataIndex"].([]interface{})

		SO_Class.Fmt.Println(false, "Search Value:", searchValue)
		SO_Class.Fmt.Println(false, "Data Index:")

		colConditions := ""
		for i, index := range dataIndex {
			colConditions += " lower((" + index.(string) + ")::bpchar) like lower('%" + searchValue + "%')"
			if i < len(dataIndex)-1 {
				colConditions += " or "
			}
		}

		SO_Class.Fmt.Println(false, "Sform SetQueryCondition : ", colConditions)
		queryKondisi += " and (" + colConditions + ")"

	}

	return queryKondisi
}

func (f *FormType) GetColumnsProperties(cols []Kolum, columns []map[string]interface{}) ([]string, []map[string]interface{}) {

	arrFields := make(map[string]interface{})
	fields := make([]string, 0, len(cols))
	for _, k := range cols {
		fields = append(fields, k.Code)
		arrFields[k.Code] = k.Kolum
	}

	var arr = make(map[string]interface{})
	for i, kolumn := range columns {
		SO_Class.Log.Println(false, " form GetColumnsProperties ", "i : ", i, "kolumn : ", kolumn)
		arr[kolumn["code"].(string)] = kolumn
	}

	var columnProperties []map[string]interface{}
	columnProperties = append(columnProperties, map[string]interface{}{
		// "id":             "rowNumGrid",
		"xtype":          "rownumberer",
		"resizable":      false,
		"width":          45,
		"autoSizeColumn": true,
		"locked":         true,
		// "dataIndex":      "rowNumGrid",
	})
	for i, flds := range fields {
		SO_Class.Log.Println(false, " form GetColumnsProperties ", "i : ", i, "flds : ", flds)
		// arr[flds] = cols

		properties := map[string]interface{}{}
		properties["dataIndex"] = flds
		properties["name"] = arrFields[flds].(KolumProperty).name
		properties["header"] = arrFields[flds].(KolumProperty).name
		properties["hidden"] = arrFields[flds].(KolumProperty).hidden
		properties["hideable"] = !arrFields[flds].(KolumProperty).noHideable
		properties["lockable"] = true

		field, err := arr[flds].(map[string]interface{})
		if err {
			switch field["type"].(string) {
			case "DATE":
				properties["width"] = 200
				properties["type"] = "date"
				properties["filter"] = map[string]interface{}{
					"type": "date",
				}
				// properties["formatter"] = `date("d-M-Y h:i:s")`
				properties["autoSizeColumn"] = true
			case "TIMESTAMP":
				properties["width"] = 200
				properties["type"] = "datetime"
				properties["filter"] = map[string]interface{}{
					"type": "date",
				}
				// properties["formatter"] = `date("d-M-Y h:i:s")`
				properties["autoSizeColumn"] = true
			case "BPCHAR":
				properties["width"] = 100
				properties["type"] = "string"
				properties["filter"] = map[string]interface{}{
					"type": "string",
				}
				properties["autoSizeColumn"] = true
			case "NUMERIC":
				fallthrough
			case "INT8":
				fallthrough
			case "INT4":
				properties["width"] = 200
				properties["type"] = "number"
				properties["xtype"] = "numbercolumn"
				properties["filter"] = map[string]interface{}{
					"type": "number",
				}
				properties["align"] = "right"
				properties["autoSizeColumn"] = true

				scale := field["scale"].(int64)
				// precision := field["precision"].(int64)
				if scale <= 0 {
					properties["format"] = "0,000"
				} else if scale > 0 {
					if scale >= 10 {
						scale = 10
					}
					properties["format"] = "0,000." + SO_Class.Strings.Repeat("0", int(scale))
				}

			case "TEXT":
				properties["width"] = 200
				properties["type"] = "string"
				properties["filter"] = map[string]interface{}{
					"type": "string",
				}
				properties["autoSizeColumn"] = true
			default:
				properties["width"] = 100
				properties["type"] = "string"
				properties["filter"] = map[string]interface{}{
					"type": "string",
				}
				properties["autoSizeColumn"] = true
			}

			columnProperties = append(columnProperties, properties)
		} else {
			SO_Class.Log.Println(true, " form GetColumnsProperties ", "i : ", i, "flds : ", flds, " not found (check your sql statement field!!!)")
		}
	}
	return fields, columnProperties
}

type ParamLoadGrid struct {
	c            *gin.Context
	sqlstm       string
	key          string
	columns      []Kolum
	defaultField bool
}

func (f *FormType) LoadGrid(pLG ParamLoadGrid) (hasil SO_Class.Hasil) {

	// errorData := map[string]interface{}{
	// 	"sort":    []map[string]interface{}{},
	// 	"filter":  []map[string]interface{}{},
	// 	"fields":  []string{},
	// 	"columns": []map[string]interface{}{},
	// 	"data": map[string]interface{}{
	// 		"key":   "", // propertyGrid.Key,
	// 		"items": []map[string]interface{}{},
	// 		"total": 0, // 100,
	// 	},
	// 	"queryETE": "",
	// }
	if pLG.defaultField {
		prefixName := pLG.key[:2]
		pLG.columns = append(pLG.columns, Kolum{SO_Class.Fmt.Sprint(prefixName, "rgid"), KolumProperty{name: "Entry ID"}})
		pLG.columns = append(pLG.columns, Kolum{SO_Class.Fmt.Sprint(prefixName, "rgdt"), KolumProperty{name: "Entry Date"}})
		pLG.columns = append(pLG.columns, Kolum{SO_Class.Fmt.Sprint(prefixName, "rgdt_time"), KolumProperty{name: "Entry Time"}})
		pLG.columns = append(pLG.columns, Kolum{SO_Class.Fmt.Sprint(prefixName, "chid"), KolumProperty{name: "Change ID"}})
		pLG.columns = append(pLG.columns, Kolum{SO_Class.Fmt.Sprint(prefixName, "chdt"), KolumProperty{name: "Change Date"}})
		pLG.columns = append(pLG.columns, Kolum{SO_Class.Fmt.Sprint(prefixName, "chdt_time"), KolumProperty{name: "Change Time"}})
		pLG.columns = append(pLG.columns, Kolum{SO_Class.Fmt.Sprint(prefixName, "chno"), KolumProperty{name: "Change No"}})
		pLG.columns = append(pLG.columns, Kolum{SO_Class.Fmt.Sprint(prefixName, "csid"), KolumProperty{name: "Change System ID"}})
		pLG.columns = append(pLG.columns, Kolum{SO_Class.Fmt.Sprint(prefixName, "csdt"), KolumProperty{name: "Change System Date"}})
	}

	queryWhere, mapFilter := f.SetQueryFilter(pLG.c.Param("filter"))

	if pLG.c.Param("sort") == "" {
		param := gin.Param{Key: "sort", Value: `[{"property": "` + pLG.key + `","direction": "desc"}]`}
		pLG.c.Params = append(pLG.c.Params, param)
	}
	querySort, mapSort := f.SetQuerySort(pLG.c.Param("sort"))

	queryKondisi := f.SetQueryCondition(pLG.c)
	queryFinal := SO_Class.Fmt.Sprint(pLG.sqlstm, queryWhere, queryKondisi, querySort)
	queryETE := queryFinal
	queryFinal = f.SetQueryPaging(pLG.c, queryFinal)

	queryCount := SO_Class.Fmt.Sprint(`select count(*) jmlh from ( `, queryETE, ` ) a`)
	hasilRecordCount := Form.GetRs(pLG.c, queryCount)
	if !hasilRecordCount.Sukses {
		return hasilRecordCount
	}
	jmlh := hasilRecordCount.Data.([]map[string]interface{})[0]["jmlh"].(int64)

	hasil = f.GetRecordSet(pLG.c, queryFinal)

	fields, columnProperties := f.GetColumnsProperties(pLG.columns, hasil.Data.(map[string]interface{})["columns"].([]map[string]interface{}))

	ResultGrid := map[string]interface{}{
		"sort":    mapSort,
		"filter":  mapFilter,
		"fields":  fields,
		"columns": columnProperties,
		"data": map[string]interface{}{
			"key":   pLG.key, // propertyGrid.Key,
			"items": hasil.Data.(map[string]interface{})["rows"],
			"total": jmlh, // 100,
		},
		"queryETE": queryETE,
	}

	hasil.Data = ResultGrid

	return hasil

}

//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//

type ObjGrd = SO_Object.ObjGrd
type ObjPnl = SO_Object.ObjPnl
type ObjTxt = SO_Object.ObjTxt
type ObjRmk = SO_Object.ObjRmk
type ObjNum = SO_Object.ObjNum
type ObjDtp = SO_Object.ObjDtp
type ObjCmb = SO_Object.ObjCmb
type ObjRad = SO_Object.ObjRad
type ObjChg = SO_Object.ObjChg
type ObjPop = SO_Object.ObjPop
type ObjCnt = SO_Object.ObjCnt
type ObjBtn = SO_Object.ObjBtn

func (f *FormType) CrtObj(c any) map[string]interface{} {
	return SO_Object.CrtObj(c)
}

//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//
//****************************************************//

type TxFn = SO_Module.TxFn
type TransactionResult = SO_Module.TransactionResult
type Transaction = SO_Module.Transaction
type ParamIUD = SO_Module.ParamIUD
type ParamBFCS = SO_Module.ParamBFCS

func (f *FormType) ExecQueryMultiple(c *gin.Context, userName string, fn TxFn) (err error) {
	return SO_Module.Database.ExecQueryMultiple(c, userName, fn)
}

func (f *FormType) Execute(tx Transaction, tr *TransactionResult, userName string, sqlstm string) (sql sql.Result, err error) {
	return SO_Module.Database.Execute(tx, tr, userName, sqlstm)
}

func (f *FormType) CheckRecord_BFCS(bfcs ParamBFCS) (bool, error) {
	return SO_Module.Database.CheckRecord_BFCS(bfcs)
}

func (f *FormType) GetTBLNOR(tx Transaction, userName string, table string) string {
	return SO_Module.Database.GetTBLNOR(tx, userName, table)
}

func (f *FormType) GetCurrentTime() time.Time {
	return time.Now()
}

func (f *FormType) GetSintaxSQL_IUD(iud ParamIUD) string {
	return SO_Module.Database.GetSintaxSQL_IUD(iud)
}

func (f *FormType) GetDefaultField(prefixName string) (r string) {

	r = SO_Class.Fmt.Sprint(" ",
		",", prefixName, "rgid ",
		",", prefixName, "rgdt::date as ", prefixName, "rgdt",
		",TO_CHAR(", prefixName, "rgdt, 'HH24:MI:SS.MS') as ", prefixName, "rgdt_time ",
		",", prefixName, "chid ",
		",", prefixName, "chdt::date as ", prefixName, "chdt",
		",TO_CHAR(", prefixName, "chdt, 'HH24:MI:SS.MS') as ", prefixName, "chdt_time ",
		",", prefixName, "chno  ",
		",", prefixName, "csid ",
		",", prefixName, "csdt ",
		" ")

	return r
}

func (f *FormType) SetLogPrintInitFlag(flag bool) {
	f.logPrintInitFlag = flag
}

var Form = NewFormType()
