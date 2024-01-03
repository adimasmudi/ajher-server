package otp

import "gorm.io/gorm"

type Repository interface {
	Save(otp Otp) (Otp, error)
	FindByOtpCode(otpCode string) (Otp, error)
	Update(otp Otp) (Otp, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// func (r *repository) Save(otp Otp) (Otp, error){

// }
// func (r *repository) FindByOtpCode(otpCode string) (Otp, error){

// }
// func (r *repository) Update(otp Otp) (Otp, error){

// }
