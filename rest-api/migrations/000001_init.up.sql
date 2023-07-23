CREATE TABLE "users" (
  "id" bigint,
  "name" text NOT NULL,
  "email" text NOT NULL,
  "password" text NOT NULL,
  "phone" varchar(15) NOT NULL,
  "created_at" timestamp default CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamp default CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "organizations" (
  "id" bigint,
  "name" text NOT NULL,
  "created_at" timestamp default CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamp default CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "organization_users" (
  "organization_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamp default CURRENT_TIMESTAMP NOT NULL,
  "updated_at" timestamp default CURRENT_TIMESTAMP NOT NULL
);

CREATE UNIQUE INDEX "email_unique" ON "users" ("email");
CREATE UNIQUE INDEX "phone_unique" ON "users" ("phone");
CREATE UNIQUE INDEX "organization_user_unique" ON "organization_users" ("organization_id", "user_id");

ALTER TABLE "organization_users" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON DELETE CASCADE;
ALTER TABLE "organization_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;





-- Trigger for updated_at field
CREATE OR REPLACE FUNCTION updated_timestamp_func()
RETURNS TRIGGER
LANGUAGE plpgsql AS
'
BEGIN
    NEW.updated_at = now();
    NEW.created_at = OLD.created_at;
    RETURN NEW;
END;
';

DO $$
DECLARE
    t text;
BEGIN
    FOR t IN
        SELECT table_name FROM information_schema.columns WHERE column_name = 'updated_at'
    LOOP
        EXECUTE format('CREATE TRIGGER trigger_update_timestamp
                    BEFORE UPDATE ON %I
                    FOR EACH ROW EXECUTE PROCEDURE updated_timestamp_func()', t,t);
    END loop;
END;
$$ language 'plpgsql';