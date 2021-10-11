package seeder

import (
	"fmt"
	"time"

	"01-online-store/database"
	"01-online-store/models"
)

// CategorySeeder ..
func CategorySeeder() {

	db := database.PostsqlConn()

	now := time.Now()

	var cat []models.Category

	var buku = models.Category{
		ID:        1,
		Name:      "Buku",
		CreatedAt: now,
		UpdatedAt: now,
	}
	cat = append(cat, buku)

	var minuman = models.Category{
		ID:        2,
		Name:      "Minuman",
		CreatedAt: now,
		UpdatedAt: now,
	}
	cat = append(cat, minuman)

	for _, ct := range cat {
		if err := db.Create(&ct).Error; err != nil {
			fmt.Println(err.Error())

			//log.Fatalln(err.Error())
		}
		fmt.Printf("category %s has been created\n", ct.Name)
	}
}
