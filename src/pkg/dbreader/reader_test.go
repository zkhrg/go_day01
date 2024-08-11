package dbreader_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/zkhrg/go_day01/pkg/dbreader"
)

func TestReadFile(t *testing.T) {
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
	expectedOutput := `<Recipe>
    <cake>
        <name>Red Velvet Strawberry Cake</name>
        <stovetime>45 min</stovetime>
        <ingredients>
            <item>
                <itemname>Flour</itemname>
                <itemcount>2</itemcount>
                <itemunit></itemunit>
            </item>
            <item>
                <itemname>Strawberries</itemname>
                <itemcount>8</itemcount>
                <itemunit></itemunit>
            </item>
            <item>
                <itemname>Coffee Beans</itemname>
                <itemcount>2.5</itemcount>
                <itemunit>tablespoons</itemunit>
            </item>
            <item>
                <itemname>Cinnamon</itemname>
                <itemcount>1</itemcount>
                <itemunit></itemunit>
            </item>
        </ingredients>
    </cake>
    <cake>
        <name>Moonshine Muffin</name>
        <stovetime>30 min</stovetime>
        <ingredients>
            <item>
                <itemname>Brown sugar</itemname>
                <itemcount>1</itemcount>
                <itemunit>mug</itemunit>
            </item>
            <item>
                <itemname>Blueberries</itemname>
                <itemcount>1</itemcount>
                <itemunit>mug</itemunit>
            </item>
        </ingredients>
    </cake>
</Recipe>
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

	reader, _ := dbreader.ReaderByFileExtension(oldFile.Name())
	reader.ReadFile(oldFile.Name())
	reader.Print()

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

func TestReadFile1(t *testing.T) {
	oldFileContent := `<recipes>
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
</recipes>
`
	expectedOutput := `{
    "cake": [
        {
            "name": "Red Velvet Strawberry Cake",
            "time": "32 min",
            "ingredients": [
                {
                    "ingredient_name": "Flour",
                    "ingredient_count": "3",
                    "ingredient_unit": "tanks"
                },
                {
                    "ingredient_name": "Sugar",
                    "ingredient_count": "1.5",
                    "ingredient_unit": "cups"
                }
            ]
        },
        {
            "name": "Chocolate Cake",
            "time": "60 min",
            "ingredients": [
                {
                    "ingredient_name": "Cocoa powder",
                    "ingredient_count": "0.5",
                    "ingredient_unit": "cups"
                }
            ]
        }
    ]
}
`

	oldFile, err := os.CreateTemp("", "*.xml")
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

	reader, _ := dbreader.ReaderByFileExtension(oldFile.Name())
	reader.ReadFile(oldFile.Name())
	reader.Print()

	w.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}
	os.Stdout = oldStdout
	actualOutput := buf.String()
	if actualOutput != expectedOutput {
		fmt.Printf("test results not eq, but compare it maunally becase make changed the order\n%s \n\nand\n\n %s", actualOutput, expectedOutput)
	}
}
