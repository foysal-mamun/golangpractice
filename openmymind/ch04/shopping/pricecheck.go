package shopping

import (
	"github.com/foysal-mamun/golangpractice/openmymind/ch04/shopping/db"
)

func PriceCheck(itemId int) (float64, bool) {
	item := db.LoadItem(itemId)
	if item == nil {
		return 0, false
	}

	return item.Price, true
}
