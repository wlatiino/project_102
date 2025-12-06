package SO_Object

import (
	SO_Class "SOApp_GO/class"
)

func CrtObj(c any) map[string]interface{} {

	switch v := any(c).(type) {
	case ObjGrd:
		return CrtObjGrd(v)
	case ObjPnl:
		return CrtObjPnl(v)
	case ObjTxt:
		return CrtObjTxt(v)
	case ObjRmk:
		return CrtObjRmk(v)
	case ObjNum:
		return CrtObjNum(v)
	case ObjDtp:
		return CrtObjDtp(v)
	case ObjCmb:
		return CrtObjCmb(v)
	case ObjRad:
		return CrtObjRad(v)
	case ObjChg:
		return CrtObjChg(v)
	case ObjPop:
		return CrtObjPop(v)
	case ObjCnt:
		return CrtObjCnt(v)
	case ObjBtn:
		return CrtObjBtn(v)
	default:
		return nil
	}

}

func init() {
	SO_Class.Log.Println(true, "Masuk soobj-object-init()")
}
