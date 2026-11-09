package student

type Student struct {
	Name string
	Age  int
}

func (s Student) Display() {
	println("Name:", s.Name)
	println("Age:", s.Age)
}
