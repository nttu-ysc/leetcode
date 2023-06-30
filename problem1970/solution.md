# 1970. Last Day Where You Can Still Cross
## 題目
There is a 1-based binary matrix where 0 represents land and 1 represents water. You are given integers row and col representing the number of rows and columns in the matrix, respectively.

Initially on day 0, the entire matrix is land. However, each day a new cell becomes flooded with water. You are given a 1-based 2D array cells, where cells[i] = [ $r_i$, $c_i$ ] represents that on the $i^{th}$ day, the cell on the $r_{i^{th}}$ row and $c_{i^{th}}$ column (1-based coordinates) will be covered with water (i.e., changed to 1).

You want to find the last day that it is possible to walk from the top to the bottom by only walking on land cells. You can start from any cell in the top row and end at any cell in the bottom row. You can only travel in the four cardinal directions (left, right, up, and down).

Return the last day where it is possible to walk from the top to the bottom by only walking on land cells.

---
## 想法
使用`BFS`來判斷是否達到最後一row,如果每次都從第一天開始算,當天數很大時會發生TLE,所以這邊在使用`binary search`來減少運算時間

比較需要注意的是每次`bfs`後,走訪過後的紀錄不重置(共用同一個visited array)  
因為目標是走到最後一row就完成,不必每條路線都完整畫出,如果紀錄發現已走過代表這條路線是無法成功的路線所以直接return

[完整代碼](./solution_bfs.go)
