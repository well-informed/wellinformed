ALTER TABLE interactions
  ADD COLUMN completed BOOLEAN NOT NULL DEFAULT FALSE;

ALTER TABLE interactions
  ADD COLUMN saved_for_later BOOLEAN NOT NULL DEFAULT FALSE;