package nlp_test

import (
	"fmt"

	"github.com/sjramanathan/go-foundation/nlp"
)

func ExampleTokenize() {
	text := "Who's on first?"
	tokens := nlp.Tokenize(text)
	fmt.Println(tokens)

	// output:
	// [who on first]
}
