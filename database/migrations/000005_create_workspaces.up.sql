CREATE TABLE workspaces (
                 id SERIAL PRIMARY KEY,
                 name TEXT UNIQUE,
                 owner_id INTEGER REFERENCES users(id),
                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
             );

CREATE TABLE workspace_members (
                 workspace_id INTEGER REFERENCES workspaces(id),
                 user_id INTEGER REFERENCES users(id),
                 PRIMARY KEY (workspace_id, user_id)
             );