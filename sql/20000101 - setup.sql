--------------------------------
-- begin drop all data on schema
--------------------------------
/*
DO $$ DECLARE
rec RECORD;
_schemaName varchar;
BEGIN
_schemaName = 'xxxxxxxxxx'; -- current_schema()

FOR rec IN (SELECT tablename FROM pg_tables WHERE schemaname = _schemaName) LOOP
EXECUTE 'DROP TABLE IF EXISTS ' || rec.tablename || ' CASCADE';
END LOOP;

FOR rec IN (
SELECT 'DROP FUNCTION ' || ns.nspname || '.' || proname 
       || '(' || oidvectortypes(proargtypes) || ');' as sintax
FROM pg_proc INNER JOIN pg_namespace ns ON (pg_proc.pronamespace = ns.oid)
WHERE prokind = 'f' and ns.nspname = _schemaName  order by proname
) LOOP
EXECUTE rec.sintax;
END LOOP;


FOR rec IN (
SELECT 'DROP PROCEDURE ' || ns.nspname || '.' || proname 
       || '(' || oidvectortypes(proargtypes) || ');' as sintax
FROM pg_proc INNER JOIN pg_namespace ns ON (pg_proc.pronamespace = ns.oid)
WHERE prokind = 'p' and ns.nspname = _schemaName  order by proname
) LOOP
EXECUTE rec.sintax;
END LOOP;

END $$;
*/

-----------------------------
-- begin create event trigger
-----------------------------
CREATE OR REPLACE FUNCTION public.stptblhss()
RETURNS event_trigger AS $$
DECLARE
    event_record RECORD;
BEGIN
--	raise notice 'masuk sini';

	if exists (
		SELECT * FROM pg_tables 
        WHERE schemaname = current_schema()
        AND tablename = 'tblhss'
	) then
		if exists (
	        select 
				* 
	        from pg_proc p
	        left join pg_namespace n on n.oid = p.pronamespace
	        where p.proname = 'stptblhss'
	          and n.nspname = 'public'
	    ) then
	
		    FOR event_record IN 
				SELECT * FROM pg_event_trigger_ddl_commands() 
				Where schema_name = current_schema()
					and command_tag not in ('COMMENT')
			LOOP
	--				raise notice '%', concat(event_record.object_type);
	
				insert into tblhss ( 
					tlevtp,tlschm,tlobtp,tlobnm,tlstmt,tlremk,
					tlrgid,tlrgdt,tlchid,tlchdt,tlchno,
					tldlfg,tlcsdt,tlcsid,tlsrce,tlunix
				) 
				values (
					event_record.command_tag,  
					event_record.schema_name,
					event_record.object_type,
					event_record.object_identity,
					current_query(),
					'',
					current_user,clock_timestamp(),current_user,clock_timestamp(),0,
					0,clock_timestamp(),current_user,'pgsql',''
				);
	
		--        IF event_record.command_tag IN ('CREATE PROCEDURE', 'ALTER PROCEDURE') THEN
		--            INSERT INTO procedure_changes (event_type, schema_name, procedure_name, user_name, query, changed_at)
		--            VALUES (
		--                event_record.command_tag,  
		--                event_record.schema_name,
		--                event_record.object_identity, 
		--                current_user,
		--                current_query(),
		--                NOW()
		--            );
		--        END IF;
		    END LOOP;
	    end if;
	end if;

END;
$$ LANGUAGE plpgsql;

DO $$
BEGIN
	IF not exists (
		SELECT *
	    FROM pg_event_trigger 
	    WHERE evtname = 'trstblhss'
	) then
		create EVENT TRIGGER trstblhss
		ON ddl_command_end
--		WHEN TAG IN ('CREATE PROCEDURE', 'ALTER PROCEDURE')
		EXECUTE FUNCTION public.stptblhss();
	end if;
END $$;

