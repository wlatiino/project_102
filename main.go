package main

import (
	SO_Class "SOApp_GO/class"
	SO_Form "SOApp_GO/form"
)

func init() {
	SO_Class.Fmt.Println(true, "Masuk Main Init")
}
func main() {
	SO_Class.Fmt.Println(true, "Masuk Main Main")
	SO_Form.Rute.Mulai()
}
