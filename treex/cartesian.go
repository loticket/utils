package treex


//实现商品规格的笛卡尔积
type GoodsSkus struct {
	SpecKey         int64
	SpecId          int64
	SpecName        string
	SpecValueKey   int64
	SpecValueId    int64
	SpecValue       string
}

func Cartesian(sets ...[]GoodsSkus) [][]GoodsSkus {
	lens := func(i int) int { return len(sets[i]) }
	product := [][]GoodsSkus{}
	for ix := make([]int, len(sets)); ix[0] < lens(0); nextIndex(ix, lens) {
		var r []GoodsSkus
		for j, k := range ix {
			r = append(r, sets[j][k])
		}
		product = append(product, r)
	}
	return product
}

func nextIndex(ix []int, lens func(i int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens(j) {
			return
		}
		ix[j] = 0
	}
}
