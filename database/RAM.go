package database

import (
	"OCRsearch/helpers"
	"OCRsearch/models"
	"os"
	"strings"
)

type RAM struct {
	root models.RAMNode
}

var ramInstance *RAM

func RAMInstance() *RAM {
	if ramInstance == nil {
		ramInstance = &RAM{
			root: *models.NewRAMNode(),
		}
	}
	return ramInstance
}

func (r *RAM) GetRoot() models.Node {
	return &r.root
}

func (r *RAM) Initialize() error {
	items, _ := os.ReadDir("./database/data")

	for _, item := range items {
		if item.IsDir() {
			continue
		}
		filename := item.Name()
		jsonFile, err := os.Open("./database/data/" + filename)
		if err != nil {
			return err
		}
		result := helpers.ReadJSON(jsonFile)
		for frameID, words := range result {
			for _, word := range words.([]interface{}) {
				for _, splittedWord := range strings.Fields(helpers.NormalizeUnicode(word.(string))) {
					r.insertWord(splittedWord, models.NewRAMFrame(filename[:len(filename)-5], frameID))
				}
			}
		}

		jsonFile.Close()
	}

	return nil
}

func (r *RAM) insertWord(word string, frame *models.RAMFrame) {
	currentNode := &r.root

	for _, character := range word {
		nextNode := currentNode.GetChild(character)
		if nextNode == nil {
			nextNode = models.NewRAMNode()
			currentNode.AddChild(character, nextNode)
		}
		currentNode = nextNode.(*models.RAMNode)
	}

	currentNode.AddFrame(frame)
}
