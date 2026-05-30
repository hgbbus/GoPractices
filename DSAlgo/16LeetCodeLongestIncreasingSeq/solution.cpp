#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    // Approach 1: Dynamic Programming
    int lengthOfLIS1(vector<int>& nums) {
        int n = nums.size();
        if (n == 0) return 0;
        
        int dp[n];
        int ans = 1;
        for (int i = 0; i < n; i++) {
            dp[i] = 1;
            for (int j = 0; j < i; j++) {
                if (nums[i] > nums[j]) {
                    dp[i] = max(dp[i], dp[j] + 1);
                }
            }
            ans = max(ans, dp[i]);
        }
        
        return ans;
    }

    // Approach 2: Dynamic Programming with Binary Search (Patience Sorting)
    int lengthOfLIS2(vector<int>& nums) {
        vector<int> tails;
        for (int num : nums) {
            auto it = lower_bound(tails.begin(), tails.end(), num);
            if (it == tails.end()) {
                tails.push_back(num);
            } else {
                *it = num;
            }
        }
        return tails.size();
    }

    // Approach 2a: Dynamic Programming with Binary Search (Patience Sorting) - using manual binary search
    int lengthOfLIS(vector<int>& nums) {
        int n = nums.size();
        if (n == 0) return 0;

        int tails[n];
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
};

int main() {
    Solution solution;
    vector<int> nums1 = {10, 9, 2, 5, 3, 7, 101, 18};
    vector<int> nums2 = {0, 1, 0, 3, 2, 3};
    vector<int> nums3 = {7, 7, 7, 7, 7, 7, 7};
    
    cout << solution.lengthOfLIS(nums1) << endl; // Output: 4
    cout << solution.lengthOfLIS(nums2) << endl; // Output: 4
    cout << solution.lengthOfLIS(nums3) << endl; // Output: 1
    
    return 0;
}

