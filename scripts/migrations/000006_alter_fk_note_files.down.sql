ALTER TABLE files DROP CONSTRAINT IF EXISTS fk_files_note;
ALTER TABLE files DROP COLUMN IF EXISTS note_id;
