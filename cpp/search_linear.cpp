#include <iostream>

using namespace std;

int compare(int x, int y) {
    return x == y;
}

int search_linear(int *a, int length, int key) {
    int i;
    for (i = 0; i < length; i++) {
        int res = compare(a[i], key);
        if (res == 1) {
            return i;
        }
    }
    return -1;
}

int main() {
    int a [] = {2, 4, 3, 6, 9, 7, 8, 1, 5, 0};
    int length = sizeof(a)/sizeof(a[0]);
    int key = 1;
    int res = search_linear(a, length, key);
    if (res == -1) {
        cout << "Element not found" << "\n";
    } else {
        cout << "Element found at idx: " << res << "\n";
    }
    return 0;
}
