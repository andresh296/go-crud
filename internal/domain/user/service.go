package user

type Repository interface {
	GetUserByEmail(email string) (*User,error)
	Save(user User) error
}

type Service interface {
	GetUserByEmail(email string) (*User,error)
	Save(user User) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s service)GetUserByEmail(email string) (*User,error){
	return s.repository.GetUserByEmail(email)
}



func (s service) Save(user User) (User, error) {
	user.setID()
	user.hashPassword()
	err := s.repository.Save(user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
