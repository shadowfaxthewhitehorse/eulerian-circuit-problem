package main

import (
    "fmt"
)

func eulerianCircuit(graph [][]int) []int {
    // Check if the graph has an Eulerian circuit
    if !hasEulerianCircuit(graph) {
        return []int{}
    }

    // Initialize the circuit with the starting vertex
    circuit := []int{0}

    // Keep track of the visited edges
    visitedEdges := make(map[[2]int]bool)

    // Keep track of the current vertex
    currentVertex := 0

    for len(visitedEdges) < len(graph)*len(graph) {
        // Find an edge that hasn't been visited yet
        nextVertex := -1
        for v := range graph[currentVertex] {
            if graph[currentVertex][v] > 0 && !visitedEdges[[2]int{currentVertex, v}] {
                nextVertex = v
                break
            }
        }

        // If there are no unvisited edges, backtrack
        if nextVertex == -1 {
            for v := range graph[currentVertex] {
                if graph[currentVertex][v] > 0 && degree(graph, v) > 1 && isBridge(graph, visitedEdges, currentVertex, v) {
                    nextVertex = v
                    break
                }
            }
            if nextVertex == -1 {
                return []int{}
            }
        }

        // Add the next vertex to the circuit and mark the edge as visited
        circuit = append(circuit, nextVertex)
        visitedEdges[[2]int{currentVertex, nextVertex}] = true

        // Update the current vertex
        currentVertex = nextVertex
    }

    return circuit
}

func hasEulerianCircuit(graph [][]int) bool {
    // A graph has an Eulerian circuit if and only if every vertex has even degree
    for v := range graph {
        if degree(graph, v)%2 != 0 {
            return false
        }
    }
    return true
}

func degree(graph [][]int, v int) int {
    // Calculate the degree of vertex v in the graph
    d := 0
    for i := range graph {
        d += graph[v][i]
    }
    return d
}

func isBridge(graph [][]int, visitedEdges map[[2]int]bool, u int, v int) bool {
    // Check if the edge (u, v) is a bridge in the graph
    visited := make([]bool, len(graph))
    visited[u] = true
    stack := []int{u}
    for len(stack) > 0 {
        x := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        for i := range graph[x] {
            if graph[x][i] > 0 && !visitedEdges[[2]int{x, i}] && !visited[i] {
                visited[i] = true
                stack = append(stack, i)
            }
        }
    }
    return !visited[v]
}

//func main() {
    // Test the program with a sample graph
//    graph := [][]int{{0, 1, 0, 1, 0},
//                    {1, 0, 1, 0, 1},
//                     {0, 1, 0, 1, 0},
//                     {1, 0, 1, 0, 1},
//                    {0, 1, 0, 1, 0}}
//    circuit := eulerianCircuit(graph)
