// Code generated by GoVPP binapi-generator. DO NOT EDIT.
// source: memif.api.json

/*
Package memif is a generated VPP binary API of the 'memif' VPP module.

It is generated from this file:
	memif.api.json

It contains these VPP binary API objects:
	10 messages
	5 services
*/
package memif

import "git.fd.io/govpp.git/api"
import "github.com/lunixbochs/struc"
import "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = struc.Pack
var _ = bytes.NewBuffer

/* Messages */

// MemifSocketFilenameAddDel represents the VPP binary API message 'memif_socket_filename_add_del'.
// Generated from 'memif.api.json', line 4:
//
//            "memif_socket_filename_add_del",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "is_add"
//            ],
//            [
//                "u32",
//                "socket_id"
//            ],
//            [
//                "u8",
//                "socket_filename",
//                128
//            ],
//            {
//                "crc": "0x30e3929d"
//            }
//
type MemifSocketFilenameAddDel struct {
	IsAdd          uint8
	SocketID       uint32
	SocketFilename []byte `struc:"[128]byte"`
}

func (*MemifSocketFilenameAddDel) GetMessageName() string {
	return "memif_socket_filename_add_del"
}
func (*MemifSocketFilenameAddDel) GetCrcString() string {
	return "30e3929d"
}
func (*MemifSocketFilenameAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewMemifSocketFilenameAddDel() api.Message {
	return &MemifSocketFilenameAddDel{}
}

// MemifSocketFilenameAddDelReply represents the VPP binary API message 'memif_socket_filename_add_del_reply'.
// Generated from 'memif.api.json', line 35:
//
//            "memif_socket_filename_add_del_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type MemifSocketFilenameAddDelReply struct {
	Retval int32
}

func (*MemifSocketFilenameAddDelReply) GetMessageName() string {
	return "memif_socket_filename_add_del_reply"
}
func (*MemifSocketFilenameAddDelReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MemifSocketFilenameAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewMemifSocketFilenameAddDelReply() api.Message {
	return &MemifSocketFilenameAddDelReply{}
}

// MemifCreate represents the VPP binary API message 'memif_create'.
// Generated from 'memif.api.json', line 53:
//
//            "memif_create",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u8",
//                "role"
//            ],
//            [
//                "u8",
//                "mode"
//            ],
//            [
//                "u8",
//                "rx_queues"
//            ],
//            [
//                "u8",
//                "tx_queues"
//            ],
//            [
//                "u32",
//                "id"
//            ],
//            [
//                "u32",
//                "socket_id"
//            ],
//            [
//                "u8",
//                "secret",
//                24
//            ],
//            [
//                "u32",
//                "ring_size"
//            ],
//            [
//                "u16",
//                "buffer_size"
//            ],
//            [
//                "u8",
//                "hw_addr",
//                6
//            ],
//            {
//                "crc": "0x6597cdb2"
//            }
//
type MemifCreate struct {
	Role       uint8
	Mode       uint8
	RxQueues   uint8
	TxQueues   uint8
	ID         uint32
	SocketID   uint32
	Secret     []byte `struc:"[24]byte"`
	RingSize   uint32
	BufferSize uint16
	HwAddr     []byte `struc:"[6]byte"`
}

func (*MemifCreate) GetMessageName() string {
	return "memif_create"
}
func (*MemifCreate) GetCrcString() string {
	return "6597cdb2"
}
func (*MemifCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewMemifCreate() api.Message {
	return &MemifCreate{}
}

// MemifCreateReply represents the VPP binary API message 'memif_create_reply'.
// Generated from 'memif.api.json', line 113:
//
//            "memif_create_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            {
//                "crc": "0xfda5941f"
//            }
//
type MemifCreateReply struct {
	Retval    int32
	SwIfIndex uint32
}

func (*MemifCreateReply) GetMessageName() string {
	return "memif_create_reply"
}
func (*MemifCreateReply) GetCrcString() string {
	return "fda5941f"
}
func (*MemifCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewMemifCreateReply() api.Message {
	return &MemifCreateReply{}
}

// MemifDelete represents the VPP binary API message 'memif_delete'.
// Generated from 'memif.api.json', line 135:
//
//            "memif_delete",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            {
//                "crc": "0x529cb13f"
//            }
//
type MemifDelete struct {
	SwIfIndex uint32
}

func (*MemifDelete) GetMessageName() string {
	return "memif_delete"
}
func (*MemifDelete) GetCrcString() string {
	return "529cb13f"
}
func (*MemifDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewMemifDelete() api.Message {
	return &MemifDelete{}
}

// MemifDeleteReply represents the VPP binary API message 'memif_delete_reply'.
// Generated from 'memif.api.json', line 157:
//
//            "memif_delete_reply",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "i32",
//                "retval"
//            ],
//            {
//                "crc": "0xe8d4e804"
//            }
//
type MemifDeleteReply struct {
	Retval int32
}

func (*MemifDeleteReply) GetMessageName() string {
	return "memif_delete_reply"
}
func (*MemifDeleteReply) GetCrcString() string {
	return "e8d4e804"
}
func (*MemifDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewMemifDeleteReply() api.Message {
	return &MemifDeleteReply{}
}

// MemifSocketFilenameDetails represents the VPP binary API message 'memif_socket_filename_details'.
// Generated from 'memif.api.json', line 175:
//
//            "memif_socket_filename_details",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "socket_id"
//            ],
//            [
//                "u8",
//                "socket_filename",
//                128
//            ],
//            {
//                "crc": "0xe347e32f"
//            }
//
type MemifSocketFilenameDetails struct {
	SocketID       uint32
	SocketFilename []byte `struc:"[128]byte"`
}

func (*MemifSocketFilenameDetails) GetMessageName() string {
	return "memif_socket_filename_details"
}
func (*MemifSocketFilenameDetails) GetCrcString() string {
	return "e347e32f"
}
func (*MemifSocketFilenameDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewMemifSocketFilenameDetails() api.Message {
	return &MemifSocketFilenameDetails{}
}

// MemifSocketFilenameDump represents the VPP binary API message 'memif_socket_filename_dump'.
// Generated from 'memif.api.json', line 198:
//
//            "memif_socket_filename_dump",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type MemifSocketFilenameDump struct{}

func (*MemifSocketFilenameDump) GetMessageName() string {
	return "memif_socket_filename_dump"
}
func (*MemifSocketFilenameDump) GetCrcString() string {
	return "51077d14"
}
func (*MemifSocketFilenameDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewMemifSocketFilenameDump() api.Message {
	return &MemifSocketFilenameDump{}
}

// MemifDetails represents the VPP binary API message 'memif_details'.
// Generated from 'memif.api.json', line 216:
//
//            "memif_details",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            [
//                "u32",
//                "sw_if_index"
//            ],
//            [
//                "u8",
//                "if_name",
//                64
//            ],
//            [
//                "u8",
//                "hw_addr",
//                6
//            ],
//            [
//                "u32",
//                "id"
//            ],
//            [
//                "u8",
//                "role"
//            ],
//            [
//                "u8",
//                "mode"
//            ],
//            [
//                "u32",
//                "socket_id"
//            ],
//            [
//                "u32",
//                "ring_size"
//            ],
//            [
//                "u16",
//                "buffer_size"
//            ],
//            [
//                "u8",
//                "admin_up_down"
//            ],
//            [
//                "u8",
//                "link_up_down"
//            ],
//            {
//                "crc": "0x4f5a3397"
//            }
//
type MemifDetails struct {
	SwIfIndex   uint32
	IfName      []byte `struc:"[64]byte"`
	HwAddr      []byte `struc:"[6]byte"`
	ID          uint32
	Role        uint8
	Mode        uint8
	SocketID    uint32
	RingSize    uint32
	BufferSize  uint16
	AdminUpDown uint8
	LinkUpDown  uint8
}

func (*MemifDetails) GetMessageName() string {
	return "memif_details"
}
func (*MemifDetails) GetCrcString() string {
	return "4f5a3397"
}
func (*MemifDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}
func NewMemifDetails() api.Message {
	return &MemifDetails{}
}

// MemifDump represents the VPP binary API message 'memif_dump'.
// Generated from 'memif.api.json', line 276:
//
//            "memif_dump",
//            [
//                "u16",
//                "_vl_msg_id"
//            ],
//            [
//                "u32",
//                "client_index"
//            ],
//            [
//                "u32",
//                "context"
//            ],
//            {
//                "crc": "0x51077d14"
//            }
//
type MemifDump struct{}

func (*MemifDump) GetMessageName() string {
	return "memif_dump"
}
func (*MemifDump) GetCrcString() string {
	return "51077d14"
}
func (*MemifDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}
func NewMemifDump() api.Message {
	return &MemifDump{}
}

/* Services */

type Services interface {
	DumpMemif(*MemifDump) (*MemifDetails, error)
	DumpMemifSocketFilename(*MemifSocketFilenameDump) (*MemifSocketFilenameDetails, error)
	MemifCreate(*MemifCreate) (*MemifCreateReply, error)
	MemifDelete(*MemifDelete) (*MemifDeleteReply, error)
	MemifSocketFilenameAddDel(*MemifSocketFilenameAddDel) (*MemifSocketFilenameAddDelReply, error)
}

func init() {
	api.RegisterMessage((*MemifSocketFilenameAddDel)(nil), "memif.MemifSocketFilenameAddDel")
	api.RegisterMessage((*MemifSocketFilenameAddDelReply)(nil), "memif.MemifSocketFilenameAddDelReply")
	api.RegisterMessage((*MemifCreate)(nil), "memif.MemifCreate")
	api.RegisterMessage((*MemifCreateReply)(nil), "memif.MemifCreateReply")
	api.RegisterMessage((*MemifDelete)(nil), "memif.MemifDelete")
	api.RegisterMessage((*MemifDeleteReply)(nil), "memif.MemifDeleteReply")
	api.RegisterMessage((*MemifSocketFilenameDetails)(nil), "memif.MemifSocketFilenameDetails")
	api.RegisterMessage((*MemifSocketFilenameDump)(nil), "memif.MemifSocketFilenameDump")
	api.RegisterMessage((*MemifDetails)(nil), "memif.MemifDetails")
	api.RegisterMessage((*MemifDump)(nil), "memif.MemifDump")
}
