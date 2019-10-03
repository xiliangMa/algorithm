import java.util.Arrays;

/**
 * Create by $(xiliangMa) on 2019-10-03
 */
public class bubble_sort_test {

    /**
     * 冒泡排序 Bubble Sort
     */
    public static int[] Test_BSort(int[] arrs) {
        int l = arrs.length;
        for (int i = 0; i < l - 1; i++) {
            for (int j = i + 1; j < l; j++) {
                if (arrs[i] > arrs[j]) {
                    int temp = arrs[i];
                    arrs[i] = arrs[j];
                    arrs[j] = temp;
                }
            }
        }
        System.out.println("冒泡排序---后: " + Arrays.toString(arrs));
        return arrs;
    }

    /**
     * 冒泡排序 Bubble Sort 递归
     */
    public static int[] Test_Bsort2(int[] arrs) {
        for(int i = 0; i < arrs.length - 1 ; i++) {
            if (arrs[i] > arrs[i+1]) {
                int temp = arrs[i];
                arrs[i] = arrs[i+1];
                arrs[i+1] = temp;
                Test_Bsort2(arrs);
            }
        }
        return arrs;
    }

    public static void main(String[] args) {
        int[] arrs = {3, 1, 6, 4, 5, 0, 8, 9, 7, 2};

        // 01 冒泡排序
        Test_BSort(arrs);

        Test_Bsort2(arrs);
        System.out.println("冒泡排序(递归)---后: " +  Arrays.toString(arrs));
    }
}
