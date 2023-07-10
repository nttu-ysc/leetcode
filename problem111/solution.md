# 111. Minimum Depth of Binary Tree

Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

---

## Explain

本題直接使用 `bfs` 解就能直接算出來了, 依照每層去遞回, 直到該節點的左右子節點都為 `null` 變返回 `depth`

## Code

[Code](./solution.go)