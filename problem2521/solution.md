# 2453. Destroy Sequential Targets

You are given a 0-indexed array `nums` consisting of positive integers, representing targets on a number line. You are also given an integer `space`.

You have a machine which can destroy targets. Seeding the machine with some `nums[i]` allows it to destroy all targets with values that can be represented as `nums[i] + c * space`, where `c` is any non-negative integer. You want to destroy the maximum number of targets in `nums`.

Return the minimum value of `nums[i]` you can seed the machine with to destroy the maximum number of targets.

---

## Explain

該題目就是求所有 `nums` 的元素, 裡面會有幾個質因數  
本題的想法比較直覺, 直接對每個 `num` 做質因數分解, 把質因數存進 `map` 中  
最後返回有幾個質因數

## Code 

[Code](./solution.go)
