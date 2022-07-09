package env_test

const (
	_testKey      = "TEST_KEY"
	_testValue    = "testValue"
	_defaultValue = "defaultValue"
	_emptyValue   = ""
)

type testScenario struct {
	desc               string
	setKey             string
	setValue           string
	inKey              string
	inDefaultValue     string
	inAdditionalParams []string
	wantValue          string
	wantPanic          bool
}
