package SO_Object

import (
	SO_Class "SOApp_GO/class"

	"github.com/go-playground/validator/v10"
)

// Define a struct with a Print method
type ObjPnl struct {
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

func CrtObjPnl(o ObjPnl) map[string]interface{} {
	validate := validator.New()

	err := validate.Struct(o)
	if err != nil {
		SO_Class.Fmt.Println(true, "Validation error:", err)
		return nil
	}

	c := map[string]interface{}{
		"xtype":        "xtObjPnl",
		"soObj_menuID": o.MenuId,
		"soObj_frmID":  o.FrmId,
		"soObj_id":     o.Id,
		"hidden":       o.Hidden,
		"items":        o.Items,
	}

	if o.Border != nil {
		c["border"] = o.Border
	}

	if o.Collapsible != nil {
		c["collapsible"] = o.Collapsible
	}

	if o.Collapsed != nil {
		c["collapsed"] = o.Collapsed
	}

	if o.Width != nil {
		c["width"] = o.Width
	}

	if o.Height != nil {
		c["height"] = o.Height
	}

	if o.Margin != "" {
		c["margin"] = o.Margin
	}

	if o.BodyPadding != "" {
		c["bodyPadding"] = o.BodyPadding
	}

	if o.Padding != "" {
		c["bodyStyle"] = "padding: " + o.Padding + "px"
	}

	if o.TipeVH != "" {
		var tipeVH string
		if o.TipeVH == "V" {
			tipeVH = "vbox"
		} else if o.TipeVH == "H" {
			tipeVH = "hbox"
		} else if o.TipeVH == "fit" {
			tipeVH = "fit"
		}
		c["layout"] = tipeVH
	}

	if o.Scroll {
		c["overflowY"] = "auto"
		c["overflowX"] = "auto"
	}

	if o.Title != nil {
		c["title"] = o.Title
	}

	if o.LabelWidth != nil {
		c["labelWidth"] = o.LabelWidth
	}

	if o.Html != nil {
		c["html"] = o.Html
	}

	if o.OnAjax != nil {
		c["onAjax"] = o.OnAjax
	}

	return c
}

func init() {
	SO_Class.Log.Println(true, "Masuk soobj-objPnl-init()")
}
