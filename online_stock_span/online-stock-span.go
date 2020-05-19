package stock

type StockSpanner struct {
	stocks []stock
}

type stock struct {
	price  int
	weight int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	w := 1

	for len(this.stocks) > 0 && this.stocks[len(this.stocks)-1].price <= price {
		w += this.stocks[len(this.stocks)-1].weight
		this.stocks = this.stocks[:len(this.stocks)-1]
	}

	this.stocks = append(this.stocks, stock{price: price, weight: w})

	return w
}

/*
	Your StockSpanner object will be instantiated and called as such:
	obj := Constructor();
	param_1 := obj.Next(price);
*/

// https://leetcode.com/problems/online-stock-span/
// https://leetcode.com/submissions/detail/341654614/?from=/explore/featured/card/may-leetcoding-challenge/536/week-3-may-15th-may-21st/3334/

// tags: stack
