package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/russross/blackfriday.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const (
		NoExtensions           Extensions = 0
		NoIntraEmphasis        Extensions = 1 << iota // Ignore emphasis markers inside words
		Tables                                        // Render tables
		FencedCode                                    // Render fenced code blocks
		Autolink                                      // Detect embedded URLs that are not explicitly marked
		Strikethrough                                 // Strikethrough text using ~~test~~
		LaxHTMLBlocks                                 // Loosen up HTML block parsing rules
		SpaceHeadings                                 // Be strict about prefix heading rules
		HardLineBreak                                 // Translate newlines into line breaks
		TabSizeEight                                  // Expand tabs to eight spaces instead of four
		Footnotes                                     // Pandoc-style footnotes
		NoEmptyLineBeforeBlock                        // No need to insert an empty line to start a (code, quote, ordered list, unordered list) block
		HeadingIDs                                    // specify heading IDs  with {#id}
		Titleblock                                    // Titleblock ala pandoc
		AutoHeadingIDs                                // Create the heading ID from the text
		BackslashLineBreak                            // Translate trailing backslashes into line breaks
		DefinitionLists                               // Render definition lists

		CommonHTMLFlags HTMLFlags = UseXHTML | Smartypants |
			SmartypantsFractions | SmartypantsDashes | SmartypantsLatexDashes

		CommonExtensions Extensions = NoIntraEmphasis | Tables | FencedCode |
			Autolink | Strikethrough | SpaceHeadings | HeadingIDs |
			BackslashLineBreak | DefinitionLists
	)
	dat, err := ioutil.ReadFile("./markdown/file_1.md")
	check(err)
	// fmt.Print(string(dat))
	// fmt.Println('\n')

	output := blackfriday.Run(dat, blackfriday.WithExtensions(1))
	fmt.Print(string(output))
}
