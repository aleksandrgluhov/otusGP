package hw02unpackstring

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type tests struct {
	input         string
	expectedValue string
	expectedError error
}

func TestUnpack(t *testing.T) {
	tests := []tests{
		{
			input:         "a4bc2d5e",
			expectedValue: "aaaabccddddde",
			expectedError: nil,
		},
		{
			input:         "abccd",
			expectedValue: "abccd",
			expectedError: nil,
		},
		{
			input:         "",
			expectedValue: "",
			expectedError: nil,
		},
		{
			input:         "aaa0b",
			expectedValue: "aab",
			expectedError: nil,
		},
		{
			input:         "d\n5abc",
			expectedValue: "d\n\n\n\n\nabc",
			expectedError: nil,
		},
		// uncomment if task with asterisk completed
		{
			input:         `qwe\4\5`,
			expectedValue: `qwe45`,
			expectedError: nil,
		},
		{
			input:         `qwe\45`,
			expectedValue: `qwe44444`,
			expectedError: nil,
		},
		{
			input:         `qwe\\5`,
			expectedValue: `qwe\\\\\`,
			expectedError: nil,
		},
		{
			input:         `qwe\\\3`,
			expectedValue: `qwe\3`,
			expectedError: nil,
		},
		{
			input:         `qwe\4a`,
			expectedValue: `qwe4a`,
			expectedError: nil,
		},
		{
			input:         `abc\\`,
			expectedValue: `abc\`,
			expectedError: nil,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expectedValue, result)
			require.Equal(t, tc.expectedError, nil)
		})
	}
}

func TestUnpackInvalidString(t *testing.T) {
	tests := []tests{
		{
			input:         "3abc",
			expectedValue: "",
			expectedError: ErrInvalidString,
		},
		{
			input:         "45",
			expectedValue: "",
			expectedError: ErrInvalidString,
		},
		{
			input:         "aaa10b",
			expectedValue: "",
			expectedError: ErrInvalidString,
		},
		{
			input:         "abc00",
			expectedValue: "",
			expectedError: ErrInvalidString,
		},
		// uncomment if task with asterisk completed
		{
			input:         `qw\ne`,
			expectedValue: "",
			expectedError: ErrInvalidEscape,
		},
		{
			input:         `ab\c`,
			expectedValue: "",
			expectedError: ErrInvalidEscape,
		},
		{
			input:         `abc\`,
			expectedValue: "",
			expectedError: ErrInvalidEscape,
		},
		{
			input:         `abc\\\`,
			expectedValue: "",
			expectedError: ErrInvalidEscape,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result, err := Unpack(tc.input)
			require.Truef(t, errors.Is(err, tc.expectedError), "actual error %q", err)
			require.Equal(t, tc.expectedValue, result)
		})
	}
}
