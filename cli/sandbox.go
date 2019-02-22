package cli

import (
	"fmt"

	gcli "github.com/urfave/cli"

	deviceWallet "github.com/skycoin/hardware-wallet-go/device-wallet"
	"github.com/skycoin/hardware-wallet-go/device-wallet/messages"
)

func sandbox() gcli.Command {
	name := "sandbox"
	return gcli.Command{
		Name:         name,
		Usage:        "Sandbox.",
		Description:  "",
		Flags:        []gcli.Flag{},
		OnUsageError: onCommandUsageError(name),
		Action: func(_ *gcli.Context) {
			var deviceType deviceWallet.DeviceType
			if deviceWallet.DeviceConnected(deviceWallet.DeviceTypeEmulator) {
				deviceType = deviceWallet.DeviceTypeEmulator
			} else if deviceWallet.DeviceConnected(deviceType) {
				deviceType = deviceWallet.DeviceTypeUsb
			} else {
				log.Println("no device detected")
				return
			}

			deviceWallet.WipeDevice(deviceType)

			deviceWallet.DeviceSetMnemonic(deviceType, "cloud flower upset remain green metal below cup stem infant art thank")

			var pinEnc string
			kind, _ := deviceWallet.DeviceChangePin(deviceType)
			for kind == uint16(messages.MessageType_MessageType_PinMatrixRequest) {
				log.Printf("PinMatrixRequest response: ")
				fmt.Scanln(&pinEnc)
				kind, _ = deviceWallet.DevicePinMatrixAck(deviceType, pinEnc)
			}

			// come on one-more time
			// testing what happen when we try to change an existing pin code
			kind, _ = deviceWallet.DeviceChangePin(deviceType)
			for kind == uint16(messages.MessageType_MessageType_PinMatrixRequest) {
				log.Printf("PinMatrixRequest response: ")
				fmt.Scanln(&pinEnc)
				kind, _ = deviceWallet.DevicePinMatrixAck(deviceType, pinEnc)
			}

			var data []byte
			kind, data = deviceWallet.DeviceAddressGen(deviceType, 9, 15, false)
			if kind == uint16(messages.MessageType_MessageType_PinMatrixRequest) {
				log.Printf("PinMatrixRequest response: ")
				fmt.Scanln(&pinEnc)
				kind, data = deviceWallet.DevicePinMatrixAck(deviceType, pinEnc)

				if kind == uint16(messages.MessageType_MessageType_ResponseSkycoinAddress) {
					_, addresses := deviceWallet.DecodeResponseSkycoinAddress(kind, data)
					log.Print("Successfully got address")
					log.Print(addresses)
				}
			} else {
				log.Println("Got addresses without pin code")
				_, addresses := deviceWallet.DecodeResponseSkycoinAddress(kind, data)
				log.Print(addresses)
			}
		},
	}
}
