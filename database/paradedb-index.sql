CREATE INDEX IF NOT EXISTS categories_search_idx ON categories
USING bm25 (id, name)
WITH (key_field = 'id');

CREATE INDEX IF NOT EXISTS products_search_idx ON products
USING bm25 (id, name)
WITH (key_field = 'id');
