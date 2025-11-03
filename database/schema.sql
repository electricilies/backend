-- users
CREATE TABLE users (
  id UUID PRIMARY KEY
);

-- categories
CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);

-- products
CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  views_count INTEGER NOT NULL DEFAULT 0,
  total_purchase INTEGER NOT NULL DEFAULT 0,
  trending_score FLOAT NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);

-- attributes
CREATE TABLE attributes (
  id SERIAL PRIMARY KEY,
  code VARCHAR(100) UNIQUE NOT NULL,
  name TEXT NOT NULL
);

-- attribute_values
CREATE TABLE attribute_values (
  id SERIAL PRIMARY KEY,
  attribute_id INTEGER NOT NULL REFERENCES attributes (id)
  ON UPDATE CASCADE,
  value TEXT NOT NULL
);

-- products_attribute_values
CREATE TABLE products_attribute_values (
  product_id INTEGER NOT NULL REFERENCES products (id) ON UPDATE CASCADE,
  attribute_value_id INTEGER NOT NULL REFERENCES attribute_values (id) ON UPDATE CASCADE,
  PRIMARY KEY (product_id, attribute_value_id)
);

-- product_variants
CREATE TABLE product_variants (
  id SERIAL PRIMARY KEY,
  sku TEXT UNIQUE NOT NULL,
  price DECIMAL NOT NULL,
  quantity INTEGER NOT NULL,
  purchase_count INTEGER NOT NULL DEFAULT 0,
  product_id INTEGER NOT NULL REFERENCES products (id) ON UPDATE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);

-- product_images
CREATE TABLE product_images (
  id SERIAL PRIMARY KEY,
  url TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  "order" INTEGER NOT NULL,
  product_variant_id INTEGER REFERENCES product_variants (id) ON UPDATE CASCADE
);

-- reviews
CREATE TABLE reviews (
  id SERIAL PRIMARY KEY,
  rating INTEGER NOT NULL CHECK (rate >= 1 AND rate <= 5),
  content TEXT,
  image_url TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,
  user_id UUID NOT NULL REFERENCES users (id) ON UPDATE CASCADE,
  product_id INTEGER NOT NULL REFERENCES products (id) ON UPDATE CASCADE
);

-- products_categories
CREATE TABLE products_categories (
  product_id INTEGER NOT NULL REFERENCES products (id) ON UPDATE CASCADE,
  category_id INTEGER NOT NULL REFERENCES categories (id) ON UPDATE CASCADE,
  PRIMARY KEY (product_id, category_id)
);

-- options
CREATE TABLE options (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  product_id INTEGER NOT NULL REFERENCES products (id) ON UPDATE CASCADE
);

-- option_values
CREATE TABLE option_values (
  id SERIAL PRIMARY KEY,
  value TEXT NOT NULL,
  option_id INTEGER NOT NULL REFERENCES options (id) ON UPDATE CASCADE
);

-- option_values_product_variants
CREATE TABLE option_values_product_variants (
  product_variant_id INTEGER NOT NULL REFERENCES product_variants (id) ON UPDATE CASCADE,
  option_value_id INTEGER NOT NULL REFERENCES option_values (id) ON UPDATE CASCADE,
  PRIMARY KEY (product_variant_id, option_value_id)
);

-- carts
CREATE TABLE carts (
  id SERIAL PRIMARY KEY,
  user_id UUID UNIQUE NOT NULL REFERENCES users (id) ON UPDATE CASCADE,
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_user_id
ON carts (user_id);

-- cart_items
CREATE TABLE cart_items (
  id SERIAL PRIMARY KEY,
  quantity INTEGER NOT NULL,
  cart_id INTEGER NOT NULL REFERENCES carts (id) ON UPDATE CASCADE,
  product_variant_id INTEGER NOT NULL REFERENCES product_variants (id) ON UPDATE CASCADE
);

-- order_statuses
CREATE TABLE order_statuses (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);

-- payment_methods
CREATE TABLE payment_methods (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);

-- payment_statuses
CREATE TABLE payment_statuses (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);

-- payment_providers
CREATE TABLE payment_providers (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);

-- payments
CREATE TABLE payments (
  id SERIAL PRIMARY KEY,
  amount DECIMAL NOT NULL,
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  method_id INTEGER NOT NULL REFERENCES payment_methods (id) ON UPDATE CASCADE,
  status_id INTEGER NOT NULL REFERENCES payment_statuses (id) ON UPDATE CASCADE,
  provider_id INTEGER NOT NULL REFERENCES payment_providers (id) ON UPDATE CASCADE
);

-- orders
CREATE TABLE orders (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  user_id UUID NOT NULL REFERENCES users (id) ON UPDATE CASCADE,
  status_id INTEGER NOT NULL REFERENCES order_statuses (id) ON UPDATE CASCADE,
  payment_id INTEGER NOT NULL REFERENCES payments (id) ON UPDATE CASCADE
);

-- order_items
CREATE TABLE order_items (
  id SERIAL PRIMARY KEY,
  quantity INTEGER NOT NULL,
  order_id INTEGER NOT NULL REFERENCES orders (id) ON UPDATE CASCADE,
  product_variant_id INTEGER NOT NULL REFERENCES product_variants (id) ON UPDATE CASCADE
);

-- return_request_statuses
CREATE TABLE return_request_statuses (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);

-- return_requests
CREATE TABLE return_requests (
  id SERIAL PRIMARY KEY,
  reason VARCHAR(150) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  status_id INTEGER NOT NULL REFERENCES return_request_statuses (id) ON UPDATE CASCADE,
  user_id UUID NOT NULL REFERENCES users (id) ON UPDATE CASCADE,
  order_item_id INTEGER NOT NULL REFERENCES order_items (id) ON UPDATE CASCADE
);

-- refund_statuses
CREATE TABLE refund_statuses (
  id SERIAL PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);

-- refunds
CREATE TABLE refunds (
  id SERIAL PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  status_id INTEGER NOT NULL REFERENCES refund_statuses (id) ON UPDATE CASCADE,
  payment_id INTEGER NOT NULL REFERENCES payments (id) ON UPDATE CASCADE,
  return_request_id INTEGER NOT NULL REFERENCES return_requests (id) ON UPDATE CASCADE
);
