package domain

type User struct {
	ID   string
	Name string
}

type UserRepository interface {
	GetID(id string) (error, User)
	GetName(name string) (error, User)
}

type UserUsecase interface {
	GetID(id string) (error, User)
	GetName(name string) (error, User)
}
