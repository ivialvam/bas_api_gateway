package usecase

type Login struct{}

type LoginInterface interface {
	Auth(Username, Password string) bool
}

func NewLogin() LoginInterface {
	return &Login{}
}

func (pi *Login) Auth(Username, Password string) bool {
	if Username == "admin" && Password == "admin123" {
		return true
	}
	return false
}
