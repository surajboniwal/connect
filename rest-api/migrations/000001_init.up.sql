CREATE TABLE "users" (
  "id" bigint,
  "name" text NOT NULL,
  "email" text NOT NULL,
  "password" text NOT NULL,
  "phone" varchar(15) NOT NULL,
  "created_at" timestamp default CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);

CREATE TABLE "organizations" (
  "id" bigint,
  "name" text NOT NULL,
  "created_at" timestamp default CURRENT_TIMESTAMP,
  PRIMARY KEY ("id")
);

CREATE TABLE "organization_users" (
  "organization_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamp default CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX "email_unique" ON "users" ("email");
CREATE UNIQUE INDEX "phone_unique" ON "users" ("phone");
CREATE UNIQUE INDEX "organization_user_unique" ON "organization_users" ("organization_id", "user_id");

ALTER TABLE "organization_users" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON DELETE CASCADE;
ALTER TABLE "organization_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;