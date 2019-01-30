package cli

import (
	deviceWallet "github.com/skycoin/hardware-wallet-go/device-wallet"
	gcli "github.com/urfave/cli"
)

func emulatorGenerateMnemonicCmd() gcli.Command {
	name := "emulatorGenerateMnemonic"
	return gcli.Command{
		Name:        name,
		Usage:       "Ask the device to generate a mnemonic and configure itself with it.",
		Description: "",
		Flags: []gcli.Flag{
			gcli.BoolFlag{
				Name:  "usePassphrase",
				Usage: "Configure a passphrase",
			},
			gcli.IntFlag{
				Name:  "wordCount",
				Usage: "Use a specific (12 | 24) number of words (default 12) for the Mnemonic",
				Value: 12,
			},
		},
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) {
			passphrase := c.Bool("usePassphrase")
			wordCount := uint32(c.Uint64("wordCount"))
			deviceWallet.DeviceGenerateMnemonic(deviceWallet.DeviceTypeEmulator, wordCount, passphrase)
		},
	}
}
