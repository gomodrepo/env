package env_test

import (
	"os"
	"testing"

	"github.com/gomodrepo/env"
)

func TestMustGet(t *testing.T) {
	scenarios := []testScenario{
		{
			desc:      "#00",
			inKey:     _testKey,
			wantPanic: true,
		},
		{
			desc:      "#01",
			setKey:    _testKey,
			setValue:  _testValue,
			inKey:     _testKey,
			wantValue: _testValue,
		},
		{
			desc:      "#02",
			setKey:    _testKey,
			setValue:  _emptyValue,
			inKey:     _testKey,
			wantValue: _emptyValue,
		},
	}

	for _, s := range scenarios {
		t.Run("MustGet", func(t *testing.T) {
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

			got := env.MustGet(s.inKey)
			if got != s.wantValue {
				t.Errorf("%v: got '%v' want '%v'", s.desc, got, s.wantValue)
			}
		})
	}
}
