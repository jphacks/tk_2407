-- Down Migration: Revert changes to stamps and reactions tables

BEGIN;

-- Step 1: Drop foreign key constraint on 'stamp_type'
ALTER TABLE reactions
    DROP CONSTRAINT reactions_stamp_type_fkey;

-- Step 2: Alter reactions table - replace 'stamp_type' with 'stamp_id'
ALTER TABLE reactions
    DROP COLUMN stamp_type,
    ADD COLUMN stamp_id CHAR(26) NOT NULL;

-- Step 3: Alter stamps table - re-add 'id' column and set as primary key, remove 'type' as primary key
ALTER TABLE stamps
    ADD COLUMN id CHAR(26) PRIMARY KEY,
    DROP CONSTRAINT stamps_pkey,
    ADD PRIMARY KEY (id);

-- Step 4: Re-add foreign key constraint on 'stamp_id' referencing 'stamps.id'
ALTER TABLE reactions
    ADD FOREIGN KEY (stamp_id) REFERENCES stamps (id) ON DELETE CASCADE;

COMMIT;
