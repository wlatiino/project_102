
call stpCreateTable('{
"table_name" : "ivline",
"comment" : "Table Inventory Line",
"columns" : {
	"ililnoiy" : "IY||int||not null",
	"ililno" : "Line No||int||not null",
	"ilihnoiy" : "IY||int||not null",
	"ilitnoiy" : "Item IY||int||not null",
	"ilqtys" : "qtys||decimal(24,10)||not null",
	"ilharg" : "harga||decimal(24,10)||not null",
	"iltotl" : "total||decimal(24,10)||not null"
},
"primary" : "ililnoiy",
"unique" : ["ilihnoiy||ililno","ilihnoiy||ilitnoiy"],
"foreign" : {
	"ilihnoiy" : "ivhead||ihihnoiy",
	"ilitnoiy" : "mitmas||mmitnoiy"
}
}');

/*

*/

