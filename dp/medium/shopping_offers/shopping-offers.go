package shopping

func shoppingOffers(price []int, special [][]int, needs []int) int {
	// special = append(special, priceToSpecial(price)...)
	return 0
}

// func bruteForce(special [][]int, needs []int) int {
// 	if isEnough(needs) {
// 		return 0
// 	}

// 	minPrice := 0

// 	for i := 0; i < len(special); i++ {
// 		price := bruteForce(special, needs)

// 	}

// 	return minPrice
// }

// func isEnough(needs []int) bool {
// 	for _, v := range needs {
// 		if v != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func priceToSpecial(price []int) [][]int {
// 	special := make([][]int, len(price))

// 	for i := 0; i < len(price); i++ {
// 		special[i] = make([]int, len(price)+1)
// 		special[i][i] = 1
// 		special[i][len(price)] = price[i]
// 	}

// 	return special
// }

// Shopping Offers
// https://leetcode.com/problems/shopping-offers/

// tags: dp, dynamic programming
