/* 传递地址到函数 */

package main

import "fmt"

func main() {
	var a int = 100
	var b int = 200

	swap(&a, &b)

	fmt.Printf("after swap: %d\n", a)

	fmt.Printf("after swap: %d\n", b)

}

func swap(x *int, y *int) {
	var temp int = *x
	*x = *y
	*y = temp
}
