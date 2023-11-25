package infrastructure

type Role struct {
	ID          string
	Name        string
	Description string
}

var rolesDummy = []Role{{
	ID:          "1",
	Name:        "Admin",
	Description: "Description",
}, {
	ID:          "2",
	Name:        "User",
	Description: "Description",
}}
