package cli

import (
	"fmt"

	gcli "github.com/urfave/cli"

	"log"

	deviceWallet "github.com/skycoin/hardware-wallet-go/device-wallet"
	"github.com/skycoin/hardware-wallet-go/device-wallet/messages"
)

func setPinCode() gcli.Command {
	name := "setPinCode"
	return gcli.Command{
		Name:        name,
		Usage:       "Configure a PIN code on a device.",
		Description: "",
		Flags: []gcli.Flag{
			gcli.StringFlag{
				Name:   "deviceType",
				Usage:  "Device type to send instructions to, hardware wallet (USB) or emulator.",
				EnvVar: "DEVICE_TYPE",
				Value:  "USB",
			},
		},
		OnUsageError: onCommandUsageError(name),
		Action: func(c *gcli.Context) {
			var deviceType deviceWallet.DeviceType
			switch c.String("deviceType") {
			case "USB":
				deviceType = deviceWallet.DeviceTypeUsb
			case "EMULATOR":
				deviceType = deviceWallet.DeviceTypeEmulator
			default:
				log.Println("No device detected")
				return
			}

			var pinEnc string
			kind, _ := deviceWallet.DeviceChangePin(deviceType)
			for kind == uint16(messages.MessageType_MessageType_PinMatrixRequest) {
				fmt.Printf("PinMatrixRequest response: ")
				fmt.Scanln(&pinEnc)
				kind, _ = deviceWallet.DevicePinMatrixAck(deviceType, pinEnc)
			}
		},
	}
}
