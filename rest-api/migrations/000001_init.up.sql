CREATE TABLE "users" (
  "id" bigint,
  "name" text NOT NULL,
  "email" text NOT NULL,
  "password" text NOT NULL,
  "phone" varchar(15) NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "organizations" (
  "id" bigint,
  "name" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "organization_users" (
  "organization_id" bigint NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "customers" (
  "id" bigint,
  "name" text NOT NULL,
  "organization_id" bigint,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "channels" (
  "id" bigint,
  "name" text NOT NULL,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "organization_channels" (
  "id" bigint,
  "organization_id" bigint,
  "channel_id" bigint,
  "metadata" jsonb,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now()),
  PRIMARY KEY ("id")
);

CREATE TABLE "customer_channels" (
  "customer_id" bigint,
  "channel_id" bigint,
  "metadata" jsonb,
  "created_at" timestamp DEFAULT (now()),
  "updated_at" timestamp DEFAULT (now())
);

CREATE UNIQUE INDEX "email_unique" ON "users" ("email");
CREATE UNIQUE INDEX "phone_unique" ON "users" ("phone");
CREATE UNIQUE INDEX "organization_user_unique" ON "organization_users" ("organization_id", "user_id");
CREATE UNIQUE INDEX "customer_channel_unique" ON "customer_channels" ("customer_id", "channel_id");

ALTER TABLE "customer_channels" ADD FOREIGN KEY ("channel_id") REFERENCES "channels" ("id");
ALTER TABLE "customer_channels" ADD FOREIGN KEY ("customer_id") REFERENCES "customers" ("id");
ALTER TABLE "organization_channels" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON DELETE NO ACTION;
ALTER TABLE "organization_channels" ADD FOREIGN KEY ("channel_id") REFERENCES "channels" ("id") ON DELETE NO ACTION;
ALTER TABLE "organization_users" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id") ON DELETE NO ACTION;
ALTER TABLE "organization_users" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION;
ALTER TABLE "customers" ADD FOREIGN KEY ("organization_id") REFERENCES "organizations" ("id");



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