// ************************************************************
// This file is automatically generated by genxdr. Do not edit.
// ************************************************************

package protocol

import (
	"github.com/calmh/xdr"
)

/*

header Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             magic                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                         message Type                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        message Length                         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct header {
	unsigned int magic;
	int messageType;
	int messageLength;
}

*/

func (o header) XDRSize() int {
	return 4 + 4 + 4
}

func (o header) MarshalXDR() ([]byte, error) {
	buf := make([]byte, o.XDRSize())
	m := &xdr.Marshaller{Data: buf}
	return buf, o.MarshalXDRInto(m)
}

func (o header) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o header) MarshalXDRInto(m *xdr.Marshaller) error {
	m.MarshalUint32(o.magic)
	m.MarshalUint32(uint32(o.messageType))
	m.MarshalUint32(uint32(o.messageLength))
	return m.Error
}

func (o *header) UnmarshalXDR(bs []byte) error {
	u := &xdr.Unmarshaller{Data: bs}
	return o.UnmarshalXDRFrom(u)
}
func (o *header) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	o.magic = u.UnmarshalUint32()
	o.messageType = int32(u.UnmarshalUint32())
	o.messageLength = int32(u.UnmarshalUint32())
	return u.Error
}

/*

Ping Structure:
(contains no fields)


struct Ping {
}

*/

func (o Ping) XDRSize() int {
	return 0
}
func (o Ping) MarshalXDR() ([]byte, error) {
	return nil, nil
}

func (o Ping) MustMarshalXDR() []byte {
	return nil
}

func (o Ping) MarshalXDRInto(m *xdr.Marshaller) error {
	return nil
}

func (o *Ping) UnmarshalXDR(bs []byte) error {
	return nil
}

func (o *Ping) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	return nil
}

/*

Pong Structure:
(contains no fields)


struct Pong {
}

*/

func (o Pong) XDRSize() int {
	return 0
}
func (o Pong) MarshalXDR() ([]byte, error) {
	return nil, nil
}

func (o Pong) MustMarshalXDR() []byte {
	return nil
}

func (o Pong) MarshalXDRInto(m *xdr.Marshaller) error {
	return nil
}

func (o *Pong) UnmarshalXDR(bs []byte) error {
	return nil
}

func (o *Pong) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	return nil
}

/*

JoinRelayRequest Structure:
(contains no fields)


struct JoinRelayRequest {
}

*/

func (o JoinRelayRequest) XDRSize() int {
	return 0
}
func (o JoinRelayRequest) MarshalXDR() ([]byte, error) {
	return nil, nil
}

func (o JoinRelayRequest) MustMarshalXDR() []byte {
	return nil
}

func (o JoinRelayRequest) MarshalXDRInto(m *xdr.Marshaller) error {
	return nil
}

func (o *JoinRelayRequest) UnmarshalXDR(bs []byte) error {
	return nil
}

func (o *JoinRelayRequest) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	return nil
}

/*

RelayFull Structure:
(contains no fields)


struct RelayFull {
}

*/

func (o RelayFull) XDRSize() int {
	return 0
}
func (o RelayFull) MarshalXDR() ([]byte, error) {
	return nil, nil
}

func (o RelayFull) MustMarshalXDR() []byte {
	return nil
}

func (o RelayFull) MarshalXDRInto(m *xdr.Marshaller) error {
	return nil
}

func (o *RelayFull) UnmarshalXDR(bs []byte) error {
	return nil
}

func (o *RelayFull) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	return nil
}

/*

JoinSessionRequest Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                  Key (length + padded data)                   \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct JoinSessionRequest {
	opaque Key<32>;
}

*/

func (o JoinSessionRequest) XDRSize() int {
	return 4 + len(o.Key) + xdr.Padding(len(o.Key))
}

func (o JoinSessionRequest) MarshalXDR() ([]byte, error) {
	buf := make([]byte, o.XDRSize())
	m := &xdr.Marshaller{Data: buf}
	return buf, o.MarshalXDRInto(m)
}

func (o JoinSessionRequest) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o JoinSessionRequest) MarshalXDRInto(m *xdr.Marshaller) error {
	if l := len(o.Key); l > 32 {
		return xdr.ElementSizeExceeded("Key", l, 32)
	}
	m.MarshalBytes(o.Key)
	return m.Error
}

func (o *JoinSessionRequest) UnmarshalXDR(bs []byte) error {
	u := &xdr.Unmarshaller{Data: bs}
	return o.UnmarshalXDRFrom(u)
}
func (o *JoinSessionRequest) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	o.Key = u.UnmarshalBytesMax(32)
	return u.Error
}

/*

Response Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                             Code                              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                Message (length + padded data)                 \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct Response {
	int Code;
	string Message<>;
}

*/

func (o Response) XDRSize() int {
	return 4 +
		4 + len(o.Message) + xdr.Padding(len(o.Message))
}

func (o Response) MarshalXDR() ([]byte, error) {
	buf := make([]byte, o.XDRSize())
	m := &xdr.Marshaller{Data: buf}
	return buf, o.MarshalXDRInto(m)
}

func (o Response) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o Response) MarshalXDRInto(m *xdr.Marshaller) error {
	m.MarshalUint32(uint32(o.Code))
	m.MarshalString(o.Message)
	return m.Error
}

func (o *Response) UnmarshalXDR(bs []byte) error {
	u := &xdr.Unmarshaller{Data: bs}
	return o.UnmarshalXDRFrom(u)
}
func (o *Response) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	o.Code = int32(u.UnmarshalUint32())
	o.Message = u.UnmarshalString()
	return u.Error
}

/*

ConnectRequest Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                   ID (length + padded data)                   \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct ConnectRequest {
	opaque ID<32>;
}

*/

func (o ConnectRequest) XDRSize() int {
	return 4 + len(o.ID) + xdr.Padding(len(o.ID))
}

func (o ConnectRequest) MarshalXDR() ([]byte, error) {
	buf := make([]byte, o.XDRSize())
	m := &xdr.Marshaller{Data: buf}
	return buf, o.MarshalXDRInto(m)
}

func (o ConnectRequest) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o ConnectRequest) MarshalXDRInto(m *xdr.Marshaller) error {
	if l := len(o.ID); l > 32 {
		return xdr.ElementSizeExceeded("ID", l, 32)
	}
	m.MarshalBytes(o.ID)
	return m.Error
}

func (o *ConnectRequest) UnmarshalXDR(bs []byte) error {
	u := &xdr.Unmarshaller{Data: bs}
	return o.UnmarshalXDRFrom(u)
}
func (o *ConnectRequest) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	o.ID = u.UnmarshalBytesMax(32)
	return u.Error
}

/*

SessionInvitation Structure:

 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                  From (length + padded data)                  \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                  Key (length + padded data)                   \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
/                                                               /
\                Address (length + padded data)                 \
/                                                               /
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|         16 zero bits          |             Port              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                  Server Socket (V=0 or 1)                   |V|
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+


struct SessionInvitation {
	opaque From<32>;
	opaque Key<32>;
	opaque Address<32>;
	unsigned int Port;
	bool ServerSocket;
}

*/

func (o SessionInvitation) XDRSize() int {
	return 4 + len(o.From) + xdr.Padding(len(o.From)) +
		4 + len(o.Key) + xdr.Padding(len(o.Key)) +
		4 + len(o.Address) + xdr.Padding(len(o.Address)) + 4 + 4
}

func (o SessionInvitation) MarshalXDR() ([]byte, error) {
	buf := make([]byte, o.XDRSize())
	m := &xdr.Marshaller{Data: buf}
	return buf, o.MarshalXDRInto(m)
}

func (o SessionInvitation) MustMarshalXDR() []byte {
	bs, err := o.MarshalXDR()
	if err != nil {
		panic(err)
	}
	return bs
}

func (o SessionInvitation) MarshalXDRInto(m *xdr.Marshaller) error {
	if l := len(o.From); l > 32 {
		return xdr.ElementSizeExceeded("From", l, 32)
	}
	m.MarshalBytes(o.From)
	if l := len(o.Key); l > 32 {
		return xdr.ElementSizeExceeded("Key", l, 32)
	}
	m.MarshalBytes(o.Key)
	if l := len(o.Address); l > 32 {
		return xdr.ElementSizeExceeded("Address", l, 32)
	}
	m.MarshalBytes(o.Address)
	m.MarshalUint16(o.Port)
	m.MarshalBool(o.ServerSocket)
	return m.Error
}

func (o *SessionInvitation) UnmarshalXDR(bs []byte) error {
	u := &xdr.Unmarshaller{Data: bs}
	return o.UnmarshalXDRFrom(u)
}
func (o *SessionInvitation) UnmarshalXDRFrom(u *xdr.Unmarshaller) error {
	o.From = u.UnmarshalBytesMax(32)
	o.Key = u.UnmarshalBytesMax(32)
	o.Address = u.UnmarshalBytesMax(32)
	o.Port = u.UnmarshalUint16()
	o.ServerSocket = u.UnmarshalBool()
	return u.Error
}