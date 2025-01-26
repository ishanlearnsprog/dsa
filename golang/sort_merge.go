package main

import "fmt"

func merge(a []int, b int, m int, e int) {
    l := e-b+1
    temp := make([]int, l)
    i , j, end1, end2, k := b, m+1, m, e, 0
   
    for ; i <= end1 && j <= end2; k++ {
        if a[i] < a[j] {
            temp[k] = a[i]
            i++
        } else {
            temp[k] = a[j]
            j++
        }
    }

    for ; i <= end1; i++ {
        temp[k] = a[i]
        k++
    }

    for ; j <= end2; j++ {
        temp[k] = a[j]
        k++
    }

    for k = 0; k < l; k++ {
        a[k+b] = temp[k]
    }
}

func sort_merge(a []int, b int, e int) {
    if (b < e) {
        m := (b+e)/2
        sort_merge(a, b, m)
        sort_merge(a, m+1, e)
        merge(a, b, m, e)
    }
}

func main() {
    a := []int{3, 4, 7, 5, 1, 6, 8, 2, 9, 0}
    fmt.Println("Before Sorting:", a)
    sort_merge(a, 0, len(a)-1)
    fmt.Println("After Sorting:", a)
}
