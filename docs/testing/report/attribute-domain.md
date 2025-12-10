### `domain.NewAttribute`

#### Meta

|                  |     |                                                                    |     |     |                    |     |     |     |     | 0   | 1                   | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------------------------------ | --- | --- | ------------------ | --- | --- | --- | --- | --- | ------------------- | --- | --- | ---------------- |
| Function Code    |     | A-NA-01                                                            |     |     | Function Name      |     |     |     |     |     | domain.NewAttribute |     |     |                  |
| Created By       |     | CodeCompanion                                                      |     |     | Executed By        |     |     |     |     |     |                     |     |     |                  |
| Lines of code    |     | 15                                                                 |     |     | Lack of test cases |     |     |     |     |     | 0                   |     |     |                  |
| Test requirement |     | Validate Attribute code and name boundary values (min, max, empty) |     |     |                    |     |     |     |     |     |                     |     |     |                  |
| Passed           |     | Failed                                                             |     |     | Untested           |     |     |     |     |     | N                   | A   | B   | Total Test Cases |
| 13               |     | 0                                                                  |     |     | 0                  |     |     |     |     |     | 1                   | 3   | 9   | 13               |

#### Sheet

| #   |           |                                              |     |                                                                                 |     | A-NA-01-01 | A-NA-01-02 | A-NA-01-03 | A-NA-01-04 | A-NA-01-05 | A-NA-01-06 | A-NA-01-07 | A-NA-01-08 | A-NA-01-09 | A-NA-01-10 | A-NA-01-11 | A-NA-01-12 | A-NA-01-13 |
| --- | --------- | -------------------------------------------- | --- | ------------------------------------------------------------------------------- | --- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- |
| 1   | Condition | Code Length                                  |     |                                                                                 |     |            |            |            |            |            |            |            |            |            |            |            |            |            |
| 2   |           |                                              |     | 0 (empty)                                                                       |     |            |            |            |            |            |            |            |            |            |            |            | O          |            |
| 3   |           |                                              |     | 1 (min - 1)                                                                     |     | O          |            |            |            |            |            |            |            |            |            |            |            |            |
| 4   |           |                                              |     | 2 (min)                                                                         |     |            | O          |            |            |            |            |            |            |            |            |            |            |            |
| 5   |           |                                              |     | 3 (min + 1)                                                                     |     |            |            | O          |            |            |            |            |            |            |            |            |            |            |
| 6   |           |                                              |     | 50 (max)                                                                        |     |            |            |            | O          |            |            |            |            |            |            |            |            |            |
| 7   |           |                                              |     | 51 (max + 1)                                                                    |     |            |            |            |            | O          |            |            |            |            |            |            |            |            |
| 8   |           |                                              |     | Valid (normal)                                                                  |     |            |            |            |            |            | O          | O          | O          | O          | O          |            |            | O          |
| 9   |           | Name Length                                  |     |                                                                                 |     |            |            |            |            |            |            |            |            |            |            |            |            |            |
| 10  |           |                                              |     | 0 (empty)                                                                       |     |            |            |            |            |            |            |            |            |            |            |            |            | O          |
| 11  |           |                                              |     | 1 (min - 1)                                                                     |     |            |            |            |            |            | O          |            |            |            |            |            |            |            |
| 12  |           |                                              |     | 2 (min)                                                                         |     |            |            |            |            |            |            | O          |            |            |            |            |            |            |
| 13  |           |                                              |     | 3 (min + 1)                                                                     |     |            |            |            |            |            |            |            | O          |            |            |            |            |            |
| 14  |           |                                              |     | 100 (max)                                                                       |     |            |            |            |            |            |            |            |            | O          |            |            |            |            |
| 15  |           |                                              |     | 101 (max + 1)                                                                   |     |            |            |            |            |            |            |            |            |            | O          |            |            |            |
| 16  |           |                                              |     | Valid (normal)                                                                  |     | O          | O          | O          | O          | O          |            |            |            |            |            | O          | O          |            |
| 17  | Confirm   | Return                                       |     |                                                                                 |     |            |            |            |            |            |            |            |            |            |            |            |            |            |
| 18  |           | Code                                         |     |                                                                                 |     |            |            |            |            |            |            |            |            |            |            |            |            |            |
| 19  |           |                                              |     | Input Code                                                                      |     | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          |
| 20  |           | Name                                         |     |                                                                                 |     |            |            |            |            |            |            |            |            |            |            |            |            |            |
| 21  |           |                                              |     | Input Name                                                                      |     | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          |
| 22  |           | ID                                           |     |                                                                                 |     |            |            |            |            |            |            |            |            |            |            |            |            |            |
| 23  |           |                                              |     | Not Nil                                                                         |     | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          |
| 24  |           | Values                                       |     |                                                                                 |     |            |            |            |            |            |            |            |            |            |            |            |            |            |
| 25  |           |                                              |     | Empty Array                                                                     |     | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          | O          |
| 26  |           | Error                                        |     |                                                                                 |     |            |            |            |            |            |            |            |            |            |            |            |            |            |
| 27  |           |                                              |     | Key: 'Attribute.Code' Error:Field validation for 'Code' failed on the 'gte' tag |     | O          |            |            |            |            | O          |            |            |            |            |            | O          |            |
| 28  |           |                                              |     | Key: 'Attribute.Code' Error:Field validation for 'Code' failed on the 'lte' tag |     |            |            |            |            | O          |            |            |            |            |            |            |            |            |
| 29  |           |                                              |     | Key: 'Attribute.Name' Error:Field validation for 'Name' failed on the 'gte' tag |     |            |            |            |            |            | O          | O          |            |            |            |            |            | O          |
| 30  |           |                                              |     | Key: 'Attribute.Name' Error:Field validation for 'Name' failed on the 'lte' tag |     |            |            |            |            |            |            |            |            |            | O          |            |            |            |
| 31  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                                                                                 |     | B          | B          | B          | B          | B          | A          | B          | B          | B          | B          | N          | A          | A          |
| 32  |           | Passed/Failed                                |     |                                                                                 |     | P          | P          | P          | P          | P          | P          | P          | P          | P          | P          | P          | P          | P          |
| 33  |           | Executed Date                                |     |                                                                                 |     | 2025-12-10 | 2025-12-10 | 2025-12-10 | 2025-12-10 | 2025-12-10 | 2025-12-10 | 2025-12-10 | 2025-12-10 | 2025-12-10 | 2025-12-10 | 2025-12-10 | 2025-12-10 | 2025-12-10 |
| 34  |           | Defect ID                                    |     |                                                                                 |     |            |            |            |            |            |            |            |            |            |            |            |            |            |

### `domain.NewAttributeValue`

#### Meta

|                  |     |                                                        |     |     |                    |     |     |     |     | 0   | 1                        | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------------------ | --- | --- | ------------------ | --- | --- | --- | --- | --- | ------------------------ | --- | --- | ---------------- |
| Function Code    |     | A-NAV-01                                               |     |     | Function Name      |     |     |     |     |     | domain.NewAttributeValue |     |     |                  |
| Created By       |     | CodeCompanion                                          |     |     | Executed By        |     |     |     |     |     |                          |     |     |                  |
| Lines of code    |     | 10                                                     |     |     | Lack of test cases |     |     |     |     |     | 0                        |     |     |                  |
| Test requirement |     | Validate AttributeValue value boundary values          |     |     |                    |     |     |     |     |     |                          |     |     |                  |
| Passed           |     | Failed                                                 |     |     | Untested           |     |     |     |     |     | N                        | A   | B   | Total Test Cases |
| 7                |     | 0                                                      |     |     | 0                  |     |     |     |     |     | 1                        | 1   | 5   | 7                |

#### Sheet

| #   |           |                                              |     |                                                                                      |     | A-NAV-01-01 | A-NAV-01-02 | A-NAV-01-03 | A-NAV-01-04 | A-NAV-01-05 | A-NAV-01-06 | A-NAV-01-07 |
| --- | --------- | -------------------------------------------- | --- | ------------------------------------------------------------------------------------ | --- | ----------- | ----------- | ----------- | ----------- | ----------- | ----------- | ----------- |
| 1   | Condition | Value Length                                 |     |                                                                                      |     |             |             |             |             |             |             |             |
| 2   |           |                                              |     | 0 (empty)                                                                            |     | O           |             |             |             |             |             |             |
| 3   |           |                                              |     | 1 (min)                                                                              |     |             | O           |             |             |             |             |             |
| 4   |           |                                              |     | 2 (min + 1)                                                                          |     |             |             | O           |             |             |             |             |
| 5   |           |                                              |     | 50                                                                                   |     |             |             |             | O           |             |             |             |
| 6   |           |                                              |     | 100 (max)                                                                            |     |             |             |             |             | O           |             |             |
| 7   |           |                                              |     | 101 (max + 1)                                                                        |     |             |             |             |             |             | O           |             |
| 8   |           |                                              |     | Valid (normal)                                                                       |     |             |             |             |             |             |             | O           |
| 9   | Confirm   | Return                                       |     |                                                                                      |     |             |             |             |             |             |             |             |
| 10  |           | Value                                        |     |                                                                                      |     |             |             |             |             |             |             |             |
| 11  |           |                                              |     | Input Value                                                                          |     | O           | O           | O           | O           | O           | O           | O           |
| 12  |           | ID                                           |     |                                                                                      |     |             |             |             |             |             |             |             |
| 13  |           |                                              |     | Not Nil                                                                              |     | O           | O           | O           | O           | O           | O           | O           |
| 14  |           | Error                                        |     |                                                                                      |     |             |             |             |             |             |             |             |
| 15  |           |                                              |     | Key: 'AttributeValue.Value' Error:Field validation for 'Value' failed on 'gte' tag   |     | O           |             |             |             |             |             |             |
| 16  |           |                                              |     | Key: 'AttributeValue.Value' Error:Field validation for 'Value' failed on 'lte' tag   |     |             |             |             |             |             | O           |             |
| 17  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                                                                                      |     | A           | B           | B           | B           | B           | B           | N           |
| 18  |           | Passed/Failed                                |     |                                                                                      |     | P           | P           | P           | P           | P           | P           | P           |
| 19  |           | Executed Date                                |     |                                                                                      |     | 2025-12-10  | 2025-12-10  | 2025-12-10  | 2025-12-10  | 2025-12-10  | 2025-12-10  | 2025-12-10  |
| 20  |           | Defect ID                                    |     |                                                                                      |     |             |             |             |             |             |             |             |

### `domain.Attribute.Update`

#### Meta

|                  |     |                                           |     |     |                    |     |     |     |     | 0   | 1                       | 2   | 3   | 4                |
| ---------------- | --- | ----------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ----------------------- | --- | --- | ---------------- |
| Function Code    |     | A-U-01                                    |     |     | Function Name      |     |     |     |     |     | domain.Attribute.Update |     |     |                  |
| Created By       |     | CodeCompanion                             |     |     | Executed By        |     |     |     |     |     |                         |     |     |                  |
| Lines of code    |     | 5                                         |     |     | Lack of test cases |     |     |     |     |     | 0                       |     |     |                  |
| Test requirement |     | Validate Attribute name update logic      |     |     |                    |     |     |     |     |     |                         |     |     |                  |
| Passed           |     | Failed                                    |     |     | Untested           |     |     |     |     |     | N                       | A   | B   | Total Test Cases |
| 4                |     | 0                                         |     |     | 0                  |     |     |     |     |     | 4                       | 0   | 0   | 4                |

#### Sheet

| #   |           |                                              |     |                  |     | A-U-01-01 | A-U-01-02 | A-U-01-03 | A-U-01-04 |
| --- | --------- | -------------------------------------------- | --- | ---------------- | --- | --------- | --------- | --------- | --------- |
| 1   | Condition | Initial Name                                 |     |                  |     |           |           |           |           |
| 2   |           |                                              |     | "Original Name"  |     | O         | O         | O         |           |
| 3   |           |                                              |     | "Name"           |     |           |           |           | O         |
| 4   |           | Update Name                                  |     |                  |     |           |           |           |           |
| 5   |           |                                              |     | "" (empty)       |     | O         |           |           |           |
| 6   |           |                                              |     | "Original Name"  |     |           | O         |           |           |
| 7   |           |                                              |     | "Updated Name"   |     |           |           | O         |           |
| 8   |           |                                              |     | "name"           |     |           |           |           | O         |
| 9   | Confirm   | Return                                       |     |                  |     |           |           |           |           |
| 10  |           | Name                                         |     |                  |     |           |           |           |           |
| 11  |           |                                              |     | "Original Name"  |     | O         | O         |           |           |
| 12  |           |                                              |     | "Updated Name"   |     |           |           | O         |           |
| 13  |           |                                              |     | "name"           |     |           |           |           | O         |
| 14  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                  |     | N         | N         | N         | N         |
| 15  |           | Passed/Failed                                |     |                  |     | P         | P         | P         | P         |
| 16  |           | Executed Date                                |     |                  |     | 2025-12-10| 2025-12-10| 2025-12-10| 2025-12-10|
| 17  |           | Defect ID                                    |     |                  |     |           |           |           |           |

### `domain.Attribute.UpdateValue`

#### Meta

|                  |     |                                                 |     |     |                    |     |     |     |     | 0   | 1                            | 2   | 3   | 4                |
| ---------------- | --- | ----------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ---------------------------- | --- | --- | ---------------- |
| Function Code    |     | A-UV-01                                         |     |     | Function Name      |     |     |     |     |     | domain.Attribute.UpdateValue |     |     |                  |
| Created By       |     | CodeCompanion                                   |     |     | Executed By        |     |     |     |     |     |                              |     |     |                  |
| Lines of code    |     | 11                                              |     |     | Lack of test cases |     |     |     |     |     | 0                            |     |     |                  |
| Test requirement |     | Validate AttributeValue update logic and errors |     |     |                    |     |     |     |     |     |                              |     |     |                  |
| Passed           |     | Failed                                          |     |     | Untested           |     |     |     |     |     | N                            | A   | B   | Total Test Cases |
| 3                |     | 0                                               |     |     | 0                  |     |     |     |     |     | 2                            | 1   | 0   | 3                |

#### Sheet

| #   |           |                                              |     |                     |     | A-UV-01-01 | A-UV-01-02 | A-UV-01-03 |
| --- | --------- | -------------------------------------------- | --- | ------------------- | --- | ---------- | ---------- | ---------- |
| 1   | Condition | Value ID                                     |     |                     |     |            |            |            |
| 2   |           |                                              |     | Existing            |     | O          | O          |            |
| 3   |           |                                              |     | Non-existent        |     |            |            | O          |
| 4   |           | New Value                                    |     |                     |     |            |            |            |
| 5   |           |                                              |     | "" (empty)          |     |            | O          |            |
| 6   |           |                                              |     | "Green"             |     | O          |            | O          |
| 7   | Confirm   | Return                                       |     |                     |     |            |            |            |
| 8   |           | Value                                        |     |                     |     |            |            |            |
| 9   |           |                                              |     | "Green"             |     | O          |            |            |
| 10  |           |                                              |     | Original Value      |     |            | O          |            |
| 11  |           | Error                                        |     |                     |     |            |            |            |
| 12  |           |                                              |     | ErrNotFound         |     |            |            | O          |
| 13  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                     |     | N          | N          | A          |
| 14  |           | Passed/Failed                                |     |                     |     | P          | P          | P          |
| 15  |           | Executed Date                                |     |                     |     | 2025-12-10 | 2025-12-10 | 2025-12-10 |
| 16  |           | Defect ID                                    |     |                     |     |            |            |            |

### `domain.Attribute.RemoveValue`

#### Meta

|                  |     |                                              |     |     |                    |     |     |     |     | 0   | 1                            | 2   | 3   | 4                |
| ---------------- | --- | -------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | ---------------------------- | --- | --- | ---------------- |
| Function Code    |     | A-RV-01                                      |     |     | Function Name      |     |     |     |     |     | domain.Attribute.RemoveValue |     |     |                  |
| Created By       |     | CodeCompanion                                |     |     | Executed By        |     |     |     |     |     |                              |     |     |                  |
| Lines of code    |     | 11                                           |     |     | Lack of test cases |     |     |     |     |     | 0                            |     |     |                  |
| Test requirement |     | Validate AttributeValue removal logic        |     |     |                    |     |     |     |     |     |                              |     |     |                  |
| Passed           |     | Failed                                       |     |     | Untested           |     |     |     |     |     | N                            | A   | B   | Total Test Cases |
| 3                |     | 0                                            |     |     | 0                  |     |     |     |     |     | 3                            | 0   | 0   | 3                |

#### Sheet

| #   |           |                                              |     |                  |     | A-RV-01-01 | A-RV-01-02 | A-RV-01-03 |
| --- | --------- | -------------------------------------------- | --- | ---------------- | --- | ---------- | ---------- | ---------- |
| 1   | Condition | Initial Values Count                         |     |                  |     |            |            |            |
| 2   |           |                                              |     | 1                |     |            |            | O          |
| 3   |           |                                              |     | 3                |     | O          |            |            |
| 4   |           | Value ID                                     |     |                  |     |            |            |            |
| 5   |           |                                              |     | Existing         |     | O          |            | O          |
| 6   |           |                                              |     | Non-existent     |     |            | O          |            |
| 7   | Confirm   | Return                                       |     |                  |     |            |            |            |
| 8   |           | Values Length                                |     |                  |     |            |            |            |
| 9   |           |                                              |     | 0                |     |            |            | O          |
| 10  |           |                                              |     | 1                |     |            | O          |            |
| 11  |           |                                              |     | 2                |     | O          |            |            |
| 12  | Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                  |     | N          | N          | N          |
| 13  |           | Passed/Failed                                |     |                  |     | P          | P          | P          |
| 14  |           | Executed Date                                |     |                  |     | 2025-12-10 | 2025-12-10 | 2025-12-10 |
| 15  |           | Defect ID                                    |     |                  |     |            |            |            |
