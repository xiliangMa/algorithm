import java.util.Arrays;

/**
 * Create by $(xiliangMa) on 2019-10-03
 */
public class insert_sort_test {


    public static void Test_Isort(){
        int[] arrs = {3, 1, 6, 4, 5, 0, 8, 9, 7, 2};

        for (int i = 0; i < arrs.length - 1; i++) {
            for (int j = i + 1; j > 0; j--) {
                if (arrs[j-1] > arrs[j]) {
                    int temp = arrs[j-1];
                    arrs[j-1] = arrs[j];
                    arrs[j] = temp;
                }
            }
        }
        System.out.println("插入排序---后: " +  Arrays.toString(arrs));
    }

    public static void main(String[] args) {
        Test_Isort();
    }
}
