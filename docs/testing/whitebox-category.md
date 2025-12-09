# Whitebox Test Cases

## Category

### `domain.NewCategory`

#### Meta

|                  |     |                                                               |     |     |                    |     |     |     |     | 0   | 1                    | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | -------------------- | --- | --- | ---------------- |
| Function Code    |     | C-NC-01                                                       |     |     | Function Name      |     |     |     |     |     | Category.NewCategory |     |     |                  |
| Created By       |     | CodeCompanion                                                 |     |     | Executed By        |     |     |     |     |     |                      |     |     |                  |
| Lines of code    |     | 34                                                            |     |     | Lack of test cases |     |     |     |     |     | 0                    |     |     |                  |
| Test requirement |     | Validate Category name boundary values (min, max, empty, etc) |     |     |                    |     |     |     |     |     |                      |     |     |                  |
| Passed           |     | Passed                                                        |     |     | Untested           |     |     |     |     |     | Y                    | N   | B   | Total Test Cases |
| 1                |     | 1                                                             |     |     | 0                  |     |     |     |     |     | 5                    | 2   | 7   | 7                |

#### Sheet

|           |                                              |     |         |     | C-NC-01-01 | C-NC-01-02 | C-NC-01-03 | C-NC-01-04 | C-NC-01-05 | C-NC-01-06 | C-NC-01-07 |
| --------- | -------------------------------------------- | --- | ------- | --- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- |
| Condition | Name Length                                  |     |         |     |            |            |            |            |            |            |            |
|           |                                              |     | 0       |     |            |            |            |            |            |            | O          |
|           |                                              |     | 1       |     | O          |            |            |            |            |            |            |
|           |                                              |     | 2       |     |            | O          |            |            |            |            |            |
|           |                                              |     | 3       |     |            |            | O          |            |            |            |            |
|           |                                              |     | 99      |     |            |            |            | O          |            |            |            |
|           |                                              |     | 100     |     |            |            |            |            | O          |            |            |
|           |                                              |     | 101     |     |            |            |            |            |            | O          |            |
| Confirm   | Result                                       |     |         |     |            |            |            |            |            |            |            |
|           |                                              |     | Success |     |            | O          | O          | O          | O          |            |            |
|           |                                              |     | Error   |     | O          |            |            |            |            | O          | O          |
| Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |         |     | B          | B          | B          | B          | B          | B          | A          |
|           | Passed/Failed                                |     |         |     | Passed     | Passed     | Passed     | Passed     | Passed     | Passed     | Passed     |
|           | Executed Date                                |     |         |     | 2025-12-09 | 2025-12-09 | 2025-12-09 | 2025-12-09 | 2025-12-09 | 2025-12-09 | 2025-12-09 |
|           | Defect ID                                    |     |         |     |            |            |            |            |            |            |            |

### `domain.Category.Update`

#### Meta

|                  |     |                                                                      |     |     |                    |     |     |     |     | 0   | 1               | 2   | 3   | 4                |
| ---------------- | --- | -------------------------------------------------------------------- | --- | --- | ------------------ | --- | --- | --- | --- | --- | --------------- | --- | --- | ---------------- |
| Function Code    |     | C-UV-01                                                              |     |     | Function Name      |     |     |     |     |     | Category.Update |     |     |                  |
| Created By       |     | CodeCompanion                                                        |     |     | Executed By        |     |     |     |     |     |                 |     |     |                  |
| Lines of code    |     | 24                                                                   |     |     | Lack of test cases |     |     |     |     |     | 0               |     |     |                  |
| Test requirement |     | Validate Category name update boundary values (min, max, empty, etc) |     |     |                    |     |     |     |     |     |                 |     |     |                  |
| Passed           |     | Passed                                                               |     |     | Untested           |     |     |     |     |     | Y               | N   | B   | Total Test Cases |
| 1                |     | 1                                                                    |     |     | 0                  |     |     |     |     |     | 6               | 2   | 8   | 8                |

#### Sheet

|           |                                              |     |         |     | C-UV-01-01 | C-UV-01-02 | C-UV-01-03 | C-UV-01-04 | C-UV-01-05 | C-UV-01-06 | C-UV-01-07 | C-UV-01-08 |
| --------- | -------------------------------------------- | --- | ------- | --- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- | ---------- |
| Condition | Name Length                                  |     |         |     |            |            |            |            |            |            |            |            |
|           |                                              |     | 0       |     |            |            |            |            |            |            | O          |            |
|           |                                              |     | 1       |     | O          |            |            |            |            |            |            |            |
|           |                                              |     | 2       |     |            | O          |            |            |            |            |            |            |
|           |                                              |     | 3       |     |            |            | O          |            |            |            |            |            |
|           |                                              |     | 9       |     |            |            |            |            |            |            |            | O          |
|           |                                              |     | 99      |     |            |            |            | O          |            |            |            |            |
|           |                                              |     | 100     |     |            |            |            |            | O          |            |            |            |
|           |                                              |     | 101     |     |            |            |            |            |            | O          |            |            |
| Confirm   | Result                                       |     |         |     |            |            |            |            |            |            |            |            |
|           |                                              |     | Success |     |            | O          | O          | O          | O          |            | O          | O          |
|           |                                              |     | Error   |     | O          |            |            |            |            | O          |            |            |
| Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |         |     | B          | B          | B          | B          | B          | B          | A          | N          |
|           | Passed/Failed                                |     |         |     | Passed     | Passed     | Passed     | Passed     | Passed     | Passed     | Passed     | Passed     |
|           | Executed Date                                |     |         |     | 2025-12-09 | 2025-12-09 | 2025-12-09 | 2025-12-09 | 2025-12-09 | 2025-12-09 | 2025-12-09 | 2025-12-09 |
|           | Defect ID                                    |     |         |     |            |            |            |            |            |            |            |            |
