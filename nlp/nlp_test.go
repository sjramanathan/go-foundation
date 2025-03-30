package nlp

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	text := "What's on second?"
	expected := []string{"what", "s", "on", "second"}
	tokens := Tokenize(text)

	if !reflect.DeepEqual(expected, tokens) {
		t.Fatalf("Expected: %#v, Actual: %#v", expected, tokens)
	}
}
