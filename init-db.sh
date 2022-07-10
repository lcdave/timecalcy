#!/bin/bash

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER admin;
    CREATE DATABASE timecalcy-db;
    GRANT ALL PRIVILEGES ON DATABASE newdb TO newuser;
EOSQL

psql -f /seed/dump.sql timecalcy-db