package config

import (
	"fmt"
	"os"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

var SnapClient snap.Client

func InitMidtrans() {
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	if serverKey == "" {
		fmt.Println("Error: MIDTRANS_SERVER_KEY is empty!")
		return
	}

	midtrans.ServerKey = serverKey
	midtrans.Environment = midtrans.Sandbox // Ganti ke midtrans.Production jika live

	SnapClient.New(midtrans.ServerKey, midtrans.Sandbox)
	fmt.Println("Midtrans initialized successfully")
}
