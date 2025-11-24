# Test Plan and Report Template

## Overview

This document describes the standard format for writing test plans and test reports for the project. All test documentation should follow this structure to ensure consistency and clarity.

## Test Plan Structure

### Test Case Table Format

Each feature should have a test case table with the following columns:

| Column | Description |
|--------|-------------|
| **ID** | Unique test case identifier (e.g., `MC-ATC-1`, `ATTR-CREATE-1`) |
| **Test Case Description** | Brief description of what is being tested |
| **Test Case Procedure** | Step-by-step instructions to execute the test |
| **Expected Output** | The expected result or behavior |
| **Inter-test Case Dependence** | Other test cases that must pass first (if any) |
| **Result** | Test execution status (Passed/Failed/Untested) |
| **Test Date** | Date when the test was executed |
| **Note** | Additional comments or observations |

### Test Case ID Naming Convention

- **Format**: `<FEATURE>-<ACTION>-<NUMBER>`
- **Examples**:
  - `ATTR-CREATE-1`: Attribute Creation test case #1
  - `ATTR-UPDATE-1`: Attribute Update test case #1
  - `ATTR-DELETE-1`: Attribute Delete test case #1
  - `ATTRVAL-CREATE-1`: Attribute Value Creation test case #1

### Test Case Procedure Format

Write test procedures as numbered steps:

```
1. Step one description
2. Step two description
3. Step three description
```

For API tests, include:
- HTTP method and endpoint
- Request body/parameters
- Authentication details (if required)

For UI tests, include:
- Navigation steps
- Button clicks
- Form inputs
- Verification points

### Expected Output Format

Describe expected results clearly:
- Status codes (for API tests)
- Response body structure (for API tests)
- UI changes (for UI tests)
- Database state changes
- Error messages (for negative tests)

## Example Test Plan

### Feature: Add To Cart

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| MC-ATC-1 | Verify user can add an item to the cart | 1. Navigate to any product details page.<br>2. Click "Add to Cart" button.<br>3. Observe cart icon or popup. | Display notification "Item added to cart."<br>Item appears in cart with quantity = 1.<br>Subtotal, tax, and total updated correctly. | | Untested | | |
| MC-ATC-2 | Verify adding the same item increases its quantity | 1. Navigate to the same product page.<br>2. Click "Add to Cart" again. | Item's quantity in cart increases by 1.<br>Subtotal, tax, and total updated correctly. | MC-ATC-1 | Untested | | |

## Test Execution Guidelines

### Before Testing
1. Review all test cases for the feature
2. Ensure test environment is set up correctly
3. Prepare test data if needed
4. Document environment details (version, configuration, etc.)

### During Testing
1. Follow test procedures exactly as written
2. Document any deviations from expected results
3. Take screenshots/logs for failures
4. Note any unexpected behavior

### After Testing
1. Update the Result column (Passed/Failed)
2. Fill in the Test Date
3. Add notes for any failures or observations
4. Create bug reports for failures
5. Update test cases if procedures need clarification

## Test Result Values

- **Passed**: Test executed successfully, output matches expected
- **Failed**: Test executed, but output differs from expected
- **Blocked**: Cannot execute due to dependency failure or environment issue
- **Untested**: Test has not been executed yet
- **N/A**: Test case not applicable for current release/version

## Best Practices

1. **Be Specific**: Write clear, unambiguous test procedures
2. **Be Complete**: Include all necessary steps and verification points
3. **Be Consistent**: Use the same format across all test documentation
4. **Be Traceable**: Link test cases to requirements or user stories
5. **Be Maintainable**: Update test cases when features change
6. **Be Realistic**: Expected outputs should match actual system behavior
7. **Document Dependencies**: Clearly note when tests depend on each other

## Automated Test Correlation

When possible, correlate manual test cases with automated tests:

```markdown
**Automated Test**: `TestAttribute_Create` in `internal/domain/attribute_test.go`
**Manual Test**: ATTR-CREATE-1
```

This helps ensure coverage and identifies gaps between manual and automated testing.
