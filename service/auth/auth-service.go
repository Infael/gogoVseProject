package auth

type AuthService interface {
	Login(email, password string) (string, error)
	Register(email, password string) error
}
