package main

import (
	"log"

	"github.com/data-harvesters/goapify/cmd/goapify/internal"
	"github.com/spf13/cobra"
)

func main() {
	log.SetFlags(0)
	cmd := &cobra.Command{Use: "goapify"}
	cmd.AddCommand(
		internal.NewCmd(),
	)
	_ = cmd.Execute()
}
