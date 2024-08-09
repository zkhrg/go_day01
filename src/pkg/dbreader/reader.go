package dbreader

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
)

type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit" xml:"itemunit"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Recipe struct {
	Cakes []Cake `json:"cake" xml:"cake"`
}

type JSONDBReader struct {
	recipe *Recipe
}

type XMLDBReader struct {
	recipe *Recipe
}

type DBReader interface {
	ReadFile(string) error
	Print() error
}

func (s *JSONDBReader) ReadFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()
	decoder := json.NewDecoder(file)
	tokens_to_skip := 3

	for i := 0; i < tokens_to_skip; i++ {
		_, err = decoder.Token()
		if err != nil {
			return err
		}
	}

	var recipe Recipe
	for decoder.More() {
		var cake Cake
		err := decoder.Decode(&cake)
		if err != nil {
			return err
		}
		recipe.Cakes = append(recipe.Cakes, cake)
	}

	_, err = decoder.Token()
	if err != nil {
		return err
	}
	s.recipe = &recipe
	return err
}

func (s *JSONDBReader) Print() error {
	if s.recipe == nil {
		return errors.New("recipe is empty")
	}
	indent := "    "
	data, err := xml.MarshalIndent(s.recipe, "", indent)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func (s *XMLDBReader) ReadFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := xml.NewDecoder(file)

	var recipe Recipe
	for {
		token, err := decoder.Token()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}

		switch token := token.(type) {
		case xml.StartElement:
			if token.Name.Local == "cake" {
				var cake Cake
				if err := decoder.DecodeElement(&cake, &token); err != nil {
					return err
				}
				recipe.Cakes = append(recipe.Cakes, cake)
			}
		}
	}
	s.recipe = &recipe
	return err
}

func (s *XMLDBReader) Print() error {
	if s.recipe == nil {
		return errors.New("recipe is empty")
	}
	indent := "    "
	data, err := json.MarshalIndent(s.recipe, "", indent)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
