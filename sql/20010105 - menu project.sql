
	-- select * from tblmnu order by tmnomr;

    insert into tblmnu ( 
        tmmenuiy,                
        tmnomr,tmgrup,tmmenu,tmscut,tmaces,tmbcdt,tmfwdt,tmurlw,tmsyfg,tmmntp,
        tmrgid,tmrgdt,tmchid,tmchdt,tmchno,
        tmdlfg,tmcsdt,tmcsid,tmsrce,tmunix
    )     
	select 
        (select max(tmmenuiy) from tblmnu)+(row_number() over()), * 
    from (
        select
        '0210','','MASTER GENERAL','','',99,99,'','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
        union all select
        '021005','','ITEM','GEN005','VAEDLX',99,99,'mitmas','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
        union all select
        '021010','','PERSONAL KONTRAK','GEN010','VAEDLX',99,99,'mbpmas','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
        union all select
        '10','','TRANSACTION','','',99,99,'','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
        union all select
        '1005','','INVENTORY','INV005','VAEDLX',99,99,'ivhead','W','I',
        'sysadmin',clock_timestamp(),'sysadmin',clock_timestamp(),0,
        0,clock_timestamp(),'sysadmin','pgsql',''
	) semua;      
    update tblnor set tnnour = (select max(tmmenuiy) from tblmnu) where tntabl = 'tblmnu'; 


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
