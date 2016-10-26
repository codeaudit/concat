// Code generated by protoc-gen-gogo.
// source: stmt.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Statement struct {
	Id        string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Publisher string         `protobuf:"bytes,2,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Namespace string         `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Body      *StatementBody `protobuf:"bytes,4,opt,name=body" json:"body,omitempty"`
	Timestamp int64          `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Signature []byte         `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Statement) Reset()                    { *m = Statement{} }
func (m *Statement) String() string            { return proto1.CompactTextString(m) }
func (*Statement) ProtoMessage()               {}
func (*Statement) Descriptor() ([]byte, []int) { return fileDescriptorStmt, []int{0} }

func (m *Statement) GetBody() *StatementBody {
	if m != nil {
		return m.Body
	}
	return nil
}

type StatementBody struct {
	// Types that are valid to be assigned to Body:
	//	*StatementBody_Simple
	//	*StatementBody_Compound
	//	*StatementBody_Envelope
	//	*StatementBody_Archive
	Body isStatementBody_Body `protobuf_oneof:"body"`
}

func (m *StatementBody) Reset()                    { *m = StatementBody{} }
func (m *StatementBody) String() string            { return proto1.CompactTextString(m) }
func (*StatementBody) ProtoMessage()               {}
func (*StatementBody) Descriptor() ([]byte, []int) { return fileDescriptorStmt, []int{1} }

type isStatementBody_Body interface {
	isStatementBody_Body()
}

type StatementBody_Simple struct {
	Simple *SimpleStatement `protobuf:"bytes,1,opt,name=simple,oneof"`
}
type StatementBody_Compound struct {
	Compound *CompoundStatement `protobuf:"bytes,2,opt,name=compound,oneof"`
}
type StatementBody_Envelope struct {
	Envelope *EnvelopeStatement `protobuf:"bytes,3,opt,name=envelope,oneof"`
}
type StatementBody_Archive struct {
	Archive *ArchiveStatement `protobuf:"bytes,4,opt,name=archive,oneof"`
}

func (*StatementBody_Simple) isStatementBody_Body()   {}
func (*StatementBody_Compound) isStatementBody_Body() {}
func (*StatementBody_Envelope) isStatementBody_Body() {}
func (*StatementBody_Archive) isStatementBody_Body()  {}

func (m *StatementBody) GetBody() isStatementBody_Body {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *StatementBody) GetSimple() *SimpleStatement {
	if x, ok := m.GetBody().(*StatementBody_Simple); ok {
		return x.Simple
	}
	return nil
}

func (m *StatementBody) GetCompound() *CompoundStatement {
	if x, ok := m.GetBody().(*StatementBody_Compound); ok {
		return x.Compound
	}
	return nil
}

func (m *StatementBody) GetEnvelope() *EnvelopeStatement {
	if x, ok := m.GetBody().(*StatementBody_Envelope); ok {
		return x.Envelope
	}
	return nil
}

func (m *StatementBody) GetArchive() *ArchiveStatement {
	if x, ok := m.GetBody().(*StatementBody_Archive); ok {
		return x.Archive
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*StatementBody) XXX_OneofFuncs() (func(msg proto1.Message, b *proto1.Buffer) error, func(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error), func(msg proto1.Message) (n int), []interface{}) {
	return _StatementBody_OneofMarshaler, _StatementBody_OneofUnmarshaler, _StatementBody_OneofSizer, []interface{}{
		(*StatementBody_Simple)(nil),
		(*StatementBody_Compound)(nil),
		(*StatementBody_Envelope)(nil),
		(*StatementBody_Archive)(nil),
	}
}

func _StatementBody_OneofMarshaler(msg proto1.Message, b *proto1.Buffer) error {
	m := msg.(*StatementBody)
	// body
	switch x := m.Body.(type) {
	case *StatementBody_Simple:
		_ = b.EncodeVarint(1<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Simple); err != nil {
			return err
		}
	case *StatementBody_Compound:
		_ = b.EncodeVarint(2<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Compound); err != nil {
			return err
		}
	case *StatementBody_Envelope:
		_ = b.EncodeVarint(3<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Envelope); err != nil {
			return err
		}
	case *StatementBody_Archive:
		_ = b.EncodeVarint(4<<3 | proto1.WireBytes)
		if err := b.EncodeMessage(x.Archive); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("StatementBody.Body has unexpected type %T", x)
	}
	return nil
}

func _StatementBody_OneofUnmarshaler(msg proto1.Message, tag, wire int, b *proto1.Buffer) (bool, error) {
	m := msg.(*StatementBody)
	switch tag {
	case 1: // body.simple
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(SimpleStatement)
		err := b.DecodeMessage(msg)
		m.Body = &StatementBody_Simple{msg}
		return true, err
	case 2: // body.compound
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(CompoundStatement)
		err := b.DecodeMessage(msg)
		m.Body = &StatementBody_Compound{msg}
		return true, err
	case 3: // body.envelope
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(EnvelopeStatement)
		err := b.DecodeMessage(msg)
		m.Body = &StatementBody_Envelope{msg}
		return true, err
	case 4: // body.archive
		if wire != proto1.WireBytes {
			return true, proto1.ErrInternalBadWireType
		}
		msg := new(ArchiveStatement)
		err := b.DecodeMessage(msg)
		m.Body = &StatementBody_Archive{msg}
		return true, err
	default:
		return false, nil
	}
}

func _StatementBody_OneofSizer(msg proto1.Message) (n int) {
	m := msg.(*StatementBody)
	// body
	switch x := m.Body.(type) {
	case *StatementBody_Simple:
		s := proto1.Size(x.Simple)
		n += proto1.SizeVarint(1<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *StatementBody_Compound:
		s := proto1.Size(x.Compound)
		n += proto1.SizeVarint(2<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *StatementBody_Envelope:
		s := proto1.Size(x.Envelope)
		n += proto1.SizeVarint(3<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case *StatementBody_Archive:
		s := proto1.Size(x.Archive)
		n += proto1.SizeVarint(4<<3 | proto1.WireBytes)
		n += proto1.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type SimpleStatement struct {
	Object string   `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	Refs   []string `protobuf:"bytes,2,rep,name=refs" json:"refs,omitempty"`
	Tags   []string `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty"`
	Deps   []string `protobuf:"bytes,4,rep,name=deps" json:"deps,omitempty"`
}

func (m *SimpleStatement) Reset()                    { *m = SimpleStatement{} }
func (m *SimpleStatement) String() string            { return proto1.CompactTextString(m) }
func (*SimpleStatement) ProtoMessage()               {}
func (*SimpleStatement) Descriptor() ([]byte, []int) { return fileDescriptorStmt, []int{2} }

type CompoundStatement struct {
	Body []*SimpleStatement `protobuf:"bytes,1,rep,name=body" json:"body,omitempty"`
}

func (m *CompoundStatement) Reset()                    { *m = CompoundStatement{} }
func (m *CompoundStatement) String() string            { return proto1.CompactTextString(m) }
func (*CompoundStatement) ProtoMessage()               {}
func (*CompoundStatement) Descriptor() ([]byte, []int) { return fileDescriptorStmt, []int{3} }

func (m *CompoundStatement) GetBody() []*SimpleStatement {
	if m != nil {
		return m.Body
	}
	return nil
}

type EnvelopeStatement struct {
	Body []*Statement `protobuf:"bytes,1,rep,name=body" json:"body,omitempty"`
}

func (m *EnvelopeStatement) Reset()                    { *m = EnvelopeStatement{} }
func (m *EnvelopeStatement) String() string            { return proto1.CompactTextString(m) }
func (*EnvelopeStatement) ProtoMessage()               {}
func (*EnvelopeStatement) Descriptor() ([]byte, []int) { return fileDescriptorStmt, []int{4} }

func (m *EnvelopeStatement) GetBody() []*Statement {
	if m != nil {
		return m.Body
	}
	return nil
}

type ArchiveStatement struct {
}

func (m *ArchiveStatement) Reset()                    { *m = ArchiveStatement{} }
func (m *ArchiveStatement) String() string            { return proto1.CompactTextString(m) }
func (*ArchiveStatement) ProtoMessage()               {}
func (*ArchiveStatement) Descriptor() ([]byte, []int) { return fileDescriptorStmt, []int{5} }

func init() {
	proto1.RegisterType((*Statement)(nil), "proto.Statement")
	proto1.RegisterType((*StatementBody)(nil), "proto.StatementBody")
	proto1.RegisterType((*SimpleStatement)(nil), "proto.SimpleStatement")
	proto1.RegisterType((*CompoundStatement)(nil), "proto.CompoundStatement")
	proto1.RegisterType((*EnvelopeStatement)(nil), "proto.EnvelopeStatement")
	proto1.RegisterType((*ArchiveStatement)(nil), "proto.ArchiveStatement")
}

func init() { proto1.RegisterFile("stmt.proto", fileDescriptorStmt) }

var fileDescriptorStmt = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x6a, 0xdb, 0x40,
	0x10, 0xc7, 0xbb, 0x92, 0xac, 0x56, 0xe3, 0x7e, 0xd8, 0x4b, 0x71, 0xf7, 0xd0, 0x83, 0x10, 0x3d,
	0x88, 0x1e, 0x4c, 0xb1, 0xa1, 0x90, 0x53, 0x88, 0x43, 0x20, 0x67, 0xe5, 0x09, 0xf4, 0x31, 0xb1,
	0x15, 0x2c, 0xed, 0xa2, 0x5d, 0x1b, 0xfc, 0x70, 0x79, 0xa5, 0x3c, 0x43, 0xd8, 0x0f, 0x4b, 0xfe,
	0x20, 0x27, 0x8d, 0x7e, 0xf3, 0xff, 0xcf, 0x30, 0x33, 0x0b, 0x20, 0x55, 0xa3, 0xe6, 0xa2, 0xe3,
	0x8a, 0xd3, 0x91, 0xf9, 0x24, 0xaf, 0x04, 0xa2, 0x27, 0x95, 0x2b, 0x6c, 0xb0, 0x55, 0xf4, 0x3b,
	0x78, 0x75, 0xc5, 0x48, 0x4c, 0xd2, 0x28, 0xf3, 0xea, 0x8a, 0xfe, 0x86, 0x48, 0xec, 0x8a, 0x6d,
	0x2d, 0x37, 0xd8, 0x31, 0xcf, 0xe0, 0x01, 0xe8, 0x6c, 0x9b, 0x37, 0x28, 0x45, 0x5e, 0x22, 0xf3,
	0x6d, 0xb6, 0x07, 0x34, 0x85, 0xa0, 0xe0, 0xd5, 0x81, 0x05, 0x31, 0x49, 0xc7, 0x8b, 0x9f, 0xb6,
	0xed, 0xbc, 0xef, 0xb5, 0xe2, 0xd5, 0x21, 0x33, 0x0a, 0x5d, 0x47, 0xd5, 0x0d, 0x4a, 0x95, 0x37,
	0x82, 0x8d, 0x62, 0x92, 0xfa, 0xd9, 0x00, 0x74, 0x56, 0xd6, 0xeb, 0x36, 0x57, 0xbb, 0x0e, 0x59,
	0x18, 0x93, 0xf4, 0x6b, 0x36, 0x80, 0xe4, 0x8d, 0xc0, 0xb7, 0xb3, 0x9a, 0xf4, 0x1f, 0x84, 0xb2,
	0x6e, 0xc4, 0x16, 0xcd, 0x1c, 0xe3, 0xc5, 0xec, 0xd8, 0xd9, 0xc0, 0x5e, 0xfb, 0xf8, 0x29, 0x73,
	0x3a, 0xfa, 0x1f, 0xbe, 0x94, 0xbc, 0x11, 0x7c, 0xd7, 0x56, 0x66, 0xc8, 0xf1, 0x82, 0x39, 0xcf,
	0xbd, 0xc3, 0xa7, 0xae, 0x5e, 0xab, 0x7d, 0xd8, 0xee, 0x71, 0xcb, 0x85, 0x1d, 0x7f, 0xf0, 0x3d,
	0x38, 0x7c, 0xe6, 0x3b, 0x6a, 0xe9, 0x12, 0x3e, 0xe7, 0x5d, 0xb9, 0xa9, 0xf7, 0xe8, 0x96, 0xf3,
	0xcb, 0xd9, 0xee, 0x2c, 0x3d, 0x75, 0x1d, 0x95, 0xab, 0xd0, 0xae, 0x33, 0x41, 0xf8, 0x71, 0x31,
	0x09, 0x9d, 0x41, 0xc8, 0x8b, 0x17, 0x2c, 0x95, 0xbb, 0x9c, 0xfb, 0xa3, 0x14, 0x82, 0x0e, 0x9f,
	0x25, 0xf3, 0x62, 0x3f, 0x8d, 0x32, 0x13, 0x6b, 0xa6, 0xf2, 0xb5, 0x64, 0xbe, 0x65, 0x3a, 0xd6,
	0xac, 0x42, 0x21, 0x59, 0x60, 0x99, 0x8e, 0x93, 0x5b, 0x98, 0x5e, 0x0d, 0x4f, 0xff, 0xba, 0x93,
	0x92, 0xd8, 0xff, 0x78, 0xb1, 0xf6, 0xa8, 0xc9, 0x0d, 0x4c, 0xaf, 0xb6, 0x40, 0xff, 0x9c, 0x15,
	0x98, 0x5c, 0xbe, 0x09, 0x67, 0xa5, 0x30, 0xb9, 0xdc, 0x44, 0x11, 0x1a, 0xe9, 0xf2, 0x3d, 0x00,
	0x00, 0xff, 0xff, 0xc4, 0xae, 0x82, 0x5f, 0xc3, 0x02, 0x00, 0x00,
}
