-- CreateTable
CREATE TABLE IF NOT EXISTS users (
         id SERIAL PRIMARY KEY,
         firstname TEXT NOT NULL,
         name TEXT NOT NULL,
         email TEXT NOT NULL
    );

-- Seed
INSERT INTO users (id, firstname, name, email) VALUES (1, 'Davide', 'Lo Cascio', 'locascio.davide@gmail.com');