import sys
import heapq


def prim_mst(weight_matrix, start_vertex):
    num_vertices = len(weight_matrix)
    mst = [[] for _ in range(num_vertices)]
    visited = [False] * num_vertices
    min_heap = [(0, start_vertex, -1)]
    
    while min_heap:
        weight, current_vertex, parent_vertex = heapq.heappop(min_heap)
        
        if visited[current_vertex]:
            continue
        visited[current_vertex] = True
        
        if parent_vertex != -1:
            mst[parent_vertex].append((current_vertex, weight))
            mst[current_vertex].append((parent_vertex, weight))
        
        for neighbor in range(num_vertices):
            if not visited[neighbor] and weight_matrix[current_vertex][neighbor] != -1:
                heapq.heappush(min_heap, (weight_matrix[current_vertex][neighbor], neighbor, current_vertex))
    
    return mst


def dfs_traversal(mst, start_vertex):
    visited = [False] * len(mst)
    path = []
    
    def dfs(vertex):
        visited[vertex] = True
        path.append(vertex)
        for neighbor, _ in mst[vertex]:
            if not visited[neighbor]:
                dfs(neighbor)
    
    dfs(start_vertex)
    return path


def calculate_path_weight(path, weight_matrix):
    total_weight = 0
    for i in range(len(path) - 1):
        total_weight += weight_matrix[path[i]][path[i+1]]
    return total_weight


def tsp_2_approximation(start_vertex, weight_matrix):
    mst = prim_mst(weight_matrix, start_vertex)
    path = dfs_traversal(mst, start_vertex)
    path.append(start_vertex)  
    total_weight = calculate_path_weight(path, weight_matrix)
    return total_weight, path


def main():
    input_data = sys.stdin.read().split()
    start_vertex = int(input_data[0])
    n = int((len(input_data) - 1) ** 0.5)
    weight_matrix = []
    index = 1
    for i in range(n):
        row = []
        for j in range(n):
            row.append(float(input_data[index]))
            index += 1
        weight_matrix.append(row)
    
    total_weight, path = tsp_2_approximation(start_vertex, weight_matrix)
    print(f"{total_weight:.2f}")
    print(" ".join(map(str, path)))

    
if __name__ == "__main__":
    main()