package presenter

import (
	"go-clean-architecture/src/domain/model"
	"strconv"
)

type User struct {
	ID   string
	Name string
}

func (p *Presenter) User(user *model.User) *User {
	return &User{
		ID:   strconv.Itoa(user.ID),
		Name: user.Name,
	}
}

func (p *Presenter) Users(us []*model.User) []*User {
	var results []*User
	for _, u := range us {
		results = append(results, p.User(u))
	}
	return results
}
