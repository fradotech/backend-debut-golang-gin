package infrastructure

type User struct {
	ID    string
	Name  string
	Email string
}

var usersDummy = []User{{
	ID:    "1",
	Name:  "Frado",
	Email: "Email",
}, {
	ID:    "2",
	Name:  "Frado2",
	Email: "Email",
}}
