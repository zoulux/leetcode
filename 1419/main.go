package main

import "fmt"

func main() {

	fmt.Println(minNumberOfFrogs("crocakcroraoakk")) // test 30
	fmt.Println(minNumberOfFrogs("croakcroa"))
	fmt.Println(minNumberOfFrogs("croakcroak"))
	fmt.Println(minNumberOfFrogs("crcoakroak"))
	fmt.Println(minNumberOfFrogs("croakcrook"))
	fmt.Println(minNumberOfFrogs("cccrorakrcoakorakoak"))

}

func minNumberOfFrogs(croakOfFrogs string) int {

	var previous = [...]byte{
		'c': 'k',
		'r': 'c',
		'o': 'r',
		'a': 'o',
		'k': 'a',
	}

	counter := [len(previous)]int{} //croak 数量

	for _, c := range croakOfFrogs {
		p := previous[c]
		if counter[p] > 0 {
			counter[p]--
		} else if c != 'c' {
			return -1
		}
		counter[c]++
	}

	for _, v := range "croa" {
		if counter[v] > 0 {
			return -1
		}
	}
	return counter['k']
}

func minNumberOfFrogs3(croakOfFrogs string) int {

	var previous = map[rune]rune{
		'c': 'k',
		'r': 'c',
		'o': 'r',
		'a': 'o',
		'k': 'a',
	}

	counter := map[rune]int{} //croak 数量

	for _, c := range croakOfFrogs {
		p := previous[c]
		if counter[p] > 0 {
			counter[p]--
		} else if c != 'c' {
			return -1
		}
		counter[c]++
	}

	for _, v := range "croa" {
		if counter[v] > 0 {
			return -1
		}
	}

	return counter['k']
}

func minNumberOfFrogs2(croakOfFrogs string) int {
	counter := [5]int{} //croak 数量
	//croak
	target := map[byte]int{
		'c': 0,
		'r': 1,
		'o': 2,
		'a': 3,
		'k': 4,
	}

	var ans int

	for i := 0; i < len(croakOfFrogs); i++ {
		idx := target[croakOfFrogs[i]]
		if idx == 0 {
			if counter[4] != 0 {
				counter[4]--
			} else {
				ans++
			}
			counter[0]++
		} else {
			if counter[idx-1] != 0 {
				counter[idx-1]--
				counter[idx]++
			} else {
				return -1
			}
		}
	}

	if counter[4] == ans {
		return ans
	}
	return -1
}
