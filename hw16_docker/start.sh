#!/bin/bash -x

# Copy hw15_go_sql
cp -r ../hw15_go_sql ./

docker compose up

# Run the database using Docker
# docker compose -f ../hw15_go_sql/internal/storage/postgres/docker-compose.yml up -d

# docker run -it $(docker build -q .)
# docker compose up .

# Remove hw15_go_sql
rm -rf ./hw15_go_sql