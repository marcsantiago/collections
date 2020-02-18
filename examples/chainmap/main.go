package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/marcsantiago/collections"
	chain "github.com/marcsantiago/collections/chainmap"
)

// Example of letting user specified command-line arguments take precedence over environment variables which in turn take precedence over default values:
func main() {
	defaults := collections.GenericMap{
		collections.StringValue("color"): collections.StringValue("red"),
		collections.StringValue("user"):  collections.StringValue("guest"),
	}
	combined := chain.New(args(), osENV(), defaults)
	color, ok := combined.Get(collections.StringValue("color"))
	if !ok {
		log.Fatal("missing color")
	}

	user, ok := combined.Get(collections.StringValue("user"))
	if !ok {
		log.Fatal("missing user")
	}

	fmt.Printf("color %s, user %s", color, user)
}

func osENV() collections.GenericMap {
	// try un commenting me
	//os.Setenv("COLOR", "gold")
	//os.Setenv("USER", "admin")

	color := os.Getenv("COLOR")
	user := os.Getenv("EXAMPLE-USER")
	if color != "" && user != "" {
		return collections.GenericMap{
			collections.StringValue("color"): collections.StringValue(color),
			collections.StringValue("user"):  collections.StringValue(user),
		}
	}

	return collections.GenericMap{}
}

func args() collections.GenericMap {
	// try passing in flags
	color := flag.String("color", "", "")
	user := flag.String("user", "", "")
	flag.Parse()

	if *color != "" && *user != "" {
		return collections.GenericMap{
			collections.StringValue("color"): collections.StringValue(*color),
			collections.StringValue("user"):  collections.StringValue(*user),
		}
	}

	return collections.GenericMap{}
}
