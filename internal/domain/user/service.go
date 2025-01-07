package user

type Repository interface {
	GetByEmail(email string) (*User,error)
	Save(user User) error
}

type Service interface {
	GetByEmail(email string) (*User,error)
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

func (s service)GetByEmail(email string) (*User,error){
	return s.repository.GetByEmail(email)
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
