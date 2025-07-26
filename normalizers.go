// Package textn8r provides a set of normalizers for strings.
// It includes functions to convert case, trim spaces, remove special characters,
// and more. Normalizers can be applied individually or as a collection.
// The package is designed to be flexible and extensible, allowing users to create
// custom normalizers as needed.
// // Example usage:
//
//	normalizers := textn8r.Normalizers{
//	    textn8r.UpperCaseNormalizer,
//	    textn8r.TrimSpaceNormalizer,
//	    textn8r.RemoveSpecialCharactersNormalizer,
//	}
//	result := normalizers.Apply("  Hello, World!  ")
//	fmt.Println(result) // Output: "HELLO WORLD"
//	customNormalizer := textn8r.Normalizer(func(input string) string {
//	    return strings.ReplaceAll(input, "foo", "bar")
//	})
//	result = customNormalizer.Apply("foo is replaced with bar")
//	fmt.Println(result) // Output: "bar is replaced with bar"
//	customNormalizer = textn8r.ReplaceTabNormalizer("-")
//	result = customNormalizer.Apply("This\tis\ta\ttest")
//	fmt.Println(result) // Output: "This-is-a-test"
//	customNormalizer = textn8r.ReplaceCarriageReturnNormalizer("-")
//	result = customNormalizer.Apply("This\ris\ra\rcarriage\rreturn")
//	fmt.Println(result) // Output: "This-is-a-carriage-return"
//	customNormalizer = textn8r.ReplaceNonAlphanumericNormalizer("-")
//	result = customNormalizer.Apply("This is a test! @2023")
//	fmt.Println(result) // Output: "This-is-a-test---2023"
//	customNormalizer = textn8r.ReplacePunctuationNormalizer("-")
//	result = customNormalizer.Apply("Hello, World! This is a test.")
//	fmt.Println(result) // Output: "Hello- World- This is a test-"
//	customNormalizer = textn8r.ReplaceDigitsNormalizer("-")
//	result = customNormalizer.Apply("Year 2023, Month 10, Day 15")
//	fmt.Println(result) // Output: "Year --, Month --, Day --"
//	customNormalizer = textn8r.ReplaceSpaceNormalizer("-")
//	result = customNormalizer.Apply("This is a test with spaces")
//	fmt.Println(result) // Output: "This-is-a-test-with-spaces"
//	customNormalizer = textn8r.ReplaceDiacriticsNormalizer("-")
//	result = customNormalizer.Apply("Café, résumé, naïve, jalapeño -")
//	fmt.Println(result) // Output: "Café- résumé- naïve- jalapeño -"
//	customNormalizer = textn8r.ReplaceNewLineNormalizer("-")
//	result = customNormalizer.Apply("Line 1\nLine 2\nLine 3")
//	fmt.Println(result) // Output: "Line 1-Line 2-Line 3"
//	customNormalizer = textn8r.ReplaceTildesNormalizer
//	result = customNormalizer.Apply("Cañón, año, niño")
//	fmt.Println(result) // Output: "Cañón, año, niño"
//	customNormalizer = textn8r.ReplaceAccentsNormalizer
//	result = customNormalizer.Apply("Café, résumé, naïve, jalapeño")
//	fmt.Println(result) // Output: "Cafe, resume, naive, jalapeno"
//	customNormalizer = textn8r.ReplaceSpecialCharactersNormalizer("-")
//	result = customNormalizer.Apply("Hello! @World #2023")
//	fmt.Println(result) // Output: "Hello- -World -2023"
//	customNormalizer = textn8r.ReplaceAllSpaceNormalizer
//	result = customNormalizer.Apply("This  is   a    test")
//	fmt.Println(result) // Output: "This-is-a-test"
package textn8r

import (
	"regexp"
	"strings"
)

// Normalizers is a collection of normalizers.
type Normalizers []Normalizer

// Apply applies all normalizers in the collection to the input string.
func (n Normalizers) Apply(input string) string {
	for _, normalizer := range n {
		input = normalizer(input)
	}

	return input
}

// Normalizer is a function that normalizes a string and not receive any parameter.
type Normalizer func(input string) string

// Apply applies the normalizer to the input string.
func (n Normalizer) Apply(input string) string {
	return n(input)
}

// UpperCaseNormalizer converts the input string to uppercase.
func UpperCaseNormalizer(input string) string {
	return strings.ToUpper(input)
}

// LowerCaseNormalizer converts the input string to lowercase.
func LowerCaseNormalizer(input string) string {
	return strings.ToLower(input)
}

// TrimSpaceNormalizer removes leading and trailing white spaces from the input string.
func TrimSpaceNormalizer(input string) string {
	return strings.TrimSpace(input)
}

// RemoveExtraSpaceNormalizer removes extra white spaces from the input string.
func RemoveExtraSpaceNormalizer(input string) string {
	return strings.Join(strings.Fields(input), " ")
}

// RemoveAllSpaceNormalizer removes all white spaces from the input string.
func RemoveAllSpaceNormalizer(input string) string {
	return strings.ReplaceAll(input, " ", "")
}

// RemoveCarriageReturnNormalizer removes carriage return characters from the input string.
func RemoveCarriageReturnNormalizer(input string) string {
	regex := regexp.MustCompile(`\r`)
	return regex.ReplaceAllString(input, "")
}

// RemoveNewLineNormalizer removes new line characters from the input string.
func RemoveNewLineNormalizer(input string) string {
	regex := regexp.MustCompile(`\n`)
	return regex.ReplaceAllString(input, "")
}

// RemoveTabNormalizer removes tab characters from the input string.
func RemoveTabNormalizer(input string) string {
	regex := regexp.MustCompile(`\t`)
	return regex.ReplaceAllString(input, "")
}

// RemoveNonAlphanumericNormalizer removes non-alphanumeric characters from the input string.
func RemoveNonAlphanumericNormalizer(input string) string {
	regex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	return regex.ReplaceAllString(input, "")
}

// RemoveTildesNormalizer removes tildes from the input string.
func RemoveTildesNormalizer(input string) string {
	regex := regexp.MustCompile(`[~]`)
	return regex.ReplaceAllString(input, "")
}

// RemoveDiacriticsNormalizer removes diacritics from the input string.
func RemoveDiacriticsNormalizer(input string) string {
	regex := regexp.MustCompile(`[^\x00-\x7F]+`)
	return regex.ReplaceAllString(input, "")
}

// RemoveSpecialCharactersNormalizer removes special characters from the input string.
func RemoveSpecialCharactersNormalizer(input string) string {
	regex := regexp.MustCompile(`[^a-zA-Z0-9\s\p{L}\p{N}]+`)
	return regex.ReplaceAllString(input, " ")
}

// RemovePunctuationNormalizer removes punctuation characters from the input string.
func RemovePunctuationNormalizer(input string) string {
	regex := regexp.MustCompile(`[[:punct:]]`)
	input = regex.ReplaceAllString(input, "")

	// not removed: ¿ ! ¡
	regex = regexp.MustCompile(`[¿!¡]`)
	input = regex.ReplaceAllString(input, "")

	return input
}

// RemoveDigitsNormalizer removes digit characters from the input string.
func RemoveDigitsNormalizer(input string) string {
	regex := regexp.MustCompile(`[0-9]`)
	return regex.ReplaceAllString(input, "")
}

// ReplaceSpecialCharactersNormalizer replaces special characters with a given replacement string.
func ReplaceSpecialCharactersNormalizer(input, replacement string) string {
	regex := regexp.MustCompile(`[^a-zA-Z0-9\s\p{L}\p{N}]+`)
	return regex.ReplaceAllString(input, replacement)
}

// ReplaceAccentsNormalizer replaces accented characters with their non-accented counterparts.
func ReplaceAccentsNormalizer(input string) string {
	// lowercase
	regex := regexp.MustCompile(`[áàãâä]`)
	input = regex.ReplaceAllString(input, "a")
	regex = regexp.MustCompile(`[éèêë]`)
	input = regex.ReplaceAllString(input, "e")
	regex = regexp.MustCompile(`[íìîï]`)
	input = regex.ReplaceAllString(input, "i")
	regex = regexp.MustCompile(`[óòõôö]`)
	input = regex.ReplaceAllString(input, "o")
	regex = regexp.MustCompile(`[úùûü]`)
	input = regex.ReplaceAllString(input, "u")
	regex = regexp.MustCompile(`[ç]`)
	input = regex.ReplaceAllString(input, "c")
	regex = regexp.MustCompile(`[ñ]`)
	input = regex.ReplaceAllString(input, "n")
	// uppercase
	regex = regexp.MustCompile(`[ÁÀÃÂÄ]`)
	input = regex.ReplaceAllString(input, "A")
	regex = regexp.MustCompile(`[ÉÈÊË]`)
	input = regex.ReplaceAllString(input, "E")
	regex = regexp.MustCompile(`[ÍÌÎÏ]`)
	input = regex.ReplaceAllString(input, "I")
	regex = regexp.MustCompile(`[ÓÒÕÔÖ]`)
	input = regex.ReplaceAllString(input, "O")
	regex = regexp.MustCompile(`[ÚÙÛÜ]`)
	input = regex.ReplaceAllString(input, "U")
	regex = regexp.MustCompile(`[Ç]`)
	input = regex.ReplaceAllString(input, "C")
	regex = regexp.MustCompile(`[Ñ]`)
	input = regex.ReplaceAllString(input, "N")

	return input
}

// ReplaceTildesNormalizer replaces tildes with their non-tilde counterparts.
func ReplaceTildesNormalizer(input string) string {
	regex := regexp.MustCompile(`[ñ]`)
	return regex.ReplaceAllString(input, "n")
}

// ReplaceTabNormalizer replaces tab characters with a given replacement string.
func ReplaceTabNormalizer(replacement string) Normalizer {
	return func(input string) string {
		regex := regexp.MustCompile(`\t`)
		return regex.ReplaceAllString(input, replacement)
	}
}

// ReplaceCarriageReturnNormalizer replaces carriage return characters with a given replacement string.
func ReplaceCarriageReturnNormalizer(replacement string) Normalizer {
	return func(input string) string {
		regex := regexp.MustCompile(`\r`)
		return regex.ReplaceAllString(input, replacement)
	}
}

// ReplaceNonAlphanumericNormalizer replaces non-alphanumeric characters with a given replacement string.
func ReplaceNonAlphanumericNormalizer(replacement string) Normalizer {
	return func(input string) string {
		regex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
		return regex.ReplaceAllString(input, replacement)
	}
}

// ReplacePunctuationNormalizer replaces punctuation characters with a given replacement string.
func ReplacePunctuationNormalizer(replacement string) Normalizer {
	return func(input string) string {
		regex := regexp.MustCompile(`[[:punct:]]`)
		return regex.ReplaceAllString(input, replacement)
	}
}

// ReplaceDigitsNormalizer replaces digit characters with a given replacement string.
func ReplaceDigitsNormalizer(replacement string) Normalizer {
	return func(input string) string {
		regex := regexp.MustCompile(`[0-9]`)
		return regex.ReplaceAllString(input, replacement)
	}
}

// ReplaceSpaceNormalizer replaces space characters with a given replacement string.
func ReplaceSpaceNormalizer(replacement string) Normalizer {
	return func(input string) string {
		regex := regexp.MustCompile(`\s`)
		return regex.ReplaceAllString(input, replacement)
	}
}

// ReplaceDiacriticsNormalizer replaces diacritics with a given replacement string.
func ReplaceDiacriticsNormalizer(replacement string) Normalizer {
	return func(input string) string {
		regex := regexp.MustCompile(`[^\x00-\x7F]+`)
		return regex.ReplaceAllString(input, replacement)
	}
}

// ReplaceNewLineNormalizer replaces new line characters with a given replacement string.
func ReplaceNewLineNormalizer(replacement string) Normalizer {
	return func(input string) string {
		regex := regexp.MustCompile(`\n`)
		return regex.ReplaceAllString(input, replacement)
	}
}
