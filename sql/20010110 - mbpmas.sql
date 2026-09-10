
call stpCreateTable('{
"table_name" : "mbpmas",
"comment" : "Table Master Bisnis Partner",
"columns" : {
	"bpbpnoiy" : "Bisnis Partner IY||int||not null",
	"bpbpno" : "No||bpchar(50)||not null",
	"bpbpnm" : "Name||bpchar(100)||not null",
	"bpbpds" : "Description||text||not null",
	"bpaddr" : "Address||text||not null",
	"bpmail" : "Email||bpchar(200)||not null",
	"bptelp" : "Telphone||bpchar(200)||not null",
	"bpvtcd" : "Vat Code||bpchar(5)||not null"
},
"primary" : "mmitnoiy",
"unique" : ["mmitno"]
}');


-------------------------------------------------------------------------------------------------------------------

insert into tbldsc (
	tddscd,tddsnm,tdlgth,tdsyfg,
	tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
) values (
	'VTCD','VAT CODE','5','1',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 
insert into tblsys (
	tsdscd,tssycd,tssynm,tssyv1,
	tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
) values (
	'VTCD','10','10%','10',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
), (
	'VTCD','11','11%','11',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
), (
	'VTCD','12','12%','11',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 


/*
	mmrums :
		Normal
		Serial
		Saldo
*/

