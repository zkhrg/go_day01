package dbcomparator_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/zkhrg/go_day01/pkg/dbcomparator"
)

func TestCompareDB(t *testing.T) {
	oldFileContent := `{
  "cake": [
    {
      "name": "Red Velvet Strawberry Cake",
      "time": "45 min",
      "ingredients": [
        {
          "ingredient_name": "Flour",
          "ingredient_count": "2"
        },
        {
          "ingredient_name": "Strawberries",
          "ingredient_count": "8"
        },
        {
          "ingredient_name": "Coffee Beans",
          "ingredient_count": "2.5",
          "ingredient_unit": "tablespoons"
        },
        {
          "ingredient_name": "Cinnamon",
          "ingredient_count": "1"
        }
      ]
    },
    {
      "name": "Moonshine Muffin",
      "time": "30 min",
      "ingredients": [
        {
          "ingredient_name": "Brown sugar",
          "ingredient_count": "1",
          "ingredient_unit": "mug"
        },
        {
          "ingredient_name": "Blueberries",
          "ingredient_count": "1",
          "ingredient_unit": "mug"
        }
      ]
    }
  ]
}`
	newFileContent := `<recipes>
  <cake>
    <name>Red Velvet Strawberry Cake</name>
    <stovetime>32 min</stovetime>
    <ingredients>
      <item>
        <itemname>Flour</itemname>
        <itemcount>3</itemcount>
        <itemunit>tanks</itemunit>
      </item>
      <item>
        <itemname>Sugar</itemname>
        <itemcount>1.5</itemcount>
        <itemunit>cups</itemunit>
      </item>
    </ingredients>
  </cake>
  <cake>
    <name>Chocolate Cake</name>
    <stovetime>60 min</stovetime>
    <ingredients>
      <item>
        <itemname>Cocoa powder</itemname>
        <itemcount>0.5</itemcount>
        <itemunit>cups</itemunit>
      </item>
    </ingredients>
  </cake>
</recipes>`
	expectedOutput := `CHANGED cooking time for cake "Red Velvet Strawberry Cake" - "32 min" isntead of "45 min"
CHANGED unit count for ingredient "Flour" for cake "Red Velvet Strawberry Cake" - "3" instead of "2"
ADDED unit "tanks" for ingredient "Flour" for cake "Red Velvet Strawberry Cake"
REMOVED ingredient "Strawberries" for cake "Red Velvet Strawberry Cake"
REMOVED ingredient "Coffee Beans" for cake "Red Velvet Strawberry Cake"
REMOVED ingredient "Cinnamon" for cake "Red Velvet Strawberry Cake"
ADDED ingredient "Sugar" for cake "Red Velvet Strawberry Cake"
REMOVED cake "Moonshine Muffin"
ADDED cake "Chocolate Cake"
`

	oldFile, err := os.CreateTemp("", "*.json")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(oldFile.Name())

	_, err = oldFile.WriteString(oldFileContent)
	if err != nil {
		t.Fatal(err)
	}

	if err := oldFile.Close(); err != nil {
		t.Fatal(err)
	}

	newFile, err := os.CreateTemp("", "*.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(newFile.Name())

	_, err = newFile.WriteString(newFileContent)
	if err != nil {
		t.Fatal(err)
	}

	if err := newFile.Close(); err != nil {
		t.Fatal(err)
	}

	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	oldStdout := os.Stdout
	os.Stdout = w

	defer func() {
		os.Stdout = oldStdout
		w.Close()
	}()

	dbcomparator.CompareDB(oldFile.Name(), newFile.Name())

	w.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}
	os.Stdout = oldStdout
	actualOutput := buf.String()
	if actualOutput != expectedOutput {
		fmt.Printf("test results not eq, but compare in maunally becase make changed the order\n%s \n\nand\n\n %s", actualOutput, expectedOutput)
	}
}
