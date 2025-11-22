CREATE INDEX IF NOT EXISTS attributes_search_idx ON attributes
USING bm25 (id, code, name)
WITH (key_field = 'id');

CREATE INDEX IF NOT EXISTS attribute_values_search_idx ON attribute_values
USING bm25 (id, value)
WITH (key_field = 'id');

CREATE INDEX IF NOT EXISTS categories_search_idx ON categories
USING bm25 (id, name)
WITH (key_field = 'id');

CREATE INDEX IF NOT EXISTS products_search_idx ON products
USING bm25 (id, name)
WITH (key_field = 'id');
