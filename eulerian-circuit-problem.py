from typing import List

def eulerian_circuit(graph: List[List[int]]) -> List[int]:
    # Check if the graph has an Eulerian circuit
    if not has_eulerian_circuit(graph):
        return []

    # Initialize the circuit with the starting vertex
    circuit = [0]

    # Keep track of the visited edges
    visited_edges = set()

    # Keep track of the current vertex
    current_vertex = 0

    while len(visited_edges) < len(graph) * len(graph):
        # Find an edge that hasn't been visited yet
        next_vertex = -1
        for v in range(len(graph)):
            if graph[current_vertex][v] > 0 and (current_vertex, v) not in visited_edges:
                next_vertex = v
                break

        # If there are no unvisited edges, backtrack
        if next_vertex == -1:
            for v in range(len(graph)):
                if graph[current_vertex][v] > 0 and degree(graph, v) > 1 and is_bridge(graph, visited_edges, current_vertex, v):
                    next_vertex = v
                    break
            if next_vertex == -1:
                return []

        # Add the next vertex to the circuit and mark the edge as visited
        circuit.append(next_vertex)
        visited_edges.add((current_vertex, next_vertex))

        # Update the current vertex
        current_vertex = next_vertex

    return circuit

def has_eulerian_circuit(graph: List[List[int]]) -> bool:
    # A graph has an Eulerian circuit if and only if every vertex has even degree
    for v in range(len(graph)):
        if degree(graph, v) % 2 != 0:
            return False
    return True

def degree(graph: List[List[int]], v: int) -> int:
    # Calculate the degree of vertex v in the graph
    return sum(graph[v])

def is_bridge(graph: List[List[int]], visited_edges: set, u: int, v: int) -> bool:
    # Check if the edge (u, v) is a bridge in the graph
    visited = [False] * len(graph)
    visited[u] = True
    stack = [u]
    while stack:
        x = stack.pop()
        for i in range(len(graph)):
            if graph[x][i] > 0 and (x, i) not in visited_edges and not visited[i]:
                visited[i] = True
                stack.append(i)
    return not visited[v]

# Test the program with a sample graph
graph = [[0, 1, 0, 1, 0],
         [1, 0, 1, 0, 1],
         [0, 1, 0, 1, 0],
         [1, 0, 1, 0, 1],
         [0, 1, 0, 1, 0]]
circuit = eulerian_circuit(graph)
if len(circuit) > 0:
    print("Eulerian circuit found:", circuit)
else:
    print("No Eulerian circuit found.")
