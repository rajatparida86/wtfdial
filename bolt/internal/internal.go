package internal

import (
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/rajatparida86/wtfdial"
)

//go:generate protoc --gogo_out=. internal.proto

// MarshalDial -- Marshal wtf.Dial to byte array
func MarshalDial(dial *wtf.Dial) ([]byte, error) {
	bytes, err := proto.Marshal(&Dial{
		ID:           proto.Int64(int64(dial.ID)),
		UserID:       proto.Int64(int64(dial.UserID)),
		Status:       proto.Float64(float64(dial.Status)),
		Name:         proto.String(dial.Name),
		ModifiedTime: proto.Int64(dial.ModifiedTime.UnixNano()),
	})
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// UnMarshalDial -- Converts byte array to wtf.Dial
func UnMarshalDial(bytes []byte) (*wtf.Dial, error) {
	var protobuf Dial
	if err := proto.Unmarshal(bytes, &protobuf); err != nil {
		return nil, err
	}
	var dial wtf.Dial
	dial.ID = wtf.DialID(protobuf.GetID())
	dial.UserID = wtf.UserID(protobuf.GetUserID())
	dial.Status = protobuf.GetStatus()
	dial.Name = protobuf.GetName()
	dial.ModifiedTime = time.Unix(0, protobuf.GetModifiedTime()).UTC()
	return &dial, nil
}
