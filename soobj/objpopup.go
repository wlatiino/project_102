package SO_Object

import (
	SO_Class "SOApp_GO/class"

	"github.com/go-playground/validator/v10"
)

// Define a struct with a Print method
type ObjPop struct {
	Id               string `validate:"required"`
	Name             string
	FrmId            string `validate:"required"`
	MenuId           string `validate:"required"`
	Controller       string `validate:"required"`
	Method           string `validate:"required"`
	Mode             string `validate:"required"`
	PopCode          string
	PopDesc          string
	PopCodeFf        string
	PopDescFf        string
	Width            interface{}
	AllowBlank       interface{}
	Hidden           bool
	DescHidden       bool
	Checkbox         bool
	Margin           string
	Padding          string
	HideLabel        interface{}
	LabelWidth       interface{}
	OnAjax           interface{}
	OnBeforeLoadGrid interface{}
	OnAfterPopup     interface{}
	SqlCondition     string
}

func CrtObjPop(o ObjPop) map[string]interface{} {
	validate := validator.New()

	err := validate.Struct(o)
	if err != nil {
		SO_Class.Fmt.Println(true, "Validation error:", err)
		return nil
	}

	c := map[string]interface{}{
		"xtype":            "xtObjPop",
		"soObj_menuID":     o.MenuId,
		"soObj_frmID":      o.FrmId,
		"soObj_id":         o.Id,
		"soObj_popCode":    SO_Class.Strings.ReplaceAll(o.PopCode, ".", "_"),
		"soObj_popCodeFg":  SO_Class.Strings.Contains(o.PopCode, "."),
		"soObj_popDesc":    SO_Class.Strings.ReplaceAll(o.PopDesc, ".", "_"),
		"soObj_popDescFg":  SO_Class.Strings.Contains(o.PopDesc, "."),
		"soObj_name":       o.Name,
		"soObj_mode":       o.Mode,
		"soObj_controller": o.Controller,
		"soObj_method":     o.Method,
		"hidden":           o.Hidden,
		"descHidden":       o.DescHidden,
		"checkBox":         o.Checkbox,
	}

	if o.AllowBlank != nil {
		c["allowBlank"] = o.AllowBlank
	}
	if o.SqlCondition != "" {
		c["soObj_sqlCondition"] = o.SqlCondition
	}

	if o.LabelWidth != nil {
		c["labelWidth"] = o.LabelWidth
	}

	if o.OnAjax != nil {
		c["onAjax"] = o.OnAjax
	}

	if o.Margin != "" {
		c["margin"] = o.Margin
	}

	if o.Padding != "" {
		c["bodyStyle"] = "padding: " + o.Padding + "px"
	}

	if o.HideLabel != nil {
		c["hideLabel"] = o.HideLabel
	}

	if o.Width != nil {
		c["width"] = o.Width
	}

	if o.PopCodeFf != "" {
		c["popCodeFf"] = o.PopCodeFf
	}

	if o.PopDescFf != "" {
		c["popDescFf"] = o.PopDescFf
	}

	if o.OnBeforeLoadGrid != nil {
		c["onBeforeLoadGrid"] = o.OnBeforeLoadGrid
	}

	if o.OnAfterPopup != nil {
		c["onAfterPopup"] = o.OnAfterPopup
	}

	return c
}

func init() {
	SO_Class.Log.Println(true, "Masuk soobj-objPop-init()")
}
