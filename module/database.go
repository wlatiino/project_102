package SO_Module

import (
	SO_Class "SOApp_GO/class"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// Define a struct with a Print method
type db struct{}

func (saya db) RedisDB(host, port, password string, redisDB int, redisUNIX string, redisSOCK string) *redis.Client {
	if redisUNIX == "1" {
		redisClient := redis.NewClient(&redis.Options{
			Network:  "unix",
			Addr:     redisSOCK,
			Password: password,
			DB:       redisDB,
		})
		return redisClient
	} else {
		redisClient := redis.NewClient(&redis.Options{
			Addr:     host + ":" + port,
			Password: password,
			DB:       redisDB,
		})
		return redisClient
	}
}

func (saya db) GetDbConnList() []map[string]interface{} {
	// Read JSON file path from environment variable
	jsonFilePath := os.Getenv("DBCONFIG_FILE_PATH")

	// Read JSON data from file
	jsonData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		SO_Class.Fmt.Println(true, "Error reading JSON file:", err)
	}

	// Unmarshal JSON data into struct
	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		SO_Class.Fmt.Println(true, "Error decoding JSON:", err)
	}

	var dataHasil []map[string]interface{}
	for key, values := range data {
		nilai := map[string]interface{}{
			"value": key,
			"label": values.(map[string]interface{})["label"],
		}
		dataHasil = append(dataHasil, nilai)
	}

	return dataHasil

}

func getDbConn(dbvalue string) map[string]interface{} {

	// Read JSON file path from environment variable
	jsonFilePath := os.Getenv("DBCONFIG_FILE_PATH")

	// Read JSON data from file
	jsonData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		SO_Class.Fmt.Println(true, "Error reading JSON file:", err)
	}

	// Unmarshal JSON data into struct
	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		SO_Class.Fmt.Println(true, "Error decoding JSON:", err)
	}

	// Print the decoded JSON data
	// log.Println("dbconfigdbconfigdbconfig222", data[dbvalue])

	if data[dbvalue] == nil {
		SO_Class.Fmt.Println(true, "DB CONFIG TIDAK SESUAI!")
		return nil
	}

	return data[dbvalue].(map[string]interface{})

}

func openConnection(connString string) (Conn *sql.DB, err error) {

	// sqlTipe := "pqsql" //mysql
	// switch sqlTipe {
	// case "mysql":
	// 	dbServer := "127.0.0.1"
	// 	dbName := "golang02" //golang01
	// 	dbUser := "golang"
	// 	dbPass := "go123"

	// 	DataSourceName := SO_Class.Fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", dbUser, dbPass, dbServer, dbName)
	// 	Conn, err = sql.Open("mysql", DataSourceName)

	// case "pqsql":

	// 	dbServer := os.Getenv("POSTGRESQL_HOST")
	// 	dbName := os.Getenv("POSTGRESQL_DB")
	// 	dbUser := os.Getenv("POSTGRESQL_USER")
	// 	dbPass := os.Getenv("POSTGRESQL_PASSWORD")
	// 	dbSchema := os.Getenv("POSTGRESQL_SCHEMA")

	// 	DataSourceName := SO_Class.Fmt.Sprintf("host=%s user=%s "+
	// 		"password=%s dbname=%s sslmode=disable search_path=%s",
	// 		dbServer, dbUser, dbPass, dbName, dbSchema)

	// 	Conn, err = sql.Open("postgres", DataSourceName)
	// }

	dbinfo := getDbConn(connString)
	if dbinfo == nil {
		return nil, SO_Class.Fmt.Errorf("Database Config tidak sesuai!!!")
	}

	dbServer := dbinfo["host"]
	dbName := dbinfo["db"]
	dbUser := dbinfo["user"]
	dbPass := dbinfo["password"]
	dbSchema := dbinfo["schema"]
	dbPort := dbinfo["port"]

	DataSourceName := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable search_path=%s TimeZone=Asia/Jakarta",
		dbServer, dbPort, dbUser, dbPass, dbName, dbSchema)

	// log.Println("DataSourceNameDataSourceNameDataSourceName", DataSourceName)

	Conn, err = sql.Open("postgres", DataSourceName)

	if err != nil {
		SO_Class.Log.Println(false, "host=", dbServer, " port=", dbPort, " user=", dbUser, " "+
			"password=", dbPass, " dbname=", dbName, " sslmode=disable search_path=", dbSchema, " TimeZone=Asia/Jakarta",
		)
		SO_Class.Log.Println(true, "func OpenConnection ", Conn, " error :", err)
	}
	return Conn, err

}

func (saya db) GetRs(connString string, sqlstm string) (hasil []map[string]interface{}, jumlahRecord int, errRs error) {

	jumlahRecord = 0
	Conn, err := openConnection(connString)
	if err != nil {
		return nil, jumlahRecord, err
	} else {
		defer Conn.Close()
	}

	var rs *sql.Rows

	rs, errRs = Conn.Query(sqlstm)

	if errRs != nil {
		hasil = nil
		SO_Class.Log.Println(true, "func GetRs ", err)
	} else {
		defer rs.Close()

		var results []map[string]interface{}
		columns, _ := rs.Columns()

		// looping tiap record
		for rs.Next() {

			values := make([]interface{}, len(columns))
			fields := make([]interface{}, len(columns))
			for i := range values {
				fields[i] = &values[i]
			}

			err = rs.Scan(fields...)
			if err != nil {
				hasil = nil
				SO_Class.Log.Println(true, "func OpenRs ", err)
				break
			}

			rows := make(map[string]interface{})
			columnTypes, _ := rs.ColumnTypes()

			// looping kolom untuk tiap record
			for i, fieldName := range columns {
				field := values[i]
				rows[fieldName] = getFieldNilai(string(columnTypes[i].DatabaseTypeName()), field)
			}
			results = append(results, rows)
		}

		jumlahRecord = len(results)
		hasil = results
		errRs = nil
	}

	// SO_Class.Log.Println(true, "func GetRs --> hasil : ", hasil)
	return hasil, jumlahRecord, errRs
}

func (saya db) GetRecordSet(connString string, sqlstm string) (hasil map[string]interface{}, errRs error) {

	hasil = map[string]interface{}{
		"columns": []map[string]interface{}{},
		"row":     map[string]interface{}{},
	}

	Conn, err := openConnection(connString)
	if err != nil {
		return nil, err
	} else {
		defer Conn.Close()
	}

	var rs *sql.Rows

	SO_Class.Log.Println(false, "func GetRecordSet ", sqlstm)
	rs, errRs = Conn.Query(sqlstm)

	if errRs != nil {
		hasil = nil
		SO_Class.Log.Println(true, "func GetRecordSet ", err)
	} else {
		defer rs.Close()

		var dataRows []map[string]interface{}
		var dataColumns []map[string]interface{}
		columns, _ := rs.Columns()
		columnTypes, _ := rs.ColumnTypes()

		// looping tiap kolom
		for i, colName := range columns {

			colType := columnTypes[i]
			fldName := colName
			nullable, _ := colType.Nullable()
			length, _ := colType.Length()
			precision, scale, decimalFlag := colType.DecimalSize()

			dataColumns = append(dataColumns, map[string]interface{}{
				"code":        fldName,
				"name":        colName,
				"type":        SO_Class.Fmt.Sprint(colType.DatabaseTypeName()),
				"nullable":    nullable,
				"length":      length,
				"decimalFlag": decimalFlag,
				"precision":   precision,
				"scale":       scale,
				// "ScanType":    SO_Class.Fmt.Sprint(colType.ScanType()),
			})
		}

		// looping tiap record
		for rs.Next() {

			values := make([]interface{}, len(columns))
			fields := make([]interface{}, len(columns))
			for i := range values {
				fields[i] = &values[i]
			}

			err = rs.Scan(fields...)
			if err != nil {
				hasil = nil
				SO_Class.Log.Println(true, "func GetRecordSet ", err)
				break
			}

			rows := make(map[string]interface{})
			columnTypes, _ := rs.ColumnTypes()

			// looping "kolom" untuk tiap record
			for i, fieldName := range columns {
				field := values[i]
				rows[fieldName] = getFieldNilai(string(columnTypes[i].DatabaseTypeName()), field)
			}
			dataRows = append(dataRows, rows)
		}

		hasil = map[string]interface{}{
			"columns": dataColumns,
			"rows":    dataRows,
		}

		errRs = nil
	}

	// SO_Class.Log.Println(true, "func GetRecordSet --> hasil : ", hasil)
	return hasil, errRs
}

func getFieldNilai(fieldType string, field any) (fieldValue interface{}) {

	SO_Class.Log.Println(false, "field Type : ", fieldType)
	// fieldType := fmt.Sprint(reflect.TypeOf(field))
	switch SO_Class.Strings.ToLower(fieldType) {
	case "string", "text", "varchar": // varchar, text
		fieldValue = ""
		if field != nil {
			nilai := field.(string)
			fieldValue = SO_Class.Strings.TrimRight(nilai, " ")
		}
	case "bpchar", "bit": // bpchar
		fieldValue = ""
		if field != nil {
			nilai := string(field.([]byte))
			fieldValue = SO_Class.Strings.TrimRight(nilai, " ")
		}
	case "time.time", "timestamp", "date": // timestamp
		fieldValue = field
		if field != nil {
			nilai := field.(time.Time)
			if SO_Class.Strings.Contains(nilai.String(), "00:00:00") {
				fieldValue = nilai.Format("02-Jan-2006")
			} else {
				fieldValue = nilai.Format("02-Jan-2006 15:04:05")
			}
		}
	case "numeric", "int4", "int8":
		fieldValue = 0
		if field != nil {
			fieldValue = field
		}
	default:
		fieldValue = field
	}

	switch any(fieldValue).(type) {
	case []uint8:
		fieldValue = string(field.([]byte))
	}

	return fieldValue
}

func (saya db) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(bytes), err
}

func (saya db) VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type ParamIUD struct {
	Mode       string
	UserName   string
	Source     string
	TableName  string
	StructAnda interface{}
}

func (saya db) GetSintaxSQL_IUD(iud ParamIUD) string {

	tbl := reflect.TypeOf(iud.StructAnda)
	val := reflect.ValueOf(iud.StructAnda)

	sql := ""
	var fields []string
	var values []string
	var fv []string // field and value
	var keys []string
	var prefix string

	for i := 0; i < tbl.NumField(); i++ {
		field := tbl.Field(i)
		value := val.Field(i).Interface()
		fieldName := SO_Class.Strings.ToLower(field.Name)
		if prefix == "" {
			prefix = fieldName[:2]
		}

		if field.Tag.Get("key") == "true" {
			keys = append(keys, SO_Class.Fmt.Sprintf("%s = '%s'", fieldName, value))
		}

		if field.Tag.Get("tipe") <= "2" {
			fields = append(fields, fieldName)
			values = append(values, SO_Class.Fmt.Sprintf("'%s'", value))
		}

		if field.Tag.Get("tipe") == "2" {
			fv = append(fv, SO_Class.Fmt.Sprintf("%s = '%s'", fieldName, value))
		}
	}

	currentTime := time.Now()
	switch SO_Class.Strings.ToUpper(iud.Mode) {
	case "A":
		sql = SO_Class.Fmt.Sprint(
			" insert into ", iud.TableName, " (",
			SO_Class.Strings.Join(fields, " , "),
			" , ", prefix, "rgid",
			" , ", prefix, "rgdt",
			" , ", prefix, "chid",
			" , ", prefix, "chdt",
			" , ", prefix, "chno",
			" , ", prefix, "srce",
			" , ", prefix, "csid",
			" , ", prefix, "csdt",
			") values (",
			SO_Class.Strings.Join(values, " , "),
			" , '", iud.UserName, "'",
			" , '", currentTime.Format("2006-01-02 3:4:5"), "'",
			" , '", iud.UserName, "'",
			" , '", currentTime.Format("2006-01-02 3:4:5"), "'",
			" , '0'",
			" , '", iud.Source, "'",
			" , '", iud.UserName, "'",
			" , '", currentTime.Format("2006-01-02 3:4:5"), "'",
			");")
	case "E":
		sql = SO_Class.Fmt.Sprint(
			" update ", iud.TableName,
			" set ", SO_Class.Strings.Join(fv, " , "),
			" , ", prefix, "chid = '", iud.UserName, "'",
			" , ", prefix, "chdt = '", currentTime.Format("2006-01-02 3:4:5"), "'",
			" , ", prefix, "chno = coalesce(", prefix, "chno,0)+1",
			" , ", prefix, "srce = '", iud.Source, "'",
			" , ", prefix, "csid = '", iud.UserName, "'",
			" , ", prefix, "csdt = '", currentTime.Format("2006-01-02 3:4:5"), "'",
			" where ", SO_Class.Strings.Join(keys, " and "), ";")
	case "D":
		sql = SO_Class.Fmt.Sprint(
			" delete from ", iud.TableName,
			" where ", SO_Class.Strings.Join(keys, " and "), ";")
	}

	SO_Class.Log.Println(true, "GetSintasxSQL_IUD : ", sql)

	return sql
}

type Transaction interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type TransactionResult struct {
	Error  error
	Sintax string
}

type ParamBFCS struct {
	Tx       Transaction
	Mode     string
	Table    string
	KeyField string
	KeyValue string
	CSDT     string
}

func (saya db) CheckRecord_BFCS(bfcs ParamBFCS) (bool, error) {

	if bfcs.Mode == "E" || bfcs.Mode == "D" {
		var csdt string
		var csid string
		prefix := bfcs.KeyField[:2]
		errRs := bfcs.Tx.QueryRow(
			SO_Class.Fmt.Sprint(
				" select ",
				"", prefix, "csdt, ",
				"", prefix, "csid ",
				"from ", bfcs.Table, " ",
				" where ", bfcs.KeyField, " = '", bfcs.KeyValue, "'",
			),
		).Scan(&csdt, &csid)

		if errRs != nil {
			return false, SO_Class.Fmt.Errorf(
				"Record not found! '%s'<br> Please refresh your data!<br> %s",
				bfcs.KeyValue, errRs)
		}

		t1, _ := time.Parse(time.RFC3339, csdt)
		t2, _ := time.Parse("02-Jan-2006 15:04:05", bfcs.CSDT)
		// t2, _ := time.Parse("2006-01-02 15:04:05", bfcs.CSDT)

		if !t1.Equal(t2) {
			return false, SO_Class.Fmt.Errorf(
				"This record just change by %s,<br> During time %s! (%s)",
				csid, csdt, bfcs.CSDT)
		}
	}
	return true, nil
}

func (saya db) GetTBLNOR(tx Transaction, userName string, table string) string {

	var iy string
	sqlstm :=
		SO_Class.Fmt.Sprint(
			" select stpTBLNOR('", userName, "','", table, "') ",
		)
	errTBLNOR := tx.QueryRow(sqlstm).Scan(&iy)
	if errTBLNOR != nil {
		SO_Class.Log.Println(true, "error StpTblNor ", errTBLNOR)
		iy = "0"
	}

	execTBLSLF(tx, userName, SO_Class.Fmt.Sprint(sqlstm, " -- ", iy))

	return iy
}

func execTBLSLF(tx Transaction, userName string, sqlstm string) (sql sql.Result, err error) {

	// Begin masuk ke LogFile
	currentTime := time.Now()
	sql, err = tx.Exec(fmt.Sprint("",
		" insert into tblslf (",
		" tquser, tqstmt, tqremk,",
		" tqrgid, tqrgdt, tqchid, tqchdt, tqchno, tqcsid, tqcsdt)",
		" values",
		" (",
		" '", userName, "',", "'", SO_Class.Strings.ReplaceAll(sqlstm, "'", "`"), "',", "'',",
		" '", userName, "',", "'", currentTime.Format("2006-01-02 3:4:5"), "',",
		" '", userName, "',", "'", currentTime.Format("2006-01-02 3:4:5"), "',", "'", 0, "',",
		" '", userName, "',", "'", currentTime.Format("2006-01-02 3:4:5"), "'",
		" )"))
	// log.Println("func ExecTBLSLF --> ", "sql -->", sql)
	// log.Println("func ExecTBLSLF --> ", "sqlstm -->", strings.ReplaceAll(sqlstm, "'", "``"))
	if err != nil {
		return sql, err
	}
	// End masuk ke LogFile
	return sql, err
}

func execTBLELF(c *gin.Context, sqlstm string, remark string) (sql sql.Result, err error) {

	globalDB, _ := c.Get("globalDB")
	tx, err := openConnection(globalDB.(string))
	if err != nil {
		return nil, err
	} else {
		defer tx.Close()
	}

	SO_Class.Log.Println(false, "func ExecTBLELF --> ", "sqlstm -->", sqlstm)
	SO_Class.Log.Println(false, "func ExecTBLELF --> ", "remark -->", remark)
	// Begin masuk ke LogFile
	currentTime := time.Now()
	stmt := fmt.Sprint("",
		"insert into tblelf (",
		"teuser, testmt, teremk,",
		"teerst, teerms,",
		"tergid, tergdt, techid, techdt, techno, tecsid, tecsdt)",
		"values",
		" (",
		" '", c.Param("username"), "',", "'", SO_Class.Strings.ReplaceAll(sqlstm, "'", "\""), "',", "'", "", "',",
		" '", "99999", "',", "'", SO_Class.Strings.TrimLeft(SO_Class.Strings.ReplaceAll(remark, "'", "\""), "pq: "), "',",
		" '", c.Param("username"), "',", "'", currentTime.Format("2006-01-02 3:4:5"), "',",
		" '", c.Param("username"), "',", "'", currentTime.Format("2006-01-02 3:4:5"), "',", "'", 0, "',",
		" '", c.Param("username"), "',", "'", currentTime.Format("2006-01-02 3:4:5"), "'",
		" )")
	sql, err = tx.Exec(stmt)
	if err != nil {
		SO_Class.Log.Println(true, "func ExecTBLELF error --> ", err, " stmt -->", stmt)
		return sql, err
	}
	// End masuk ke LogFile
	return sql, err
}

// // ************************************************************************************************** //

func (saya db) Execute(tx Transaction, tr *TransactionResult, userName string, sqlstm string) (sql sql.Result, err error) {

	tr.Sintax = tr.Sintax + sqlstm + `
` // ini fungsinya untuk enter (`)

	sql, err = tx.Exec(sqlstm) // Execute Sintax SQL Statement seperti Insert, Update dan Delete
	if err != nil {
		SO_Class.Log.Println(true, "Execute ", err)
		return sql, err
	}
	sql, err = execTBLSLF(tx, userName, sqlstm)
	if err != nil {
		SO_Class.Log.Println(true, "Execute ", err)
		return sql, err
	}

	return sql, err
}

// A Txfn is a function that will be called with an initialized `Transaction` object
// that can be used for executing statements and queries against a database.
type TxFn func(Transaction) TransactionResult

// ExecQueryMultiple creates a new transaction and handles rollback/commit based on the
// error object returned by the `TxFn`
// conn fungsinya untuk koneksi ke database manapun kalau kirim nil maka sesuai local setting
func (saya db) ExecQueryMultiple(c *gin.Context, userName string, fn TxFn) (err error) {
	globalDB, _ := c.Get("globalDB")
	conn, err := openConnection(globalDB.(string))
	if err != nil {
		return err
	} else {
		defer conn.Close()
	}

	tx, err := conn.Begin()
	if err != nil {
		return
	}
	SO_Class.Log.Println(false, "func ExecQueryMultiple --> ", "Begin Transaction")
	var data TransactionResult
	defer func() {

		if p := recover(); p != nil {
			SO_Class.Log.Println(false, "func ExecQueryMultiple --> ", "RollBack Transaction -->", p)
			// a panic occurred, rollback and repanic
			tx.Rollback()
			execTBLELF(c, data.Sintax, data.Error.Error())
			SO_Class.Log.Println(false, "Error ExecTBLELF", p)
		} else if err != nil {
			SO_Class.Log.Println(false, "func ExecQueryMultiple --> ", err)
			SO_Class.Log.Println(false, "func ExecQueryMultiple --> ", "RollBack Transaction")
			// something went wrong, rollback
			tx.Rollback()
			execTBLELF(c, data.Sintax, data.Error.Error())
		} else {
			SO_Class.Log.Println(false, "func ExecQueryMultiple --> ", "Commit Transaction")
			// all good, commit
			err = tx.Commit()
		}
	}()

	data = fn(tx)
	err = data.Error
	return err
}

// Exported instance
var Database db
