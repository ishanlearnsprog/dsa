package main

import "fmt"

func sort_insertion(a []int) {
    l := len(a)

    for i := 1; i < l; i++ {
        mini := a[i]
        j := i
        for ; j-1 >= 0 && a[j-1] > mini; j-- {
            a[j] = a[j-1]
        }
        a[j] = mini
    }
}

func main() {
    a := []int{3, 4, 7, 5, 1, 6, 8, 2, 9, 0}
    fmt.Println("Before Sorting:", a)
    sort_insertion(a)
    fmt.Println("After Sorting:", a)
}
