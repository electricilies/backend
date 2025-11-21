-- users
CREATE TABLE users (
  id UUID PRIMARY KEY
);

-- categories
CREATE TABLE categories (
  id UUID PRIMARY KEY,
  name TEXT UNIQUE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);

-- products
CREATE TABLE products (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  price DECIMAL(12, 0) NOT NULL,
  views_count INTEGER NOT NULL DEFAULT 0,
  total_purchase INTEGER NOT NULL DEFAULT 0,
  rating REAL NOT NULL DEFAULT 0,
  trending_score REAL NOT NULL DEFAULT 0,
  category_id UUID NOT NULL REFERENCES categories (id) ON UPDATE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);

-- attributes
CREATE TABLE attributes (
  id UUID PRIMARY KEY,
  code VARCHAR(100) UNIQUE NOT NULL,
  name TEXT NOT NULL,
  deleted_at TIMESTAMP
);

-- attribute_values
CREATE TABLE attribute_values (
  id UUID PRIMARY KEY,
  attribute_id UUID NOT NULL REFERENCES attributes (id) ON UPDATE CASCADE ON DELETE CASCADE,
  value TEXT NOT NULL
);

-- products_attribute_values
CREATE TABLE products_attribute_values (
  product_id UUID NOT NULL REFERENCES products (id) ON UPDATE CASCADE,
  attribute_value_id UUID NOT NULL REFERENCES attribute_values (id) ON UPDATE CASCADE,
  PRIMARY KEY (product_id, attribute_value_id)
);

-- product_variants
CREATE TABLE product_variants (
  id UUID PRIMARY KEY,
  sku TEXT UNIQUE NOT NULL,
  price DECIMAL(12, 0) NOT NULL,
  quantity INTEGER NOT NULL,
  purchase_count INTEGER NOT NULL DEFAULT 0,
  product_id UUID NOT NULL REFERENCES products (id) ON UPDATE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP
);

-- product_images
CREATE TABLE product_images (
  id UUID PRIMARY KEY,
  url TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  "order" INTEGER NOT NULL,
  product_id UUID REFERENCES products (id) ON UPDATE CASCADE,
  product_variant_id UUID REFERENCES product_variants (id) ON UPDATE CASCADE
);

-- reviews
CREATE TABLE reviews (
  id UUID PRIMARY KEY,
  rating SMALLINT NOT NULL CHECK (rating >= 1 AND rating <= 5),
  content TEXT,
  image_url TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  deleted_at TIMESTAMP,
  user_id UUID NOT NULL REFERENCES users (id) ON UPDATE CASCADE,
  order_item_id UUID NOT NULL REFERENCES order_items (id) ON UPDATE CASCADE
);

-- options
CREATE TABLE options (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  product_id UUID NOT NULL REFERENCES products (id) ON UPDATE CASCADE,
  deleted_at TIMESTAMP
);

-- option_values
CREATE TABLE option_values (
  id UUID PRIMARY KEY,
  value TEXT NOT NULL,
  option_id UUID NOT NULL REFERENCES options (id) ON UPDATE CASCADE ON DELETE CASCADE
);

-- option_values_product_variants
CREATE TABLE option_values_product_variants (
  product_variant_id UUID NOT NULL REFERENCES product_variants (id) ON UPDATE CASCADE,
  option_value_id UUID NOT NULL REFERENCES option_values (id) ON UPDATE CASCADE,
  PRIMARY KEY (product_variant_id, option_value_id)
);

-- carts
CREATE TABLE carts (
  id UUID PRIMARY KEY,
  user_id UUID UNIQUE NOT NULL REFERENCES users (id) ON UPDATE CASCADE,
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- cart_items
CREATE TABLE cart_items (
  id UUID PRIMARY KEY DEFAULT UUID_GENERATE_V4(),
  quantity INTEGER NOT NULL,
  cart_id UUID NOT NULL REFERENCES carts (id) ON UPDATE CASCADE,
  product_variant_id UUID NOT NULL REFERENCES product_variants (id) ON UPDATE CASCADE
);

-- order_statuses
CREATE TABLE order_statuses (
  id UUID PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);


-- order_providers
CREATE TABLE order_providers (
  id UUID PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);

-- orders
CREATE TABLE orders (
  id UUID PRIMARY KEY,
  address TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  total_amount DECIMAL(12, 0) NOT NULL,
  is_paid BOOLEAN NOT NULL DEFAULT FALSE,
  user_id UUID NOT NULL REFERENCES users (id) ON UPDATE CASCADE,
  status_id UUID NOT NULL REFERENCES order_statuses (id) ON UPDATE CASCADE,
  provider_id UUID NOT NULL REFERENCES order_providers (id) ON UPDATE CASCADE
);

-- order_items
CREATE TABLE order_items (
  id UUID PRIMARY KEY,
  quantity INTEGER NOT NULL,
  order_id UUID NOT NULL REFERENCES orders (id) ON UPDATE CASCADE,
  price DECIMAL(12, 0) NOT NULL,
  product_variant_id UUID NOT NULL REFERENCES product_variants (id) ON UPDATE CASCADE
);

-- return_request_statuses
CREATE TABLE return_request_statuses (
  id UUID PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);

-- return_requests
CREATE TABLE return_requests (
  id UUID PRIMARY KEY,
  reason VARCHAR(150) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  status_id UUID NOT NULL DEFAULT 1 REFERENCES return_request_statuses (id) ON UPDATE CASCADE,
  user_id UUID NOT NULL REFERENCES users (id) ON UPDATE CASCADE,
  order_item_id UUID NOT NULL REFERENCES order_items (id) ON UPDATE CASCADE
);

-- refund_statuses
CREATE TABLE refund_statuses (
  id UUID PRIMARY KEY,
  name TEXT UNIQUE NOT NULL
);

-- refunds
CREATE TABLE refunds (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  status_id UUID NOT NULL DEFAULT 1 REFERENCES refund_statuses (id) ON UPDATE CASCADE,
  order_item_id UUID NOT NULL REFERENCES order_items (id) ON UPDATE CASCADE,
  return_request_id UUID NOT NULL REFERENCES return_requests (id) ON UPDATE CASCADE
);
