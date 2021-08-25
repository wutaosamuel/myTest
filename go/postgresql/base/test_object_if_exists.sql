CREATE SCHEMA IF NOT EXISTS test_schema;

CREATE TABLE IF NOT EXISTS test_schema.test_object (
	"name" varchar(255) PRIMARY KEY,
	"id" int NOT NULL,
	"value" double precision DEFAULT 0
);