DO $$
BEGIN
    IF EXISTS (
        SELECT 1 FROM pg_class c
        JOIN pg_namespace n ON n.oid = c.relnamespace
        WHERE c.relname = 'idx_short_url' AND n.nspname = 'public'
    ) THEN
        DROP INDEX idx_short_url;
    END IF;
END $$;

DROP TABLE IF EXISTS urls;