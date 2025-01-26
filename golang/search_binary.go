package main

import "fmt"

func search_binary(a []int, target int) int {
    b := 0
    e := len(a) - 1

    for ; b <= e; {
        m := (b + e) / 2
        if a[m] == target {
            return m
        } else if a[m] < target {
            b = m + 1
        } else {
            e = m - 1
        }
    }

    return -1
}

func main() {
    a := []int{3, 4, 7, 5, 1, 6, 8, 2, 9, 0}
    target := 1
    fmt.Println("Array Elements:", a)
    res := search_binary(a, target)

    if res != -1 {
        fmt.Println("Element found at index:", res)
    } else {
        fmt.Println("Element not found")
    }
}
