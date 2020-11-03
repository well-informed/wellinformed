ALTER TABLE content_items
  ADD COLUMN source_type varchar DEFAULT 'SrcRSSFeed' NOT NULL;