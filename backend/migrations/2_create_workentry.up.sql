-- CreateTable
CREATE TABLE IF NOT EXISTS work_entries
(
    id          SERIAL PRIMARY KEY,
    start_time  varchar NOT NULL,
    end_time    varchar NOT NULL,
    break_time  varchar NOT NULL
);