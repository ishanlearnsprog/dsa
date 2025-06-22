class search_linear {
    public static void main(String args[]) {
        int a[] = {2, 0, 4, 5, 1, 3, 6, 7, 8, 9};
        int key = 8;
        int res = -1;

        for (int i = 0; i < a.length; i++) {
            if (a[i] == key) {
                res = i;
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
