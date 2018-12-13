package cli

import (
	"fmt"

	gcli "github.com/urfave/cli"

	deviceWallet "github.com/skycoin/hardware-wallet-go/device-wallet"
	"github.com/skycoin/hardware-wallet-go/device-wallet/messages"
)

func emulatorRecoveryCmd() gcli.Command {
	name := "emulatorRecovery"
	return gcli.Command{
		Name:        name,
		Usage:       "Ask the emulator to perform the seed recovery procedure.",
		Description: "",

		Flags: []gcli.Flag{
			gcli.BoolFlag{
				Name:  "usePassphrase",
				Usage: "Configure a passphrase",
			},
		},
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) {
			passphrase := c.Bool("usePassphrase")
			msg := deviceWallet.RecoveryDevice(deviceWallet.DeviceTypeEmulator, passphrase)
			for msg.Kind == uint16(messages.MessageType_MessageType_WordRequest) {
				var word string
				fmt.Printf("Word: ")
				fmt.Scanln(&word)
				msg = deviceWallet.DeviceWordAck(deviceWallet.DeviceTypeEmulator, word)
			}

			if msg.Kind == uint16(messages.MessageType_MessageType_ButtonRequest) {
				// Send ButtonAck
				msg = deviceWallet.DeviceButtonAck(deviceWallet.DeviceTypeEmulator)
			}

			if msg.Kind == uint16(messages.MessageType_MessageType_Failure) {
				failMsg := deviceWallet.DecodeFailMsg(msg.Kind, msg.Data)
				fmt.Println("Failed with code: ", failMsg)
				return
			}

			if msg.Kind == uint16(messages.MessageType_MessageType_Success) {
				successMsg := deviceWallet.DecodeSuccessMsg(msg.Kind, msg.Data)
				fmt.Println("Success with code: ", successMsg)
				return
			}
		},
	}
}
