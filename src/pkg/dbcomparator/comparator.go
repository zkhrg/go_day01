package dbcomparator

import (
	"fmt"

	"github.com/zkhrg/go_day01/pkg/dbreader"
)

func CompareDB(old_db_filename, new_db_filename string) error {
	var old_file_reader dbreader.DBReader
	var new_file_reader dbreader.DBReader
	var err error

	old_file_reader, err = dbreader.ReaderByFileExtension(old_db_filename)
	if err != nil {
		return err
	}
	new_file_reader, err = dbreader.ReaderByFileExtension(new_db_filename)
	if err != nil {
		return err
	}
	old_file_reader.ReadFile(old_db_filename)
	new_file_reader.ReadFile(new_db_filename)

	CompareDBPrintUtil(old_file_reader, new_file_reader)

	return err
}

func CompareDBPrintUtil(old_file_reader, new_file_reader dbreader.DBReader) {
	old_file_map := dbreader.OriginalRecipeToMapRecipe(old_file_reader.GetRecipe())
	new_file_map := dbreader.OriginalRecipeToMapRecipe(new_file_reader.GetRecipe())

	for k, v := range old_file_map.Cakes {
		cake, ok := new_file_map.Cakes[k]
		if !ok {
			fmt.Printf("REMOVED cake \"%s\"\n", k)
			continue
		}
		if v.Time != cake.Time {
			fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" isntead of \"%s\"\n", k, cake.Time, v.Time)
		}
		for k1, v1 := range v.Ingredients {
			ingredient, ok := new_file_map.Cakes[k].Ingredients[k1]
			if !ok {
				fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", k1, k)
				continue
			}
			if v1.Count != ingredient.Count {
				fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", k1, k, ingredient.Count, v1.Count)
			}
			if v1.Unit != ingredient.Unit && ingredient.Unit != "" && v1.Unit != "" {
				fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", k1, k, ingredient.Unit, v1.Unit)
			} else if v1.Unit != ingredient.Unit && ingredient.Unit == "" {
				fmt.Printf("REMOVED unit for ingredient \"%s\" for cake \"%s\"\n", k1, k)
			} else if v1.Unit != ingredient.Unit && ingredient.Unit != "" && v1.Unit == "" {
				fmt.Printf("ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", ingredient.Unit, k1, k)
			}
		}
		for k2 := range cake.Ingredients {
			_, ok := old_file_map.Cakes[k].Ingredients[k2]
			if !ok {
				fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", k2, k)
			}
		}
	}
	for k := range new_file_map.Cakes {
		_, ok := old_file_map.Cakes[k]
		if !ok {
			fmt.Printf("ADDED cake \"%s\"\n", k)
		}
	}
}
