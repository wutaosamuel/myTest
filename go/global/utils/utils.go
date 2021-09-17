package utils

var Auth = map[string]string{
	"Zero": "0",
}

func SetAuth(key, value string) {
	Auth[key] = value
}
