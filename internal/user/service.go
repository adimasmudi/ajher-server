package user

import (
	"ajher-server/internal/otp"
	"ajher-server/utils"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterUserInput) (User, error)
	Login(input LoginUserInput) (User, error)
	GetUserById(userId string) (User, error)
	GoogleAuth(input GoogleOAuthInput) (User, error)
	GenerateAndSendEmail(input ResetPasswordInput) (otp.Otp, error)
	ChangePassword(input ChangePasswordUserInput) (User, error)
}

type service struct {
	repository    Repository
	otpRepository otp.Repository
}

func NewService(repository Repository, otpRepository otp.Repository) *service {
	return &service{repository, otpRepository}
}

func (s *service) Register(input RegisterUserInput) (User, error) {

	user := User{}
	if !utils.IsEmailValid(input.Email) {
		return user, errors.New("email is not valid")
	}

	user, err := s.repository.FindByEmail(input.Email, "users")

	if err != nil {
		return user, err
	}

	if user.ID != "" {
		return user, errors.New("user already exist")
	}

	user.Email = input.Email
	user.Username = input.Username
	user.FullName = input.Username
	user.LastLogin = time.Now()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)

	newUser, err := s.repository.Save(user, "users")

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginUserInput) (User, error) {
	user, err := s.repository.FindByEmail(input.Email, "users")

	if err != nil {
		return user, errors.New("user with that email doesn't exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return user, err
	}

	user.LastLogin = time.Now()
	user.UpdatedAt = time.Now()

	newUpdatedLoginUser, err := s.repository.Update(user, "users")

	if err != nil {
		return newUpdatedLoginUser, err
	}

	return user, nil
}

func (s *service) GetUserById(userId string) (User, error) {
	user, err := s.repository.GetById(userId, "users")

	if err != nil {
		return user, err
	}

	if user.ID == "" {
		return user, errors.New("user with that id doesn't exist")
	}

	return user, nil
}

func (s *service) GoogleAuth(input GoogleOAuthInput) (User, error) {
	var user User

	response, err := utils.VerifyIdToken(input.OAuthAccessToken)

	if err != nil {
		return user, errors.New("error validation token")
	}

	var googleUser utils.GoogleUser

	googleUser.Id = response.UserId
	googleUser.Email = response.Email
	googleUser.Picture = response.Audience
	googleUser.VerifiedEmail = response.VerifiedEmail

	isUserExist, err := s.repository.FindByEmail(googleUser.Email, "users")

	if err == nil && isUserExist.ID != "" {
		isUserExist.ID = "document_id_from_firestore"
		isUserExist.LastLogin = time.Now()
		isUserExist.UpdatedAt = time.Now()

		newUpdatedLoginUser, err := s.repository.Update(isUserExist, "users")

		if err != nil {
			return newUpdatedLoginUser, err
		}
		return isUserExist, nil
	}

	userName := strings.Split(googleUser.Email, "@")[0]

	user.Email = googleUser.Email
	user.Username = userName
	user.FullName = userName
	user.Picture = googleUser.Picture
	user.LastLogin = time.Now()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	newUser, err := s.repository.Save(user, "users")

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) GenerateAndSendEmail(input ResetPasswordInput) (otp.Otp, error) {
	otp := otp.Otp{}
	if !utils.IsEmailValid(input.Email) {
		return otp, errors.New("email is not valid")
	}

	user, err := s.repository.FindByEmail(input.Email, "users")

	if err != nil {
		return otp, err
	}

	if user.ID == "" {
		return otp, errors.New("user with that email doesn't exist")
	}

	otpString := utils.EncodeToString(4)

	to := []string{user.Email}
	cc := []string{os.Getenv("CONFIG_AUTH_EMAIL")}
	subject := "Reset Password OTP Code"
	message := fmt.Sprintf("Your OTP is %s", otpString)

	err = utils.SendMail(to, cc, subject, message)

	if err != nil {
		return otp, err
	}

	otp.UserId = user.ID
	otp.Otpcode = otpString
	otp.Status = "valid"
	otp.ValidUntil = time.Now().UTC().Add(time.Minute)

	savedOtp, err := s.otpRepository.Save(otp, "otps")

	if err != nil {
		return otp, err
	}

	return savedOtp, nil
}

func (s *service) ChangePassword(input ChangePasswordUserInput) (User, error) {
	user := User{}

	if input.Password != input.ConfirmPassword {
		return user, errors.New("password is not same with confirm password")
	}
	_, err := s.otpRepository.FindByOtpCode(input.OtpCode, "otps")

	if err != nil {
		return user, err
	}

	oldUser, err := s.repository.GetById("document_id_from_firestore", "users")

	if err != nil {
		return user, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	oldUser.Password = string(passwordHash)
	oldUser.UpdatedAt = time.Now()

	newUser, err := s.repository.Update(oldUser, "users")

	if err != nil {
		return user, err
	}

	return newUser, nil
}
