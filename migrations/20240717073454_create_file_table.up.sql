CREATE TABLE IF NOT EXISTS files (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "name" VARCHAR NOT NULL,
    "size" INT NOT NULL,
    "content_type" VARCHAR NOT NULL,
    "url" VARCHAR NOT NULL,
    "note_id" UUID NOT NULL REFERENCES notes("id") ON DELETE CASCADE,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "updated_at" TIMESTAMPTZ
);

CREATE TRIGGER files_updated_at
BEFORE UPDATE ON files
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
