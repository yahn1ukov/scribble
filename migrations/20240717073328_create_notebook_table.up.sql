CREATE TABLE IF NOT EXISTS notebooks (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "user_id" UUID NOT NULL REFERENCES users("id") ON DELETE CASCADE,
    "title" VARCHAR UNIQUE NOT NULL,
    "description" TEXT,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now(),
    "updated_at" TIMESTAMPTZ
);

CREATE TRIGGER notebooks_updated_at
BEFORE UPDATE ON notebooks
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
