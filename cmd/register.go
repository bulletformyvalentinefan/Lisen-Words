package cmd

import (
	"fmt"
	"lisen-words/pkg/db"
	"github.com/spf13/cobra"
)

var (
	regUsername string
	regPassword string
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Registrar un nuevo usuario",
	Run: func(cmd *cobra.Command, args []string) {
		database, err := db.InitDB("./lisen.db")
		if err != nil {
			fmt.Println("Error al conectar con la base de datos:", err)
			return
		}
		defer database.Close()

		err = db.CreateUser(database, regUsername, regPassword)
		if err != nil {
			fmt.Println("Error al registrar usuario:", err)
			return
		}

		fmt.Printf("¡Usuario '%s' registrado exitosamente!\n", regUsername)
	},
}

func init() {
	registerCmd.Flags().StringVarP(&regUsername, "user", "u", "", "Nombre de usuario (requerido)")
	registerCmd.Flags().StringVarP(&regPassword, "pass", "p", "", "Contraseña del usuario (requerido)")
	registerCmd.MarkFlagRequired("user")
	registerCmd.MarkFlagRequired("pass")

	rootCmd.AddCommand(registerCmd)
}