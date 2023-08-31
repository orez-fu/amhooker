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
	// GlobalConfig.Print()
	if GlobalConfig.AlertConfigPath == "" {
		log.Fatal("alert configuration file not found. Plase passing amhooker configuration file from env \"AMHOOKER_CONFIG_FILE=<file_path>\" or command option \"--config_file=<file_path>\"")
	}

	amhookerApp := amhooker.NewAMHookerApp(
		GlobalConfig.Port,
		GlobalConfig.DebugMode,
		GlobalConfig.AlertConfigPath,
	)

	return amhookerApp.Start()
}
