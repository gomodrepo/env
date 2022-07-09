package env_test

import (
	"os"
	"strings"
	"testing"

	"github.com/gomodrepo/env"
)

func TestMustGetIn(t *testing.T) {
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
			wantPanic: true,
		},
		{
			desc:               "#03",
			setKey:             _testKey,
			setValue:           _testValue,
			inKey:              _testKey,
			inAdditionalParams: []string{"testVal", "testValues"},
			wantPanic:          true,
		},
		{
			desc:               "#04",
			setKey:             _testKey,
			setValue:           _testValue,
			inKey:              _testKey,
			inAdditionalParams: []string{"testVal", "testValues", _testValue},
			wantValue:          _testValue,
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
			wantPanic:          true,
		},
		{
			desc:               "#07",
			setKey:             _testKey,
			setValue:           _emptyValue,
			inKey:              _testKey,
			inAdditionalParams: []string{"testVal", "testValues", _emptyValue},
			wantValue:          _emptyValue,
		},
	}

	for _, s := range scenarios {
		t.Run("MustGetIn", func(t *testing.T) {
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

			got := env.MustGetIn(s.inKey, s.inAdditionalParams...)
			if got != s.wantValue {
				t.Errorf("%v: got '%v' want '%v'", s.desc, got, s.wantValue)
			}
		})
	}
}
