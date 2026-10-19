# syntax=docker/dockerfile:1

FROM liquibase/liquibase:5.0.4

ADD --checksum=sha256:73914527305a40cce504b0d3d90b23caf565136912d607ae8ae7c5895512332c \
    --chown=liquibase:liquibase \
    https://repo.maven.apache.org/maven2/org/postgresql/postgresql/42.7.14/postgresql-42.7.14.jar \
    /liquibase/lib/postgresql-42.7.14.jar
