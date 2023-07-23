-- Drop foreign keys
ALTER TABLE "organizations" DROP CONSTRAINT IF EXISTS "organizations_id_fkey";
ALTER TABLE "organization_users" DROP CONSTRAINT IF EXISTS "organization_users_organization_id_fkey";
ALTER TABLE "organization_users" DROP CONSTRAINT IF EXISTS "organization_users_user_id_fkey";
ALTER TABLE "organization_channels" DROP CONSTRAINT IF EXISTS "organization_channels_organization_id_fkey";
ALTER TABLE "organization_channels" DROP CONSTRAINT IF EXISTS "organization_channels_channel_id_fkey";
ALTER TABLE "customer_channels" DROP CONSTRAINT IF EXISTS "customer_channels_channel_id_fkey";
ALTER TABLE "customer_channels" DROP CONSTRAINT IF EXISTS "customer_channels_customer_id_fkey";

-- Drop triggers
DROP TRIGGER IF EXISTS trigger_update_timestamp ON "users";
DROP TRIGGER IF EXISTS trigger_update_timestamp ON "organizations";
DROP TRIGGER IF EXISTS trigger_update_timestamp ON "organization_users";
DROP TRIGGER IF EXISTS trigger_update_timestamp ON "customers";
DROP TRIGGER IF EXISTS trigger_update_timestamp ON "channels";
DROP TRIGGER IF EXISTS trigger_update_timestamp ON "organization_channels";
DROP TRIGGER IF EXISTS trigger_update_timestamp ON "customer_channels";

-- Drop indexes
DROP INDEX IF EXISTS "customer_channel_unique";
DROP INDEX IF EXISTS "organization_user_unique";
DROP INDEX IF EXISTS "phone_unique";
DROP INDEX IF EXISTS "email_unique";

-- Drop tables
DROP TABLE IF EXISTS "customer_channels";
DROP TABLE IF EXISTS "organization_channels";
DROP TABLE IF EXISTS "channels";
DROP TABLE IF EXISTS "customers";
DROP TABLE IF EXISTS "organization_users";
DROP TABLE IF EXISTS "organizations";
DROP TABLE IF EXISTS "users";
