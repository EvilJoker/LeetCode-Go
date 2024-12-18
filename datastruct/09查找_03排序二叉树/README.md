# 定义

排序二叉树（Binary Search Tree，简称 BST）是一种二叉树数据结构，其中每个节点的值大于其左子树的所有节点值，小于其右子树的所有节点值。排序二叉树有助于高效的查找、插入、删除和遍历操作。

## 示例

```golang
package main

import "fmt"

// 二叉树节点结构
type Node struct {
    value int
    left  *Node
    right *Node
}

// 插入节点
// 递归，没有叶子节点，就构造并返回
func insert(root *Node, value int) *Node {
    if root == nil {
        return &Node{value: value}
    }
    if value < root.value {
        root.left = insert(root.left, value)
    } else {
        root.right = insert(root.right, value)
    }
    return root
}

// 查找节点
// 二分查找
func search(root *Node, value int) *Node {
    if root == nil || root.value == value {
        return root
    }
    if value < root.value {
        return search(root.left, value)
    }
    return search(root.right, value)
}

// 删除节点
func delete(root *Node, value int) *Node {
    if root == nil {
        return root
    }
    if value < root.value {
        root.left = delete(root.left, value)
    } else if value > root.value {
        root.right = delete(root.right, value)
    } else {
        //被删除的节点就是自己
        // 节点只有一个子节点或没有子节点,直接返回继承者
        if root.left == nil {
            return root.right
        } else if root.right == nil {
            return root.left
        }
        // 节点有两个子节点，找到右子树的最小节点
        // 性质，右子树最小节点比左子树大
        minNode := minValueNode(root.right)
        // 赋值
        root.value = minNode.value
        // 删除，最小节点在左下角，直接删除即可
        root.right = delete(root.right, minNode.value)
    }
    return root
}

// 获取最小节点
func minValueNode(root *Node) *Node {
    current := root
    for current != nil && current.left != nil {
        current = current.left
    }
    return current
}

// 中序遍历
func inorder(root *Node) {
    if root != nil {
        inorder(root.left)
        fmt.Println(root.value)
        inorder(root.right)
    }
}

// 前序遍历
func preorder(root *Node) {
    if root != nil {
        fmt.Println(root.value)
        preorder(root.left)
        preorder(root.right)
    }
}

// 后序遍历
func postorder(root *Node) {
    if root != nil {
        postorder(root.left)
        postorder(root.right)
        fmt.Println(root.value)
    }
}

func main() {
    // 创建根节点
    root := &Node{value: 10}

    // 插入节点
    insert(root, 5)
    insert(root, 15)
    insert(root, 3)
    insert(root, 7)

    // 查找节点
    node := search(root, 7)
    if node != nil {
        fmt.Printf("Found node with value: %d\n", node.value)
    } else {
        fmt.Println("Node not found")
    }

    // 打印排序二叉树的中序遍历结果
    fmt.Println("Inorder traversal:")
    inorder(root)

    // 打印排序二叉树的前序遍历结果
    fmt.Println("Preorder traversal:")
    preorder(root)

    // 打印排序二叉树的后序遍历结果
    fmt.Println("Postorder traversal:")
    postorder(root)

    // 删除节点
    root = delete(root, 5)
    fmt.Println("After deleting 5:")
    inorder(root)
}
```

## 特点

查找操作：排序二叉树能在 O(log n) 时间复杂度内进行查找（假设树是平衡的）。
插入和删除操作：插入和删除节点也在 O(log n) 时间内完成（对于平衡树）。
递归性质：树的操作（插入、查找、遍历）通常使用递归进行。
平衡性：如果树是平衡的，排序二叉树非常高效，但如果树变得不平衡（例如：按顺序插入数据），性能可能退化为 O(n)。
遍历方式：
中序遍历：按照节点值的升序遍历树。
前序遍历：先访问根节点，然后是左子树，再是右子树。
后序遍历：先访问左子树，再是右子树，最后访问根节点。
删除节点：删除节点时，如果节点有两个子树，则用右子树的最小节点替换当前节点值，然后删除该最小节点。