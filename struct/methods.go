package main

import (
	"fmt"
		
	"time"
)


type Signin interface {
	login() []string
}

// facebook user 
func (f Facebook) login() []string {
	user := []string{f.username, f.email}
	fmt.Println(" ")
	fmt.Println("SignedIn With Facebook")
	fmt.Printf("Facebook Username: %v\n", f.username)
	fmt.Println(" ")
	return user
}

// google user
func (g Google) login() []string {
	user := []string{g.username, g.email}
	fmt.Println(" ")
	fmt.Println("SignedIn With Google")
	fmt.Printf("Google Username: %v\n", g.username)
	fmt.Println(" ")
	return user
}

// logger to signin login types
func verify(s Signin) []string {
	fmt.Println(" ")
	fmt.Println("User signed in")
	fmt.Println(time.Now().Date())
	return s.login()
	
}