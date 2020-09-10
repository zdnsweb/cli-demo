package main

import (
	"github.com/danielgtaylor/openapi-cli-generator/cli"
    "fmt"
    // "os"
    // "io/ioutil"
    "gopkg.in/h2non/gentleman.v2/context"
)

func main() {
	cli.Init(&cli.Config{
		AppName:   "zddi",
		EnvPrefix: "ZDDI",
		Version:   "1.0.0",
	})

	// TODO: Add register commands here.
    viewRegister(false)
    completionRegister()


    // Register a request middleware handler to print the path.
    cli.Client.UseRequest(func(ctx *context.Context, h context.Handler) {
	    fmt.Printf("Request path: %s\n", ctx.Request.URL.Path)
	    h.Next(ctx)
    })

    // Register a response middleware handler to print the status code.
    cli.Client.UseResponse(func(ctx *context.Context, h context.Handler) {
        fmt.Printf("Response status: %d\n", ctx.Response.StatusCode)
        // body, err := ioutil.ReadAll(ctx.Response.Body)
        // if err != nil {
        //     fmt.Fprintf(os.Stderr, "Errors: %s\n", err)
        // }
        // fmt.Printf("Response body: %s\n", body)
	    h.Next(ctx)
    })

	cli.Root.Execute()
}
