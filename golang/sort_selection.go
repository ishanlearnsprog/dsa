package main

import "fmt"

func sort_selection(a []int) {
    l := len(a)

    for i := 0; i < l-1; i++ {
        mini := a[i]
        idx := i
        for j := i+1; j < l; j++ {
            if a[j] < mini {
                mini = a[j]
                idx = j
            }
        }
        if i != idx {
            a[idx], a[i] = a[i], a[idx]
        }
    }
}

func main() {
    a := []int{3, 4, 7, 5, 1, 6, 8, 2, 9, 0}
    fmt.Println("Before Sorting:", a)
    sort_selection(a)
    fmt.Println("After Sorting:", a)
}
