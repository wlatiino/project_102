
call stpCreateTable('{
"table_name" : "coline",
"comment" : "Table Customer Order Line",
"columns" : {
	"cdcdnoiy" : "IY||int||not null",
	"cdcdno" : "Line No||int||not null",
	"cdchnoiy" : "IY||int||not null",
	"cditnoiy" : "Item IY||int||not null",
	"cdqtys" : "Qty||decimal(24,10)||not null",
	"cdunms" : "Unit Measurement||bpchar(5)||not null",
	"cdorqt" : "Original Qty or Qty UM||decimal(24,10)||not null",
	"cdharg" : "Harga||decimal(24,10)||not null",
	"cdtotl" : "Line Amount||decimal(24,10)||not null"
},
"primary" : "cdcdnoiy",
"unique" : ["cdchnoiy||cdcdno"],
"foreign" : {
	"cdchnoiy" : "cohead||chchnoiy",
	"cditnoiy" : "mitmas||mmitnoiy"
}
}');

/*

*/

