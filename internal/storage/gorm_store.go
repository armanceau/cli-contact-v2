package storage

import (
	"github.com/armanceau/cli-contact-v2/internal/database"
	"github.com/armanceau/cli-contact-v2/internal/models"
)

type GormStore struct{}

func NewGormStore() *GormStore {
	return &GormStore{}
}

func (g *GormStore) Ajouter(c Contact) Contact {
	contact := models.Contact{
		Nom:   c.Nom,
		Email: c.Email,
	}
	database.DB.Create(&contact)
	return Contact{
		ID:    int(contact.ID),
		Nom:   contact.Nom,
		Email: contact.Email,
	}
}

func (g *GormStore) Lister() []Contact {
	var contacts []models.Contact
	database.DB.Find(&contacts)

	var result []Contact
	for _, c := range contacts {
		result = append(result, Contact{
			ID:    int(c.ID),
			Nom:   c.Nom,
			Email: c.Email,
		})
	}
	return result
}

func (g *GormStore) Supprimer(ID int) bool {
	res := database.DB.Delete(&models.Contact{}, ID)
	return res.RowsAffected > 0
}

func (g *GormStore) MettreAJour(c Contact) (Contact, bool) {
	var contact models.Contact
	if err := database.DB.First(&contact, c.ID).Error; err != nil {
		return Contact{}, false
	}
	contact.Nom = c.Nom
	contact.Email = c.Email
	database.DB.Save(&contact)
	return Contact{ID: int(contact.ID), Nom: contact.Nom, Email: contact.Email}, true
}

func (g *GormStore) Recuperer(ID int) (Contact, bool) {
	var contact models.Contact
	if err := database.DB.First(&contact, ID).Error; err != nil {
		return Contact{}, false
	}
	return Contact{ID: int(contact.ID), Nom: contact.Nom, Email: contact.Email}, true
}

func (g *GormStore) NextID() int {
	var last models.Contact
	database.DB.Last(&last)
	return int(last.ID) + 1
}
