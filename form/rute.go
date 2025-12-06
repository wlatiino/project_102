package SO_Form

import (
	SO_Class "SOApp_GO/class"
	SO_Module "SOApp_GO/module"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type rRute struct{}

func init() {
	SO_Class.Log.Println(true, "Masuk form-rute-init()")
	if err := godotenv.Load(); err != nil {
		SO_Class.Log.Println(true, "No .env file found!")
	}
}

func (r rRute) Mulai() {

	SO_Class.Log.Println(true, "")
	SO_Class.Log.Println(true, "***** ***** ***** Start Server ***** ***** *****")

	var router = gin.Default()

	if SO_Class.Strings.ToUpper(os.Getenv("GIN_MODE")) == "RELEASE" {
		gin.SetMode(gin.ReleaseMode)
	}
	if SO_Class.Strings.ToUpper(os.Getenv("GIN_MODE")) == "DEV" {
		gin.SetMode(gin.DebugMode)
		router.Use(cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			// AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
			// AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", " X-Requested-With", "Content-Type", "Accept", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "*"
			},
			MaxAge: 12 * time.Hour,
		}))
	}

	var hasil = SO_Class.Hasil{}
	// var rd = SO_Module.NewAuth(redisClient)
	var tk = SO_Module.NewToken()

	Form.BuatIdentitasForm(nil, tk)

	appAddr := ":" + os.Getenv("PORT")
	srv := &http.Server{
		Addr:    appAddr,
		Handler: router,
	}

	router.NoRoute(func(c *gin.Context) {
		SO_Class.Fmt.Println(true, "url : ", c.Request.URL, " (tidak ada rute) ")
		hasil.Sukses = false
		hasil.Kode = "PAGE_NOT_FOUND"
		hasil.Pesan = "page_not_found"
		c.JSON(404, hasil)
	})

	router.POST("/login", Login.Klik)
	router.POST("/refreshToken", Login.Refresh)
	router.GET("/getConnList", Login.GetConnList)
	router.GET("/xxx", XXX.Coba1)

	router.GET("/dataso", SO_Module.Middleware.TokenAuthMiddleware(), Rute.Panggil)
	router.POST("/dataso", SO_Module.Middleware.TokenAuthMiddleware(), Rute.Panggil)
	router.PUT("/dataso", SO_Module.Middleware.TokenAuthMiddleware(), Rute.Panggil)
	router.DELETE("/dataso", SO_Module.Middleware.TokenAuthMiddleware(), Rute.Panggil)
	// router.POST("/dataso", TBLDSC.LoadGrid)
	// router.PUT("/dataso", TBLDSC.LoadGrid)
	// router.DELETE("/dataso", TBLDSC.LoadGrid)

	go func() {
		SO_Class.Log.Println(true, "Listen And Serve", srv.ListenAndServe())
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			SO_Class.Log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	SO_Class.Log.Println(true, "***** ***** *****  End Server  ***** ***** *****")
	SO_Class.Log.Println(true, "")
}

func (r rRute) Panggil(c *gin.Context) {
	flagPrint := false

	SO_Class.Fmt.Println(flagPrint, "Full URL:", c.Request.URL.String())    // path + query string
	SO_Class.Fmt.Println(flagPrint, "Path:", c.Request.URL.Path)            // only path
	SO_Class.Fmt.Println(flagPrint, "Query params:", c.Request.URL.Query()) // map[string][]string
	SO_Class.Fmt.Println(flagPrint, "Method:", c.Request.Method)            // string
	// Example to get full URL with scheme and host
	SO_Class.Fmt.Println(flagPrint, "Scheme + Host + URL:", c.Request.Host, c.Request.URL.String())

	var hasil SO_Class.Hasil
	var data map[string]interface{}

	var err error
	method := c.Request.Method
	if method == "GET" {
		err = SO_Class.Fungsi.ConvertQueryParamToJSONWithStruct(c, &data)
	} else {
		err = SO_Class.Fungsi.ConvertPostFormToJSONWithStruct(c, &data)
	}
	// err := SO_Class.Fungsi.ConvertRawDataToJSONWithStruct(c, &data)
	if err != nil {
		hasil.Sukses = false
		hasil.Pesan = err.Error()
		c.JSON(http.StatusUnprocessableEntity, hasil)
		return
	}

	// Begin Cari Rute
	controller, errController := data["Controller"].(string)
	if !errController {
		hasil.Sukses = false
		hasil.Pesan = "param controller's not found"
		c.JSON(400, hasil)
		return
	}
	method, errMethod := data["Method"].(string)
	if !errMethod {
		hasil.Sukses = false
		hasil.Pesan = "param method's not found"
		c.JSON(400, hasil)
		return
	}

	var dataParams []gin.Param
	for key, value := range data {
		param := gin.Param{Key: key, Value: SO_Class.Fmt.Sprint(value)}
		SO_Class.Log.Println(false, "func Rutes Panggil --> key : ", key, " , value : ", value, " , param : ", param)
		dataParams = append(dataParams, param)
	}
	c.Params = dataParams

	HasilAkhir := Rute.Cari(controller, method, c)

	if !HasilAkhir.Sukses {
		// c.JSON(400, HasilAkhir)
		c.JSON(http.StatusOK, HasilAkhir)
		return
	}
	// End Cari Rute

	c.JSON(http.StatusOK, HasilAkhir)
}

func (r rRute) Cari(controllerName string, methodName string, params ...interface{}) (hasil SO_Class.Hasil) {

	hasil.Sukses = false
	hasil.Pesan = ""
	hasil.Data = nil
	hasil.Kode = ""

	if controllerName == "" || methodName == "" {
		hasil.Sukses = false
		hasil.Pesan = SO_Class.Fmt.Sprint("Controller and Method cannot be empty!")
		return hasil
	}

	f := Form
	SO_Class.Fmt.Println(false, "List Files :", f.forms)

	obj := f.forms[controllerName]
	if obj == nil {
		hasil.Sukses = false
		hasil.Pesan = SO_Class.Fmt.Sprint("Controller : ", controllerName, " and Method : ", methodName, " not found!")
		return hasil
	}
	v := reflect.ValueOf(obj)

	flagListMethod := false
	if flagListMethod {
		t := reflect.TypeOf(obj)
		for i := 0; i < t.NumMethod(); i++ {
			SO_Class.Fmt.Println(true, "Method:", t.Method(i).Name)
		}
	}

	// Find method by name
	method := v.MethodByName(methodName)
	if !method.IsValid() {
		SO_Class.Log.Println(true, "form-rute-cari -> Controller : ", controllerName, " and Method : ", methodName, " not found")
		hasil.Sukses = false
		hasil.Pesan = SO_Class.Fmt.Sprint("Controller : ", controllerName, " and Method : ", methodName, "not found")
		return hasil
	}

	// Convert params to []reflect.Value
	args := make([]reflect.Value, len(params))
	for i, p := range params {
		args[i] = reflect.ValueOf(p)
	}

	// Call method
	var res []reflect.Value
	var results any

	res = method.Call(args)
	if len(res) == 0 {
		results = res
	} else {
		results = res[0].Interface()
	}

	// If method returns something, you can process results here
	SO_Class.Fmt.Println(false, "form-rute-cari -> Results:", results)
	hasil = results.(SO_Class.Hasil)

	return hasil
}

/*
"go mod init App01"
"go mod tidy"
"go mod vendor"
*/

/*
$ go get -u github.com/golang-jwt/jwt
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/gin-contrib/cors
$ go get -u github.com/go-sql-driver/mysql v1.5.0
*/

var Rute rRute
