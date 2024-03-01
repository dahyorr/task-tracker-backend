CREATE TABLE sessions (
                 id SERIAL PRIMARY KEY,
                 user_id INTEGER REFERENCES users(id),
                 token TEXT UNIQUE,
                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                 expires_at TIMESTAMP
             );
