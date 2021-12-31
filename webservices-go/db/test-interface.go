package db

type IDb interface {
	SayHello() string
}

type MyType struct{}

func (m *MyType) SayHello() string {
	return "Hello"
}
