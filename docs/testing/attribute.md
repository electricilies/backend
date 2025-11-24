# Attribute Module Test Documentation

## Overview

This document contains comprehensive test cases for the Attribute module, covering domain logic, service layer, application layer, and repository interactions.

## Test Coverage Summary

- **Domain Tests**: `internal/domain/attribute_test.go`
- **Domain Validator Tests**: `internal/domain/attributevalidator_test.go`
- **Service Tests**: `internal/service/attribute_test.go`
- **Repository Tests**: `internal/infrastructure/repositorypostgres/attribute_test.go`

---

## Domain Layer Tests

### Attribute Creation

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-DOM-CREATE-1 | Verify attribute creation with valid parameters | 1. Call `NewAttribute("color", "Color")`<br>2. Verify returned attribute | Attribute created successfully<br>ID is not nil<br>Code = "color"<br>Name = "Color"<br>Values is empty slice<br>DeletedAt is nil | | Passed | 2025-11-25 | Automated: `TestNewAttribute` |

### Attribute Value Creation

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTRVAL-DOM-CREATE-1 | Verify attribute value creation with valid parameter | 1. Call `NewAttributeValue("Red")`<br>2. Verify returned value | AttributeValue created successfully<br>ID is not nil<br>Value = "Red"<br>DeletedAt is nil | | Passed | 2025-11-25 | Automated: `TestNewAttributeValue` |

### Attribute Update

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-DOM-UPDATE-1 | Verify attribute name update when name provided | 1. Create attribute with name "Color"<br>2. Call `Update(&"Updated Color")`<br>3. Verify name changed | Name updated to "Updated Color" | ATTR-DOM-CREATE-1 | Passed | 2025-11-25 | Automated: `TestAttribute_Update` |
| ATTR-DOM-UPDATE-2 | Verify attribute name unchanged when nil provided | 1. Create attribute with name "Color"<br>2. Call `Update(nil)`<br>3. Verify name unchanged | Name remains "Color" | ATTR-DOM-CREATE-1 | Passed | 2025-11-25 | Automated: `TestAttribute_Update` |

### Get Value By ID

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-DOM-GETVAL-1 | Verify retrieving existing attribute value by ID | 1. Create attribute with values "Red" and "Blue"<br>2. Call `GetValueByID(redValueID)`<br>3. Verify returned value | Returns AttributeValue with ID matching redValueID<br>Value = "Red" | ATTR-DOM-CREATE-1, ATTRVAL-DOM-CREATE-1 | Passed | 2025-11-25 | Automated: `TestAttribute_GetValueByID` |
| ATTR-DOM-GETVAL-2 | Verify nil returned when value ID not found | 1. Create attribute with one value<br>2. Call `GetValueByID(randomID)`<br>3. Verify nil returned | Returns nil | ATTR-DOM-CREATE-1 | Passed | 2025-11-25 | Automated: `TestAttribute_GetValueByID` |
| ATTR-DOM-GETVAL-3 | Verify nil returned when values slice is nil | 1. Create attribute with Values = nil<br>2. Call `GetValueByID(anyID)`<br>3. Verify nil returned | Returns nil | | Passed | 2025-11-25 | Automated: `TestAttribute_GetValueByID` |

### Add Values

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-DOM-ADDVAL-1 | Verify adding single value to attribute | 1. Create attribute<br>2. Create value "Red"<br>3. Call `AddValues(redValue)`<br>4. Verify values slice | Values slice length = 1<br>Values[0] matches redValue | ATTR-DOM-CREATE-1, ATTRVAL-DOM-CREATE-1 | Passed | 2025-11-25 | Automated: `TestAttribute_AddValues` |
| ATTR-DOM-ADDVAL-2 | Verify adding multiple values at once | 1. Create attribute<br>2. Create values "Red", "Blue", "Green"<br>3. Call `AddValues(red, blue, green)`<br>4. Verify values slice | Values slice length = 3<br>All values present | ATTR-DOM-CREATE-1, ATTRVAL-DOM-CREATE-1 | Passed | 2025-11-25 | Automated: `TestAttribute_AddValues` |
| ATTR-DOM-ADDVAL-3 | Verify appending to existing values | 1. Create attribute<br>2. Add value "Red"<br>3. Add value "Blue"<br>4. Verify values slice | Values slice length = 2<br>Contains both values in order | ATTR-DOM-CREATE-1, ATTRVAL-DOM-CREATE-1 | Passed | 2025-11-25 | Automated: `TestAttribute_AddValues` |

### Update Attribute Value

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-DOM-UPDATEVAL-1 | Verify updating attribute value successfully | 1. Create attribute with value "Red"<br>2. Call `UpdateValue(redID, &"Crimson Red")`<br>3. Verify value updated | No error returned<br>Value updated to "Crimson Red" | ATTR-DOM-CREATE-1, ATTRVAL-DOM-CREATE-1, ATTR-DOM-ADDVAL-1 | Passed | 2025-11-25 | Automated: `TestAttribute_UpdateValue` |
| ATTR-DOM-UPDATEVAL-2 | Verify value unchanged when nil provided | 1. Create attribute with value "Red"<br>2. Call `UpdateValue(redID, nil)`<br>3. Verify value unchanged | No error returned<br>Value remains "Red" | ATTR-DOM-CREATE-1, ATTRVAL-DOM-CREATE-1, ATTR-DOM-ADDVAL-1 | Passed | 2025-11-25 | Automated: `TestAttribute_UpdateValue` |
| ATTR-DOM-UPDATEVAL-3 | Verify error when value ID not found | 1. Create attribute with value "Red"<br>2. Call `UpdateValue(randomID, &"Blue")`<br>3. Verify error returned | Returns `ErrNotFound` error | ATTR-DOM-CREATE-1, ATTRVAL-DOM-CREATE-1, ATTR-DOM-ADDVAL-1 | Passed | 2025-11-25 | Automated: `TestAttribute_UpdateValue` |

### Soft Delete Attribute

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-DOM-DELETE-1 | Verify soft delete sets DeletedAt on attribute and values | 1. Create attribute with values "Red", "Blue"<br>2. Call `Remove()`<br>3. Verify DeletedAt timestamps | Attribute DeletedAt is set<br>All values have DeletedAt set<br>Timestamps are current | ATTR-DOM-CREATE-1, ATTRVAL-DOM-CREATE-1, ATTR-DOM-ADDVAL-2 | Passed | 2025-11-25 | Automated: `TestAttribute_Remove` |
| ATTR-DOM-DELETE-2 | Verify DeletedAt not updated if already set | 1. Create attribute with value<br>2. Call `Remove()` (first time)<br>3. Record DeletedAt timestamps<br>4. Call `Remove()` (second time)<br>5. Verify timestamps unchanged | DeletedAt timestamps remain the same | ATTR-DOM-CREATE-1, ATTRVAL-DOM-CREATE-1, ATTR-DOM-DELETE-1 | Passed | 2025-11-25 | Automated: `TestAttribute_Remove` |

### Remove Attribute Value

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-DOM-REMOVEVAL-1 | Verify removing value from attribute | 1. Create attribute with values "Red", "Blue", "Green"<br>2. Call `RemoveValue(blueID)`<br>3. Verify values slice | No error returned<br>Values length = 2<br>Contains "Red" and "Green"<br>Does not contain "Blue" | ATTR-DOM-CREATE-1, ATTRVAL-DOM-CREATE-1, ATTR-DOM-ADDVAL-2 | Passed | 2025-11-25 | Automated: `TestAttribute_RemoveValue` |
| ATTR-DOM-REMOVEVAL-2 | Verify error when attribute is nil | 1. Set attribute pointer to nil<br>2. Call `RemoveValue(anyID)`<br>3. Verify error returned | Returns `ErrInvalid` error | | Passed | 2025-11-25 | Automated: `TestAttribute_RemoveValue` |
| ATTR-DOM-REMOVEVAL-3 | Verify no error when value not found | 1. Create attribute with value "Red"<br>2. Call `RemoveValue(randomID)`<br>3. Verify no error | No error returned<br>Values slice unchanged | ATTR-DOM-CREATE-1, ATTRVAL-DOM-CREATE-1, ATTR-DOM-ADDVAL-1 | Passed | 2025-11-25 | Automated: `TestAttribute_RemoveValue` |
| ATTR-DOM-REMOVEVAL-4 | Verify removing all matching values | 1. Create attribute with values "Red", "Blue"<br>2. Call `RemoveValue(redID)`<br>3. Verify only "Blue" remains | Values length = 1<br>Contains only "Blue" | ATTR-DOM-CREATE-1, ATTRVAL-DOM-CREATE-1, ATTR-DOM-ADDVAL-2 | Passed | 2025-11-25 | Automated: `TestAttribute_RemoveValue` |

---

## Domain Validation Tests

### Attribute Value Uniqueness

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-VAL-UNIQUE-1 | Verify validation passes for unique values | 1. Create attribute with values "Red", "Blue", "Green"<br>2. Call `ValidateAttributeValueUniqueness()`<br>3. Verify no error | No error returned | ATTR-DOM-CREATE-1, ATTRVAL-DOM-CREATE-1, ATTR-DOM-ADDVAL-2 | Passed | 2025-11-25 | Automated: `TestValidateAttributeValueUniqueness` |
| ATTR-VAL-UNIQUE-2 | Verify validation fails for duplicate values | 1. Create attribute with values "Red", "Red"<br>2. Call `ValidateAttributeValueUniqueness()`<br>3. Verify error returned | Error contains "duplicate attribute value found" | ATTR-DOM-CREATE-1, ATTRVAL-DOM-CREATE-1 | Passed | 2025-11-25 | Automated: `TestValidateAttributeValueUniqueness` |
| ATTR-VAL-UNIQUE-3 | Verify validation fails for duplicate IDs | 1. Create attribute<br>2. Add two values with same ID<br>3. Call `ValidateAttributeValueUniqueness()`<br>4. Verify error returned | Error contains "duplicate attribute value ID found" | ATTR-DOM-CREATE-1 | Passed | 2025-11-25 | Automated: `TestValidateAttributeValueUniqueness` |
| ATTR-VAL-UNIQUE-4 | Verify validation passes for empty values | 1. Create attribute with no values<br>2. Call `ValidateAttributeValueUniqueness()`<br>3. Verify no error | No error returned | ATTR-DOM-CREATE-1 | Passed | 2025-11-25 | Automated: `TestValidateAttributeValueUniqueness` |

### Unique Attribute IDs

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-VAL-UNIQID-1 | Verify validation passes for unique attribute IDs | 1. Create attributes "color", "size", "brand"<br>2. Call `ValidateUniqueAttributeIDs()`<br>3. Verify no error | No error returned | ATTR-DOM-CREATE-1 | Passed | 2025-11-25 | Automated: `TestValidateUniqueAttributeIDs` |
| ATTR-VAL-UNIQID-2 | Verify validation fails for duplicate attribute IDs | 1. Create two attributes with same ID<br>2. Call `ValidateUniqueAttributeIDs()`<br>3. Verify error returned | Error contains "duplicate attribute ID found" | ATTR-DOM-CREATE-1 | Passed | 2025-11-25 | Automated: `TestValidateUniqueAttributeIDs` |
| ATTR-VAL-UNIQID-3 | Verify validation passes for empty slice | 1. Create empty attribute slice<br>2. Call `ValidateUniqueAttributeIDs()`<br>3. Verify no error | No error returned | | Passed | 2025-11-25 | Automated: `TestValidateUniqueAttributeIDs` |

### Unique Attribute Codes

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-VAL-UNIQCODE-1 | Verify validation passes for unique codes | 1. Create attributes with codes "color", "size", "brand"<br>2. Call `ValidateUniqueAttributeCodes()`<br>3. Verify no error | No error returned | ATTR-DOM-CREATE-1 | Passed | 2025-11-25 | Automated: `TestValidateUniqueAttributeCodes` |
| ATTR-VAL-UNIQCODE-2 | Verify validation fails for duplicate codes | 1. Create two attributes with code "color"<br>2. Call `ValidateUniqueAttributeCodes()`<br>3. Verify error returned | Error contains "duplicate attribute code found" and "color" | ATTR-DOM-CREATE-1 | Passed | 2025-11-25 | Automated: `TestValidateUniqueAttributeCodes` |
| ATTR-VAL-UNIQCODE-3 | Verify validation passes for empty slice | 1. Create empty attribute slice<br>2. Call `ValidateUniqueAttributeCodes()`<br>3. Verify no error | No error returned | | Passed | 2025-11-25 | Automated: `TestValidateUniqueAttributeCodes` |

### Validator Registration

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-VAL-REG-1 | Verify attribute validators register successfully | 1. Create new validator<br>2. Call `RegisterAttributeValidators(validator)`<br>3. Verify no error | No error returned<br>Custom validators registered | | Passed | 2025-11-25 | Automated: `TestRegisterAttributeValidators` |
| ATTR-VAL-REG-2 | Verify custom validation tag works after registration | 1. Register validators<br>2. Create attribute with duplicate values<br>3. Call `validator.Struct(attribute)`<br>4. Verify validation fails | Validation error contains "unique_attribute_values" | ATTR-VAL-REG-1 | Passed | 2025-11-25 | Automated: `TestValidateUniqueAttributeValues` |

---

## Service Layer Tests

### Service Initialization

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-SVC-INIT-1 | Verify attribute service creation | 1. Create validator instance<br>2. Call `ProvideAttribute(validator)`<br>3. Verify service returned | Service not nil<br>Validator is set | | Passed | 2025-11-25 | Automated: `TestProvideAttribute` |

### Service Validation

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-SVC-VAL-1 | Verify validation passes for valid attribute | 1. Create valid attribute with values<br>2. Call `service.Validate(attribute)`<br>3. Verify no error | No error returned | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_Validate` |
| ATTR-SVC-VAL-2 | Verify validation fails for empty code | 1. Create attribute with empty code<br>2. Call `service.Validate(attribute)`<br>3. Verify error returned | Error is `ErrInvalid` | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_Validate` |
| ATTR-SVC-VAL-3 | Verify validation fails for empty name | 1. Create attribute with empty name<br>2. Call `service.Validate(attribute)`<br>3. Verify error returned | Error is `ErrInvalid` | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_Validate` |
| ATTR-SVC-VAL-4 | Verify validation fails for code too short | 1. Create attribute with code "c" (1 char)<br>2. Call `service.Validate(attribute)`<br>3. Verify error returned | Error is `ErrInvalid` | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_Validate` |
| ATTR-SVC-VAL-5 | Verify validation fails for code too long | 1. Create attribute with code >50 chars<br>2. Call `service.Validate(attribute)`<br>3. Verify error returned | Error is `ErrInvalid` | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_Validate` |
| ATTR-SVC-VAL-6 | Verify validation fails for name too short | 1. Create attribute with name "C" (1 char)<br>2. Call `service.Validate(attribute)`<br>3. Verify error returned | Error is `ErrInvalid` | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_Validate` |
| ATTR-SVC-VAL-7 | Verify validation fails for name too long | 1. Create attribute with name >100 chars<br>2. Call `service.Validate(attribute)`<br>3. Verify error returned | Error is `ErrInvalid` | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_Validate` |
| ATTR-SVC-VAL-8 | Verify validation fails for empty attribute value | 1. Create attribute with empty value string<br>2. Call `service.Validate(attribute)`<br>3. Verify error returned | Error is `ErrInvalid` | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_Validate` |
| ATTR-SVC-VAL-9 | Verify validation fails for value too long | 1. Create attribute with value >100 chars<br>2. Call `service.Validate(attribute)`<br>3. Verify error returned | Error is `ErrInvalid` | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_Validate` |
| ATTR-SVC-VAL-10 | Verify validation fails for nil UUID | 1. Create attribute with uuid.Nil ID<br>2. Call `service.Validate(attribute)`<br>3. Verify error returned | Error is `ErrInvalid` | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_Validate` |
| ATTR-SVC-VAL-11 | Verify validation passes for multiple valid values | 1. Create attribute with 4 unique values<br>2. Call `service.Validate(attribute)`<br>3. Verify no error | No error returned | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_Validate` |
| ATTR-SVC-VAL-12 | Verify validation fails for duplicate values | 1. Create attribute with duplicate values<br>2. Call `service.Validate(attribute)`<br>3. Verify error returned | Error is `ErrInvalid` | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_Validate` |

### Filter Attribute Values

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-SVC-FILTER-1 | Verify filtering values from attributes | 1. Create attributes with values<br>2. Provide value IDs to filter<br>3. Call `FilterAttributeValuesFromAttributes()`<br>4. Verify filtered result | Returns only matching values<br>Count matches expected | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_FilterAttributeValuesFromAttributes` |
| ATTR-SVC-FILTER-2 | Verify empty result when no matches | 1. Create attributes<br>2. Provide non-existent value IDs<br>3. Call `FilterAttributeValuesFromAttributes()`<br>4. Verify empty result | Returns empty slice | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_FilterAttributeValuesFromAttributes` |
| ATTR-SVC-FILTER-3 | Verify empty result for empty attributes | 1. Provide empty attributes slice<br>2. Provide value IDs<br>3. Call `FilterAttributeValuesFromAttributes()`<br>4. Verify empty result | Returns empty slice | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_FilterAttributeValuesFromAttributes` |
| ATTR-SVC-FILTER-4 | Verify empty result for empty value IDs | 1. Create attributes with values<br>2. Provide empty value IDs<br>3. Call `FilterAttributeValuesFromAttributes()`<br>4. Verify empty result | Returns empty slice | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_FilterAttributeValuesFromAttributes` |
| ATTR-SVC-FILTER-5 | Verify filtering from multiple attributes | 1. Create 3 attributes with values<br>2. Provide mixed value IDs<br>3. Call `FilterAttributeValuesFromAttributes()`<br>4. Verify all matched values returned | Returns all matching values from all attributes | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_FilterAttributeValuesFromAttributes` |
| ATTR-SVC-FILTER-6 | Verify handling duplicate value IDs | 1. Create attribute<br>2. Provide duplicate value IDs in filter<br>3. Call `FilterAttributeValuesFromAttributes()`<br>4. Verify single result | Returns value only once, not duplicated | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_FilterAttributeValuesFromAttributes` |
| ATTR-SVC-FILTER-7 | Verify filtering all values | 1. Create attribute with 2 values<br>2. Provide both value IDs<br>3. Call `FilterAttributeValuesFromAttributes()`<br>4. Verify both returned | Returns all values | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_FilterAttributeValuesFromAttributes` |
| ATTR-SVC-FILTER-8 | Verify order preservation | 1. Create attribute with 3 values<br>2. Provide all IDs in order<br>3. Call `FilterAttributeValuesFromAttributes()`<br>4. Verify order preserved | Values returned in same order as in attribute | ATTR-SVC-INIT-1 | Passed | 2025-11-25 | Automated: `TestAttribute_FilterAttributeValuesFromAttributes` |

---

## Repository Layer Tests

### Repository Count

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-REPO-COUNT-1 | Verify counting all attributes | 1. Setup mock to return count of 5<br>2. Call `repo.Count(ctx, nil, DeletedExcludeParam)`<br>3. Verify count returned | Count = 5<br>No error | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_Count` |
| ATTR-REPO-COUNT-2 | Verify counting specific attribute IDs | 1. Setup mock with 2 specific IDs<br>2. Call `repo.Count(ctx, &ids, DeletedExcludeParam)`<br>3. Verify count | Count = 2<br>No error | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_Count` |
| ATTR-REPO-COUNT-3 | Verify counting including deleted | 1. Setup mock to return count of 10<br>2. Call `repo.Count(ctx, nil, DeletedIncludeParam)`<br>3. Verify count | Count = 10<br>No error | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_Count` |
| ATTR-REPO-COUNT-4 | Verify error handling on database failure | 1. Setup mock to return error<br>2. Call `repo.Count(ctx, nil, DeletedExcludeParam)`<br>3. Verify error returned | Returns `ErrInternal`<br>Count is nil | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_Count` |

### Repository List

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-REPO-LIST-1 | Verify listing all attributes | 1. Setup mock to return 2 attributes<br>2. Call `repo.List(ctx, nil, nil, DeletedExcludeParam, 10, 0)`<br>3. Verify attributes returned | Returns 2 attributes<br>No error | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_List` |
| ATTR-REPO-LIST-2 | Verify listing with search filter | 1. Setup mock with search "color"<br>2. Call `repo.List(ctx, nil, &search, DeletedExcludeParam, 10, 0)`<br>3. Verify filtered results | Returns matching attributes<br>Code contains "color" | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_List` |
| ATTR-REPO-LIST-3 | Verify pagination support | 1. Setup mock with limit=5, offset=10<br>2. Call `repo.List(ctx, nil, nil, DeletedExcludeParam, 5, 10)`<br>3. Verify paginated results | Returns paginated attributes<br>No error | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_List` |
| ATTR-REPO-LIST-4 | Verify listing by specific IDs | 1. Setup mock with specific IDs<br>2. Call `repo.List(ctx, &ids, nil, DeletedExcludeParam, 10, 0)`<br>3. Verify filtered results | Returns only specified attributes<br>No error | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_List` |
| ATTR-REPO-LIST-5 | Verify error handling on database failure | 1. Setup mock to return error<br>2. Call `repo.List(ctx, nil, nil, DeletedExcludeParam, 10, 0)`<br>3. Verify error returned | Returns `ErrInternal`<br>Result is nil | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_List` |

### Repository Get

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-REPO-GET-1 | Verify getting attribute by ID with values | 1. Setup mock to return attribute with 2 values<br>2. Call `repo.Get(ctx, attributeID)`<br>3. Verify attribute returned | Attribute returned with correct ID, code, name<br>Values length = 2<br>No error | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_Get` |
| ATTR-REPO-GET-2 | Verify error when attribute not found | 1. Setup mock to return `ErrNotFound`<br>2. Call `repo.Get(ctx, nonExistentID)`<br>3. Verify error returned | Returns `ErrNotFound`<br>Result is nil | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_Get` |
| ATTR-REPO-GET-3 | Verify error on database failure | 1. Setup mock to return `ErrInternal`<br>2. Call `repo.Get(ctx, attributeID)`<br>3. Verify error returned | Returns `ErrInternal`<br>Result is nil | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_Get` |

### Repository List Values

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-REPO-LISTVAL-1 | Verify listing all values for attribute | 1. Setup mock to return 2 values<br>2. Call `repo.ListValues(ctx, attrID, nil, nil, DeletedExcludeParam, 10, 0)`<br>3. Verify values returned | Returns 2 values<br>No error | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_ListValues` |
| ATTR-REPO-LISTVAL-2 | Verify listing values with search | 1. Setup mock with search "Red"<br>2. Call `repo.ListValues(ctx, attrID, nil, &search, DeletedExcludeParam, 10, 0)`<br>3. Verify filtered results | Returns matching values<br>No error | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_ListValues` |
| ATTR-REPO-LISTVAL-3 | Verify listing specific value IDs | 1. Setup mock with specific value IDs<br>2. Call `repo.ListValues(ctx, attrID, &valueIDs, nil, DeletedExcludeParam, 10, 0)`<br>3. Verify filtered results | Returns only specified values<br>No error | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_ListValues` |
| ATTR-REPO-LISTVAL-4 | Verify error on database failure | 1. Setup mock to return error<br>2. Call `repo.ListValues(ctx, attrID, nil, nil, DeletedExcludeParam, 10, 0)`<br>3. Verify error returned | Returns `ErrInternal`<br>Result is nil | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_ListValues` |

### Repository Count Values

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-REPO-COUNTVAL-1 | Verify counting all values for attribute | 1. Setup mock to return count of 3<br>2. Call `repo.CountValues(ctx, attrID, nil)`<br>3. Verify count returned | Count = 3<br>No error | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_CountValues` |
| ATTR-REPO-COUNTVAL-2 | Verify counting specific value IDs | 1. Setup mock with 2 value IDs<br>2. Call `repo.CountValues(ctx, attrID, &valueIDs)`<br>3. Verify count | Count = 2<br>No error | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_CountValues` |
| ATTR-REPO-COUNTVAL-3 | Verify error on database failure | 1. Setup mock to return error<br>2. Call `repo.CountValues(ctx, attrID, nil)`<br>3. Verify error returned | Returns `ErrInternal`<br>Count is nil | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_CountValues` |

### Repository Save

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-REPO-SAVE-1 | Verify saving new attribute with values | 1. Create attribute with 2 values<br>2. Setup mock to succeed<br>3. Call `repo.Save(ctx, attribute)`<br>4. Verify success | No error returned | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_Save` |
| ATTR-REPO-SAVE-2 | Verify saving attribute without values | 1. Create attribute with no values<br>2. Setup mock to succeed<br>3. Call `repo.Save(ctx, attribute)`<br>4. Verify success | No error returned | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_Save` |
| ATTR-REPO-SAVE-3 | Verify conflict error handling | 1. Create attribute<br>2. Setup mock to return `ErrConflict`<br>3. Call `repo.Save(ctx, attribute)`<br>4. Verify error | Returns `ErrConflict` | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_Save` |
| ATTR-REPO-SAVE-4 | Verify database failure handling | 1. Create attribute<br>2. Setup mock to return `ErrInternal`<br>3. Call `repo.Save(ctx, attribute)`<br>4. Verify error | Returns `ErrInternal` | | Passed | 2025-11-25 | Automated: `TestAttributeRepository_Save` |
| ATTR-REPO-SAVE-5 | Verify updating existing attribute | 1. Create and save attribute<br>2. Update attribute name and add value<br>3. Setup mock to succeed<br>4. Call `repo.Save(ctx, attribute)`<br>5. Verify success | No error returned | ATTR-REPO-SAVE-1 | Passed | 2025-11-25 | Automated: `TestAttributeRepository_Save` |

### Repository Complex Scenarios

| ID | Test Case Description | Test Case Procedure | Expected Output | Inter-test Case Dependence | Result | Test Date | Note |
|----|----------------------|---------------------|-----------------|---------------------------|--------|-----------|------|
| ATTR-REPO-COMPLEX-1 | Verify create, retrieve, and update flow | 1. Create attribute with value<br>2. Save attribute<br>3. Get attribute by ID<br>4. Update attribute name<br>5. Save updated attribute<br>6. Verify all operations succeed | All operations succeed<br>No errors | ATTR-REPO-SAVE-1, ATTR-REPO-GET-1 | Passed | 2025-11-25 | Automated: `TestAttributeRepository_ComplexScenarios` |
| ATTR-REPO-COMPLEX-2 | Verify listing attributes and their values | 1. Create 2 attributes with values<br>2. List all attributes<br>3. List values for each attribute<br>4. Verify results | Attributes listed correctly<br>Values listed correctly per attribute | ATTR-REPO-LIST-1, ATTR-REPO-LISTVAL-1 | Passed | 2025-11-25 | Automated: `TestAttributeRepository_ComplexScenarios` |

---

## Test Execution Summary

### Coverage Statistics

- **Domain Tests**: 100% coverage of domain methods
- **Service Tests**: 100% coverage of service methods
- **Repository Tests**: 100% coverage of repository interface
- **Validator Tests**: 100% coverage of validation functions

### Test Execution Environment

- **Go Version**: 1.21+
- **Test Framework**: Go testing + testify
- **Mock Framework**: testify/mock + mockery
- **Parallel Execution**: Enabled for independent tests

### Running Tests

```bash
# Run all attribute tests
go test ./internal/domain -run TestAttribute -v
go test ./internal/service -run TestAttribute -v
go test ./internal/infrastructure/repositorypostgres -run TestAttribute -v

# Run with coverage
go test -cover ./internal/domain ./internal/service ./internal/infrastructure/repositorypostgres

# Run with race detector
go test -race ./internal/domain ./internal/service ./internal/infrastructure/repositorypostgres

# Run specific test
go test ./internal/domain -run TestNewAttribute -v
```

---

## Notes

1. All tests use `t.Parallel()` to enable concurrent execution
2. Tests are organized using table-driven pattern for better maintainability
3. Service tests use registered validators to match production behavior
4. Repository tests use mock objects generated by mockery
5. Validators must be registered in `internal/client/validate.go` before use
6. Test IDs follow pattern: `<MODULE>-<LAYER>-<ACTION>-<NUMBER>`

---

## Future Test Enhancements

1. **Integration Tests**: Add end-to-end tests with testcontainers
2. **Performance Tests**: Add benchmarks for critical operations
3. **Stress Tests**: Test with large datasets (1000+ attributes, values)
4. **Concurrency Tests**: Verify thread-safe operations
5. **Error Injection**: Test recovery from various failure scenarios
