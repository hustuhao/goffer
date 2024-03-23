package v1

import "math"

const DjInf = math.MaxInt32

// Dijkstra 算法，返回距离和路径
// 参数 graph[i][j] 表示从点 i 到点 j 的距离，为0表示边不存在。这里假设不同点的边权重是正数，不会是0。
// 返回起点到达每一个点的最近距离，以及到达该点最优路径的前一个点
func Dijkstra(graph [][]int, start int) ([]int, []int) {
	n := len(graph)          // 图中顶点的数量。
	dist := make([]int64, n) // dist 数组用于存储从起始点到每个顶点的最短距离。初始时，所有距离都设置为无穷大（INF），表示目前尚未找到路径。
	prev := make([]int, n)   // prev 数组用于跟踪到达每个顶点的最短路径上的前一个顶点。

	visited := make([]bool, n) // visited 数组用于标记一个顶点是否已被处理。初始时，所有顶点都未被访问。被访问即为最优路径上的点。

	for i := range dist {
		dist[i] = DjInf
		prev[i] = -1 // 表示未找到路径通向i
	}
	dist[start] = 0 // 将起始点到自身的距离设为 0。

	for i := 0; i < n; i++ { // 遍历每个顶点
		u := -1
		for v := 0; v < n; v++ { // 找到未被访问的点中，距离起始点最近的点 u。
			if !visited[v] && (u == -1 || dist[v] < dist[u]) {
				u = v
			}
		}

		visited[u] = true // 标记被访问过的点，
		for v := 0; v < n; v++ {
			if graph[u][v] != 0 && dist[u]+int64(graph[u][v]) < dist[v] { // 更新从 u 到其他所有顶点 v 的最短路径
				dist[v] = dist[u] + int64(graph[u][v])
				prev[v] = u // 记录v的前一个点为u
			}
		}
	}

	retDis := make([]int, len(dist))
	for i, d := range dist {
		retDis[i] = int(d)
	}
	return retDis, prev
}
