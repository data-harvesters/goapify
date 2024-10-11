package internal

import (
	"log"

	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "new actorName",
		Short:   "Creates a new actor environment",
		Example: "goapify new airbnb-review-scraper",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, names []string) {
			environment := newEnv(names[0])

			err := environment.setup()
			if err != nil {
				log.Fatalln(err)
			}
		},
	}

	return cmd
}
