package SO_Form

import (
	SO_Class "SOApp_GO/class"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Define a struct with a Print method
type lLogin struct{}

type user struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Database  string `json:"database"`
	VersionId string `json:"version_id"`
	Version   string `json:"version"`
	Date      string `json:"date"`
	DeviceId  string `json:"device_id"`
}

func (lLogin) Klik(c *gin.Context) {
	SO_Class.Log.Println(true, "Masuk Login-Klik()")

	var hasil SO_Class.Hasil
	hasil.Sukses = false

	var u user
	// err := SO_Class.Fungsi.ConvertRawDataToJSONWithStruct(c, &u)
	err := SO_Class.Fungsi.ConvertPostFormToJSONWithStruct(c, &u)
	if err != nil {
		hasil.Sukses = false
		hasil.Pesan = err.Error()
		c.JSON(http.StatusUnprocessableEntity, hasil)
		return
	}

	c.Set("globalDB", u.Database)

	sqlstm := SO_Class.Fmt.Sprint(`select * from tblusr where tuuser = '` + u.Username + `' `)
	SO_Class.Log.Println(false, sqlstm)
	tblusr := Form.GetRs(c, sqlstm)
	if tblusr.Data == nil {
		hasil.Sukses = false
		hasil.Pesan = "Username or Password not match!"
		c.JSON(http.StatusUnauthorized, hasil)
		return
	}
	rs := tblusr.Data.([]map[string]interface{})[0]
	matchHash := Form.VerifyPassword(u.Password, rs["tupswd"].(string))
	if !matchHash || u.Username != rs["tuuser"].(string) {
		hasil.Sukses = false
		hasil.Pesan = "Username or Password not match!"
		c.JSON(http.StatusUnauthorized, hasil)
		return
	}

	ts, err := Form.tk.CreateToken(
		u.Username,
		u.Database,
		u.Version,
		u.VersionId,
		u.DeviceId,
		"loginId")
	if err != nil {
		hasil.Sukses = false
		hasil.Pesan = err.Error()
		c.JSON(http.StatusUnprocessableEntity, hasil)
		return
	}

	SO_Class.Fmt.Println(true, "ts :", ts)

	currentTime := Form.GetCurrentTime()
	tokens := map[string]interface{}{
		"username": u.Username,
		// "db":            u.Database,
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
		"dateInfo":      currentTime.Format("2006-01-02 3:4:5"),
		// "expiredFlag":   "0",
		// "pinFlag":       "0",
		"name":   rs["tuname"].(string),
		"useriy": rs["tuuseriy"].(int64),
		"email":  "email@contoh.com",
	}

	hasil.Sukses = true
	hasil.Pesan = "Login Success"
	hasil.Data = tokens
	c.JSON(http.StatusOK, hasil)
}

func (lLogin) Refresh(c *gin.Context) {

	var hasil SO_Class.Hasil
	hasil.Sukses = false

	mapToken := map[string]string{}
	err := SO_Class.Fungsi.ConvertRawDataToJSONWithStruct(c, &mapToken)
	if err != nil {
		hasil.Pesan = SO_Class.Fmt.Sprint("unauthorized 0001!", err.Error())
		c.JSON(http.StatusUnprocessableEntity, hasil)
		return
	}

	refreshTokenString := mapToken["refresh_token"]

	token, err := Form.tk.RefreshToken(refreshTokenString)
	if err != nil {
		hasil.Pesan = "Refresh token expired"
		c.JSON(http.StatusUnauthorized, hasil)
		return
	}

	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		// refreshUUID, ok := claims["refresh_uuid"].(string) //convert the interface to string
		// if !ok {
		// 	hasil.Pesan = SO_Class.Fmt.Sprint("unauthorized 0003!")
		// 	c.JSON(http.StatusUnprocessableEntity, hasil)
		// 	return
		// }

		// _, err := Form.rd.FetchAuth(refreshUUID)
		// if err != nil {
		// 	hasil.Pesan = SO_Class.Fmt.Sprint("unauthorized 1", err.Error())
		// 	c.JSON(http.StatusUnauthorized, hasil)
		// 	return
		// }
		userId, userIdOk := claims["user_id"].(string)
		if !userIdOk {
			hasil.Pesan = SO_Class.Fmt.Sprint("unauthorized 2")
			c.JSON(http.StatusUnprocessableEntity, hasil)
			return
		}
		database, dbOk := claims["database"].(string)
		if !dbOk {
			hasil.Pesan = SO_Class.Fmt.Sprint("unauthorized 3")
			c.JSON(http.StatusUnprocessableEntity, hasil)
			return
		}
		version, versionOk := claims["version"].(string)
		if !versionOk {
			hasil.Pesan = SO_Class.Fmt.Sprint("unauthorized 4")
			c.JSON(http.StatusUnprocessableEntity, hasil)
			return
		}
		loginId, loginIdOk := claims["login_id"].(string)
		if !loginIdOk {
			hasil.Pesan = SO_Class.Fmt.Sprint("unauthorized 5")
			c.JSON(http.StatusUnprocessableEntity, hasil)
			return
		}

		// //Delete the previous Refresh Token
		// delErr := Form.rd.DeleteRefresh(refreshUUID)
		// if delErr != nil { //if any goes wrong
		// 	hasil.Pesan = SO_Class.Fmt.Sprint("unauthorized 6", delErr.Error())
		// 	c.JSON(http.StatusUnauthorized, hasil)
		// 	return
		// }

		deviceId, _ := claims["device_id"].(string)
		versionId, _ := claims["version_id"].(string)

		//Create new pairs of refresh and access tokens
		ts, createErr := Form.tk.CreateToken(userId, database, version, versionId, deviceId, loginId)
		if createErr != nil {
			hasil.Pesan = SO_Class.Fmt.Sprint("unauthorized 0004!", createErr.Error())
			c.JSON(http.StatusUnprocessableEntity, hasil)
			c.JSON(http.StatusForbidden, hasil)
			return
		}
		// //save the tokens metadata to redis
		// saveErr := Form.rd.CreateAuth(strings.ToUpper(userId)+"||"+strings.ToUpper(database)+"||"+strings.ToUpper(versi), ts)
		// if saveErr != nil {
		// 	hasil.Pesan = SO_Class.Fmt.Sprint("unauthorized 7", saveErr.Error())
		// 	c.JSON(http.StatusForbidden, hasil)
		// 	return
		// }
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		hasil.Sukses = true
		hasil.Pesan = SO_Class.Fmt.Sprint("refresh token sukses")
		hasil.Data = tokens
		c.JSON(http.StatusCreated, hasil)
	} else {
		hasil.Pesan = SO_Class.Fmt.Sprint("refresh token expired")
		c.JSON(http.StatusUnauthorized, hasil)
	}
}

func (lLogin) GetConnList(c *gin.Context) {
	SO_Class.Log.Println(true, "Masuk Login-GetConnList()")
	c.JSON(http.StatusOK, Form.GetDbConnList(c))
}

func init() {
	if Form.logPrintInitFlag {
		SO_Class.Log.Println(true, "Masuk form-login-init()")
	}
}

// Exported instance
var Login lLogin
