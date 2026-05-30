public class Solution {
    public static void main(String[] args) {
        // test case 1
        int[] nums = {10, 9, 2, 5, 3, 7, 101, 18};
        System.out.println(lengthOfLIS(nums));

        // test case 2
        int[] nums2 = {0, 1, 0, 3, 2, 3};
        System.out.println(lengthOfLIS(nums2));

        // test case 3
        int[] nums3 = {7, 7, 7, 7, 7, 7, 7};
        System.out.println(lengthOfLIS(nums3));
    }

    // Approach 1: Dynamic Programming
    public static int lengthOfLIS1(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }

        int n = nums.length;
        int[] dp = new int[n];
        int maxLength = 1;

        for (int i = 0; i < n; i++) {
            dp[i] = 1; // Each element is an increasing subsequence of length 1
            for (int j = 0; j < i; j++) {
                if (nums[i] > nums[j]) {
                    dp[i] = Math.max(dp[i], dp[j] + 1);
                }
            }
            maxLength = Math.max(maxLength, dp[i]);
        }

        return maxLength;
    }

    // Approach 2: Binary Search
    public static int lengthOfLIS(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }

        int n = nums.length;
        int[] tails = new int[n];
        int size = 0;
        for (int num : nums) {
            int lo = 0, hi = size;
            while (lo < hi) {
                int mid = lo + (hi - lo) / 2;
                if (tails[mid] < num) {
                    lo = mid + 1;
                } else {
                    hi = mid;
                }
            }
            tails[lo] = num;
            if (lo == size) {
                size++;
            }
        }
        return size;
    }
}
