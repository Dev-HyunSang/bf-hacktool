package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "brute force hacktool",
	Short: "brute force hacktool by @dev-hyunsang",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringP("author", "a", "HyunSang Park")
	viper.SetDefault("author", "HyunSang Park <me@hyunsang.dev>")
	viper.SetDefault("license", "MIT")
}

func initConfig() {

}
