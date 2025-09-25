package database

import (
	"log"

	"github.com/armanceau/cli-contact-v2/internal/models"
	"gorm.io/gorm"
	// "gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("contacts.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Echec de la connexion à la base de données : %v", err)
	}
	log.Println("Connexion à la base de données SQLite réussie !")

	err = DB.AutoMigrate(&models.Contact{})
	if err != nil {
		log.Fatalf("Échec de la migration de la base de données pour Product: %v", err)
	}
	log.Println("Migration de la base de données pour Product réussie !")
}
