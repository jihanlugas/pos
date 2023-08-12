package user

type AuthenticationRepository interface {
	//GetById() string
}

type authenticationRepository struct {
}

//func (r authenticationRepository) GetById() string {
//	return "Get By Id"
//}

func NewAuthenticationRepository() AuthenticationRepository {
	return authenticationRepository{}
}
