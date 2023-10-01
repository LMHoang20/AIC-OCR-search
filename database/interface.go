package database

import "OCRsearch/models"

type Database interface {
	Initialize() error
	GetRoot() models.Node
}

func Instance(dbType string) Database {
	switch dbType {
	default:
		return RAMInstance()
	}
}
