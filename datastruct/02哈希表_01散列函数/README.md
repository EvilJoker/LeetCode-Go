# 定义

哈希表: 是一种通过散列函数将数据的键映射到数组下标来实现快速数据查找的数据结构。

散列函数: 是将键映射为固定范围整数（数组索引）的函数，用于高效定位存储位置。

## 示例

```golang
package main

import (
    "fmt"
    "hash/fnv"
)

// 定义一个简单的哈希函数
func hash(key string, size int) int {
    h := fnv.New32a()
    h.Write([]byte(key))
    return int(h.Sum32()) % size
}

func main() {
    // 哈希表大小
    size := 10

    // 示例键值
    keys := []string{"apple", "banana", "cherry"}

    // 打印每个键的哈希值
    for _, key := range keys {
        index := hash(key, size)
        fmt.Printf("Key: %s, Hash Index: %d\n", key, index)
    }
}
```

## 特点

+ 快速查找: 时间复杂度接近 O(1)，适合快速检索。
+ 散列函数设计影响性能: 需均匀分布避免冲突。
+ 冲突处理: 常见方法包括链地址法和开放地址法。
+ 哈希表扩容: 当负载因子过高时，需要动态扩容以维持性能。
