# Integration Testing Strategy

## Scope

- **Application:** Integration test, use testcontainers.
- **Focus:** Test full lifecycles (Create->Get->Update->Delete) and side effects (Cache, DB).
- There are some services that cannot be connected to real external systems (e.g., payment gateway). In that case, use mocking (mockery) for those external systems only.

## Best Practices

### 1. Testify Suite & Testcontainers

- Use `testify/suite` for integration tests requiring setup/teardown.
- Use testcontainers for external dependencies (Redis, S3, DB).
- Example structure:

  ```go
  type AttributeTestSuite struct {
      suite.Suite
      containers *component.Containers
      app        http.AttributeApplication
  }

  func (s *AttributeTestSuite) SetupSuite() {
      // Initialize containers and application
  }

  func (s *AttributeTestSuite) TearDownSuite() {
      s.containers.Cleanup(s.T())
  }
  ```

### 2. Lifecycle Testing

- Verify state changes across multiple operations.
- Check side effects like Cache hits/misses/invalidation.
- Example:
  ```go
  s.Run("Create resource", func() { ... })
  s.Run("Get resource (cache miss)", func() { ... })
  s.Run("Get resource (cache hit)", func() { ... })
  s.Run("Update resource (cache invalidation)", func() { ... })
  s.Run("Delete resource", func() { ... })
  ```
- Use `s.Require().NoError(err)` for critical assertions that should stop the test on failure.

### 3. Seeding

- In integration tests, use seeding setting to prepare initial data state.
  ```go
  func (s *OrderTestSuite) newContainersConfig() *component.ContainersConfig {
      containersConfig := component.NewContainersConfig(&component.NewContainersConfigParam{
          DBEnabled: true,
      })
      containersConfig.DB.Seed = true
      return containersConfig
  }
  ```

#### Seed Data Reference

The seed data includes the following records. There are much more records seeded, but there are few shown below here. **Do not read the `seed.sql` file directly as it is very large.**

- **Attributes:**

  ```sql
  INSERT INTO attributes (id, code, name) VALUES
    ('00000000-0000-7000-0000-000000000007', 'battery_capacity', 'Dung lượng pin'),
    ('00000000-0000-7000-0000-000000000010', 'brand', 'Thương hiệu'),
    ('00000000-0000-7000-0000-000000000057', 'is_warranty_applied', 'Sản phẩm có được bảo hành không?'),
  ...
  ```

- **Attribute Values:**

  ```sql
  INSERT INTO attribute_values (id, attribute_id, value) VALUES
    ('00000000-0000-7000-0000-000000000017', '00000000-0000-7000-0000-000000000007', '2500'),
    ('00000000-0000-7000-0000-000000000018', '00000000-0000-7000-0000-000000000007', '4400'),
    ('00000000-0000-7000-0000-000000000019', '00000000-0000-7000-0000-000000000007', '5000 mAh'),
    ('00000000-0000-7000-0000-000000000020', '00000000-0000-7000-0000-000000000007', '5000mAh'),
    ('00000000-0000-7000-0000-000000000021', '00000000-0000-7000-0000-000000000007', '5100mAh'),
    ...
  ```

- **Category:** `00000000-0000-7000-0000-000000001796`

- **Products:**

  ```sql
  INSERT INTO products (id, name, price, rating, description, category_id) VALUES
    ('00000000-0000-7000-0000-000278469304', 'Điện thoại Masstel Izi 56 4G (LTE) Gọi HD Call ,Pin khủng ,loa lớn - Hàng Chính Hãng', 499000, 0, ..., '00000000-0000-7000-0000-000000001796'),
    ('00000000-0000-7000-0000-000278469345', 'Điện thoại ZTE Family 4GB/128GB, Màn OLED Full HD+, Dimensity 700, Kháng nước IP67, Sạc 22,5W - Mới nguyên seal - Hàng nhập khẩu nhật', 2599000, 0, ..., '00000000-0000-7000-0000-000000001795'),
    ...
  ```

- **Products - Attribute Values:**

  ```sql
  INSERT INTO products_attribute_values (product_id, attribute_value_id) VALUES
    ('00000000-0000-7000-0000-000278469304', '00000000-0000-7000-0000-000000000017'),
    ('00000000-0000-7000-0000-000278469345', '00000000-0000-7000-0000-000000000018'),
    ...
  ```

- **Options:**

  ```sql
  INSERT INTO options (id, name, product_id) VALUES
    ('00000000-0000-7000-0000-000000000001', 'Màu Sắc', '00000000-0000-7000-0000-000278469304'),
    ('00000000-0000-7000-0000-000000000002', 'Màu sắc', '00000000-0000-7000-0000-000278469345')
  ```

- **Option Values:**

  ```sql
  INSERT INTO option_values (id, value, option_id) VALUES
    ('00000000-0000-7000-0000-000000000001', 'Black/Đen', '00000000-0000-7000-0000-000000000001'),
    ('00000000-0000-7000-0000-000000000002', 'Blue/Xanh', '00000000-0000-7000-0000-000000000001'),
    ('00000000-0000-7000-0000-000000000003', 'vàng', '00000000-0000-7000-0000-000000000001'),
    ('00000000-0000-7000-0000-000000000004', ' Trắng', '00000000-0000-7000-0000-000000000002'),
    ('00000000-0000-7000-0000-000000000005', ' Đen', '00000000-0000-7000-0000-000000000002')
  ```

- **Product Variants:**

  ```sql
  INSERT INTO product_variants (id, sku, price, quantity, purchase_count, product_id) VALUES
    ('00000000-0000-7000-0000-000278469308', '4868714459472', 499000, 1000, 100, '00000000-0000-7000-0000-000278469304'),
    ('00000000-0000-7000-0000-000278469306', '5829048963687', 499000, 1000, 100, '00000000-0000-7000-0000-000278469304'),
    ('00000000-0000-7000-0000-000278469310', '4493764476839', 499000, 1000, 100, '00000000-0000-7000-0000-000278469304'),
    ('00000000-0000-7000-0000-000278469347', '4432167913574', 2599000, 1000, 100, '00000000-0000-7000-0000-000278469345'),
    ('00000000-0000-7000-0000-000278469350', '7717071852400', 2599000, 1000, 100, '00000000-0000-7000-0000-000278469345'),
  ...
  ```

- **Option Values - Product Variants:**

  ```sql
  INSERT INTO option_values_product_variants (product_variant_id, option_value_id) VALUES
    ('00000000-0000-7000-0000-000278469308', '00000000-0000-7000-0000-000000000002'),
    ('00000000-0000-7000-0000-000278469306', '00000000-0000-7000-0000-000000000003'),
    ('00000000-0000-7000-0000-000278469310', '00000000-0000-7000-0000-000000000001'),
    ('00000000-0000-7000-0000-000278469347', '00000000-0000-7000-0000-000000000004'),
    ('00000000-0000-7000-0000-000278469350', '00000000-0000-7000-0000-000000000005')
  ```

- **Product Images:**

  ```sql
  INSERT INTO product_images (id, url, "order", product_id, product_variant_id) VALUES
    ('00000000-0000-7000-0000-000000000002', 'https://salt.tikicdn.com/ts/product/9a/e4/9f/9ca6f5af6469f9a0855ff256ba7130fd.jpg', 1, '00000000-0000-7000-0000-000278469304', NULL),
    ('00000000-0000-7000-0000-000000000003', 'https://salt.tikicdn.com/ts/product/b8/46/23/3ba27ae38e439003342d3598de1ab5db.jpg', 2, '00000000-0000-7000-0000-000278469304', NULL),
    ('00000000-0000-7000-0000-000000000004', 'https://salt.tikicdn.com/ts/product/7e/fa/d0/1b6d949a66c6e2b1a2788a008b86f42b.jpg', 3, '00000000-0000-7000-0000-000278469304', NULL),
    ('00000000-0000-7000-0000-000000000005', 'https://salt.tikicdn.com/cache/w1200/ts/product/11/f4/ff/6dec330d7aeb6250d7e4d8142a4b373e.jpg', 4, '00000000-0000-7000-0000-000278469304', '00000000-0000-7000-0000-000278469308'),
    ...
  ```

- **Cart:**

  ```sql
  INSERT INTO carts (id, user_id) VALUES
    ('00000000-0000-7000-0000-000000000001', '00000000-0000-7000-0000-000000000003') -- This is the only user that has cart, name is 'customer'
  ```

- **Cart Items:**
  ```sql
  INSERT INTO cart_items (id, cart_id, product_variant_id, quantity) VALUES
    ('00000000-0000-7000-0000-000000000001', '00000000-0000-7000-0000-000000000001', '00000000-0000-7000-0000-000278469308', 10),
    ('00000000-0000-7000-0000-000000000002', '00000000-0000-7000-0000-000000000001', '00000000-0000-7000-0000-000278620836', 1)
  ```

## Rules

1. **Application:** Integration tests with Testcontainers & Testify Suite.
2. **Lifecycle:** Test sequences (Create->Get->Update->Delete) and cache side-effects.
3. **Seeding:** Use `containersConfig.DB.Seed = true` for tests requiring pre-populated data.
4. **Assertions:** Use `s.Require()` for critical steps, `s.Assert()` (or `s.Equal`, etc.) for checks.
