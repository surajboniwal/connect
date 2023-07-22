CREATE TABLE "users" (
  "id" bigint,
  "name" text NOT NULL,
  "email" text NOT NULL,
  "password" text NOT NULL,
  "phone" varchar(15) NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "organizations" (
  "id" bigint,
  "name" text NOT NULL,
  PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "email_unique" ON "users" ("email");

CREATE UNIQUE INDEX "phone_unique" ON "users" ("phone");