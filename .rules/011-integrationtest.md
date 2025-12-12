# Integration Testing Strategy

## Scope

- **Application:** Integration test, use testcontainers.
- **Focus:** Test full lifecycles (Create->Get->Update->Delete) and side effects (Cache, DB).

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

The seed data includes:
- **Attributes:** `battery_capacity`, `brand`, `is_warranty_applied`, etc.
- **Attribute Values:** Specific values for the above attributes.
- **Category:** `00000000-0000-7000-0000-000000001796`
- **Products:**
  - `00000000-0000-7000-0000-000278469304` (Masstel Izi 56 4G)
  - `00000000-0000-7000-0000-000278469345` (ZTE Family)
- **Product Variants:**
  - `00000000-0000-7000-0000-000278469308` (Variant of Masstel)
  - `00000000-0000-7000-0000-000278469347` (Variant of ZTE)
- **User:** `00000000-0000-7000-0000-000000000003` (Customer with cart)
- **Cart:** `00000000-0000-7000-0000-000000000001`

*Note: Do not read the `seed.sql` file directly as it is very large.*

## Rules

1. **Application:** Integration tests with Testcontainers & Testify Suite.
2. **Lifecycle:** Test sequences (Create->Get->Update->Delete) and cache side-effects.
3. **Seeding:** Use `containersConfig.DB.Seed = true` for tests requiring pre-populated data.
4. **Assertions:** Use `s.Require()` for critical steps, `s.Assert()` (or `s.Equal`, etc.) for checks.
