package cmd

import (
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/hoop33/jwtbot/model"
	"github.com/spf13/cobra"
)

var (
	algo string
	sub  string
	key  string
)

// CreateCmd lists the supported algorithms
var CreateCmd = &cobra.Command{
	Use:   "create [flags]",
	Short: "create a JWT",
	Long:  "creates a JWT using the specified flags and key/value pairs",
	Run: func(cmd *cobra.Command, args []string) {

		algorithm, ok := model.AlgorithmForName(algo)
		if !ok {
			fmt.Printf("algorithm '%s' not valid\n", algo)
			os.Exit(1)
		}

		claims := &jwt.StandardClaims{
			Subject: sub,
		}

		var signedString string
		var err error
		token := jwt.NewWithClaims(algorithm, claims)

		if algorithm == jwt.SigningMethodNone {
			signedString, err = token.SignedString(jwt.UnsafeAllowNoneSignatureType)
		} else {
			signedString, err = token.SignedString([]byte(key))
		}
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println(signedString)
	},
}

func init() {
	CreateCmd.Flags().StringVarP(&algo, "algorithm", "a", "hs256", "algorithm for signing the JWT")
	CreateCmd.Flags().StringVarP(&key, "key", "k", "", "signing key")
	CreateCmd.Flags().StringVarP(&sub, "subject", "s", "", "subject for the JWT")

	RootCmd.AddCommand(CreateCmd)
}
