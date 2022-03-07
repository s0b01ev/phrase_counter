package main

import (
	"bufio"
        "fmt"
	"io"
	"os"
	"log"
	"unicode"
	"strings"
	"sort"
	"sync"
)	

// main "phrase: number" map
type M map[string]int
var phrases []M

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

// sort and print out top 100 
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

// converts phrase to string
func phraseToStr(phr [3][]rune) string {
	words := []string{}
	for i := 0; i <= 2; i++ {
		words = append(words, strings.ToLower(string(phr[i])))
	}	
	return strings.Join(words, " ")
}

// inserts new element or increments existing one
func updatePhrases(phr string, a int) {
	if _, ok := phrases[a][phr]; ok {
		phrases[a][phr]++
	} else {
		phrases[a][phr] = 1
        }		

}	
//main data processor
// r - data
// a - index to use when storing in phrases 
func processInput(r bufio.Reader, a int) {
	var word []rune
	var phrase [3][]rune
	var pi int
	tmpPunctAdd := false // we add "." "'" "`" temporarily 

	for  {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}	
		} else {
		        if unicode.IsPunct(c) || unicode.IsSpace(c) { // 
                            if len(word) != 0 {  // no onfinished word
				    if unicode.IsPunct(c) || (unicode.IsSpace(c) && tmpPunctAdd) { 
					    if ! tmpPunctAdd {  // 1st punct char
						   word = append(word, c)
						   tmpPunctAdd = true
						   continue
					    } else {		              // 2nd punct char in a row,  
						   word = word[:len(word)-1] // delete previous punct/space
					    }		   
			           }		   
				    phrase[pi] = word
				    if pi < 2 {
					    pi++
				    } else { // phrase completed	
					    updatePhrases(phraseToStr(phrase), a)				    
					    phrase[0] = phrase[1]
					    phrase[1] = phrase[2]
				    }
				    tmpPunctAdd = false
				    word = nil
			     } // if len(word) != 0

		        }  else if unicode.IsLetter(c) || unicode.IsDigit(c) { // 
			    word = append(word, c)
			    tmpPunctAdd = false
                        }
		}	
        }
	return
}

// merges several maps in one
func mergeMaps(maps []M) map[string]int {
	lenIdxMap := map[int]int{}
	lens:= make([]int, len(maps))
	for i := 0; i < len(maps); i++  {
		lenIdxMap[len(maps[i])] = i
	}	
	i := 0
	for k := range lenIdxMap{
	    lens[i] = k
	    i++
	}
	// keys now contains sorted array of mapa keys 
	// 1st element  - longest map, we will be merge all others into it
	sort.Sort(sort.Reverse(sort.IntSlice(lens)))
	
        i = 1
	idS := 0
        idT := lenIdxMap[lens[0]]
	for  {
		if i >= len(maps) {
			break
		}	
		idS = lenIdxMap[lens[i]]
		i++
		for ks, vs := range maps[idS] {
                       if vt, ok := maps[idT][ks]; ok {
			       maps[idT][ks] = vt + vs
                       } else {
			       maps[idT][ks] = vs
                       }
		}	
	}	
	return maps[idT]
}	

func main() {
	var wg sync.WaitGroup
	if len(os.Args) == 1 { // stdin
	        phrases = append(phrases, M{})
		r := bufio.NewReader(os.Stdin)
		processInput(*r, 0)
	} else { 	
	        wg.Add(len(os.Args)-1)
		for a := 1; a < len(os.Args); a++ {
	                phrases = append(phrases, M{})
			go func(a int) {
				defer wg.Done()
				f, _:= os.Open(os.Args[a])
				r := bufio.NewReader(f)
				processInput(*r, a-1)
			}(a)	
		}	
		wg.Wait()
	}	

	phrasesJoined := mergeMaps(phrases) 
	sortByValue(phrasesJoined)
}	
