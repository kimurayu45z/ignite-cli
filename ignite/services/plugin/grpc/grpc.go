package grpc

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	v1 "github.com/ignite/cli/ignite/services/plugin/grpc/v1"
)

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

// HandshakeConfig are used to just do a basic handshake between a plugin and host.
// If the handshake fails, a user friendly error is shown. This prevents users from
// executing bad plugins or executing a plugin directory. It is a UX feature, not a
// security feature.
func HandshakeConfig() plugin.HandshakeConfig {
	return handshakeConfig
}

// NewPlugin returns a new gRPC plugin that implements the interface over gRPC.
func NewPlugin(impl Interface) plugin.Plugin {
	return grpcPlugin{impl: impl}
}

type grpcPlugin struct {
	// This is required by the Hashicorp plugin implementation for gRPC plugins
	plugin.Plugin

	impl Interface
}

// GRPCServer returns a new server that implements the plugin interface over gRPC.
func (p grpcPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	v1.RegisterInterfaceServiceServer(s, &server{impl: p.impl})
	return nil
}

// GRPCClient returns a new plugin client that allows calling the plugin interface over gRPC.
func (p grpcPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &client{grpc: v1.NewInterfaceServiceClient(c)}, nil
}

type client struct{ grpc v1.InterfaceServiceClient }

func (c client) Manifest() (*Manifest, error) {
	r, err := c.grpc.Manifest(context.Background(), &v1.ManifestRequest{})
	if err != nil {
		return nil, err
	}
	return r.Manifest, nil
}

func (c client) Execute(cmd *ExecutedCommand) error {
	_, err := c.grpc.Execute(context.Background(), &v1.ExecuteRequest{Cmd: cmd})
	return err
}

func (c client) ExecuteHookPre(h *ExecutedHook) error {
	_, err := c.grpc.ExecuteHookPre(context.Background(), &v1.ExecuteHookPreRequest{Hook: h})
	return err
}

func (c client) ExecuteHookPost(h *ExecutedHook) error {
	_, err := c.grpc.ExecuteHookPost(context.Background(), &v1.ExecuteHookPostRequest{Hook: h})
	return err
}

func (c client) ExecuteHookCleanUp(h *ExecutedHook) error {
	_, err := c.grpc.ExecuteHookCleanUp(context.Background(), &v1.ExecuteHookCleanUpRequest{Hook: h})
	return err
}

type server struct {
	v1.UnimplementedInterfaceServiceServer

	impl Interface
}

func (s server) Manifest(ctx context.Context, _ *v1.ManifestRequest) (*v1.ManifestResponse, error) {
	m, err := s.impl.Manifest(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.ManifestResponse{Manifest: m}, nil
}

func (s server) Execute(ctx context.Context, r *v1.ExecuteRequest) (*v1.ExecuteResponse, error) {
	err := s.impl.Execute(ctx, r.GetCmd())
	if err != nil {
		return nil, err
	}

	return &v1.ExecuteResponse{}, nil
}

func (s server) ExecuteHookPre(ctx context.Context, r *v1.ExecuteHookPreRequest) (*v1.ExecuteHookPreResponse, error) {
	err := s.impl.ExecuteHookPre(ctx, r.GetHook())
	if err != nil {
		return nil, err
	}

	return &v1.ExecuteHookPreResponse{}, nil
}

func (s server) ExecuteHookPost(ctx context.Context, r *v1.ExecuteHookPostRequest) (*v1.ExecuteHookPostResponse, error) {
	err := s.impl.ExecuteHookPost(ctx, r.GetHook())
	if err != nil {
		return nil, err
	}

	return &v1.ExecuteHookPostResponse{}, nil
}

func (s server) ExecuteHookCleanUp(ctx context.Context, r *v1.ExecuteHookCleanUpRequest) (*v1.ExecuteHookCleanUpResponse, error) {
	err := s.impl.ExecuteHookCleanUp(ctx, r.GetHook())
	if err != nil {
		return nil, err
	}

	return &v1.ExecuteHookCleanUpResponse{}, nil
}