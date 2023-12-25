package domain

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: User{}},
		{Model: Student{}},
		{Model: Teacher{}},
	}
}