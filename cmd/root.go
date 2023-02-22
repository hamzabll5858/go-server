/*
Copyright © 2023 HAMZA BILAL <hamza.cs@outlook.com>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	grpc_server "go-server/pkg/grpc"
	"go-server/pkg/routers"
	"go-server/pkg/smtp"
	"os"
	"strings"
	"sync"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-server",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		grpc_enabled, _ := cmd.Flags().GetBool("grpc-server")
		wg.Add(1)
		if grpc_enabled {
			go grpc_server.InitGRPC()
		}
		http_enabled, _ := cmd.Flags().GetBool("http-server")
		if http_enabled {
			go routers.InitRouter(&wg)
		}
		send_mail, _ := cmd.Flags().GetBool("send-mail")
		if send_mail {
			smtp.SMTPSend()
			wg.Done()
		}
		wg.Wait()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolP("grpc-server", "G", false, "Run GRPC server")
	rootCmd.Flags().BoolP("http-server", "S", false, "Run HTTP server")
	rootCmd.Flags().BoolP("send-mail", "m", false, "send a test mail")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name "config" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath("./")
		viper.AddConfigPath("config/")
		viper.SetConfigType("env")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	viper.SetEnvKeyReplacer(strings.NewReplacer(`.`, `_`))
	viper.SetEnvPrefix("GOSERVER")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
