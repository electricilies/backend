### `domain.NewOrderItem`

#### Meta

|                  |     |                                                          |     |     |                    |     |     |     |     | 0   | 1                  | 2   | 3   | 4                |
| ---------------- | --- | -------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ------------------ | --- | --- | ---------------- |
| Function Code    |     | O-NOI-01                                                 |     |     | Function Name      |     |     |     |     |     | domain.NewOrderItem|     |     |                  |
| Created By       |     | CodeCompanion                                            |     |     | Executed By        |     |     |     |     |     |                    |     |     |                  |
| Lines of code    |     | 14                                                       |     |     | Lack of test cases |     |     |     |     |     | 0                  |     |     |                  |
| Test requirement |     | Validate OrderItem quantity and price boundary values    |     |     |                    |     |     |     |     |     |                    |     |     |                  |
| Passed           |     | Failed                                                   |     |     | Untested           |     |     |     |     |     | N                  | A   | B   | Total Test Cases |
| 11               |     | 0                                                        |     |     | 0                  |     |     |     |     |     | 1                  | 3   | 7   | 11               |

#### Sheet

| #   |           |                                              |     |                                                                                      |     | O-NOI-01-01 | O-NOI-01-02 | O-NOI-01-03 | O-NOI-01-04 | O-NOI-01-05 | O-NOI-01-06 | O-NOI-01-07 | O-NOI-01-08 | O-NOI-01-09 | O-NOI-01-10 | O-NOI-01-11 |
| --- | --------- | -------------------------------------------- | --- | ------------------------------------------------------------------------------------ | --- | ----------- | ----------- | ----------- | ----------- | ----------- | ----------- | ----------- | ----------- | ----------- | ----------- | ----------- |
| 1   | Condition | Quantity                                     |     |                                                                                      |     |             |             |             |             |             |             |             |             |             |             |             |
| 2   |           |                                              |     | -1 (negative)                                                                        |     |             |             |             |             |             |             | O           |             |             |             |             |
| 3   |           |                                              |     | 0 (invalid)                                                                          |     | O           |             |             |             |             |             |             |             |             |             |             |
| 4   |           |                                              |     | 1 (min)                                                                              |     |             | O           |             |             |             |             |             |             | O           |             | O           |
| 5   |           |                                              |     | 2 (min + 1)                                                                          |     |             |             | O           |             |             |             |             |             |             |             |             |
| 6   |           |                                              |     | 5                                                                                    |     |             |             |             |             |             |             |             |             |             |             |             |
| 7   |           |                                              |     | 50                                                                                   |     |             |             |             | O           |             |             |             |             |             |             |             |
| 8   |           |                                              |     | 99 (max)                                                                             |     |             |             |             |             | O           |             |             |             |             |             |             |
| 9   |           |                                              |     | 100 (max + 1, invalid)                                                               |     |             |             |             |             |             | O           |             |             |             |             |             |
| 10  |           | Price                                        |     |                                                                                      |     |             |             |             |             |             |             |             |             |             |             |             |
| 11  |           |                                              |     | -1 (negative)                                                                        |     |             |             |             |             |             |             |             |             |             | O           |             |
| 12  |           |                                              |     | 0 (invalid)                                                                          |     |             |             |             |             |             |             |             | O           |             |             |             |
| 13  |           |                                              |     | 1 (min)                                                                              |     |             |             |             |             |             |             |             |             | O           |             |             |
| 14  |           |                                              |     | 1000                                                                                 |     | O           | O           | O           | O           | O           | O           | O           |             |             |             |             |
| 15  |           |                                              |     | 10000                                                                                |     |             |             |             |             |             |             |             |             |             |             | O           |
| 16  | Confirm   | Return                                       |     |                                                                                      |     |             |             |             |             |             |             |             |             |             |             |             |
| 17  |           | Quantity                                     |     |                                                                                      |     |             |             |             |             |             |             |             |             |             |             |             |
| 18  |           |                                              |     | Input Quantity                                                                       |     | O           | O           | O           | O           | O           | O           | O           | O           | O           | O           | O           |
| 19  |           | Price                                        |     |                                                                                      |     |             |             |             |             |             |             |             |             |             |             |             |
| 20  |           |                                              |     | Input Price                                                                          |     | O           | O           | O           | O           | O           | O           | O           | O           | O           | O           | O           |
| 21  |           | ID                                           |     |                                                                                      |     |             |             |             |             |             |             |             |             |             |             |             |
| 22  |           |                                              |     | Not Nil                                                                              |     | O           | O           | O           | O           | O           | O           | O           | O           | O           | O           | O           |
| 23  |           | Error                                        |     |                                                                                      |     |             |             |             |             |             |             |             |             |             |             |             |
| 24  |           |                                              |     | Key: 'OrderItem.Quantity' Error:Field validation for 'Quantity' failed on 'gt' tag   |     | O           |             |             |             |             |             | O           |             |             |             |             |
| 25  |           |                                              |     | Key: 'OrderItem.Quantity' Error:Field validation for 'Quantity' failed on 'lt' tag   |     |             |             |             |             |             | O           |             |             |             |             |             |
| 26  |           |                                              |     | Key: 'OrderItem.Price' Error:Field validation for 'Price' failed on 'gt' tag         |     |             |             |             |             |             |             |             | O           |             | O           |             |
| 27  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                                                                                      |     | A           | B           | B           | B           | B           | B           | A           | A           | B           | A           | N           |
| 28  |           | Passed/Failed                                |     |                                                                                      |     | P           | P           | P           | P           | P           | P           | P           | P           | P           | P           | P           |
| 29  |           | Executed Date                                |     |                                                                                      |     | 2025-12-10  | 2025-12-10  | 2025-12-10  | 2025-12-10  | 2025-12-10  | 2025-12-10  | 2025-12-10  | 2025-12-10  | 2025-12-10  | 2025-12-10  | 2025-12-10  |
| 30  |           | Defect ID                                    |     |                                                                                      |     |             |             |             |             |             |             |             |             |             |             |             |

### `domain.NewOrder`

#### Meta

|                  |     |                                                              |     |     |                    |     |     |     |     | 0   | 1              | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------------------------ | --- | --- | ------------------ | --- | --- | --- | --- | --- | -------------- | --- | --- | ---------------- |
| Function Code    |     | O-NO-01                                                      |     |     | Function Name      |     |     |     |     |     | domain.NewOrder|     |     |                  |
| Created By       |     | CodeCompanion                                                |     |     | Executed By        |     |     |     |     |     |                |     |     |                  |
| Lines of code    |     | 22                                                           |     |     | Lack of test cases |     |     |     |     |     | 0              |     |     |                  |
| Test requirement |     | Validate Order creation and total amount calculation         |     |     |                    |     |     |     |     |     |                |     |     |                  |
| Passed           |     | Failed                                                       |     |     | Untested           |     |     |     |     |     | N              | A   | B   | Total Test Cases |
| 5                |     | 0                                                            |     |     | 0                  |     |     |     |     |     | 4              | 1   | 0   | 5                |

#### Sheet

| #   |           |                                              |     |                                |     | O-NO-01-01 | O-NO-01-02 | O-NO-01-03 | O-NO-01-04 | O-NO-01-05 |
| --- | --------- | -------------------------------------------- | --- | ------------------------------ | --- | ---------- | ---------- | ---------- | ---------- | ---------- |
| 1   | Condition | Items Count                                  |     |                                |     |            |            |            |            |            |
| 2   |           |                                              |     | 0 (empty)                      |     | O          |            |            |            |            |
| 3   |           |                                              |     | 1 (Qty:2, Price:100)           |     |            | O          |            |            |            |
| 4   |           |                                              |     | 2 (Qty:2/3, Price:100/200)     |     |            |            | O          |            |            |
| 5   |           |                                              |     | 1 (Qty:10, Price:5000)         |     |            |            |            | O          |            |
| 6   |           |                                              |     | 3 (Qty:1/2/3, Price:100/250/500)|     |            |            |            |            | O          |
| 7   | Confirm   | Return                                       |     |                                |     |            |            |            |            |            |
| 8   |           | TotalAmount                                  |     |                                |     |            |            |            |            |            |
| 9   |           |                                              |     | 0                              |     | O          |            |            |            |            |
| 10  |           |                                              |     | 200                            |     |            | O          |            |            |            |
| 11  |           |                                              |     | 800                            |     |            |            | O          |            |            |
| 12  |           |                                              |     | 2100                           |     |            |            |            |            | O          |
| 13  |           |                                              |     | 50000                          |     |            |            |            | O          |            |
| 14  |           | Status                                       |     |                                |     |            |            |            |            |            |
| 15  |           |                                              |     | Pending                        |     | O          | O          | O          | O          | O          |
| 16  |           | IsPaid                                       |     |                                |     |            |            |            |            |            |
| 17  |           |                                              |     | false                          |     | O          | O          | O          | O          | O          |
| 18  |           | ID                                           |     |                                |     |            |            |            |            |            |
| 19  |           |                                              |     | Not Nil                        |     | O          | O          | O          | O          | O          |
| 20  |           | Error                                        |     |                                |     |            |            |            |            |            |
| 21  |           |                                              |     | Key: 'Order.Items' validation  |     | O          |            |            |            |            |
| 22  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                                |     | A          | N          | N          | N          | N          |
| 23  |           | Passed/Failed                                |     |                                |     | P          | P          | P          | P          | P          |
| 24  |           | Executed Date                                |     |                                |     | 2025-12-10 | 2025-12-10 | 2025-12-10 | 2025-12-10 | 2025-12-10 |
| 25  |           | Defect ID                                    |     |                                |     |            |            |            | DF-O-VA-01 |            |

### `domain.Order.Update`

#### Meta

|                  |     |                                                         |     |     |                    |     |     |     |     | 0   | 1                  | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ------------------ | --- | --- | ---------------- |
| Function Code    |     | O-U-01                                                  |     |     | Function Name      |     |     |     |     |     | domain.Order.Update|     |     |                  |
| Created By       |     | CodeCompanion                                           |     |     | Executed By        |     |     |     |     |     |                    |     |     |                  |
| Lines of code    |     | 16                                                      |     |     | Lack of test cases |     |     |     |     |     | 0                  |     |     |                  |
| Test requirement |     | Validate Order update logic and UpdatedAt timestamp     |     |     |                    |     |     |     |     |     |                    |     |     |                  |
| Passed           |     | Failed                                                  |     |     | Untested           |     |     |     |     |     | N                  | A   | B   | Total Test Cases |
| 5                |     | 0                                                       |     |     | 0                  |     |     |     |     |     | 5                  | 0   | 0   | 5                |

#### Sheet

| #   |           |                                              |     |                      |     | O-U-01-01 | O-U-01-02 | O-U-01-03 | O-U-01-04 | O-U-01-05 |
| --- | --------- | -------------------------------------------- | --- | -------------------- | --- | --------- | --------- | --------- | --------- | --------- |
| 1   | Condition | Field Changed                                |     |                      |     |           |           |           |           |           |
| 2   |           |                                              |     | Address only         |     | O         |           |           |           |           |
| 3   |           |                                              |     | Status only          |     |           | O         |           |           |           |
| 4   |           |                                              |     | IsPaid only          |     |           |           | O         |           |           |
| 5   |           |                                              |     | All fields           |     |           |           |           | O         |           |
| 6   |           |                                              |     | None (no change)     |     |           |           |           |           | O         |
| 7   | Confirm   | Return                                       |     |                      |     |           |           |           |           |           |
| 8   |           | Address                                      |     |                      |     |           |           |           |           |           |
| 9   |           |                                              |     | Updated              |     | O         |           |           | O         |           |
| 10  |           |                                              |     | Unchanged            |     |           | O         | O         |           | O         |
| 11  |           | Status                                       |     |                      |     |           |           |           |           |           |
| 12  |           |                                              |     | Updated              |     |           | O         |           | O         |           |
| 13  |           |                                              |     | Unchanged            |     | O         |           | O         |           | O         |
| 14  |           | IsPaid                                       |     |                      |     |           |           |           |           |           |
| 15  |           |                                              |     | Updated              |     |           |           | O         | O         |           |
| 16  |           |                                              |     | Unchanged            |     | O         | O         |           |           | O         |
| 17  |           | UpdatedAt                                    |     |                      |     |           |           |           |           |           |
| 18  |           |                                              |     | Changed              |     | O         | O         | O         | O         |           |
| 19  |           |                                              |     | Unchanged            |     |           |           |           |           | O         |
| 20  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                      |     | N         | N         | N         | N         | N         |
| 21  |           | Passed/Failed                                |     |                      |     | P         | P         | P         | P         | P         |
| 22  |           | Executed Date                                |     |                      |     | 2025-12-10| 2025-12-10| 2025-12-10| 2025-12-10| 2025-12-10|
| 23  |           | Defect ID                                    |     |                      |     |           |           |           |           |           |
