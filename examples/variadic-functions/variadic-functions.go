// [_Variadic functions_](http://en.wikipedia.org/wiki/Variadic_function)
// can be called with any number of trailing arguments.
// For example, `fmt.Println` is a common variadic
// function.

package main

import "fmt"

// Here's a function that will take an arbitrary number
// of `ints` as arguments.
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}


//用interface{}传递任意类型数据是Go语言的惯例用法。使用interface{}仍然是类型安
//全的，这和 C/C++ 不太一样。关于它的用法，可参阅3.5节的内容。代码清单2-4示范了如何分派
//传入interface{}类型的数据。
func MyPrintf(args ...interface{}) {
    for _, arg := range args {
        switch arg.(type) {
        case int:
            fmt.Println(arg, "is an int value.")
        case string:
            fmt.Println(arg, "is a string value.")
        case int64:
            fmt.Println(arg, "is an int64 value.")
        default:
            fmt.Println(arg, "is an unknown type.")
        }
    }
}

func main() {

    // Variadic functions can be called in the usual way
    // with individual arguments.
    sum(1, 2)
    sum(1, 2, 3)

    // If you already have multiple args in a slice,
    // apply them to a variadic function using
    // `func(slice...)` like this.
    nums := []int{1, 2, 3, 4}
    sum(nums...)


    var v1 int = 1
    var v2 int64 = 234
    var v3 string = "hello"
    var v4 float32 = 1.234
    MyPrintf(v1, v2, v3, v4)

}
