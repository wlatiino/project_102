package SO_Object

import (
	SO_Class "SOApp_GO/class"

	"github.com/go-playground/validator/v10"
)

// Define a struct with a Print method
type ObjCnt struct {
	Id string `validate:"required"`
	// Controller  string `validate:"required"`
	// Method      string `validate:"required"`
	FrmId       string `validate:"required"`
	MenuId      string `validate:"required"`
	Items       []interface{}
	Title       interface{}
	Border      interface{}
	Collapsible interface{}
	Collapsed   interface{}
	Width       interface{}
	Height      interface{}
	Hidden      bool
	Margin      string
	Padding     string
	BodyPadding string
	TipeVH      string
	Scroll      bool
	LabelWidth  interface{}
	Html        interface{}
	OnAjax      interface{}
}

func CrtObjCnt(o ObjCnt) map[string]interface{} {
	validate := validator.New()

	err := validate.Struct(o)
	if err != nil {
		SO_Class.Fmt.Println(true, "Validation error:", err)
		return nil
	}

	c := map[string]interface{}{
		"xtype":        "xtObjCnt",
		"soObj_menuID": o.MenuId,
		"soObj_frmID":  o.FrmId,
		"soObj_id":     o.Id,
		"hidden":       o.Hidden,
		"items":        o.Items,
	}

	if o.LabelWidth != nil {
		c["labelWidth"] = o.LabelWidth
	}

	if o.TipeVH != "" {
		var tipeVH string
		if o.TipeVH == "V" {
			tipeVH = "vbox"
		} else if o.TipeVH == "H" {
			tipeVH = "hbox"
		}
		c["layout"] = tipeVH
	}

	if o.Margin != "" {
		c["margin"] = o.Margin
	}

	return c
}

func init() {
	SO_Class.Log.Println(true, "Masuk soobj-objPnl-init()")
}
