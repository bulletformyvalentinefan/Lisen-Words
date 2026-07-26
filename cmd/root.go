package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lisen",
	Short: "Lisen Words - CLI para buscar, guardar y escuchar música",
	Long:  `Lisen Words es una aplicación de consola desarrollada en Go que te permite autenticarte, buscar canciones en Deezer y escucharlas directamente desde la terminal.`,
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}