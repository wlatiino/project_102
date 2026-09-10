
call stpCreateTable('{
"table_name" : "ivhead",
"comment" : "Table Inventory Header",
"columns" : {
	"ihihnoiy" : "IY||int||not null",
	"ihtrno" : "No Transaksi||bpchar(50)||not null",
	"ihtrdt" : "Tanggal Transaksi||bpchar(8)||not null",
	"ihtype" : "Tipe Transaksi||bpchar(5)||not null"
},
"primary" : "ihihnoiy",
"unique" : ["ihtrno"]
}');

-------------------------------------------------------------------------------------------------------------------

insert into tbldsc (
	tddscd,tddsnm,tdlgth,tdsyfg,
	tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
) values (
	'IVHEAD_TYPE','Inventory Type','5','1',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 
insert into tblsys (
	tsdscd,tssycd,tssynm,
	tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
) values (
	'IVHEAD_TYPE','AJI','Adjustment In',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
), (
	'IVHEAD_TYPE','AJO','Adjustment Out',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 


/*

*/

