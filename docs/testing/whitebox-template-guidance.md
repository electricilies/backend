## Guidance of whitebox testing

- This is an template for testing a function (whitebox)
- Each of the call in a test, calling a function should produce a single sheet for each
- For those value of the argument which aren't provided (use default value in go), don't need to mention
- The function name should be in the format of "Attribute.Create", "Category.Update",..., and the function code you can decide yourself, for ex "At-C-01" which mean attribute create 01, "C-UV-01" which mean Category Update value 01
- The defect ID should be mentioned if there is any defect found during the test, otherwise leave it blank. The ID should be in form of "DF-{Module}-{FunctionName}-01", for ex "DF-At-C-01" for Attribute Create function, "DF-P-UV-01" for Product Update variant function
- In the integration test, there are tests with lifecycles, the precondition should note previous function calls from another sheet if any by ID. And the precondition should mention spinning up which containers, the DB need seed or not.

## Template

### Meta

|                  |     |                                                                          |     |     |                    |     |     |     |     | 0   | 1          | 2   | 3   | 4                |
| ---------------- | --- | ------------------------------------------------------------------------ | --- | --- | ------------------ | --- | --- | --- | --- | --- | ---------- | --- | --- | ---------------- |
| Function Code    |     | Function1                                                                |     |     | Function Name      |     |     |     |     |     | Function A |     |     |                  |
| Created By       |     | <Developer Name>                                                         |     |     | Executed By        |     |     |     |     |     |            |     |     |                  |
| Lines of code    |     | 100                                                                      |     |     | Lack of test cases |     |     |     |     |     | -5         |     |     |                  |
| Test requirement |     | <Brief description about requirements which are tested in this function> |     |     |                    |     |     |     |     |     |            |     |     |                  |
| Passed           |     | Failed                                                                   |     |     | Untested           |     |     |     |     |     | N          | A   | B   | Total Test Cases |
| 1                |     | 1                                                                        |     |     | 5                  |     |     |     |     |     | 5          | 1   | 1   | 7                |

### Sheet

|           |                                              |     |                       |     | UTCID01             | UTCID02             | UTCID02             | UTCID02             | UTCID02             | UTCID02             | UTCID07             |
| --------- | -------------------------------------------- | --- | --------------------- | --- | ------------------- | ------------------- | ------------------- | ------------------- | ------------------- | ------------------- | ------------------- |
| Condition | Precondition                                 |     |                       |     |                     |                     |                     |                     |                     |                     |                     |
|           | a                                            |     |                       |     |                     |                     |                     |                     |                     |                     |                     |
|           |                                              |     | -2                    |     | O                   |                     |                     |                     |                     |                     |                     |
|           |                                              |     | -1                    |     |                     |                     |                     |                     |                     |                     | O                   |
|           |                                              |     | 0                     |     |                     | O                   | O                   | O                   |                     |                     |                     |
|           |                                              |     | 1                     |     |                     |                     |                     |                     | O                   | O                   |                     |
|           | b                                            |     |                       |     |                     |                     |                     |                     |                     |                     |                     |
|           |                                              |     | 0                     |     |                     | O                   | O                   |                     |                     |                     |                     |
|           |                                              |     | -2                    |     |                     |                     |                     |                     | O                   | O                   | O                   |
|           |                                              |     | 2                     |     |                     |                     |                     | O                   |                     |                     |                     |
|           | c                                            |     |                       |     |                     |                     |                     |                     |                     |                     |                     |
|           |                                              |     | 0                     |     |                     | O                   |                     |                     |                     |                     |                     |
|           |                                              |     | 1                     |     |                     |                     | O                   | O                   | O                   |                     |                     |
|           |                                              |     | 3                     |     |                     |                     |                     |                     |                     |                     | O                   |
|           |                                              |     | 5                     |     |                     |                     |                     |                     |                     | O                   |                     |
| Confirm   | Return                                       |     |                       |     |                     |                     |                     |                     |                     |                     |                     |
|           | list                                         |     |                       |     |                     |                     |                     |                     |                     |                     |                     |
|           |                                              |     |                       |     | O                   |                     | O                   |                     |                     | O                   |                     |
|           |                                              |     | size = 0              |     |                     | O                   |                     |                     |                     |                     |                     |
|           |                                              |     | {-1/2}                |     |                     |                     |                     | O                   |                     |                     |                     |
|           |                                              |     | {1,1}                 |     |                     |                     |                     |                     | O                   |                     |                     |
|           |                                              |     | {1,-3}                |     |                     |                     |                     |                     |                     |                     | O                   |
|           | Exception                                    |     |                       |     |                     |                     |                     |                     |                     |                     |                     |
|           |                                              |     |                       |     |                     |                     |                     |                     |                     |                     |                     |
|           | Log message                                  |     |                       |     |                     |                     |                     |                     |                     |                     |                     |
|           |                                              |     | "please input a>= -1" |     | O                   |                     |                     |                     |                     |                     |                     |
| Result    | Type(N : Normal, A : Abnormal, B : Boundary) |     |                       |     | A                   | N                   | N                   | N                   | N                   | N                   | B                   |
|           | Passed/Failed                                |     |                       |     | P                   | F                   |                     |                     |                     |                     |                     |
|           | Executed Date                                |     |                       |     | 2007-02-26 00:00:00 | 2007-02-26 00:00:00 | 2007-02-26 00:00:00 | 2007-02-26 00:00:00 | 2007-02-26 00:00:00 | 2007-02-26 00:00:00 | 2007-03-03 00:00:00 |
|           | Defect ID                                    |     |                       |     |                     | DF-M-FN-01          |                     |                     |                     |                     |                     |
