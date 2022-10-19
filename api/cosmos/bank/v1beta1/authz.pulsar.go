// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package bankv1beta1

import (
	v1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	_ "cosmossdk.io/api/cosmos/msg/v1/legacy_amino"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_SendAuthorization_1_list)(nil)

type _SendAuthorization_1_list struct {
	list *[]*v1beta1.Coin
}

func (x *_SendAuthorization_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_SendAuthorization_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_SendAuthorization_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	(*x.list)[i] = concreteValue
}

func (x *_SendAuthorization_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*v1beta1.Coin)
	*x.list = append(*x.list, concreteValue)
}

func (x *_SendAuthorization_1_list) AppendMutable() protoreflect.Value {
	v := new(v1beta1.Coin)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_SendAuthorization_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_SendAuthorization_1_list) NewElement() protoreflect.Value {
	v := new(v1beta1.Coin)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_SendAuthorization_1_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_SendAuthorization_2_list)(nil)

type _SendAuthorization_2_list struct {
	list *[]string
}

func (x *_SendAuthorization_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_SendAuthorization_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_SendAuthorization_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_SendAuthorization_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_SendAuthorization_2_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message SendAuthorization at list field AllowList as it is not of Message kind"))
}

func (x *_SendAuthorization_2_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_SendAuthorization_2_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_SendAuthorization_2_list) IsValid() bool {
	return x.list != nil
}

var (
	md_SendAuthorization             protoreflect.MessageDescriptor
	fd_SendAuthorization_spend_limit protoreflect.FieldDescriptor
	fd_SendAuthorization_allow_list  protoreflect.FieldDescriptor
)

func init() {
	file_cosmos_bank_v1beta1_authz_proto_init()
	md_SendAuthorization = File_cosmos_bank_v1beta1_authz_proto.Messages().ByName("SendAuthorization")
	fd_SendAuthorization_spend_limit = md_SendAuthorization.Fields().ByName("spend_limit")
	fd_SendAuthorization_allow_list = md_SendAuthorization.Fields().ByName("allow_list")
}

var _ protoreflect.Message = (*fastReflection_SendAuthorization)(nil)

type fastReflection_SendAuthorization SendAuthorization

func (x *SendAuthorization) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SendAuthorization)(x)
}

func (x *SendAuthorization) slowProtoReflect() protoreflect.Message {
	mi := &file_cosmos_bank_v1beta1_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SendAuthorization_messageType fastReflection_SendAuthorization_messageType
var _ protoreflect.MessageType = fastReflection_SendAuthorization_messageType{}

type fastReflection_SendAuthorization_messageType struct{}

func (x fastReflection_SendAuthorization_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SendAuthorization)(nil)
}
func (x fastReflection_SendAuthorization_messageType) New() protoreflect.Message {
	return new(fastReflection_SendAuthorization)
}
func (x fastReflection_SendAuthorization_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SendAuthorization
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SendAuthorization) Descriptor() protoreflect.MessageDescriptor {
	return md_SendAuthorization
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SendAuthorization) Type() protoreflect.MessageType {
	return _fastReflection_SendAuthorization_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SendAuthorization) New() protoreflect.Message {
	return new(fastReflection_SendAuthorization)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SendAuthorization) Interface() protoreflect.ProtoMessage {
	return (*SendAuthorization)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SendAuthorization) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.SpendLimit) != 0 {
		value := protoreflect.ValueOfList(&_SendAuthorization_1_list{list: &x.SpendLimit})
		if !f(fd_SendAuthorization_spend_limit, value) {
			return
		}
	}
	if len(x.AllowList) != 0 {
		value := protoreflect.ValueOfList(&_SendAuthorization_2_list{list: &x.AllowList})
		if !f(fd_SendAuthorization_allow_list, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_SendAuthorization) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "cosmos.bank.v1beta1.SendAuthorization.spend_limit":
		return len(x.SpendLimit) != 0
	case "cosmos.bank.v1beta1.SendAuthorization.allow_list":
		return len(x.AllowList) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.SendAuthorization"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.SendAuthorization does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SendAuthorization) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "cosmos.bank.v1beta1.SendAuthorization.spend_limit":
		x.SpendLimit = nil
	case "cosmos.bank.v1beta1.SendAuthorization.allow_list":
		x.AllowList = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.SendAuthorization"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.SendAuthorization does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SendAuthorization) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "cosmos.bank.v1beta1.SendAuthorization.spend_limit":
		if len(x.SpendLimit) == 0 {
			return protoreflect.ValueOfList(&_SendAuthorization_1_list{})
		}
		listValue := &_SendAuthorization_1_list{list: &x.SpendLimit}
		return protoreflect.ValueOfList(listValue)
	case "cosmos.bank.v1beta1.SendAuthorization.allow_list":
		if len(x.AllowList) == 0 {
			return protoreflect.ValueOfList(&_SendAuthorization_2_list{})
		}
		listValue := &_SendAuthorization_2_list{list: &x.AllowList}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.SendAuthorization"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.SendAuthorization does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SendAuthorization) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "cosmos.bank.v1beta1.SendAuthorization.spend_limit":
		lv := value.List()
		clv := lv.(*_SendAuthorization_1_list)
		x.SpendLimit = *clv.list
	case "cosmos.bank.v1beta1.SendAuthorization.allow_list":
		lv := value.List()
		clv := lv.(*_SendAuthorization_2_list)
		x.AllowList = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.SendAuthorization"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.SendAuthorization does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SendAuthorization) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.bank.v1beta1.SendAuthorization.spend_limit":
		if x.SpendLimit == nil {
			x.SpendLimit = []*v1beta1.Coin{}
		}
		value := &_SendAuthorization_1_list{list: &x.SpendLimit}
		return protoreflect.ValueOfList(value)
	case "cosmos.bank.v1beta1.SendAuthorization.allow_list":
		if x.AllowList == nil {
			x.AllowList = []string{}
		}
		value := &_SendAuthorization_2_list{list: &x.AllowList}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.SendAuthorization"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.SendAuthorization does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SendAuthorization) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "cosmos.bank.v1beta1.SendAuthorization.spend_limit":
		list := []*v1beta1.Coin{}
		return protoreflect.ValueOfList(&_SendAuthorization_1_list{list: &list})
	case "cosmos.bank.v1beta1.SendAuthorization.allow_list":
		list := []string{}
		return protoreflect.ValueOfList(&_SendAuthorization_2_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: cosmos.bank.v1beta1.SendAuthorization"))
		}
		panic(fmt.Errorf("message cosmos.bank.v1beta1.SendAuthorization does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SendAuthorization) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in cosmos.bank.v1beta1.SendAuthorization", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SendAuthorization) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SendAuthorization) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_SendAuthorization) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SendAuthorization) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SendAuthorization)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.SpendLimit) > 0 {
			for _, e := range x.SpendLimit {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.AllowList) > 0 {
			for _, s := range x.AllowList {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*SendAuthorization)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.AllowList) > 0 {
			for iNdEx := len(x.AllowList) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.AllowList[iNdEx])
				copy(dAtA[i:], x.AllowList[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.AllowList[iNdEx])))
				i--
				dAtA[i] = 0x12
			}
		}
		if len(x.SpendLimit) > 0 {
			for iNdEx := len(x.SpendLimit) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.SpendLimit[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*SendAuthorization)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SendAuthorization: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SendAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field SpendLimit", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.SpendLimit = append(x.SpendLimit, &v1beta1.Coin{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.SpendLimit[len(x.SpendLimit)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AllowList", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.AllowList = append(x.AllowList, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: cosmos/bank/v1beta1/authz.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SendAuthorization allows the grantee to spend up to spend_limit coins from
// the granter's account.
//
// Since: cosmos-sdk 0.43
type SendAuthorization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SpendLimit []*v1beta1.Coin `protobuf:"bytes,1,rep,name=spend_limit,json=spendLimit,proto3" json:"spend_limit,omitempty"`
	// allow_list specifies an optional list of addresses to whom the grantee can send tokens on behalf of the
	// granter. If omitted, any recipient is allowed.
	//
	// Since: cosmos-sdk 0.47
	AllowList []string `protobuf:"bytes,2,rep,name=allow_list,json=allowList,proto3" json:"allow_list,omitempty"`
}

func (x *SendAuthorization) Reset() {
	*x = SendAuthorization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cosmos_bank_v1beta1_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendAuthorization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendAuthorization) ProtoMessage() {}

// Deprecated: Use SendAuthorization.ProtoReflect.Descriptor instead.
func (*SendAuthorization) Descriptor() ([]byte, []int) {
	return file_cosmos_bank_v1beta1_authz_proto_rawDescGZIP(), []int{0}
}

func (x *SendAuthorization) GetSpendLimit() []*v1beta1.Coin {
	if x != nil {
		return x.SpendLimit
	}
	return nil
}

func (x *SendAuthorization) GetAllowList() []string {
	if x != nil {
		return x.AllowList
	}
	return nil
}

var File_cosmos_bank_v1beta1_authz_proto protoreflect.FileDescriptor

var file_cosmos_bank_v1beta1_authz_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x26, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6d,
	0x73, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79, 0x5f, 0x61, 0x6d, 0x69,
	0x6e, 0x6f, 0x2f, 0x61, 0x6d, 0x69, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd9, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x71, 0x0a, 0x0b, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x5f, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x35, 0xc8, 0xde, 0x1f, 0x00, 0xaa, 0xdf, 0x1f, 0x28, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x73, 0xa0, 0xe7, 0xb0, 0x2a, 0x01, 0x52, 0x0a, 0x73, 0x70,
	0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x3a, 0x32, 0xca, 0xb4, 0x2d, 0x0d, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x8a, 0xe7, 0xb0, 0x2a, 0x1c, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x53, 0x65, 0x6e, 0x64, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0xc5, 0x01, 0x0a, 0x17,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x7a, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b,
	0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62,
	0x61, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x62, 0x61, 0x6e, 0x6b,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x42, 0x58, 0xaa, 0x02, 0x13,
	0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x2e, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xca, 0x02, 0x13, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5c, 0x42, 0x61, 0x6e,
	0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x1f, 0x43, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x5c, 0x42, 0x61, 0x6e, 0x6b, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x43, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x3a, 0x3a, 0x42, 0x61, 0x6e, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cosmos_bank_v1beta1_authz_proto_rawDescOnce sync.Once
	file_cosmos_bank_v1beta1_authz_proto_rawDescData = file_cosmos_bank_v1beta1_authz_proto_rawDesc
)

func file_cosmos_bank_v1beta1_authz_proto_rawDescGZIP() []byte {
	file_cosmos_bank_v1beta1_authz_proto_rawDescOnce.Do(func() {
		file_cosmos_bank_v1beta1_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_cosmos_bank_v1beta1_authz_proto_rawDescData)
	})
	return file_cosmos_bank_v1beta1_authz_proto_rawDescData
}

var file_cosmos_bank_v1beta1_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cosmos_bank_v1beta1_authz_proto_goTypes = []interface{}{
	(*SendAuthorization)(nil), // 0: cosmos.bank.v1beta1.SendAuthorization
	(*v1beta1.Coin)(nil),      // 1: cosmos.base.v1beta1.Coin
}
var file_cosmos_bank_v1beta1_authz_proto_depIdxs = []int32{
	1, // 0: cosmos.bank.v1beta1.SendAuthorization.spend_limit:type_name -> cosmos.base.v1beta1.Coin
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cosmos_bank_v1beta1_authz_proto_init() }
func file_cosmos_bank_v1beta1_authz_proto_init() {
	if File_cosmos_bank_v1beta1_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cosmos_bank_v1beta1_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendAuthorization); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cosmos_bank_v1beta1_authz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cosmos_bank_v1beta1_authz_proto_goTypes,
		DependencyIndexes: file_cosmos_bank_v1beta1_authz_proto_depIdxs,
		MessageInfos:      file_cosmos_bank_v1beta1_authz_proto_msgTypes,
	}.Build()
	File_cosmos_bank_v1beta1_authz_proto = out.File
	file_cosmos_bank_v1beta1_authz_proto_rawDesc = nil
	file_cosmos_bank_v1beta1_authz_proto_goTypes = nil
	file_cosmos_bank_v1beta1_authz_proto_depIdxs = nil
}
