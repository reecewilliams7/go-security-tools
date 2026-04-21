package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	internaljwt "github.com/reecewilliams7/go-security-tools/internal/jwt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	decodeJWTCmd.Flags().StringP(JWTTokenFlag, "t", "", "The JWT to decode. Reads from stdin when not provided.")
	jwtCmd.AddCommand(decodeJWTCmd)
}

var decodeJWTCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decodes a JWT and displays its header and payload claims",
	Long:  "Decodes a JWT (without signature verification) and pretty-prints its header and payload.",
	PreRunE: func(cmd *cobra.Command, _ []string) error {
		return viper.BindPFlag(JWTTokenFlag, cmd.Flags().Lookup(JWTTokenFlag))
	},
	RunE: func(_ *cobra.Command, _ []string) error {
		token := viper.GetString(JWTTokenFlag)
		if token == "" {
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				token = strings.TrimSpace(scanner.Text())
			}
		}
		if token == "" {
			return fmt.Errorf("no JWT provided: use --token or pipe via stdin")
		}

		out, err := internaljwt.Decode(token)
		if err != nil {
			return err
		}

		headerJSON, _ := json.MarshalIndent(out.Header, "", "  ")
		payloadJSON, _ := json.MarshalIndent(out.Payload, "", "  ")

		fmt.Println("── Header ──────────────────────────────────────────────────")
		fmt.Println(string(headerJSON))
		fmt.Println("\n── Payload ─────────────────────────────────────────────────")
		fmt.Println(string(payloadJSON))

		if out.ExpiresAt != nil {
			fmt.Printf("\nExpires: %s", out.ExpiresAt.Format(time.RFC3339))
			if out.IsExpired {
				fmt.Print("  [EXPIRED]")
			}
			fmt.Println()
		}

		return nil
	},
}
