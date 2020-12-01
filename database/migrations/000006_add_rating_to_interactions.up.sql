ALTER TABLE interactions
  ADD COLUMN rating smallint CHECK (rating >= 0 AND rating <= 10);