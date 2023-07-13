# 207. Course Schedule

There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

- For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.

Return `true` if you can finish all courses. Otherwise, return `false`.

## Explain

本題思路先找出課程要求數為 `0` 的丟到 `queue` 先開始
```go
indegree := make([]int, numCourses)
for _, pre := range prerequisites {
    indegree[pre[0]]++
}
var queue []int
for course, degree := range indegree {
    if degree == 0 {
        queue = append(queue, course)
    }
}
```

然後從 `queue` 拿出來的每個 `course` 去找 `prerequisites` 裡面有要求當堂 `course` 把 `indegree - 1 `.  
如果 
- `indegree < 0` 不可能出現
- `indegree = 0` 代表課程已經可修了, 加到 `queue`
- `indegree > 0` 不做事

```go
cnt := len(queue)

for len(queue) != 0 {
    course := queue[0]
    queue = queue[1:]

    for _, pre := range prerequisites {
        if course == pre[1] {
            indegree[pre[0]]--

            if indegree[pre[0]] < 0 {
                return false
            }

            if indegree[pre[0]] == 0 {
                cnt++
                queue = append(queue, pre[0])
            }
        }
    }
}
```

## Code

[Solution](./main.go)
