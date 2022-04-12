package user

// Ingoing Port
type Repository interface {
	//Save(user Users) (Users, error)
	//Update(user Users) (Users, error)
	//Delete(Id int) (err error)
	FindById(Id int) (user *Users, err error)
	//FindAll() (users []Users, err error)
}