package authentication

type Repository interface {
	//GetById() string
}

type authenticationRepository struct {
}

//func (r authenticationRepository) GetById() string {
//	return "Get By Id"
//}

func NewAuthenticationRepository() Repository {
	return authenticationRepository{}
}
