/**
 * 练习：等价二叉树
 * -------------
可以用多种不同的二叉树的叶子节点存储相同的数列值。例如，这里有两个二叉树保存了序列 1，1，2，3，5，8，13。


用于检查两个二叉树是否存储了相同的序列的函数在多数语言中都是相当复杂的。这里将使用 Go 的并发和 channel 来编写一个简单的解法。

这个例子使用了 tree 包，定义了类型：

	type Tree struct {
		Left  *Tree
		Value int
		Right *Tree
	}

1. 实现 Walk 函数。

2. 测试 Walk 函数。

函数 tree.New(k) 构造了一个随机结构的二叉树，保存了值 k，2k，3k，...，10k。 创建一个新的 channel ch 并且对其进行步进：

go Walk(tree.New(1), ch)
然后从 channel 中读取并且打印 10 个值。应当是值 1，2，3，...，10。

3. 用 Walk 实现 Same 函数来检测是否 t1 和 t2 存储了相同的值。

4. 测试 Same 函数。

Same(tree.New(1), tree.New(1)) 应当返回 true，而 Same(tree.New(1), tree.New(2)) 应当返回 false。
*/
package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

//  发送value，结束后关闭channel
func Walk(t *tree.Tree, ch chan int){
	sendValue(t,ch)
	close(ch)
}
//  递归向channel传值
func sendValue(t *tree.Tree, ch chan int){
	if t != nil {
		sendValue(t.Left, ch)
		ch <- t.Value
		sendValue(t.Right, ch)
	}
}

// 使用写好的Walk函数来确定两个tree对象  是否一样 原理还是判断value值
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1,ch1)
	go Walk(t2,ch2)
	for i := range ch1 {   // ch1 关闭后   for循环自动跳出
		if i != <- ch2 {
			return false
		}
	}
	return true
}

func main() {

	// 打印 tree.New(1)的值
	var ch = make(chan int)
	go Walk(tree.New(1),ch)
	for v := range ch {
		fmt.Println(v)
	}

	//  比较两个tree的value值是否相等
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}