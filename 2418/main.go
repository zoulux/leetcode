package main

import "sort"

func main() {

}

func sortPeople(names []string, heights []int) []string {
	ln := len(names)
	type person struct {
		name   string
		height int
	}

	list := make([]*person, 0, ln)
	for i := 0; i < ln; i++ {
		list = append(list, &person{
			name:   names[i],
			height: heights[i],
		})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].height > list[j].height
	})

	ans := make([]string, 0, ln)
	for _, p := range list {
		ans = append(ans, p.name)
	}
	return ans
}

func sortPeople2(names []string, heights []int) []string {
	ln := len(names)
	type person struct {
		name   string
		height int
	}

	list := make([]person, ln)
	for i := 0; i < ln; i++ {
		list[i] = person{
			name:   names[i],
			height: heights[i],
		}
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].height > list[j].height
	})

	ans := make([]string, 0, ln)

	for _, p := range list {
		ans = append(ans, p.name)
	}
	return ans
}
