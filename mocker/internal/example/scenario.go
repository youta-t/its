//go:generate go run github.com/youta-t/its/mocker
//go:generate go run github.com/youta-t/its/structer

package example

import "errors"

type User struct {
	Id   string
	Name string
}

type UserRegistry interface {
	Get(userId string) (User, error)
	Update(User) error
	Delete(User) error
}

type SessionStore func(cookie string) (userId string, ok bool)

func UpdateUser(
	sess SessionStore,
	registry UserRegistry,
) func(cookie string, newName string) error {
	return func(cookie, newName string) error {
		userId, ok := sess(cookie)
		if !ok {
			return errors.New("you are not logged in")
		}
		user, err := registry.Get(userId)
		if err != nil {
			return err
		}
		user.Name = newName
		return registry.Update(user)
	}
}
