-- ChatGPT did it for me, just because we use atlas free bruh?

DO $$
DECLARE
    trig RECORD;
    func RECORD;
BEGIN
    FOR trig IN
        SELECT event_object_table AS table_name, trigger_name
        FROM information_schema.triggers
        WHERE trigger_name LIKE 'ele_%'
    LOOP
        EXECUTE format('DROP TRIGGER IF EXISTS %I ON %I CASCADE;', trig.trigger_name, trig.table_name);
        RAISE NOTICE 'Dropped trigger: % on table: %', trig.trigger_name, trig.table_name;
    END LOOP;

    FOR func IN
        SELECT
            p.oid,
            n.nspname AS schema_name,
            p.proname AS function_name,
            pg_get_function_identity_arguments(p.oid) AS args
        FROM pg_proc p
        JOIN pg_namespace n ON n.oid = p.pronamespace
        WHERE n.nspname = 'public'
          AND p.proname LIKE 'ele_%'
    LOOP
        EXECUTE format(
            'DROP FUNCTION IF EXISTS %I.%I(%s) CASCADE;',
            func.schema_name,
            func.function_name,
            func.args
        );
        RAISE NOTICE 'Dropped function: %.%(%s)', func.schema_name, func.function_name, func.args;
    END LOOP;
END $$;
