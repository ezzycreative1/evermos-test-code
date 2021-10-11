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

}
