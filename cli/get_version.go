package cli

import (
	"fmt"

	gcli "github.com/urfave/cli"

	"log"

	deviceWallet "github.com/skycoin/hardware-wallet-go/device-wallet"
)

func getVersionCmd() gcli.Command {
	name := "getVersion"
	return gcli.Command{
		Name:         name,
		Usage:        "Ask firmware version.",
		Description:  "",
		OnUsageError: onCommandUsageError(name),
		Flags: []gcli.Flag{
			gcli.StringFlag{
				Name:   "deviceType",
				Usage:  "Device type to send instructions to, hardware wallet (USB) or emulator.",
				EnvVar: "DEVICE_TYPE",
				Value:  "USB",
			},
		},
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

			version := deviceWallet.DeviceGetVersion(deviceType)
			fmt.Printf("Firmware version is %s\n", version)
		},
	}
}
