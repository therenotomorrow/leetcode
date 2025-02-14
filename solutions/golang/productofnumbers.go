package golang

type ProductOfNumbers struct {
	data []int
}

func ProductOfNumbersConstructor() ProductOfNumbers {
	return ProductOfNumbers{data: []int{1}}
}

func (p *ProductOfNumbers) Add(num int) {
	if num == 0 {
		p.data = []int{1}
	} else {
		p.data = append(p.data, num*p.data[len(p.data)-1])
	}
}

func (p *ProductOfNumbers) GetProduct(k int) int {
	size := len(p.data) - 1

	if k > size {
		return 0
	}

	return p.data[size] / p.data[size-k]
}
