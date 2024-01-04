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

func (r *repository) Save(otp Otp) (Otp, error) {
	err := r.db.Create(&otp).Error

	if err != nil {
		return otp, err
	}

	return otp, nil
}

func (r *repository) FindByOtpCode(otpCode string) (Otp, error) {
	var otp Otp
	err := r.db.Where("otpcode=?", otpCode).Find(&otp).Error

	if err != nil {
		return otp, err
	}

	return otp, nil
}

func (r *repository) Update(otp Otp) (Otp, error) {
	err := r.db.Save(&otp).Error

	if err != nil {
		return otp, err
	}

	return otp, nil
}
