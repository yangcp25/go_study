package facade

type IUser interface {
	Login(name string, password string) (*User, error)
	Register(name string, password string) (*User, error)
}

type User struct {
	name string
}

type UserService struct{}

func (user UserService) Login(name string, password string) (*User, error) {
	return &User{name: name}, nil
}

func (user UserService) Register(name string, password string) (*User, error) {
	return &User{name: name}, nil
}

type UserFacade interface {
	UserLoginOrRegister(name string, password string) (*User, error)
}

func (user UserService) UserLoginOrRegister(name string, password string) (*User, error) {
	userObj, err := user.Login(name, password)
	if err != nil {
		return nil, err
	}

	if userObj != nil {
		return userObj, nil
	}

	return user.Register(name, password)
}
