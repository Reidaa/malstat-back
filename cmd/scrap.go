package cmd

import (
	"malstat/scrapper/internal"
	"malstat/scrapper/pkg/utils"

	"github.com/urfave/cli"
)

// ScrapCmd processes the scraping command by managing the database connection,
// CSV file output, and scraping operation. It takes a CLI context containing:
//   - db: database connection string
//   - csv: output CSV file path (optional)
//   - top: number of records to process
//
// The function executes the scraping operation through internal.Scrap and
// returns any encountered errors during the process.
func ScrapCmd(ctx *cli.Context) error {
	var connStr string = ctx.String("db")
	var csvFile string = ctx.String("csv")
	var top int = ctx.Int("top")

	if csvFile != "" {
		utils.Info.Println("Output to", csvFile)
	}

	err := internal.Scrap(top, connStr, csvFile)
	if err != nil {
		return err
	}

	return nil
}
