package cli

import (
	"fmt"

	gcli "github.com/urfave/cli"

	deviceWallet "github.com/skycoin/hardware-wallet-go/device-wallet"
	"github.com/skycoin/hardware-wallet-go/device-wallet/messages"
)

func deviceRecoveryCmd() gcli.Command {
	name := "deviceRecovery"
	return gcli.Command{
		Name:         name,
		Usage:        "Ask the device to perform the seed recovery procedure.",
		Description:  "",
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) {
			msg := deviceWallet.RecoveryDevice(deviceWallet.DeviceTypeUsb)
			for msg.Kind == uint16(messages.MessageType_MessageType_WordRequest) {
				var word string
				fmt.Printf("Word: ")
				fmt.Scanln(&word)
				msg = deviceWallet.DeviceWordAck(deviceWallet.DeviceTypeUsb, word)
			}
			if msg.Kind == uint16(messages.MessageType_MessageType_Failure) {
				_, failMsg := deviceWallet.DecodeFailMsg(msg.Kind, msg.Data)
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
