# 2049. Count Nodes With the Highest Score
## 題目
There is a binary tree rooted at 0 consisting of n nodes. The nodes are labeled from 0 to n - 1. You are given a 0-indexed integer array parents representing the tree, where parents[i] is the parent of node i. Since node 0 is the root, parents[0] == -1.

Each node has a score. To find the score of a node, consider if the node and the edges connected to it were removed. The tree would become one or more non-empty subtrees. The size of a subtree is the number of the nodes in it. The score of the node is the product of the sizes of all those subtrees.

Return the number of nodes that have the highest score.

---

## 想法
#### 使用DFS把節點的左右節點數量存下來,在依序移除節點來計算每個節點的最大分數

首先要先確認每個節點連接的左右節點  
`children` 表示每個節點裡面左子節點和右子節點  
e.g. children[2] = [3,1] -> 節點 2 有左子節點 3 和右子節點 1

```go []
	children := make([][]int, n)
	for k := range children {
		children[k] = make([]int, 0)
	}
	for i := 1; i < n; i++ {
		children[parents[i]] = append(children[parents[i]], i)
	}
```

然後再根據children去遞迴出每個節點的左右節點數量 `counts`
```go
func count(node int, children [][]int, counts *[][2]int) {
	if len(children[node]) == 2 {
		count(children[node][0], children, counts)
		count(children[node][1], children, counts)
		l := (*counts)[children[node][0]]
		r := (*counts)[children[node][1]]
		left := l[0] + l[1] + 1
		right := r[0] + r[1] + 1
		(*counts)[node] = [2]int{left, right}
	} else if len(children[node]) == 1 {
		count(children[node][0], children, counts)
		l := (*counts)[children[node][0]]
		left := l[0] + l[1] + 1
		(*counts)[node] = [2]int{left}
	}
}
```

最後就可以算出每個節點最大的分數,並找到最大分數個數
```go
	var ans int
	var counts = make([][2]int, n)
	count(0, children, &counts)
	scores := make(map[int]int)
	for i := 0; i < n; i++ {
		l, r, remain := counts[i][0], counts[i][1], (n - 1 - counts[i][0] - counts[i][1])
		if l == 0 {
			l = 1
		}
		if r == 0 {
			r = 1
		}
		if remain == 0 {
			remain = 1
		}
		scores[l*r*remain]++
		ans = max(ans, l*r*remain)
	}
	return scores[ans]
```
[完整代碼](./solution.go)