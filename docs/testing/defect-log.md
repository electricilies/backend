## Defects

| Defect ID  | Module       | Description                                                                                                                                                                                                                                                                               | Type         | Severity | Priority | Status | Created Date |
| ---------- | ------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------ | -------- | -------- | ------ | ------------ |
| DF-O-VA-01 | Order Domain | ValidateOrderTotalAmount function incorrectly calculated total amount by only summing prices without multiplying by quantity. Fixed to multiply each item's price by its quantity before summing.<br/>Expected: TotalAmount = Sum(Price \* Quantity)<br/>Actual: TotalAmount = Sum(Price) | Coding Logic | Serious  | High     | Closed | 10-Dec-2025  |
| DF-O-IP-01 | Order Domain | IsPaid field had `validate:"required"` tag which caused validation to fail when the value is `false` (zero value for bool). Removed the 'required' validation since `false` is a valid value.<br/>Expected: IsPaid can be true or false<br/>Actual: Validation fails when IsPaid is false | Coding Logic | Medium   | Medium   | Closed | 10-Dec-2025  |

## Note

- Type: User Interface, Business Logic, Feature missing, Coding Logic
- Severity: Fatal, Serious, Medium, Cosmetic
- Priority: Immediately, High, Medium, Low
- Status: Open, Pending, Closed
- Created Date format: DD-MMM-YYYY

Ex:

| Defect ID | Module        | Description                                                                                                                                                                              | Type                                  | Severity                            | Priority                        | Status | Created Date                               |
| --------- | ------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------- | ----------------------------------- | ------------------------------- | ------ | ------------------------------------------ |
| 1         | Title section | <Describe defect and expected result>\nEx: Default value of Status field = blank is incorrect\nExpected result: when create new document, default value of [Status] field must be "Open" | <Type of defect>\nEx:\nUser Interface | <Severity of defect>\nEx:\nCosmetic | <Priority of defect>\nEx:\nHigh | Open   | <Created date of defect>\nEx:\n21-Sep-2008 |
