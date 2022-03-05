package treex

import (
	"fmt"
	"testing"
)

func TestCat(t *testing.T) {
	var GoodsSkusArr1 []GoodsSkus = []GoodsSkus{
		GoodsSkus{
			SpecKey:      0,
			SpecId:       10005,
			SpecName:     "大小",
			SpecValueKey: 0,
			SpecValueId:  10012,
			SpecValue:    "1x",
		},
		GoodsSkus{
			SpecKey:      0,
			SpecId:       10005,
			SpecName:     "大小",
			SpecValueKey: 3,
			SpecValueId:  10013,
			SpecValue:    "2x",
		},
		GoodsSkus{
			SpecKey:      0,
			SpecId:       10005,
			SpecName:     "大小",
			SpecValueKey: 3,
			SpecValueId:  10015,
			SpecValue:    "3x",
		},
		GoodsSkus{
			SpecKey:      0,
			SpecId:       10005,
			SpecName:     "大小",
			SpecValueKey: 4,
			SpecValueId:  10016,
			SpecValue:    "4x",
		},
	}

	var GoodsSkusArr2 []GoodsSkus = []GoodsSkus{
		GoodsSkus{
			SpecKey:      1,
			SpecId:       10006,
			SpecName:     "喜好",
			SpecValueKey: 0,
			SpecValueId:  10014,
			SpecValue:    "抹抹嘴",
		},
		GoodsSkus{
			SpecKey:      1,
			SpecId:       10006,
			SpecName:     "喜好",
			SpecValueKey: 1,
			SpecValueId:  10017,
			SpecValue:    "耸耸肩",
		},
	}

	var GoodsSkusArr3 []GoodsSkus = []GoodsSkus{
		GoodsSkus{
			SpecKey:      2,
			SpecId:       10007,
			SpecName:     "容量",
			SpecValueKey: 0,
			SpecValueId:  10018,
			SpecValue:    "8+128",
		},
		GoodsSkus{
			SpecKey:      2,
			SpecId:       10007,
			SpecName:     "容量",
			SpecValueKey: 1,
			SpecValueId:  10019,
			SpecValue:    "8+256",
		},
	}

	dd := CartesianArr(GoodsSkusArr1, GoodsSkusArr2, GoodsSkusArr3)
	fmt.Println(len(dd))
	fmt.Println(dd)
}
