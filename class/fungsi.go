package SO_Class

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type fFungsi struct{}

func (saya fFungsi) IsJSON(rawData []byte) bool {
	var js interface{}
	return json.Unmarshal(rawData, &js) == nil
}

func (saya fFungsi) ConvertStringToJSON(rawData []byte) (hasil Hasil) {

	hasil.Sukses = false
	hasil.Pesan = ""
	hasil.Data = nil
	hasil.Kode = ""

	var myJSOnData map[string]interface{}
	if err := json.Unmarshal(rawData, &myJSOnData); err != nil {
		Fmt.Println(true, "Convert String To JSOn ", err)
		// return nil, errors.New("Invalid JSON from reading body raw")
		hasil.Sukses = false
		hasil.Pesan = "Invalid JSON from reading body raw"
		hasil.Data = nil
		return hasil
	}

	hasil.Sukses = true
	hasil.Data = myJSOnData
	return hasil

}

func (saya fFungsi) ConvertToJSON(s any) (hasil Hasil) {

	hasil.Sukses = false
	hasil.Pesan = ""
	hasil.Data = nil
	hasil.Kode = ""

	var myJSOnData map[string]interface{}
	var err error
	switch val := s.(type) {
	case []byte:
		err = json.Unmarshal(val, &myJSOnData)
	case string:
		err = json.Unmarshal([]byte(val), &myJSOnData)
	case map[string]interface{}:
		jsonBytes, _ := json.Marshal(val)
		err = json.Unmarshal(jsonBytes, &myJSOnData)
	default:
		err = errors.New("invalid type")
	}

	if err != nil {
		Fmt.Println(true, "Error Convert To JSOn ", err)
		hasil.Sukses = false
		hasil.Pesan = Fmt.Sprint("Error Convert To JSOn ", err)
		hasil.Data = nil
		return hasil
	}

	hasil.Sukses = true
	hasil.Data = myJSOnData
	return hasil

}

func (saya fFungsi) ConvertRawDataToJSONWithStruct(c *gin.Context, structAnda any) error {

	// Read the raw body
	rawData, err := c.GetRawData()
	if err != nil {
		return errors.New(Fmt.Sprint("Failed to read body (GetRawData): %v", err))
	}

	hasil := Fungsi.ConvertToJSON(rawData)
	if !hasil.Sukses {
		return errors.New(Fmt.Sprint(hasil.Pesan))
	}
	jsonBytes, _ := json.Marshal(hasil.Data)
	err = json.Unmarshal(jsonBytes, &structAnda)

	if err != nil {
		return errors.New(Fmt.Sprint("Gagal isi struct anda : %v", err))
	}

	return nil
}

func (saya fFungsi) ConvertQueryParamToJSONWithStruct(c *gin.Context, structAnda any) error {

	rawData := c.Request.URL.Query()
	if len(rawData) == 0 {
		return nil
	}
	hasil := Fungsi.ConvertToJSON(rawData["Data"][0])
	if !hasil.Sukses {
		return errors.New(Fmt.Sprint(hasil.Pesan))
	}

	dataParam := hasil.Data.(map[string]interface{})
	for key, values := range rawData {
		if key != "Data" {
			dataParam[key] = Strings.Join(values, ",")
		}
	}

	jsonBytes, _ := json.Marshal(dataParam)
	err := json.Unmarshal(jsonBytes, &structAnda)

	if err != nil {
		return errors.New(Fmt.Sprint("Gagal isi struct anda : %v", err))
	}

	return nil
}

func (saya fFungsi) ConvertPostFormToJSONWithStruct(c *gin.Context, structAnda any) error {

	if c.PostForm("Data") == "" {
		return nil
	}

	rawData := c.PostForm("Data")

	hasil := Fungsi.ConvertToJSON(rawData)
	if !hasil.Sukses {
		return errors.New(Fmt.Sprint(hasil.Pesan))
	}
	jsonBytes, _ := json.Marshal(hasil.Data)
	err := json.Unmarshal(jsonBytes, &structAnda)

	if err != nil {
		return errors.New(Fmt.Sprint("Gagal isi struct anda : %v", err))
	}

	return nil
}

func (saya fFungsi) BindParamsToStruct(c *gin.Context, structAnda interface{}) error {

	paramsMap := make(map[string]interface{})
	for i, param := range c.Params {
		Log.Println(false, "BindParamsToStruct ", i, param)
		key := param.Key
		value := param.Value
		// Assuming you want to convert values to string, adjust as needed
		paramsMap[key] = value
	}

	structValue := reflect.ValueOf(structAnda).Elem()
	flagPrint := false
	for fieldName, fieldValue := range paramsMap {
		// Get the corresponding struct field by name
		fName := Strings.ToUpper(fieldName)
		Log.Println(flagPrint, "Field: ", fName)

		field := structValue.FieldByName(fName)

		if !field.IsValid() {
			Log.Println(flagPrint, "BindParamsToStruct - Field not found: ", fName)
		} else {
			// Set the struct field value
			field.Set(reflect.ValueOf(fieldValue))
		}
	}

	return nil
}

// Exported instance
var Fungsi fFungsi
