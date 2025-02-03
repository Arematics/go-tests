package tests

import (
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"os"
	"testing"
)

var Test *testing.T

// InitializeTests is a generic utility to run Godog tests for any package.
// It reduces redundant initialization logic across packages.
func InitializeTests(t *testing.T, initializer func(ctx *godog.ScenarioContext), featurePath string, suiteName string) {
	Test = t
	options := &godog.Options{
		Format: "pretty",                  // Fancy output for test results
		Output: colors.Colored(os.Stdout), // Enable colored output
		Paths:  []string{featurePath},     // Path to the feature files
	}

	// Build the Godog test suite
	suite := godog.TestSuite{
		Name:                suiteName,
		ScenarioInitializer: initializer, // Initialize package-specific steps dynamically
		Options:             options,
	}

	// Run the test suite and fail the test execution on failure
	if status := suite.Run(); status != 0 {
		t.Fail()
	}
}
