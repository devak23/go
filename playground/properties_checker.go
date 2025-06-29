package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//
//func PropertiesComparator() {
//	root := "/path/to/your/properties/files"
//
//}

func ParserProperties(file string) (map[string]string, error) {
	props := make(map[string]string)
	f, err := os.OpenFile(file, os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %w", file, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "!") {
			continue // Skip empty lines and comments
		}
		// Split the line into key and value
		var key, value string
		n, err := fmt.Sscanf(line, "%[^=]=%s", &key, &value)
		if n == 2 && err == nil {
			props[key] = value
		} else {
			return nil, fmt.Errorf("error parsing line '%s' in file %s: %w", line, file, err)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", file, err)
	}
	return props, nil
}

func ComparePropertiesFile(file1, file2 string) {
	props1, err1 := ParserProperties(file1)
	props2, err2 := ParserProperties(file2)

	if err1 != nil || err2 != nil {
		fmt.Println("Error parsing properties files:", err1, err2)
	}

	printUniqueKeysFoundIn(file1, file2, props1, props2)

	printUniqueKeysFoundIn(file2, file1, props2, props1)

	printKeysWithDifferentValues(file1, file2, props1, props2)
}

func printKeysWithDifferentValues(file1 string, file2 string, props1 map[string]string, props2 map[string]string) {
	for key, val1 := range props1 {
		if val2, exist := props2[key]; exist && val1 != val2 {
			fmt.Printf("Key '%s' has different values: '%s' in %s and '%s' in %s.\n", key, val1, file1, val2, file2)
		}
	}
	for key, val2 := range props2 {
		if val1, exist := props1[key]; exist && val1 != val2 {
			fmt.Printf("Key '%s' has different values: '%s' in %s and '%s' in %s.\n", key, val2, file2, val1, file1)
		}
	}
}

func printUniqueKeysFoundIn(sourceFile string, destFile string, source map[string]string, target map[string]string) {
	for key := range source {
		if _, exists := target[key]; !exists {
			fmt.Printf("Key '%s' found in %s but not in %s.\n", key, sourceFile, destFile)
		}
	}
}
