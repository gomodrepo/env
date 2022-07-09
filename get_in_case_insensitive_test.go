package env_test

import (
	"os"
	"strings"
	"testing"

	"github.com/gomodrepo/env"
)

func TestGetInCaseInsensitive(t *testing.T) {
	scenarios := []testScenario{
		{
			desc:           "#00",
			inKey:          _testKey,
			inDefaultValue: _defaultValue,
			wantValue:      _defaultValue,
		},
		{
			desc:           "#01",
			inKey:          _testKey,
			inDefaultValue: _emptyValue,
			wantValue:      _emptyValue,
		},
		{
			desc:               "#02",
			inKey:              _testKey,
			inDefaultValue:     _defaultValue,
			inAdditionalParams: []string{"testVal", "testValues"},
			wantValue:          _defaultValue,
		},
		{
			desc:               "#03",
			inKey:              _testKey,
			inDefaultValue:     _defaultValue,
			inAdditionalParams: []string{"testVal", "testValues", _defaultValue},
			wantValue:          _defaultValue,
		},
		{
			desc:               "#04",
			inKey:              _testKey,
			inDefaultValue:     _defaultValue,
			inAdditionalParams: []string{"testVal", "testValues", strings.ToUpper(_defaultValue)},
			wantValue:          _defaultValue,
		},
		{
			desc:           "#05",
			setKey:         _testKey,
			setValue:       _testValue,
			inKey:          _testKey,
			inDefaultValue: _defaultValue,
			wantValue:      _defaultValue,
		},
		{
			desc:               "#06",
			setKey:             _testKey,
			setValue:           _testValue,
			inKey:              _testKey,
			inDefaultValue:     _defaultValue,
			inAdditionalParams: []string{"testVal", "testValues"},
			wantValue:          _defaultValue,
		},
		{
			desc:               "#07",
			setKey:             _testKey,
			setValue:           _testValue,
			inKey:              _testKey,
			inDefaultValue:     _defaultValue,
			inAdditionalParams: []string{"testVal", "testValues", _defaultValue},
			wantValue:          _defaultValue,
		},
		{
			desc:               "#08",
			setKey:             _testKey,
			setValue:           _testValue,
			inKey:              _testKey,
			inDefaultValue:     _defaultValue,
			inAdditionalParams: []string{"testVal", "testValues", strings.ToUpper(_defaultValue)},
			wantValue:          _defaultValue,
		},
		{
			desc:               "#09",
			setKey:             _testKey,
			setValue:           _testValue,
			inKey:              _testKey,
			inDefaultValue:     _defaultValue,
			inAdditionalParams: []string{"testVal", "testValues", _testValue},
			wantValue:          _testValue,
		},
		{
			desc:               "#10",
			setKey:             _testKey,
			setValue:           _testValue,
			inKey:              _testKey,
			inDefaultValue:     _defaultValue,
			inAdditionalParams: []string{"testVal", "testValues", strings.ToUpper(_testValue)},
			wantValue:          _testValue,
		},
		{
			desc:               "#11",
			setKey:             _testKey,
			setValue:           _emptyValue,
			inKey:              _testKey,
			inDefaultValue:     _defaultValue,
			inAdditionalParams: []string{"testVal", "testValues"},
			wantValue:          _defaultValue,
		},
		{
			desc:               "#12",
			setKey:             _testKey,
			setValue:           _emptyValue,
			inKey:              _testKey,
			inDefaultValue:     _defaultValue,
			inAdditionalParams: []string{"testVal", "testValues", _emptyValue},
			wantValue:          _emptyValue,
		},
	}

	for _, s := range scenarios {
		t.Run("GetInCaseInsensitive", func(t *testing.T) {
			backup, ok := os.LookupEnv(s.setKey)
			defer func() {
				if ok {
					os.Setenv(s.setKey, backup)
				} else {
					os.Unsetenv(s.setKey)
				}
			}()

			os.Setenv(s.setKey, s.setValue)

			got := env.GetInCaseInsensitive(s.inKey, s.inDefaultValue, s.inAdditionalParams...)
			if got != s.wantValue {
				t.Errorf("%v: got '%v' want '%v'", s.desc, got, s.wantValue)
			}
		})
	}
}
