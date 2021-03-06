// Code generated by protoc-gen-go.
// source: examples/add/addservice/addservice.proto
// DO NOT EDIT!

package addservice

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

import "net"
import "net/rpc"
import "github.com/pguelpa/go-rpcgen/codec"
import "net/url"
import "net/http"
import "github.com/pguelpa/go-rpcgen/webrpc"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type AddMessage struct {
	X                *int32 `protobuf:"varint,1,req,name=x" json:"x,omitempty"`
	Y                *int32 `protobuf:"varint,2,req,name=y" json:"y,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *AddMessage) Reset()         { *this = AddMessage{} }
func (this *AddMessage) String() string { return proto.CompactTextString(this) }
func (*AddMessage) ProtoMessage()       {}

func (this *AddMessage) GetX() int32 {
	if this != nil && this.X != nil {
		return *this.X
	}
	return 0
}

func (this *AddMessage) GetY() int32 {
	if this != nil && this.Y != nil {
		return *this.Y
	}
	return 0
}

type SumMessage struct {
	Z                *int32 `protobuf:"varint,1,req,name=z" json:"z,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (this *SumMessage) Reset()         { *this = SumMessage{} }
func (this *SumMessage) String() string { return proto.CompactTextString(this) }
func (*SumMessage) ProtoMessage()       {}

func (this *SumMessage) GetZ() int32 {
	if this != nil && this.Z != nil {
		return *this.Z
	}
	return 0
}

func init() {
}

// AddService is an interface satisfied by the generated client and
// which must be implemented by the object wrapped by the server.
type AddService interface {
	Add(in *AddMessage, out *SumMessage) error
}

// internal wrapper for type-safe RPC calling
type rpcAddServiceClient struct {
	*rpc.Client
}

func (this rpcAddServiceClient) Add(in *AddMessage, out *SumMessage) error {
	return this.Call("AddService.Add", in, out)
}

// NewAddServiceClient returns an *rpc.Client wrapper for calling the methods of
// AddService remotely.
func NewAddServiceClient(conn net.Conn) AddService {
	return rpcAddServiceClient{rpc.NewClientWithCodec(codec.NewClientCodec(conn))}
}

// ServeAddService serves the given AddService backend implementation on conn.
func ServeAddService(conn net.Conn, backend AddService) error {
	srv := rpc.NewServer()
	if err := srv.RegisterName("AddService", backend); err != nil {
		return err
	}
	srv.ServeCodec(codec.NewServerCodec(conn))
	return nil
}

// DialAddService returns a AddService for calling the AddService servince at addr (TCP).
func DialAddService(addr string) (AddService, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return NewAddServiceClient(conn), nil
}

// ListenAndServeAddService serves the given AddService backend implementation
// on all connections accepted as a result of listening on addr (TCP).
func ListenAndServeAddService(addr string, backend AddService) error {
	clients, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	srv := rpc.NewServer()
	if err := srv.RegisterName("AddService", backend); err != nil {
		return err
	}
	for {
		conn, err := clients.Accept()
		if err != nil {
			return err
		}
		go srv.ServeCodec(codec.NewServerCodec(conn))
	}
	panic("unreachable")
}

// AddServiceWeb is the web-based RPC version of the interface which
// must be implemented by the object wrapped by the webrpc server.
type AddServiceWeb interface {
	Add(r *http.Request, in *AddMessage, out *SumMessage) error
}

// internal wrapper for type-safe webrpc calling
type rpcAddServiceWebClient struct {
	remote   *url.URL
	protocol webrpc.Protocol
}

func (this rpcAddServiceWebClient) Add(in *AddMessage, out *SumMessage) error {
	return webrpc.Post(this.protocol, this.remote, "/AddService/Add", in, out)
}

// Register a AddServiceWeb implementation with the given webrpc ServeMux.
// If mux is nil, the default webrpc.ServeMux is used.
func RegisterAddServiceWeb(this AddServiceWeb, mux webrpc.ServeMux) error {
	if mux == nil {
		mux = webrpc.DefaultServeMux
	}
	if err := mux.Handle("/AddService/Add", func(c *webrpc.Call) error {
		in, out := new(AddMessage), new(SumMessage)
		if err := c.ReadRequest(in); err != nil {
			return err
		}
		if err := this.Add(c.Request, in, out); err != nil {
			return err
		}
		return c.WriteResponse(out)
	}); err != nil {
		return err
	}
	return nil
}

// NewAddServiceWebClient returns a webrpc wrapper for calling the methods of AddService
// remotely via the web.  The remote URL is the base URL of the webrpc server.
func NewAddServiceWebClient(pro webrpc.Protocol, remote *url.URL) AddService {
	return rpcAddServiceWebClient{remote, pro}
}
