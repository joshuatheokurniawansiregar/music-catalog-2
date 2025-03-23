package memberships

import "gorm.io/gorm"

type(
	User struct{
		*gorm.Model
		Email string `gorm:"type:varchar(255);unique; not null"`
		Username string `gorm:"type:varchar(255);unique; not null"`
		Password string `gorm:"not null"`
		CreatedBy string `gorm:"type:varchar(255); not null"`
		UpdatedBy string `gorm:"type:varchar(255); not null"`
	}
)

type(
	SignUpRequest struct{
		Email string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginRequest struct{
		Email string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

type(
	LoginResponse struct{
		AccessToken string `json:"accessToken"`
	}
)