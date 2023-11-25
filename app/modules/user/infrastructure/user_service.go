package infrastructure

type UserService struct{}

func (us *UserService) find() []User {
	return usersDummy
}
