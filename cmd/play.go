package cmd

import (
	"fmt"
	"lisen-words/pkg/audio"
	"lisen-words/pkg/deezer"
	"github.com/spf13/cobra"
)

var playCmd = &cobra.Command{
	Use:   "play [query]",
	Short: "Buscar y reproducir una canción en tiempo real",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]

		res, err := deezer.SearchTrack(query)
		if err != nil || len(res.Data) == 0 {
			fmt.Println("No se encontró la canción.")
			return
		}

		track := res.Data[0]
		fmt.Printf("\n▶ Reproduciendo: %s - %s\n", track.Title, track.Artist.Name)
		fmt.Println("Presiona Ctrl+C para detener la reproducción.")

		err = audio.PlayUrl(track.Preview)
		if err != nil {
			fmt.Println("Error al reproducir audio:", err)
			return
		}

		fmt.Println("Reproducción terminada.")
	},
}

func init() {
	rootCmd.AddCommand(playCmd)
}