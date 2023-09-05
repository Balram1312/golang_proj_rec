package database

import(
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"golang_test/models"
)

type App struct{
	DB *gorm.DB
}



func Connect() (*gorm.DB,error){
	dsn := "host=localhost user=postgres password=tiger dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		return nil,err
	}

	err = db.AutoMigrate(&models.Users{})
	if err != nil{
		return nil,err
	}

	return db,nil
}