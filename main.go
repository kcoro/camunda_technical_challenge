package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// UsersAPI represents the api's json response
// Only concerned with needed data, total pages, and the array of user data
type UsersAPI struct {
	TotalPages int       `json:"total_pages"`
	Data       []UserDTO `json:"data"`
}

// UserDTO represents the structure of a user object in the json response
type UserDTO struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// queryAPI is a function that calls the REST API endpoint,
// converts json to a go data struct, pulls out user structs,
// and returns a list of users.
func queryAPI() []UserDTO {
	apiJson := UsersAPI{}
	users := make([]UserDTO, 0)
	pageNum := 1
	totalPages := 1

	// Make initial get request to REST API endpoint
	// If there are additional pages, make additional requests
	for pageNum <= totalPages {
		apiURL := "https://reqres.in/api/users?page=" + strconv.Itoa(pageNum)
		response, err := http.Get(apiURL)
		// Check if error was returned by http.Get(), log error
		if err != nil {
			log.Fatal(err)
		}

		// Retrieve response body from *Http.Response
		body, _ := io.ReadAll(response.Body)
		defer response.Body.Close()

		// Convert json to struct
		json.Unmarshal(body, &apiJson)

		// Get total number of pages of data available
		totalPages = apiJson.TotalPages
		pageNum++

		// Iterate over api's returned list of users
		// Add each user to users slice
		for _, user := range apiJson.Data {
			newUser := UserDTO{
				Id:        user.Id,
				Email:     user.Email,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}

			users = append(users, newUser)
		}
	}

	return users
}

func main() {
	// Run api query, and retrieve list of users
	users := queryAPI()

	// Print first and last names of users to stdout.
	for _, user := range users {
		fmt.Println(user.FirstName + " " + user.LastName)
	}
}
