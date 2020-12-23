package advent

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func parseIngredientsAndAllergens(filename string) (ingredients, allergens [][]string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "(")
		ingredients = append(ingredients, strings.Split(line[0][:len(line[0])-1], " "))
		allergens = append(allergens, strings.Split(line[1][9:len(line[1])-1], ", "))
	}
	return
}

func possibleAllergens(ingredients, allergens [][]string) (allergenMap map[string]map[string]bool) {
	allergenMap = make(map[string]map[string]bool)
	for ii, _allergens := range allergens {
		for _, allergen := range _allergens {
			mapping, ok := allergenMap[allergen]
			if !ok {
				mapping = make(map[string]bool)
				for _, ingredient := range ingredients[ii] {
					mapping[ingredient] = true
				}
				allergenMap[allergen] = mapping
			} else {
				for ingredient := range mapping {
					if !containsString(ingredients[ii], ingredient) {
						delete(mapping, ingredient)
					}
				}
			}
		}
	}
	return allergenMap
}

func splitIngredients(ingredients, allergens [][]string) (goodIngredients, badIngredients []string) {
	allergenMap := possibleAllergens(ingredients, allergens)
	var allIngredients []string
	for _, _ingredients := range ingredients {
		for _, ingredient := range _ingredients {
			if !containsString(allIngredients, ingredient) {
				allIngredients = append(allIngredients, ingredient)
			}
		}
	}
	for _, ingredient := range allIngredients {
		good := true
		for _, mapping := range allergenMap {
			if mapping[ingredient] {
				good = false
				break
			}
		}
		if good {
			goodIngredients = append(goodIngredients, ingredient)
		} else {
			badIngredients = append(badIngredients, ingredient)
		}
	}
	return
}

func countNonAllergens(filename string) int {
	ingredients, allergens := parseIngredientsAndAllergens(filename)
	goodIngredients, _ := splitIngredients(ingredients, allergens)

	output := 0
	for _, ingredient := range goodIngredients {
		for _, _ingredients := range ingredients {
			if containsString(_ingredients, ingredient) {
				output += 1
			}
		}
	}
	return output
}

func allergenNames(allergenMap map[string]map[string]bool) (allergenNames []string) {
	for allergen := range allergenMap {
		allergenNames = append(allergenNames, allergen)
	}
	sort.Strings(allergenNames)
	return
}

func valuesFromSet(mapping map[string]bool) (values []string) {
	for key := range mapping {
		values = append(values, key)
	}
	return
}

func findAllergens(filename string) string {
	ingredients, allergens := parseIngredientsAndAllergens(filename)
	allergenMap := possibleAllergens(ingredients, allergens)
	badIngredientMap := make(map[string]string)
	allergenNames := allergenNames(allergenMap)

	for len(allergenMap) > 0 {
		// identify entries with just one allowed ingredient
		for allergen, mapping := range allergenMap {
			values := valuesFromSet(mapping)
			if len(values) == 1 {
				badIngredientMap[allergen] = values[0]
			}
		}
		// removes identified allergen/ingredient from other
		// allergen lists
		for allergen, ingredient := range badIngredientMap {
			delete(allergenMap, allergen)
			for _, mapping := range allergenMap {
				delete(mapping, ingredient)
			}
		}
	}
	output := ""
	for _, allergen := range allergenNames {
		output += badIngredientMap[allergen] + ","
	}
	return output[:len(output)-1]
}

func Day21() {
	fmt.Println("Test")
	fmt.Println("Part 1: ", countNonAllergens("inputs/day21_test.txt"))
	fmt.Println("Part 2: ", findAllergens("inputs/day21_test.txt"))
	fmt.Println("Main")
	fmt.Println("Part 1: ", countNonAllergens("inputs/day21.txt"))
	fmt.Println("Part 2: ", findAllergens("inputs/day21.txt"))
}
