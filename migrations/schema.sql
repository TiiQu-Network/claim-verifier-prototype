CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT NOT NULL
);
CREATE UNIQUE INDEX "version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "recipients" (
"id" TEXT PRIMARY KEY,
"first_name" TEXT NOT NULL,
"last_name" TEXT NOT NULL,
"dob" DATETIME NOT NULL,
"address_id" integer NOT NULL,
"institution_id" integer NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "institutions" (
"id" TEXT PRIMARY KEY,
"name" text NOT NULL,
"salt" text NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "addresses" (
"id" TEXT PRIMARY KEY,
"first_line" text NOT NULL,
"second_line" text NOT NULL,
"town" text NOT NULL,
"county" text NOT NULL,
"postcode" text NOT NULL,
"country" text NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "data_hashes" (
"id" TEXT PRIMARY KEY,
"personal_data" text NOT NULL,
"personal_constituents" TEXT,
"certification_data" text NOT NULL,
"certification_constituents" TEXT,
"claimant_address" TEXT,
"institution_id" integer,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "issueables" (
"id" TEXT PRIMARY KEY,
"name" text NOT NULL,
"institution_reference" text NOT NULL,
"institution_id" integer,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "issues" (
"id" TEXT PRIMARY KEY,
"recipient_id" integer NOT NULL,
"recipient_ref" text NOT NULL,
"issueable_id" integer NOT NULL,
"issueable_level_id" integer NOT NULL,
"issue_date" NUMERIC NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "platform_members" (
"id" TEXT PRIMARY KEY,
"first_name" text NOT NULL,
"last_name" text NOT NULL,
"dob" DATETIME NOT NULL,
"blockchain_address" text NOT NULL,
"recipient_id" integer,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "platform_member_claims" (
"id" TEXT PRIMARY KEY,
"member_id" integer NOT NULL,
"address_id" integer NOT NULL,
"recipient_ref" TEXT NOT NULL,
"issue_id" integer NOT NULL,
"issue_level_id" integer NOT NULL,
"issue_date" DATETIME NOT NULL,
"verified" bool,
"contract_id" integer,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
