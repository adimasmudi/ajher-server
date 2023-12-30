package user

type UserFormatter struct {
	ID       int    `json:"userId"`
	FullName string `json:"name"`
	Email    string `json:"email"`
	Picture  string `json:"picture"`
	Username string `json:"username"`
	Gender   string `json:"gender"`
}

func FormatUser(user User) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Picture:  user.Picture,
		Username: user.Username,
		Gender:   user.Gender,
	}

	return formatter
}
