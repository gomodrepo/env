package env_test

import (
	"os"
	"strings"
	"testing"

	"github.com/gomodrepo/env"
)

func TestMustGetExceptCaseInsensitive(t *testing.T) {
	scenarios := []testScenario{
		{
			desc:      "#00",
			inKey:     _testKey,
			wantPanic: true,
		},
		{
			desc:               "#01",
			inKey:              _testKey,
			inAdditionalParams: []string{"testVal", "testValues"},
			wantPanic:          true,
		},
		{
			desc:      "#02",
			setKey:    _testKey,
			setValue:  _testValue,
			inKey:     _testKey,
			wantValue: _testValue,
		},
		{
			desc:               "#03",
			setKey:             _testKey,
			setValue:           _testValue,
			inKey:              _testKey,
			inAdditionalParams: []string{"testVal", "testValues"},
			wantValue:          _testValue,
		},
		{
			desc:               "#04",
			setKey:             _testKey,
			setValue:           _testValue,
			inKey:              _testKey,
			inAdditionalParams: []string{"testVal", "testValues", _testValue},
			wantPanic:          true,
		},
		{
			desc:               "#05",
			setKey:             _testKey,
			setValue:           _testValue,
			inKey:              _testKey,
			inAdditionalParams: []string{"testVal", "testValues", strings.ToUpper(_testValue)},
			wantPanic:          true,
		},
		{
			desc:               "#06",
			setKey:             _testKey,
			setValue:           _emptyValue,
			inKey:              _testKey,
			inAdditionalParams: []string{"testVal", "testValues"},
			wantValue:          _emptyValue,
		},
		{
			desc:               "#07",
			setKey:             _testKey,
			setValue:           _emptyValue,
			inKey:              _testKey,
			inAdditionalParams: []string{"testVal", "testValues", _emptyValue},
			wantPanic:          true,
		},
	}

	for _, s := range scenarios {
		t.Run("MustGetExceptCaseInsensitive", func(t *testing.T) {
			backup, ok := os.LookupEnv(s.setKey)
			defer func() {
				if ok {
					os.Setenv(s.setKey, backup)
				} else {
					os.Unsetenv(s.setKey)
				}

				p := recover()
				if (p == nil && s.wantPanic) || (p != nil && !s.wantPanic) {
					t.Errorf("%v: gotPanic '%v' wantPanic '%v'", s.desc, p, s.wantPanic)
				}
			}()

			os.Setenv(s.setKey, s.setValue)

			got := env.MustGetExceptCaseInsensitive(s.inKey, s.inAdditionalParams...)
			if got != s.wantValue {
				t.Errorf("%v: got '%v' want '%v'", s.desc, got, s.wantValue)
			}
		})
	}
}
