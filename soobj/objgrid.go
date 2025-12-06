package SO_Object

import (
	SO_Class "SOApp_GO/class"

	"github.com/go-playground/validator/v10"
)

// Define a struct with a Print method
type ObjGrd struct {
	Id                    string
	Controller            string
	Method                string
	FrmId                 string
	MenuId                string
	SqlCondition          string
	Title                 interface{}
	Border                interface{}
	Collapsible           interface{}
	Collapsed             interface{}
	Scroll                bool
	Width                 interface{}
	Height                interface{}
	Frame                 interface{}
	Action                string
	DetailItems           []interface{}
	Checkbox              bool
	ShowGridAdvSearch     interface{}
	OnActionGrid          interface{}
	OnActionGridSubmit    interface{}
	OnActionGridShow      interface{}
	OnBeforeLoadGrid      interface{}
	OnAfterLoadGrid       interface{}
	OnAfterRender         interface{}
	OnDoubleKlik          interface{}
	OnCellDoubleKlik      interface{}
	OnCheckChange         interface{}
	WindowWidth           interface{}
	WindowHeight          interface{}
	WindowResizable       interface{}
	GetAllRecordsCheckbox bool
	NoPaging              bool
}

func CrtObjGrd(o ObjGrd) map[string]interface{} {
	validate := validator.New()

	err := validate.Struct(o)
	if err != nil {
		SO_Class.Fmt.Println(true, "Validation error:", err)
		return nil
	}

	xTipe := "xtObjGrp"
	tipeGrid := "WithPaging"
	if o.NoPaging {
		tipeGrid = "NoPaging"
		xTipe = "xtObjGrd"
	}

	c := map[string]interface{}{
		"xtype":                 xTipe,
		"soObj_menuID":          o.MenuId,
		"soObj_frmID":           o.FrmId,
		"soObj_id":              o.Id,
		"soObj_controller":      o.Controller,
		"soObj_method":          o.Method,
		"soObj_tipeGrid":        tipeGrid,
		"soObj_filters":         "",
		"soObj_sorters":         "",
		"checkbox":              o.Checkbox,
		"getAllRecordsCheckbox": o.GetAllRecordsCheckbox,
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

	if o.Title != nil {
		c["title"] = o.Title
	}

	if o.Frame != nil {
		c["frame"] = o.Border
	}

	if o.Action != "" {
		c["soObj_action"] = o.Action
	}

	if o.DetailItems != nil {
		c["detailItems"] = o.DetailItems
	}

	if o.OnActionGrid != nil {
		c["onActionGrid"] = o.OnActionGrid
	}

	if o.OnActionGridSubmit != nil {
		c["onActionGridSubmit"] = o.OnActionGridSubmit
	}

	if o.OnActionGridShow != nil {
		c["onActionGridShow"] = o.OnActionGridShow
	}

	if o.OnBeforeLoadGrid != nil {
		c["onBeforeLoadGrid"] = o.OnBeforeLoadGrid
	}

	if o.OnAfterLoadGrid != nil {
		c["onAfterLoadGrid"] = o.OnAfterLoadGrid
	}

	if o.ShowGridAdvSearch != nil {
		c["showGridAdvSearch"] = o.ShowGridAdvSearch
	}

	if o.OnAfterRender != nil {
		c["onAfterRender"] = o.OnAfterRender
	}

	if o.WindowWidth != nil {
		c["windowWidth"] = o.WindowWidth
	}

	if o.WindowHeight != nil {
		c["windowHeight"] = o.WindowHeight
	}

	if o.WindowResizable != nil {
		c["windowResizable"] = o.WindowResizable
	}

	if o.SqlCondition != "" {
		c["soObj_sqlCondition"] = o.SqlCondition
	}

	if o.OnDoubleKlik != nil {
		c["onDoubleKlik"] = o.OnDoubleKlik
	}

	if o.OnCellDoubleKlik != nil {
		c["onCellDoubleKlik"] = o.OnCellDoubleKlik
	}

	if o.OnCheckChange != nil {
		c["onCheckChange"] = o.OnCheckChange
	}

	return c
}

func init() {
	SO_Class.Log.Println(true, "Masuk soobj-objGrd-init()")
}
