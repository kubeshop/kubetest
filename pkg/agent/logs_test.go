package agent_test

import (
	"context"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/kubeshop/testkube/pkg/agent"
	"github.com/kubeshop/testkube/pkg/cloud"
	"github.com/kubeshop/testkube/pkg/executor/output"
	"github.com/kubeshop/testkube/pkg/log"
	"github.com/kubeshop/testkube/pkg/ui"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func TestLogStream(t *testing.T) {
	url := "localhost:8997"

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		lis, err := net.Listen("tcp", url)
		if err != nil {
			panic(err)
		}

		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		cloud.RegisterTestKubeCloudAPIServer(grpcServer, newLogStreamServer(ctx))
		grpcServer.Serve(lis)
	}()

	m := func(ctx *fasthttp.RequestCtx) {
		h := &ctx.Response.Header
		h.Add("Content-type", "application/json")
		fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
	}

	config := agent.Config{Insecure: true}
	grpcConn, err := agent.NewGRPCConnection(context.Background(), url, config, log.DefaultLogger)
	ui.ExitOnError("error creating gRPC connection", err)
	defer grpcConn.Close()

	grpcClient := cloud.NewTestKubeCloudAPIClient(grpcConn)

	var msgCnt int32
	logStreamFunc := func(ctx context.Context, executionID string) (chan output.Output, error) {
		ch := make(chan output.Output, 5)

		ch <- output.Output{
			Type_:   output.TypeLogLine,
			Content: "log1",
			Time:    time.Now(),
		}
		msgCnt++
		return ch, nil
	}

	logger, _ := zap.NewDevelopment()
	agent, err := agent.NewAgent(logger.Sugar(), m, "api-key", grpcClient, 5, 5, logStreamFunc, "", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	g, groupCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return agent.Run(groupCtx)
	})

	time.Sleep(100 * time.Millisecond)
	cancel()

	g.Wait()

	assert.NotZero(t, msgCnt)
}

type CloudLogsServer struct {
	ctx context.Context
	cloud.UnimplementedTestKubeCloudAPIServer
}

func (cs *CloudLogsServer) ExecuteAsync(srv cloud.TestKubeCloudAPI_ExecuteAsyncServer) error {
	<-cs.ctx.Done()
	return nil
}
func (cs *CloudLogsServer) GetLogsStream(srv cloud.TestKubeCloudAPI_GetLogsStreamServer) error {
	md, ok := metadata.FromIncomingContext(srv.Context())
	if !ok {
		panic("no metadata")
	}
	apiKey := md.Get("api-key")
	if apiKey[0] != "api-key" {
		panic("error bad api-key")
	}

	req := &cloud.LogsStreamRequest{
		StreamId:    "streamID",
		ExecutionId: "executionID",
	}
	err := srv.Send(req)
	if err != nil {
		return err
	}

	resp, err := srv.Recv()
	if err != nil {
		return err
	}
	fmt.Println(resp)

	return nil
}

func newLogStreamServer(ctx context.Context) *CloudLogsServer {
	return &CloudLogsServer{ctx: ctx}
}
