# Business Logic Rules

## Soft Delete Strategy

**Soft deleted:** Product, ProductVariant, Category, Attribute, AttributeValue, Option, OptionValue, Review, Order
**Hard deleted:** CartItem (transient, no historical value)

## Aggregates

1. **Attribute** - Filterable properties (Color: Red, Material: Cotton)
2. **Cart** - Shopping cart (one per user)
3. **Category** - Product categorization
4. **Order** - Purchases, order items
5. **Product** - Products, variants, options, images
6. **Review** - Product ratings

## Product Structure

```
Product
├─ belongs to 1 Category (required)
├─ has multiple Attributes (for filtering)
├─ has multiple Options (e.g. "Size", "Color")
│  └─ each Option has OptionValues (e.g. "S/M/L", "Red/Blue")
├─ has multiple ProductVariants
│  ├─ combination of OptionValues
│  ├─ own price, quantity, SKU
│  └─ own images (optional)
└─ has ProductImages (fallback)
```

## Key Rules

### Product-Category

- Every product MUST have exactly one category
- Category cannot be null

### Attributes vs Options

- **Attributes:** For filtering/search (Material, Brand, Season), can be reused across products
- **Options:** Create variants (Size, Color)

### Variants

- A simple product must have exact one variant
- A configurable product must have at least one variant
- Variants are combinations of option values
- Each variant has unique SKU
- Own price, quantity, images
- Partial matrix allowed (not all combinations required)
- Can be soft-deleted independently

**Example:**

```
Product: T-Shirt
Options:
  Size: [S, M, L]
  Color: [Red, Blue]

Variants (partial matrix):
  ✅ S + Red
  ✅ M + Blue
  ✅ L + Red
  ❌ S + Blue (not in stock, not created)
```

### Product Images

- A product must have at least one image
- Variant isn't required to have images

### Life cycle

- When a product is created, it is not allowed to add more options or option values later
- Only product variants can be created based on existing options and option values
- Product variant can be created, updated, soft-deleted independently

### Product Metrics

- `views_count` - Incremented on product page view
- `total_purchase` - Incremented on order confirmation
- `trending_score` - Based on recent views/purchases
- `rating` - Average of approved reviews

## Cart Rules

### One Cart Per User

- ✅ Each user has exactly ONE active cart
- ✅ User has to create his cart himself due to not using automatic creation (e.g. on login)
- ✅ Persists across sessions

### Cart Items

- ✅ References specific ProductVariant (not just Product)
- ✅ Has quantity
- ✅ Stores price snapshot
- ✅ **Hard deleted** when removed
- ❌ Cannot add deleted/out-of-stock variants
- ❌ Quantity cannot exceed available stock

## Order Workflow

### Order Status Flow

```
pending → confirmed → shipped → delivered
    ↓
cancelled
```

- The pending is only applied to orders that has provider payment online integration (VNPAY, MOMO, etc..), otherwise the order is created as confirmed directly (COD)
- The COD provider payment method is Cash on Delivery, so the order is created as confirmed directly

### Order Confirmation (⚠️ Critical Transaction)

**When status: pending → confirmed:**

1. **Decrease Inventory**
   - Decrease `product_variants.quantity`
   - Validate sufficient stock BEFORE confirming
   - Reject if insufficient

2. **Clean Cart**
   - No cleaning cart items associated with order user

3. **Update Metrics**
   - Increment `products.total_purchase`
   - Increment `product_variants.purchase_count`

⚠️ **MUST be atomic transaction** (all or nothing)

### Order Cancellation

**When order cancelled:**

1. **Restore Inventory**
   - Increase `product_variants.quantity`
   - Only if previously confirmed/shipped

2. **Update Metrics**
   - Decrement `products.total_purchase`
   - Decrement `product_variants.purchase_count`

3. **Process Refund**

### Order Item Snapshots

- ✅ Store product name, variant SKU, price at purchase time
- ✅ Immutable - preserve even if product changes/deleted

## Review Rules

- ✅ One review per user per product
- ✅ Rating: 1-5 stars (integer)
- ✅ User must have purchased product
- ✅ Reviews can be soft-deleted (moderation)
- ✅ `products.rating` = AVG of non-deleted reviews

## Inventory Management

### Stock Tracking

- ✅ Tracked at ProductVariant level (not Product)
- ✅ Decreases on order confirmation
- ✅ Increases on order cancellation

### Out of Stock

- ✅ Variant with `quantity = 0` is out of stock
- ❌ Cannot add to cart
- ✅ Show "Out of Stock" badge
- ✅ Product page still accessible

## Validation Summary

### Product

- ✅ Must have category
- ✅ Name: 3-200 chars
- ✅ Description: min 10 chars
- ✅ Price > 0
- ✅ At least 1 image (product or variant)

### ProductVariant

- ✅ Unique SKU
- ✅ Price > 0
- ✅ Quantity >= 0
- ✅ At least 1 option value

### Cart

- ✅ Item quantity > 0
- ✅ Quantity ≤ available stock
- ❌ No deleted/OOS variants

### Order

- ✅ At least 1 item
- ✅ Total = sum of item prices
- ❌ Cannot confirm if insufficient inventory
- ❌ Cannot cancel if shipped/delivered

## Event-Driven Side Effects

**Order Confirmed:**

1. Decrease variant quantities ⚠️
2. Remove cart items ⚠️
3. Update metrics
4. Check restock threshold

**Order Cancelled:**

1. Restore variant quantities
2. Update metrics
3. Process refund

**Review Created/Updated:**

1. Recalculate product rating
2. Update trending score

**Product Viewed:**

1. Increment views_count
2. Update trending score

## Concurrency

**Problem:** Multiple users order last item
**Solution:** Use `SELECT ... FOR UPDATE` with transactions
**Pattern:** Row-level locking during order confirmation

## Testing Critical Paths

2. ✅ Handles insufficient inventory gracefully
3. ✅ Transaction rollback on failure
4. ✅ Variant partial matrix creation
5. ✅ One cart per user
6. ✅ Hard delete cart items
7. ✅ Soft delete behavior

## Quick Rules

1. ✅ Soft delete by default (except CartItem)
2. ✅ One cart per user, persists across sessions
3. ✅ Product has 1 category, multiple attributes
4. ✅ Variants = partial matrix of option values
5. ✅ Order confirmation = atomic (inventory - cart + metrics)
6. ✅ Order items = immutable snapshots
7. ✅ Reviews update product rating
8. ✅ Inventory at variant level
9. ✅ All critical ops use transactions
10. ✅ Concurrency: row-level locking
