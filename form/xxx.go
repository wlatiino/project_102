package SO_Form

import (
	SO_Class "SOApp_GO/class"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type xXXX struct{}

func (xXXX) Coba1(c *gin.Context) {
	SO_Class.Log.Println(true, "Masuk XXX-Coba1()")

	var hasil SO_Class.Hasil

	v1 := SO_Class.Fmt.Sprint(`[{"property":"tmnomr","direction":"asc"}]`)

	v2 := SO_Class.Fmt.Sprint(
		`[{"property": "tsdscd","direction": "asc"}`,
		`,{"property": "tssycd","direction": "asc"}]`,
	)

	prop := []string{`[{"property": "tmnomr", "direction": "asc"}]`, `[{"property": "tmnomr", "direction": "asc"}]`}
	v3 := SO_Class.Fmt.Sprint(prop)

	v4 := SO_Class.Fmt.Sprint(
		`[{"property": "tsdscd","direction": "asc"}]`,
		`,[{"property": "tssycd","direction": "asc"}]`,
	)

	hasil.Sukses = true
	hasil.Pesan = "XXXXXXXXXXXXXXXXXXXXX"
	hasil.Data = map[string]interface{}{
		"v1": v1,
		"v2": v2,
		"v3": v3,
		"v4": v4,
	}
	c.JSON(http.StatusOK, hasil)
}

func init() {
	if Form.logPrintInitFlag {
		SO_Class.Log.Println(true, "Masuk form-xxx-init()")
	}
}

// Exported instance
var XXX xXXX
