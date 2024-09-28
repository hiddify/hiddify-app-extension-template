package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"

	ex "github.com/hiddify/hiddify-app-example-extension/hiddify_extension"
)

// // go:embed resources/en.i18n.json
// var enJsonData []byte
func getAllFilenames(efs *embed.FS) (files []string, err error) {
	if err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
}
func main() {
	var result map[string]interface{}

	// // Unmarshal the embedded JSON data
	// err := json.Unmarshal(enJsonData, &result)
	// if err != nil {
	// 	fmt.Println("Error loading JSON:", err)
	// 	return
	// }
	content1, _ := ex.Resources.ReadFile("translations/en.i18n.json")
	print(string(content1))

	v, e := getAllFilenames(&ex.Resources)
	fmt.Printf("%++v, %v", v, e)
	err := fs.WalkDir(ex.Resources, "translations/", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
		if err != nil {
			return err
		}

		// If it's a file, print its name and content
		if !d.IsDir() {
			fmt.Println("Found file:", path)
			content, err := ex.Resources.ReadFile(path)
			if err != nil {
				return err
			}
			fmt.Printf("Content of %s:\n%s\n", path, string(content))
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Loaded JSON:", result)
}
