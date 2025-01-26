package main

import "fmt"

func search_linear(a []int, t int) int {
    for i, e := range a {
        if e == t {
            return i
        }
    }
    return -1
}

func main() {
    a := []int{3, 4, 7, 5, 1, 6, 8, 2, 9, 0}
    target := 11
    fmt.Println("Array Elements:", a)
    res := search_linear(a, target)

    if res != -1 {
        fmt.Println("Element found at index:", res)
    } else {
        fmt.Println("Element not found")
    }
}
