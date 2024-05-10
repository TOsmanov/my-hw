#!/bin/bash -x

# Copy hw15_go_sql
cp -r ../hw15_go_sql ./

# Run the database and server using Docker
docker compose up

# Remove hw15_go_sql
rm -rf ./hw15_go_sql