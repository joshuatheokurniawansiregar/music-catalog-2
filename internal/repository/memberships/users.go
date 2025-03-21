package memberships

import (
	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/models/memberships"
	"gorm.io/gorm"
)

// for register
func (r *Repository) CreateUser(model memberships.User) error {
	return r.db.Create(&model).Error
}

// for login
func(r *Repository) GetUser(email string, username string, id uint)(*memberships.User, error){
	var user memberships.User = memberships.User{}
	var response *gorm.DB = r.db.Where("email = ?", email).Or("username = ?", username).Or("id = ?", id).First(&user)
	if response.Error != nil{
		return nil, response.Error 
	}

	return &user, nil
}