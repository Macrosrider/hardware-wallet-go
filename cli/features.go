package cli

import (
	gcli "github.com/urfave/cli"

	"log"

	deviceWallet "github.com/skycoin/hardware-wallet-go/device-wallet"
)

func featuresCmd() gcli.Command {
	name := "features"
	return gcli.Command{
		Name:         name,
		Usage:        "Ask the device Features.",
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

			deviceWallet.DeviceGetFeatures(deviceType)
		},
	}
}
