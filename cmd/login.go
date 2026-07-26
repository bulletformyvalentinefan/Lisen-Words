package cmd

import (
	"fmt"
	"lisen-words/pkg/db"
	"github.com/spf13/cobra"
)

var (
	loginUsername string
	loginPassword string
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Iniciar sesión",
	Run: func(cmd *cobra.Command, args []string) {
		database, err := db.InitDB("./lisen.db")
		if err != nil {
			fmt.Println("Error al conectar con la base de datos:", err)
			return
		}
		defer database.Close()

		userID, err := db.AuthenticateUser(database, loginUsername, loginPassword)
		if err != nil {
			fmt.Println("Error de autenticación: Credenciales inválidas.")
			return
		}

		fmt.Printf("¡Bienvenido %s! Inicio de sesión exitoso (ID: %d).\n", loginUsername, userID)
	},
}

func init() {
	loginCmd.Flags().StringVarP(&loginUsername, "user", "u", "", "Nombre de usuario (requerido)")
	loginCmd.Flags().StringVarP(&loginPassword, "pass", "p", "", "Contraseña del usuario (requerido)")
	loginCmd.MarkFlagRequired("user")
	loginCmd.MarkFlagRequired("pass")

	rootCmd.AddCommand(loginCmd)
}