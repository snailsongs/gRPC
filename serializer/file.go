package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

/**
1.write protobuf data to binary file
2.read protobuf data from binary file
3.write protobuf data to json file
4.compare size of binary file and json file
*/

// WriteProtoToBinary write protobuf data to binary file
func WriteProtoToBinary(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("could not encode message to binary: %v", err)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("could not write binary file: %v", err)
	}
	return nil
}

// ReadProtoToBinary read protobuf data from binary file
func ReadProtoToBinary(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("could not read binary file: %v", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("could not decode binary file: %v", err)
	}
	return nil
}

func WriteProtoToJson(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("could not encode message to json: %v", err)
	}
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("could not write json file: %v", err)
	}
	return nil
}
