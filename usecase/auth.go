package usecase

type Login struct{}

type LoginInterface interface {
	Auth(username, password string) bool
}

func TaskLogin() LoginInterface {
	return &Login{}
}

func (pi *Login) Auth(username, password string) bool {
	if username == "ivialva" && password == "123456" {
		return true
	}
	return false
}
