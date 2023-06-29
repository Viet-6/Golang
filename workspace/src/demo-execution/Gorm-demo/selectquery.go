package gormdemo

import (
	"fmt"

	"gorm.io/gorm"
)

func selectQuery(db *gorm.DB) {
	// query the first record order by ID
	product := Product{}

	db.First(&product)
	fmt.Println(product)

	product = Product{}
	// Get one record, no specified order
	fmt.Println(product)
	db.Take(&product)
	fmt.Println(product)

	// Get the newest record order by ID (id desc)
	product = Product{}
	db.Last(&product)
	fmt.Println(product)

	// Get all records in database
	products := []Product{}

	db.Find(products)
	fmt.Println(products)

	// To avoid the record not found error, use Find() instead.
	// Remember to add Limit(1) else it will take all the records
	product = Product{}
	result := db.Limit(1).Find(&product)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)

}
