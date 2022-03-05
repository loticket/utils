package stringx

import (
	"fmt"
	"math/rand"
	"time"
)

func StrSn(storeId int64) string {
	var dates string = time.Now().Local().Format("20060102150405")
	var start int = 100000
	var end int = 999999
	if storeId < 10 {
		start = 100000000
		end = 999999999
	} else if storeId < 100 && storeId > 9 { //10 - 99
		start = 10000000
		end = 99999999
	} else if storeId < 1000 && storeId > 99 { //100 - 999
		start = 1000000
		end = 9999999
	}

	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(end - start)
	random = start + random
	return fmt.Sprintf("%s%d%d", dates, storeId, random)
}
