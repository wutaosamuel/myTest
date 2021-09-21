package utils

var Auth = map[string]string{
	"Zero": "0",
}

func SetAuth(key, value string) {
	Auth[key] = value
}

var NewO = map[string]interface{} {
	"object": NewObject(),
}

type Object struct {
	Name string
	ID int
}

func NewObject() *Object {
	return &Object{
		Name: "",
		ID: -1,
	}
}