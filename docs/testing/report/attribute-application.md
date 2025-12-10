### `application.Attribute.Create`

#### Meta

|                  |     |                                                        |     |     |                    |     |     |     |     | 0   | 1                            | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------------------ | --- | --- | ------------------ | --- | --- | --- | --- | --- | ---------------------------- | --- | --- | ---------------- |
| Function Code    |     | C-APP-05                                               |     |     | Function Name      |     |     |     |     |     | application.Attribute.Create |     |     |                  |
| Created By       |     | CodeCompanion                                          |     |     | Executed By        |     |     |     |     |     |                              |     |     |                  |
| Lines of code    |     | 18                                                     |     |     | Lack of test cases |     |     |     |     |     | 0                            |     |     |                  |
| Test requirement |     | Validate Attribute creation with valid and invalid data|     |     |                    |     |     |     |     |     |                              |     |     |                  |
| Passed           |     | Failed                                                 |     |     | Untested           |     |     |     |     |     | N                            | A   | B   | Total Test Cases |
| 3                |     | 0                                                      |     |     | 0                  |     |     |     |     |     | 2                            | 1   | 0   | 3                |

#### Sheet

| #   |           |                                              |     |                   |     | C-APP-05-01 | C-APP-05-02 | C-APP-05-03 |
| --- | --------- | -------------------------------------------- | --- | ----------------- | --- | ----------- | ----------- | ----------- |
| 1   | Condition | Precondition                                 |     |                   |     |             |             |             |
| 2   |           |                                              |     | DB/Redis Up       |     | O           | O           | O           |
| 3   |           |                                              |     | C-APP-05-01       |     |             | O           |             |
| 4   |           | Code                                         |     |                   |     |             |             |             |
| 5   |           |                                              |     | "color"           |     | O           | O           |             |
| 6   |           |                                              |     | "size"            |     |             |             | O           |
| 7   | Confirm   | Return                                       |     |                   |     |             |             |             |
| 8   |           | ID                                           |     |                   |     |             |             |             |
| 9   |           |                                              |     | Not Nil           |     | O           |             | O           |
| 10  |           | Name                                         |     |                   |     |             |             |             |
| 11  |           |                                              |     | "Color"           |     | O           |             |             |
| 12  |           |                                              |     | "Size"            |     |             |             | O           |
| 13  |           | Error                                        |     |                   |     |             |             |             |
| 14  |           |                                              |     | Not Nil           |     |             | O           |             |
| 15  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                   |     | N           | A           | N           |
| 16  |           | Passed/Failed                                |     |                   |     | P           | P           | P           |
| 17  |           | Executed Date                                |     |                   |     | 2025-12-10  | 2025-12-10  | 2025-12-10  |
| 18  |           | Defect ID                                    |     |                   |     |             |             |             |

### `application.Attribute.CreateValue`

#### Meta

|                  |     |                                                             |     |     |                    |     |     |     |     | 0   | 1                                 | 2   | 3   | 4                |
| ---------------- | --- | ----------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | --------------------------------- | --- | --- | ---------------- |
| Function Code    |     | C-APP-06                                                    |     |     | Function Name      |     |     |     |     |     | application.Attribute.CreateValue |     |     |                  |
| Created By       |     | CodeCompanion                                               |     |     | Executed By        |     |     |     |     |     |                                   |     |     |                  |
| Lines of code    |     | 18                                                          |     |     | Lack of test cases |     |     |     |     |     | 0                                 |     |     |                  |
| Test requirement |     | Validate Attribute Value creation                           |     |     |                    |     |     |     |     |     |                                   |     |     |                  |
| Passed           |     | Failed                                                      |     |     | Untested           |     |     |     |     |     | N                                 | A   | B   | Total Test Cases |
| 3                |     | 0                                                           |     |     | 0                  |     |     |     |     |     | 2                                 | 1   | 0   | 3                |

#### Sheet

| #   |           |                                              |     |                   |     | C-APP-06-01 | C-APP-06-02 | C-APP-06-03 |
| --- | --------- | -------------------------------------------- | --- | ----------------- | --- | ----------- | ----------- | ----------- |
| 1   | Condition | Precondition                                 |     |                   |     |             |             |             |
| 2   |           |                                              |     | C-APP-05-03       |     | O           | O           |             |
| 3   |           | AttributeID                                  |     |                   |     |             |             |             |
| 4   |           |                                              |     | Existing (Size)   |     | O           | O           |             |
| 5   |           |                                              |     | Non-Existent      |     |             |             | O           |
| 6   |           | Value                                        |     |                   |     |             |             |             |
| 7   |           |                                              |     | "Small"           |     | O           |             |             |
| 8   |           |                                              |     | "Medium"          |     |             | O           |             |
| 9   |           |                                              |     | "Test"            |     |             |             | O           |
| 10  | Confirm   | Return                                       |     |                   |     |             |             |             |
| 11  |           | ID                                           |     |                   |     |             |             |             |
| 12  |           |                                              |     | Not Nil           |     | O           | O           |             |
| 13  |           | Value                                        |     |                   |     |             |             |             |
| 14  |           |                                              |     | "Small"           |     | O           |             |             |
| 15  |           |                                              |     | "Medium"          |     |             | O           |             |
| 16  |           | Error                                        |     |                   |     |             |             |             |
| 17  |           |                                              |     | Not Nil           |     |             |             | O           |
| 18  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                   |     | N           | N           | A           |
| 19  |           | Passed/Failed                                |     |                   |     | P           | P           | P           |
| 20  |           | Executed Date                                |     |                   |     | 2025-12-10  | 2025-12-10  | 2025-12-10  |
| 21  |           | Defect ID                                    |     |                   |     |             |             |             |

### `application.Attribute.List`

#### Meta

|                  |     |                                                          |     |     |                    |     |     |     |     | 0   | 1                          | 2   | 3   | 4                |
| ---------------- | --- | -------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | -------------------------- | --- | --- | ---------------- |
| Function Code    |     | C-APP-07                                                 |     |     | Function Name      |     |     |     |     |     | application.Attribute.List |     |     |                  |
| Created By       |     | CodeCompanion                                            |     |     | Executed By        |     |     |     |     |     |                            |     |     |                  |
| Lines of code    |     | 40                                                       |     |     | Lack of test cases |     |     |     |     |     | 0                          |     |     |                  |
| Test requirement |     | Validate Attribute listing, pagination, and cache        |     |     |                    |     |     |     |     |     |                            |     |     |                  |
| Passed           |     | Failed                                                   |     |     | Untested           |     |     |     |     |     | N                          | A   | B   | Total Test Cases |
| 3                |     | 0                                                        |     |     | 0                  |     |     |     |     |     | 3                          | 0   | 0   | 3                |

#### Sheet

| #   |           |                                              |     |                  |     | C-APP-07-01 | C-APP-07-02 | C-APP-07-03 |
| --- | --------- | -------------------------------------------- | --- | ---------------- | --- | ----------- | ----------- | ----------- |
| 1   | Condition | Precondition                                 |     |                  |     |             |             |             |
| 2   |           |                                              |     | C-APP-05-01 & 03 |     | O           |             | O           |
| 3   |           |                                              |     | C-APP-12-01      |     |             | O           |             |
| 4   |           | Page                                         |     |                  |     |             |             |             |
| 5   |           |                                              |     | 1                |     | O           | O           | O           |
| 6   |           | Limit                                        |     |                  |     |             |             |             |
| 7   |           |                                              |     | 10               |     | O           | O           | O           |
| 8   | Confirm   | Return                                       |     |                  |     |             |             |             |
| 9   |           | TotalItems                                   |     |                  |     |             |             |             |
| 10  |           |                                              |     | 2                |     | O           |             | O           |
| 11  |           |                                              |     | 1                |     |             | O           |             |
| 12  |           | Data Length                                  |     |                  |     |             |             |             |
| 13  |           |                                              |     | 2                |     | O           |             | O           |
| 14  |           |                                              |     | 1                |     |             | O           |             |
| 15  |           | Cache                                        |     |                  |     |             |             |             |
| 16  |           |                                              |     | Hit              |     |             |             | O           |
| 17  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                  |     | N           | N           | N           |
| 18  |           | Passed/Failed                                |     |                  |     | P           | P           | P           |
| 19  |           | Executed Date                                |     |                  |     | 2025-12-10  | 2025-12-10  | 2025-12-10  |
| 20  |           | Defect ID                                    |     |                  |     |             |             |             |

### `application.Attribute.Get`

#### Meta

|                  |     |                                                                 |     |     |                    |     |     |     |     | 0   | 1                         | 2   | 3   | 4                |
| ---------------- | --- | --------------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ------------------------- | --- | --- | ---------------- |
| Function Code    |     | C-APP-08                                                        |     |     | Function Name      |     |     |     |     |     | application.Attribute.Get |     |     |                  |
| Created By       |     | CodeCompanion                                                   |     |     | Executed By        |     |     |     |     |     |                           |     |     |                  |
| Lines of code    |     | 17                                                              |     |     | Lack of test cases |     |     |     |     |     | 0                         |     |     |                  |
| Test requirement |     | Validate Attribute retrieval, cache behavior, and error handling|     |     |                    |     |     |     |     |     |                           |     |     |                  |
| Passed           |     | Failed                                                          |     |     | Untested           |     |     |     |     |     | N                         | A   | B   | Total Test Cases |
| 4                |     | 0                                                               |     |     | 0                  |     |     |     |     |     | 2                         | 2   | 0   | 4                |

#### Sheet

| #   |           |                                              |     |                       |     | C-APP-08-01 | C-APP-08-02 | C-APP-08-03 | C-APP-08-04 |
| --- | --------- | -------------------------------------------- | --- | --------------------- | --- | ----------- | ----------- | ----------- | ----------- |
| 1   | Condition | Precondition                                 |     |                       |     |             |             |             |             |
| 2   |           |                                              |     | C-APP-05-01           |     | O           |             |             | O           |
| 3   |           |                                              |     | C-APP-12-01           |     |             | O           |             |             |
| 4   |           | AttributeID                                  |     |                       |     |             |             |             |             |
| 5   |           |                                              |     | Existing (Color)      |     | O           | O           |             | O           |
| 6   |           |                                              |     | Non-Existent          |     |             |             | O           |             |
| 7   | Confirm   | Return                                       |     |                       |     |             |             |             |             |
| 8   |           | Name                                         |     |                       |     |             |             |             |             |
| 9   |           |                                              |     | "Color"               |     | O           |             |             | O           |
| 10  |           | Error                                        |     |                       |     |             |             |             |             |
| 11  |           |                                              |     | Not Nil               |     |             | O           | O           |             |
| 12  |           | Cache                                        |     |                       |     |             |             |             |             |
| 13  |           |                                              |     | Hit                   |     |             |             |             | O           |
| 14  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                       |     | N           | A           | A           | N           |
| 15  |           | Passed/Failed                                |     |                       |     | P           | P           | P           | P           |
| 16  |           | Executed Date                                |     |                       |     | 2025-12-10  | 2025-12-10  | 2025-12-10  | 2025-12-10  |
| 17  |           | Defect ID                                    |     |                       |     |             |             |             |             |

### `application.Attribute.ListValues`

#### Meta

|                  |     |                                                               |     |     |                    |     |     |     |     | 0   | 1                                | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | -------------------------------- | --- | --- | ---------------- |
| Function Code    |     | C-APP-09                                                      |     |     | Function Name      |     |     |     |     |     | application.Attribute.ListValues |     |     |                  |
| Created By       |     | CodeCompanion                                                 |     |     | Executed By        |     |     |     |     |     |                                  |     |     |                  |
| Lines of code    |     | 40                                                            |     |     | Lack of test cases |     |     |     |     |     | 0                                |     |     |                  |
| Test requirement |     | Validate Attribute Values listing, pagination, and cache      |     |     |                    |     |     |     |     |     |                                  |     |     |                  |
| Passed           |     | Failed                                                        |     |     | Untested           |     |     |     |     |     | N                                | A   | B   | Total Test Cases |
| 3                |     | 0                                                             |     |     | 0                  |     |     |     |     |     | 3                                | 0   | 0   | 3                |

#### Sheet

| #   |           |                                              |     |                  |     | C-APP-09-01 | C-APP-09-02 | C-APP-09-03 |
| --- | --------- | -------------------------------------------- | --- | ---------------- | --- | ----------- | ----------- | ----------- |
| 1   | Condition | Precondition                                 |     |                  |     |             |             |             |
| 2   |           |                                              |     | C-APP-06-01 & 02 |     | O           |             | O           |
| 3   |           |                                              |     | C-APP-13-01      |     |             | O           |             |
| 4   |           | AttributeID                                  |     |                  |     |             |             |             |
| 5   |           |                                              |     | Existing (Size)  |     | O           | O           | O           |
| 6   | Confirm   | Return                                       |     |                  |     |             |             |             |
| 7   |           | TotalItems                                   |     |                  |     |             |             |             |
| 8   |           |                                              |     | 2                |     | O           |             | O           |
| 9   |           |                                              |     | 1                |     |             | O           |             |
| 10  |           | Data Length                                  |     |                  |     |             |             |             |
| 11  |           |                                              |     | 2                |     | O           |             | O           |
| 12  |           |                                              |     | 1                |     |             | O           |             |
| 13  |           | Cache                                        |     |                  |     |             |             |             |
| 14  |           |                                              |     | Hit              |     |             |             | O           |
| 15  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                  |     | N           | N           | N           |
| 16  |           | Passed/Failed                                |     |                  |     | P           | P           | P           |
| 17  |           | Executed Date                                |     |                  |     | 2025-12-10  | 2025-12-10  | 2025-12-10  |
| 18  |           | Defect ID                                    |     |                  |     |             |             |             |

### `application.Attribute.Update`

#### Meta

|                  |     |                                                                  |     |     |                    |     |     |     |     | 0   | 1                            | 2   | 3   | 4                |
| ---------------- | --- | ---------------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ---------------------------- | --- | --- | ---------------- |
| Function Code    |     | C-APP-10                                                         |     |     | Function Name      |     |     |     |     |     | application.Attribute.Update |     |     |                  |
| Created By       |     | CodeCompanion                                                    |     |     | Executed By        |     |     |     |     |     |                              |     |     |                  |
| Lines of code    |     | 18                                                               |     |     | Lack of test cases |     |     |     |     |     | 0                            |     |     |                  |
| Test requirement |     | Validate Attribute update, error handling, and cache invalidation|     |     |                    |     |     |     |     |     |                              |     |     |                  |
| Passed           |     | Failed                                                           |     |     | Untested           |     |     |     |     |     | N                            | A   | B   | Total Test Cases |
| 3                |     | 0                                                                |     |     | 0                  |     |     |     |     |     | 2                            | 1   | 0   | 3                |

#### Sheet

| #   |           |                                              |     |                           |     | C-APP-10-01 | C-APP-10-02 | C-APP-10-03 |
| --- | --------- | -------------------------------------------- | --- | ------------------------- | --- | ----------- | ----------- | ----------- |
| 1   | Condition | Precondition                                 |     |                           |     |             |             |             |
| 2   |           |                                              |     | C-APP-05-01               |     | O           |             |             |
| 3   |           |                                              |     | C-APP-05-03               |     |             |             | O           |
| 4   |           | AttributeID                                  |     |                           |     |             |             |             |
| 5   |           |                                              |     | Existing (Color)          |     | O           |             |             |
| 6   |           |                                              |     | Existing (Size)           |     |             |             | O           |
| 7   |           |                                              |     | Non-Existent              |     |             | O           |             |
| 8   |           | Name                                         |     |                           |     |             |             |             |
| 9   |           |                                              |     | "Updated Color"           |     | O           |             |             |
| 10  |           |                                              |     | "Cache Invalidation Test" |     |             |             | O           |
| 11  |           |                                              |     | "New Name"                |     |             | O           |             |
| 12  | Confirm   | Return                                       |     |                           |     |             |             |             |
| 13  |           | Name                                         |     |                           |     |             |             |             |
| 14  |           |                                              |     | "Updated Color"           |     | O           |             |             |
| 15  |           |                                              |     | "Cache Invalidation Test" |     |             |             | O           |
| 16  |           | Error                                        |     |                           |     |             |             |             |
| 17  |           |                                              |     | Not Nil                   |     |             | O           |             |
| 18  |           | Cache                                        |     |                           |     |             |             |             |
| 19  |           |                                              |     | Invalidated               |     |             |             | O           |
| 20  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                           |     | N           | A           | N           |
| 21  |           | Passed/Failed                                |     |                           |     | P           | P           | P           |
| 22  |           | Executed Date                                |     |                           |     | 2025-12-10  | 2025-12-10  | 2025-12-10  |
| 23  |           | Defect ID                                    |     |                           |     |             |             |             |

### `application.Attribute.UpdateValue`

#### Meta

|                  |     |                                                                        |     |     |                    |     |     |     |     | 0   | 1                                 | 2   | 3   | 4                |
| ---------------- | --- | ---------------------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | --------------------------------- | --- | --- | ---------------- |
| Function Code    |     | C-APP-11                                                               |     |     | Function Name      |     |     |     |     |     | application.Attribute.UpdateValue |     |     |                  |
| Created By       |     | CodeCompanion                                                          |     |     | Executed By        |     |     |     |     |     |                                   |     |     |                  |
| Lines of code    |     | 20                                                                     |     |     | Lack of test cases |     |     |     |     |     | 0                                 |     |     |                  |
| Test requirement |     | Validate Attribute Value update, error handling, and cache invalidation|     |     |                    |     |     |     |     |     |                                   |     |     |                  |
| Passed           |     | Failed                                                                 |     |     | Untested           |     |     |     |     |     | N                                 | A   | B   | Total Test Cases |
| 3                |     | 0                                                                      |     |     | 0                  |     |     |     |     |     | 2                                 | 1   | 0   | 3                |

#### Sheet

| #   |           |                                              |     |                    |     | C-APP-11-01 | C-APP-11-02 | C-APP-11-03 |
| --- | --------- | -------------------------------------------- | --- | ------------------ | --- | ----------- | ----------- | ----------- |
| 1   | Condition | Precondition                                 |     |                    |     |             |             |             |
| 2   |           |                                              |     | C-APP-06-01        |     | O           |             | O           |
| 3   |           | AttributeID                                  |     |                    |     |             |             |             |
| 4   |           |                                              |     | Existing (Size)    |     | O           | O           | O           |
| 5   |           | AttributeValueID                             |     |                    |     |             |             |             |
| 6   |           |                                              |     | Existing (Small)   |     | O           |             | O           |
| 7   |           |                                              |     | Non-Existent       |     |             | O           |             |
| 8   |           | Value                                        |     |                    |     |             |             |             |
| 9   |           |                                              |     | "Extra Small"      |     | O           |             |             |
| 10  |           |                                              |     | "Cache Test Value" |     |             |             | O           |
| 11  |           |                                              |     | "New Value"        |     |             | O           |             |
| 12  | Confirm   | Return                                       |     |                    |     |             |             |             |
| 13  |           | Value                                        |     |                    |     |             |             |             |
| 14  |           |                                              |     | "Extra Small"      |     | O           |             |             |
| 15  |           |                                              |     | "Cache Test Value" |     |             |             | O           |
| 16  |           | Error                                        |     |                    |     |             |             |             |
| 17  |           |                                              |     | Not Nil            |     |             | O           |             |
| 18  |           | Cache                                        |     |                    |     |             |             |             |
| 19  |           |                                              |     | Invalidated        |     |             |             | O           |
| 20  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                    |     | N           | A           | N           |
| 21  |           | Passed/Failed                                |     |                    |     | P           | P           | P           |
| 22  |           | Executed Date                                |     |                    |     | 2025-12-10  | 2025-12-10  | 2025-12-10  |
| 23  |           | Defect ID                                    |     |                    |     |             |             |             |

### `application.Attribute.Delete`

#### Meta

|                  |     |                                                        |     |     |                    |     |     |     |     | 0   | 1                            | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------------------ | --- | --- | ------------------ | --- | --- | --- | --- | --- | ---------------------------- | --- | --- | ---------------- |
| Function Code    |     | C-APP-12                                               |     |     | Function Name      |     |     |     |     |     | application.Attribute.Delete |     |     |                  |
| Created By       |     | CodeCompanion                                          |     |     | Executed By        |     |     |     |     |     |                              |     |     |                  |
| Lines of code    |     | 15                                                     |     |     | Lack of test cases |     |     |     |     |     | 0                            |     |     |                  |
| Test requirement |     | Validate Attribute deletion and error handling         |     |     |                    |     |     |     |     |     |                              |     |     |                  |
| Passed           |     | Failed                                                 |     |     | Untested           |     |     |     |     |     | N                            | A   | B   | Total Test Cases |
| 2                |     | 0                                                      |     |     | 0                  |     |     |     |     |     | 1                            | 1   | 0   | 2                |

#### Sheet

| #   |           |                                              |     |                  |     | C-APP-12-01 | C-APP-12-02 |
| --- | --------- | -------------------------------------------- | --- | ---------------- | --- | ----------- | ----------- |
| 1   | Condition | Precondition                                 |     |                  |     |             |             |
| 2   |           |                                              |     | C-APP-05-01      |     | O           |             |
| 3   |           | AttributeID                                  |     |                  |     |             |             |
| 4   |           |                                              |     | Existing (Color) |     | O           |             |
| 5   |           |                                              |     | Non-Existent     |     |             | O           |
| 6   | Confirm   | Return                                       |     |                  |     |             |             |
| 7   |           | Error                                        |     |                  |     |             |             |
| 8   |           |                                              |     | Nil              |     | O           |             |
| 9   |           |                                              |     | Not Nil          |     |             | O           |
| 10  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                  |     | N           | A           |
| 11  |           | Passed/Failed                                |     |                  |     | P           | P           |
| 12  |           | Executed Date                                |     |                  |     | 2025-12-10  | 2025-12-10  |
| 13  |           | Defect ID                                    |     |                  |     |             |             |

### `application.Attribute.DeleteValue`

#### Meta

|                  |     |                                                            |     |     |                    |     |     |     |     | 0   | 1                                 | 2   | 3   | 4                |
| ---------------- | --- | ---------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | --------------------------------- | --- | --- | ---------------- |
| Function Code    |     | C-APP-13                                                   |     |     | Function Name      |     |     |     |     |     | application.Attribute.DeleteValue |     |     |                  |
| Created By       |     | CodeCompanion                                              |     |     | Executed By        |     |     |     |     |     |                                   |     |     |                  |
| Lines of code    |     | 15                                                         |     |     | Lack of test cases |     |     |     |     |     | 0                                 |     |     |                  |
| Test requirement |     | Validate Attribute Value deletion                          |     |     |                    |     |     |     |     |     |                                   |     |     |                  |
| Passed           |     | Failed                                                     |     |     | Untested           |     |     |     |     |     | N                                 | A   | B   | Total Test Cases |
| 1                |     | 0                                                          |     |     | 0                  |     |     |     |     |     | 1                                 | 0   | 0   | 1                |

#### Sheet

| #   |           |                                              |     |                  |     | C-APP-13-01 |
| --- | --------- | -------------------------------------------- | --- | ---------------- | --- | ----------- |
| 1   | Condition | Precondition                                 |     |                  |     |             |
| 2   |           |                                              |     | C-APP-06-02      |     | O           |
| 3   |           | AttributeID                                  |     |                  |     |             |
| 4   |           |                                              |     | Existing (Size)  |     | O           |
| 5   |           | AttributeValueID                             |     |                  |     |             |
| 6   |           |                                              |     | Existing (Medium)|     | O           |
| 7   | Confirm   | Return                                       |     |                  |     |             |
| 8   |           | Error                                        |     |                  |     |             |
| 9   |           |                                              |     | Nil              |     | O           |
| 10  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                  |     | N           |
| 11  |           | Passed/Failed                                |     |                  |     | P           |
| 12  |           | Executed Date                                |     |                  |     | 2025-12-10  |
| 13  |           | Defect ID                                    |     |                  |     |             |
