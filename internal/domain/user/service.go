package user

type Repository interface {
	Save(user User) error
}

type Service interface {
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



func (s service) Save(user User) (User, error) {
	user.setID()

	user.hashPassword()

	err := s.repository.Save(user)
	if err != nil {
		return User{}, err
	}
	
	return user, nil
}