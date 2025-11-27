-- attribute_values_temp
CREATE TEMPORARY TABLE temp_attribute_values (
  id UUID PRIMARY KEY,
  attribute_id UUID NOT NULL,
  value TEXT NOT NULL,
  deleted_at TIMESTAMPTZ
);

-- cart_items_temp
CREATE TEMPORARY TABLE temp_cart_items (
  id UUID PRIMARY KEY,
  quantity INTEGER NOT NULL,
  cart_id UUID NOT NULL,
  product_variant_id UUID NOT NULL
);

-- product_variants_temp
CREATE TEMPORARY TABLE temp_product_variants (
  id UUID PRIMARY KEY,
  sku TEXT NOT NULL,
  price DECIMAL(12, 0) NOT NULL,
  quantity INTEGER NOT NULL,
  purchase_count INTEGER NOT NULL,
  product_id UUID NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ
);

-- product_images_temp
CREATE TEMPORARY TABLE temp_product_images (
  id UUID PRIMARY KEY,
  url TEXT NOT NULL,
  "order" INTEGER NOT NULL,
  product_id UUID NOT NULL,
  product_variant_id UUID,
  created_at TIMESTAMPTZ NOT NULL,
  deleted_at TIMESTAMPTZ
);

-- products_attribute_values_temp
CREATE TEMPORARY TABLE temp_products_attribute_values (
  product_id UUID NOT NULL,
  attribute_value_id UUID NOT NULL,
  PRIMARY KEY (product_id, attribute_value_id)
);

-- name: CreateTempTableOptions :exec
CREATE TEMPORARY TABLE temp_options (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  product_id UUID NOT NULL,
  deleted_at TIMESTAMPTZ
) ON COMMIT DROP;

-- option_values_temp
CREATE TEMPORARY TABLE temp_option_values (
  id UUID PRIMARY KEY,
  value TEXT NOT NULL,
  option_id UUID NOT NULL,
  deleted_at TIMESTAMPTZ
);

-- option_values_product_variants_temp
CREATE TEMPORARY TABLE temp_option_values_product_variants (
  product_variant_id UUID NOT NULL,
  option_value_id UUID NOT NULL,
  PRIMARY KEY (product_variant_id, option_value_id)
);

-- name: CreateTempTableOrderItems :exec
CREATE TEMPORARY TABLE temp_order_items (
  id UUID PRIMARY KEY,
  quantity INTEGER NOT NULL,
  order_id UUID NOT NULL,
  price NUMERIC NOT NULL,
  product_variant_id UUID NOT NULL
) ON COMMIT DROP;
