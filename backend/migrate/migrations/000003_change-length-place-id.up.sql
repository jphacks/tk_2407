-- Up Migration: Apply changes to stamps and reactions tables

BEGIN;

-- Step 1: Drop existing foreign key in reactions table
ALTER TABLE reactions
    DROP CONSTRAINT reactions_stamp_id_fkey;

-- Step 2: Alter stamps table - drop 'id' column, set 'type' as primary key
ALTER TABLE stamps
    DROP COLUMN id,
    ADD PRIMARY KEY (type);

-- Step 3: Alter reactions table - replace 'stamp_id' column with 'stamp_type'
ALTER TABLE reactions
    DROP COLUMN stamp_id,
    ADD COLUMN stamp_type VARCHAR(255) NOT NULL;

-- Step 4: Add foreign key constraint on 'stamp_type' referencing 'stamps.type'
ALTER TABLE reactions
    ADD FOREIGN KEY (stamp_type) REFERENCES stamps (type) ON DELETE CASCADE;

COMMIT;
