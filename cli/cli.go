package cli

import (
	"log"

	"github.com/leebenson/conform"
)

func init() {
	conform.AddSanitizer("redact", func(_ string) string { return "*****"})
	if err := configInit(); err != nil {
		log.Panic(err)
	}
}

func Execute() error {
	return rootCmd.Execute()
}