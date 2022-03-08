package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

// main "phrase: number" map
var phrases = map[string]int{}

// map sort-by-value
type Pair struct {
	Key   string
	Value int
}

type Pairs []Pair

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Pairs) Less(i, j int) bool { return p[i].Value > p[j].Value }

// END of map sort-by-value

func sortByValue(phrases map[string]int) {
	m := len(phrases)
	p := make(Pairs, m)
	i := 0
	for k, v := range phrases {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)

	if m > 100 {
		m = 100
	}
	for k := 0; k < m; k++ {
		fmt.Printf("%v - %v\n", p[k].Key, p[k].Value)
	}
}

func phraseToStr(phr [3][]rune) string {
	words := []string{}
	for i := 0; i <= 2; i++ {
		words = append(words, strings.ToLower(string(phr[i])))
	}
	return strings.Join(words, " ")
}

func updatePhrases(phr string) {
	if _, ok := phrases[phr]; ok {
		phrases[phr]++
	} else {
		phrases[phr] = 1
	}

}

func processInput(r bufio.Reader) {
	var word []rune
	var phrase [3][]rune
	var pi int
	tmpPunctAdd := false // we add "." "'" "`" temporarily

	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			if unicode.IsPunct(c) || unicode.IsSpace(c) { //
				if len(word) != 0 { // no onfinished word
					if unicode.IsPunct(c) || (unicode.IsSpace(c) && tmpPunctAdd) {
						if !tmpPunctAdd { // 1st punct char
							word = append(word, c)
							tmpPunctAdd = true
							continue
						} else { // 2nd punct char in a row,
							word = word[:len(word)-1] // delete previous punct/space
						}
					}
					phrase[pi] = word
					if pi < 2 {
						pi++
					} else { // phrase completed
						updatePhrases(phraseToStr(phrase))
						phrase[0] = phrase[1]
						phrase[1] = phrase[2]
					}
					tmpPunctAdd = false
					word = nil
				} // if len(word) != 0

			} else if unicode.IsLetter(c) || unicode.IsDigit(c) { //
				word = append(word, c)
				tmpPunctAdd = false
			}
		}
	}

	return
}

func main() {
	if len(os.Args) == 1 { // stdin
		r := bufio.NewReader(os.Stdin)
		processInput(*r)
	} else {
		for a := 1; a < len(os.Args); a++ {
			f, _ := os.Open(os.Args[a])
			r := bufio.NewReader(f)
			processInput(*r)
		}
	}
	sortByValue(phrases)
}
