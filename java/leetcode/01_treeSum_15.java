class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        int n = nums.length;
        Set<List<Integer>> data = new HashSet<>();

        for(int i = 0; i < n; i++){
            int l = i + 1;
            int r = n - 1;
            while (l < r){
                if (nums[i] + nums[l] + nums[r] == 0){
                    data.add(Arrays.asList(nums[i], nums[l], nums[r]));
                    l++;
                    r--;
                } else if(nums[i] + nums[l] + nums[r] < 0) {
                    l++;
                } else {
                    r--;
                }
            }
        }
        List<List<Integer>> result = new ArrayList<>();
        result.addAll(data);
        return result;
    }
}