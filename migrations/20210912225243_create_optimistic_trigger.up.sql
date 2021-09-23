CREATE OR REPLACE FUNCTION trigger_update_optimistic()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated = NOW();
    NEW.version = NEW.version + 1;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
