## Guidance of blackbox testing

- In the integration test, there are tests with lifecycles, the Inter-test case should note previous test case. And the precondition should mention spinning up which containers, the DB need seed or not.
- The test case description should be in word instead of code style

### Meta

- Precondition:
  - Container database: up, seeded
  - Container redis: up

| Pass | Fail | Untested | N/A | Number of Test cases |
| ---- | ---- | -------- | --- | -------------------- |
| 0    | 0    | 10       | 0   | 10                   |

### Sheet

| #   | ID      | Test Case Description                             | Test Case Procedure                     | Expected Output                                                                    | Inter-test case Dependence | Result   | Test date | Test data                                                    | Note |
| --- | ------- | ------------------------------------------------- | --------------------------------------- | ---------------------------------------------------------------------------------- | -------------------------- | -------- | --------- | ------------------------------------------------------------ | ---- |
| -   |         | TestAttributeLifecycle (Test suite function name) |                                         |                                                                                    |                            |          |           |                                                              |      |
| -   | A-AL-01 | Create first attribute                            | - create attribute with `attributeData` | - Attribute successfully created and has the same name, code with the create param |                            | Untested |           | - `attributeData`:<br/> - Code: `color`<br/> - Name: `Color` |      |

NOTE:

- A-AL-01: Mean Attribute - Attribute Lifecycle - 01
- Result: Pass, Fail, Untested, N/A
