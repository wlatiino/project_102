-----------------------------
-- begin seeding initial data
-----------------------------


    select stptblnor('sysadmin', 'tblmnu');
    insert into tblmnu ( 
        tmmenuiy,                
        tmnomr,tmgrup,tmmenu,tmscut,tmaces,tmbcdt,tmfwdt,tmurlw,tmsyfg,tmmntp,
        tmrgid,tmrgdt,tmchid,tmchdt,tmchno,
        tmdlfg,tmcsdt,tmcsid,tmsrce,tmunix
    ) values (
        '1','01','','FILE','','','99','99','','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        '2',                
        '0105','form01','SINTAX LOG FILE','FIL005','VXL','99','99','tblslf','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        '3','0110','form01','ERROR LOG FILE','FIL010','VXL','99','99','tblelf','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        '4','0113','form01','PROCESS BATCH LOG FILE','FIL013','VEXL','99','99','tblplf','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        '5','0115','','MENU','FIL015','VAEDLXP','99','99','tblmnu','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        '6','0120','','SET USER','FIL020','VAEDLXP','99','99','tblusr','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        '7','0123','','SET USER DISABLE/ENABLE','FIL023','VAELXP','99','99','tblued','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        '8','0125','','SET USER ACCESS MENU','FIL025','VAX','99','99','tbluah','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        '9','0190','','UPLOAD DATA TABLE','FIL090','VOL','99','99','upload','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        '10','02','','MASTER','','','99','99','','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        '11','0205','','MASTER SYSTEM','','','99','99','','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        '12','020505','','TABLE PARAMETER','SYS005','VAEDLX','99','99','tblprm','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        '13','020510','','TABLE DESCRIPTION *','SYS010','VAEDLX','99','99','tbldsc','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        '14','020515','','TABLE SYSTEM *','SYS015','VAEDLX','99','99','tblsys','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
    );
    update tblnor set tnnour = (select max(tmmenuiy) from tblmnu) where tntabl = 'tblmnu'; 



    select stptblnor('sysadmin', 'tblusr');
    insert into tblusr ( 
        tuuseriy,
        tuuser,tuname,tupswd,tuuslv,tuemid,tudept,tumail,tuwelc,tuexpp,tuexpd,tuexpv,tumntp,
        turgid,turgdt,tuchid,tuchdt,tuchno,
        tudlfg,tucsdt,tucsid,tusrce,tuunix
    ) values (
        '1',
        'admin','administrator','$2a$04$kqK.rSR1kgl5u7NljTpy3.3UUC0LdbKhGP.HSXAjsEtE4epg8UUsO','00','admin','IT','admin@it.com','','0','20301231','24','IE',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
    );
    update tblnor set tnnour = (select max(tuuseriy) from tblusr) where tntabl = 'tblusr';


    insert into tbluam (
        tauseriy, tamenuiy, taaces
        ,taremk,taitrm
        ,targid,targdt,tachid,tachdt,tachno,tadlfg,tadpfg,tadsfg,tasrce,tacsdt,tacsid
    )
    select tuuseriy, tmmenuiy, case when tuuser = 'admin' then tmaces else '' end taaces
        , '' taremk, '' taitrm
        ,'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,1,0,'pgsql',clock_timestamp(),'sysadmin'
    from tblmnu 
    left join tblusr on 1 = 1
    left join tbluam on tauseriy = tuuseriy and tamenuiy = tmmenuiy
    where tamenuiy is null;


    insert into tbldsc (
        tddscd,tddsnm,tdlgth,tdsyfg,
        tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
    ) values (
        'MODE','MODE HAK AKSES','1','1',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 
    insert into tblsys (
        tsdscd,tssycd,tssynm,
        tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
    ) values (
        'MODE','A','ADD',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'MODE','E','EDIT',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'MODE','D','DELETE',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'MODE','L','VIEW LIST',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'MODE','P','PRINT',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'MODE','R','APPROVAL',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'MODE','V','VIEW',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'MODE','X','EXPORT TO EXCEL',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'MODE','O','PROCESS',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 

---------------------------------------------------------------------------------------------------------------------------


    insert into tbldsc (
        tddscd,tddsnm,tdlgth,tdsyfg,
        tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
    ) values (
        'USLV','USER LEVL','2','1',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 
    insert into tblsys (
        tsdscd,tssycd,tssynm,
        tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
    ) values (
        'USLV','00','OWNER',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'USLV','10','SUPER ADMINISTRATOR',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'USLV','20','USER ADMINISTRATOR',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'USLV','30','USER',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 

---------------------------------------------------------------------------------------------------------------------------


    insert into tbldsc (
        tddscd,tddsnm,tdlgth,tdsyfg,
        tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
    ) values (
        'MNTP','MENU TYPE','3','1',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 
    insert into tblsys (
        tsdscd,tssycd,tssynm,
        tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
    ) values (
        'MNTP','I','INTERNAL',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'MNTP','E','EXTERNAL',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 

---------------------------------------------------------------------------------------------------------------------------

    insert into tbldsc (
        tddscd,tddsnm,tdlgth,tdsyfg,
        tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
    ) values (
        'AKTV','Aktive','1','1',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 
    insert into tblsys (
        tsdscd,tssycd,tssynm,
        tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
    ) values (
        'AKTV','0','Not Active',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'AKTV','1','Active',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 

---------------------------------------------------------------------------------------------------------------------------

    insert into tbldsc (
        tddscd,tddsnm,tdlgth,tdsyfg,
        tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
    ) values (
        'DSPLY','DISPLAY','1','1',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 
    insert into tblsys (
        tsdscd,tssycd,tssynm,
        tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
    ) values (
        'DSPLY','0','Not Display',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'DSPLY','1','Display',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 

---------------------------------------------------------------------------------------------------------------------------

    insert into tbldsc (
        tddscd,tddsnm,tdlgth,tdsyfg,
        tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
    ) values (
        'ET','ERROR TYPE','2','1',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 
    insert into tblsys (
        tsdscd,tssycd,tssynm,
        tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
    ) values (
        'ET','AP','APPLICATION',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'ET','SP','STORE PROCEDURE',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 

---------------------------------------------------------------------------------------------------------------------------

    insert into tbldsc (
        tddscd,tddsnm,tdlgth,tdsyfg,
        tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
    ) values (
        'STAS','STATUS','1','1',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 
    insert into tblsys (
        tsdscd,tssycd,tssynm,
        tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
    ) values (
        'STAS','0','Not Available',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'STAS','1','Available',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 

---------------------------------------------------------------------------------------------------------------------------

    insert into tbldsc (
        tddscd,tddsnm,tdlgth,tdsyfg,
        tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
    ) values (
        'YESNO','YES NO','1','1',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 
    insert into tblsys (
        tsdscd,tssycd,tssynm,
        tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
    ) values (
        'YESNO','0','NO',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'YESNO','1','YES',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 

---------------------------------------------------------------------------------------------------------------------------

    insert into tbldsc (
        tddscd,tddsnm,tdlgth,tdsyfg,
        tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
    ) values (
        'BPMSG','Batch Proses Messsage','1','1',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 
    insert into tblsys (
        tsdscd,tssycd,tssynm,
        tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
    ) values (
        'BPMSG','1','Initial Process',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'BPMSG','2','OnGoing Process',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'BPMSG','3','Done Process',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'BPMSG','4','Error Process',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 

---------------------------------------------------------------------------------------------------------------------------

    insert into tbldsc (
        tddscd,tddsnm,tdlgth,tdsyfg,
        tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
    ) values (
        'UPLOAD_TABLE','List Table','15','1',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 
    insert into tblsys (
        tsdscd,tssycd,tssynm,
        tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
    ) values (
        'UPLOAD_TABLE','tblsys','Table System',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'UPLOAD_TABLE','hlispl','LIMIT SPL INDIVIDU',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'UPLOAD_TABLE','hlgspl','LIMIT SPL GABUNGAN',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 

---------------------------------------------------------------------------------------------------------------------------

    insert into tbldsc (
        tddscd,tddsnm,tdlgth,tdsyfg,
        tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
    ) values (
        'STED','Enable / Disable','1','1',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 
    insert into tblsys (
        tsdscd,tssycd,tssynm,
        tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
    ) values (
        'STED','E','Enable',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'STED','D','Disable',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 


---------------------------------------------------------------------------------------------------------------------------

    insert into tbldsc (
        tddscd,tddsnm,tdlgth,tdsyfg,
        tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
    ) values (
        'AUTH_FG','AUTHORIZED FLAG','1','1',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 
    insert into tblsys (
        tsdscd,tssycd,tssynm,
        tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
    ) values (
        'AUTH_FG','-','UnAuthorized (-)',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ), (
        'AUTH_FG','+','Authorized (+)',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,0,clock_timestamp(),'sysadmin','pgsql',''
    ); 


/*


  select 
  	concat('','
            insert into tbldsc (
                tddscd,tddsnm,tdlgth,tdsyfg,
                tdrgid,tdrgdt,tdchid,tdchdt,tdchno,tddlfg,tdcsdt,tdcsid,tdsrce,tdunix
            ) 
            values (
                ''',rtrim(tddscd),''',''',rtrim(tddsnm),''',''',tdlgth,''',''',tdsyfg,''',
                ''sysadmin'',clock_timestamp(),''sysadmin'',clock_timestamp(),0,0,clock_timestamp(),''sysadmin'',''pgsql'',''''
            ); 
  ')
  from tbldsc 
--  where tddscd in ('MODE','USLV','MNTP','AKTV','DSPLY','ET','STAS','YESNO','BPMSG','UPLOAD_TABLE','STED');
  where tddscd in ('AUTH_FG');
  
  select 
  	concat('','
            insert into tblsys (
                tsdscd,tssycd,tssynm,
                tsrgid,tsrgdt,tschid,tschdt,tschno,tsdlfg,tscsdt,tscsid,tssrce,tsunix
            ) 
            values (
                ''',rtrim(tsdscd),''',''',rtrim(tssycd),''',''',rtrim(tssynm),''',
                ''sysadmin'',clock_timestamp(),''sysadmin'',clock_timestamp(),0,0,clock_timestamp(),''sysadmin'',''pgsql'',''''
            ); 
  ')
  from tblsys 
--  where tsdscd in ('MODE','USLV','MNTP','AKTV','DSPLY','ET','STAS','YESNO','BPMSG','UPLOAD_TABLE','STED','AUTH_FG');
  where tsdscd in ('AUTH_FG');
  
*/