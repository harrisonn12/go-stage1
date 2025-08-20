package main

import "fmt"

func main() {
	dict := map[string]string{
		"hello": "hola",
		"world": "mundo",
		"cat":   "gato",
		"dog":   "perro",
		"food":  "comida",
	}

	word := "world"
	if v, ok := dict[word]; ok {
		fmt.Printf("%q -> %q\n", word, v)
	} else {
		fmt.Println("Word not found")
	}

	delete(dict, "cat")

	for k, v := range dict {
		fmt.Printf("%s => %s\n", k, v)
	}
}
