package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/zkhrg/go_day01/pkg/dbreader"
)

func main() {
	// // Пример обработки JSON
	// jsonFile, err := os.Open("../datasets/recipes001.json")
	// if err != nil {
	// 	fmt.Println("Error opening JSON file:", err)
	// 	return
	// }
	// defer jsonFile.Close()
	// decoder := json.NewDecoder(jsonFile)

	// _, err = decoder.Token()
	// if err != nil {
	// 	fmt.Println("Error reading '{' token:", err)
	// 	return
	// }

	// // Пропускаем ключ "cake"
	// _, err = decoder.Token()
	// if err != nil {
	// 	fmt.Println("Error reading 'cake' token:", err)
	// 	return
	// }

	// // Пропускаем открывающую скобку массива
	// _, err = decoder.Token()
	// if err != nil {
	// 	fmt.Println("Error reading array token:", err)
	// 	return
	// }

	// // Читаем объекты по одному
	// for decoder.More() {
	// 	var cake dbreader.Cake
	// 	err := decoder.Decode(&cake)
	// 	if err != nil {
	// 		fmt.Println("Error decoding JSON-CAKE:", err)
	// 		return
	// 	}
	// 	// Обрабатываем каждый контакт
	// 	fmt.Printf("Name: %s, Time: %s\n", cake.Name, cake.Time)
	// }

	// // Пропускаем закрывающую скобку массива
	// _, err = decoder.Token()
	// if err != nil {
	// 	fmt.Println("Error reading '}' token:", err)
	// }

	xmlFile, err := os.Open("/Users/diamondp/Projects/Go_Day01-1/datasets/recipes001.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}

	defer xmlFile.Close()

	var dbRecipe dbreader.Recipe
	xmlDecoder := xml.NewDecoder(xmlFile)

	if err := xmlDecoder.Decode(&dbRecipe); err != nil {
		fmt.Println("Error decoding XML:", err)
		return
	}
	for _, cake := range dbRecipe.Cakes {
		fmt.Printf("XML | Cake: %s | Time: %s\n", cake.Name, cake.Time)
		fmt.Printf("XML |     Ingredients: \n")
		for _, ingredient := range cake.Ingredients {
			fmt.Printf("XML |         Name: %s\n", ingredient.Name)
			fmt.Printf("XML |         Count: %s\n", ingredient.Count)
			fmt.Printf("XML |         Unit: %s\n", ingredient.Unit)
		}
	}
	// fmt.Printf("JSON - ID: %d, Name: %s\n", jsonPerson.ID, jsonPerson.Name)
	// for _, contact := range jsonPerson.Contacts {
	// 	fmt.Printf("  Contact Type: %s, Value: %s\n", contact.Type, contact.Value)
	// }

	// // Пример обработки XML
	// xmlFile, err := os.Open("data.xml")
	// if err != nil {
	// 	fmt.Println("Error opening XML file:", err)
	// 	return
	// }
	// defer xmlFile.Close()

	// var xmlPerson Person
	// xmlDecoder := xml.NewDecoder(xmlFile)
	// if err := xmlDecoder.Decode(&xmlPerson); err != nil {
	// 	fmt.Println("Error decoding XML:", err)
	// 	return
	// }
	// fmt.Printf("XML - ID: %d, Name: %s\n", xmlPerson.ID, xmlPerson.Name)
	// for _, contact := range xmlPerson.Contacts {
	// 	fmt.Printf("  Contact Type: %s, Value: %s\n", contact.Type, contact.Value)
	// }
}

// func main() {
// 	// Открываем файл
// 	file, err := os.Open("cakes.xml")
// 	if err != nil {
// 			fmt.Println("Error opening file:", err)
// 			return
// 	}
// 	defer file.Close()

// 	// Создаем декодер
// 	decoder := xml.NewDecoder(file)

// 	// Обрабатываем каждый элемент <cake>
// 	for {
// 			token, err := decoder.Token()
// 			if err != nil {
// 					if err.Error() == "EOF" {
// 							break
// 					}
// 					fmt.Println("Error reading token:", err)
// 					return
// 			}

// 			switch token := token.(type) {
// 			case xml.StartElement:
// 					if token.Name.Local == "cake" {
// 							var cake Cake
// 							if err := decoder.DecodeElement(&cake, &token); err != nil {
// 									fmt.Println("Error decoding XML:", err)
// 									return
// 							}
// 							// Обрабатываем каждый торт
// 							fmt.Printf("Cake Name: %s\n", cake.Name)
// 							fmt.Printf("Bake Time: %s\n", cake.Time)
// 							for _, ing := range cake.Ingredients {
// 									fmt.Printf("  Ingredient: %s, Count: %s, Unit: %s\n", ing.IngredientName, ing.IngredientCount, ing.IngredientUnit)
// 							}
// 					}
// 			}
// 	}
// }
