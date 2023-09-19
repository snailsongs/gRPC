package serializer_test

import (
	"fmt"
	"testing"

	"gRPC/pb"
	"gRPC/sample"
	"gRPC/serializer"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

/**
1.write protobuf data to binary file
*/

func TestFileSerializer(t *testing.T) {
	t.Parallel()
	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtoToBinary(laptop1, binaryFile)
	if err != nil {
		t.Fatal(err)
	}

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtoToBinary(binaryFile, laptop2)
	if err != nil {
		return
	}
	if proto.Equal(laptop1, laptop2) == true {
		t.Logf("laptop1: %v", laptop1)
		t.Logf("laptop2: %v", laptop2)
		fmt.Println("equal")
	}

	err = serializer.WriteProtoToJson(laptop1, jsonFile)
	require.NoError(t, err)
}
