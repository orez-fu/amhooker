package cli

import (
	"amhooker/amhooker"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "amhooker",
	Short: "cli to start example server & client",
	Long:  "cli to start example server & client",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		readGlobalConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := runRootJob()
		if err != nil {
			log.Println(err.Error())
		}
	},
}

func runRootJob() (err error) {
	GlobalConfig.Print()

	amhookerApp := amhooker.NewAMHookerApp(
		GlobalConfig.Port,
		GlobalConfig.DebugMode,
		GlobalConfig.AlertConfigPath,
	)

	return amhookerApp.Start()
}
