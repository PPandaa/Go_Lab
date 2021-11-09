package cobraLab

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "my-cmd",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Cobra CLI")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of cobrademo",
	Long:  `All software has versions. This is cobrademo's`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			switch strings.ToLower(args[0]) {
			case "python":
				fmt.Println("Python Version is v3.10")
			case "go":
				fmt.Println("Golang Version is v1.17.3")
			}
		} else {
			fmt.Println("CLI Version is v86.05.02")
		}
	},
}

var Source string
var lengthCmd = &cobra.Command{
	Use:   "length",
	Short: "Print the length of string",
	Long:  `All String has length. This is cobrademo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(len(Source))
	},
}

func init() {

	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(versionCmd)
	lengthCmd.Flags().StringVarP(&Source, "source", "s", "", "Source String")
	rootCmd.AddCommand(lengthCmd)

}

func initConfig() {

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".my-cmd")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
