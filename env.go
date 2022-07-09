// Package env implements a environment variable retrieval utility.
package env

import (
	"os"
	"regexp"
	"strings"
)

// Get returns the environment variable set in 'key'.
// If value is not set for 'key', it returns 'defaultValue'.
func Get(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return value
}

// GetIn returns the environment variable set in 'key'.
// If value is not set for 'key' or different from 'in', it returns 'defaultValue'.
// 'in' is case sensitive.
func GetIn(key, defaultValue string, in ...string) string {
	value := Get(key, defaultValue)
	if value == defaultValue {
		return defaultValue
	}

	for _, v := range in {
		if value == v {
			return value
		}
	}

	return defaultValue
}

// GetInCaseInsensitive returns the environment variable set to 'key'.
// If value is not set for 'key' or different from 'in', it returns 'defaultValue'.
// 'in' is not case sensitive.
func GetInCaseInsensitive(key, defaultValue string, in ...string) string {
	value := Get(key, defaultValue)
	if value == defaultValue {
		return defaultValue
	}

	for _, v := range in {
		if strings.ToLower(value) == strings.ToLower(v) {
			return value
		}
	}

	return defaultValue
}

// GetInRegex returns the environment variable set to 'key'.
// If value is not set for 'key' or does not match the regular expression 'regex', it returns 'defaultValue'.
func GetInRegex(key, defaultValue string, regex ...string) string {
	value := Get(key, defaultValue)
	if value == defaultValue {
		return defaultValue
	}

	for _, r := range regex {
		re, err := regexp.Compile(r)
		if err != nil {
			return defaultValue
		}

		if re.MatchString(value) {
			return value
		}
	}

	return defaultValue
}

// GetExcept returns the environment variable set to 'key'.
// If value is not set for 'key' or equal to 'except', it returns 'defaultValue'.
// 'except' is case sensitive.
func GetExcept(key, defaultValue string, except ...string) string {
	value := Get(key, defaultValue)
	if value == defaultValue {
		return defaultValue
	}

	for _, v := range except {
		if value == v {
			return defaultValue
		}
	}

	return value
}

// GetExceptCaseInsensitive returns the environment variable set to 'key'.
// If value is not set for 'key' or equal to 'except', it returns 'defaultValue'.
// 'except' is not case sensitive.
func GetExceptCaseInsensitive(key, defaultValue string, except ...string) string {
	value := Get(key, defaultValue)
	if value == defaultValue {
		return defaultValue
	}

	for _, v := range except {
		if strings.ToLower(value) == strings.ToLower(v) {
			return defaultValue
		}
	}

	return value
}

// GetExceptRegex returns the environment variable set to 'key'.
// If value is not set for 'key' or matches the regular expression 'regex', it returns 'defaultValue'.
func GetExceptRegex(key, defaultValue string, regex ...string) string {
	value := Get(key, defaultValue)
	if value == defaultValue {
		return defaultValue
	}

	for _, r := range regex {
		re, err := regexp.Compile(r)
		if err != nil {
			return defaultValue
		}

		if re.MatchString(value) {
			return defaultValue
		}
	}

	return value
}

// MustGet returns the environment variable set to 'key'.
// If value is not set for 'key', it raises a panic.
func MustGet(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic("env: can not find key: " + key)
	}

	return value
}

// MustGetIn returns the environment variable set in 'key'.
// If value is not set for 'key' or different from 'in', it raises a panic.
// 'in' is case sensitive.
func MustGetIn(key string, in ...string) string {
	value := MustGet(key)

	for _, v := range in {
		if value == v {
			return value
		}
	}

	panic("env: value is not in: " + key)
}

// MustGetInCaseInsensitive returns the environment variable set to 'key'.
// If value is not set for 'key' or different from 'in', it raises a panic.
// 'in' is not case sensitive.
func MustGetInCaseInsensitive(key string, in ...string) string {
	value := MustGet(key)

	for _, v := range in {
		if strings.ToLower(value) == strings.ToLower(v) {
			return value
		}
	}

	panic("env: value is not in: " + key)
}

// MustGetInRegex returns the environment variable set to 'key'.
// If value is not set for 'key' or does not match the regular expression 'regex', it raises a panic.
func MustGetInRegex(key string, regex ...string) string {
	value := MustGet(key)

	for _, r := range regex {
		re, err := regexp.Compile(r)
		if err != nil {
			panic("env: failed to compile regex: " + key)
		}

		if re.MatchString(value) {
			return value
		}
	}

	panic("env: value is not in: " + key)
}

// MustGetExcept returns the environment variable set to 'key'.
// If value is not set for 'key' or equal to 'except', it raises a panic.
// 'except' is case sensitive.
func MustGetExcept(key string, except ...string) string {
	value := MustGet(key)

	for _, v := range except {
		if value == v {
			panic("env: value is not except: " + key)
		}
	}

	return value
}

// MustGetExceptCaseInsensitive returns the environment variable set to 'key'.
// If value is not set for 'key' or equal to 'except', it raises a panic.
// 'except' is not case sensitive.
func MustGetExceptCaseInsensitive(key string, except ...string) string {
	value := MustGet(key)

	for _, v := range except {
		if strings.ToLower(value) == strings.ToLower(v) {
			panic("env: value is not except: " + key)
		}
	}

	return value
}

// MustGetExceptRegex returns the environment variable set to 'key'.
// If value is not set for 'key' or matches the regular expression 'regex', it raises a panic.
func MustGetExceptRegex(key string, regex ...string) string {
	value := MustGet(key)

	for _, r := range regex {
		re, err := regexp.Compile(r)
		if err != nil {
			panic("env: failed to compile regex: " + key)
		}

		if re.MatchString(value) {
			panic("env: value is not except: " + key)
		}
	}

	return value
}
