### `domain.NewProduct`

#### Meta

|                  |     |                                                                 |     |     |                    |     |     |     |     | 0   | 1                 | 2   | 3   | 4                |
| ---------------- | --- | --------------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ----------------- | --- | --- | ---------------- |
| Function Code    |     | P-NP-01                                                         |     |     | Function Name      |     |     |     |     |     | domain.NewProduct |     |     |                  |
| Created By       |     | CodeCompanion                                                   |     |     | Executed By        |     |     |     |     |     |                   |     |     |                  |
| Lines of code    |     | 20                                                              |     |     | Lack of test cases |     |     |     |     |     | 0                 |     |     |                  |
| Test requirement |     | Validate Product name, description, category ID boundary values |     |     |                    |     |     |     |     |     |                   |     |     |                  |
| Passed           |     | Failed                                                          |     |     | Untested           |     |     |     |     |     | N                 | A   | B   | Total Test Cases |
| 12               |     | 0                                                               |     |     | 0                  |     |     |     |     |     | 3                 | 2   | 7   | 12               |

#### Sheet

| #   |           |                                              |     |                                                                                                |     | P-NP-01-01 | P-NP-01-02 | P-NP-01-03 | P-NP-01-04 | P-NP-01-05 | P-NP-01-06 | P-NP-01-07 | P-NP-01-08 | P-NP-01-09 | P-NP-01-10 | P-NP-01-11 | P-NP-01-12 |
| --- | --------- | -------------------------------------------- | --- | ---------------------------------------------------------------------------------------------- | --- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- |
| 1   | Condition | Name Length                                  |     |                                                                                                |     |            |            |            |            |            |            |            |            |            |            |            |            |
| 2   |           |                                              |     | 0 (empty)                                                                                      |     |            |            |            |            |            |            |            |            |            |            | O          |            |
| 3   |           |                                              |     | 2 (min - 1)                                                                                    |     | O          |            |            |            |            |            |            |            |            |            |            |            |
| 4   |           |                                              |     | 3 (min)                                                                                        |     |            | O          |            |            |            |            |            |            |            |            |            |            |
| 5   |           |                                              |     | 4 (min + 1)                                                                                    |     |            |            | O          |            |            |            |            |            |            |            |            |            |
| 6   |           |                                              |     | 200 (max)                                                                                      |     |            |            |            | O          |            |            |            |            |            |            |            |            |
| 7   |           |                                              |     | 201 (max + 1)                                                                                  |     |            |            |            |            | O          |            |            |            |            |            |            |            |
| 8   |           | Description Length                           |     |                                                                                                |     |            |            |            |            |            |            |            |            |            |            |            |            |
| 9   |           |                                              |     | 0 (empty)                                                                                      |     |            |            |            |            |            |            |            |            |            |            |            | O          |
| 10  |           |                                              |     | 9 (min - 1)                                                                                    |     |            |            |            |            |            | O          |            |            |            |            |            |            |
| 11  |           |                                              |     | 10 (min)                                                                                       |     |            |            |            |            |            |            | O          |            |            |            |            |            |
| 12  |           |                                              |     | 11 (min + 1)                                                                                   |     |            |            |            |            |            |            |            | O          |            |            |            |            |
| 13  |           |                                              |     | Valid (normal)                                                                                 |     | O          | O          | O          | O          | O          |            | O          | O          | O          | O          |            |            |
| 14  |           | Category ID                                  |     |                                                                                                |     |            |            |            |            |            |            |            |            |            |            |            |            |
| 15  |           |                                              |     | Nil UUID                                                                                       |     |            |            |            |            |            |            |            |            |            |            |            | O          |
| 16  |           |                                              |     | Valid UUID                                                                                     |     | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          |            |
| 17  | Confirm   | Return                                       |     |                                                                                                |     |            |            |            |            |            |            |            |            |            |            |            |            |
| 18  |           | Name                                         |     |                                                                                                |     |            |            |            |            |            |            |            |            |            |            |            |            |
| 19  |           |                                              |     | Input Name                                                                                     |     | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          |
| 20  |           | Description                                  |     |                                                                                                |     |            |            |            |            |            |            |            |            |            |            |            |            |
| 21  |           |                                              |     | Input Description                                                                              |     | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          |
| 22  |           | Category ID                                  |     |                                                                                                |     |            |            |            |            |            |            |            |            |            |            |            |            |
| 23  |           |                                              |     | Input Category ID                                                                              |     | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          |            |
| 24  |           | ID                                           |     |                                                                                                |     |            |            |            |            |            |            |            |            |            |            |            |            |
| 25  |           |                                              |     | Not Nil                                                                                        |     | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          |
| 26  |           | CreatedAt & UpdatedAt                        |     |                                                                                                |     |            |            |            |            |            |            |            |            |            |            |            |            |
| 27  |           |                                              |     | Equal and Not Zero                                                                             |     | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          |
| 28  |           | Error                                        |     |                                                                                                |     |            |            |            |            |            |            |            |            |            |            |            |            |
| 29  |           |                                              |     | Key: 'Product.Name' Error:Field validation for 'Name' failed on the 'gte' tag                  |     | O          |            |            |            |            |            |            |            |            |            | O          |            |
| 30  |           |                                              |     | Key: 'Product.Name' Error:Field validation for 'Name' failed on the 'lte' tag                  |     |            |            |            |            | O          |            |            |            |            |            |            |            |
| 31  |           |                                              |     | Key: 'Product.Description' Error:Field validation for 'Description' failed on the 'gte' tag    |     |            |            |            |            |            | O          |            |            |            |            |            | O          |
| 32  |           |                                              |     | Key: 'Product.CategoryID' Error:Field validation for 'CategoryID' failed on the 'required' tag |     |            |            |            |            |            |            |            |            |            |            |            | O          |
| 33  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                                                                                                |     | B          | B          | B          | B          | B          | B          | B          | B          | N          | N          | A          | A          |
| 34  |           | Passed/Failed                                |     |                                                                                                |     | P          | P          | P          | P          | P          | P          | P          | P          | P          | P          | P          | P          |
| 35  |           | Executed Date                                |     |                                                                                                |     | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 |
| 36  |           | Defect ID                                    |     |                                                                                                |     |            |            |            |            |            |            |            |            |            |            |            |            |

### `domain.NewProductOption`

#### Meta

|                  |     |                                                     |     |     |                    |     |     |     |     | 0   | 1                       | 2   | 3   | 4                |
| ---------------- | --- | --------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ----------------------- | --- | --- | ---------------- |
| Function Code    |     | P-NPO-01                                            |     |     | Function Name      |     |     |     |     |     | domain.NewProductOption |     |     |                  |
| Created By       |     | CodeCompanion                                       |     |     | Executed By        |     |     |     |     |     |                         |     |     |                  |
| Lines of code    |     | 10                                                  |     |     | Lack of test cases |     |     |     |     |     | 0                       |     |     |                  |
| Test requirement |     | Validate ProductOption creation and name validation |     |     |                    |     |     |     |     |     |                         |     |     |                  |
| Passed           |     | Failed                                              |     |     | Untested           |     |     |     |     |     | N                       | A   | B   | Total Test Cases |
| 3                |     | 0                                                   |     |     | 0                  |     |     |     |     |     | 2                       | 1   | 0   | 3                |

#### Sheet

| #   |           |                                              |     |                  |     | P-NPO-01-01 | P-NPO-01-02 | P-NPO-01-03 |
| --- | --------- | -------------------------------------------- | --- | ---------------- | --- | ----------- | ----------- | ----------- |
| 1   | Condition | Option Name                                  |     |                  |     |             |             |             |
| 2   |           |                                              |     | "" (empty)       |     | O           |             |             |
| 3   |           |                                              |     | Single char      |     |             |             | O           |
| 4   |           |                                              |     | Valid (normal)   |     |             | O           |             |
| 5   | Confirm   | Return                                       |     |                  |     |             |             |             |
| 6   |           | Name                                         |     |                  |     |             |             |             |
| 7   |           |                                              |     | Input Name       |     | O           | O           | O           |
| 8   |           | ID                                           |     |                  |     |             |             |             |
| 9   |           |                                              |     | Not Nil          |     | O           | O           | O           |
| 10  |           | Error                                        |     |                  |     |             |             |             |
| 11  |           |                                              |     | Validation Error |     | O           |             |             |
| 12  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                  |     | A           | N           | N           |
| 13  |           | Passed/Failed                                |     |                  |     | P           | P           | P           |
| 14  |           | Executed Date                                |     |                  |     | 2025-12-12  | 2025-12-12  | 2025-12-12  |
| 15  |           | Defect ID                                    |     |                  |     |             |             |             |

### `domain.NewProductImage`

#### Meta

|                  |     |                                             |     |     |                    |     |     |     |     | 0   | 1                      | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ---------------------- | --- | --- | ---------------- |
| Function Code    |     | P-NPI-01                                    |     |     | Function Name      |     |     |     |     |     | domain.NewProductImage |     |     |                  |
| Created By       |     | CodeCompanion                               |     |     | Executed By        |     |     |     |     |     |                        |     |     |                  |
| Lines of code    |     | 15                                          |     |     | Lack of test cases |     |     |     |     |     | 0                      |     |     |                  |
| Test requirement |     | Validate ProductImage order boundary values |     |     |                    |     |     |     |     |     |                        |     |     |                  |
| Passed           |     | Failed                                      |     |     | Untested           |     |     |     |     |     | N                      | A   | B   | Total Test Cases |
| 3                |     | 0                                           |     |     | 0                  |     |     |     |     |     | 2                      | 1   | 0   | 3                |

#### Sheet

| #   |           |                                              |     |                  |     | P-NPI-01-01 | P-NPI-01-02 | P-NPI-01-03 |
| --- | --------- | -------------------------------------------- | --- | ---------------- | --- | ----------- | ----------- | ----------- |
| 1   | Condition | Order Value                                  |     |                  |     |             |             |             |
| 2   |           |                                              |     | -1 (negative)    |     | O           |             |             |
| 3   |           |                                              |     | 1 (min + 1)      |     |             | O           |             |
| 4   |           |                                              |     | 100 (large)      |     |             |             | O           |
| 5   | Confirm   | Return                                       |     |                  |     |             |             |             |
| 6   |           | Order                                        |     |                  |     |             |             |             |
| 7   |           |                                              |     | Input Order      |     | O           | O           | O           |
| 8   |           | ID                                           |     |                  |     |             |             |             |
| 9   |           |                                              |     | Not Nil          |     | O           | O           | O           |
| 10  |           | URL                                          |     |                  |     |             |             |             |
| 11  |           |                                              |     | Valid URL        |     | O           | O           | O           |
| 12  |           | CreatedAt                                    |     |                  |     |             |             |             |
| 13  |           |                                              |     | Not Zero         |     | O           | O           | O           |
| 14  |           | Error                                        |     |                  |     |             |             |             |
| 15  |           |                                              |     | Validation Error |     | O           |             |             |
| 16  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                  |     | A           | N           | B           |
| 17  |           | Passed/Failed                                |     |                  |     | P           | P           | P           |
| 18  |           | Executed Date                                |     |                  |     | 2025-12-12  | 2025-12-12  | 2025-12-12  |
| 19  |           | Defect ID                                    |     |                  |     |             |             |             |

### `domain.NewVariant`

#### Meta

|                  |     |                                                            |     |     |                    |     |     |     |     | 0   | 1                 | 2   | 3   | 4                |
| ---------------- | --- | ---------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ----------------- | --- | --- | ---------------- |
| Function Code    |     | P-NV-01                                                    |     |     | Function Name      |     |     |     |     |     | domain.NewVariant |     |     |                  |
| Created By       |     | CodeCompanion                                              |     |     | Executed By        |     |     |     |     |     |                   |     |     |                  |
| Lines of code    |     | 15                                                         |     |     | Lack of test cases |     |     |     |     |     | 0                 |     |     |                  |
| Test requirement |     | Validate ProductVariant price and quantity boundary values |     |     |                    |     |     |     |     |     |                   |     |     |                  |
| Passed           |     | Failed                                                     |     |     | Untested           |     |     |     |     |     | N                 | A   | B   | Total Test Cases |
| 8                |     | 0                                                          |     |     | 0                  |     |     |     |     |     | 3                 | 2   | 3   | 8                |

#### Sheet

| #   |           |                                              |     |                                                                           |     | P-NV-01-01 | P-NV-01-02 | P-NV-01-03 | P-NV-01-04 | P-NV-01-05 | P-NV-01-06 | P-NV-01-07 | P-NV-01-08 |
| --- | --------- | -------------------------------------------- | --- | ------------------------------------------------------------------------- | --- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- |
| 1   | Condition | Price                                        |     |                                                                           |     |            |            |            |            |            |            |            |            |
| 2   |           |                                              |     | 0 (invalid)                                                               |     | O          |            |            |            |            |            |            |            |
| 3   |           |                                              |     | 1 (min)                                                                   |     |            | O          |            |            |            |            |            |            |
| 4   |           |                                              |     | Negative                                                                  |     |            |            | O          |            |            |            |            |            |
| 5   |           |                                              |     | Valid (normal)                                                            |     |            |            |            |            |            |            | O          | O          |
| 6   |           | Quantity                                     |     |                                                                           |     |            |            |            |            |            |            |            |            |
| 7   |           |                                              |     | -1 (negative)                                                             |     |            |            |            |            | O          |            |            |            |
| 8   |           |                                              |     | 0 (min)                                                                   |     |            |            |            | O          |            |            | O          |            |
| 9   |           |                                              |     | 1 (min + 1)                                                               |     |            |            |            |            |            |            |            | O          |
| 10  |           | SKU                                          |     |                                                                           |     |            |            |            |            |            |            |            |            |
| 11  |           |                                              |     | "" (empty)                                                                |     |            |            |            |            |            | O          |            |            |
| 12  |           |                                              |     | Valid (normal)                                                            |     | O          | O          | O          | O          | O          |            | O          | O          |
| 13  | Confirm   | Return                                       |     |                                                                           |     |            |            |            |            |            |            |            |            |
| 14  |           | SKU                                          |     |                                                                           |     |            |            |            |            |            |            |            |            |
| 15  |           |                                              |     | Input SKU                                                                 |     | O          | O          | O          | O          | O          | O          | O          | O          |
| 16  |           | Price                                        |     |                                                                           |     |            |            |            |            |            |            |            |            |
| 17  |           |                                              |     | Input Price                                                               |     | O          | O          | O          | O          | O          | O          | O          | O          |
| 18  |           | Quantity                                     |     |                                                                           |     |            |            |            |            |            |            |            |            |
| 19  |           |                                              |     | Input Quantity                                                            |     | O          | O          | O          | O          | O          | O          | O          | O          |
| 20  |           | PurchaseCount                                |     |                                                                           |     |            |            |            |            |            |            |            |            |
| 21  |           |                                              |     | 0 (init)                                                                  |     | O          | O          | O          | O          | O          | O          | O          | O          |
| 22  |           | ID                                           |     |                                                                           |     |            |            |            |            |            |            |            |            |
| 23  |           |                                              |     | Not Nil                                                                   |     | O          | O          | O          | O          | O          | O          | O          | O          |
| 24  |           | CreatedAt & UpdatedAt                        |     |                                                                           |     |            |            |            |            |            |            |            |            |
| 25  |           |                                              |     | Not Zero                                                                  |     | O          | O          | O          | O          | O          | O          | O          | O          |
| 26  |           | Error                                        |     |                                                                           |     |            |            |            |            |            |            |            |            |
| 27  |           |                                              |     | Key: 'ProductVariant.Price' Error:Field validation failed on 'gt' tag     |     | O          |            |            |            |            |            |            |            |
| 28  |           |                                              |     | Key: 'ProductVariant.Price' Error:Field validation failed on 'gt' tag     |     |            |            | O          |            |            |            |            |            |
| 29  |           |                                              |     | Key: 'ProductVariant.Quantity' Error:Field validation failed on 'gte' tag |     |            |            |            |            | O          |            |            |            |
| 30  |           |                                              |     | Key: 'ProductVariant.SKU' Error:Field validation failed on 'required' tag |     |            |            |            |            |            | O          |            |            |
| 31  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                                                                           |     | B          | B          | B          | B          | B          | A          | N          | N          |
| 32  |           | Passed/Failed                                |     |                                                                           |     | P          | P          | P          | P          | P          | P          | P          | P          |
| 33  |           | Executed Date                                |     |                                                                           |     | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 |
| 34  |           | Defect ID                                    |     |                                                                           |     |            |            |            |            |            |            |            |            |

### `domain.CreateOptionValues`

#### Meta

|                  |     |                                     |     |     |                    |     |     |     |     | 0   | 1                         | 2   | 3   | 4                |
| ---------------- | --- | ----------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-COV-01                            |     |     | Function Name      |     |     |     |     |     | domain.CreateOptionValues |     |     |                  |
| Created By       |     | CodeCompanion                       |     |     | Executed By        |     |     |     |     |     |                           |     |     |                  |
| Lines of code    |     | 12                                  |     |     | Lack of test cases |     |     |     |     |     | 0                         |     |     |                  |
| Test requirement |     | Validate batch OptionValue creation |     |     |                    |     |     |     |     |     |                           |     |     |                  |
| Passed           |     | Failed                              |     |     | Untested           |     |     |     |     |     | N                         | A   | B   | Total Test Cases |
| 4                |     | 0                                   |     |     | 0                  |     |     |     |     |     | 4                         | 0   | 0   | 4                |

#### Sheet

| #   |           |                                              |     |                   |     | P-COV-01-01 | P-COV-01-02 | P-COV-01-03 | P-COV-01-04 |
| --- | --------- | -------------------------------------------- | --- | ----------------- | --- | ----------- | ----------- | ----------- | ----------- |
| 1   | Condition | Values List                                  |     |                   |     |             |             |             |             |
| 2   |           |                                              |     | Empty array       |     | O           |             |             |             |
| 3   |           |                                              |     | Single value      |     |             | O           |             |             |
| 4   |           |                                              |     | Multiple values   |     |             |             | O           |             |
| 5   |           |                                              |     | With empty string |     |             |             |             | O           |
| 6   | Confirm   | Return                                       |     |                   |     |             |             |             |             |
| 7   |           | Count                                        |     |                   |     |             |             |             |             |
| 8   |           |                                              |     | 0                 |     | O           |             |             |             |
| 9   |           |                                              |     | 1                 |     |             | O           |             |             |
| 10  |           |                                              |     | 3                 |     |             |             | O           |             |
| 11  |           |                                              |     | 3 (with empty)    |     |             |             |             | O           |
| 12  |           | ID                                           |     |                   |     |             |             |             |             |
| 13  |           |                                              |     | All Not Nil       |     | O           | O           | O           | O           |
| 14  |           | Value                                        |     |                   |     |             |             |             |             |
| 15  |           |                                              |     | Match Input       |     | O           | O           | O           | O           |
| 16  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                   |     | N           | N           | N           | N           |
| 17  |           | Passed/Failed                                |     |                   |     | P           | P           | P           | P           |
| 18  |           | Executed Date                                |     |                   |     | 2025-12-12  | 2025-12-12  | 2025-12-12  | 2025-12-12  |
| 19  |           | Defect ID                                    |     |                   |     |             |             |             |             |

### `domain.Product.Update`

#### Meta

|                  |     |                                              |     |     |                    |     |     |     |     | 0   | 1                     | 2   | 3   | 4                |
| ---------------- | --- | -------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | --------------------- | --- | --- | ---------------- |
| Function Code    |     | P-U-01                                       |     |     | Function Name      |     |     |     |     |     | domain.Product.Update |     |     |                  |
| Created By       |     | CodeCompanion                                |     |     | Executed By        |     |     |     |     |     |                       |     |     |                  |
| Lines of code    |     | 13                                           |     |     | Lack of test cases |     |     |     |     |     | 0                     |     |     |                  |
| Test requirement |     | Validate Product update logic and timestamps |     |     |                    |     |     |     |     |     |                       |     |     |                  |
| Passed           |     | Failed                                       |     |     | Untested           |     |     |     |     |     | N                     | A   | B   | Total Test Cases |
| 6                |     | 0                                            |     |     | 0                  |     |     |     |     |     | 6                     | 0   | 0   | 6                |

#### Sheet

| #   |           |                                              |     |                  |     | P-U-01-01  | P-U-01-02  | P-U-01-03  | P-U-01-04  | P-U-01-05  | P-U-01-06  |
| --- | --------- | -------------------------------------------- | --- | ---------------- | --- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- |
| 1   | Condition | Update Fields                                |     |                  |     |            |            |            |            |            |            |
| 2   |           |                                              |     | Name only        |     | O          |            |            |            |            |            |
| 3   |           |                                              |     | Description only |     |            | O          |            |            |            |            |
| 4   |           |                                              |     | Category ID only |     |            |            | O          |            |            |            |
| 5   |           |                                              |     | All fields       |     |            |            |            | O          |            |            |
| 6   |           |                                              |     | Empty values     |     |            |            |            |            | O          |            |
| 7   |           |                                              |     | Same values      |     |            |            |            |            |            | O          |
| 8   | Confirm   | Return                                       |     |                  |     |            |            |            |            |            |            |
| 9   |           | Name                                         |     |                  |     |            |            |            |            |            |            |
| 10  |           |                                              |     | Updated          |     | O          |            |            | O          | O          | O          |
| 11  |           |                                              |     | Original         |     |            | O          | O          |            |            |            |
| 12  |           | Description                                  |     |                  |     |            |            |            |            |            |            |
| 13  |           |                                              |     | Updated          |     |            | O          |            | O          |            |            |
| 14  |           |                                              |     | Original         |     | O          |            | O          |            | O          | O          |
| 15  |           | UpdatedAt                                    |     |                  |     |            |            |            |            |            |            |
| 16  |           |                                              |     | Changed          |     | O          | O          | O          | O          |            |            |
| 17  |           |                                              |     | Unchanged        |     |            |            |            |            | O          | O          |
| 18  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                  |     | N          | N          | N          | N          | N          | N          |
| 19  |           | Passed/Failed                                |     |                  |     | P          | P          | P          | P          | P          | P          |
| 20  |           | Executed Date                                |     |                  |     | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 |
| 21  |           | Defect ID                                    |     |                  |     |            |            |            |            |            |            |

### `domain.Product.UpdateVariant`

#### Meta

|                  |     |                                                   |     |     |                    |     |     |     |     | 0   | 1                            | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ---------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-UV-01                                           |     |     | Function Name      |     |     |     |     |     | domain.Product.UpdateVariant |     |     |                  |
| Created By       |     | CodeCompanion                                     |     |     | Executed By        |     |     |     |     |     |                              |     |     |                  |
| Lines of code    |     | 17                                                |     |     | Lack of test cases |     |     |     |     |     | 0                            |     |     |                  |
| Test requirement |     | Validate variant price/quantity update and errors |     |     |                    |     |     |     |     |     |                              |     |     |                  |
| Passed           |     | Failed                                            |     |     | Untested           |     |     |     |     |     | N                            | A   | B   | Total Test Cases |
| 6                |     | 0                                                 |     |     | 0                  |     |     |     |     |     | 5                            | 1   | 0   | 6                |

#### Sheet

| #   |           |                                              |     |                         |     | P-UV-01-01 | P-UV-01-02 | P-UV-01-03 | P-UV-01-04 | P-UV-01-05 | P-UV-01-06 |
| --- | --------- | -------------------------------------------- | --- | ----------------------- | --- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- |
| 1   | Condition | Update Fields                                |     |                         |     |            |            |            |            |            |            |
| 2   |           |                                              |     | Price only              |     | O          |            |            |            |            |            |
| 3   |           |                                              |     | Quantity only           |     |            | O          |            |            |            |            |
| 4   |           |                                              |     | Both price and quantity |     |            |            | O          |            |            |            |
| 5   |           |                                              |     | Zero values (no change) |     |            |            |            |            | O          |            |
| 6   |           |                                              |     | Different variant index |     |            |            |            |            |            | O          |
| 7   |           | Variant Existence                            |     |                         |     |            |            |            |            |            |            |
| 8   |           |                                              |     | Existing                |     | O          | O          | O          | O          | O          | O          |
| 9   |           |                                              |     | Non-existent            |     |            |            |            |            |            |            |
| 10  | Confirm   | Return                                       |     |                         |     |            |            |            |            |            |            |
| 11  |           | Price                                        |     |                         |     |            |            |            |            |            |            |
| 12  |           |                                              |     | Updated                 |     | O          |            | O          |            |            |            |
| 13  |           |                                              |     | Original                |     |            | O          |            | O          | O          | O          |
| 14  |           | Quantity                                     |     |                         |     |            |            |            |            |            |            |
| 15  |           |                                              |     | Updated                 |     |            | O          | O          |            |            | O          |
| 16  |           |                                              |     | Original                |     | O          |            |            | O          | O          |            |
| 17  |           | UpdatedAt                                    |     |                         |     |            |            |            |            |            |            |
| 18  |           |                                              |     | Changed                 |     | O          | O          | O          |            |            | O          |
| 19  |           |                                              |     | Unchanged               |     |            |            |            | O          | O          |            |
| 20  |           | Error                                        |     |                         |     |            |            |            |            |            |            |
| 21  |           |                                              |     | ErrNotFound             |     |            |            |            |            |            |            |
| 22  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                         |     | N          | N          | N          | N          | N          | N          |
| 23  |           | Passed/Failed                                |     |                         |     | P          | P          | P          | P          | P          | P          |
| 24  |           | Executed Date                                |     |                         |     | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 |
| 25  |           | Defect ID                                    |     |                         |     |            |            |            |            |            |            |

### `domain.Product.UpdateOption`

#### Meta

|                  |     |                                              |     |     |                    |     |     |     |     | 0   | 1                           | 2   | 3   | 4                |
| ---------------- | --- | -------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | --------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-UO-01                                      |     |     | Function Name      |     |     |     |     |     | domain.Product.UpdateOption |     |     |                  |
| Created By       |     | CodeCompanion                                |     |     | Executed By        |     |     |     |     |     |                             |     |     |                  |
| Lines of code    |     | 11                                           |     |     | Lack of test cases |     |     |     |     |     | 0                           |     |     |                  |
| Test requirement |     | Validate option name update logic and errors |     |     |                    |     |     |     |     |     |                             |     |     |                  |
| Passed           |     | Failed                                       |     |     | Untested           |     |     |     |     |     | N                           | A   | B   | Total Test Cases |
| 4                |     | 0                                            |     |     | 0                  |     |     |     |     |     | 3                           | 1   | 0   | 4                |

#### Sheet

| #   |           |                                              |     |                   |     | P-UO-01-01 | P-UO-01-02 | P-UO-01-03 | P-UO-01-04 |
| --- | --------- | -------------------------------------------- | --- | ----------------- | --- | ---------- | ---------- | ---------- | ---------- |
| 1   | Condition | New Option Name                              |     |                   |     |            |            |            |            |
| 2   |           |                                              |     | Valid new name    |     | O          |            |            |            |
| 3   |           |                                              |     | Empty (no change) |     |            | O          |            |            |
| 4   |           |                                              |     | Different index   |     |            |            |            | O          |
| 5   |           | Option Existence                             |     |                   |     |            |            |            |            |
| 6   |           |                                              |     | Existing          |     | O          | O          | O          |            |
| 7   |           |                                              |     | Non-existent      |     |            |            |            | O          |
| 8   | Confirm   | Return                                       |     |                   |     |            |            |            |            |
| 9   |           | Name                                         |     |                   |     |            |            |            |            |
| 10  |           |                                              |     | Updated           |     | O          |            |            |            |
| 11  |           |                                              |     | Original          |     |            | O          | O          |            |
| 12  |           | Error                                        |     |                   |     |            |            |            |            |
| 13  |           |                                              |     | ErrNotFound       |     |            |            |            | O          |
| 14  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                   |     | N          | N          | N          | A          |
| 15  |           | Passed/Failed                                |     |                   |     | P          | P          | P          | P          |
| 16  |           | Executed Date                                |     |                   |     | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 |
| 17  |           | Defect ID                                    |     |                   |     |            |            |            |            |

### `domain.Product.UpdateOptionValue`

#### Meta

|                  |     |                                               |     |     |                    |     |     |     |     | 0   | 1                                | 2   | 3   | 4                |
| ---------------- | --- | --------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | -------------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-UOV-01                                      |     |     | Function Name      |     |     |     |     |     | domain.Product.UpdateOptionValue |     |     |                  |
| Created By       |     | CodeCompanion                                 |     |     | Executed By        |     |     |     |     |     |                                  |     |     |                  |
| Lines of code    |     | 22                                            |     |     | Lack of test cases |     |     |     |     |     | 0                                |     |     |                  |
| Test requirement |     | Validate option value update logic and errors |     |     |                    |     |     |     |     |     |                                  |     |     |                  |
| Passed           |     | Failed                                        |     |     | Untested           |     |     |     |     |     | N                                | A   | B   | Total Test Cases |
| 4                |     | 0                                             |     |     | 0                  |     |     |     |     |     | 2                                | 2   | 0   | 4                |

#### Sheet

| #   |           |                                              |     |                   |     | P-UOV-01-01 | P-UOV-01-02 | P-UOV-01-03 | P-UOV-01-04 |
| --- | --------- | -------------------------------------------- | --- | ----------------- | --- | ----------- | ----------- | ----------- | ----------- |
| 1   | Condition | New Option Value                             |     |                   |     |             |             |             |             |
| 2   |           |                                              |     | Valid new value   |     | O           |             |             |             |
| 3   |           |                                              |     | Empty (no change) |     |             | O           |             |             |
| 4   |           | Option/Value Existence                       |     |                   |     |             |             |             |             |
| 5   |           |                                              |     | Both exist        |     | O           | O           |             |             |
| 6   |           |                                              |     | Option not exist  |     |             |             | O           |             |
| 7   |           |                                              |     | Value not exist   |     |             |             |             | O           |
| 8   | Confirm   | Return                                       |     |                   |     |             |             |             |             |
| 9   |           | Value                                        |     |                   |     |             |             |             |             |
| 10  |           |                                              |     | Updated           |     | O           |             |             |             |
| 11  |           |                                              |     | Original          |     |             | O           |             |             |
| 12  |           | Error                                        |     |                   |     |             |             |             |             |
| 13  |           |                                              |     | ErrNotFound       |     |             |             | O           | O           |
| 14  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                   |     | N           | N           | A           | A           |
| 15  |           | Passed/Failed                                |     |                   |     | P           | P           | P           | P           |
| 16  |           | Executed Date                                |     |                   |     | 2025-12-12  | 2025-12-12  | 2025-12-12  | 2025-12-12  |
| 17  |           | Defect ID                                    |     |                   |     |             |             |             |             |

### `domain.Product.GetOptionByID`

#### Meta

|                  |     |                                 |     |     |                    |     |     |     |     | 0   | 1                            | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ---------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-GOI-01                        |     |     | Function Name      |     |     |     |     |     | domain.Product.GetOptionByID |     |     |                  |
| Created By       |     | CodeCompanion                   |     |     | Executed By        |     |     |     |     |     |                              |     |     |                  |
| Lines of code    |     | 7                               |     |     | Lack of test cases |     |     |     |     |     | 0                            |     |     |                  |
| Test requirement |     | Validate option retrieval by ID |     |     |                    |     |     |     |     |     |                              |     |     |                  |
| Passed           |     | Failed                          |     |     | Untested           |     |     |     |     |     | N                            | A   | B   | Total Test Cases |
| 2                |     | 0                               |     |     | 0                  |     |     |     |     |     | 2                            | 0   | 0   | 2                |

#### Sheet

| #   |           |                                              |     |                 |     | P-GOI-01-01 | P-GOI-01-02 |
| --- | --------- | -------------------------------------------- | --- | --------------- | --- | ----------- | ----------- |
| 1   | Condition | Option ID                                    |     |                 |     |             |             |
| 2   |           |                                              |     | Existing        |     | O           |             |
| 3   |           |                                              |     | Non-existent    |     |             | O           |
| 4   | Confirm   | Return                                       |     |                 |     |             |             |
| 5   |           | Option                                       |     |                 |     |             |             |
| 6   |           |                                              |     | Found (pointer) |     | O           |             |
| 7   |           |                                              |     | Nil             |     |             | O           |
| 8   | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                 |     | N           | N           |
| 9   |           | Passed/Failed                                |     |                 |     | P           | P           |
| 10  |           | Executed Date                                |     |                 |     | 2025-12-12  | 2025-12-12  |
| 11  |           | Defect ID                                    |     |                 |     |             |             |

### `domain.Product.GetOptionsByIDs`

#### Meta

|                  |     |                                 |     |     |                    |     |     |     |     | 0   | 1                              | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ------------------------------ | --- | --- | ---------------- |
| Function Code    |     | P-GOIS-01                       |     |     | Function Name      |     |     |     |     |     | domain.Product.GetOptionsByIDs |     |     |                  |
| Created By       |     | CodeCompanion                   |     |     | Executed By        |     |     |     |     |     |                                |     |     |                  |
| Lines of code    |     | 12                              |     |     | Lack of test cases |     |     |     |     |     | 0                              |     |     |                  |
| Test requirement |     | Validate batch option retrieval |     |     |                    |     |     |     |     |     |                                |     |     |                  |
| Passed           |     | Failed                          |     |     | Untested           |     |     |     |     |     | N                              | A   | B   | Total Test Cases |
| 6                |     | 0                               |     |     | 0                  |     |     |     |     |     | 6                              | 0   | 0   | 6                |

#### Sheet

| #   |           |                                              |     |                                 |     | P-GOIS-01-01 | P-GOIS-01-02 | P-GOIS-01-03 | P-GOIS-01-04 | P-GOIS-01-05 | P-GOIS-01-06 |
| --- | --------- | -------------------------------------------- | --- | ------------------------------- | --- | ------------ | ------------ | ------------ | ------------ | ------------ | ------------ |
| 1   | Condition | Option ID List                               |     |                                 |     |              |              |              |              |              |              |
| 2   |           |                                              |     | Single existing                 |     | O            |              |              |              |              |              |
| 3   |           |                                              |     | Multiple existing               |     |              | O            |              |              |              |              |
| 4   |           |                                              |     | All options                     |     |              |              | O            |              |              |              |
| 5   |           |                                              |     | All non-existent                |     |              |              |              | O            |              |              |
| 6   |           |                                              |     | Mixed existing and non-existent |     |              |              |              |              | O            |              |
| 7   |           |                                              |     | Empty list                      |     |              |              |              |              |              | O            |
| 8   | Confirm   | Return                                       |     |                                 |     |              |              |              |              |              |              |
| 9   |           | Count                                        |     |                                 |     |              |              |              |              |              |              |
| 10  |           |                                              |     | 1                               |     | O            |              |              |              |              |              |
| 11  |           |                                              |     | 2                               |     |              | O            |              |              |              |              |
| 12  |           |                                              |     | 3                               |     |              |              | O            |              |              |              |
| 13  |           |                                              |     | 0                               |     |              |              |              | O            | O            | O            |
| 14  |           |                                              |     | 1 (mixed)                       |     |              |              |              |              | O            |              |
| 15  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                                 |     | N            | N            | N            | N            | N            | N            |
| 16  |           | Passed/Failed                                |     |                                 |     | P            | P            | P            | P            | P            | P            |
| 17  |           | Executed Date                                |     |                                 |     | 2025-12-12   | 2025-12-12   | 2025-12-12   | 2025-12-12   | 2025-12-12   | 2025-12-12   |
| 18  |           | Defect ID                                    |     |                                 |     |              |              |              |              |              |              |

### `domain.Product.GetVariantByID`

#### Meta

|                  |     |                                  |     |     |                    |     |     |     |     | 0   | 1                             | 2   | 3   | 4                |
| ---------------- | --- | -------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ----------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-GVI-01                         |     |     | Function Name      |     |     |     |     |     | domain.Product.GetVariantByID |     |     |                  |
| Created By       |     | CodeCompanion                    |     |     | Executed By        |     |     |     |     |     |                               |     |     |                  |
| Lines of code    |     | 7                                |     |     | Lack of test cases |     |     |     |     |     | 0                             |     |     |                  |
| Test requirement |     | Validate variant retrieval by ID |     |     |                    |     |     |     |     |     |                               |     |     |                  |
| Passed           |     | Failed                           |     |     | Untested           |     |     |     |     |     | N                             | A   | B   | Total Test Cases |
| 2                |     | 0                                |     |     | 0                  |     |     |     |     |     | 2                             | 0   | 0   | 2                |

#### Sheet

| #   |           |                                              |     |                 |     | P-GVI-01-01 | P-GVI-01-02 |
| --- | --------- | -------------------------------------------- | --- | --------------- | --- | ----------- | ----------- |
| 1   | Condition | Variant ID                                   |     |                 |     |             |             |
| 2   |           |                                              |     | Existing        |     | O           |             |
| 3   |           |                                              |     | Non-existent    |     |             | O           |
| 4   | Confirm   | Return                                       |     |                 |     |             |             |
| 5   |           | Variant                                      |     |                 |     |             |             |
| 6   |           |                                              |     | Found (pointer) |     | O           |             |
| 7   |           |                                              |     | Nil             |     |             | O           |
| 8   | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                 |     | N           | N           |
| 9   |           | Passed/Failed                                |     |                 |     | P           | P           |
| 10  |           | Executed Date                                |     |                 |     | 2025-12-12  | 2025-12-12  |
| 11  |           | Defect ID                                    |     |                 |     |             |             |

### `domain.Product.UpdateMinPrice`

#### Meta

|                  |     |                                    |     |     |                    |     |     |     |     | 0   | 1                             | 2   | 3   | 4                |
| ---------------- | --- | ---------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ----------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-UMP-01                           |     |     | Function Name      |     |     |     |     |     | domain.Product.UpdateMinPrice |     |     |                  |
| Created By       |     | CodeCompanion                      |     |     | Executed By        |     |     |     |     |     |                               |     |     |                  |
| Lines of code    |     | 10                                 |     |     | Lack of test cases |     |     |     |     |     | 0                             |     |     |                  |
| Test requirement |     | Validate minimum price calculation |     |     |                    |     |     |     |     |     |                               |     |     |                  |
| Passed           |     | Failed                             |     |     | Untested           |     |     |     |     |     | N                             | A   | B   | Total Test Cases |
| 6                |     | 0                                  |     |     | 0                  |     |     |     |     |     | 6                             | 0   | 0   | 6                |

#### Sheet

| #   |           |                                              |     |                       |     | P-UMP-01-01 | P-UMP-01-02 | P-UMP-01-03 | P-UMP-01-04 | P-UMP-01-05 | P-UMP-01-06 |
| --- | --------- | -------------------------------------------- | --- | --------------------- | --- | ----------- | ----------- | ----------- | ----------- | ----------- | ----------- |
| 1   | Condition | Variant Prices                               |     |                       |     |             |             |             |             |             |             |
| 2   |           |                                              |     | Single variant        |     | O           |             |             |             |             |             |
| 3   |           |                                              |     | Multiple ascending    |     |             | O           |             |             |             |             |
| 4   |           |                                              |     | Multiple descending   |     |             |             | O           |             |             |             |
| 5   |           |                                              |     | Multiple random order |     |             |             |             | O           |             |             |
| 6   |           |                                              |     | All same prices       |     |             |             |             |             | O           |             |
| 7   |           |                                              |     | No variants           |     |             |             |             |             |             | O           |
| 8   | Confirm   | Return                                       |     |                       |     |             |             |             |             |             |             |
| 9   |           | Price                                        |     |                       |     |             |             |             |             |             |             |
| 10  |           |                                              |     | 10000 (min)           |     | O           | O           | O           | O           |             |             |
| 11  |           |                                              |     | 20000 (same)          |     |             |             |             |             | O           |             |
| 12  |           |                                              |     | 0 (empty)             |     |             |             |             |             |             | O           |
| 13  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                       |     | N           | N           | N           | N           | N           | N           |
| 14  |           | Passed/Failed                                |     |                       |     | P           | P           | P           | P           | P           | P           |
| 15  |           | Executed Date                                |     |                       |     | 2025-12-12  | 2025-12-12  | 2025-12-12  | 2025-12-12  | 2025-12-12  | 2025-12-12  |
| 16  |           | Defect ID                                    |     |                       |     |             |             |             |             |             |             |

### `domain.Product.AddVariantImages`

#### Meta

|                  |     |                                         |     |     |                    |     |     |     |     | 0   | 1                               | 2   | 3   | 4                |
| ---------------- | --- | --------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ------------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-AVI-01                                |     |     | Function Name      |     |     |     |     |     | domain.Product.AddVariantImages |     |     |                  |
| Created By       |     | CodeCompanion                           |     |     | Executed By        |     |     |     |     |     |                                 |     |     |                  |
| Lines of code    |     | 13                                      |     |     | Lack of test cases |     |     |     |     |     | 0                               |     |     |                  |
| Test requirement |     | Validate variant image attachment logic |     |     |                    |     |     |     |     |     |                                 |     |     |                  |
| Passed           |     | Failed                                  |     |     | Untested           |     |     |     |     |     | N                               | A   | B   | Total Test Cases |
| 4                |     | 0                                       |     |     | 0                  |     |     |     |     |     | 2                               | 2   | 0   | 4                |

#### Sheet

| #   |           |                                              |     |               |     | P-AVI-01-01 | P-AVI-01-02 | P-AVI-01-03 | P-AVI-01-04 |
| --- | --------- | -------------------------------------------- | --- | ------------- | --- | ----------- | ----------- | ----------- | ----------- |
| 1   | Condition | Variant Existence                            |     |               |     |             |             |             |             |
| 2   |           |                                              |     | Existing      |     | O           |             | O           | O           |
| 3   |           |                                              |     | Non-existent  |     |             | O           |             |             |
| 4   |           | Images Count                                 |     |               |     |             |             |             |             |
| 5   |           |                                              |     | 0 (no images) |     |             |             | O           |             |
| 6   |           |                                              |     | 2 images      |     | O           | O           |             |             |
| 7   |           |                                              |     | 5 images      |     |             |             |             | O           |
| 8   | Confirm   | Return                                       |     |               |     |             |             |             |             |
| 9   |           | Images Count                                 |     |               |     |             |             |             |             |
| 10  |           |                                              |     | 2             |     | O           |             |             |             |
| 11  |           |                                              |     | 0             |     |             |             | O           |             |
| 12  |           |                                              |     | 5             |     |             |             |             | O           |
| 13  |           | Error                                        |     |               |     |             |             |             |             |
| 14  |           |                                              |     | ErrNotFound   |     |             | O           |             |             |
| 15  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |               |     | N           | A           | N           | N           |
| 16  |           | Passed/Failed                                |     |               |     | P           | P           | P           | P           |
| 17  |           | Executed Date                                |     |               |     | 2025-12-12  | 2025-12-12  | 2025-12-12  | 2025-12-12  |
| 18  |           | Defect ID                                    |     |               |     |             |             |             |             |

### `domain.Product.Remove` (Soft Delete)

#### Meta

|                  |     |                                                      |     |     |                    |     |     |     |     | 0   | 1                     | 2   | 3   | 4                |
| ---------------- | --- | ---------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | --------------------- | --- | --- | ---------------- |
| Function Code    |     | P-R-01                                               |     |     | Function Name      |     |     |     |     |     | domain.Product.Remove |     |     |                  |
| Created By       |     | CodeCompanion                                        |     |     | Executed By        |     |     |     |     |     |                       |     |     |                  |
| Lines of code    |     | 10                                                   |     |     | Lack of test cases |     |     |     |     |     | 0                     |     |     |                  |
| Test requirement |     | Validate soft delete cascades to all nested entities |     |     |                    |     |     |     |     |     |                       |     |     |                  |
| Passed           |     | Failed                                               |     |     | Untested           |     |     |     |     |     | N                     | A   | B   | Total Test Cases |
| 1                |     | 0                                                    |     |     | 0                  |     |     |     |     |     | 1                     | 0   | 0   | 1                |

#### Sheet

| #   |           |                                              |     |                                |     | P-R-01-01  |
| --- | --------- | -------------------------------------------- | --- | ------------------------------ | --- | ---------- |
| 1   | Condition | Product Content                              |     |                                |     |            |
| 2   |           |                                              |     | With options, variants, images |     | O          |
| 3   | Confirm   | Return                                       |     |                                |     |            |
| 4   |           | Product DeletedAt                            |     |                                |     |            |
| 5   |           |                                              |     | Not Zero                       |     | O          |
| 6   |           | Product UpdatedAt                            |     |                                |     |            |
| 7   |           |                                              |     | Not Zero                       |     | O          |
| 8   |           | Option DeletedAt                             |     |                                |     |            |
| 9   |           |                                              |     | All Not Zero                   |     | O          |
| 10  |           | OptionValue DeletedAt                        |     |                                |     |            |
| 11  |           |                                              |     | All Not Zero                   |     | O          |
| 12  |           | Variant DeletedAt                            |     |                                |     |            |
| 13  |           |                                              |     | All Not Zero                   |     | O          |
| 14  |           | ProductImage DeletedAt                       |     |                                |     |            |
| 15  |           |                                              |     | All Not Zero                   |     | O          |
| 16  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                                |     | N          |
| 17  |           | Passed/Failed                                |     |                                |     | P          |
| 18  |           | Executed Date                                |     |                                |     | 2025-12-12 |
| 19  |           | Defect ID                                    |     |                                |     |            |

### `domain.Option.GetValueByID`

#### Meta

|                  |     |                                       |     |     |                    |     |     |     |     | 0   | 1                          | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | -------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-OGV-01                              |     |     | Function Name      |     |     |     |     |     | domain.Option.GetValueByID |     |     |                  |
| Created By       |     | CodeCompanion                         |     |     | Executed By        |     |     |     |     |     |                            |     |     |                  |
| Lines of code    |     | 7                                     |     |     | Lack of test cases |     |     |     |     |     | 0                          |     |     |                  |
| Test requirement |     | Validate option value retrieval by ID |     |     |                    |     |     |     |     |     |                            |     |     |                  |
| Passed           |     | Failed                                |     |     | Untested           |     |     |     |     |     | N                          | A   | B   | Total Test Cases |
| 2                |     | 0                                     |     |     | 0                  |     |     |     |     |     | 2                          | 0   | 0   | 2                |

#### Sheet

| #   |           |                                              |     |                 |     | P-OGV-01-01 | P-OGV-01-02 |
| --- | --------- | -------------------------------------------- | --- | --------------- | --- | ----------- | ----------- |
| 1   | Condition | Value ID                                     |     |                 |     |             |             |
| 2   |           |                                              |     | Existing        |     | O           |             |
| 3   |           |                                              |     | Non-existent    |     |             | O           |
| 4   | Confirm   | Return                                       |     |                 |     |             |             |
| 5   |           | OptionValue                                  |     |                 |     |             |             |
| 6   |           |                                              |     | Found (pointer) |     | O           |             |
| 7   |           |                                              |     | Nil             |     |             | O           |
| 8   | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                 |     | N           | N           |
| 9   |           | Passed/Failed                                |     |                 |     | P           | P           |
| 10  |           | Executed Date                                |     |                 |     | 2025-12-12  | 2025-12-12  |
| 11  |           | Defect ID                                    |     |                 |     |             |             |

### `domain.Option.GetValuesByIDs`

#### Meta

|                  |     |                                       |     |     |                    |     |     |     |     | 0   | 1                            | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ---------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-OGVS-01                             |     |     | Function Name      |     |     |     |     |     | domain.Option.GetValuesByIDs |     |     |                  |
| Created By       |     | CodeCompanion                         |     |     | Executed By        |     |     |     |     |     |                              |     |     |                  |
| Lines of code    |     | 12                                    |     |     | Lack of test cases |     |     |     |     |     | 0                            |     |     |                  |
| Test requirement |     | Validate batch option value retrieval |     |     |                    |     |     |     |     |     |                              |     |     |                  |
| Passed           |     | Failed                                |     |     | Untested           |     |     |     |     |     | N                            | A   | B   | Total Test Cases |
| 6                |     | 0                                     |     |     | 0                  |     |     |     |     |     | 6                            | 0   | 0   | 6                |

#### Sheet

| #   |           |                                              |     |                                 |     | P-OGVS-01-01 | P-OGVS-01-02 | P-OGVS-01-03 | P-OGVS-01-04 | P-OGVS-01-05 | P-OGVS-01-06 |
| --- | --------- | -------------------------------------------- | --- | ------------------------------- | --- | ------------ | ------------ | ------------ | ------------ | ------------ | ------------ |
| 1   | Condition | Value ID List                                |     |                                 |     |              |              |              |              |              |              |
| 2   |           |                                              |     | Single existing                 |     | O            |              |              |              |              |              |
| 3   |           |                                              |     | Multiple existing               |     |              | O            |              |              |              |              |
| 4   |           |                                              |     | All values                      |     |              |              | O            |              |              |              |
| 5   |           |                                              |     | All non-existent                |     |              |              |              | O            |              |              |
| 6   |           |                                              |     | Mixed existing and non-existent |     |              |              |              |              | O            |              |
| 7   |           |                                              |     | Empty list                      |     |              |              |              |              |              | O            |
| 8   | Confirm   | Return                                       |     |                                 |     |              |              |              |              |              |              |
| 9   |           | Count                                        |     |                                 |     |              |              |              |              |              |              |
| 10  |           |                                              |     | 1                               |     | O            |              |              |              |              |              |
| 11  |           |                                              |     | 2                               |     |              | O            |              |              |              |              |
| 12  |           |                                              |     | 4                               |     |              |              | O            |              |              |              |
| 13  |           |                                              |     | 0                               |     |              |              |              | O            | O            | O            |
| 14  |           |                                              |     | 1 (mixed)                       |     |              |              |              |              | O            |              |
| 15  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                                 |     | N            | N            | N            | N            | N            | N            |
| 16  |           | Passed/Failed                                |     |                                 |     | P            | P            | P            | P            | P            | P            |
| 17  |           | Executed Date                                |     |                                 |     | 2025-12-12   | 2025-12-12   | 2025-12-12   | 2025-12-12   | 2025-12-12   | 2025-12-12   |
| 18  |           | Defect ID                                    |     |                                 |     |              |              |              |              |              |              |

### `domain.ProductVariant.DecreaseQuantity`

#### Meta

|                  |     |                                                      |     |     |                    |     |     |     |     | 0   | 1                                      | 2   | 3   | 4                |
| ---------------- | --- | ---------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | -------------------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-DQ-01                                              |     |     | Function Name      |     |     |     |     |     | domain.ProductVariant.DecreaseQuantity |     |     |                  |
| Created By       |     | CodeCompanion                                        |     |     | Executed By        |     |     |     |     |     |                                        |     |     |                  |
| Lines of code    |     | 11                                                   |     |     | Lack of test cases |     |     |     |     |     | 0                                      |     |     |                  |
| Test requirement |     | Validate quantity decrease and purchase count update |     |     |                    |     |     |     |     |     |                                        |     |     |                  |
| Passed           |     | Failed                                               |     |     | Untested           |     |     |     |     |     | N                                      | A   | B   | Total Test Cases |
| 6                |     | 0                                                    |     |     | 0                  |     |     |     |     |     | 6                                      | 0   | 0   | 6                |

#### Sheet

| #   |           |                                              |     |                               |     | P-DQ-01-01 | P-DQ-01-02 | P-DQ-01-03 | P-DQ-01-04 | P-DQ-01-05 | P-DQ-01-06 |
| --- | --------- | -------------------------------------------- | --- | ----------------------------- | --- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- |
| 1   | Condition | Initial Quantity                             |     |                               |     |            |            |            |            |            |            |
| 2   |           |                                              |     | 10                            |     | O          | O          | O          | O          | O          |            |
| 3   |           |                                              |     | 0                             |     |            |            |            |            |            | O          |
| 4   |           | Decrease By                                  |     |                               |     |            |            |            |            |            |            |
| 5   |           |                                              |     | 5 (valid)                     |     | O          |            |            |            |            |            |
| 6   |           |                                              |     | 10 (equal)                    |     |            | O          |            |            |            |            |
| 7   |           |                                              |     | 15 (more than available)      |     |            |            | O          |            |            |            |
| 8   |           |                                              |     | 0 (zero)                      |     |            |            |            | O          |            |            |
| 9   |           |                                              |     | -5 (negative)                 |     |            |            |            |            | O          |            |
| 10  |           |                                              |     | 5 (from zero quantity)        |     |            |            |            |            |            | O          |
| 11  | Confirm   | Return                                       |     |                               |     |            |            |            |            |            |            |
| 12  |           | Quantity                                     |     |                               |     |            |            |            |            |            |            |
| 13  |           |                                              |     | 5 (decreased)                 |     | O          |            |            |            |            |            |
| 14  |           |                                              |     | 0 (floored to zero)           |     |            | O          | O          |            |            |            |
| 15  |           |                                              |     | 10 (unchanged)                |     |            |            |            | O          | O          |            |
| 16  |           | PurchaseCount                                |     |                               |     |            |            |            |            |            |            |
| 17  |           |                                              |     | 5 (incremented)               |     | O          |            |            |            |            |            |
| 18  |           |                                              |     | 10 (incremented)              |     |            | O          |            |            |            |            |
| 19  |           |                                              |     | 15 (incremented beyond stock) |     |            |            | O          |            |            |            |
| 20  |           |                                              |     | 0 (unchanged)                 |     |            |            |            | O          | O          |            |
| 21  |           |                                              |     | 5 (from zero)                 |     |            |            |            |            |            | O          |
| 22  |           | UpdatedAt                                    |     |                               |     |            |            |            |            |            |            |
| 23  |           |                                              |     | Changed                       |     | O          | O          | O          |            |            | O          |
| 24  |           |                                              |     | Unchanged                     |     |            |            |            | O          | O          |            |
| 25  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                               |     | N          | N          | N          | N          | N          | N          |
| 26  |           | Passed/Failed                                |     |                               |     | P          | P          | P          | P          | P          | P          |
| 27  |           | Executed Date                                |     |                               |     | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 | 2025-12-12 |
| 28  |           | Defect ID                                    |     |                               |     |            |            |            |            |            |            |

### `domain.Product.AddAttributeIDs`

#### Meta

|                  |     |                                      |     |     |                    |     |     |     |     | 0   | 1                              | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------ | --- | --- | ------------------ | --- | --- | --- | --- | --- | ------------------------------ | --- | --- | ---------------- |
| Function Code    |     | P-AAI-01                             |     |     | Function Name      |     |     |     |     |     | domain.Product.AddAttributeIDs |     |     |                  |
| Created By       |     | CodeCompanion                        |     |     | Executed By        |     |     |     |     |     |                                |     |     |                  |
| Lines of code    |     | 2                                    |     |     | Lack of test cases |     |     |     |     |     | 0                              |     |     |                  |
| Test requirement |     | Validate batch attribute ID addition |     |     |                    |     |     |     |     |     |                                |     |     |                  |
| Passed           |     | Failed                               |     |     | Untested           |     |     |     |     |     | N                              | A   | B   | Total Test Cases |
| 1                |     | 0                                    |     |     | 0                  |     |     |     |     |     | 1                              | 0   | 0   | 1                |

#### Sheet

| #   |           |                                              |     |                      |     | P-AAI-01-01 |
| --- | --------- | -------------------------------------------- | --- | -------------------- | --- | ----------- |
| 1   | Condition | Add Operations                               |     |                      |     |             |
| 2   |           |                                              |     | Single then multiple |     | O           |
| 3   | Confirm   | Return                                       |     |                      |     |             |
| 4   |           | AttributeIDs Count                           |     |                      |     |             |
| 5   |           |                                              |     | 3 (1+2)              |     | O           |
| 6   |           | AttributeIDs Values                          |     |                      |     |             |
| 7   |           |                                              |     | Correct order        |     | O           |
| 8   | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                      |     | N           |
| 9   |           | Passed/Failed                                |     |                      |     | P           |
| 10  |           | Executed Date                                |     |                      |     | 2025-12-12  |
| 11  |           | Defect ID                                    |     |                      |     |             |

### `domain.Product.AddAttributeValueIDs`

#### Meta

|                  |     |                                            |     |     |                    |     |     |     |     | 0   | 1                                   | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------ | --- | --- | ------------------ | --- | --- | --- | --- | --- | ----------------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-AAVI-01                                  |     |     | Function Name      |     |     |     |     |     | domain.Product.AddAttributeValueIDs |     |     |                  |
| Created By       |     | CodeCompanion                              |     |     | Executed By        |     |     |     |     |     |                                     |     |     |                  |
| Lines of code    |     | 2                                          |     |     | Lack of test cases |     |     |     |     |     | 0                                   |     |     |                  |
| Test requirement |     | Validate batch attribute value ID addition |     |     |                    |     |     |     |     |     |                                     |     |     |                  |
| Passed           |     | Failed                                     |     |     | Untested           |     |     |     |     |     | N                                   | A   | B   | Total Test Cases |
| 1                |     | 0                                          |     |     | 0                  |     |     |     |     |     | 1                                   | 0   | 0   | 1                |

#### Sheet

| #   |           |                                              |     |                      |     | P-AAVI-01-01 |
| --- | --------- | -------------------------------------------- | --- | -------------------- | --- | ------------ |
| 1   | Condition | Add Operations                               |     |                      |     |              |
| 2   |           |                                              |     | Single then multiple |     | O            |
| 3   | Confirm   | Return                                       |     |                      |     |              |
| 4   |           | AttributeValueIDs Count                      |     |                      |     |              |
| 5   |           |                                              |     | 2 (1+1)              |     | O            |
| 6   |           | AttributeValueIDs Values                     |     |                      |     |              |
| 7   |           |                                              |     | Correct order        |     | O            |
| 8   | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                      |     | N            |
| 9   |           | Passed/Failed                                |     |                      |     | P            |
| 10  |           | Executed Date                                |     |                      |     | 2025-12-12   |
| 11  |           | Defect ID                                    |     |                      |     |              |

### `domain.Product.AddOptions`

#### Meta

|                  |     |                                |     |     |                    |     |     |     |     | 0   | 1                         | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------ | --- | --- | ------------------ | --- | --- | --- | --- | --- | ------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-AO-01                        |     |     | Function Name      |     |     |     |     |     | domain.Product.AddOptions |     |     |                  |
| Created By       |     | CodeCompanion                  |     |     | Executed By        |     |     |     |     |     |                           |     |     |                  |
| Lines of code    |     | 2                              |     |     | Lack of test cases |     |     |     |     |     | 0                         |     |     |                  |
| Test requirement |     | Validate batch option addition |     |     |                    |     |     |     |     |     |                           |     |     |                  |
| Passed           |     | Failed                         |     |     | Untested           |     |     |     |     |     | N                         | A   | B   | Total Test Cases |
| 1                |     | 0                              |     |     | 0                  |     |     |     |     |     | 1                         | 0   | 0   | 1                |

#### Sheet

| #   |           |                                              |     |                      |     | P-AO-01-01 |
| --- | --------- | -------------------------------------------- | --- | -------------------- | --- | ---------- |
| 1   | Condition | Add Operations                               |     |                      |     |            |
| 2   |           |                                              |     | Single then multiple |     | O          |
| 3   | Confirm   | Return                                       |     |                      |     |            |
| 4   |           | Options Count                                |     |                      |     |            |
| 5   |           |                                              |     | 2 (1+1)              |     | O          |
| 6   |           | Options Values                               |     |                      |     |            |
| 7   |           |                                              |     | Correct order        |     | O          |
| 8   | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                      |     | N          |
| 9   |           | Passed/Failed                                |     |                      |     | P          |
| 10  |           | Executed Date                                |     |                      |     | 2025-12-12 |
| 11  |           | Defect ID                                    |     |                      |     |            |

### `domain.Product.AddVariants`

#### Meta

|                  |     |                                 |     |     |                    |     |     |     |     | 0   | 1                          | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | -------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-AV-01                         |     |     | Function Name      |     |     |     |     |     | domain.Product.AddVariants |     |     |                  |
| Created By       |     | CodeCompanion                   |     |     | Executed By        |     |     |     |     |     |                            |     |     |                  |
| Lines of code    |     | 2                               |     |     | Lack of test cases |     |     |     |     |     | 0                          |     |     |                  |
| Test requirement |     | Validate batch variant addition |     |     |                    |     |     |     |     |     |                            |     |     |                  |
| Passed           |     | Failed                          |     |     | Untested           |     |     |     |     |     | N                          | A   | B   | Total Test Cases |
| 1                |     | 0                               |     |     | 0                  |     |     |     |     |     | 1                          | 0   | 0   | 1                |

#### Sheet

| #   |           |                                              |     |                      |     | P-AV-01-01 |
| --- | --------- | -------------------------------------------- | --- | -------------------- | --- | ---------- |
| 1   | Condition | Add Operations                               |     |                      |     |            |
| 2   |           |                                              |     | Single then multiple |     | O          |
| 3   | Confirm   | Return                                       |     |                      |     |            |
| 4   |           | Variants Count                               |     |                      |     |            |
| 5   |           |                                              |     | 2 (1+1)              |     | O          |
| 6   |           | Variants Values                              |     |                      |     |            |
| 7   |           |                                              |     | Correct order        |     | O          |
| 8   | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                      |     | N          |
| 9   |           | Passed/Failed                                |     |                      |     | P          |
| 10  |           | Executed Date                                |     |                      |     | 2025-12-12 |
| 11  |           | Defect ID                                    |     |                      |     |            |

### `domain.Product.AddImages`

#### Meta

|                  |     |                               |     |     |                    |     |     |     |     | 0   | 1                        | 2   | 3   | 4                |
| ---------------- | --- | ----------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ------------------------ | --- | --- | ---------------- |
| Function Code    |     | P-AI-01                       |     |     | Function Name      |     |     |     |     |     | domain.Product.AddImages |     |     |                  |
| Created By       |     | CodeCompanion                 |     |     | Executed By        |     |     |     |     |     |                          |     |     |                  |
| Lines of code    |     | 2                             |     |     | Lack of test cases |     |     |     |     |     | 0                        |     |     |                  |
| Test requirement |     | Validate batch image addition |     |     |                    |     |     |     |     |     |                          |     |     |                  |
| Passed           |     | Failed                        |     |     | Untested           |     |     |     |     |     | N                        | A   | B   | Total Test Cases |
| 1                |     | 0                             |     |     | 0                  |     |     |     |     |     | 1                        | 0   | 0   | 1                |

#### Sheet

| #   |           |                                              |     |                      |     | P-AI-01-01 |
| --- | --------- | -------------------------------------------- | --- | -------------------- | --- | ---------- |
| 1   | Condition | Add Operations                               |     |                      |     |            |
| 2   |           |                                              |     | Single then multiple |     | O          |
| 3   | Confirm   | Return                                       |     |                      |     |            |
| 4   |           | Images Count                                 |     |                      |     |            |
| 5   |           |                                              |     | 2 (1+1)              |     | O          |
| 6   |           | Images Values                                |     |                      |     |            |
| 7   |           |                                              |     | Correct order        |     | O          |
| 8   | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                      |     | N          |
| 9   |           | Passed/Failed                                |     |                      |     | P          |
| 10  |           | Executed Date                                |     |                      |     | 2025-12-12 |
| 11  |           | Defect ID                                    |     |                      |     |            |

### `domain.Option.AddOptionValues`

#### Meta

|                  |     |                                      |     |     |                    |     |     |     |     | 0   | 1                             | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------ | --- | --- | ------------------ | --- | --- | --- | --- | --- | ----------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-OAV-01                             |     |     | Function Name      |     |     |     |     |     | domain.Option.AddOptionValues |     |     |                  |
| Created By       |     | CodeCompanion                        |     |     | Executed By        |     |     |     |     |     |                               |     |     |                  |
| Lines of code    |     | 2                                    |     |     | Lack of test cases |     |     |     |     |     | 0                             |     |     |                  |
| Test requirement |     | Validate batch option value addition |     |     |                    |     |     |     |     |     |                               |     |     |                  |
| Passed           |     | Failed                               |     |     | Untested           |     |     |     |     |     | N                             | A   | B   | Total Test Cases |
| 1                |     | 0                                    |     |     | 0                  |     |     |     |     |     | 1                             | 0   | 0   | 1                |

#### Sheet

| #   |           |                                              |     |                      |     | P-OAV-01-01 |
| --- | --------- | -------------------------------------------- | --- | -------------------- | --- | ----------- |
| 1   | Condition | Add Operations                               |     |                      |     |             |
| 2   |           |                                              |     | Single then multiple |     | O           |
| 3   | Confirm   | Return                                       |     |                      |     |             |
| 4   |           | Values Count                                 |     |                      |     |             |
| 5   |           |                                              |     | 3 (1+2)              |     | O           |
| 6   |           | Values Data                                  |     |                      |     |             |
| 7   |           |                                              |     | Correct order        |     | O           |
| 8   | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                      |     | N           |
| 9   |           | Passed/Failed                                |     |                      |     | P           |
| 10  |           | Executed Date                                |     |                      |     | 2025-12-12  |
| 11  |           | Defect ID                                    |     |                      |     |             |

### `domain.Option.Remove`

#### Meta

|                  |     |                                                        |     |     |                    |     |     |     |     | 0   | 1                    | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------------------ | --- | --- | ------------------ | --- | --- | --- | --- | --- | -------------------- | --- | --- | ---------------- |
| Function Code    |     | P-OR-01                                                |     |     | Function Name      |     |     |     |     |     | domain.Option.Remove |     |     |                  |
| Created By       |     | CodeCompanion                                          |     |     | Executed By        |     |     |     |     |     |                      |     |     |                  |
| Lines of code    |     | 6                                                      |     |     | Lack of test cases |     |     |     |     |     | 0                    |     |     |                  |
| Test requirement |     | Validate soft delete on Option and nested OptionValues |     |     |                    |     |     |     |     |     |                      |     |     |                  |
| Passed           |     | Failed                                                 |     |     | Untested           |     |     |     |     |     | N                    | A   | B   | Total Test Cases |
| 1                |     | 0                                                      |     |     | 0                  |     |     |     |     |     | 1                    | 0   | 0   | 1                |

#### Sheet

| #   |           |                                              |     |                             |     | P-OR-01-01 |
| --- | --------- | -------------------------------------------- | --- | --------------------------- | --- | ---------- |
| 1   | Condition | Option Content                               |     |                             |     |            |
| 2   |           |                                              |     | With multiple option values |     | O          |
| 3   | Confirm   | Return                                       |     |                             |     |            |
| 4   |           | Option DeletedAt                             |     |                             |     |            |
| 5   |           |                                              |     | Not Zero                    |     | O          |
| 6   |           | OptionValue DeletedAt                        |     |                             |     |            |
| 7   |           |                                              |     | All Not Zero                |     | O          |
| 8   | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                             |     | N          |
| 9   |           | Passed/Failed                                |     |                             |     | P          |
| 10  |           | Executed Date                                |     |                             |     | 2025-12-12 |
| 11  |           | Defect ID                                    |     |                             |     |            |

### `domain.OptionValue.Remove`

#### Meta

|                  |     |                                     |     |     |                    |     |     |     |     | 0   | 1                         | 2   | 3   | 4                |
| ---------------- | --- | ----------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-OVR-01                            |     |     | Function Name      |     |     |     |     |     | domain.OptionValue.Remove |     |     |                  |
| Created By       |     | CodeCompanion                       |     |     | Executed By        |     |     |     |     |     |                           |     |     |                  |
| Lines of code    |     | 3                                   |     |     | Lack of test cases |     |     |     |     |     | 0                         |     |     |                  |
| Test requirement |     | Validate soft delete on OptionValue |     |     |                    |     |     |     |     |     |                           |     |     |                  |
| Passed           |     | Failed                              |     |     | Untested           |     |     |     |     |     | N                         | A   | B   | Total Test Cases |
| 1                |     | 0                                   |     |     | 0                  |     |     |     |     |     | 1                         | 0   | 0   | 1                |

#### Sheet

| #   |           |                                              |     |          |     | P-OVR-01-01 |
| --- | --------- | -------------------------------------------- | --- | -------- | --- | ----------- |
| 1   | Condition | OptionValue                                  |     |          |     |             |
| 2   |           |                                              |     | Valid    |     | O           |
| 3   | Confirm   | Return                                       |     |          |     |             |
| 4   |           | DeletedAt                                    |     |          |     |             |
| 5   |           |                                              |     | Not Zero |     | O           |
| 6   | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |          |     | N           |
| 7   |           | Passed/Failed                                |     |          |     | P           |
| 8   |           | Executed Date                                |     |          |     | 2025-12-12  |
| 9   |           | Defect ID                                    |     |          |     |             |

### `domain.ProductVariant.Remove`

#### Meta

|                  |     |                                        |     |     |                    |     |     |     |     | 0   | 1                            | 2   | 3   | 4                |
| ---------------- | --- | -------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ---------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-VR-01                                |     |     | Function Name      |     |     |     |     |     | domain.ProductVariant.Remove |     |     |                  |
| Created By       |     | CodeCompanion                          |     |     | Executed By        |     |     |     |     |     |                              |     |     |                  |
| Lines of code    |     | 4                                      |     |     | Lack of test cases |     |     |     |     |     | 0                            |     |     |                  |
| Test requirement |     | Validate soft delete on ProductVariant |     |     |                    |     |     |     |     |     |                              |     |     |                  |
| Passed           |     | Failed                                 |     |     | Untested           |     |     |     |     |     | N                            | A   | B   | Total Test Cases |
| 1                |     | 0                                      |     |     | 0                  |     |     |     |     |     | 1                            | 0   | 0   | 1                |

#### Sheet

| #   |           |                                              |     |          |     | P-VR-01-01 |
| --- | --------- | -------------------------------------------- | --- | -------- | --- | ---------- |
| 1   | Condition | ProductVariant                               |     |          |     |            |
| 2   |           |                                              |     | Valid    |     | O          |
| 3   | Confirm   | Return                                       |     |          |     |            |
| 4   |           | DeletedAt                                    |     |          |     |            |
| 5   |           |                                              |     | Not Zero |     | O          |
| 6   |           | UpdatedAt                                    |     |          |     |            |
| 7   |           |                                              |     | Not Zero |     | O          |
| 8   | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |          |     | N          |
| 9   |           | Passed/Failed                                |     |          |     | P          |
| 10  |           | Executed Date                                |     |          |     | 2025-12-12 |
| 11  |           | Defect ID                                    |     |          |     |            |

### `domain.ProductImage.Remove`

#### Meta

|                  |     |                                      |     |     |                    |     |     |     |     | 0   | 1                          | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------ | --- | --- | ------------------ | --- | --- | --- | --- | --- | -------------------------- | --- | --- | ---------------- |
| Function Code    |     | P-IR-01                              |     |     | Function Name      |     |     |     |     |     | domain.ProductImage.Remove |     |     |                  |
| Created By       |     | CodeCompanion                        |     |     | Executed By        |     |     |     |     |     |                            |     |     |                  |
| Lines of code    |     | 3                                    |     |     | Lack of test cases |     |     |     |     |     | 0                          |     |     |                  |
| Test requirement |     | Validate soft delete on ProductImage |     |     |                    |     |     |     |     |     |                            |     |     |                  |
| Passed           |     | Failed                               |     |     | Untested           |     |     |     |     |     | N                          | A   | B   | Total Test Cases |
| 1                |     | 0                                    |     |     | 0                  |     |     |     |     |     | 1                          | 0   | 0   | 1                |

#### Sheet

| #   |           |                                              |     |          |     | P-IR-01-01 |
| --- | --------- | -------------------------------------------- | --- | -------- | --- | ---------- |
| 1   | Condition | ProductImage                                 |     |          |     |            |
| 2   |           |                                              |     | Valid    |     | O          |
| 3   | Confirm   | Return                                       |     |          |     |            |
| 4   |           | DeletedAt                                    |     |          |     |            |
| 5   |           |                                              |     | Not Zero |     | O          |
| 6   | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |          |     | N          |
| 7   |           | Passed/Failed                                |     |          |     | P          |
| 8   |           | Executed Date                                |     |          |     | 2025-12-12 |
| 9   |           | Defect ID                                    |     |          |     |            |
