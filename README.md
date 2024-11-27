# Order Management System with Unit Testing in Go

This project demonstrates the implementation of an **Order Management System** in Go with a focus on **unit testing**. The system calculates the total cost of an order containing items with quantities and unit prices. Additionally, it showcases two variations of a `ComputeTotal` function to handle or fail for an empty order list, serving as an example of testing edge cases.

---

## What is Unit Testing?

Unit testing is a software testing method where individual units or components of a program are tested in isolation. The goal is to verify that each unit functions as expected under various scenarios.

### Writing Effective Test Cases

1. **Understand the Function**  
   Study the input, output, and behavior of the function in different scenarios. Ask yourself:
   - What is the function supposed to do in normal cases (e.g., valid inputs)?
   - What are potential edge cases (e.g., empty lists, zero quantities, or invalid data)?
   - Are there any constraints or assumptions made in the code (e.g., `UnitPrice` should never be `nil`)?

2. **Check if All Scenarios Are Covered**  
   Compare the function's behavior with expected outcomes for all possible scenarios:
   - If the function already handles a case correctly (like empty orders in this example), write a test case to confirm that behavior.
   - If the function doesn't handle a case (e.g., inconsistent currency codes in the item list), write tests to highlight these gaps.

3. **Write Thorough Tests**  
   For each scenario, include:
   - **Input**: The data being tested (e.g., empty list, zero quantity).
   - **Output**: The expected result (e.g., total = 0, or error message).
   - **Edge Cases**: Unexpected but possible situations (e.g., `nil` prices or missing items).

4. **Ensure the Code Matches the Tests**  
   - After writing test cases, run them against the function.
   - If a test fails, review the code and suggest necessary changes.

5. **Iterate and Refine**  
   Testing isn’t just about current cases—it’s about anticipating future changes. Write tests robust enough to catch regressions when updating the codebase.

---

## How to Write Unit Tests in Go

Go has a built-in `testing` package to support unit testing. Here’s how it works:

1. **Test Functions**: Test functions must start with `Test` and accept a single argument of type `*testing.T`.
2. **Assertions**: Use libraries like [Testify](https://github.com/stretchr/testify) for readable assertions, such as `assert.NoError` and `assert.Equal`.
3. **Run Tests**: Execute `go test ./...` from the root of the project to run all tests.

---

## Project Overview

This project defines:
- **Order**: Represents an order with a unique ID, currency, and a list of items.
- **Item**: Represents an item with a unique ID, quantity, and unit price.
- **ComputeTotalWithZero**: Returns a total of `0` for an empty order without raising an error.
- **ComputeTotalWithoutZero**: Returns an error if the order list is empty.

### Functions:

1. **ComputeTotalWithZero**:  
   - Returns a total of `0` if the order has no items.  
   - Ensures the program continues gracefully without errors in such cases.

2. **ComputeTotalWithoutZero**:  
   - Throws an error if the order has no items.  
   - Demonstrates how edge cases like empty orders should be explicitly handled.

---

## Unit Tests

Unit tests validate the behavior of the `ComputeTotalWithZero` and `ComputeTotalWithoutZero` functions. The tests cover:

### Scenarios:

1. **Valid Order**:  
   - Single item with a known quantity and unit price.  
   - Verifies total calculation accuracy.

2. **Empty Order**:  
   - For `ComputeTotalWithZero`: Verifies that the total is `0` without errors.  
   - For `ComputeTotalWithoutZero`: Ensures an error is raised with the message "order contains no items."

### Example: Intentionally Failing Test

For demonstration purposes, you can modify `ComputeTotalWithoutZero` to return a `0` total for empty orders instead of an error. The corresponding test case would fail, highlighting the incorrect behavior.

#### Example of a Failing Test Output:

```text
=== RUN   TestComputeTotalWithoutZero_EmptyOrder
    order_test.go:61:
                Error Trace:    order_test.go:61
                Error:          An error is expected but got nil.
FAIL


### Run Tests

To run the tests:

bash
go test -v ./...