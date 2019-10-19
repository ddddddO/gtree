FROM postgres:12-alpine
ENV POSTGRES_DB foodb

COPY /db/create-table.sql /docker-entrypoint-initdb.d/

EXPOSE 5432
