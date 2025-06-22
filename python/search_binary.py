a = [1, 2, 3, 4, 5, 6, 7, 8, 9, 0]
key = 4
res = -1

low = 0
high = len(a) - 1

while low <= high:
    mid = (low + high) // 2
    if a[mid] < key:
        low = mid + 1
    elif a[mid] > key:
        high = mid - 1
    else:
        res = mid
        break

if res == -1:
    print("Element not found")
else:
    print("Element found at idx:", res)
