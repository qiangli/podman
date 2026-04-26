// Package embed provides pure-Go components from the Podman project for
// embedding in external Go applications.
//
// This package re-exports Podman's pure-Go REST API client bindings
// (pkg/bindings) and entity types for container, image, network, pod,
// and volume management. These packages communicate with a running Podman
// service via Unix socket or HTTP — no CGo dependencies.
//
// The Podman service itself is managed as an external process by the host
// application, similar to the Ollama embedding pattern.
package embed

import (
	"context"

	"go.podman.io/podman/v6/pkg/bindings"
	"go.podman.io/podman/v6/pkg/bindings/containers"
	"go.podman.io/podman/v6/pkg/bindings/images"
	"go.podman.io/podman/v6/pkg/bindings/network"
	"go.podman.io/podman/v6/pkg/bindings/pods"
	"go.podman.io/podman/v6/pkg/bindings/system"
	"go.podman.io/podman/v6/pkg/bindings/volumes"
	entTypes "go.podman.io/podman/v6/pkg/domain/entities/types"
)

// Re-export connection management.
type Connection = bindings.Connection

// NewConnection creates a connection context to a Podman service.
// The uri should be a Unix socket path (e.g., "unix:///run/podman/podman.sock")
// or an HTTP endpoint.
func NewConnection(ctx context.Context, uri string) (context.Context, error) {
	return bindings.NewConnection(ctx, uri)
}

// --- Container operations ---

var (
	ContainerList       = containers.List
	ContainerInspect    = containers.Inspect
	ContainerStart      = containers.Start
	ContainerStop       = containers.Stop
	ContainerRemove     = containers.Remove
	ContainerCreateSpec = containers.CreateWithSpec
	ContainerWait       = containers.Wait
	ContainerLogs       = containers.Logs
	ContainerExecCreate  = containers.ExecCreate
	ContainerExecStart   = containers.ExecStart
	ContainerExecInspect = containers.ExecInspect
)

type (
	ContainerListOptions    = containers.ListOptions
	ContainerInspectOptions = containers.InspectOptions
	ContainerStartOptions   = containers.StartOptions
	ContainerStopOptions    = containers.StopOptions
	ContainerRemoveOptions  = containers.RemoveOptions
	ContainerCreateOptions  = containers.CreateOptions
	ContainerLogsOptions    = containers.LogOptions
	ListContainer           = entTypes.ListContainer
)

// --- Image operations ---

var (
	ImageList   = images.List
	ImageExists = images.Exists
	ImagePull   = images.Pull
	ImageRemove = images.Remove
	ImageBuild  = images.Build
)

type (
	ImageListOptions   = images.ListOptions
	ImagePullOptions   = images.PullOptions
	ImageRemoveOptions = images.RemoveOptions
	ImageSummary       = entTypes.ImageSummary
)

// --- Network operations ---

var (
	NetworkCreate  = network.Create
	NetworkList    = network.List
	NetworkRemove  = network.Remove
	NetworkInspect = network.Inspect
)

type (
	NetworkCreateOptions  = network.CreateOptions
	NetworkListOptions    = network.ListOptions
	NetworkRemoveOptions  = network.RemoveOptions
	NetworkInspectOptions = network.InspectOptions
)

// --- Pod operations ---

var (
	PodCreateFromSpec = pods.CreatePodFromSpec
	PodStart          = pods.Start
	PodStop           = pods.Stop
	PodRemove         = pods.Remove
	PodList           = pods.List
	PodInspect        = pods.Inspect
)

type (
	PodStartOptions  = pods.StartOptions
	PodStopOptions   = pods.StopOptions
	PodRemoveOptions = pods.RemoveOptions
	PodListOptions   = pods.ListOptions
)

// --- Volume operations ---

var (
	VolumeCreate  = volumes.Create
	VolumeList    = volumes.List
	VolumeRemove  = volumes.Remove
	VolumeInspect = volumes.Inspect
)

type (
	VolumeCreateOptions = volumes.CreateOptions
	VolumeListOptions   = volumes.ListOptions
	VolumeRemoveOptions = volumes.RemoveOptions
)

// --- System operations ---

var (
	SystemInfo    = system.Info
	SystemVersion = system.Version
)
