CREATE TABLE IF NOT EXISTS notes (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "notebook_id" UUID NOT NULL REFERENCES notebooks("id") ON DELETE CASCADE,
    "title" VARCHAR NOT NULL,
    "content" TEXT,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "updated_at" TIMESTAMPTZ
);

CREATE TRIGGER notes_updated_at
BEFORE UPDATE ON notes
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
