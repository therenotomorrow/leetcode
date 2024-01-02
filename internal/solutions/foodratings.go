package solutions

import "container/heap"

type FoodItem struct {
	food    string
	rate    int
	cuisine string
}

type FoodItemsPQ struct {
	data []*FoodItem
}

func (pq *FoodItemsPQ) Len() int {
	return len(pq.data)
}

func (pq *FoodItemsPQ) Less(i, j int) bool {
	if pq.data[i].rate == pq.data[j].rate {
		return pq.data[i].food <= pq.data[j].food
	}

	return pq.data[i].rate > pq.data[j].rate
}

func (pq *FoodItemsPQ) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq *FoodItemsPQ) Push(x interface{}) {
	val, ok := x.(*FoodItem)

	if !ok {
		panic("wrong FoodItem")
	}

	pq.data = append(pq.data, val)
}

func (pq *FoodItemsPQ) Pop() interface{} {
	old := pq.data
	n := len(old)
	item := old[n-1]
	pq.data = old[:pq.Len()-1]

	return item
}

type FoodRatings struct {
	data map[string]*FoodItemsPQ
	rate map[string]*FoodItem
}

func FoodRatingsConstructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	data := make(map[string]*FoodItemsPQ, len(cuisines))
	rate := make(map[string]*FoodItem)

	for i, cuisine := range cuisines {
		if _, ok := data[cuisine]; !ok {
			data[cuisine] = &FoodItemsPQ{data: nil}
		}

		rate[foods[i]] = &FoodItem{food: foods[i], rate: ratings[i], cuisine: cuisine}

		heap.Push(data[cuisine], rate[foods[i]])
	}

	return FoodRatings{data: data, rate: rate}
}

func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	fr.rate[food].food = food
	fr.rate[food].rate = newRating

	heap.Init(fr.data[fr.rate[food].cuisine])
}

func (fr *FoodRatings) HighestRated(cuisine string) string {
	return fr.data[cuisine].data[0].food
}
