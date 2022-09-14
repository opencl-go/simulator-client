package client

import (
	"context"

	"github.com/opencl-go/simulator-grpc/sim"
)

// Platform represents interactions on the platform layer.
type Platform struct {
	client sim.PlatformClient
}

// PrepareExtensionFunction requests to prepare an extension function for testing.
// The returned value is the address value of the function pointer. Note that the simulator has a limited amount
// of functions to return.
func (platform *Platform) PrepareExtensionFunction(ctx context.Context) (string, uintptr, error) {
	request := sim.PlatformPrepareExtensionFunctionRequest{}
	response, err := platform.client.PrepareExtensionFunction(ctx, &request)
	if err != nil {
		return "", 0, err
	}
	return response.Name, uintptr(response.Address), nil
}

// ReleaseExtensionFunction requests to release a previously prepared extension function.
// Call this function at the end of the test, to ensure resources are removed.
func (platform *Platform) ReleaseExtensionFunction(ctx context.Context, name string) error {
	request := sim.PlatformReleaseExtensionFunctionRequest{Name: name}
	_, err := platform.client.ReleaseExtensionFunction(ctx, &request)
	return err
}
