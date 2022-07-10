FROM postgres:11

ADD ./init-db.sh /docker-entrypoint-initdb.d/
ADD ./seed/dump.sql /db-dumps/