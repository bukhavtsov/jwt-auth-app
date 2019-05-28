FROM postgres:alpine
COPY script/init.sql /docker-entrypoint-initdb.d/db_init.sql
EXPOSE 5432