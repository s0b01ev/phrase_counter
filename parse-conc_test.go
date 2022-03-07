package main

import (
	"testing"
	"reflect"
	"strings"
	"bufio"
)

//func Init() {
//      phrases = append(phrases, M{})
//}	

func TestPhraseToStr(t *testing.T) {
	var phr [3][]rune
	r1 := []rune("aaa")
	r2 := []rune("bbb")
	r3 := []rune("ccc")
	phr[0] = r1 
	phr[1] = r2 
	phr[2] = r3 
	expected := "aaa bbb ccc"
	actual := phraseToStr(phr)
	if actual != expected {
		t.Errorf("expected: %v \nactual = %v \n", actual, expected)
	}	
}

func TestMergeMaps(t *testing.T) {
	map1 := map[string]int{"aaa": 2, "bbb": 2," ccc": 1, "ggg": 7}
	map2 := map[string]int{"aaa": 3, "bbb": 1," ccc": 1, "zzz": 9, "yyy": 3}
	map3 := map[string]int{"ddd": 2, "eee": 2," fff": 1}
	maps := []M{map1, map2, map3}
	expected := map[string]int{"aaa": 5,  "bbb": 3,  "ccc": 2,  "ddd": 2,  "eee": 2,  "fff": 1,  "ggg": 7,  "yyy": 3, "zzz": 9}
        actual := mergeMaps(maps)

	eq := reflect.DeepEqual(actual, expected)
	if  eq {
	     t.Errorf("Expected result of merging maps:\n %v is:\n%v\nactual:\n%v", maps, expected, actual)
	} 
}	


func TestProcessInput1(t *testing.T) {
	// type M map[string]int
	// var phrases []M
        phrases = append(phrases, M{})
	s := "aa'a bbb, ccc : ddd"
        sr := strings.NewReader(s)
	br := bufio.NewReader(sr)
	expected := map[string]int{"a'aa bbb ccc": 1, "bbb ccc ddd": 1}
	processInput(*br, 0)
	actual := phrases[0]
	eq := reflect.DeepEqual(actual, expected)
	if  eq {
		t.Errorf("expected: %v \nactual = %v \n", expected, actual)
	}

}

func TestProcessInput2(t *testing.T) {
	// type M map[string]int
	// var phrases []M
        phrases = append(phrases, M{})
	s := "aaa bbb ccc ddd aaa bbb ccc ddd"
        sr := strings.NewReader(s)
	br := bufio.NewReaderSize(sr, 64)
	expected := map[string]int{"aaa bbb ccc ddd": 2, "bbb ccc ddd": 2, "ccc ddd aaa": 1, "ddd aaa bbb": 1}
	processInput(*br, 1)
	actual := phrases[1]
	eq := reflect.DeepEqual(actual, expected)
	if  eq {
		t.Errorf("expected: %v \nactual = %v \n", expected, actual)
	}

}

func TestProcessInput3(t *testing.T) {
	// type M map[string]int
	// var phrases []M
        phrases = append(phrases, M{})
	s := "aaa bbb"
        sr := strings.NewReader(s)
	br := bufio.NewReaderSize(sr, 19)
	expected := map[string]int{}
	processInput(*br, 2)
	actual := phrases[2]
	eq := reflect.DeepEqual(actual, expected)
	if  eq {
		t.Errorf("expected: %v \nactual = %v \n", actual, expected)
	}

}

func TestProcessInput4(t *testing.T) {
	// type M map[string]int
	// var phrases []M
        phrases = append(phrases, M{})
	s := ",aaa -  bbb ;ccc"
        sr := strings.NewReader(s)
	br := bufio.NewReaderSize(sr, 19)
	expected := map[string]int{"aaa bbb ccc" : 1}
	processInput(*br, 3)
	actual := phrases[3]
	eq := reflect.DeepEqual(actual, expected)
	if  eq {
		t.Errorf("expected: %v \nactual = %v \n", actual, expected)
	}

}

func TestProcessInputUnicode(t *testing.T) {
	// type M map[string]int
	// var phrases []M
        phrases = append(phrases, M{})
	s := "อยาลางผลาญฤๅเขนฆาบฑาใคร ไมถอโทษโกรธแชงซดฮดฮดดา หดอภยเหมอนกฬาอชฌาสย"
        sr := strings.NewReader(s)
	br := bufio.NewReaderSize(sr, 19)
	expected := map[string]int{"อยาลางผลาญฤๅเขนฆาบฑาใคร ไมถอโทษโกรธแชงซดฮดฮดดา หดอภยเหมอนกฬาอชฌาสย" : 1}
	processInput(*br, 4)
	actual := phrases[4]
	eq := reflect.DeepEqual(actual, expected)
	if  eq {
		t.Errorf("expected: %v \nactual = %v \n", actual, expected)
	}

}
