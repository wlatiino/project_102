
call stpCreateTable('{
"table_name" : "mitmas",
"comment" : "Table Master Item",
"columns" : {
	"mmitnoiy" : "Item IY||int||not null",
	"mmitno"   : "Item No||bpchar(50)||not null",
	"mmitds"   : "Item Description||bpchar(200)||not null",
	"mmitcl"   : "Item Category||bpchar(10)||not null",
	"mmitgr"   : "Item Group||bpchar(10)||not null",
	"mmitty"   : "Item Type||bpchar(10)||not null",
	"mmitbr"   : "Item Brand||bpchar(10)||not null",
	"mmunms"   : "Unit Measurement||bpchar(5)||not null",
	"mmrums"   : "Rumus||bpchar(2)||not null"
},
"primary" : "mmitnoiy",
"unique" : ["mmitno"]
}');


-------------------------------------------------------------------------------------------------------------------

insert into tbldsc (
	tddscd,tddsnm,tdlgth,tdsyfg,
	tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
) values (
	'ITCL','Item Category','10','1',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 
insert into tblsys (
	tsdscd,tssycd,tssynm,
	tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
) values (
	'ITCL','NA','NOT AVAILABLE',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 

-------------------------------------------------------------------------------------------------------------------

insert into tbldsc (
	tddscd,tddsnm,tdlgth,tdsyfg,
	tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
) values (
	'ITGR','Item Group','10','1',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 
insert into tblsys (
	tsdscd,tssycd,tssynm,
	tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
) values (
	'ITGR','NA','NOT AVAILABLE',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 

-------------------------------------------------------------------------------------------------------------------

insert into tbldsc (
	tddscd,tddsnm,tdlgth,tdsyfg,
	tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
) values (
	'ITTY','Item Type','10','1',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 
insert into tblsys (
	tsdscd,tssycd,tssynm,
	tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
) values (
	'ITTY','NA','NOT AVAILABLE',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 

-------------------------------------------------------------------------------------------------------------------

insert into tbldsc (
	tddscd,tddsnm,tdlgth,tdsyfg,
	tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
) values (
	'ITBR','Item Brand','10','1',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 
insert into tblsys (
	tsdscd,tssycd,tssynm,
	tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
) values (
	'ITBR','NA','NOT AVAILABLE',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 

-------------------------------------------------------------------------------------------------------------------

insert into tbldsc (
	tddscd,tddsnm,tdlgth,tdsyfg,
	tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
) values (
	'UNMS','Unit Measurement','5','1',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 
insert into tblsys (
	tsdscd,tssycd,tssynm,
	tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
) values (
	'UNMS','NA','NOT AVAILABLE',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
), (
	'UNMS','PCS','PIECES',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 

-------------------------------------------------------------------------------------------------------------------

insert into tbldsc (
	tddscd,tddsnm,tdlgth,tdsyfg,
	tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
) values (
	'MITMAS_RUMS','Rumus Perhitungan Stock','2','1',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 
insert into tblsys (
	tsdscd,tssycd,tssynm,
	tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
) values (
	'MITMAS_RUMS','NL','NORMAL',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
), (
	'MITMAS_RUMS','SL','SERIAL',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
), (
	'MITMAS_RUMS','SO','SALDO',
	'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
); 

/*
	mmrums :
		Normal
		Serial
		Saldo
*/

