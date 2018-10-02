package main

import (
	"fmt"
	"strconv"
)

type User struct {
	id        int
	firstName string
	lastName  string
	email     string
	userRate  float64
	mobile    string
}

func getUserRate(userId int) float64 {
	if userId == 155 {
		return 499998.23
	} else {
		return 0.0
	}
}

// (value receiver) Method
func (u User) information() string {
	return "Hi! My name is " + u.firstName + " " + u.lastName + ". My UserID is " + strconv.Itoa(u.id) + ". My email address " + u.email + ". Hourly wage => " + strconv.FormatFloat(u.userRate, 'f', 6, 64) + "."
}

// (pointer receiver) Method
func (u *User) updateEmail() {
	u.email = "harish31@flowace.in"
}

// (pointer receiver) Method to update userRate
func (u *User) upadteUserRate(userRate float64) {
	u.userRate = userRate
}

func main() {

	user1 := User{id: 155, firstName: "Harish", lastName: "Matta", email: "harish@flowace.in", mobile: "+91 9866863834"}

	user1.upadteUserRate(getUserRate(user1.id))

	fmt.Println(user1)

	fmt.Println(user1.information())
}
