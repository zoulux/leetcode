package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(adventureCamp(
		[]string{
			"xO->xO->xO", "xO->BKbWDH", "xO->BKbWDH", "BKbWDH->mWXW", "LSAYWW->LSAYWW", "oAibBvPdJ", "LSAYWW->u", "LSAYWW->LSAYWW",
		},
	))

	fmt.Println(adventureCamp(
		[]string{
			"", "Gryffindor->Slytherin->Gryffindor", "Hogwarts->Hufflepuff->Ravenclaw",
		},
	))
	fmt.Println(adventureCamp(
		[]string{
			"Alice->Dex", "", "Dex",
		},
	))

	fmt.Println(adventureCamp(
		[]string{
			"leet->code", "leet->code->Campsite->Leet", "leet->code->leet->courier",
		},
	))
	fmt.Println(adventureCamp(
		[]string{
			"leet->code", "leet->code->Campsite->Leet", "leet->code->leet->courier",
		},
	))

}

func adventureCamp(expeditions []string) int {
	initArr := strings.Split(expeditions[0], "->")
	m0 := make(map[string]bool)
	for _, v := range initArr {
		m0[v] = true
	}

	idx := -1
	ans := 0
	for i := 1; i < len(expeditions); i++ {
		arr := strings.Split(expeditions[i], "->")
		m1 := make(map[string]bool)
		for _, v := range arr {
			if v != "" {
				m1[v] = true
			}
		}

		s := 0
		for k := range m1 {
			if !m0[k] {
				s++
				m0[k] = true
			}
		}
		if s > ans {
			ans = s
			idx = i
		}
	}

	return idx
}
