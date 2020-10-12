package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/bianjieai/bsnhub-service-demo/common"
	"github.com/bianjieai/bsnhub-service-demo/examples/nft-service-provider/iservice"
)

var (
	KeysCmd = &cobra.Command{
		Use:   "keys",
		Short: "iService Key management commands",
	}
)

// KeysAddCmd implements the keys add command
func KeysAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name] [passphrase] [config-file]",
		Short: "Generate a new key",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := ""

			if len(args) == 2 {
				configFileName = common.DefaultConfigFileName
			} else {
				configFileName = args[2]
			}

			config, err := common.LoadYAMLConfig(configFileName)
			if err != nil {
				return err
			}

			iserviceClient := iservice.MakeServiceClientWrapper(iservice.NewConfig(config))

			addr, mnemonic, err := iserviceClient.AddKey(args[0], args[1])
			if err != nil {
				return err
			}

			fmt.Printf("key generated successfully: \n\nname: %s\naddress: %s\nmnemonic: %s\n\n", args[0], addr, mnemonic)

			return nil
		},
	}

	return cmd
}

// KeysShowCmd implements the keys show command
func KeysShowCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [name] [passphrase] [config-file]",
		Short: "Show the key information by name",
		Args:  cobra.RangeArgs(2, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := ""

			if len(args) == 2 {
				configFileName = common.DefaultConfigFileName
			} else {
				configFileName = args[2]
			}

			config, err := common.LoadYAMLConfig(configFileName)
			if err != nil {
				return err
			}

			iserviceClient := iservice.MakeServiceClientWrapper(iservice.NewConfig(config))

			addr, err := iserviceClient.ShowKey(args[0], args[1])
			if err != nil {
				return err
			}

			fmt.Printf("%s\n", addr)

			return nil
		},
	}

	return cmd
}

// KeysImportCmd implements the keys import command
func KeysImportCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import [name] [passphrase] [key-armor] [config-file]",
		Short: "Import a key from the private key armor",
		Args:  cobra.RangeArgs(3, 4),
		RunE: func(cmd *cobra.Command, args []string) error {
			configFileName := ""

			if len(args) == 3 {
				configFileName = common.DefaultConfigFileName
			} else {
				configFileName = args[2]
			}

			config, err := common.LoadYAMLConfig(configFileName)
			if err != nil {
				return err
			}

			iserviceClient := iservice.MakeServiceClientWrapper(iservice.NewConfig(config))

			addr, err := iserviceClient.ImportKey(args[0], args[1], args[2])
			if err != nil {
				return err
			}

			fmt.Printf("key imported successfully: %s\n", addr)

			return nil
		},
	}

	return cmd
}

func init() {
	KeysCmd.AddCommand(
		KeysAddCmd(),
		KeysShowCmd(),
		KeysImportCmd(),
	)
}
