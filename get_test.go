package env_test

import (
	"os"
	"testing"

	"github.com/gomodrepo/env"
)

func TestGet(t *testing.T) {
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
			desc:           "#02",
			setKey:         _testKey,
			setValue:       _testValue,
			inKey:          _testKey,
			inDefaultValue: _defaultValue,
			wantValue:      _testValue,
		},
		{
			desc:           "#03",
			setKey:         _testKey,
			setValue:       _emptyValue,
			inKey:          _testKey,
			inDefaultValue: _defaultValue,
			wantValue:      _emptyValue,
		},
	}

	for _, s := range scenarios {
		t.Run("Get", func(t *testing.T) {
			backup, ok := os.LookupEnv(s.setKey)
			defer func() {
				if ok {
					os.Setenv(s.setKey, backup)
				} else {
					os.Unsetenv(s.setKey)
				}
			}()

			os.Setenv(s.setKey, s.setValue)

			got := env.Get(s.inKey, s.inDefaultValue)
			if got != s.wantValue {
				t.Errorf("%v: got '%v' want '%v'", s.desc, got, s.wantValue)
			}
		})
	}
}
