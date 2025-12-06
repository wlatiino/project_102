--------------------------
-- begin create core table
--------------------------

call stpCreateTable('{
"table_name" : "tblslf",
"comment" : "Table Sintax Logfile",
"columns" : {
	"tqnomriy" : "IY||serial||not null",
	"tquser" : "User||bpchar(50)||not null",
	"tqstmt" : "SQL Statement||text||not null"
},
"primary" : "tqnomriy"
}');

call stpCreateTable('{
"table_name" : "tblxlf",
"comment" : "Table Sintax Logfile Exec",
"columns" : {
	"txnomriy" : "IY||serial||not null",
	"txuser" : "User||bpchar(50)||not null",
	"txstmt" : "SQL Statement||text||not null"
},
"primary" : "txnomriy"
}');

call stpCreateTable('{
"table_name" : "tblelf",
"comment" : "Table Error Logfile",
"columns" : {
	"tenomriy" : "IY||serial||not null",
	"teuser" : "User||bpchar(50)||not null",
	"teerst" : "Error State||bpchar(50)||not null",
	"teerms" : "Error Message||text||not null",
	"testmt" : "SQL Statement||text||not null"
},
"primary" : "tenomriy"
}');

call stpCreateTable('{
"table_name" : "tblnor",
"comment" : "Table Running No",
"columns" : {
	"tnnomriy" : "IY||serial||not null",
	"tntabl" : "Table Name||bpchar(20)||not null",
	"tnnour" : "Last Running No||int||not null"
},
"primary" : "tnnomriy",
"unique" : ["tntabl"]
}');

call stpCreateTable('{
"table_name" : "tblprm",
"comment" : "Table Parameter",
"columns" : {
	"trprcd" : "Parameter Code||bpchar(20)||not null",
    "trprnm" : "Parameter Name||bpchar(200)||not null",
    "trsyv1" : "Value 1||decimal(24,10)||null",
    "trsyv2" : "Value 2||decimal(24,10)||null",
    "trsyv3" : "Value 3||decimal(24,10)||null",
    "trsyt1" : "Text 1||text||null",
    "trsyt2" : "Text 2||text||null",
    "trsyt3" : "Text 3||text||null"
},
"primary" : "trprcd"
}');

call stpCreateTable('{
"table_name" : "tbldsc",
"comment" : "Table Description",
"columns" : {
	"tddscd" : "Description Code||bpchar(20)||not null",
    "tddsnm" : "Description Name||bpchar(200)||not null",
    "tdlgth" : "Character Length||int||not null",
    "tdsyfg" : "System Flag||bit(1)||not null"
},
"primary" : "tddscd"
}');

call stpCreateTable('{
"table_name" : "tblsys",
"comment" : "Table System",
"columns" : {
	"tsnomriy" : "IY||serial||not null",
	"tsdscd" : "Description Code||bpchar(20)||not null",
	"tssycd" : "System Code||bpchar(20)||not null",
    "tssynm" : "System Name||bpchar(200)||not null",
    "tssyv1" : "Value 1||decimal(24,10)||null",
    "tssyv2" : "Value 2||decimal(24,10)||null",
    "tssyv3" : "Value 3||decimal(24,10)||null",
    "tssyt1" : "Text 1||text||null",
    "tssyt2" : "Text 2||text||null",
    "tssyt3" : "Text 3||text||null",
    "tslsv1" : "Label Value 1||bpchar(200)||null",
    "tslsv2" : "Label Value 2||bpchar(200)||null",
    "tslsv3" : "Label Value 3||bpchar(200)||null",
    "tslst1" : "Label Text 1||bpchar(200)||null",
    "tslst2" : "Label Text 2||bpchar(200)||null",
    "tslst3" : "Label Text 3||bpchar(200)||null"
},
"primary" : "tsnomriy",
"unique" : ["tsdscd||tssycd"],
"foreign" : {
	"tsdscd" : "tbldsc||tddscd"
}
}');

call stpCreateTable('{
"table_name" : "tblmnu",
"comment" : "Table Menu",
"columns" : {
	"tmmenuiy" : "System Menu IY||int||not null",
	"tmnomr" : "Nomor Urut||bpchar(20)||not null",
	"tmgrup" : "Kelompok Menu||bpchar(250)||null",
	"tmmenu" : "Menu||bpchar(200)||not null",
	"tmdesc" : "Menu Deskirpsi||text||null",
	"tmscut" : "Short Cut||bpchar(20)||null",
	"tmaces" : "Menu Akses||bpchar(20)||null",
	"tmbcdt" : "BackDate||int||null",
	"tmfwdt" : "ForwardDate||int||null",
	"tmurlw" : "URL||bpchar(200)||null",
	"tmsyfg" : "System Flag||bpchar(10)||null",
	"tmmntp" : "Menu Type||bpchar(10)||null",
	"tmusct" : "User Hit Count||int||null",
	"tmlsdt" : "User Hit Last Date||timestamp(3)||null",
	"tmlsby" : "User Hit Last By||bpchar(50)||null",
	"tmrldt" : "Release Date||bpchar(8)||null",
	"tmgrid" : "Grid Query||text||null",
	"tmjson" : "JSON||json||null"
},
"primary" : "tmmenuiy",
"unique" : ["tmnomr"]
}');

call stpCreateTable('{
"table_name" : "tblusr",
"comment" : "Table User",
"columns" : {
	"tuuseriy" : "IY||int||not null",
	"tuuser" : "User Login||bpchar(50)||not null",
	"tuname" : "User Name||bpchar(100)||not null",
	"tupswd" : "Password||varchar(100)||null",
	"tuepin" : "Pin||varchar(100)||null",
	"tuuslv" : "User Level||bpchar(100)||not null",
	"tuemid" : "Employee ID||bpchar(50)||null",
	"tudept" : "Department||bpchar(100)||null",
	"tumail" : "Mail||bpchar(100)||null",
	"tuwelc" : "Welcome Text||text||null",
	"tuexpp" : "Expired||bit(1)||null",
	"tuexpd" : "Expired Date||bpchar(8)||null",
	"tuexpv" : "Expired Value||int||null",
	"tumntp" : "Menu Type||bpchar(10)||null",
	"tulgct" : "Login Counter||int||null",
	"tulsli" : "Last Login||timestamp(3)||null",
	"tulslo" : "Last Logoff||timestamp(3)||null",
	"tutokn" : "token||text||null"
},
"primary" : "tuuseriy",
"unique" : ["tuuser","tuemid"]
}');

call stpCreateTable('{
"table_name" : "tblush",
"comment" : "Table User Other Access",
"columns" : {
	"tvuseriy" : "IY||int4||not null"
},
"primary" : "tvuseriy",
"foreign" :{
	"tvuseriy" : "tblusr||tuuseriy"
}
}');

call stpCreateTable('{
"table_name" : "tbluph",
"comment" : "Table User Password History",
"columns" : {
	"tpnomriy" : "IY||serial||not null",
	"tpuser" : "User Login||bpchar(50)||not null",
	"tppswd" : "Password||varchar(100)||not null"
},
"primary" : "tpnomriy"
}');

call stpCreateTable('{
"table_name" : "tblued",
"comment" : "Table User Enable Disable History",
"columns" : {
	"tenomriy" : "IY||int||not null",
	"teuseriy" : "User IY||int||not null",
	"tested" : "Status||bpchar(1)||not null",
	"tereas" : "Reason||text||not null"
},
"primary" : "tenomriy",
"foreign" :{
	"teuseriy" : "tblusr||tuuseriy"
}
}');

call stpCreateTable('{
"table_name" : "tbluam",
"comment" : "Table User Access Menu",
"columns" : {
	"tanomriy" : "IY||serial||not null",
	"tauseriy" : "tblusr IY||int||not null",
	"tamenuiy" : "tblmnu IY||int||not null",
	"taaces" : "access menu||bpchar(20)||null",
	"talsdt" : "last date use||timestamp(3)||null",
	"tausct" : "use count||int||null",
	"tafavo" : "favourite||bit(1)||null"
},
"primary" : "tanomriy",
"unique" : ["tauseriy||tamenuiy"],
"foreign" :{
	"tauseriy" : "tblusr||tuuseriy",
	"tamenuiy" : "tblmnu||tmmenuiy"
}
}');

call stpCreateTable('{
"table_name" : "tbluah",
"comment" : "Table User Access Menu - History",
"columns" : {
	"tbnomriy" : "IY||serial||not null",
	"tbuseriy" : "TBLUSR IY||int||not null",
	"tbmenuiy" : "TBLMNU IY||int||not null",
	"tbaces" : "Access Menu User||bpchar(2)||null",
	"tbacec" : "Access Menu Creator||bpchar(20)||null"
},
"primary" : "tbnomriy",
"foreign" :{
	"tbuseriy" : "tblusr||tuuseriy",
	"tbmenuiy" : "tblmnu||tmmenuiy"
}
}');

call stpCreateTable('{
"table_name" : "tblplf",
"comment" : "Table Proses Batch Logfile",
"columns" : {
	"tznomriy" : "IY||serial||not null",
	"tznotr" : "No Transaksi||bpchar(50)||not null",
	"tzmenuiy" : "TBLMNU IY||int||not null",
	"tzkeyc" : "Key Code||text||not null",
	"tzkeyd" : "Key Duplication||text||not null",
	"tzstmt" : "SqlStatement||text||null",
	"tzstat" : "Status||bpchar(2)||not null",
	"tzsttm" : "Start Date||timestamp(3)||not null",
	"tzentm" : "End Date||timestamp(3)||null",
	"tzjbid" : "Lumen Job ID||int||not null"
},
"primary" : "tznomriy",
"unique" : ["tznotr"],
"foreign" :{
	"tzmenuiy" : "tblmnu||tmmenuiy"
}
}');

call stpCreateTable('{
"table_name" : "tblhsl",
"comment" : "Table History Login",
"columns" : {
	"thnomriy" : "IY||serial||not null",
	"thuser" : "User||bpchar(50)||not null",
	"thipad" : "IP Address||bpchar(50)||not null",
	"thpcnm" : "PC Name||bpchar(100)||not null",
	"thlsli" : "Login Date||timestamp(3)||not null",
	"thlslo" : "Logout Date||timestamp(3)||null",
	"thdvvs" : "Device System||bpchar(100)||not null",
	"thsyvs" : "Device Version||bpchar(100)||not null"
},
"primary" : "thnomriy"
}');

call stpCreateTable('{
"table_name" : "csynbr",
"comment" : "Table Running Transaction No",
"columns" : {
	"cnnomriy" : "No Running IY||serial||not null",
	"cncode" : "Code||bpchar(20)||not null",
	"cnnbty" : "Modul Name||bpchar(5)||not null",
	"cnnbnr" : "Last Number Used||bpchar(8)||not null",
	"cnnomr" : "Last Transaction No||bpchar(100)||not null",
	"cnyear" : "Year||bpchar(4)||not null",
	"cnmnth" : "Month||bpchar(2)||not null"
},
"primary" : "cnnomriy",
"unique" : ["cncode||cnnbty||cnyear||cnmnth"]
}');

call stpCreateTable('{
  "table_name": "tblhss",
  "comment": "Table History Sintax",
  "columns": {
    	"tlnomriy": "IY||serial||not null",
    	"tlevtp": "Event Type||text||not null",
    	"tlschm": "Schema||text||not null",
    	"tlobtp": "Object Type||text||not null",
    	"tlobnm": "Object Name||text||not null",
    	"tlstmt": "Sintax||text||not null"
},
"primary": "tlnomriy"
}');

