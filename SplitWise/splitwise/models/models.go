package models

//* This represents the entities for our application

//* Here we will be having main entities of our application like : User, Group, Expense, OwnedBy(Debt)

type User struct {
	ID    int
	Name  string
	Email string
}

type Group struct {
	ID       int
	Name     string
	Members  []User
	Expenses []Expense
}

type Expense struct {
	ID          int
	Description string
	Amount      float64
	PaidBy      int //UserId
	OwedBy      []OwedBy
}

type OwedBy struct {
	UserId      int
	ShareAmount float64
}
