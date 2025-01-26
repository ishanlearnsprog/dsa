package main

import "fmt"

func heapify(a []int, n int, i int) {
    largest := i
    left := 2 * i + 1
    right := 2 * i + 2

    if left < n && a[left] > a[largest] {
        largest = left
    }

    if right < n && a[right] > a[largest] {
        largest = right
    }

    if largest != i {
        a[i], a[largest] = a[largest], a[i]
        heapify(a, n, largest)
    }
}

func sort_heap(a []int) {
    l := len(a)

    for i := l / 2 - 1; i >= 0; i-- {
        heapify(a, l, i)
    }

    for i := l - 1; i > 0; i-- {
        a[0], a[i] = a[i], a[0]
        heapify(a, i, 0)
    }
}

func main() {
    a := []int{3, 4, 7, 5, 1, 6, 8, 2, 9, 0}
    fmt.Println("Before Sorting:", a)
    sort_heap(a)
    fmt.Println("After Sorting:", a)
}
