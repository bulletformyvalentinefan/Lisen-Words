package cmd

import (
	"fmt"
	"lisen-words/pkg/deezer"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Buscar canciones en Deezer",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]

		res, err := deezer.SearchTrack(query)
		if err != nil {
			fmt.Println("Error al consultar Deezer:", err)
			return
		}

		if len(res.Data) == 0 {
			fmt.Println("No se encontraron canciones.")
			return
		}

		for _, track := range res.Data {
			fmt.Printf("ID: %d | %s - %s (%ds)\n", track.ID, track.Title, track.Artist.Name, track.Duration)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}