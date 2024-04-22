package main

import "fmt"

/*
 * @lc app=leetcode id=1971 lang=golang
 *
 * [1971] Find if Path Exists in Graph
 */
func addEdges(graph map[int][]int, u int, v int) map[int][]int {
	graph[u] = append(graph[u], v)
	graph[v] = append(graph[v], u)

	return graph
}

func search(graph map[int][]int, visited map[string]bool, source int, destination int) []bool {
	found := []bool{}

	for i := 0; i < len(graph[source]); i++ {
		if graph[source][i] == destination {
			found = append(found, true)
			break
		}
		if !visited[fmt.Sprint(graph[source][i])] {
			visited[fmt.Sprint(graph[source][i])] = true
			found = append(found, search(graph, visited, graph[source][i], destination)...)
		}
	}

	return found
}

func validPath(_ int, edges [][]int, source int, destination int) bool {
	if len(edges) <= 0 || source == destination {
		return true
	}
	graph := make(map[int][]int)
	visited := make(map[string]bool)

	for i := 0; i < len(edges); i++ {
		graph = addEdges(graph, edges[i][0], edges[i][1])
	}

	result := search(graph, visited, source, destination)
	response := false

	for _, found := range result {
		if found {
			response = true
			break
		}
	}

	return response
}

// @lc code=end

func main() {
	fmt.Println(validPath(3, [][]int{{0, 1}, {1, 2}, {2, 0}}, 0, 2))
	fmt.Println(validPath(3, [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}}, 5, 9))
	fmt.Println(validPath(6, [][]int{{0, 1}, {0, 2}, {3, 5}, {4, 3}}, 0, 5))

	fmt.Println(validPath(50, [][]int{{18, 46}, {8, 48}, {13, 30}, {28, 29}, {2, 16}, {7, 36}, {12, 19}, {31, 16}, {11, 46}, {6, 46}, {19, 27}, {4, 24}, {10, 37}, {14, 37}, {39, 31}, {10, 22}, {23, 2}, {47, 11}, {40, 7}, {21, 17}, {9, 3}, {34, 10}, {48, 1}, {21, 35}, {43, 48}, {27, 5}, {36, 11}, {43, 36}, {31, 48}, {25, 33}, {46, 19}, {31, 30}, {16, 45}, {30, 10}, {35, 47}, {35, 13}, {37, 48}, {49, 3}, {7, 26}, {2, 30}, {0, 27}, {25, 9}, {28, 27}, {39, 18}, {32, 6}, {14, 43}, {9, 27}, {27, 4}, {6, 0}, {21, 43}}, 48, 2))
}
