package db

/*type Item struct {
	Price float64
}

func LoadItem(id int) *Item {
	return &Item{
		Price: 9.0001,
	}
}*/

import (
	"github.com/foysal-mamun/golangpractice/openmymind/ch04/shopping/models"
)

func LoadItem(id int) *models.Item {
	return &models.Item{
		Price: 9.0001,
	}
}
