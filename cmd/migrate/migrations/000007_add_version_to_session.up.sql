ALTER TABLE
    sessions
    ADD COLUMN version
        INT NOT NULL DEFAULT 0;