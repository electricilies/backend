### `domain.NewCartItem`

#### Meta

|                  |     |                                                         |     |     |                    |     |     |     |     | 0   | 1                |     | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ---------------- | --- | --- | --- | ---------------- |
| Function Code    |     | CR-NCI-01                                               |     |     | Function Name      |     |     |     |     |     | domain.NewCartItem|     |     |     |                  |
| Created By       |     | CodeCompanion                                           |     |     | Executed By        |     |     |     |     |     |                  |     |     |     |                  |
| Lines of code    |     | 14                                                      |     |     | Lack of test cases |     |     |     |     |     | 0                |     |     |     |                  |
| Test requirement |     | Validate CartItem quantity boundary values              |     |     |                    |     |     |     |     |     |                  |     |     |     |                  |
| Passed           |     | Failed                                                  |     |     | Untested           |     |     |     |     |     | N                | A   | B   |     | Total Test Cases |
| 8                |     | 0                                                       |     |     | 0                  |     |     |     |     |     | 1                | 2   | 5   |     | 8                |

#### Sheet

| #   |           |                                              |     |                                                                                     |     | CR-NCI-01-01 | CR-NCI-01-02 | CR-NCI-01-03 | CR-NCI-01-04 | CR-NCI-01-05 | CR-NCI-01-06 | CR-NCI-01-07 | CR-NCI-01-08 |
| --- | --------- | -------------------------------------------- | --- | ----------------------------------------------------------------------------------- | --- | ------------ | ------------ | ------------ | ------------ | ------------ | ------------ | ------------ | ------------ |
| 1   | Condition | Quantity                                     |     |                                                                                     |     |              |              |              |              |              |              |              |              |
| 2   |           |                                              |     | -1 (negative)                                                                       |     |              |              |              |              |              |              |              | O            |
| 3   |           |                                              |     | 0 (invalid)                                                                         |     | O            |              |              |              |              |              |              |              |
| 4   |           |                                              |     | 1 (min)                                                                             |     |              | O            |              |              |              |              |              |              |
| 5   |           |                                              |     | 2 (min + 1)                                                                         |     |              |              | O            |              |              |              |              |              |
| 6   |           |                                              |     | 50                                                                                  |     |              |              |              | O            |              |              |              |              |
| 7   |           |                                              |     | 99 (max - 1)                                                                        |     |              |              |              |              | O            |              |              |              |
| 8   |           |                                              |     | 100 (max)                                                                           |     |              |              |              |              |              | O            |              |              |
| 9   |           |                                              |     | 101 (max + 1)                                                                       |     |              |              |              |              |              |              | O            |              |
| 10  | Confirm   | Return                                       |     |                                                                                     |     |              |              |              |              |              |              |              |              |
| 11  |           | Quantity                                     |     |                                                                                     |     |              |              |              |              |              |              |              |              |
| 12  |           |                                              |     | Input Quantity                                                                      |     | O            | O            | O            | O            | O            | O            | O            | O            |
| 13  |           | ID                                           |     |                                                                                     |     |              |              |              |              |              |              |              |              |
| 14  |           |                                              |     | Not Nil                                                                             |     | O            | O            | O            | O            | O            | O            | O            | O            |
| 15  |           | ProductID                                    |     |                                                                                     |     |              |              |              |              |              |              |              |              |
| 16  |           |                                              |     | Input ProductID                                                                     |     | O            | O            | O            | O            | O            | O            | O            | O            |
| 17  |           | ProductVariantID                             |     |                                                                                     |     |              |              |              |              |              |              |              |              |
| 18  |           |                                              |     | Input ProductVariantID                                                              |     | O            | O            | O            | O            | O            | O            | O            | O            |
| 19  |           | Error                                        |     |                                                                                     |     |              |              |              |              |              |              |              |              |
| 20  |           |                                              |     | Key: 'CartItem.Quantity' Error:Field validation for 'Quantity' failed on 'gt' tag   |     | O            |              |              |              |              |              |              | O            |
| 21  |           |                                              |     | Key: 'CartItem.Quantity' Error:Field validation for 'Quantity' failed on 'lte' tag  |     |              |              |              |              |              |              | O            |              |
| 22  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                                                                                     |     | A            | B            | B            | B            | B            | B            | B            | A            |
| 23  |           | Passed/Failed                                |     |                                                                                     |     | P            | P            | P            | P            | P            | P            | P            | P            |
| 24  |           | Executed Date                                |     |                                                                                     |     | 2025-12-10   | 2025-12-10   | 2025-12-10   | 2025-12-10   | 2025-12-10   | 2025-12-10   | 2025-12-10   | 2025-12-10   |
| 25  |           | Defect ID                                    |     |                                                                                     |     |              |              |              |              |              |              |              |              |

### `domain.Cart.UpsertItem`

#### Meta

|                  |     |                                                       |     |     |                    |     |     |     |     | 0   | 1                    | 2   | 3   | 4                |
| ---------------- | --- | ----------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | -------------------- | --- | --- | ---------------- |
| Function Code    |     | CR-UI-01                                              |     |     | Function Name      |     |     |     |     |     | domain.Cart.UpsertItem|     |     |                  |
| Created By       |     | CodeCompanion                                         |     |     | Executed By        |     |     |     |     |     |                      |     |     |                  |
| Lines of code    |     | 10                                                    |     |     | Lack of test cases |     |     |     |     |     | 0                    |     |     |                  |
| Test requirement |     | Validate cart item upsert logic (add/update quantity) |     |     |                    |     |     |     |     |     |                      |     |     |                  |
| Passed           |     | Failed                                                |     |     | Untested           |     |     |     |     |     | N                    | A   | B   | Total Test Cases |
| 3                |     | 0                                                     |     |     | 0                  |     |     |     |     |     | 3                    | 0   | 0   | 3                |

#### Sheet

| #   |           |                                              |     |                               |     | CR-UI-01-01 | CR-UI-01-02 | CR-UI-01-03 |
| --- | --------- | -------------------------------------------- | --- | ----------------------------- | --- | ----------- | ----------- | ----------- |
| 1   | Condition | Existing Items                               |     |                               |     |             |             |             |
| 2   |           |                                              |     | 0 (empty cart)                |     | O           |             |             |
| 3   |           |                                              |     | 1 (same ProductVariantID)     |     |             | O           |             |
| 4   |           |                                              |     | 1 (different ProductVariantID)|     |             |             | O           |
| 5   |           | New Item Quantity                            |     |                               |     |             |             |             |
| 6   |           |                                              |     | 5                             |     | O           |             |             |
| 7   |           |                                              |     | 2                             |     |             | O           | O           |
| 8   | Confirm   | Return                                       |     |                               |     |             |             |             |
| 9   |           | Items Length                                 |     |                               |     |             |             |             |
| 10  |           |                                              |     | 1                             |     | O           | O           |             |
| 11  |           |                                              |     | 2                             |     |             |             | O           |
| 12  |           | Quantity                                     |     |                               |     |             |             |             |
| 13  |           |                                              |     | 5                             |     | O           | O           |             |
| 14  |           |                                              |     | 2                             |     |             |             | O           |
| 15  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                               |     | N           | N           | N           |
| 16  |           | Passed/Failed                                |     |                               |     | P           | P           | P           |
| 17  |           | Executed Date                                |     |                               |     | 2025-12-10  | 2025-12-10  | 2025-12-10  |
| 18  |           | Defect ID                                    |     |                               |     |             |             |             |

### `domain.Cart.UpdateItem`

#### Meta

|                  |     |                                                          |     |     |                    |     |     |     |     | 0   | 1                     | 2   | 3   | 4                |
| ---------------- | --- | -------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | --------------------- | --- | --- | ---------------- |
| Function Code    |     | CR-UDI-01                                                |     |     | Function Name      |     |     |     |     |     | domain.Cart.UpdateItem|     |     |                  |
| Created By       |     | CodeCompanion                                            |     |     | Executed By        |     |     |     |     |     |                       |     |     |                  |
| Lines of code    |     | 12                                                       |     |     | Lack of test cases |     |     |     |     |     | 0                     |     |     |                  |
| Test requirement |     | Validate cart item update and removal logic              |     |     |                    |     |     |     |     |     |                       |     |     |                  |
| Passed           |     | Failed                                                   |     |     | Untested           |     |     |     |     |     | N                     | A   | B   | Total Test Cases |
| 3                |     | 0                                                        |     |     | 0                  |     |     |     |     |     | 2                     | 0   | 1   | 3                |

#### Sheet

| #   |           |                                              |     |                             |     | CR-UDI-01-01 | CR-UDI-01-02 | CR-UDI-01-03 |
| --- | --------- | -------------------------------------------- | --- | --------------------------- | --- | ------------ | ------------ | ------------ |
| 1   | Condition | Initial Quantity                             |     |                             |     |              |              |              |
| 2   |           |                                              |     | 5                           |     | O            | O            |              |
| 3   |           |                                              |     | 10                          |     |              |              | O            |
| 4   |           | Update Quantity                              |     |                             |     |              |              |              |
| 5   |           |                                              |     | 0 (removes item)            |     |              | O            |              |
| 6   |           |                                              |     | 1                           |     |              |              | O            |
| 7   |           |                                              |     | 10                          |     | O            |              |              |
| 8   | Confirm   | Return                                       |     |                             |     |              |              |              |
| 9   |           | Items Length                                 |     |                             |     |              |              |              |
| 10  |           |                                              |     | 0                           |     |              | O            |              |
| 11  |           |                                              |     | 1                           |     | O            |              | O            |
| 12  |           | Quantity                                     |     |                             |     |              |              |              |
| 13  |           |                                              |     | 1                           |     |              |              | O            |
| 14  |           |                                              |     | 10                          |     | O            |              |              |
| 15  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                             |     | N            | B            | N            |
| 16  |           | Passed/Failed                                |     |                             |     | P            | P            | P            |
| 17  |           | Executed Date                                |     |                             |     | 2025-12-10   | 2025-12-10   | 2025-12-10   |
| 18  |           | Defect ID                                    |     |                             |     |              |              |              |

### `domain.Cart.RemoveItem`

#### Meta

|                  |     |                                            |     |     |                    |     |     |     |     | 0   | 1                     | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------ | --- | --- | ------------------ | --- | --- | --- | --- | --- | --------------------- | --- | --- | ---------------- |
| Function Code    |     | CR-RI-01                                   |     |     | Function Name      |     |     |     |     |     | domain.Cart.RemoveItem|     |     |                  |
| Created By       |     | CodeCompanion                              |     |     | Executed By        |     |     |     |     |     |                       |     |     |                  |
| Lines of code    |     | 7                                          |     |     | Lack of test cases |     |     |     |     |     | 0                     |     |     |                  |
| Test requirement |     | Validate cart item removal logic           |     |     |                    |     |     |     |     |     |                       |     |     |                  |
| Passed           |     | Failed                                     |     |     | Untested           |     |     |     |     |     | N                     | A   | B   | Total Test Cases |
| 3                |     | 0                                          |     |     | 0                  |     |     |     |     |     | 3                     | 0   | 0   | 3                |

#### Sheet

| #   |           |                                              |     |                  |     | CR-RI-01-01 | CR-RI-01-02 | CR-RI-01-03 |
| --- | --------- | -------------------------------------------- | --- | ---------------- | --- | ----------- | ----------- | ----------- |
| 1   | Condition | Initial Items Count                          |     |                  |     |             |             |             |
| 2   |           |                                              |     | 1                |     |             |             | O           |
| 3   |           |                                              |     | 2                |     |             | O           |             |
| 4   |           |                                              |     | 3                |     | O           |             |             |
| 5   |           | Item ID                                      |     |                  |     |             |             |             |
| 6   |           |                                              |     | Existing         |     | O           |             | O           |
| 7   |           |                                              |     | Non-existent     |     |             | O           |             |
| 8   | Confirm   | Return                                       |     |                  |     |             |             |             |
| 9   |           | Items Length                                 |     |                  |     |             |             |             |
| 10  |           |                                              |     | 0                |     |             |             | O           |
| 11  |           |                                              |     | 2                |     | O           | O           |             |
| 12  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                  |     | N           | N           | N           |
| 13  |           | Passed/Failed                                |     |                  |     | P           | P           | P           |
| 14  |           | Executed Date                                |     |                  |     | 2025-12-10  | 2025-12-10  | 2025-12-10  |
| 15  |           | Defect ID                                    |     |                  |     |             |             |             |

### `domain.Cart.ClearItems`

#### Meta

|                  |     |                                          |     |     |                    |     |     |     |     | 0   | 1                     | 2   | 3   | 4                |
| ---------------- | --- | ---------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | --------------------- | --- | --- | ---------------- |
| Function Code    |     | CR-CI-01                                 |     |     | Function Name      |     |     |     |     |     | domain.Cart.ClearItems|     |     |                  |
| Created By       |     | CodeCompanion                            |     |     | Executed By        |     |     |     |     |     |                       |     |     |                  |
| Lines of code    |     | 3                                        |     |     | Lack of test cases |     |     |     |     |     | 0                     |     |     |                  |
| Test requirement |     | Validate cart clear items logic          |     |     |                    |     |     |     |     |     |                       |     |     |                  |
| Passed           |     | Failed                                   |     |     | Untested           |     |     |     |     |     | N                     | A   | B   | Total Test Cases |
| 2                |     | 0                                        |     |     | 0                  |     |     |     |     |     | 2                     | 0   | 0   | 2                |

#### Sheet

| #   |           |                                              |     |                    |     | CR-CI-01-01 | CR-CI-01-02 |
| --- | --------- | -------------------------------------------- | --- | ------------------ | --- | ----------- | ----------- |
| 1   | Condition | Initial Items Count                          |     |                    |     |             |             |
| 2   |           |                                              |     | 0 (empty)          |     |             | O           |
| 3   |           |                                              |     | 5                  |     | O           |             |
| 4   | Confirm   | Return                                       |     |                    |     |             |             |
| 5   |           | Items Length                                 |     |                    |     |             |             |
| 6   |           |                                              |     | 0                  |     | O           | O           |
| 7   |           | Items Not Nil                                |     |                    |     |             |             |
| 8   |           |                                              |     | true               |     | O           | O           |
| 9   | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                    |     | N           | N           |
| 10  |           | Passed/Failed                                |     |                    |     | P           | P           |
| 11  |           | Executed Date                                |     |                    |     | 2025-12-10  | 2025-12-10  |
| 12  |           | Defect ID                                    |     |                    |     |             |             |
