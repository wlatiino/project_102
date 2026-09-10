
call stpCreateTable('{
"table_name" : "pohead",
"comment" : "Table Purchase Order Header",
"columns" : {
	"phphnoiy" : "IY||int||not null",
	"phpono" : "No PO||bpchar(50)||not null",
	"phpodt" : "Tanggal PO||bpchar(8)||not null",
	"phrfno" : "Reference No||bpchar(50)||not null",
	"phbpnoiy" : "Bisnis Partner IY||int||not null",
	"phvtcd" : "Vat Code||bpchar(5)||not null",
	"phsubt" : "Sub Total||decimal(24,10)||not null",
	"phvtam" : "Vat Amount||decimal(24,10)||not null",
	"phtotl" : "Total||decimal(24,10)||not null"
},
"primary" : "phphnoiy",
"unique" : ["phpono"],
"foreign" : {
	"phbpnoiy" : "mbpmas||bpbpnoiy"
}
}');

-------------------------------------------------------------------------------------------------------------------


/*

*/

