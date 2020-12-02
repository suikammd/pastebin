package models

type PasteText struct {
	Content string `json:"content" gorm:"column:content"`
	Path    string `json:"path,omitempty" gorm:"column:path;primary_key"`
}
