package SO_Object

import (
	SO_Class "SOApp_GO/class"

	"github.com/go-playground/validator/v10"
)

// Define a struct with a Print method
type ObjNum struct {
	Id               string `validate:"required"`
	Name             string
	FrmId            string `validate:"required"`
	MenuId           string `validate:"required"`
	Mode             string `validate:"required"`
	Width            interface{}
	AllowBlank       interface{}
	Hidden           bool
	Margin           string
	Padding          string
	HideLabel        interface{}
	LabelWidth       interface{}
	OnAjax           interface{}
	Max              interface{}
	Min              interface{}
	DecimalSeparator interface{}
	DecimalPrecision int
}

func CrtObjNum(o ObjNum) map[string]interface{} {
	validate := validator.New()

	err := validate.Struct(o)
	if err != nil {
		SO_Class.Fmt.Println(true, "Validation error:", err)
		return nil
	}

	c := map[string]interface{}{
		"xtype":            "xtObjNum",
		"soObj_menuID":     o.MenuId,
		"soObj_frmID":      o.FrmId,
		"soObj_id":         o.Id,
		"soObj_name":       o.Name,
		"soObj_mode":       o.Mode,
		"hidden":           o.Hidden,
		"decimalPrecision": o.DecimalPrecision,
	}

	if o.AllowBlank != nil {
		c["allowBlank"] = o.AllowBlank
	}

	if o.DecimalSeparator != nil {
		c["decimalSeparator"] = o.DecimalSeparator
	}

	if o.Width != nil {
		c["width"] = o.Width
	}

	if o.LabelWidth != nil {
		c["labelWidth"] = o.LabelWidth
	}

	if o.Max != nil {
		c["maxValue"] = o.Max
	}

	if o.Min != nil {
		c["minValue"] = o.Min
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
	SO_Class.Log.Println(true, "Masuk soobj-objNum-init()")
}
