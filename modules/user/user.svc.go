package user

type UserServices struct {
	repository *UserRepository
}

var Service UserServices

func (s *UserServices) GetUser() *[]User {
	return s.repository.GetUser()
}
