import java.util.Arrays;

/**
 * Create by $(xiliangMa) on 2019-10-03
 */
public class select_sort_test {
    public static int[] Test_Ssort(int[] arrs) {

        for (int i = 0; i < arrs.length - 1 ; i++) {
            int min = arrs[i];
            if (min > arrs[i+1]) {
                min = arrs[i+1];
                arrs[i+1] = arrs[i];
                arrs[i] = min;
                Test_Ssort(arrs);
            }
        }
        return arrs;
    }

    public static void main(String[] args) {
        int[] arrs = {3, 1, 6, 4, 5, 0, 8, 9, 7, 2};
        Test_Ssort(arrs);

        System.out.println("选择排序(递归)---后: " +  Arrays.toString(arrs));
    }
}
