/**
 * for 是 Go 的 “while”
 * -------------
如果省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环。
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	var i int
	for {
		i++
		fmt.Println(i)
		time.Sleep(time.Second/10)
	}
}
