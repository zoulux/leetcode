package main

import "fmt"

func main() {

	fmt.Println(calcEquation([][]string{
		{"a", "e"},
		{"b", "e"},
	},
		[]float64{4.0, 3.0},
		[][]string{
			{"a", "b"},
			{"e", "e"},
			{"x", "x"},
		}))

	//fmt.Println(calcEquation([][]string{
	//	{"a", "b"},
	//	{"b", "c"},
	//}, []float64{2.0, 3.0}, [][]string{
	//	{"a", "c"},
	//	{"b", "a"},
	//	{"a", "e"},
	//	{"a", "a"},
	//	{"x", "x"},
	//}))

	//fmt.Println(calcEquation([][]string{
	//	{"a", "b"},
	//	{"a", "d"},
	//	{"d", "c"},
	//},
	//	[]float64{1.0, 2.0, 3.0},
	//	[][]string{
	//		{"a", "c"},
	//		{"b", "d"},
	//		{"b", "a"},
	//		{"d", "c"},
	//	}))
}

func calcEquation2(equations [][]string, values []float64, queries [][]string) []float64 {

	graph := make(map[string]map[string]float64)
	for i, v := range equations {

		if graph[v[0]] == nil {
			graph[v[0]] = make(map[string]float64)
		}

		if graph[v[1]] == nil {
			graph[v[1]] = make(map[string]float64)
		}

		graph[v[0]][v[1]] = values[i]
		graph[v[1]][v[0]] = 1 / values[i]
	}

	var dfs func(from, to string, visit map[string]bool) float64
	dfs = func(from, to string, visit map[string]bool) float64 {
		if graph[from] == nil || graph[to] == nil {
			return -1
		} else if from == to {
			// 相等返回 1
			return 1
		} else if v, ok := graph[from][to]; ok {
			return v
		} else {
			for k, v := range graph[from] {
				if visit[k] {
					continue
				}

				visit[k] = true
				if t := dfs(k, to, visit); t != -1 {
					return t * v
				}
				visit[k] = false

			}
			return -1
		}
	}

	var ans []float64

	for _, q := range queries {
		ans = append(ans, dfs(q[0], q[1], make(map[string]bool)))
	}
	return ans
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	type item struct {
		token string
		v     float64
	}
	tokens := make(map[string]struct{})
	for _, v := range equations {
		for _, vv := range v {
			tokens[vv] = struct{}{}
		}
	}
	parent := make(map[string]*item, len(tokens))
	// x 的父节点是 b
	// x / y   = v

	for k := range tokens {
		parent[k] = &item{token: k, v: 1}
	}

	var find func(t string) *item
	find = func(t string) *item {
		if parent[t].token != t {
			tv := find(parent[t].token)
			parent[t] = &item{
				v:     tv.v * parent[t].v,
				token: tv.token,
			}
		}

		return parent[t]
	}

	union := func(a, b string, v float64) {
		ait := find(a)
		bit := find(b)

		parent[bit.token] = &item{
			token: ait.token,
			v:     bit.v * v * ait.v,
		}
	}

	for idx, v := range equations {
		union(v[0], v[1], values[idx])
	}

	var ans []float64
	for _, q := range queries {
		if _, ok := tokens[q[0]]; !ok {
			ans = append(ans, -1)
		} else if _, ok := tokens[q[1]]; !ok {
			ans = append(ans, -1)
		} else {
			v0, v1 := find(q[0]), find(q[1])
			if v0.token == v1.token {
				ans = append(ans, v1.v/v0.v)
			} else {
				ans = append(ans, -1)
			}
		}
	}
	return ans
}
