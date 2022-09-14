package client

import (
	"context"

	"github.com/opencl-go/simulator-grpc/sim"
)

// Devices represents interactions on the devices layer.
type Devices struct {
	client sim.DevicesClient
}

// Create requests to create a new device. The returned value is the ID value of the created device.
// Use this function for any new test that you intend to run in parallel to others.
func (devices *Devices) Create(ctx context.Context) (uintptr, error) {
	request := sim.DeviceCreateRequest{}
	response, err := devices.client.Create(ctx, &request)
	if err != nil {
		return 0, err
	}
	return uintptr(response.ID), nil
}

// Delete requests to remove a previously created device.
// Call this function at the end of the test, to ensure resources are removed.
func (devices *Devices) Delete(ctx context.Context, id uintptr) error {
	request := sim.DeviceDeleteRequest{ID: uint64(id)}
	_, err := devices.client.Delete(ctx, &request)
	return err
}
