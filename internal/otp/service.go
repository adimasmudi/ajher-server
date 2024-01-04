package otp

import (
	"errors"
	"time"
)

type Service interface {
	VerifyOtp(input VerifyOtpInput) (Otp, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) VerifyOtp(input VerifyOtpInput) (Otp, error) {

	otp, err := s.repository.FindByOtpCode(input.OtpCode)

	if err != nil {
		return otp, err
	}

	if otp.ID == 0 {
		return otp, errors.New("otp code is not found")
	}

	if time.Now().UTC().After(otp.ValidUntil) {
		return otp, errors.New("otp is expired")
	}

	if otp.Status != "valid" {
		return otp, errors.New("otp is not valid")
	}

	otp.Status = "invalid"
	updatedOtp, err := s.repository.Update(otp)

	if err != nil {
		return otp, err
	}

	return updatedOtp, nil
}
