package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// houses on a line
func rob(nums []int) int {
    var pre1 int    // best amount robbed before previous house
    var pre2 int    // best amount robbed before current house

    var cur int
    for _, num := range nums {
        cur = max(pre1 + num, pre2)
        pre1 = pre2
        pre2 = cur
    }

    return cur
}

// houses in a circle
func rob2(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    // return the maximum of robbing the first house or not robbing it
    return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

func rob3(root *TreeNode) int {
	robRoot, notRobRoot := dfs(root)
	return max(robRoot, notRobRoot)
}

func dfs(cur *TreeNode) (robCurrent int, notRobCurrent int) {
	if cur == nil {
		return 0, 0
	}
	
	leftRob, leftNotRob := dfs(cur.Left)
	rightRob, rightNotRob := dfs(cur.Right)

	robCurrent = cur.Val + leftNotRob + rightNotRob
	notRobCurrent = max(leftRob, leftNotRob) + max(rightRob, rightNotRob)

	return robCurrent, notRobCurrent
}

func main() {
	// test case 1
	nums1 := []int{1, 2, 3, 1}
	println(rob(nums1)) // expected output: 4
	println(rob2(nums1)) // expected output: 4
	println(rob2([]int{2, 3, 2})) // expected output: 3

	// test case 2
	nums2 := []int{2, 7, 9, 3, 1}
	println(rob(nums2)) // expected output: 12
	println(rob2(nums2)) // expected output: 12
	println(rob2([]int{1, 2, 3})) // expected output: 3
}