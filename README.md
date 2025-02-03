<p align="center">
  <img src="https://arematics.com/assets/banner/full_banner_transparent.png" width="380" height="144" alt="Arematics Banner">
</p>

# Test Initialization Utility

This repository provides a utility function to streamline the initialization and execution of [Godog](https://github.com/cucumber/godog) BDD (Behavior Driven Development) tests in Go projects. It introduces a reusable `InitializeTests` function to handle boilerplate code for running feature tests with dynamic steps and paths.

## Features

- **Dynamic Scenario Initialization:** Clean separation of initialization logic for each package.
- **Improved Usability:** Reduces repetitive setup code for tests.
- **Colored Output Support:** Provides clear and colored test summaries.

## Getting Started

### Installation

To use this utility, first ensure you have the required dependencies:

```bash
go get -t github.com/Arematics/go-tests
```

### Usage Example

Here’s a basic structure on how to use the `InitializeTests` function:

```go
package tests

import (
  "github.com/Arematics/go-tests"
	"github.com/cucumber/godog"
	"testing"
)

func TestAnyFeatures(t *testing.T) {
	tests.InitializeTests(t, InitializeScenario, "./example.feature", "Example Suite")
}

// InitializeScenario defines the setup for your feature's steps.
func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I have (\d+) apples$`, func(apples int) error {
		// Step implementation here
		return nil
	})
	ctx.Step(`^I eat (\d+) apples$`, func(apples int) error {
		// Step implementation here
		return nil
	})
}
```

### Function Signature

```go
func InitializeTests(t *testing.T, initializer func(ctx *godog.ScenarioContext), featurePath string, suiteName string)
```

**Parameters:**

- `t *testing.T`: Active test reference to link the BDD suite to a test run.
- `initializer func(ctx *godog.ScenarioContext)`: Scenario initializer to register feature-specific steps dynamically.
- `featurePath string`: Path to the `.feature` file containing the scenarios to execute.
- `suiteName string`: A name for your test suite to make results identifiable.

### Benefits

The `InitializeTests` function simplifies Godog-based projects with:

- Separation of test logic and step initialization.
- Easy reusability and reduced code duplication.

## Code of Conduct

This project is governed by the [MIT License](LICENSE)
