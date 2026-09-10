
call stpCreateTable('{
"table_name" : "cohead",
"comment" : "Table Customer Order Header",
"columns" : {
	"chchnoiy" : "IY||int||not null",
	"chcono" : "No PO||bpchar(50)||not null",
	"chcodt" : "Tanggal PO||bpchar(8)||not null",
	"chrfno" : "Reference No||bpchar(50)||not null",
	"chbpnoiy" : "Bisnis Partner IY||int||not null",
	"chvtcd" : "Vat Code||bpchar(5)||not null",
	"chsubt" : "Sub Total||decimal(24,10)||not null",
	"chvtam" : "Vat Amount||decimal(24,10)||not null",
	"chtotl" : "Total||decimal(24,10)||not null"
},
"primary" : "chchnoiy",
"unique" : ["chcono"],
"foreign" : {
	"chbpnoiy" : "mbpmas||bpbpnoiy"
}
}');

-------------------------------------------------------------------------------------------------------------------


/*

*/

