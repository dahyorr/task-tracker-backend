

CREATE  FUNCTION modify_updated_at_on_update()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';


CREATE TABlE tasks (
                id SERIAL PRIMARY KEY,
                name TEXT,
                description TEXT,
                status TEXT,
                created_by INTEGER REFERENCES users(id),
                workspace_id INTEGER REFERENCES workspaces(id),
                due_date TIMESTAMP,
                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
            );

CREATE TRIGGER update_task_updated_on
    BEFORE UPDATE
    ON
        tasks
    FOR EACH ROW
EXECUTE PROCEDURE modify_updated_at_on_update();