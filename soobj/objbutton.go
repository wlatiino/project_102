package SO_Object

import (
	SO_Class "SOApp_GO/class"

	"github.com/go-playground/validator/v10"
)

// Define a struct with a Print method
type ObjBtn struct {
	Id         string `validate:"required"`
	Name       string
	FrmId      string `validate:"required"`
	MenuId     string `validate:"required"`
	Mode       string `validate:"required"`
	Text       string
	Scale      string
	IconCls    string
	IconAlign  string
	Width      interface{}
	Hidden     bool
	Margin     string
	LabelWidth interface{}
	OnKlik     interface{}
}

func CrtObjBtn(o ObjBtn) map[string]interface{} {
	validate := validator.New()

	err := validate.Struct(o)
	if err != nil {
		SO_Class.Fmt.Println(true, "Validation error:", err)
		return nil
	}

	c := map[string]interface{}{
		"xtype":        "xtObjBtn",
		"soObj_menuID": o.MenuId,
		"soObj_frmID":  o.FrmId,
		"soObj_id":     o.Id,
		"soObj_name":   o.Name,
		"soObj_mode":   o.Mode,
		"hidden":       o.Hidden,
	}

	if o.LabelWidth != nil {
		c["labelWidth"] = o.LabelWidth
	}

	if o.Margin != "" {
		c["margin"] = o.Margin
	}

	if o.Width != nil {
		c["width"] = o.Width
	}

	if o.Text != "" {
		c["text"] = o.Text
	}

	if o.Scale != "" {
		c["scale"] = o.Scale
	}

	if o.IconCls != "" {
		c["iconCls"] = o.IconCls
	}

	if o.IconAlign != "" {
		c["iconAlign"] = o.IconAlign
	}

	if o.OnKlik != nil {
		c["onKlik"] = o.OnKlik
	}

	return c
}

func init() {
	SO_Class.Log.Println(true, "Masuk soobj-objTxt-init()")
}
