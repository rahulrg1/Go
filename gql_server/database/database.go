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
var location =[]string{"delhi","kolkata","mumbai"} 
var Data = []*model.User{
	&model.User{ID: uuid.New().String(), Name: "rr", Location: &location[0], Age: 21},
	&model.User{ID: uuid.New().String(), Name: "gg", Location: &location[1], Age: 22},
	&model.User{ID: uuid.New().String(), Name: "ii", Location: &location[2], Age: 23},
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

func UpdateUser(id string,input model.UpdateUser) (*model.User, error) {
	var user *model.User
	for _,u:= range Data {
		if(id==u.ID) {
			u.Name = input.Name
			u.Location = input.Location
			u.Age=input.Age
			user=u
		}
	}
	// if(user.ID!=id) {
	// 	return user, nil
	// }
	return user,nil
}