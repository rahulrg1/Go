package database

import (
	"gql_server/graph/model"

	"github.com/google/uuid"
)

// type Database struct {
// 	Id       string `json:"id,omitempty"`
// 	Name     string `json:"name,omitempty"`
// 	Location string `json:"location,omitempty"`
// 	Age      int    `json:"age,omitempty`
// }

var Data = []*model.User{
	&model.User{ID: uuid.New().String(), Name: "rr", Location: "delhi", Age: 21},
	&model.User{ID: uuid.New().String(), Name: "gg", Location: "pune", Age: 22},
	&model.User{ID: uuid.New().String(), Name: "ii", Location: "assam", Age: 23},
}

func AddData(input *model.User) []*model.User {

	Data = append(Data, input)
	return Data
}

func CreateUser(input model.NewUser) (*model.User, error) {
	uuid := uuid.New().String()
	newdata := &model.User{
		ID:       uuid,
		Name:     input.Name,
		Location: input.Location,
		Age:      input.Age,
	}
	AddData(newdata)
	return newdata, nil
}

func GetUser() ([]*model.User, error) {
	return Data, nil
}

func DeleteUser(id string) (*model.User, error) {
	var user *model.User
	var j int
	for i, u := range Data {
		if id == Data[i].ID {
			user = u
			j = i
		}
	}
	Data = append(Data[:j], Data[j+1:]...)
	return user, nil
}
