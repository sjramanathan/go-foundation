package nlp

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

func TestTokenize(t *testing.T) {
	text := "What's on second?"
	expected := []string{"what", "on", "second"}

	tokens := Tokenize(text)

	if !reflect.DeepEqual(expected, tokens) {
		t.Fatalf("Expected: %#v, Actual: %#v", expected, tokens)
	}
}

var testCases = []struct {
	text   string
	tokens []string
}{
	{"Who's on first?", []string{"who", "on", "first"}},
	{"", nil},
}

func TestTokenizeTable(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.text, func(t *testing.T) {
			tokens := Tokenize(tc.text)
			require.Equal(t, tc.tokens, tokens)
		})
	}
}

type tokenizeCase struct {
	text   string
	tokens []string
}

func loadTokenizeCases(t *testing.T) []tokenizeCase {
	data, err := ioutil.ReadFile("tokenize_cases.toml")
	var testCases struct {
		Cases []tokenizeCase
	}

	err = toml.Unmarshal(data, &testCases)
	require.NoError(t, err, "Unmarshal TOML")

	return testCases.Cases
}

// Exercise: Read test cases from tokenize_cases.toml
func TestTokenizeToml(t *testing.T) {
	tokenizeTestCases := loadTokenizeCases(t)

	fmt.Printf("%#v\n", tokenizeTestCases)

	for _, tc := range tokenizeTestCases {
		t.Run(tc.text, func(t *testing.T) {
			tokens := Tokenize(tc.text)
			require.Equal(t, tc.tokens, tokens)
		})
	}
}

func FuzzTokenize(f *testing.F) {
	f.Fuzz(func(t *testing.T, text string) {
		tokens := Tokenize(text)
		lText := strings.ToLower(text)

		for _, tok := range tokens {
			if !strings.Contains(lText, tok) {
				t.Fatal(tok)
			}
		}
	})
}
