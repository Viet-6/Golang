package gormdemo

/*
	use gorm library to working with database
	install:
	> go get -u gorm.io/gorm
	> go get -u gorm.io/driver/<database>
*/

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// defined an model
type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float32
}

// make connection to mysql
func mysqlConnect() (db *gorm.DB, err error) {
	dsn := "root:mydb_p@ssw0rd@tcp(localhost:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

func GormD() {
	db, err := mysqlConnect()

	if err != nil {
		panic("failed to connect database")
	}

	// createStatement(db)
	queryData(db)

}

func createStatement(db *gorm.DB) {
	// product := Product{
	// 	Name:        "pro 1",
	// 	Description: "new product",
	// 	Price:       1.2,
	// }

	// products := []Product{
	// 	{
	// 		Name:        "pro 2",
	// 		Description: "new product 2",
	// 		Price:       99.89,
	// 	},
	// 	{
	// 		Name:        "pro 3",
	// 		Description: "new product 3",
	// 		Price:       199.89,
	// 	},
	// }

	// result := db.Create(&product)

	// fmt.Println(product)
	// fmt.Println(result.Error)
	// fmt.Println(result.RowsAffected)

	// result2 := db.Create(products)

	// fmt.Println(products)
	// fmt.Println(result2.Error)
	// fmt.Println(result2.RowsAffected)

	// // The reason why product need to add &(address pointer) and products was not:
	// // product is a struct and it a value type
	// // -> after execute a db.Create() it need to manipulate on that product data
	// // so we need to add & to point directly to where the product stored in the memory
	// // products in another hand, is a slice and it is a referance type
	// // referance type in this case is a slice is basicly like a list of addresss so we dont need to add & to it

	// // Create Record With Selected Fields
	// db.Select("Name", "Price").Create(&product)
	// fmt.Println(products)
	// fmt.Println(result2.Error)
	// fmt.Println(result2.RowsAffected)

	// product1 := Product{
	// 	Name:        "some product",
	// 	Description: "new product dclq",
	// 	Price:       1.5,
	// }
	// // Create Record & Ignore Seleceted Fields
	// db.Omit("Name").Create(&product1)
	// fmt.Println(product1)
	// fmt.Println(result2.Error)
	// fmt.Println(result2.RowsAffected)

	// // Batch Insert
	// // What it do is divide number of insert query and group it in a smaller batch
	// // EX: insert into tables (val1, val2) values (1, 2), (2, 3), ...
	errs := db.CreateInBatches(productSeeder(), 2)

	fmt.Println(*errs)

}

// Hooks
// Hooks is a function that are called before or after perform creation/querying/updating/delection
// if Hook get error it will rollback the transaction
// define Hooks
func (product *Product) BeforeSave(tx *gorm.DB) (err error) {

	if product.ID != 1 {
		err = errors.New("over limiting ID")
	}

	return
}

func productSeeder() (products []Product) {
	for i := 0; i < 500; i++ {
		product := Product{
			Name:        "product " + strconv.Itoa(rand.Intn(100)),
			Description: "new product " + strconv.Itoa(rand.Intn(100)),
			Price:       rand.Float32(),
		}
		products = append(products, product)
	}

	return
}

func queryData(db *gorm.DB) {
	product := Product{}
	db.First(&product)

	fmt.Println(product.Name)
}
