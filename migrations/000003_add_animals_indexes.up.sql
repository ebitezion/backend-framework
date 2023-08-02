CREATE INDEX IF NOT EXISTS animals_title_idx ON movies USING GIN (to_tsvector('simple', title));

CREATE INDEX IF NOT EXISTS animals_genres_idx ON movies USING GIN (genres);