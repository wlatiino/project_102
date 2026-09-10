package SO_Form

import (
	SO_Class "SOApp_GO/class"

	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type ivline struct {
	table struct {
		ILILNOIY string `tipe:"1" json:"ID" key:"true" required:"true"`
		ILILNO   string `tipe:"1" json:"Line No" required:"true"`
		ILIHNOIY string `tipe:"1" json:"Inventory Header IY" required:"true"`
		ILITNOIY string `tipe:"1" json:"Item IY" required:"true"`
		ILQTYS   string `tipe:"2" json:"Qty"`
		ILHARG   string `tipe:"2" json:"Harga"`
		ILTOTL   string `tipe:"2" json:"Total" `
		ILREMK   string `tipe:"2"  json:"Remark"`
		ILUSRM   string `tipe:"2" `
	}
}

func (saya ivline) StpSave(tx Transaction, c *gin.Context,
	mode string, userName string, source string, csdt string, table ivline,
) (tr TransactionResult) {

	data := table.table
	tr.Sintax = " -- ivline.StpSave() "

	_, errBFCS := Form.CheckRecord_BFCS(ParamBFCS{
		Tx:       tx,
		Mode:     mode,
		Table:    "ivline",
		KeyField: "ihihnoiy",
		KeyValue: data.ILILNOIY,
		CSDT:     csdt,
	})
	if errBFCS != nil {
		tr.Error = errBFCS
		return tr
	}

	var err error

	if mode == "A" {
		data.ILILNOIY = Form.GetTBLNOR(tx, userName, "ivline")
	}

	sqlstm := Form.GetSintaxSQL_IUD(ParamIUD{
		TableName:  "ivline",
		UserName:   userName,
		Source:     source,
		Mode:       mode,
		StructAnda: data,
	})
	SO_Class.Log.CetakKunci(true, c, " ivline.stpSave ", sqlstm)

	_, err = Form.Execute(tx, &tr, userName, sqlstm)
	if err != nil {
		tr.Error = err
		return tr
	}

	return tr
}

func init() {
	Form.Add("IVLINE", IVLINE)
	if Form.logPrintInitFlag {
		SO_Class.Log.Println(true, "Masuk form-ivline-init()")
	}
}

// Exported instance
var IVLINE ivline
