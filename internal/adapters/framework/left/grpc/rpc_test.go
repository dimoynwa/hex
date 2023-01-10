package rpc

import (
	"context"
	"hex/internal/adapters/app/api"
	"hex/internal/adapters/core/arithmetic"
	"hex/internal/adapters/framework/left/grpc/pb"
	"hex/internal/adapters/framework/right/db"
	"hex/internal/ports"
	"log"
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	var err error

	lis = bufconn.Listen(bufSize)
	grpcServer := grpc.NewServer()

	// Ports
	var dbAdapter ports.DbPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var grpcAdapter ports.GRPCPort

	dbDriver := os.Getenv("DB_DRIVER") // "postgres"
	dsName := os.Getenv("DS_NAME")     // "jdbc:postgresql://localhost:5432/postgres"

	dbAdapter, err = db.NewAdapter(dbDriver, dsName)
	if err != nil {
		log.Fatalf("failed to start db: %v", err)
	}
	defer dbAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbAdapter, core)

	grpcAdapter = NewAdapter(appAdapter)

	pb.RegisterArithmeticServiceServer(grpcServer, grpcAdapter)

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("testing server start error: %v", err)
		}
	}()
}

func bufDialor(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func getGRPCConnection(ctx context.Context, t *testing.T) *grpc.ClientConn {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialor), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("fail to dial bufnet: %v", err)
	}
	return conn
}

func TestGetAddition(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{A: 1, B: 1}

	ans, err := client.GetAddition(ctx, params)

	if err != nil {
		t.Fatalf("expected 2, but was: %v", err)
	}

	require.Equal(t, int32(2), ans.GetValue())
}

func TestGetSubtraction(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{A: 1, B: 1}

	ans, err := client.GetSubtraction(ctx, params)

	if err != nil {
		t.Fatalf("expected 0, but was: %v", err)
	}

	require.Equal(t, int32(0), ans.GetValue())
}

func TestGetMultiplication(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{A: 1, B: 1}

	ans, err := client.GetMultiplication(ctx, params)

	if err != nil {
		t.Fatalf("expected 1, but was: %v", err)
	}

	require.Equal(t, int32(1), ans.GetValue())
}

func TestGetDivision(t *testing.T) {
	ctx := context.Background()
	conn := getGRPCConnection(ctx, t)
	defer conn.Close()

	client := pb.NewArithmeticServiceClient(conn)
	params := &pb.OperationParameters{A: 1, B: 1}

	ans, err := client.GetDivision(ctx, params)

	if err != nil {
		t.Fatalf("expected 1, but was: %v", err)
	}

	require.Equal(t, int32(1), ans.GetValue())
}
