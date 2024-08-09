package dbreader

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

// type DBReader interface {
// 	OpenFile(string)
// 	ReadFile(string) (Recipe, error)
// }

// type JSONDBReader struct {
// 	current_cake string
// }

// func (s *JSONDBReader) OpenFile(filename string) (Recipe, error) {
// 	return Recipe{}, nil
// }

// func (s *JSONDBReader) ReadCakeFe
