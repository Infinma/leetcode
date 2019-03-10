/**
 * mem = <target - value, first index of value>
 */

class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> mem = new HashMap<>();
        int ans[] = {-1, -1};
        for (int i = 0; i < nums.length; i++) {
            if (mem.containsKey(nums[i])) {
                ans[0] = mem.get(nums[i]);
                ans[1] = i;
                return ans;
            } else {
                mem.put(target - nums[i], i);
            }
        }
        return ans;
    }
}
