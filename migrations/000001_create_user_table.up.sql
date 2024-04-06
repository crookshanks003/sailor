CREATE TYPE gender as ENUM('male', 'female');

CREATE TABLE "users" (
	"id" serial PRIMARY KEY,
	"name" varchar,
	"age" integer,
	"birthdate" date,
	"gender" gender,
	"location" varchar,
	"phone" varchar unique,
	"created_at" timestamp default current_timestamp,
	"password" varchar
);
