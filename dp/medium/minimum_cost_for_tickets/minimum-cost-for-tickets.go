package ticket

// Time complexity: O(N)
// Space complexity: O(N)
func mincostTickets(days []int, costs []int) int {
	dayMap := make(map[int]struct{}, len(days))
	for _, d := range days {
		dayMap[d] = struct{}{}
	}

	lastDay := days[len(days)-1]
	memo := make([]int, lastDay+1)

	return dy(1, dayMap, costs, lastDay, memo)
}

func dy(day int, dayMap map[int]struct{}, costs []int, lastDay int, memo []int) int {
	if day > lastDay {
		return 0
	}

	if memo[day] != 0 {
		return memo[day]
	}

	if _, ok := dayMap[day]; ok {
		memo[day] = minInts(
			costs[0]+dy(day+1, dayMap, costs, lastDay, memo),
			costs[1]+dy(day+7, dayMap, costs, lastDay, memo),
			costs[2]+dy(day+30, dayMap, costs, lastDay, memo),
		)

	} else {
		memo[day] = dy(day+1, dayMap, costs, lastDay, memo)
	}

	return memo[day]
}

func minInts(ints ...int) int {
	min := 1<<63 - 1
	for _, n := range ints {
		if n < min {
			min = n
		}
	}
	return min
}

// Minimum Cost For Tickets
// https://leetcode.com/problems/minimum-cost-for-tickets/

// tags: dp, dynamic programming
