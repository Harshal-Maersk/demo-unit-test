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
   - If the function already handles a case correctly (like empty orders in this example), simply write a test case to confirm that behavior.
   - If the function doesn't handle a case (e.g., inconsistent currency codes in the item list), it’s your responsibility as a tester to identify and report this gap to the developer.

3. **Write Thorough Tests**  
   For each scenario, include:
   - **Input**: What data is being tested? (e.g., empty list, zero quantity).
   - **Output**: What result do you expect? (e.g., total = 0).
   - **Edge Cases**: What unexpected but possible situations should the code handle gracefully? (e.g., `UnitPrice` is `nil`).

4. **Ensure the Code Matches the Tests**  
   - After writing test cases, run them against the function.
   - If the code passes all tests, it’s handling the scenarios correctly.
   - If a test fails, review the code and suggest necessary changes to the programmer.

5. **Iterate and Refine**  
   Testing isn’t just about covering known cases—it’s also about anticipating future changes. Once a function’s behavior is covered:
   - Think about what might change in future versions.
   - Make your tests robust enough to catch regressions (breaking old functionality while adding new features).

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
   - Ensures the program continues without errors in such cases.

2. **ComputeTotalWithoutZero**:
   - Throws an error if the order has no items.
   - Highlights how such edge cases could break functionality if not handled.

---

## Tests

Unit tests validate the behavior of the `ComputeTotalWithZero` and `ComputeTotalWithoutZero` functions. The tests cover:

1. **Valid Order**: A single item with a known quantity and unit price.
2. **Empty Order**:
   - For `ComputeTotalWithZero`: The total should be `0` without errors.
   - For `ComputeTotalWithoutZero`: An error should be raised with the message "order contains no items."

### Run Tests

To run the tests:

```bash
go test -v ./...
