# Business Logic Rules

This document describes the core business rules and domain logic for the e-commerce system.

## Core Principles

### Soft Delete Strategy

**Default Behavior:** Most entities use soft delete (marked with `deleted_at` timestamp)

**Soft-deleted entities:**

- ✅ Product
- ✅ ProductVariant
- ✅ Category
- ✅ Attribute
- ✅ AttributeValue
- ✅ Option
- ✅ OptionValue
- ✅ Review
- ✅ Order

**Hard-deleted entities:**

- ❌ **CartItem** - Deleted entirely from database when removed

**Why hard delete CartItem?** Cart items are transient data that don't need historical tracking. Once removed, they have no business value.

---

## Aggregates

The system is organized into the following aggregates (bounded contexts):

1. **Attribute** - Product attributes and their values (e.g., Color: Red, Blue)
2. **Cart** - Shopping cart management
3. **Category** - Product categorization hierarchy
4. **Order** - Purchase orders and order items
5. **Product** - Products, variants, options, and images
6. **Review** - Product reviews and ratings

---

## Product Aggregate

### Product Structure

```
Product
├── belongs to 1 Category (required)
├── has multiple Attributes (optional)
├── has multiple Options (e.g., "Size", "Color")
│   └── each Option has multiple OptionValues (e.g., "Small", "Red")
├── has multiple ProductVariants
│   ├── each variant has a combination of OptionValues
│   ├── each variant has its own price, quantity, SKU
│   └── each variant can have its own images
└── has multiple ProductImages (at product level)
```

### Business Rules

#### Product-Category Relationship

- **Rule:** Every product MUST belong to exactly one category
- **Constraint:** `category_id` is required and must reference an active (non-deleted) category
- **Behavior:** If category is deleted, products remain but queries should handle gracefully

#### Product Attributes

- **Rule:** A product can have multiple attribute values (e.g., Material: Cotton, Season: Summer)
- **Purpose:** Attributes are for filtering/faceted search, not for variants
- **Example:** A shirt has Material=Cotton (attribute), but Size=M/L/XL (options → variants)

#### Options and Variants

- **Rule:** Each Option (e.g., "Size") has multiple OptionValues (e.g., "Small", "Medium", "Large")
- **Rule:** ProductVariants are combinations of OptionValues from different options
- **Important:** Variants do NOT need to cover all possible combinations (partial matrix allowed)

**Example:**

```
Product: T-Shirt
Options:
  - Size: [Small, Medium, Large]
  - Color: [Red, Blue, Black]

Possible full matrix: 3 × 3 = 9 variants
Actual variants (partial matrix allowed):
  ✅ Small + Red
  ✅ Small + Blue
  ✅ Medium + Red
  ✅ Large + Black
  ❌ Small + Black (out of stock, not created)
  ❌ Medium + Blue (out of stock, not created)
  ❌ Medium + Black (out of stock, not created)
  ❌ Large + Red (out of stock, not created)
  ❌ Large + Blue (out of stock, not created)
```

#### Variant Management

- **Rule:** Each variant MUST have a unique SKU
- **Rule:** Each variant has its own price (can differ from base product price)
- **Rule:** Each variant has its own quantity (inventory tracking)
- **Rule:** Each variant can have specific images (product images are fallback)
- **Rule:** Variants can be soft-deleted independently

#### Product Images

- **Product-level images:** Shown when no variant is selected
- **Variant-level images:** Override product images when variant is selected
- **Rule:** Products should have at least 1 image at product or variant level

#### Product Metrics

- **views_count:** Incremented when product detail page is viewed
- **total_purchase:** Incremented when an order with this product is confirmed
- **trending_score:** Calculated based on recent views and purchases (algorithm TBD)
- **rating:** Average of all approved reviews for this product

---

## Cart Aggregate

### Cart Structure

```
User
└── has 1 Cart (singleton per user)
    └── has multiple CartItems
        ├── references ProductVariant
        ├── has quantity
        └── snapshot of price at time of addition
```

### Business Rules

#### One Cart Per User

- **Rule:** Each user has exactly ONE active cart
- **Rule:** Cart is created automatically on first item addition
- **Rule:** Cart persists across sessions (saved in database)
- **Behavior:** Guest users may have session-based carts (not persisted)

#### Cart Items

- **Rule:** Each CartItem references a specific ProductVariant (not just Product)
- **Rule:** CartItem stores quantity (user can have multiple of same variant)
- **Rule:** CartItem may store price snapshot (to detect price changes)
- **Rule:** CartItems are **hard deleted** when removed from cart

#### Cart Item Validation

- **Rule:** Cannot add deleted ProductVariant to cart
- **Rule:** Cannot add out-of-stock variant (quantity = 0) to cart
- **Rule:** Cart quantity cannot exceed variant available quantity
- **Behavior:** When variant quantity changes, cart should validate and notify user

#### Price Changes

- **Rule:** If product/variant price changes after adding to cart, show notification
- **Behavior:** User should be informed before checkout if prices changed

---

## Order Aggregate

### Order Structure

```
Order
├── belongs to 1 User
├── has status (pending, confirmed, shipped, delivered, cancelled)
├── has multiple OrderItems
│   ├── references ProductVariant
│   ├── snapshot of product/variant data at time of order
│   ├── quantity ordered
│   └── price at time of order
└── has total amount, shipping address, payment info
```

### Order Status Flow

```
pending → confirmed → shipped → delivered
    ↓
cancelled (can cancel from pending or confirmed)
```

### Business Rules

#### Order Creation

- **Rule:** Order is created with status = "pending" when user initiates checkout
- **Rule:** OrderItems are snapshots (copy product/variant data, don't just reference)
- **Why snapshot?** Product data may change; order should reflect purchase-time state

#### Order Confirmation (Critical Transaction)

**When order status changes from "pending" → "confirmed":**

1. **Decrease Inventory:**
   - **Rule:** Decrease `product_variants.quantity` for each ordered variant
   - **Validation:** Ensure sufficient quantity available (check before confirming)
   - **Behavior:** If insufficient quantity, reject confirmation and notify user

2. **Update Cart:**
   - **Rule:** Remove corresponding CartItems from user's cart
   - **Why?** Items are now ordered; cart should be cleaned up
   - **Behavior:** Only remove items that match ordered variants and quantities

3. **Update Product Metrics:**
   - **Rule:** Increment `products.total_purchase` for each product in order
   - **Rule:** Increment `product_variants.purchase_count` for each variant

**Transaction Requirement:** All inventory decreases and cart updates MUST happen atomically (database transaction)

#### Order Cancellation

**When order is cancelled:**

1. **Restore Inventory:**
   - **Rule:** Increase `product_variants.quantity` back for each cancelled item
   - **Only if:** Order was previously "confirmed" or "shipped" (not "pending")

2. **Update Metrics:**
   - **Rule:** Decrement `products.total_purchase`
   - **Rule:** Decrement `product_variants.purchase_count`

3. **Refund Handling:**
   - Process refund according to payment method
   - Update order with refund status and amount

#### Order Item Snapshots

- **Rule:** OrderItem stores product name, variant SKU, price at time of order
- **Why?** Product may be deleted or changed later; order should preserve history
- **Fields to snapshot:**
  - Product name, description, images
  - Variant SKU, option values (e.g., "Size: M, Color: Red")
  - Price at time of order
  - Any discounts applied

---

## Category Aggregate

### Business Rules

- **Rule:** Categories can be hierarchical (parent-child relationships)
- **Rule:** Soft delete - deleted categories don't cascade delete products
- **Rule:** Products in deleted categories should still be accessible by direct link
- **Behavior:** Queries should filter deleted categories from navigation/listings

---

## Attribute Aggregate

### Attribute Structure

```
Attribute (e.g., "Color", "Material", "Brand")
└── has multiple AttributeValues (e.g., "Red", "Blue", "Cotton")
```

### Business Rules

- **Rule:** Attributes are used for product filtering/faceted search
- **Rule:** Attributes are NOT used for creating variants (use Options for that)
- **Distinction:**
  - **Attribute:** Characteristics for filtering (Material, Brand, Season)
  - **Option:** Characteristics that create variants (Size, Color for variants)

**Example:**

```
Product: Winter Jacket
Attributes:
  - Brand: Nike
  - Material: Polyester
  - Season: Winter
Options (create variants):
  - Size: S, M, L, XL
  - Color: Black, Navy
```

---

## Review Aggregate

### Business Rules

- **Rule:** Each review is for a specific Product (not ProductVariant)
- **Rule:** User can only review a product they have purchased
- **Rule:** User can only submit one review per product
- **Rule:** Reviews have rating (1-5 stars) and optional text comment
- **Rule:** Reviews can be soft-deleted (moderation)

#### Product Rating Calculation

- **Rule:** `products.rating` is the average of all approved (non-deleted) reviews
- **Trigger:** Update product rating when review is created, updated, or deleted
- **Formula:** `AVG(reviews.rating) WHERE product_id = X AND deleted_at IS NULL`

---

## Inventory Management

### Stock Tracking

- **Rule:** Inventory is tracked at ProductVariant level (not Product level)
- **Rule:** `product_variants.quantity` represents available stock
- **Rule:** Stock decreases on order confirmation
- **Rule:** Stock increases on order cancellation

### Out of Stock Handling

- **Rule:** Variants with `quantity = 0` are considered out of stock
- **Behavior:** Out-of-stock variants:
  - Cannot be added to cart
  - Show "Out of Stock" badge in product listings
  - Can still be viewed (product page remains accessible)

### Low Stock Alerts

- **Optional Rule:** Notify admin when `quantity < threshold` (e.g., quantity < 5)
- **Behavior:** Trigger restock workflow

---

## Validation Rules Summary

### Product

- ✅ Must have category
- ✅ Name: 3-200 characters
- ✅ Description: minimum 10 characters
- ✅ Price: must be positive
- ✅ At least 1 image (product or variant level)

### ProductVariant

- ✅ Unique SKU required
- ✅ Price must be positive
- ✅ Quantity must be >= 0
- ✅ Must have at least 1 option value

### Cart

- ✅ CartItem quantity must be > 0
- ✅ CartItem quantity cannot exceed variant available quantity
- ✅ Cannot add deleted or out-of-stock variants

### Order

- ✅ Must have at least 1 order item
- ✅ Total amount must match sum of item prices
- ✅ Cannot confirm order if insufficient inventory
- ✅ Cannot cancel order if already shipped/delivered

### Review

- ✅ Rating: 1-5 stars (integer)
- ✅ User must have purchased the product
- ✅ One review per user per product

---

## Event-Driven Side Effects

Certain operations trigger side effects that should be handled:

### When Order Confirmed

1. Decrease variant quantities ⚠️ **Critical**
2. Remove items from cart ⚠️ **Critical**
3. Update product metrics (total_purchase)
4. Send order confirmation email
5. Trigger inventory restock check

### When Order Cancelled

1. Restore variant quantities
2. Update product metrics
3. Process refund
4. Send cancellation email

### When Review Created/Updated

1. Recalculate product rating
2. Update trending score

### When Product Viewed

1. Increment views_count
2. Update trending score

---

## Business Constraints

### Concurrency Handling

- **Problem:** Multiple users order last available item simultaneously
- **Solution:** Use database transactions with row-level locking
- **Implementation:** `SELECT ... FOR UPDATE` when confirming order

### Data Consistency

- **Rule:** Order confirmation MUST be atomic (all-or-nothing)
- **Rule:** Inventory decrease and cart cleanup happen in same transaction
- **Rule:** Use database constraints to enforce business rules where possible

### Audit Trail

- **Rule:** Soft-deleted entities retain history
- **Rule:** Order items are immutable snapshots
- **Rule:** Track order status changes with timestamps

---

## Testing Requirements

### Critical Paths to Test

1. **Order Confirmation Flow:**
   - ✅ Inventory decreases correctly
   - ✅ Cart items are removed
   - ✅ Handles insufficient inventory gracefully
   - ✅ Transaction rollback on failure

2. **Variant Creation:**
   - ✅ Partial matrix is allowed
   - ✅ Duplicate SKU is rejected
   - ✅ Price and quantity validations

3. **Cart Management:**
   - ✅ One cart per user
   - ✅ Cannot exceed available quantity
   - ✅ Hard delete cart items

4. **Soft Delete Behavior:**
   - ✅ Deleted entities don't appear in listings
   - ✅ Deleted entities remain accessible by ID (for orders/history)
   - ✅ Can restore soft-deleted entities

---

## Rules Summary

1. ✅ Soft delete by default (except CartItem)
2. ✅ One cart per user
3. ✅ Product has 1 category, multiple attributes
4. ✅ Variants are partial matrix of option values
5. ✅ Order confirmation decreases inventory atomically
6. ✅ Order confirmation removes cart items
7. ✅ Order items are immutable snapshots
8. ✅ Reviews update product rating
9. ✅ Inventory tracked at variant level
10. ✅ All critical operations use transactions
