drop table if exists "test_object";

CREATE TABLE "test_object" (
	"name" varchar(255) PRIMARY KEY,
	"id" int NOT NULL,
	"value" double precision DEFAULT 0
);