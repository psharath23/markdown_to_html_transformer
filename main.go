package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"markdowntohtml.com/transformer"
)

func main() {

	argsWithoutProg := os.Args[1:]

	markdownFile, err := os.ReadFile(argsWithoutProg[0])
	if err != nil {
		log.Fatal(err)
	}

	htmlTransformer, err := transformer.GetTransformer("markdown", "html")
	if err != nil {
		log.Fatal(err)
	}

	htmlTransformer.AddInlineRule(transformer.HyperlinkRegex, transformer.GetHyperLink)
	htmlTransformer.AddLineRule(transformer.HeadingRegex, transformer.GetHeading)
	htmlTransformer.AddDefaultLineRule(transformer.GetParagraph)
	htmlTransformer.AddBlankLineRule(transformer.GetBlankLine)

	html := htmlTransformer.Transform(string(markdownFile))

	f, err := os.Create(fmt.Sprintf("%s.html", strings.Split(argsWithoutProg[0], ".")[0]))
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(html)
	f.Close()
}
