class search_binary {
    public static void main(String args[]) {
        int a[] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9};
        int key = 8;
        int res = -1;

        int low = 0;
        int high = a.length - 1;

        while (low <= high) {
            int mid = (low + high) / 2;
            if (a[mid] < key) {
                low = mid + 1;
            } else if (a[mid] > key) {
                high = mid - 1;
            } else {
                res = mid;
                break;
            }
        }

        if (res == -1) {
            System.out.println("Element not found");
        } else {
            System.out.println("Element found at idx: " + res);
        }
    }
}
