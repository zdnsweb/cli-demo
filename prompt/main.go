package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/danielgtaylor/openapi-cli-generator/cli"
	demo "github.com/zdnsweb/cli-demo/cmd"
)

func completer(d prompt.Document) []prompt.Suggest {
	cli.Init(&cli.Config{
		AppName:   "zddi",
		EnvPrefix: "ZDDI",
		Version:   "1.0.0",
	})
	root := cli.Root
	demo.Register()

	s := []prompt.Suggest{}
	for _, c := range root.Commands() {
		// fmt.Printf("%+v\n", c)
		s = append(s, prompt.Suggest{Text: c.Name(), Description: c.Short})
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	fmt.Println("Please select table.")
	t := prompt.Input("> ", completer)
	fmt.Println("You selected " + t)
}
