CREATE INDEX search_idx ON products
USING bm25 (id, name)
WITH (key_field = 'id');
