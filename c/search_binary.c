#include <stdio.h>

int search_binary(int *a, int length, int key) {
    int b = 0;
    int e = length - 1;

    while (b <= e) {
        int m = (b + e) / 2;

        if (a[m] < key) {
            b = m + 1;
        } else if (a[m] > key) {
            e = m - 1;
        } else {
            return m;
        }
    }

    return -1;
}

int main() {
    int a [] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 0};
    int length = sizeof(a)/sizeof(a[0]);
    int key = 8;
    int res = search_binary(a, length, key);

    if (res == -1) {
        printf("Element not found\n");
    } else {
        printf("Element found at idx: %d\n", res);
    }

    return 0;
}
