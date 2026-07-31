package auth

type Service interface {
	Login() error
	Logout() error
}

func NewService() Service {
	return &service{}
}

type service struct {
}

func (s *service) Login() error {
	return nil
}

func (s *service) Logout() error {
	return nil
}
