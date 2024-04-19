// Code generated by protoc-gen-go-lite. DO NOT EDIT.
// protoc-gen-go-lite version: v0.1.2
// source: github.com/aperturerobotics/protobuf-go-lite/testproto/proto3opt/opt.proto

package proto3opt

import (
	base64 "encoding/base64"
	binary "encoding/binary"
	io "io"
	math "math"
	strconv "strconv"
	unsafe "unsafe"

	v2 "github.com/Jeffail/gabs/v2"
	protohelpers "github.com/aperturerobotics/protobuf-go-lite/protohelpers"
	errors "github.com/pkg/errors"
	fastjson "github.com/valyala/fastjson"
)

type SimpleEnum int32

const (
	SimpleEnum_ENUM1 SimpleEnum = 0
	SimpleEnum_ENUM2 SimpleEnum = 1
	SimpleEnum_ENUM3 SimpleEnum = 2
)

// Enum value maps for SimpleEnum.
var (
	SimpleEnum_name = map[int32]string{
		0: "ENUM1",
		1: "ENUM2",
		2: "ENUM3",
	}
	SimpleEnum_value = map[string]int32{
		"ENUM1": 0,
		"ENUM2": 1,
		"ENUM3": 2,
	}
)

func (x SimpleEnum) Enum() *SimpleEnum {
	p := new(SimpleEnum)
	*p = x
	return p
}

func (x SimpleEnum) String() string {
	name, valid := SimpleEnum_name[int32(x)]
	if valid {
		return name
	}
	return strconv.Itoa(int(x))
}

type OptionalFieldInProto3 struct {
	unknownFields    []byte
	OptionalInt32    *int32      `protobuf:"varint,1,opt,name=optional_int32,json=optionalInt32,proto3,oneof" json:"optional_int32,omitempty"`
	OptionalInt64    *int64      `protobuf:"varint,2,opt,name=optional_int64,json=optionalInt64,proto3,oneof" json:"optional_int64,omitempty"`
	OptionalUint32   *uint32     `protobuf:"varint,3,opt,name=optional_uint32,json=optionalUint32,proto3,oneof" json:"optional_uint32,omitempty"`
	OptionalUint64   *uint64     `protobuf:"varint,4,opt,name=optional_uint64,json=optionalUint64,proto3,oneof" json:"optional_uint64,omitempty"`
	OptionalSint32   *int32      `protobuf:"zigzag32,5,opt,name=optional_sint32,json=optionalSint32,proto3,oneof" json:"optional_sint32,omitempty"`
	OptionalSint64   *int64      `protobuf:"zigzag64,6,opt,name=optional_sint64,json=optionalSint64,proto3,oneof" json:"optional_sint64,omitempty"`
	OptionalFixed32  *uint32     `protobuf:"fixed32,7,opt,name=optional_fixed32,json=optionalFixed32,proto3,oneof" json:"optional_fixed32,omitempty"`
	OptionalFixed64  *uint64     `protobuf:"fixed64,8,opt,name=optional_fixed64,json=optionalFixed64,proto3,oneof" json:"optional_fixed64,omitempty"`
	OptionalSfixed32 *int32      `protobuf:"fixed32,9,opt,name=optional_sfixed32,json=optionalSfixed32,proto3,oneof" json:"optional_sfixed32,omitempty"`
	OptionalSfixed64 *int64      `protobuf:"fixed64,10,opt,name=optional_sfixed64,json=optionalSfixed64,proto3,oneof" json:"optional_sfixed64,omitempty"`
	OptionalFloat    *float32    `protobuf:"fixed32,11,opt,name=optional_float,json=optionalFloat,proto3,oneof" json:"optional_float,omitempty"`
	OptionalDouble   *float64    `protobuf:"fixed64,12,opt,name=optional_double,json=optionalDouble,proto3,oneof" json:"optional_double,omitempty"`
	OptionalBool     *bool       `protobuf:"varint,13,opt,name=optional_bool,json=optionalBool,proto3,oneof" json:"optional_bool,omitempty"`
	OptionalString   *string     `protobuf:"bytes,14,opt,name=optional_string,json=optionalString,proto3,oneof" json:"optional_string,omitempty"`
	OptionalBytes    []byte      `protobuf:"bytes,15,opt,name=optional_bytes,json=optionalBytes,proto3,oneof" json:"optional_bytes,omitempty"`
	OptionalEnum     *SimpleEnum `protobuf:"varint,16,opt,name=optional_enum,json=optionalEnum,proto3,oneof" json:"optional_enum,omitempty"`
}

func (x *OptionalFieldInProto3) Reset() {
	*x = OptionalFieldInProto3{}
}

func (*OptionalFieldInProto3) ProtoMessage() {}

func (x *OptionalFieldInProto3) GetOptionalInt32() int32 {
	if x != nil && x.OptionalInt32 != nil {
		return *x.OptionalInt32
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalInt64() int64 {
	if x != nil && x.OptionalInt64 != nil {
		return *x.OptionalInt64
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalUint32() uint32 {
	if x != nil && x.OptionalUint32 != nil {
		return *x.OptionalUint32
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalUint64() uint64 {
	if x != nil && x.OptionalUint64 != nil {
		return *x.OptionalUint64
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalSint32() int32 {
	if x != nil && x.OptionalSint32 != nil {
		return *x.OptionalSint32
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalSint64() int64 {
	if x != nil && x.OptionalSint64 != nil {
		return *x.OptionalSint64
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalFixed32() uint32 {
	if x != nil && x.OptionalFixed32 != nil {
		return *x.OptionalFixed32
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalFixed64() uint64 {
	if x != nil && x.OptionalFixed64 != nil {
		return *x.OptionalFixed64
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalSfixed32() int32 {
	if x != nil && x.OptionalSfixed32 != nil {
		return *x.OptionalSfixed32
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalSfixed64() int64 {
	if x != nil && x.OptionalSfixed64 != nil {
		return *x.OptionalSfixed64
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalFloat() float32 {
	if x != nil && x.OptionalFloat != nil {
		return *x.OptionalFloat
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalDouble() float64 {
	if x != nil && x.OptionalDouble != nil {
		return *x.OptionalDouble
	}
	return 0
}

func (x *OptionalFieldInProto3) GetOptionalBool() bool {
	if x != nil && x.OptionalBool != nil {
		return *x.OptionalBool
	}
	return false
}

func (x *OptionalFieldInProto3) GetOptionalString() string {
	if x != nil && x.OptionalString != nil {
		return *x.OptionalString
	}
	return ""
}

func (x *OptionalFieldInProto3) GetOptionalBytes() []byte {
	if x != nil {
		return x.OptionalBytes
	}
	return nil
}

func (x *OptionalFieldInProto3) GetOptionalEnum() SimpleEnum {
	if x != nil && x.OptionalEnum != nil {
		return *x.OptionalEnum
	}
	return SimpleEnum_ENUM1
}

func (m *OptionalFieldInProto3) CloneVT() *OptionalFieldInProto3 {
	if m == nil {
		return (*OptionalFieldInProto3)(nil)
	}
	r := new(OptionalFieldInProto3)
	if rhs := m.OptionalInt32; rhs != nil {
		tmpVal := *rhs
		r.OptionalInt32 = &tmpVal
	}
	if rhs := m.OptionalInt64; rhs != nil {
		tmpVal := *rhs
		r.OptionalInt64 = &tmpVal
	}
	if rhs := m.OptionalUint32; rhs != nil {
		tmpVal := *rhs
		r.OptionalUint32 = &tmpVal
	}
	if rhs := m.OptionalUint64; rhs != nil {
		tmpVal := *rhs
		r.OptionalUint64 = &tmpVal
	}
	if rhs := m.OptionalSint32; rhs != nil {
		tmpVal := *rhs
		r.OptionalSint32 = &tmpVal
	}
	if rhs := m.OptionalSint64; rhs != nil {
		tmpVal := *rhs
		r.OptionalSint64 = &tmpVal
	}
	if rhs := m.OptionalFixed32; rhs != nil {
		tmpVal := *rhs
		r.OptionalFixed32 = &tmpVal
	}
	if rhs := m.OptionalFixed64; rhs != nil {
		tmpVal := *rhs
		r.OptionalFixed64 = &tmpVal
	}
	if rhs := m.OptionalSfixed32; rhs != nil {
		tmpVal := *rhs
		r.OptionalSfixed32 = &tmpVal
	}
	if rhs := m.OptionalSfixed64; rhs != nil {
		tmpVal := *rhs
		r.OptionalSfixed64 = &tmpVal
	}
	if rhs := m.OptionalFloat; rhs != nil {
		tmpVal := *rhs
		r.OptionalFloat = &tmpVal
	}
	if rhs := m.OptionalDouble; rhs != nil {
		tmpVal := *rhs
		r.OptionalDouble = &tmpVal
	}
	if rhs := m.OptionalBool; rhs != nil {
		tmpVal := *rhs
		r.OptionalBool = &tmpVal
	}
	if rhs := m.OptionalString; rhs != nil {
		tmpVal := *rhs
		r.OptionalString = &tmpVal
	}
	if rhs := m.OptionalBytes; rhs != nil {
		tmpBytes := make([]byte, len(rhs))
		copy(tmpBytes, rhs)
		r.OptionalBytes = tmpBytes
	}
	if rhs := m.OptionalEnum; rhs != nil {
		tmpVal := *rhs
		r.OptionalEnum = &tmpVal
	}
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *OptionalFieldInProto3) CloneMessageVT() any {
	return m.CloneVT()
}

func (this *OptionalFieldInProto3) EqualVT(that *OptionalFieldInProto3) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	if p, q := this.OptionalInt32, that.OptionalInt32; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalInt64, that.OptionalInt64; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalUint32, that.OptionalUint32; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalUint64, that.OptionalUint64; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalSint32, that.OptionalSint32; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalSint64, that.OptionalSint64; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalFixed32, that.OptionalFixed32; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalFixed64, that.OptionalFixed64; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalSfixed32, that.OptionalSfixed32; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalSfixed64, that.OptionalSfixed64; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalFloat, that.OptionalFloat; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalDouble, that.OptionalDouble; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalBool, that.OptionalBool; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalString, that.OptionalString; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	if p, q := this.OptionalBytes, that.OptionalBytes; (p == nil && q != nil) || (p != nil && q == nil) || string(p) != string(q) {
		return false
	}
	if p, q := this.OptionalEnum, that.OptionalEnum; (p == nil && q != nil) || (p != nil && (q == nil || *p != *q)) {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *OptionalFieldInProto3) EqualMessageVT(thatMsg any) bool {
	that, ok := thatMsg.(*OptionalFieldInProto3)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (m *OptionalFieldInProto3) MarshalJSON() ([]byte, error) {
	container := v2.New()
	if m.OptionalInt32 != nil {
		container.Set(m.OptionalInt32, "optionalInt32")
	}
	if m.OptionalInt64 != nil {
		container.Set(m.OptionalInt64, "optionalInt64")
	}
	if m.OptionalUint32 != nil {
		container.Set(m.OptionalUint32, "optionalUint32")
	}
	if m.OptionalUint64 != nil {
		container.Set(m.OptionalUint64, "optionalUint64")
	}
	if m.OptionalSint32 != nil {
		// Unsupported type sint32
	}
	if m.OptionalSint64 != nil {
		// Unsupported type sint64
	}
	if m.OptionalFixed32 != nil {
		// Unsupported type fixed32
	}
	if m.OptionalFixed64 != nil {
		// Unsupported type fixed64
	}
	if m.OptionalSfixed32 != nil {
		// Unsupported type sfixed32
	}
	if m.OptionalSfixed64 != nil {
		// Unsupported type sfixed64
	}
	if m.OptionalFloat != nil {
		container.Set(m.OptionalFloat, "optionalFloat")
	}
	if m.OptionalDouble != nil {
		container.Set(m.OptionalDouble, "optionalDouble")
	}
	if m.OptionalBool != nil {
		container.Set(m.OptionalBool, "optionalBool")
	}
	if m.OptionalString != nil {
		container.Set(m.OptionalString, "optionalString")
	}
	if m.OptionalBytes != nil {
		container.Set(base64.StdEncoding.EncodeToString(m.OptionalBytes), "optionalBytes")
	}
	if m.OptionalEnum != nil {
		container.Set(m.OptionalEnum.String(), "optionalEnum")
	}
	return container.MarshalJSON()
}

func (m *OptionalFieldInProto3) UnmarshalJSON(data []byte) error {
	var p fastjson.Parser
	v, err := p.ParseBytes(data)
	if err != nil {
		return err
	}
	return m.UnmarshalJSONValue(v)
}

func (m *OptionalFieldInProto3) UnmarshalJSONValue(v *fastjson.Value) error {
	if v == nil {
		return nil
	}
	if v.Exists("optionalInt32") {
		m.OptionalInt32 = int32(v.GetInt("optionalInt32"))
	} else if v.Exists("optional_int32") {
		m.OptionalInt32 = int32(v.GetInt("optional_int32"))
	}
	if v.Exists("optionalInt64") {
		m.OptionalInt64 = v.GetInt64("optionalInt64")
	} else if v.Exists("optional_int64") {
		m.OptionalInt64 = v.GetInt64("optional_int64")
	}
	if v.Exists("optionalUint32") {
		m.OptionalUint32 = v.GetUint64("optionalUint32")
	} else if v.Exists("optional_uint32") {
		m.OptionalUint32 = v.GetUint64("optional_uint32")
	}
	if v.Exists("optionalUint64") {
		m.OptionalUint64 = v.GetUint64("optionalUint64")
	} else if v.Exists("optional_uint64") {
		m.OptionalUint64 = v.GetUint64("optional_uint64")
	}
	if v.Exists("optionalSint32") {
		// Unsupported type sint32
	} else if v.Exists("optional_sint32") {
		// Unsupported type sint32
	}
	if v.Exists("optionalSint64") {
		// Unsupported type sint64
	} else if v.Exists("optional_sint64") {
		// Unsupported type sint64
	}
	if v.Exists("optionalFixed32") {
		// Unsupported type fixed32
	} else if v.Exists("optional_fixed32") {
		// Unsupported type fixed32
	}
	if v.Exists("optionalFixed64") {
		// Unsupported type fixed64
	} else if v.Exists("optional_fixed64") {
		// Unsupported type fixed64
	}
	if v.Exists("optionalSfixed32") {
		// Unsupported type sfixed32
	} else if v.Exists("optional_sfixed32") {
		// Unsupported type sfixed32
	}
	if v.Exists("optionalSfixed64") {
		// Unsupported type sfixed64
	} else if v.Exists("optional_sfixed64") {
		// Unsupported type sfixed64
	}
	if v.Exists("optionalFloat") {
		m.OptionalFloat = float32(v.GetFloat64("optionalFloat"))
	} else if v.Exists("optional_float") {
		m.OptionalFloat = float32(v.GetFloat64("optional_float"))
	}
	if v.Exists("optionalDouble") {
		m.OptionalDouble = v.GetFloat64("optionalDouble")
	} else if v.Exists("optional_double") {
		m.OptionalDouble = v.GetFloat64("optional_double")
	}
	if v.Exists("optionalBool") {
		m.OptionalBool = v.GetBool("optionalBool")
	} else if v.Exists("optional_bool") {
		m.OptionalBool = v.GetBool("optional_bool")
	}
	if v.Exists("optionalString") {
		m.OptionalString = string(v.GetStringBytes("optionalString"))
	} else if v.Exists("optional_string") {
		m.OptionalString = string(v.GetStringBytes("optional_string"))
	}
	if v.Exists("optionalBytes") {
		jsonBytes := v.GetStringBytes("optionalBytes")
		var err error
		m.OptionalBytes, err = base64.StdEncoding.DecodeString(string(jsonBytes))
		if err != nil {
			return err
		}
	} else if v.Exists("optional_bytes") {
		jsonBytes := v.GetStringBytes("optional_bytes")
		var err error
		m.OptionalBytes, err = base64.StdEncoding.DecodeString(string(jsonBytes))
		if err != nil {
			return err
		}
	}
	if v.Exists("optionalEnum") {
		m.OptionalEnum = SimpleEnum(v.GetInt("optionalEnum"))
	} else if v.Exists("optional_enum") {
		m.OptionalEnum = SimpleEnum(v.GetInt("optional_enum"))
	}
	return nil
}

func (m *OptionalFieldInProto3) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OptionalFieldInProto3) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *OptionalFieldInProto3) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.OptionalEnum != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalEnum))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.OptionalBytes != nil {
		i -= len(m.OptionalBytes)
		copy(dAtA[i:], m.OptionalBytes)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.OptionalBytes)))
		i--
		dAtA[i] = 0x7a
	}
	if m.OptionalString != nil {
		i -= len(*m.OptionalString)
		copy(dAtA[i:], *m.OptionalString)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(*m.OptionalString)))
		i--
		dAtA[i] = 0x72
	}
	if m.OptionalBool != nil {
		i--
		if *m.OptionalBool {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x68
	}
	if m.OptionalDouble != nil {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.OptionalDouble))))
		i--
		dAtA[i] = 0x61
	}
	if m.OptionalFloat != nil {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(*m.OptionalFloat))))
		i--
		dAtA[i] = 0x5d
	}
	if m.OptionalSfixed64 != nil {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(*m.OptionalSfixed64))
		i--
		dAtA[i] = 0x51
	}
	if m.OptionalSfixed32 != nil {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(*m.OptionalSfixed32))
		i--
		dAtA[i] = 0x4d
	}
	if m.OptionalFixed64 != nil {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(*m.OptionalFixed64))
		i--
		dAtA[i] = 0x41
	}
	if m.OptionalFixed32 != nil {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(*m.OptionalFixed32))
		i--
		dAtA[i] = 0x3d
	}
	if m.OptionalSint64 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64((uint64(*m.OptionalSint64)<<1)^uint64((*m.OptionalSint64>>63))))
		i--
		dAtA[i] = 0x30
	}
	if m.OptionalSint32 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64((uint32(*m.OptionalSint32)<<1)^uint32((*m.OptionalSint32>>31))))
		i--
		dAtA[i] = 0x28
	}
	if m.OptionalUint64 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalUint64))
		i--
		dAtA[i] = 0x20
	}
	if m.OptionalUint32 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalUint32))
		i--
		dAtA[i] = 0x18
	}
	if m.OptionalInt64 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalInt64))
		i--
		dAtA[i] = 0x10
	}
	if m.OptionalInt32 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalInt32))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OptionalFieldInProto3) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OptionalFieldInProto3) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *OptionalFieldInProto3) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	if m.OptionalEnum != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalEnum))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.OptionalBytes != nil {
		i -= len(m.OptionalBytes)
		copy(dAtA[i:], m.OptionalBytes)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(m.OptionalBytes)))
		i--
		dAtA[i] = 0x7a
	}
	if m.OptionalString != nil {
		i -= len(*m.OptionalString)
		copy(dAtA[i:], *m.OptionalString)
		i = protohelpers.EncodeVarint(dAtA, i, uint64(len(*m.OptionalString)))
		i--
		dAtA[i] = 0x72
	}
	if m.OptionalBool != nil {
		i--
		if *m.OptionalBool {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x68
	}
	if m.OptionalDouble != nil {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(*m.OptionalDouble))))
		i--
		dAtA[i] = 0x61
	}
	if m.OptionalFloat != nil {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(math.Float32bits(float32(*m.OptionalFloat))))
		i--
		dAtA[i] = 0x5d
	}
	if m.OptionalSfixed64 != nil {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(*m.OptionalSfixed64))
		i--
		dAtA[i] = 0x51
	}
	if m.OptionalSfixed32 != nil {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(*m.OptionalSfixed32))
		i--
		dAtA[i] = 0x4d
	}
	if m.OptionalFixed64 != nil {
		i -= 8
		binary.LittleEndian.PutUint64(dAtA[i:], uint64(*m.OptionalFixed64))
		i--
		dAtA[i] = 0x41
	}
	if m.OptionalFixed32 != nil {
		i -= 4
		binary.LittleEndian.PutUint32(dAtA[i:], uint32(*m.OptionalFixed32))
		i--
		dAtA[i] = 0x3d
	}
	if m.OptionalSint64 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64((uint64(*m.OptionalSint64)<<1)^uint64((*m.OptionalSint64>>63))))
		i--
		dAtA[i] = 0x30
	}
	if m.OptionalSint32 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64((uint32(*m.OptionalSint32)<<1)^uint32((*m.OptionalSint32>>31))))
		i--
		dAtA[i] = 0x28
	}
	if m.OptionalUint64 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalUint64))
		i--
		dAtA[i] = 0x20
	}
	if m.OptionalUint32 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalUint32))
		i--
		dAtA[i] = 0x18
	}
	if m.OptionalInt64 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalInt64))
		i--
		dAtA[i] = 0x10
	}
	if m.OptionalInt32 != nil {
		i = protohelpers.EncodeVarint(dAtA, i, uint64(*m.OptionalInt32))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *OptionalFieldInProto3) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OptionalInt32 != nil {
		n += 1 + protohelpers.SizeOfVarint(uint64(*m.OptionalInt32))
	}
	if m.OptionalInt64 != nil {
		n += 1 + protohelpers.SizeOfVarint(uint64(*m.OptionalInt64))
	}
	if m.OptionalUint32 != nil {
		n += 1 + protohelpers.SizeOfVarint(uint64(*m.OptionalUint32))
	}
	if m.OptionalUint64 != nil {
		n += 1 + protohelpers.SizeOfVarint(uint64(*m.OptionalUint64))
	}
	if m.OptionalSint32 != nil {
		n += 1 + protohelpers.SizeOfZigzag(uint64(*m.OptionalSint32))
	}
	if m.OptionalSint64 != nil {
		n += 1 + protohelpers.SizeOfZigzag(uint64(*m.OptionalSint64))
	}
	if m.OptionalFixed32 != nil {
		n += 5
	}
	if m.OptionalFixed64 != nil {
		n += 9
	}
	if m.OptionalSfixed32 != nil {
		n += 5
	}
	if m.OptionalSfixed64 != nil {
		n += 9
	}
	if m.OptionalFloat != nil {
		n += 5
	}
	if m.OptionalDouble != nil {
		n += 9
	}
	if m.OptionalBool != nil {
		n += 2
	}
	if m.OptionalString != nil {
		l = len(*m.OptionalString)
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.OptionalBytes != nil {
		l = len(m.OptionalBytes)
		n += 1 + l + protohelpers.SizeOfVarint(uint64(l))
	}
	if m.OptionalEnum != nil {
		n += 2 + protohelpers.SizeOfVarint(uint64(*m.OptionalEnum))
	}
	n += len(m.unknownFields)
	return n
}

func (m *OptionalFieldInProto3) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return errors.Errorf("proto: OptionalFieldInProto3: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return errors.Errorf("proto: OptionalFieldInProto3: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalInt32", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalInt32 = &v
		case 2:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalInt64", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalInt64 = &v
		case 3:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalUint32", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalUint32 = &v
		case 4:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalUint64", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalUint64 = &v
		case 5:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalSint32", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.OptionalSint32 = &v
		case 6:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalSint64", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			v2 := int64(v)
			m.OptionalSint64 = &v2
		case 7:
			if wireType != 5 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalFixed32", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.OptionalFixed32 = &v
		case 8:
			if wireType != 1 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalFixed64", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.OptionalFixed64 = &v
		case 9:
			if wireType != 5 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalSfixed32", wireType)
			}
			var v int32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = int32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.OptionalSfixed32 = &v
		case 10:
			if wireType != 1 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalSfixed64", wireType)
			}
			var v int64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.OptionalSfixed64 = &v
		case 11:
			if wireType != 5 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalFloat", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			v2 := float32(math.Float32frombits(v))
			m.OptionalFloat = &v2
		case 12:
			if wireType != 1 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalDouble", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.OptionalDouble = &v2
		case 13:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalBool", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.OptionalBool = &b
		case 14:
			if wireType != 2 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.OptionalString = &s
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OptionalBytes = append(m.OptionalBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.OptionalBytes == nil {
				m.OptionalBytes = []byte{}
			}
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalEnum", wireType)
			}
			var v SimpleEnum
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= SimpleEnum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalEnum = &v
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OptionalFieldInProto3) UnmarshalVTUnsafe(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return errors.Errorf("proto: OptionalFieldInProto3: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return errors.Errorf("proto: OptionalFieldInProto3: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalInt32", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalInt32 = &v
		case 2:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalInt64", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalInt64 = &v
		case 3:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalUint32", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalUint32 = &v
		case 4:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalUint64", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalUint64 = &v
		case 5:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalSint32", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = int32((uint32(v) >> 1) ^ uint32(((v&1)<<31)>>31))
			m.OptionalSint32 = &v
		case 6:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalSint64", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			v = (v >> 1) ^ uint64((int64(v&1)<<63)>>63)
			v2 := int64(v)
			m.OptionalSint64 = &v2
		case 7:
			if wireType != 5 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalFixed32", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.OptionalFixed32 = &v
		case 8:
			if wireType != 1 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalFixed64", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.OptionalFixed64 = &v
		case 9:
			if wireType != 5 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalSfixed32", wireType)
			}
			var v int32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = int32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			m.OptionalSfixed32 = &v
		case 10:
			if wireType != 1 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalSfixed64", wireType)
			}
			var v int64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = int64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.OptionalSfixed64 = &v
		case 11:
			if wireType != 5 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalFloat", wireType)
			}
			var v uint32
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint32(binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
			v2 := float32(math.Float32frombits(v))
			m.OptionalFloat = &v2
		case 12:
			if wireType != 1 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalDouble", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			v2 := float64(math.Float64frombits(v))
			m.OptionalDouble = &v2
		case 13:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalBool", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.OptionalBool = &b
		case 14:
			if wireType != 2 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var stringValue string
			if intStringLen > 0 {
				stringValue = unsafe.String(&dAtA[iNdEx], intStringLen)
			}
			s := stringValue
			m.OptionalString = &s
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return protohelpers.ErrInvalidLength
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return protohelpers.ErrInvalidLength
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OptionalBytes = dAtA[iNdEx:postIndex]
			iNdEx = postIndex
		case 16:
			if wireType != 0 {
				return errors.Errorf("proto: wrong wireType = %d for field OptionalEnum", wireType)
			}
			var v SimpleEnum
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protohelpers.ErrIntOverflow
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= SimpleEnum(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptionalEnum = &v
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
