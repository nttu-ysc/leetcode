# 1601. Maximum Number of Achievable Transfer Requests
## 題目
We have `n` buildings numbered from `0` to `n - 1`. Each building has a number of employees. It's transfer season, and some employees want to change the building they reside in.

You are given an array `requests` where requests[i] = [$from_i$, $to_i$] represents an employee's request to transfer from building $from_i$ to building $to_i$.

All buildings are full, so a list of requests is achievable only if for each building, the net change in employee transfers is zero. This means the number of employees leaving is equal to the number of employees moving in. For example if `n = 3` and two employees are leaving building `0`, one is leaving building `1`, and one is leaving building `2`, there should be two employees moving to building `0`, one employee moving to building `1`, and one employee moving to building `2`.

Return the *maximum number of achievable* requests.

---
## 想法

使用遞迴來查找所有的可能性  
假設請求 `requests` 有 $k$ 個, 每個請求有 $2$ 種可能 (達到請求或不達到)  
所以會有 $2^k$ 種狀況, 我們可以把所有的狀況枚舉出來  
最後再判斷是否達到人數淨變化為 0

## 完整代碼
[Code](./solution.go)