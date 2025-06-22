a = [2, 0, 4, 5, 1, 3, 6, 7, 8, 9]
key = 8
res = -1

for i in range(0, len(a)):
    if a[i] == key:
        res = i
        break

if res == -1:
    print("Element not found")
else:
    print("Element found at idx:", res)
