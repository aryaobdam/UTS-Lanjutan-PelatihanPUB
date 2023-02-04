package connection

import (
	"fmt"
	"log"

	"UTS-LANJUTAN-ARYA/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"postgres",
		"arya14",
		"bimbel_uts_lanjutan",
		"5432",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Program{}, &models.Student{}, &models.Teacher{}, &models.ProgramSD{}, &models.ProgramSMP{}, &models.ProgramSMAIpa{}, &models.ProgramSMAIps{})

	return db
}
