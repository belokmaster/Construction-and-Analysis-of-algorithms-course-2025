from sys import stdin


def input_graph(input_str):
    graph = {}
    for i in range(len(input_str)):
        k, node, value = input_str[i].split()
        updated_ends = graph.get(k, {})
        updated_ends[node] = float(value)
        graph[k] = updated_ends
    return graph


def greedy_traversal(start, end, graph: dict, passed=""):
    if start == end:
        return start 
    if start not in graph.keys() or not graph[start] or start in passed:
        return None
    for k, _ in sorted(graph[start].items(), key=lambda pairs: pairs[1]):
        rslt = greedy_traversal(k, end, graph, passed=passed+start)
        if rslt != None:
            return start+rslt
    

if __name__ == "__main__":
    input_str = stdin.readlines()
    start, end = input_str[0].split()
    graph = input_graph(input_str[1:])
    print(greedy_traversal(start, end, graph))


import sys
from typing import TypeVar, Generic

T = TypeVar('T')


class WeightedDiGraph(Generic[T]):
    _adj: dict[T, dict[T, float]]
    _nodes: set[T]

    def __init__(self):
        self._adj = {}
        self._nodes = set()

    def __iter__(self):
        return iter(self._nodes)

    def __contains__(self, node: T):
        return node in self._nodes

    def __len__(self):
        return len(self._nodes)

    def __getitem__(self, node: T):
        return self.get_neighbors(node)

    def size(self):
        return sum(len(neighbors) for neighbors in self._adj.values())

    def add_edge(self, node1: T, node2: T, weight: float):
        if node1 not in self._adj:
            self._adj[node1] = {}
            self._nodes.add(node1)
        self._adj[node1][node2] = weight
        self._nodes.add(node2)

    def add_edges(self, edges: list[tuple[T, T, float]]):
        for edge in edges:
            self.add_edge(*edge)

    def get_neighbors(self, node: T):
        return self._adj.get(node, {})


def greedy_search(graph: WeightedDiGraph[T], start: T, end: T) -> list[T]:
    if start == end:
        return [start]
    path = []
    current_node = start
    visited = set(start)
    while current_node != end:
        neighbors = graph.get_neighbors(current_node)
        if not neighbors:
            current_node = path.pop()
            continue
        min_weight = float('inf')
        min_node = None
        for node, weight in neighbors.items():
            if node not in visited and weight < min_weight:
                min_weight = weight
                min_node = node
        if min_node is None:
            current_node = path.pop()
            continue

        path.append(current_node)
        visited.add(min_node)
        current_node = min_node

    if path:
        path.append(end)
    return path

def a_star_search(graph: WeightedDiGraph[T], start: T, end: T) -> list[T]:
    pass

def main():
    start, end = input().split()
    edge_list = []
    for line in sys.stdin:
        try:
            if len(line) <= 1:
                break
            node1, node2, weight = line.split()
            weight = float(weight)
            edge_list.append((node1, node2, weight))
        except EOFError:
            break

    g = WeightedDiGraph()
    g.add_edges(edge_list)
    # print(g._adj)
    # print(g._nodes)
    res = greedy_search(g, start, end)
    print(*res, sep='')


if __name__ == '__main__':
    main()