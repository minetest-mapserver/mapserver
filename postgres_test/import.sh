#!/bin/sh
cat minetest.postgres.sql | sudo docker exec -i postgres_test_postgres_1 psql -U postgres
