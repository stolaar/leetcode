package main

import (
	"fmt"
	"math"
)

func bfs(graph map[int][]int, src int, dest int) int {
	q := []int{src}
	visited := map[int]bool{}
	distances := make(map[int]int, len(graph))

	for i := 0; i < len(graph); i++ {
		distances[i] = math.MaxInt64
	}

	distances[src] = 0

	for len(q) > 0 {
		n := q[0]
		if len(q) > 1 {
			q = q[1:]
		} else {
			q = []int{}
		}

		for _, neighbour := range graph[n] {
			if distances[neighbour] == math.MaxInt64 {
				distances[neighbour] = distances[n] + 1
			}

			if distances[neighbour] > distances[n]+1 {
				distances[neighbour] = distances[n] + 1
			}

			if _, ok := visited[neighbour]; ok {
				continue
			}
			visited[neighbour] = true

			q = append(q, neighbour)
		}
	}

	return distances[dest]
}

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	graph := make(map[int][]int, n)

	for i := 0; i < n; i++ {
		if i == n-1 {
			graph[i] = []int{}
			break
		}
		graph[i] = []int{i + 1}
	}
	result := make([]int, len(queries))

	for idx, query := range queries {
		graph[query[0]] = append(graph[query[0]], query[1])
		result[idx] = bfs(graph, 0, n-1)
	}

	return result
}

func main() {
	fmt.Println(shortestDistanceAfterQueries(5, [][]int{{2, 4}, {0, 2}, {0, 4}}))
}
