# 贪心之单元最短路径

## 题目介绍

单元最短路径问题（Unit Shortest Path Problem）是一个常见的图算法问题，目的是寻找从起点到终点的最短路径。与经典的最短路径问题（如 Dijkstra）类似，单元最短路径问题要求图中的每条边的权重为 1，或者所有边的权重是相等的。此时，我们可以采用广度优先搜索（BFS）来解决问题，因为 BFS 是一种基于层次遍历的搜索算法，非常适合处理边权相同的情况。

**输入**：一个无权图，其中包含若干个节点和边。

**输出**：从起点到终点的最短路径，或者表示无法到达的情况。

## 解题思路

- 由于图的边权相同（单位边权），我们可以使用广度优先搜索（BFS）来解决最短路径问题。
- 通过 BFS 遍历图的所有节点，并逐层扩展节点，直到找到终点或遍历完所有可能的路径。
- 使用队列（queue）来实现 BFS，记录每个节点的访问状态及其到起点的距离。

### 步骤

1. 使用队列保存待遍历的节点，初始化队列包含起点。
2. 对于队列中的每个节点，检查其相邻节点是否已访问过，如果未访问则将其加入队列，并记录其距离。
3. 重复上述步骤，直到遍历到终点或者队列为空。
4. 如果最终找到了终点，则返回其到起点的距离，否则返回-1表示无法到达。

## 示例

给定如下图，图中每个节点之间的边权均为1：
0 -- 1 -- 2 | | 3 -- 4

起点为 0，终点为 4，最短路径为 0 -> 1 -> 4，最短距离为 2。

## Golang 示例

```golang
package main

import (
    "fmt"
    "container/list"
)

// 定义图结构
type Graph struct {
    vertices int
    edges    map[int][]int
}

// 创建图
func NewGraph(vertices int) *Graph {
    return &Graph{
        vertices: vertices,
        edges:    make(map[int][]int),
    }
}

// 添加边
func (g *Graph) AddEdge(u, v int) {
    g.edges[u] = append(g.edges[u], v)
    g.edges[v] = append(g.edges[v], u) // 因为是无向图
}

// 使用广度优先搜索（BFS）解决单元最短路径问题
func (g *Graph) BFS(start, end int) int {
    // 初始化队列
    queue := list.New()
    queue.PushBack(start)

    // 初始化访问数组
    dist := make([]int, g.vertices)
    for i := range dist {
        dist[i] = -1 // -1 表示未访问
    }
    dist[start] = 0 // 起点的距离为 0

    // 广度优先搜索
    for queue.Len() > 0 {
        node := queue.Remove(queue.Front()).(int)

        // 如果到达终点
        if node == end {
            return dist[node]
        }

        // 遍历相邻节点
        for _, neighbor := range g.edges[node] {
            if dist[neighbor] == -1 { // 如果该节点未访问
                dist[neighbor] = dist[node] + 1
                queue.PushBack(neighbor)
            }
        }
    }

    // 如果无法到达终点，返回 -1
    return -1
}

func main() {
    // 创建图
    g := NewGraph(5)
    g.AddEdge(0, 1)
    g.AddEdge(1, 2)
    g.AddEdge(1, 4)
    g.AddEdge(0, 3)
    g.AddEdge(3, 4)

    // 寻找最短路径
    start, end := 0, 4
    shortestPath := g.BFS(start, end)

    if shortestPath != -1 {
        fmt.Printf("从节点 %d 到节点 %d 的最短路径长度为: %d\n", start, end, shortestPath)
    } else {
        fmt.Println("无法到达目标节点")
    }
}
```

## 关键注释说明

NewGraph：创建一个新的图，初始化图的节点数和边的列表。
AddEdge：在图中添加一条无向边。
BFS：广度优先搜索算法，用于计算从起点到终点的最短路径。
dist 数组用于记录每个节点的距离。
如果节点未访问过，则将其加入队列，并更新其距离。
当遍历到终点时，返回其距离；如果遍历完所有节点仍未找到终点，返回 -1。
特点
时间复杂度：O(V + E)，其中 V 是图中的节点数，E 是图中的边数。因为每个节点和边最多被访问一次。
空间复杂度：O(V)，用于存储节点的访问状态和距离信息。
适用场景：适用于所有无权图或单位边权图的最短路径问题。
