package validate

var Messages = map[string]string{
	"email":    "The %s field must be an valid email%s",
	"eqfield":  "The %s field is not equal to the %s field",
	"required": "The %s field is required%s",
	"url":      "The %s field must be an valid url (e.g. https://www.example.com)%s",
}
