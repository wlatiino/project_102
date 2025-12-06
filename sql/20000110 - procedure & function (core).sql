-----------------------------------------
-- begin create core function / procedure
-----------------------------------------

create or replace function fntSplitString (
   string varchar,
   delimeter varchar
) 
returns table ( splitdata text ) 
language plpgsql
as $$
begin
	return query select unnest(string_to_array(string,delimeter))::text;
end
$$;

create or replace procedure StpCreateTable(json_table json)
language plpgsql     
as $$
declare
	table_name 		varchar;
	table_columns 	varchar;
	sintax 			text;
	sintax_comment 	text;
	sintax_unique 	text;
	sintax_foreign 	text;
	_key   			text;
   	_value 			text;
    _prefix			char(2);
    _elem			text;
begin
	table_name = json_table->>'table_name';

	for _key, _value in
       select * from json_each_text(cast(json_table->>'columns' as json))
    loop
       table_columns := concat(table_columns,_key,' ',split_part(_value,'||',2),' ',split_part(_value,'||',3),',',E'\n');
       sintax_comment := concat(sintax_comment,'COMMENT ON COLUMN ',table_name,'.',_key,' IS ''',split_part(_value,'||',1),''';',E'\n');
    end loop;

   	_prefix = left(table_columns,2);

	sintax = concat('CREATE TABLE ',current_schema(),'.',table_name,' (
		',table_columns,'
		',_prefix,'remk text NULL,
		',_prefix,'rgid bpchar(50) NOT NULL,
		',_prefix,'rgdt timestamp(3) NOT NULL,
		',_prefix,'chid bpchar(50) NOT NULL,
		',_prefix,'chdt timestamp(3) NOT NULL,
		',_prefix,'chno int4 NOT NULL DEFAULT 0,
		',_prefix,'dlfg bpchar(1) NOT NULL DEFAULT ''0''::bpchar,
		',_prefix,'dpfg bpchar(1) NOT NULL DEFAULT ''1''::bpchar,
		',_prefix,'dsfg bpchar(1) NULL DEFAULT ''0''::bpchar,
		',_prefix,'ptfg bpchar(1) NULL DEFAULT ''0''::bpchar,
		',_prefix,'ptct int4 NULL DEFAULT 0,
		',_prefix,'ptid bpchar(50) NULL,
		',_prefix,'ptdt timestamp(3) NULL,
		',_prefix,'srce bpchar(50) NULL,
		',_prefix,'usrm text NULL,
		',_prefix,'itrm text NULL,
		',_prefix,'csdt timestamp(3) NOT NULL,
		',_prefix,'csid bpchar(50) NOT NULL,
		',_prefix,'csno bpchar(50) NULL,
		',_prefix,'unix bpchar(50) NULL,
		CONSTRAINT ',table_name,'_pkey PRIMARY KEY (',json_table->>'primary',')
	);');
--	raise notice '%', sintax;
	execute sintax;

	sintax = concat('
		COMMENT ON TABLE ',table_name,' IS ''',json_table->>'comment',''';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'remk IS ''Remark'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'rgid IS ''Add By'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'rgdt IS ''Add On'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'chid IS ''Change By'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'chdt IS ''Change On'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'chno IS ''Change No'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'dlfg IS ''Delete Flag'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'dpfg IS ''Display Flag'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'dsfg IS ''Disable Flag'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'ptfg IS ''Print Flag'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'ptct IS ''Print Count'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'ptid IS ''Print By'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'ptdt IS ''Print Date'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'srce IS ''Source'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'usrm IS ''Internal Use Remark'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'itrm IS ''IT Remark'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'csid IS ''Change System By'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'csdt IS ''Change System Date'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'csno IS ''Change System No'';
		COMMENT ON COLUMN ',table_name,'.',_prefix,'unix IS ''Unix Index'';
	');
--	raise notice '%', sintax;
	execute sintax;

--	raise notice '%', sintax_comment;
	execute sintax_comment;

	if json_table->>'unique' is not null then
		for _elem in 
			select json_array_elements_text(cast(json_table->>'unique' as json)) 
		loop
			sintax_unique := concat(sintax_unique,'ALTER TABLE ',table_name,' ADD CONSTRAINT ',table_name,'_',replace(_elem,'||','_'),'_unique UNIQUE (',replace(_elem,'||',','),');',E'\n');
		end loop;
	
--		raise notice '%', sintax_unique;
		execute sintax_unique;
	end if;

	if json_table->>'foreign' is not null then
		for _key, _value in
	       	select * from json_each_text(cast(json_table->>'foreign' as json))
	    loop
	      	sintax_foreign := concat(sintax_foreign,'ALTER TABLE ',table_name,' ADD CONSTRAINT ',table_name,'_',_key,'_foreign FOREIGN KEY (',_key,') REFERENCES ',split_part(_value,'||',1),'(',split_part(_value,'||',2),');',E'\n');
	    end loop;

--		raise notice '%', sintax_foreign;
		execute sintax_foreign;
	end if;
end
$$;


create or replace function stpTBLNOR(
	_username	varchar,
	_tntabl	varchar
)
returns int
language plpgsql     
as $$
declare 
	_sqlstm	text;
	_nourut int := 0;
begin
	select coalesce(tnnour,0) into _nourut from tblnor where tntabl = _tntabl;

   	if ( coalesce(_nourut,0) = 0 ) then
    	_nourut = 1;
    
    	_sqlstm = concat('
			 insert into tblnor ( 
				  tntabl,tnnour,tnremk 
				 ,tnrgid,tnrgdt,tnchid,tnchdt,tnchno,tndlfg,tndpfg,tndsfg,tnptfg,tnptct,tnptid,tnptdt,tnsrce,tncsdt,tncsid 
			 ) 
			 values ( 
				 ''',_tntabl,''',''1'','''' 
				,''',_username,''',''',clock_timestamp(),''',''',_username,''',''',clock_timestamp(),''',0,0,1,0,0,0,'''',''',clock_timestamp(),''',''system'',''',clock_timestamp(),''',''system'' 
			 );
		');
	
		execute _sqlstm;    
	else 
    	_sqlstm = concat('update tblnor set tnnour = coalesce(tnnour,0)+1 where tntabl = ''',_tntabl,''';');
	    execute _sqlstm; 

		select coalesce(tnnour,0) into _nourut from tblnor where tntabl = _tntabl;
	end if;

	return _nourut;
--	raise notice '%', _nourut;
end
$$;

create or replace function FnsHitungMasaKerja(
	_tipe	varchar,
	_frdt	varchar,
	_todt 	varchar
)
returns varchar
language plpgsql     
as $$
declare 
	_hasil varchar;
begin
	select extract(year from age(_todt::date, _frdt::date)) || ' thn ' || 
       	   extract(month from age(_todt::date, _frdt::date)) || ' bln' into _hasil;

	return _hasil;
end
$$;

create or replace function fntGetListTanggal (
   perdfrom bpchar(8),
   perduntil bpchar(8)
) 
returns table ( tanggal text ) 
language plpgsql
as $$
begin
	return query 
	SELECT 
		to_char(generate_series, 'YYYYMMDD') as tanggal
	FROM generate_series(
	    perdfrom::date,
	    perduntil::date,
	    '1 day'::interval
	);
end
$$;

create or replace procedure stpTBLUED(
    _tipe      varchar,
    _username  varchar,
    _source    varchar,
    _sysdate   varchar,
    _data      json
)
language plpgsql     
as $$
declare 
    _sqlstm     text;
    _key        varchar;
    _val        varchar;
    _tblcol     text;
    _tblval     text;
    _tblcolval  text;
	_nomriy 	varchar;
	_uslvlogin	varchar;
	_uslv		varchar;
begin
   
    _tblcol = '';
    for _key, _val in
        select * from json_each_text(_data)
    loop
       
        if _key = 'tenomriy' then
            _nomriy = _val;
        else
            _tblcol := concat(_tblcol,_key,',');
            _tblval := concat(_tblval,'''',_val,''',');
            _tblcolval := concat(_tblcolval,_key,' = ''',_val,''',');
        end if;
    end loop;
   
    call stpCheckBFCS(_tipe,_username,(concat('{"table":"tblued","field":"tenomriy","value":"',_nomriy,'","sysdate":"',_sysdate,'"}'))::json);

	-- begin validasi status user
	if exists (
		select * from tblusr
		where tuuseriy = (_data->>'teuseriy')::int
			and (case when (_data->>'tested') = 'D' then '1' else '0' end) = tudsfg
	) then
		raise exception 'User status already %!', (case when (_data->>'tested') = 'D' then 'Disabled' else 'Enabled' end);
	end if;
	-- end validasi status user

	-- begin validasi user level
	select tuuslv into _uslvlogin from tblusr where tuuser = _username;
	select tuuslv into _uslv from tblusr
			where tuuseriy = (_data->>'teuseriy')::int;

	if _uslvlogin >= _uslv then
		raise exception 'Your user level % is not qualified!', _uslvlogin;
	end if;
	-- end validasi user level

    if _tipe = '1' then
        select stpTBLNOR(_username,'tblued') into _nomriy;
        
        _sqlstm = concat('
            insert into tblued (
                tenomriy,
                ',_tblcol,'
                tergid,tergdt,techid,techdt,techno,
                tedlfg,tecsdt,tecsid,tesrce,teunix
            ) 
            values (
                ''',_nomriy,''',
                ',_tblval,'
                ''',_username,''',''',clock_timestamp(),''',''',_username,''',''',clock_timestamp(),''',0,
                0,''',clock_timestamp(),''',''',_username,''',''',_source,''',''''
            );
        ');
        execute _sqlstm;
    elsif _tipe = '2' then
        _sqlstm = concat('
            update tblued set
                ',_tblcolval,'
                techid = ''',_username,''',techdt = ''',clock_timestamp(),''',
                techno = coalesce(techno,0)+1,tesrce = ''',_source,''',
                tecsid = ''',_username,''',tecsdt = ''',clock_timestamp(),''',
                teunix = ''''
            where tenomriy = ''',_nomriy,''';
        ');
        execute _sqlstm;
    elsif _tipe = '3' then  
        raise exception 'Permission Denied!';
    end if;

    if _tipe in ('1') then
 		_sqlstm = concat('
            update tblusr set
                tudsfg = ''',(case when (_data->>'tested') = 'D' then '1' else '0' end),''',
                tuchid = ''',_username,''',tuchdt = ''',clock_timestamp(),''',
                tuchno = coalesce(tuchno,0)+1,tusrce = ''',_source,''',
                tucsid = ''',_username,''',tucsdt = ''',clock_timestamp(),''',
                tuunix = ''''
            where tuuseriy = ''',(_data->>'teuseriy'),''';
        ');
        execute _sqlstm;
    end if;
    call stpTBLSLF(_username,_sqlstm);
        
end
$$;


create or replace function FnsDateDiff(
	_tipe 	   varchar,
	_start_date date, 
	_end_date   date
)
returns integer
language plpgsql     
as $$
begin
	if _tipe = 'M' then
    	return date_part('year', age(_end_date, _start_date)) * 12 +
           	   date_part('month', age(_end_date, _start_date));
	elsif _tipe = 'D' then
    	return (_end_date - _start_date);
	end if;
end
$$;

CREATE OR REPLACE FUNCTION fnsReplaceNewLine(text_input TEXT)
RETURNS JSON AS $$
BEGIN
  RETURN (REPLACE(REPLACE(text_input, chr(10), '\n'), chr(13), ''))::json;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION fnsConvertDateIndo(_date_str TEXT, _tipe varchar)
RETURNS TEXT AS $$
DECLARE
    _dt DATE;
    _day_num INT;
    _month_num INT;
    _day_name TEXT;
    _month_name TEXT;
BEGIN
    -- Convert input string to DATE
    _dt := TO_DATE(_date_str, 'YYYYMMDD');

    -- Get day of week (0 = Sunday, 1 = Monday, ..., 6 = Saturday)
    _day_num := EXTRACT(DOW FROM _dt);
    _month_num := EXTRACT(MONTH FROM _dt);

    -- Map day name
    _day_name := CASE _day_num
        WHEN 0 THEN 'Minggu'
        WHEN 1 THEN 'Senin'
        WHEN 2 THEN 'Selasa'
        WHEN 3 THEN 'Rabu'
        WHEN 4 THEN 'Kamis'
        WHEN 5 THEN 'Jumat'
        WHEN 6 THEN 'Sabtu'
    END;

    -- Map month name
    _month_name := CASE _month_num
        WHEN 1 THEN 'Januari'
        WHEN 2 THEN 'Februari'
        WHEN 3 THEN 'Maret'
        WHEN 4 THEN 'April'
        WHEN 5 THEN 'Mei'
        WHEN 6 THEN 'Juni'
        WHEN 7 THEN 'Juli'
        WHEN 8 THEN 'Agustus'
        WHEN 9 THEN 'September'
        WHEN 10 THEN 'Oktober'
        WHEN 11 THEN 'November'
        WHEN 12 THEN 'Desember'
    END;

    if _tipe = '1' then
	    -- Return the formatted Indonesian date
	    RETURN _day_name || ', ' || EXTRACT(DAY FROM _dt)::INT || ' ' || _month_name || ' ' || EXTRACT(YEAR FROM _dt)::INT;
	else
		RETURN 'Tipe Not Found!';
	end if;
END;
$$ LANGUAGE plpgsql;

create or replace function fnsFormatUang(
	_nilai	      numeric,
	_decimalPoint int,
	_tipe	      varchar
)
returns varchar
language plpgsql     
as $$
declare 
	_hasil varchar;
begin
	if _decimalPoint = 0 then
		select to_char(_nilai, 'FM999,999,999,990') into _hasil;
	else
		select to_char(_nilai, concat('FM999,999,999,990.', repeat('0',_decimalPoint))) into _hasil;
	end if;

	if _tipe = 'ind' then
		_hasil = replace(_hasil,'.','|');
		_hasil = replace(_hasil,',','.');
		_hasil = replace(_hasil,'|',',');
	elsif _tipe = 'ina' then
		_hasil = replace(_hasil,'.','|');
		_hasil = replace(_hasil,',','.');
		_hasil = replace(_hasil,'|',',');
		if _nilai < 0 then
			_hasil = replace(_hasil,'-','');
			_hasil = concat('(',_hasil,')');
		else 
			_hasil = concat(_hasil,' ');
		end if;
	end if;

	return _hasil;
end
$$;

CREATE OR REPLACE FUNCTION fnsGetDescription(p_table_name text)
RETURNS TABLE (
    column_name text,
    column_comment text
) AS $$
BEGIN
    RETURN QUERY
    SELECT
        cols.column_name::text,   -- cast to text
        pgd.description::text     -- cast to text
    FROM
        pg_catalog.pg_statio_all_tables st
        INNER JOIN pg_catalog.pg_description pgd ON pgd.objoid = st.relid
        INNER JOIN information_schema.columns cols
            ON cols.table_schema = st.schemaname
           AND cols.table_name = st.relname
           AND cols.ordinal_position = pgd.objsubid
    WHERE
        cols.table_schema = current_schema()
        AND cols.table_name = p_table_name;
END;
$$ LANGUAGE plpgsql STABLE;

