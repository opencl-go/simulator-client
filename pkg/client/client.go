package client

import (
	"github.com/opencl-go/simulator-grpc/sim"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// PlatformServerAddressInfo is the PlatformInfoName that shall hold the server address of the simulator.
	// This is a forward from the grpc interface, so that the user of this package doesn't additionally have to
	// import the grpc package, just for this constant.
	PlatformServerAddressInfo = sim.PlatformServerAddressInfo
)

// Client represents the simulator client, to be used in tests.
type Client struct {
	connection *grpc.ClientConn
	platform   Platform
	devices    Devices
}

// NewClient attempts to connect to the given address. Retrieve the value of the address by querying the
// OpenCL platform for the PlatformServerAddressInfo string.
func NewClient(address string) (*Client, error) {
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := &Client{
		connection: connection,
		platform:   Platform{client: sim.NewPlatformClient(connection)},
		devices:    Devices{client: sim.NewDevicesClient(connection)},
	}
	return client, nil
}

// Disconnect shuts down the connection.
func (client *Client) Disconnect() {
	_ = client.connection.Close()
}

// Platform is the accessor for platform-specific interaction.
func (client *Client) Platform() *Platform {
	return &client.platform
}

// Devices is the accessor for device-specific interaction.
func (client *Client) Devices() *Devices {
	return &client.devices
}
