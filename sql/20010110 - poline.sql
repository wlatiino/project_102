
call stpCreateTable('{
"table_name" : "poline",
"comment" : "Table Purchase Order Line",
"columns" : {
	"pdpdnoiy" : "IY||int||not null",
	"pdpdno" : "Line No||int||not null",
	"pdphnoiy" : "IY||int||not null",
	"pditnoiy" : "Item IY||int||not null",
	"pdqtys" : "Qty||decimal(24,10)||not null",
	"pdunms" : "Unit Measurement||bpchar(5)||not null",
	"pdorqt" : "Original Qty or Qty UM||decimal(24,10)||not null",
	"pdharg" : "Harga||decimal(24,10)||not null",
	"pdtotl" : "Line Amount||decimal(24,10)||not null"
},
"primary" : "pdpdnoiy",
"unique" : ["pdphnoiy||pdpdno"],
"foreign" : {
	"pdphnoiy" : "pohead||phphnoiy",
	"pditnoiy" : "mitmas||mmitnoiy"
}
}');

/*

*/

