package cmd

import (
	"github.com/fzxiehui/todo_serve/config"
	"github.com/fzxiehui/todo_serve/log"
	"github.com/spf13/cobra"
)

var configPath string

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a TODO server",
	Long:  `Start a TODO server`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Debug("start called")
		log.Debug("configPath:", configPath)
		if configPath != "" {
			err := config.ReadViperConfigFromFile(configPath)
			if err != nil {
				log.Fatal(err)
				return
			}
		}
		cfg := config.Config()
		log.Debug("loglevel:", cfg.GetString("loglevel"))
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&configPath, "config", "c", "", "config file path")
}
