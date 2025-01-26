package main

import "fmt"

func partition(a []int, b int, e int) int {
    p, i, j := a[b], b+1, e

    for ; i < j; {
        
        for ; a[i] < p; i++ {}

        for ; a[j] > p; j-- {}

        if i < j {
            a[i], a[j] = a[j], a[i]
        }
    }

    a[j], a[b] = a[b], a[j]
    return j
}

func sort_quick(a []int, b int, e int) {
    if (b < e) {
        pivot := partition(a, b, e)
        sort_quick(a, b, pivot-1)
        sort_quick(a, pivot + 1, e)
    }
}

func main() {
    a := []int{3, 4, 7, 5, 1, 6, 8, 2, 9, 0}
    fmt.Println("Before Sorting:", a)
    sort_quick(a, 0, len(a)-1)
    fmt.Println("After Sorting:", a)
}
