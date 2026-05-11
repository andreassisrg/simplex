package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/andreassisrg/simplex/internal/dto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootFlags = dto.RootCmdFlags{}

	rootCmd = &cobra.Command{
		Use:   "simplex",
		Short: "A simplex solver for linear programming problems",
		Args:  cobra.ExactArgs(0),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initializeConfig(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("it works")
			decimals := viper.GetUint("decimals")
			fmt.Println("decimals flags", decimals)

			return nil
		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&rootFlags.Filename, "file", "f", "", "file containg the linear programming")
	rootCmd.MarkFlagRequired("file")
	rootCmd.Flags().UintP("decimals", "d", 2, "number of digits after the decimal point for every printed number")
	rootCmd.Flags().UintVarP(&rootFlags.Digits, "digits", "i", 3, "total number of digits for every printed number")
	rootCmd.Flags().StringVarP(&rootFlags.Policy, "policy", "p", "largest", "should be either 'largest', 'bland' or 'smallest'")
	rootCmd.PersistentFlags().StringVarP(&rootFlags.ConfigFile, "config", "c", "", "optional config.yaml file")
}

func initializeConfig(cmd *cobra.Command) error {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "*", "-", "*"))
	viper.AutomaticEnv()

	if rootFlags.ConfigFile != "" {
		viper.SetConfigFile(rootFlags.ConfigFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		fmt.Println(home)
		viper.AddConfigPath(".")
		viper.AddConfigPath(home + "/.simplex")
		viper.SetConfigName("config")
		viper.SetConfigFile("yaml")
	}

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundErr viper.ConfigFileNotFoundError
		if errors.Is(err, &configFileNotFoundErr) {
			return err
		}
	}

	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		return err
	}

	fmt.Println("configuration initalized. using config file", viper.ConfigFileUsed())
	return nil
}
