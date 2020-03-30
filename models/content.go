package models


type Content struct {
	Id       int
	Name     string
	Root     string
	Content  string
	SPage    string
	XPage    string
	List     string
}

func (m *Content) TableName() string {
	return TableName("content")
}