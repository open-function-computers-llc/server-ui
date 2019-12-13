package models

type model interface {
	GetSchema(string) string
}
