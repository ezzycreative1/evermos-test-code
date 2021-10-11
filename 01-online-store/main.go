package main

import (
	"fmt"
	"log"
	"os"

	CustomerRepo "01-online-store/app/customer/repository"
	CustomerUsecase "01-online-store/app/customer/usecase"

	SupplierRepo "01-online-store/app/supplier/repository"
	SupplierUsecase "01-online-store/app/supplier/usecase"

	CategoryRepo "01-online-store/app/category/repository"
	CategoryUsecase "01-online-store/app/category/usecase"

	"01-online-store/database"
	"01-online-store/database/migration"
	"01-online-store/database/seeder"
	routes "01-online-store/routes"

	//"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	err error
)

func main() {

	err = godotenv.Load()
	db := database.PostsqlConn()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}
	dbEvent := os.Getenv("DBEVENT")
	if dbEvent == "rollback_migrate" || dbEvent == "rollback" {
		migration.RunRollback()
	}
	if dbEvent == "migrate_only" {
		migration.RunMigration()
	}
	if dbEvent == "migrate" || dbEvent == "rollback_migrate" {
		migration.RunMigration()
		// Type Seeders
		seeder.CategorySeeder()
	}

	router := gin.New()
	router.Use(gin.Recovery())

	// CORS
	router.Use(CORSMiddleware())
	//router.Use(cors.Default())

	//Logging
	//router.Use(middle.SetUp())

	// Size Images
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// Middleware Token
	//router.Use(middle.AuthMiddleware())

	// Customer
	customerRepo := CustomerRepo.NewCustomerRepository(db)
	customerUsecase := CustomerUsecase.NewCustomerUsecase(customerRepo)

	// Supplier
	supplierRepo := SupplierRepo.NewSupplierRepository(db)
	supplierUsecase := SupplierUsecase.NewSupplierUsecase(supplierRepo)

	// Category
	categoryRepo := CategoryRepo.NewCategoryRepository(db)
	categoryUsecase := CategoryUsecase.NewCategoryUsecase(categoryRepo)

	// Health check
	routes.HealthCheckHTTPHandler(router)
	routes.CustomerHTTPHandler(router, customerUsecase)
	routes.SupplierHTTPHandler(router, supplierUsecase)
	routes.CategoryHTTPHandler(router, categoryUsecase)

	// Server
	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))); err != nil {
		log.Fatal(err)
	}
}

// CORSMiddleware ..
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			//c.Next()
			return
		}

		c.Next()
	}
}
