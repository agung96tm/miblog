package migrate

import (
	"github.com/agung96tm/miblog/lib"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

var configFile string
var executeAs string
var filename string

func init() {
	pf := StartCmd.PersistentFlags()
	pf.StringVarP(&configFile, "config", "c", "config/config.yaml", "parameter used to start service")
	pf.StringVarP(&executeAs, "executeAs", "e", "up", "execute name, support up, down, and redo")
	pf.StringVarP(&filename, "filename", "f", "", "migration file name")

	_ = cobra.MarkFlagRequired(pf, "executeAs")
}

var StartCmd = &cobra.Command{
	Use:          "migrate",
	Short:        "Migrate database",
	SilenceUsage: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		lib.SetConfigPath(configFile)
	},
	Run: func(cmd *cobra.Command, args []string) {
		config := lib.NewConfig()
		database := lib.NewDatabase(config)
		migrate := lib.NewMigration(config)

		if err := migrate.Migrate(executeAs, filename, database); err != nil {
			log.Fatalf("Error to migrate: %v", err)
		}
	},
}
