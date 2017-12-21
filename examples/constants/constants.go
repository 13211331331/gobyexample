// Go supports _constants_ of character, string, boolean,
// and numeric values.

package main

import "fmt"
import "math"

// `const` declares a constant value.
const s string = "constant"

func main() {
    fmt.Println(s)

    // A `const` statement can appear anywhere a `var`
    // statement can.
    const n = 500000000

    // Constant expressions perform arithmetic with
    // arbitrary precision.
    const d = 3e20 / n
    fmt.Println(d)

    // A numeric constant has no type until it's given
    // one, such as by an explicit cast.
    fmt.Println(int64(d))

    // A number can be given a type by using it in a
    // context that requires one, such as a variable
    // assignment or function call. For example, here
    // `math.Sin` expects a `float64`.
    fmt.Println(math.Sin(n))






    const ( // iota被重设为0
        c0 = iota // c0 == 0
        c1 = iota // c1 == 1
        c2 = iota // c2 == 2
    )
    fmt.Println(c0)
    fmt.Println(c1)
    fmt.Println(c2)

    const x = iota // x == 0 (因为iota又被重设为0了)
    const y = iota // y == 0 (同上)
    fmt.Println(x)
    fmt.Println(y)


    const ( // iota被重设为0
        D0 = iota // c0 == 0
        D1 // c1 == 1
        D2 // c2 == 2
    )
    fmt.Println(D0)
    fmt.Println(D1)
    fmt.Println(D2)

    const (
        a = 1 <<iota // a == 1 (iota在每个const开头被重设为0)
        b // b == 2
        c // c == 4
    )
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

    const (
        Sunday = iota
        Monday
        Tuesday
        Wednesday
        Thursday
        Friday
        Saturday
        numberOfDays // 这个常量没有导出
    )
    fmt.Println(numberOfDays)
    /**
    同Go语言的其他符号（symbol）一样，以大写字母开头的常量在包外可见。
    以上例子中numberOfDays为包内私有，其他符号则可被其他包访问。
    **/

    str := "Hello,世界"
    m := len(str)
    for i := 0; i < m; i++ {
        ch := str[i] // 依据下标取字符串中的字符，类型为byte
        fmt.Println(i, ch)

    }

}
