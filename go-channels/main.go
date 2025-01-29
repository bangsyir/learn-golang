package main

import (
	"flag"
	"fmt"
	"go-channels/gemini"
	"go-channels/grok"
)

func main() {
	fromPtr := flag.String("from", "grok", "the learning source")

	flag.Parse()

	fromSource := *fromPtr

	fmt.Println("from", fromSource)

	switch fromSource {
	case "gemini":
		gemini.FromGemini()
	default:
		grok.FromGrok()
	}
}
