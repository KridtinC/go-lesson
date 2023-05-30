package model

type Girl struct {
	Person
}

func (g Girl) Sing() string {
	return "Lalala"
}

// Override method
func (g *Girl) GetName() string {
	return "girl: " + g.name
}

func NewGirl(name string, age int) Girl {
	return Girl{
		Person: Person{
			name: name,
			age:  age,
		},
	}
}
