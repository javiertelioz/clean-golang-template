package say_hello

import (
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func InitializeTestSuite(ctx *godog.TestSuiteContext) {}

func InitializeScenario(ctx *godog.ScenarioContext) {
	helloContext := NewHelloFeatureContext()
	helloContext.InitializeScenario(ctx)
}

func TestMain(m *testing.M) {
	status := godog.TestSuite{
		Name:                 "hello",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options: &godog.Options{
			Format:              "pretty",
			Output:              colors.Colored(os.Stdout),
			Paths:               []string{"say_hello.feature"},
			Randomize:           -1,
			ShowStepDefinitions: false,
			NoColors:            false,
		},
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
