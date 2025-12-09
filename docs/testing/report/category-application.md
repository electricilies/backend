### `application.Category.Create`

#### Meta

|                  |     |                                                        |     |     |                    |     |     |     |     | 0   | 1                           | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------------------ | --- | --- | ------------------ | --- | --- | --- | --- | --- | --------------------------- | --- | --- | ---------------- |
| Function Code    |     | C-APP-01                                               |     |     | Function Name      |     |     |     |     |     | application.Category.Create |     |     |                  |
| Created By       |     | CodeCompanion                                          |     |     | Executed By        |     |     |     |     |     |                             |     |     |                  |
| Lines of code    |     | 18                                                     |     |     | Lack of test cases |     |     |     |     |     | 0                           |     |     |                  |
| Test requirement |     | Validate Category creation with valid and invalid data |     |     |                    |     |     |     |     |     |                             |     |     |                  |
| Passed           |     | Failed                                                 |     |     | Untested           |     |     |     |     |     | N                           | A   | B   | Total Test Cases |
| 3                |     | 0                                                      |     |     | 0                  |     |     |     |     |     | 2                           | 1   | 0   | 3                |

#### Sheet

| #   |           |                                              |     |                   |     | C-APP-01-01 | C-APP-01-02 | C-APP-01-03 |
| --- | --------- | -------------------------------------------- | --- | ----------------- | --- | ----------- | ----------- | ----------- |
| 1   | Condition | Precondition                                 |     |                   |     |             |             |             |
| 2   |           |                                              |     | DB/Redis Up       |     | O           | O           | O           |
| 3   |           | Name                                         |     |                   |     |             |             |             |
| 4   |           |                                              |     | "Electronics"     |     | O           |             |             |
| 5   |           |                                              |     | "Home Appliances" |     |             | O           |             |
| 6   |           |                                              |     | ""                |     |             |             | O           |
| 7   | Confirm   | Return                                       |     |                   |     |             |             |             |
| 8   |           | ID                                           |     |                   |     |             |             |             |
| 9   |           |                                              |     | Not Nil           |     | O           | O           |             |
| 10  |           | Name                                         |     |                   |     |             |             |             |
| 11  |           |                                              |     | "Electronics"     |     | O           |             |             |
| 12  |           |                                              |     | "Home Appliances" |     |             | O           |             |
| 13  |           | Error                                        |     |                   |     |             |             |             |
| 14  |           |                                              |     | Not Nil           |     |             |             | O           |
| 15  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                   |     | N           | N           | A           |
| 16  |           | Passed/Failed                                |     |                   |     | P           | P           | P           |
| 17  |           | Executed Date                                |     |                   |     | 2025-12-09  | 2025-12-09  | 2025-12-09  |
| 18  |           | Defect ID                                    |     |                   |     |             |             |             |

### `application.Category.Get`

#### Meta

|                  |     |                                                                 |     |     |                    |     |     |     |     | 0   | 1                        | 2   | 3   | 4                |
| ---------------- | --- | --------------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ------------------------ | --- | --- | ---------------- |
| Function Code    |     | C-APP-02                                                        |     |     | Function Name      |     |     |     |     |     | application.Category.Get |     |     |                  |
| Created By       |     | CodeCompanion                                                   |     |     | Executed By        |     |     |     |     |     |                          |     |     |                  |
| Lines of code    |     | 17                                                              |     |     | Lack of test cases |     |     |     |     |     | 0                        |     |     |                  |
| Test requirement |     | Validate Category retrieval, cache behavior, and error handling |     |     |                    |     |     |     |     |     |                          |     |     |                  |
| Passed           |     | Failed                                                          |     |     | Untested           |     |     |     |     |     | N                        | A   | B   | Total Test Cases |
| 4                |     | 0                                                               |     |     | 0                  |     |     |     |     |     | 3                        | 1   | 0   | 4                |

#### Sheet

| #   |           |                                              |     |                       |     | C-APP-02-01 | C-APP-02-02 | C-APP-02-03 | C-APP-02-04 |
| --- | --------- | -------------------------------------------- | --- | --------------------- | --- | ----------- | ----------- | ----------- | ----------- |
| 1   | Condition | Precondition                                 |     |                       |     |             |             |             |             |
| 2   |           |                                              |     | C-APP-01-01           |     | O           |             |             | O           |
| 3   |           |                                              |     | C-APP-04-01           |     |             | O           |             |             |
| 4   |           |                                              |     | DB/Redis Up           |     |             |             | O           |             |
| 5   |           | CategoryID                                   |     |                       |     |             |             |             |             |
| 6   |           |                                              |     | Existing (First)      |     | O           | O           |             | O           |
| 7   |           |                                              |     | Non-Existent          |     |             |             | O           |             |
| 8   | Confirm   | Return                                       |     |                       |     |             |             |             |             |
| 9   |           | Name                                         |     |                       |     |             |             |             |             |
| 10  |           |                                              |     | "Electronics"         |     | O           |             |             | O           |
| 11  |           |                                              |     | "Updated Electronics" |     |             | O           |             |             |
| 12  |           | Error                                        |     |                       |     |             |             |             |             |
| 13  |           |                                              |     | Not Nil               |     |             |             | O           |             |
| 14  |           | Cache                                        |     |                       |     |             |             |             |             |
| 15  |           |                                              |     | Hit                   |     |             |             |             | O           |
| 16  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                       |     | N           | N           | A           | N           |
| 17  |           | Passed/Failed                                |     |                       |     | P           | P           | P           | P           |
| 18  |           | Executed Date                                |     |                       |     | 2025-12-09  | 2025-12-09  | 2025-12-09  | 2025-12-09  |
| 19  |           | Defect ID                                    |     |                       |     |             |             |             |             |

### `application.Category.List`

#### Meta

|                  |     |                                                          |     |     |                    |     |     |     |     | 0   | 1                         | 2   | 3   | 4                |
| ---------------- | --- | -------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ------------------------- | --- | --- | ---------------- |
| Function Code    |     | C-APP-03                                                 |     |     | Function Name      |     |     |     |     |     | application.Category.List |     |     |                  |
| Created By       |     | CodeCompanion                                            |     |     | Executed By        |     |     |     |     |     |                           |     |     |                  |
| Lines of code    |     | 40                                                       |     |     | Lack of test cases |     |     |     |     |     | 0                         |     |     |                  |
| Test requirement |     | Validate Category listing, pagination, search, and cache |     |     |                    |     |     |     |     |     |                           |     |     |                  |
| Passed           |     | Failed                                                   |     |     | Untested           |     |     |     |     |     | N                         | A   | B   | Total Test Cases |
| 5                |     | 0                                                        |     |     | 0                  |     |     |     |     |     | 5                         | 0   | 0   | 5                |

#### Sheet

| #   |           |                                              |     |                  |     | C-APP-03-01 | C-APP-03-02 | C-APP-03-03 | C-APP-03-04 | C-APP-03-05 |
| --- | --------- | -------------------------------------------- | --- | ---------------- | --- | ----------- | ----------- | ----------- | ----------- | ----------- |
| 1   | Condition | Precondition                                 |     |                  |     |             |             |             |             |             |
| 2   |           |                                              |     | C-APP-01-01 & 02 |     | O           | O           | O           | O           | O           |
| 3   |           | Page                                         |     |                  |     |             |             |             |             |             |
| 4   |           |                                              |     | 1                |     | O           | O           | O           |             | O           |
| 5   |           |                                              |     | 2                |     |             |             |             | O           |             |
| 6   |           | Limit                                        |     |                  |     |             |             |             |             |             |
| 7   |           |                                              |     | 1                |     |             |             | O           | O           |             |
| 8   |           |                                              |     | 10               |     | O           | O           |             |             | O           |
| 9   |           | Search                                       |     |                  |     |             |             |             |             |             |
| 10  |           |                                              |     | ""               |     | O           |             | O           | O           | O           |
| 11  |           |                                              |     | "Electronics"    |     |             | O           |             |             |             |
| 12  | Confirm   | Return                                       |     |                  |     |             |             |             |             |             |
| 13  |           | TotalItems                                   |     |                  |     |             |             |             |             |             |
| 14  |           |                                              |     | 2                |     | O           |             | O           | O           | O           |
| 15  |           |                                              |     | >= 1             |     |             | O           |             |             |             |
| 16  |           | Data Length                                  |     |                  |     |             |             |             |             |             |
| 17  |           |                                              |     | 2                |     | O           |             |             |             | O           |
| 18  |           |                                              |     | 1                |     |             |             | O           | O           |             |
| 19  |           |                                              |     | >= 1             |     |             | O           |             |             |             |
| 20  |           | Cache                                        |     |                  |     |             |             |             |             |             |
| 21  |           |                                              |     | Hit              |     |             |             |             |             | O           |
| 22  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                  |     | N           | N           | N           | N           | N           |
| 23  |           | Passed/Failed                                |     |                  |     | P           | P           | P           | P           | P           |
| 24  |           | Executed Date                                |     |                  |     | 2025-12-09  | 2025-12-09  | 2025-12-09  | 2025-12-09  | 2025-12-09  |
| 25  |           | Defect ID                                    |     |                  |     |             |             |             |             |             |

### `application.Category.Update`

#### Meta

|                  |     |                                                                  |     |     |                    |     |     |     |     | 0   | 1                           | 2   | 3   | 4                |
| ---------------- | --- | ---------------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | --------------------------- | --- | --- | ---------------- |
| Function Code    |     | C-APP-04                                                         |     |     | Function Name      |     |     |     |     |     | application.Category.Update |     |     |                  |
| Created By       |     | CodeCompanion                                                    |     |     | Executed By        |     |     |     |     |     |                             |     |     |                  |
| Lines of code    |     | 18                                                               |     |     | Lack of test cases |     |     |     |     |     | 0                           |     |     |                  |
| Test requirement |     | Validate Category update, error handling, and cache invalidation |     |     |                    |     |     |     |     |     |                             |     |     |                  |
| Passed           |     | Failed                                                           |     |     | Untested           |     |     |     |     |     | N                           | A   | B   | Total Test Cases |
| 4                |     | 0                                                                |     |     | 0                  |     |     |     |     |     | 2                           | 2   | 0   | 4                |

#### Sheet

| #   |           |                                              |     |                           |     | C-APP-04-01 | C-APP-04-02 | C-APP-04-03 | C-APP-04-04 |
| --- | --------- | -------------------------------------------- | --- | ------------------------- | --- | ----------- | ----------- | ----------- | ----------- |
| 1   | Condition | Precondition                                 |     |                           |     |             |             |             |             |
| 2   |           |                                              |     | C-APP-01-01               |     | O           | O           |             |             |
| 3   |           |                                              |     | C-APP-01-02               |     |             |             |             | O           |
| 4   |           |                                              |     | DB/Redis Up               |     |             |             | O           |             |
| 5   |           | CategoryID                                   |     |                           |     |             |             |             |             |
| 6   |           |                                              |     | Existing (First)          |     | O           | O           |             |             |
| 7   |           |                                              |     | Existing (Second)         |     |             |             |             | O           |
| 8   |           |                                              |     | Non-Existent              |     |             |             | O           |             |
| 9   |           | Name                                         |     |                           |     |             |             |             |             |
| 10  |           |                                              |     | "Updated Electronics"     |     | O           |             |             |             |
| 11  |           |                                              |     | ""                        |     |             | O           |             |             |
| 12  |           |                                              |     | "New Name"                |     |             |             | O           |             |
| 13  |           |                                              |     | "Cache Invalidation Test" |     |             |             |             | O           |
| 14  | Confirm   | Return                                       |     |                           |     |             |             |             |             |
| 15  |           | Name                                         |     |                           |     |             |             |             |             |
| 16  |           |                                              |     | "Updated Electronics"     |     | O           |             |             |             |
| 17  |           |                                              |     | "Cache Invalidation Test" |     |             |             |             | O           |
| 18  |           | Error                                        |     |                           |     |             |             |             |             |
| 19  |           |                                              |     | Not Nil                   |     |             | O           | O           |             |
| 20  |           | Cache                                        |     |                           |     |             |             |             |             |
| 21  |           |                                              |     | Invalidated               |     |             |             |             | O           |
| 22  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                           |     | N           | A           | A           | N           |
| 23  |           | Passed/Failed                                |     |                           |     | P           | P           | P           | P           |
| 24  |           | Executed Date                                |     |                           |     | 2025-12-09  | 2025-12-09  | 2025-12-09  | 2025-12-09  |
| 25  |           | Defect ID                                    |     |                           |     |             |             |             |             |
