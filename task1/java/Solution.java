import java.util.HashMap;

class Solution {
    public int[] twoSum(int[] nums, int target) {

        HashMap<Integer, Integer> visited = new HashMap<Integer, Integer>();
        
        for ( int i = 0;  i < nums.length; i++) {
            if (visited.containsKey(target - nums[i])){
                return new int[]{i, visited.get(target - nums[i])};
            }

            visited.put(nums[i], i);
        }

        return new int[]{-1, -1};
    }
}