package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultOrEnv(t *testing.T) {
	testCases := []struct {
		env, input, expected string
	}{
		{
			env:      "LOG_ANALYZER_FILE",
			input:    "log.txt",
			expected: "log.txt",
		},
		{
			env:      "LOG_ANALYZER_LEVEL",
			input:    "warning",
			expected: "warning",
		},
		{
			env:      "LOG_ANALYZER_OUTPUT",
			input:    "stats",
			expected: "stats",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.env, func(t *testing.T) {
			os.Setenv(tC.env, tC.input)
			s := defaultOrEnv(tC.env, "")
			os.Unsetenv(tC.env)
			assert.Equal(t, tC.expected, s)
		})
	}
}

func TestCoverage(t *testing.T) {
	testCases := []struct {
		level    string
		expected []string
	}{
		{
			level: "info",
			expected: []string{
				"info",
				"warning",
				"error",
			},
		},
		{
			level: "warning",
			expected: []string{
				"warning",
				"error",
			},
		},
		{
			level: "error",
			expected: []string{
				"error",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.level, func(t *testing.T) {
			list := coverage(tC.level)
			assert.Equal(t, tC.expected, list)
		})
	}
}

func TestAnalysis(t *testing.T) {
	testCases := []struct {
		path, level string
		expected    map[string]int
	}{
		{
			path:  "log.txt",
			level: "info",
			expected: map[string]int{
				"info":    114,
				"warning": 13,
				"error":   3,
			},
		},
		{
			path:  "log.txt",
			level: "warning",
			expected: map[string]int{
				"warning": 13,
				"error":   3,
			},
		},
		{
			path:  "log.txt",
			level: "error",
			expected: map[string]int{
				"error": 3,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.level, func(t *testing.T) {
			c, err := analysis(tC.path, tC.level)
			assert.Nil(t, err)
			assert.Equal(t, tC.expected, c)
		})
	}
}

func TestPrepareOutput(t *testing.T) {
	testCases := []struct {
		counter            map[string]int
		filename, expected string
	}{
		{
			counter: map[string]int{
				"info":    114,
				"warning": 13,
				"error":   3,
			},
			filename: "test1",
			expected: `Statistics of the test1

|  Level   |  Count   |
| -------- | -------- |
| info     | 114      |
| warning  | 13       |
| error    | 3        |
`,
		},
		{
			counter: map[string]int{
				"warning": 13,
				"error":   3,
			},
			filename: "test2",
			expected: `Statistics of the test2

|  Level   |  Count   |
| -------- | -------- |
| warning  | 13       |
| error    | 3        |
`,
		},
		{
			counter: map[string]int{
				"error": 3,
			},
			filename: "test3",
			expected: `Statistics of the test3

|  Level   |  Count   |
| -------- | -------- |
| error    | 3        |
`,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.filename, func(t *testing.T) {
			str := prepareOutput(tC.counter, tC.filename)
			assert.Equal(t, tC.expected, str)
		})
	}
}
