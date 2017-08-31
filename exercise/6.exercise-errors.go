/**
 * 练习：错误
 * -------------
从先前的练习中复制 Sqrt 函数，并修改使其返回 error 值。

由于不支持负数，当 Sqrt 接收到一个负数时，应当返回一个非 nil 的错误值。

创建一个新类型

	type ErrNegativeSqrt float64
为其实现

	func (e ErrNegativeSqrt) Error() string
使其成为一个 error， 该方法就可以让 ErrNegativeSqrt(-2).Error() 返回 `"cannot Sqrt negative number: -2"`。

*注意：* 在 Error 方法内调用 fmt.Sprint(e) 将会让程序陷入死循环。可以通过先转换 e 来避免这个问题：fmt.Sprint(float64(e))。请思考这是为什么呢？

修改 Sqrt 函数，使其接受一个负数时，返回 ErrNegativeSqrt 值。
*/
package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	// return fmt.Sprint(e) // 死循环问题
	// 死循环原因：fmt.Sprint 会通过接口查询知道 e 是一个ErrNegativeSqrt接口类型，所以就会调用 e 的Error函数，
	// 但这个 fmt.Sprint 本身就是在 Error 函数里调用的，所以就构成循环调用了。
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := float64(1)
	precision := float64(100000000)

	for {
		old := z
		z = z - (z*z-x)/(2*z)
		if uint(z*precision) == uint(old*precision) {
			return z, nil
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}