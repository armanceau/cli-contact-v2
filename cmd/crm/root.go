package crm

import (
	"fmt"
	"os"

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
		database.ConnectDB()
		store := storage.NewGormStore()

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
