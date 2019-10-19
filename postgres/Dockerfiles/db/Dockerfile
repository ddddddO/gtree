FROM postgres:12-alpine
ENV POSTGRES_DB foodb

COPY /db/init-dump.sql /docker-entrypoint-initdb.d/

EXPOSE 5432
