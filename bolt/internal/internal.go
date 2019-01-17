package internal

import (
	"github.com/gogo/protobuf/proto"
	"github.com/rajatparida86/wtfdial"
)

//go:generate protoc --gogo_out=. internal.proto

func MarshalDial(dial *wtf.Dial) ([]byte, error) {
	bytes, err := proto.Marshal(&Dial{
		ID:           proto.Int64(int64(dial.ID)),
		UserID:       proto.Int64(int64(dial.UserID)),
		Status:       proto.Float64(float64(dial.Status)),
		ModifiedTime: proto.Int64(dial.ModifiedTime.UnixNano()),
	})
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func UnMarshalDial() {

}
