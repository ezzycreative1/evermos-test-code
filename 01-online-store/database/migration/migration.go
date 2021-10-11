package migration

import (
	"fmt"
	"log"

	"01-online-store/database"
	"01-online-store/models"
)

// RunRollback ..
func RunRollback() {

	db := database.PostsqlConn()

	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}

	if exist := db.HasTable("customers"); exist {
		fmt.Println("drop table customers")
		err := db.DropTable("customers")
		if err == nil {
			fmt.Println("success drop table customers")
		}
	}

	if exist := db.HasTable("categories"); exist {
		fmt.Println("drop table categories")
		err := db.DropTable("categories")
		if err == nil {
			fmt.Println("success drop table categories")
		}
	}

	if exist := db.HasTable("products"); exist {
		fmt.Println("drop table products")
		err := db.DropTable("products")
		if err == nil {
			fmt.Println("success drop table products")
		}
	}

	if exist := db.HasTable("products"); exist {
		fmt.Println("drop table products")
		err := db.DropTable("products")
		if err == nil {
			fmt.Println("success drop table products")
		}
	}

	if exist := db.HasTable("suppliers"); exist {
		fmt.Println("drop table suppliers")
		err := db.DropTable("suppliers")
		if err == nil {
			fmt.Println("success drop table suppliers")
		}
	}

	if exist := db.HasTable("orders"); exist {
		fmt.Println("drop table orders")
		err := db.DropTable("orders")
		if err == nil {
			fmt.Println("success drop table orders")
		}
	}

	if exist := db.HasTable("order_details"); exist {
		fmt.Println("drop table order_details")
		err := db.DropTable("order_details")
		if err == nil {
			fmt.Println("success drop table order_details")
		}
	}

}

// RunMigration ..
func RunMigration() {
	db := database.PostsqlConn()

	if db.Error != nil {
		log.Fatalln(db.Error.Error())
	}

	if exist := db.HasTable("customers"); !exist {
		fmt.Println("migrate table customers")
		err := db.CreateTable(&models.Customer{})
		if err == nil {
			fmt.Println("success migrate table customers")
		}
	}

	if exist := db.HasTable("categories"); !exist {
		fmt.Println("migrate table categories")
		err := db.CreateTable(&models.Category{})
		if err == nil {
			fmt.Println("success migrate table categories")
		}
	}

	if exist := db.HasTable("products"); !exist {
		fmt.Println("migrate table products")
		err := db.CreateTable(&models.Product{})
		if err == nil {
			fmt.Println("success migrate table products")
		}
	}

	if exist := db.HasTable("suppliers"); !exist {
		fmt.Println("migrate table suppliers")
		err := db.CreateTable(&models.Supplier{})
		if err == nil {
			fmt.Println("success migrate table suppliers")
		}
	}

	if exist := db.HasTable("orders"); !exist {
		fmt.Println("migrate table orders")
		err := db.CreateTable(&models.Order{})
		if err == nil {
			fmt.Println("success migrate table orders")
		}
	}

	if exist := db.HasTable("order_details"); !exist {
		fmt.Println("migrate table order_details")
		err := db.CreateTable(&models.OrderDetail{})
		if err == nil {
			fmt.Println("success migrate table order_details")
		}
	}
}
