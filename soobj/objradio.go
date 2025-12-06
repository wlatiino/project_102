package SO_Object

import (
	SO_Class "SOApp_GO/class"
	SO_Module "SOApp_GO/module"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Define a struct with a Print method
type ObjRad struct {
	C            *gin.Context
	Id           string `validate:"required"`
	Name         string
	FrmId        string `validate:"required"`
	MenuId       string `validate:"required"`
	Mode         string `validate:"required"`
	Width        interface{}
	AllowBlank   interface{}
	Hidden       bool
	Margin       string
	Padding      string
	HideLabel    interface{}
	LabelWidth   interface{}
	OnAjax       interface{}
	Sqlstm       string
	SqlCondition string
	OrderBy      string
	Table        string
	Value        string
	// Max              interface{}
	// Min              interface{}
	// DecimalSeparator interface{}
	// DecimalPrecision int
}

func CrtObjRad(o ObjRad) map[string]interface{} {
	validate := validator.New()

	err := validate.Struct(o)
	if err != nil {
		SO_Class.Fmt.Println(true, "Validation error:", err)
		return nil
	}

	var sqlstm string
	if o.Sqlstm == "" {
		if o.Table == "" {
			SO_Class.Log.Println(false, "radio", "object Struct 'table' cannot empty")
			return nil
		}
		orderBySqlstm := "tssycd asc"
		if o.OrderBy != "" {
			orderBySqlstm = o.OrderBy
		}
		sqlstm =
			SO_Class.Fmt.Sprint(
				`
				Select 
				`,
				" '", o.Table, "' as \"name\", ",
				" rtrim(tssycd) as \"inputValue\", ",
				" rtrim(tssynm) as \"boxLabel\", ",
				" cast('0 0 0 5' as bpchar) as \"margin\"",
				`
				From tblsys 
				Left Join tbldsc on tddscd = tsdscd 
				Where tsdlfg = '0' 
				and tsdscd = '`+o.Table+`' 
				`,
				o.SqlCondition,
				" Order By ", orderBySqlstm)

	} else {
		sqlstm = o.Sqlstm
	}

	db, _ := o.C.Get("globalDB")
	SO_Class.Log.Println(false, db)
	data, jmlhRec, err := SO_Module.Database.GetRs(db.(string), sqlstm)
	if err != nil {
		SO_Class.Log.Println(true, "radio record tidak ada!!! (no rows effected) ")
		SO_Class.Log.Println(true, "sqlstm : ", sqlstm)
		SO_Class.Log.Println(true, "err : ", err)
		return nil
	} else {
		if data == nil {
			SO_Class.Log.Println(false, "radio record tidak ada!!! (no rows effected) ")
			SO_Class.Log.Println(false, "sqlstm : ", sqlstm)
		}
		SO_Class.Log.Println(false, "jmlhRec : ", jmlhRec)
	}

	c := map[string]interface{}{
		"xtype":        "xtObjRad",
		"soObj_menuID": o.MenuId,
		"soObj_frmID":  o.FrmId,
		"soObj_id":     o.Id,
		"soObj_name":   o.Name,
		"soObj_mode":   o.Mode,
		"name":         o.Table,
		"items":        data,
		"value":        o.Value,
		"hidden":       o.Hidden,
	}

	if o.AllowBlank != nil {
		c["allowBlank"] = o.AllowBlank
	}

	if o.Width != nil {
		c["width"] = o.Width
	}

	if o.LabelWidth != nil {
		c["labelWidth"] = o.LabelWidth
	}

	if o.Margin != "" {
		c["margin"] = o.Margin
	}

	if o.Padding != "" {
		c["bodyStyle"] = "padding: " + o.Padding + "px"
	}

	if o.OnAjax != nil {
		c["onAjax"] = o.OnAjax
	}

	if o.HideLabel != nil {
		c["hideLabel"] = o.HideLabel
	}

	return c
}

func init() {
	SO_Class.Log.Println(true, "Masuk soobj-objRad-init()")
}
