package main


import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    pre := 0.0
    ret := 1.5
    count := 0
    for math.Abs(ret - pre) > 0.1 {
        pre = ret
        ret = math.Abs(ret - (ret*ret - x)/2*ret)
        count++
    }
    fmt.Println(" - Sqrt execution count:", count)
    return math.Abs(ret)
}

func main() {
    x := 1.0
    for x <= 3 {
        fmt.Println("       Target value: ", x)
        fmt.Println("Created Sqrt result: ", Sqrt(x))
        fmt.Println("   math.Sqrt result: ", math.Sqrt(x))
        fmt.Println("")
        x = x + 1
    }
}
