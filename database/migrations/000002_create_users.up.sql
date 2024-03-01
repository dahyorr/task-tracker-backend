CREATE TABLE users (
                 id SERIAL PRIMARY KEY,
                 username TEXT UNIQUE,
                 email TEXT NOT NULL UNIQUE,
                 password TEXT
             );


