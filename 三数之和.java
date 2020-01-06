class Solution {
    public List<List<Integer>> threeSum(int[] nums) {
        Arrays.sort(nums);
        List<List<Integer>> res = new ArrayList<>();
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) map.put(nums[i], i); // val , last index

        for(int i = 0 ; i < nums.length - 2; i++){
            for(int j = i + 1; j < nums.length -1; j++){
                int target = 0 - nums[i] - nums[j];

                if (map.containsKey(target) && map.get(target) > j){
                    res.add(Arrays.asList(nums[i], nums[j], target));
                    j = map.get(nums[j]); // Taking the last index of j to remove duplicate 
                }
            }
            i = map.get(nums[i]); // Taking the last index of i to remove duplicate
        }

        return res;
    }
}
