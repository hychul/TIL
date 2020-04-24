package main

import (
	"fmt"
	"html/template"
	"os"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func (u User) IsAdmin() bool {
	return u.ID == 0
}

func main() {
	user0 := User{ID: 0, Name: "hychul", Email: "hychome@gmail.com"}
	user1 := User{ID: 1, Name: "tester", Email: "test@test.com"}
	users := []User{user0, user1}

	templ, err := template.New("Tmpl0").Parse("ID : {{.ID}}\nName : {{.Name}}\nEmail: {{.Email}}")
	if err != nil {
		panic(err)
	}

	fmt.Println("template string\n")

	templ.Execute(os.Stdout, user0)
	fmt.Println()
	templ.Execute(os.Stdout, user1)

	fmt.Println("\ntemplate file\n")

	templFile, err := template.New("Tmpl1").ParseFiles("templates/tmpl0.tmpl", "templates/tmpl1.tmpl", "templates/tmpl2.tmpl")
	if err != nil {
		panic(err)
	}

	templFile.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user0)
	fmt.Println()
	templFile.ExecuteTemplate(os.Stdout, "tmpl1.tmpl", user1)

	fmt.Println("\ntemplate file range\n")

	templFile.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)

}
