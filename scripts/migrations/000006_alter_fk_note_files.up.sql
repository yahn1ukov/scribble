ALTER TABLE files ADD COLUMN note_id UUID NOT NULL REFERENCES notes(id) ON DELETE CASCADE;