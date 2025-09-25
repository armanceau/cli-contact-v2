package crm

import (
	"fmt"
	"log"
	"os"

	"github.com/armanceau/cli-contact-v2/config"
	"github.com/armanceau/cli-contact-v2/internal/app"
	"github.com/armanceau/cli-contact-v2/internal/database"
	"github.com/armanceau/cli-contact-v2/internal/storage"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-contact",
	Short: "cli-contact - outil de stockage de contact",
	Long:  "cli-contact est un outil CLI qui permet le stockage des contacts en les ajoutants, supprimants, modifants ...",
	Run: func(cmd *cobra.Command, args []string) {
		config.InitConfig()

		var store storage.Storer
		switch config.Config.Storage.Type {
		case "gorm":
			database.ConnectDB(config.Config.Storage.SqliteFile)
			store = storage.NewGormStore()
		case "json":
			store = storage.NewJsonStore(config.Config.Storage.JsonFile)
		case "memory":
			store = storage.NewMemoryStore()
		default:
			log.Fatalf("Backend inconnu : %s", config.Config.Storage.Type)
		}

		fmt.Printf("Storage utilisé : %s\n", config.Config.Storage.Type)

		app.Run(store)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {}
