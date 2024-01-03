package domain

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: User{}},
		{Model: Student{}},
		{Model: Teacher{}},
		{Model: Class{}},
		{Model: Subject{}},
		{Model: SchoolYear{}},
		{Model: Schedule{}},
		{Model: Admin{}},
	}
}