// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dota_commonmessages.proto

package dota

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EDOTAStatPopupTypes int32

const (
	EDOTAStatPopupTypes_k_EDOTA_SPT_Textline  EDOTAStatPopupTypes = 0
	EDOTAStatPopupTypes_k_EDOTA_SPT_Basic     EDOTAStatPopupTypes = 1
	EDOTAStatPopupTypes_k_EDOTA_SPT_Poll      EDOTAStatPopupTypes = 2
	EDOTAStatPopupTypes_k_EDOTA_SPT_Grid      EDOTAStatPopupTypes = 3
	EDOTAStatPopupTypes_k_EDOTA_SPT_DualImage EDOTAStatPopupTypes = 4
	EDOTAStatPopupTypes_k_EDOTA_SPT_Movie     EDOTAStatPopupTypes = 5
)

var EDOTAStatPopupTypes_name = map[int32]string{
	0: "k_EDOTA_SPT_Textline",
	1: "k_EDOTA_SPT_Basic",
	2: "k_EDOTA_SPT_Poll",
	3: "k_EDOTA_SPT_Grid",
	4: "k_EDOTA_SPT_DualImage",
	5: "k_EDOTA_SPT_Movie",
}

var EDOTAStatPopupTypes_value = map[string]int32{
	"k_EDOTA_SPT_Textline":  0,
	"k_EDOTA_SPT_Basic":     1,
	"k_EDOTA_SPT_Poll":      2,
	"k_EDOTA_SPT_Grid":      3,
	"k_EDOTA_SPT_DualImage": 4,
	"k_EDOTA_SPT_Movie":     5,
}

func (x EDOTAStatPopupTypes) Enum() *EDOTAStatPopupTypes {
	p := new(EDOTAStatPopupTypes)
	*p = x
	return p
}

func (x EDOTAStatPopupTypes) String() string {
	return proto.EnumName(EDOTAStatPopupTypes_name, int32(x))
}

func (x *EDOTAStatPopupTypes) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EDOTAStatPopupTypes_value, data, "EDOTAStatPopupTypes")
	if err != nil {
		return err
	}
	*x = EDOTAStatPopupTypes(value)
	return nil
}

func (EDOTAStatPopupTypes) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ea12725bdc87a759, []int{0}
}

type DotaunitorderT int32

const (
	DotaunitorderT_DOTA_UNIT_ORDER_NONE                              DotaunitorderT = 0
	DotaunitorderT_DOTA_UNIT_ORDER_MOVE_TO_POSITION                  DotaunitorderT = 1
	DotaunitorderT_DOTA_UNIT_ORDER_MOVE_TO_TARGET                    DotaunitorderT = 2
	DotaunitorderT_DOTA_UNIT_ORDER_ATTACK_MOVE                       DotaunitorderT = 3
	DotaunitorderT_DOTA_UNIT_ORDER_ATTACK_TARGET                     DotaunitorderT = 4
	DotaunitorderT_DOTA_UNIT_ORDER_CAST_POSITION                     DotaunitorderT = 5
	DotaunitorderT_DOTA_UNIT_ORDER_CAST_TARGET                       DotaunitorderT = 6
	DotaunitorderT_DOTA_UNIT_ORDER_CAST_TARGET_TREE                  DotaunitorderT = 7
	DotaunitorderT_DOTA_UNIT_ORDER_CAST_NO_TARGET                    DotaunitorderT = 8
	DotaunitorderT_DOTA_UNIT_ORDER_CAST_TOGGLE                       DotaunitorderT = 9
	DotaunitorderT_DOTA_UNIT_ORDER_HOLD_POSITION                     DotaunitorderT = 10
	DotaunitorderT_DOTA_UNIT_ORDER_TRAIN_ABILITY                     DotaunitorderT = 11
	DotaunitorderT_DOTA_UNIT_ORDER_DROP_ITEM                         DotaunitorderT = 12
	DotaunitorderT_DOTA_UNIT_ORDER_GIVE_ITEM                         DotaunitorderT = 13
	DotaunitorderT_DOTA_UNIT_ORDER_PICKUP_ITEM                       DotaunitorderT = 14
	DotaunitorderT_DOTA_UNIT_ORDER_PICKUP_RUNE                       DotaunitorderT = 15
	DotaunitorderT_DOTA_UNIT_ORDER_PURCHASE_ITEM                     DotaunitorderT = 16
	DotaunitorderT_DOTA_UNIT_ORDER_SELL_ITEM                         DotaunitorderT = 17
	DotaunitorderT_DOTA_UNIT_ORDER_DISASSEMBLE_ITEM                  DotaunitorderT = 18
	DotaunitorderT_DOTA_UNIT_ORDER_MOVE_ITEM                         DotaunitorderT = 19
	DotaunitorderT_DOTA_UNIT_ORDER_CAST_TOGGLE_AUTO                  DotaunitorderT = 20
	DotaunitorderT_DOTA_UNIT_ORDER_STOP                              DotaunitorderT = 21
	DotaunitorderT_DOTA_UNIT_ORDER_TAUNT                             DotaunitorderT = 22
	DotaunitorderT_DOTA_UNIT_ORDER_BUYBACK                           DotaunitorderT = 23
	DotaunitorderT_DOTA_UNIT_ORDER_GLYPH                             DotaunitorderT = 24
	DotaunitorderT_DOTA_UNIT_ORDER_EJECT_ITEM_FROM_STASH             DotaunitorderT = 25
	DotaunitorderT_DOTA_UNIT_ORDER_CAST_RUNE                         DotaunitorderT = 26
	DotaunitorderT_DOTA_UNIT_ORDER_PING_ABILITY                      DotaunitorderT = 27
	DotaunitorderT_DOTA_UNIT_ORDER_MOVE_TO_DIRECTION                 DotaunitorderT = 28
	DotaunitorderT_DOTA_UNIT_ORDER_PATROL                            DotaunitorderT = 29
	DotaunitorderT_DOTA_UNIT_ORDER_VECTOR_TARGET_POSITION            DotaunitorderT = 30
	DotaunitorderT_DOTA_UNIT_ORDER_RADAR                             DotaunitorderT = 31
	DotaunitorderT_DOTA_UNIT_ORDER_SET_ITEM_COMBINE_LOCK             DotaunitorderT = 32
	DotaunitorderT_DOTA_UNIT_ORDER_CONTINUE                          DotaunitorderT = 33
	DotaunitorderT_DOTA_UNIT_ORDER_VECTOR_TARGET_CANCELED            DotaunitorderT = 34
	DotaunitorderT_DOTA_UNIT_ORDER_CAST_RIVER_PAINT                  DotaunitorderT = 35
	DotaunitorderT_DOTA_UNIT_ORDER_PREGAME_ADJUST_ITEM_ASSIGNMENT    DotaunitorderT = 36
	DotaunitorderT_DOTA_UNIT_ORDER_DROP_ITEM_AT_FOUNTAIN             DotaunitorderT = 37
	DotaunitorderT_DOTA_UNIT_ORDER_TAKE_ITEM_FROM_NEUTRAL_ITEM_STASH DotaunitorderT = 38
)

var DotaunitorderT_name = map[int32]string{
	0:  "DOTA_UNIT_ORDER_NONE",
	1:  "DOTA_UNIT_ORDER_MOVE_TO_POSITION",
	2:  "DOTA_UNIT_ORDER_MOVE_TO_TARGET",
	3:  "DOTA_UNIT_ORDER_ATTACK_MOVE",
	4:  "DOTA_UNIT_ORDER_ATTACK_TARGET",
	5:  "DOTA_UNIT_ORDER_CAST_POSITION",
	6:  "DOTA_UNIT_ORDER_CAST_TARGET",
	7:  "DOTA_UNIT_ORDER_CAST_TARGET_TREE",
	8:  "DOTA_UNIT_ORDER_CAST_NO_TARGET",
	9:  "DOTA_UNIT_ORDER_CAST_TOGGLE",
	10: "DOTA_UNIT_ORDER_HOLD_POSITION",
	11: "DOTA_UNIT_ORDER_TRAIN_ABILITY",
	12: "DOTA_UNIT_ORDER_DROP_ITEM",
	13: "DOTA_UNIT_ORDER_GIVE_ITEM",
	14: "DOTA_UNIT_ORDER_PICKUP_ITEM",
	15: "DOTA_UNIT_ORDER_PICKUP_RUNE",
	16: "DOTA_UNIT_ORDER_PURCHASE_ITEM",
	17: "DOTA_UNIT_ORDER_SELL_ITEM",
	18: "DOTA_UNIT_ORDER_DISASSEMBLE_ITEM",
	19: "DOTA_UNIT_ORDER_MOVE_ITEM",
	20: "DOTA_UNIT_ORDER_CAST_TOGGLE_AUTO",
	21: "DOTA_UNIT_ORDER_STOP",
	22: "DOTA_UNIT_ORDER_TAUNT",
	23: "DOTA_UNIT_ORDER_BUYBACK",
	24: "DOTA_UNIT_ORDER_GLYPH",
	25: "DOTA_UNIT_ORDER_EJECT_ITEM_FROM_STASH",
	26: "DOTA_UNIT_ORDER_CAST_RUNE",
	27: "DOTA_UNIT_ORDER_PING_ABILITY",
	28: "DOTA_UNIT_ORDER_MOVE_TO_DIRECTION",
	29: "DOTA_UNIT_ORDER_PATROL",
	30: "DOTA_UNIT_ORDER_VECTOR_TARGET_POSITION",
	31: "DOTA_UNIT_ORDER_RADAR",
	32: "DOTA_UNIT_ORDER_SET_ITEM_COMBINE_LOCK",
	33: "DOTA_UNIT_ORDER_CONTINUE",
	34: "DOTA_UNIT_ORDER_VECTOR_TARGET_CANCELED",
	35: "DOTA_UNIT_ORDER_CAST_RIVER_PAINT",
	36: "DOTA_UNIT_ORDER_PREGAME_ADJUST_ITEM_ASSIGNMENT",
	37: "DOTA_UNIT_ORDER_DROP_ITEM_AT_FOUNTAIN",
	38: "DOTA_UNIT_ORDER_TAKE_ITEM_FROM_NEUTRAL_ITEM_STASH",
}

var DotaunitorderT_value = map[string]int32{
	"DOTA_UNIT_ORDER_NONE":                              0,
	"DOTA_UNIT_ORDER_MOVE_TO_POSITION":                  1,
	"DOTA_UNIT_ORDER_MOVE_TO_TARGET":                    2,
	"DOTA_UNIT_ORDER_ATTACK_MOVE":                       3,
	"DOTA_UNIT_ORDER_ATTACK_TARGET":                     4,
	"DOTA_UNIT_ORDER_CAST_POSITION":                     5,
	"DOTA_UNIT_ORDER_CAST_TARGET":                       6,
	"DOTA_UNIT_ORDER_CAST_TARGET_TREE":                  7,
	"DOTA_UNIT_ORDER_CAST_NO_TARGET":                    8,
	"DOTA_UNIT_ORDER_CAST_TOGGLE":                       9,
	"DOTA_UNIT_ORDER_HOLD_POSITION":                     10,
	"DOTA_UNIT_ORDER_TRAIN_ABILITY":                     11,
	"DOTA_UNIT_ORDER_DROP_ITEM":                         12,
	"DOTA_UNIT_ORDER_GIVE_ITEM":                         13,
	"DOTA_UNIT_ORDER_PICKUP_ITEM":                       14,
	"DOTA_UNIT_ORDER_PICKUP_RUNE":                       15,
	"DOTA_UNIT_ORDER_PURCHASE_ITEM":                     16,
	"DOTA_UNIT_ORDER_SELL_ITEM":                         17,
	"DOTA_UNIT_ORDER_DISASSEMBLE_ITEM":                  18,
	"DOTA_UNIT_ORDER_MOVE_ITEM":                         19,
	"DOTA_UNIT_ORDER_CAST_TOGGLE_AUTO":                  20,
	"DOTA_UNIT_ORDER_STOP":                              21,
	"DOTA_UNIT_ORDER_TAUNT":                             22,
	"DOTA_UNIT_ORDER_BUYBACK":                           23,
	"DOTA_UNIT_ORDER_GLYPH":                             24,
	"DOTA_UNIT_ORDER_EJECT_ITEM_FROM_STASH":             25,
	"DOTA_UNIT_ORDER_CAST_RUNE":                         26,
	"DOTA_UNIT_ORDER_PING_ABILITY":                      27,
	"DOTA_UNIT_ORDER_MOVE_TO_DIRECTION":                 28,
	"DOTA_UNIT_ORDER_PATROL":                            29,
	"DOTA_UNIT_ORDER_VECTOR_TARGET_POSITION":            30,
	"DOTA_UNIT_ORDER_RADAR":                             31,
	"DOTA_UNIT_ORDER_SET_ITEM_COMBINE_LOCK":             32,
	"DOTA_UNIT_ORDER_CONTINUE":                          33,
	"DOTA_UNIT_ORDER_VECTOR_TARGET_CANCELED":            34,
	"DOTA_UNIT_ORDER_CAST_RIVER_PAINT":                  35,
	"DOTA_UNIT_ORDER_PREGAME_ADJUST_ITEM_ASSIGNMENT":    36,
	"DOTA_UNIT_ORDER_DROP_ITEM_AT_FOUNTAIN":             37,
	"DOTA_UNIT_ORDER_TAKE_ITEM_FROM_NEUTRAL_ITEM_STASH": 38,
}

func (x DotaunitorderT) Enum() *DotaunitorderT {
	p := new(DotaunitorderT)
	*p = x
	return p
}

func (x DotaunitorderT) String() string {
	return proto.EnumName(DotaunitorderT_name, int32(x))
}

func (x *DotaunitorderT) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DotaunitorderT_value, data, "DotaunitorderT")
	if err != nil {
		return err
	}
	*x = DotaunitorderT(value)
	return nil
}

func (DotaunitorderT) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ea12725bdc87a759, []int{1}
}

type CDOTAMsg_LocationPing struct {
	X                    *int32   `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y                    *int32   `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
	Target               *int32   `protobuf:"varint,3,opt,name=target" json:"target,omitempty"`
	DirectPing           *bool    `protobuf:"varint,4,opt,name=direct_ping,json=directPing" json:"direct_ping,omitempty"`
	Type                 *int32   `protobuf:"varint,5,opt,name=type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CDOTAMsg_LocationPing) Reset()         { *m = CDOTAMsg_LocationPing{} }
func (m *CDOTAMsg_LocationPing) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_LocationPing) ProtoMessage()    {}
func (*CDOTAMsg_LocationPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea12725bdc87a759, []int{0}
}

func (m *CDOTAMsg_LocationPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CDOTAMsg_LocationPing.Unmarshal(m, b)
}
func (m *CDOTAMsg_LocationPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CDOTAMsg_LocationPing.Marshal(b, m, deterministic)
}
func (m *CDOTAMsg_LocationPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CDOTAMsg_LocationPing.Merge(m, src)
}
func (m *CDOTAMsg_LocationPing) XXX_Size() int {
	return xxx_messageInfo_CDOTAMsg_LocationPing.Size(m)
}
func (m *CDOTAMsg_LocationPing) XXX_DiscardUnknown() {
	xxx_messageInfo_CDOTAMsg_LocationPing.DiscardUnknown(m)
}

var xxx_messageInfo_CDOTAMsg_LocationPing proto.InternalMessageInfo

func (m *CDOTAMsg_LocationPing) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *CDOTAMsg_LocationPing) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *CDOTAMsg_LocationPing) GetTarget() int32 {
	if m != nil && m.Target != nil {
		return *m.Target
	}
	return 0
}

func (m *CDOTAMsg_LocationPing) GetDirectPing() bool {
	if m != nil && m.DirectPing != nil {
		return *m.DirectPing
	}
	return false
}

func (m *CDOTAMsg_LocationPing) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

type CDOTAMsg_ItemAlert struct {
	X                    *int32   `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y                    *int32   `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
	ItemAbilityId        *int32   `protobuf:"varint,3,opt,name=item_ability_id,json=itemAbilityId" json:"item_ability_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CDOTAMsg_ItemAlert) Reset()         { *m = CDOTAMsg_ItemAlert{} }
func (m *CDOTAMsg_ItemAlert) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_ItemAlert) ProtoMessage()    {}
func (*CDOTAMsg_ItemAlert) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea12725bdc87a759, []int{1}
}

func (m *CDOTAMsg_ItemAlert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CDOTAMsg_ItemAlert.Unmarshal(m, b)
}
func (m *CDOTAMsg_ItemAlert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CDOTAMsg_ItemAlert.Marshal(b, m, deterministic)
}
func (m *CDOTAMsg_ItemAlert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CDOTAMsg_ItemAlert.Merge(m, src)
}
func (m *CDOTAMsg_ItemAlert) XXX_Size() int {
	return xxx_messageInfo_CDOTAMsg_ItemAlert.Size(m)
}
func (m *CDOTAMsg_ItemAlert) XXX_DiscardUnknown() {
	xxx_messageInfo_CDOTAMsg_ItemAlert.DiscardUnknown(m)
}

var xxx_messageInfo_CDOTAMsg_ItemAlert proto.InternalMessageInfo

func (m *CDOTAMsg_ItemAlert) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *CDOTAMsg_ItemAlert) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *CDOTAMsg_ItemAlert) GetItemAbilityId() int32 {
	if m != nil && m.ItemAbilityId != nil {
		return *m.ItemAbilityId
	}
	return 0
}

type CDOTAMsg_MapLine struct {
	X                    *int32   `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y                    *int32   `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
	Initial              *bool    `protobuf:"varint,3,opt,name=initial" json:"initial,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CDOTAMsg_MapLine) Reset()         { *m = CDOTAMsg_MapLine{} }
func (m *CDOTAMsg_MapLine) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_MapLine) ProtoMessage()    {}
func (*CDOTAMsg_MapLine) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea12725bdc87a759, []int{2}
}

func (m *CDOTAMsg_MapLine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CDOTAMsg_MapLine.Unmarshal(m, b)
}
func (m *CDOTAMsg_MapLine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CDOTAMsg_MapLine.Marshal(b, m, deterministic)
}
func (m *CDOTAMsg_MapLine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CDOTAMsg_MapLine.Merge(m, src)
}
func (m *CDOTAMsg_MapLine) XXX_Size() int {
	return xxx_messageInfo_CDOTAMsg_MapLine.Size(m)
}
func (m *CDOTAMsg_MapLine) XXX_DiscardUnknown() {
	xxx_messageInfo_CDOTAMsg_MapLine.DiscardUnknown(m)
}

var xxx_messageInfo_CDOTAMsg_MapLine proto.InternalMessageInfo

func (m *CDOTAMsg_MapLine) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *CDOTAMsg_MapLine) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *CDOTAMsg_MapLine) GetInitial() bool {
	if m != nil && m.Initial != nil {
		return *m.Initial
	}
	return false
}

type CDOTAMsg_WorldLine struct {
	X                    *int32   `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y                    *int32   `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
	Z                    *int32   `protobuf:"varint,3,opt,name=z" json:"z,omitempty"`
	Initial              *bool    `protobuf:"varint,4,opt,name=initial" json:"initial,omitempty"`
	End                  *bool    `protobuf:"varint,5,opt,name=end" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CDOTAMsg_WorldLine) Reset()         { *m = CDOTAMsg_WorldLine{} }
func (m *CDOTAMsg_WorldLine) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_WorldLine) ProtoMessage()    {}
func (*CDOTAMsg_WorldLine) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea12725bdc87a759, []int{3}
}

func (m *CDOTAMsg_WorldLine) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CDOTAMsg_WorldLine.Unmarshal(m, b)
}
func (m *CDOTAMsg_WorldLine) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CDOTAMsg_WorldLine.Marshal(b, m, deterministic)
}
func (m *CDOTAMsg_WorldLine) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CDOTAMsg_WorldLine.Merge(m, src)
}
func (m *CDOTAMsg_WorldLine) XXX_Size() int {
	return xxx_messageInfo_CDOTAMsg_WorldLine.Size(m)
}
func (m *CDOTAMsg_WorldLine) XXX_DiscardUnknown() {
	xxx_messageInfo_CDOTAMsg_WorldLine.DiscardUnknown(m)
}

var xxx_messageInfo_CDOTAMsg_WorldLine proto.InternalMessageInfo

func (m *CDOTAMsg_WorldLine) GetX() int32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *CDOTAMsg_WorldLine) GetY() int32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *CDOTAMsg_WorldLine) GetZ() int32 {
	if m != nil && m.Z != nil {
		return *m.Z
	}
	return 0
}

func (m *CDOTAMsg_WorldLine) GetInitial() bool {
	if m != nil && m.Initial != nil {
		return *m.Initial
	}
	return false
}

func (m *CDOTAMsg_WorldLine) GetEnd() bool {
	if m != nil && m.End != nil {
		return *m.End
	}
	return false
}

type CDOTAMsg_SendStatPopup struct {
	Style                *EDOTAStatPopupTypes `protobuf:"varint,1,opt,name=style,enum=dota.EDOTAStatPopupTypes,def=0" json:"style,omitempty"`
	StatStrings          []string             `protobuf:"bytes,2,rep,name=stat_strings,json=statStrings" json:"stat_strings,omitempty"`
	StatImages           []int32              `protobuf:"varint,3,rep,name=stat_images,json=statImages" json:"stat_images,omitempty"`
	StatImageTypes       []int32              `protobuf:"varint,4,rep,name=stat_image_types,json=statImageTypes" json:"stat_image_types,omitempty"`
	Duration             *float32             `protobuf:"fixed32,5,opt,name=duration" json:"duration,omitempty"`
	UseHtml              *bool                `protobuf:"varint,6,opt,name=use_html,json=useHtml" json:"use_html,omitempty"`
	MovieName            *string              `protobuf:"bytes,7,opt,name=movie_name,json=movieName" json:"movie_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CDOTAMsg_SendStatPopup) Reset()         { *m = CDOTAMsg_SendStatPopup{} }
func (m *CDOTAMsg_SendStatPopup) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_SendStatPopup) ProtoMessage()    {}
func (*CDOTAMsg_SendStatPopup) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea12725bdc87a759, []int{4}
}

func (m *CDOTAMsg_SendStatPopup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CDOTAMsg_SendStatPopup.Unmarshal(m, b)
}
func (m *CDOTAMsg_SendStatPopup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CDOTAMsg_SendStatPopup.Marshal(b, m, deterministic)
}
func (m *CDOTAMsg_SendStatPopup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CDOTAMsg_SendStatPopup.Merge(m, src)
}
func (m *CDOTAMsg_SendStatPopup) XXX_Size() int {
	return xxx_messageInfo_CDOTAMsg_SendStatPopup.Size(m)
}
func (m *CDOTAMsg_SendStatPopup) XXX_DiscardUnknown() {
	xxx_messageInfo_CDOTAMsg_SendStatPopup.DiscardUnknown(m)
}

var xxx_messageInfo_CDOTAMsg_SendStatPopup proto.InternalMessageInfo

const Default_CDOTAMsg_SendStatPopup_Style EDOTAStatPopupTypes = EDOTAStatPopupTypes_k_EDOTA_SPT_Textline

func (m *CDOTAMsg_SendStatPopup) GetStyle() EDOTAStatPopupTypes {
	if m != nil && m.Style != nil {
		return *m.Style
	}
	return Default_CDOTAMsg_SendStatPopup_Style
}

func (m *CDOTAMsg_SendStatPopup) GetStatStrings() []string {
	if m != nil {
		return m.StatStrings
	}
	return nil
}

func (m *CDOTAMsg_SendStatPopup) GetStatImages() []int32 {
	if m != nil {
		return m.StatImages
	}
	return nil
}

func (m *CDOTAMsg_SendStatPopup) GetStatImageTypes() []int32 {
	if m != nil {
		return m.StatImageTypes
	}
	return nil
}

func (m *CDOTAMsg_SendStatPopup) GetDuration() float32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *CDOTAMsg_SendStatPopup) GetUseHtml() bool {
	if m != nil && m.UseHtml != nil {
		return *m.UseHtml
	}
	return false
}

func (m *CDOTAMsg_SendStatPopup) GetMovieName() string {
	if m != nil && m.MovieName != nil {
		return *m.MovieName
	}
	return ""
}

type CDOTAMsg_DismissAllStatPopups struct {
	TimeDelay            *float32 `protobuf:"fixed32,1,opt,name=time_delay,json=timeDelay" json:"time_delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CDOTAMsg_DismissAllStatPopups) Reset()         { *m = CDOTAMsg_DismissAllStatPopups{} }
func (m *CDOTAMsg_DismissAllStatPopups) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_DismissAllStatPopups) ProtoMessage()    {}
func (*CDOTAMsg_DismissAllStatPopups) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea12725bdc87a759, []int{5}
}

func (m *CDOTAMsg_DismissAllStatPopups) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CDOTAMsg_DismissAllStatPopups.Unmarshal(m, b)
}
func (m *CDOTAMsg_DismissAllStatPopups) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CDOTAMsg_DismissAllStatPopups.Marshal(b, m, deterministic)
}
func (m *CDOTAMsg_DismissAllStatPopups) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CDOTAMsg_DismissAllStatPopups.Merge(m, src)
}
func (m *CDOTAMsg_DismissAllStatPopups) XXX_Size() int {
	return xxx_messageInfo_CDOTAMsg_DismissAllStatPopups.Size(m)
}
func (m *CDOTAMsg_DismissAllStatPopups) XXX_DiscardUnknown() {
	xxx_messageInfo_CDOTAMsg_DismissAllStatPopups.DiscardUnknown(m)
}

var xxx_messageInfo_CDOTAMsg_DismissAllStatPopups proto.InternalMessageInfo

func (m *CDOTAMsg_DismissAllStatPopups) GetTimeDelay() float32 {
	if m != nil && m.TimeDelay != nil {
		return *m.TimeDelay
	}
	return 0
}

type CDOTAMsg_CoachHUDPing struct {
	X                    *uint32  `protobuf:"varint,1,opt,name=x" json:"x,omitempty"`
	Y                    *uint32  `protobuf:"varint,2,opt,name=y" json:"y,omitempty"`
	Tgtpath              *string  `protobuf:"bytes,3,opt,name=tgtpath" json:"tgtpath,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CDOTAMsg_CoachHUDPing) Reset()         { *m = CDOTAMsg_CoachHUDPing{} }
func (m *CDOTAMsg_CoachHUDPing) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_CoachHUDPing) ProtoMessage()    {}
func (*CDOTAMsg_CoachHUDPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea12725bdc87a759, []int{6}
}

func (m *CDOTAMsg_CoachHUDPing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CDOTAMsg_CoachHUDPing.Unmarshal(m, b)
}
func (m *CDOTAMsg_CoachHUDPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CDOTAMsg_CoachHUDPing.Marshal(b, m, deterministic)
}
func (m *CDOTAMsg_CoachHUDPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CDOTAMsg_CoachHUDPing.Merge(m, src)
}
func (m *CDOTAMsg_CoachHUDPing) XXX_Size() int {
	return xxx_messageInfo_CDOTAMsg_CoachHUDPing.Size(m)
}
func (m *CDOTAMsg_CoachHUDPing) XXX_DiscardUnknown() {
	xxx_messageInfo_CDOTAMsg_CoachHUDPing.DiscardUnknown(m)
}

var xxx_messageInfo_CDOTAMsg_CoachHUDPing proto.InternalMessageInfo

func (m *CDOTAMsg_CoachHUDPing) GetX() uint32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *CDOTAMsg_CoachHUDPing) GetY() uint32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *CDOTAMsg_CoachHUDPing) GetTgtpath() string {
	if m != nil && m.Tgtpath != nil {
		return *m.Tgtpath
	}
	return ""
}

type CDOTAMsg_UnitOrder struct {
	Issuer               *int32          `protobuf:"zigzag32,1,opt,name=issuer,def=-1" json:"issuer,omitempty"`
	OrderType            *DotaunitorderT `protobuf:"varint,2,opt,name=order_type,json=orderType,enum=dota.DotaunitorderT,def=0" json:"order_type,omitempty"`
	Units                []int32         `protobuf:"varint,3,rep,name=units" json:"units,omitempty"`
	TargetIndex          *int32          `protobuf:"varint,4,opt,name=target_index,json=targetIndex" json:"target_index,omitempty"`
	AbilityIndex         *int32          `protobuf:"varint,5,opt,name=ability_index,json=abilityIndex" json:"ability_index,omitempty"`
	Position             *CMsgVector     `protobuf:"bytes,6,opt,name=position" json:"position,omitempty"`
	Queue                *bool           `protobuf:"varint,7,opt,name=queue" json:"queue,omitempty"`
	SequenceNumber       *int32          `protobuf:"varint,8,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	Flags                *uint32         `protobuf:"varint,9,opt,name=flags" json:"flags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CDOTAMsg_UnitOrder) Reset()         { *m = CDOTAMsg_UnitOrder{} }
func (m *CDOTAMsg_UnitOrder) String() string { return proto.CompactTextString(m) }
func (*CDOTAMsg_UnitOrder) ProtoMessage()    {}
func (*CDOTAMsg_UnitOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea12725bdc87a759, []int{7}
}

func (m *CDOTAMsg_UnitOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CDOTAMsg_UnitOrder.Unmarshal(m, b)
}
func (m *CDOTAMsg_UnitOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CDOTAMsg_UnitOrder.Marshal(b, m, deterministic)
}
func (m *CDOTAMsg_UnitOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CDOTAMsg_UnitOrder.Merge(m, src)
}
func (m *CDOTAMsg_UnitOrder) XXX_Size() int {
	return xxx_messageInfo_CDOTAMsg_UnitOrder.Size(m)
}
func (m *CDOTAMsg_UnitOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_CDOTAMsg_UnitOrder.DiscardUnknown(m)
}

var xxx_messageInfo_CDOTAMsg_UnitOrder proto.InternalMessageInfo

const Default_CDOTAMsg_UnitOrder_Issuer int32 = -1
const Default_CDOTAMsg_UnitOrder_OrderType DotaunitorderT = DotaunitorderT_DOTA_UNIT_ORDER_NONE

func (m *CDOTAMsg_UnitOrder) GetIssuer() int32 {
	if m != nil && m.Issuer != nil {
		return *m.Issuer
	}
	return Default_CDOTAMsg_UnitOrder_Issuer
}

func (m *CDOTAMsg_UnitOrder) GetOrderType() DotaunitorderT {
	if m != nil && m.OrderType != nil {
		return *m.OrderType
	}
	return Default_CDOTAMsg_UnitOrder_OrderType
}

func (m *CDOTAMsg_UnitOrder) GetUnits() []int32 {
	if m != nil {
		return m.Units
	}
	return nil
}

func (m *CDOTAMsg_UnitOrder) GetTargetIndex() int32 {
	if m != nil && m.TargetIndex != nil {
		return *m.TargetIndex
	}
	return 0
}

func (m *CDOTAMsg_UnitOrder) GetAbilityIndex() int32 {
	if m != nil && m.AbilityIndex != nil {
		return *m.AbilityIndex
	}
	return 0
}

func (m *CDOTAMsg_UnitOrder) GetPosition() *CMsgVector {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *CDOTAMsg_UnitOrder) GetQueue() bool {
	if m != nil && m.Queue != nil {
		return *m.Queue
	}
	return false
}

func (m *CDOTAMsg_UnitOrder) GetSequenceNumber() int32 {
	if m != nil && m.SequenceNumber != nil {
		return *m.SequenceNumber
	}
	return 0
}

func (m *CDOTAMsg_UnitOrder) GetFlags() uint32 {
	if m != nil && m.Flags != nil {
		return *m.Flags
	}
	return 0
}

func init() {
	proto.RegisterEnum("dota.EDOTAStatPopupTypes", EDOTAStatPopupTypes_name, EDOTAStatPopupTypes_value)
	proto.RegisterEnum("dota.DotaunitorderT", DotaunitorderT_name, DotaunitorderT_value)
	proto.RegisterType((*CDOTAMsg_LocationPing)(nil), "dota.CDOTAMsg_LocationPing")
	proto.RegisterType((*CDOTAMsg_ItemAlert)(nil), "dota.CDOTAMsg_ItemAlert")
	proto.RegisterType((*CDOTAMsg_MapLine)(nil), "dota.CDOTAMsg_MapLine")
	proto.RegisterType((*CDOTAMsg_WorldLine)(nil), "dota.CDOTAMsg_WorldLine")
	proto.RegisterType((*CDOTAMsg_SendStatPopup)(nil), "dota.CDOTAMsg_SendStatPopup")
	proto.RegisterType((*CDOTAMsg_DismissAllStatPopups)(nil), "dota.CDOTAMsg_DismissAllStatPopups")
	proto.RegisterType((*CDOTAMsg_CoachHUDPing)(nil), "dota.CDOTAMsg_CoachHUDPing")
	proto.RegisterType((*CDOTAMsg_UnitOrder)(nil), "dota.CDOTAMsg_UnitOrder")
}

func init() { proto.RegisterFile("dota_commonmessages.proto", fileDescriptor_ea12725bdc87a759) }

var fileDescriptor_ea12725bdc87a759 = []byte{
	// 1229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xdb, 0x72, 0xdb, 0xb6,
	0x16, 0x0d, 0x75, 0x71, 0x24, 0xf8, 0x86, 0x20, 0xb6, 0x43, 0x3b, 0x71, 0x22, 0x2b, 0x97, 0xe3,
	0x93, 0x69, 0x3d, 0x13, 0xcf, 0xf4, 0x25, 0x0f, 0x9d, 0xa1, 0x28, 0x44, 0x62, 0x2c, 0x91, 0x1a,
	0x10, 0x72, 0x9b, 0x27, 0x0c, 0x23, 0xa1, 0x0a, 0x1b, 0x5e, 0x14, 0x12, 0x6a, 0xed, 0x3c, 0xe5,
	0x5f, 0xda, 0xdf, 0xea, 0x37, 0xf4, 0x17, 0x3a, 0x00, 0x69, 0x46, 0x91, 0x69, 0xe7, 0x45, 0x23,
	0xac, 0xbd, 0xb8, 0xb1, 0xf6, 0x5e, 0x0b, 0x60, 0x7f, 0x1a, 0x0b, 0x8f, 0x4d, 0xe2, 0x30, 0x8c,
	0xa3, 0x90, 0xa7, 0xa9, 0x37, 0xe3, 0xe9, 0xc9, 0x3c, 0x89, 0x45, 0x8c, 0x6a, 0xb2, 0x74, 0xb0,
	0x17, 0x71, 0xf1, 0x67, 0x9c, 0x7c, 0x7c, 0xef, 0xa5, 0x5c, 0x5c, 0xce, 0xaf, 0xaa, 0xed, 0x2f,
	0x1a, 0xd8, 0x35, 0xbb, 0x0e, 0x35, 0x86, 0xe9, 0x8c, 0x0d, 0xe2, 0x89, 0x27, 0xfc, 0x38, 0x1a,
	0xf9, 0xd1, 0x0c, 0x6d, 0x00, 0xed, 0x42, 0xd7, 0x5a, 0xda, 0x71, 0x9d, 0x68, 0x17, 0xf2, 0x74,
	0xa9, 0x57, 0xb2, 0xd3, 0x25, 0xda, 0x03, 0x6b, 0xc2, 0x4b, 0x66, 0x5c, 0xe8, 0x55, 0x05, 0xe5,
	0x27, 0xf4, 0x04, 0xac, 0x4f, 0xfd, 0x84, 0x4f, 0x04, 0x9b, 0xfb, 0xd1, 0x4c, 0xaf, 0xb5, 0xb4,
	0xe3, 0x06, 0x01, 0x19, 0xa4, 0x9a, 0x22, 0x50, 0x93, 0xb7, 0xeb, 0x75, 0xf5, 0x99, 0xfa, 0xdf,
	0xfe, 0x15, 0xa0, 0x42, 0x81, 0x25, 0x78, 0x68, 0x04, 0x3c, 0x11, 0xb7, 0x5e, 0xff, 0x02, 0x6c,
	0xfb, 0x82, 0x87, 0xcc, 0x7b, 0xef, 0x07, 0xbe, 0xb8, 0x64, 0xfe, 0x34, 0xd7, 0xb1, 0x29, 0x61,
	0x23, 0x43, 0xad, 0x69, 0xbb, 0x0f, 0x60, 0xd1, 0x79, 0xe8, 0xcd, 0x07, 0x7e, 0xc4, 0x6f, 0xed,
	0xab, 0x83, 0xbb, 0x7e, 0xe4, 0x0b, 0xdf, 0x0b, 0x54, 0xbf, 0x06, 0xb9, 0x3a, 0xb6, 0x7f, 0x5f,
	0xd2, 0xf8, 0x4b, 0x9c, 0x04, 0xd3, 0xef, 0xf6, 0xda, 0x00, 0xda, 0xe7, 0x5c, 0x95, 0xf6, 0x79,
	0xb9, 0x73, 0xed, 0x9b, 0xce, 0x08, 0x82, 0x2a, 0x8f, 0xa6, 0x6a, 0x21, 0x0d, 0x22, 0xff, 0xb6,
	0xff, 0xaa, 0x80, 0xbd, 0xe2, 0x32, 0x97, 0x47, 0x53, 0x57, 0x78, 0x62, 0x14, 0xcf, 0x17, 0x73,
	0x84, 0x41, 0x3d, 0x15, 0x97, 0x01, 0x57, 0x97, 0x6e, 0x9d, 0xee, 0x9f, 0x48, 0x6f, 0x4f, 0xb0,
	0x24, 0x17, 0x24, 0x2a, 0xdd, 0x7d, 0xbd, 0xf3, 0x91, 0x29, 0x98, 0xb9, 0x23, 0xca, 0x28, 0xbf,
	0x10, 0x81, 0x1f, 0x71, 0x92, 0x7d, 0x8d, 0x8e, 0xc0, 0x46, 0x2a, 0x3c, 0xc1, 0x52, 0x91, 0xf8,
	0xd1, 0x2c, 0xd5, 0x2b, 0xad, 0xea, 0x71, 0x93, 0xac, 0x4b, 0xcc, 0xcd, 0x20, 0xe9, 0xa4, 0xa2,
	0xf8, 0xa1, 0x8c, 0x92, 0x5e, 0x6d, 0x55, 0x8f, 0xeb, 0x04, 0x48, 0xc8, 0x52, 0x08, 0x3a, 0x06,
	0xf0, 0x2b, 0x81, 0xa9, 0x48, 0xe9, 0x35, 0xc5, 0xda, 0x2a, 0x58, 0x4a, 0x0a, 0x3a, 0x00, 0x8d,
	0xe9, 0x22, 0x51, 0xc1, 0x52, 0x63, 0x56, 0x48, 0x71, 0x46, 0xfb, 0xa0, 0xb1, 0x48, 0x39, 0xfb,
	0x20, 0xc2, 0x40, 0x5f, 0xcb, 0x16, 0xb3, 0x48, 0x79, 0x5f, 0x84, 0x01, 0x3a, 0x04, 0x20, 0x8c,
	0xff, 0xf0, 0x39, 0x8b, 0xbc, 0x90, 0xeb, 0x77, 0x5b, 0xda, 0x71, 0x93, 0x34, 0x15, 0x62, 0x7b,
	0x21, 0x6f, 0xff, 0x0c, 0x0e, 0x8b, 0x25, 0x75, 0xfd, 0x34, 0xf4, 0xd3, 0xd4, 0x08, 0x82, 0x62,
	0x0b, 0xa9, 0xfc, 0x5e, 0xf8, 0x21, 0x67, 0x53, 0x1e, 0x78, 0x97, 0x6a, 0x61, 0x15, 0xd2, 0x94,
	0x48, 0x57, 0x02, 0xed, 0xe1, 0x52, 0xee, 0xcd, 0xd8, 0x9b, 0x7c, 0xe8, 0x8f, 0xbb, 0xdf, 0xe6,
	0x7e, 0xf3, 0x1b, 0x53, 0x37, 0xf3, 0x80, 0x88, 0x99, 0x98, 0x7b, 0xe2, 0x83, 0xb2, 0xb6, 0x49,
	0xae, 0x8e, 0xed, 0x7f, 0x2a, 0x4b, 0x09, 0x19, 0x47, 0xbe, 0x70, 0x92, 0x29, 0x4f, 0xd0, 0x01,
	0x58, 0xf3, 0xd3, 0x74, 0xc1, 0x13, 0xd5, 0xf1, 0xde, 0xeb, 0xca, 0x8f, 0xaf, 0x48, 0x8e, 0xa0,
	0xb7, 0x00, 0xc4, 0x92, 0xa4, 0x96, 0xa7, 0xee, 0xd8, 0x3a, 0xdd, 0xcd, 0x1c, 0x95, 0x3f, 0x8b,
	0xc8, 0x17, 0x79, 0xfd, 0xf5, 0x8e, 0xb2, 0x72, 0x6c, 0x5b, 0x94, 0x39, 0xa4, 0x8b, 0x09, 0xb3,
	0x1d, 0x1b, 0x93, 0xa6, 0x2a, 0xcb, 0x25, 0xa3, 0x1d, 0x50, 0x97, 0xfc, 0x2b, 0xa3, 0xb2, 0x83,
	0xf4, 0x39, 0x7b, 0x98, 0xcc, 0x8f, 0xa6, 0xfc, 0x42, 0x45, 0xaf, 0x4e, 0xd6, 0x33, 0xcc, 0x92,
	0x10, 0x7a, 0x0a, 0x36, 0x8b, 0x57, 0xa4, 0x38, 0xd9, 0xcb, 0xdc, 0xc8, 0xc1, 0x8c, 0xf4, 0x03,
	0x68, 0xcc, 0xe3, 0xd4, 0x57, 0x0e, 0x4a, 0x97, 0xd6, 0x4f, 0x61, 0xa6, 0xd3, 0x1c, 0xa6, 0xb3,
	0x73, 0x3e, 0x11, 0x71, 0x42, 0x0a, 0x86, 0xd4, 0xf2, 0x69, 0xc1, 0x17, 0x99, 0x67, 0x0d, 0x92,
	0x1d, 0xd0, 0xff, 0xc0, 0x76, 0xca, 0x3f, 0x2d, 0x78, 0x34, 0xe1, 0x2c, 0x5a, 0x84, 0xef, 0x79,
	0xa2, 0x37, 0xd4, 0x55, 0x5b, 0x57, 0xb0, 0xad, 0x50, 0xf9, 0xf9, 0x6f, 0x81, 0x37, 0x4b, 0xf5,
	0xa6, 0xda, 0x7a, 0x76, 0x78, 0xf9, 0xb7, 0x06, 0xee, 0x97, 0xe4, 0x1c, 0xe9, 0xa0, 0x34, 0xe9,
	0xf0, 0x0e, 0xda, 0x05, 0xf7, 0x96, 0x2b, 0x1d, 0x2f, 0xf5, 0x27, 0x50, 0x43, 0x3b, 0x00, 0x2e,
	0xc3, 0xa3, 0x38, 0x08, 0x60, 0x65, 0x15, 0xed, 0x25, 0xfe, 0x14, 0x56, 0xd1, 0x3e, 0xd8, 0x5d,
	0x46, 0xbb, 0x0b, 0x2f, 0x50, 0xb9, 0x86, 0xb5, 0xd5, 0xee, 0x43, 0x99, 0x4b, 0x58, 0x7f, 0xf9,
	0x2f, 0x00, 0xdb, 0x2b, 0xe6, 0x49, 0x89, 0x65, 0xf6, 0xc1, 0x3b, 0xe8, 0x19, 0x68, 0xad, 0x56,
	0x86, 0xce, 0x39, 0x66, 0xd4, 0x61, 0x23, 0xc7, 0xb5, 0xa8, 0xe5, 0xd8, 0x50, 0x43, 0x6d, 0xf0,
	0xf8, 0x26, 0x16, 0x35, 0x48, 0x0f, 0x53, 0x58, 0x41, 0x4f, 0xc0, 0xc3, 0x55, 0x8e, 0x41, 0xa9,
	0x61, 0x9e, 0x29, 0x2a, 0xac, 0xa2, 0x23, 0x70, 0x78, 0x03, 0x21, 0xef, 0x51, 0x2b, 0xa3, 0x98,
	0x86, 0x4b, 0xbf, 0x4a, 0xa9, 0x97, 0x5d, 0xa3, 0x28, 0x79, 0x8f, 0xb5, 0xb2, 0x89, 0x96, 0x08,
	0x8c, 0x12, 0x8c, 0xe1, 0xdd, 0xb2, 0x89, 0x14, 0xcb, 0x2e, 0x26, 0x6a, 0xdc, 0x7c, 0x95, 0xd3,
	0xeb, 0x0d, 0x30, 0x6c, 0x96, 0xc9, 0xed, 0x3b, 0x83, 0xee, 0x57, 0xb9, 0xa0, 0x8c, 0x42, 0x89,
	0x61, 0xd9, 0xcc, 0xe8, 0x58, 0x03, 0x8b, 0xbe, 0x83, 0xeb, 0xe8, 0x10, 0xec, 0xaf, 0x52, 0xba,
	0xc4, 0x19, 0x31, 0x8b, 0xe2, 0x21, 0xdc, 0x28, 0x2b, 0xf7, 0xac, 0x73, 0x9c, 0x95, 0x37, 0xcb,
	0x44, 0x8e, 0x2c, 0xf3, 0x6c, 0x9c, 0x7f, 0xbf, 0x75, 0x0b, 0x81, 0x8c, 0x6d, 0x0c, 0xb7, 0xcb,
	0x24, 0x8e, 0xc6, 0xc4, 0xec, 0x1b, 0x6e, 0x7e, 0x09, 0x2c, 0xd3, 0xe0, 0xe2, 0xc1, 0x20, 0x2b,
	0xdf, 0x2b, 0x5b, 0x79, 0xd7, 0x72, 0x0d, 0xd7, 0xc5, 0xc3, 0xce, 0x20, 0x6f, 0x82, 0xca, 0x9a,
	0xa8, 0x10, 0xa9, 0xf2, 0xfd, 0x9b, 0x7d, 0x53, 0xdb, 0x66, 0xc6, 0x98, 0x3a, 0x70, 0xa7, 0x2c,
	0xc9, 0x2e, 0x75, 0x46, 0x70, 0x57, 0xbe, 0x94, 0x6b, 0x9b, 0x36, 0xc6, 0x36, 0x85, 0x7b, 0xe8,
	0x21, 0x78, 0xb0, 0x5a, 0xea, 0x8c, 0xdf, 0x75, 0x0c, 0xf3, 0x0c, 0x3e, 0x28, 0xfb, 0xae, 0x37,
	0x78, 0x37, 0xea, 0x43, 0x1d, 0xfd, 0x1f, 0x3c, 0x5f, 0x2d, 0xe1, 0xb7, 0xd8, 0xa4, 0x4a, 0x32,
	0x7b, 0x43, 0x9c, 0x21, 0x73, 0xa9, 0xe1, 0xf6, 0xe1, 0x7e, 0xd9, 0x70, 0x4a, 0xbd, 0xda, 0xf1,
	0x01, 0x6a, 0x81, 0x47, 0xd7, 0x4d, 0xb0, 0x7b, 0x45, 0x0a, 0x1e, 0xa2, 0xe7, 0xe0, 0xe8, 0xa6,
	0x27, 0xd6, 0xb5, 0x08, 0x36, 0x55, 0x9e, 0x1e, 0xa1, 0x03, 0xb0, 0x77, 0xad, 0x91, 0x41, 0x89,
	0x33, 0x80, 0x87, 0xe8, 0x25, 0x78, 0xb1, 0x5a, 0x3b, 0xc7, 0x26, 0x75, 0xc8, 0x55, 0xf6, 0x8b,
	0x5c, 0x3e, 0x2e, 0x9b, 0x9a, 0x18, 0x5d, 0x83, 0xc0, 0x27, 0x65, 0x53, 0xbb, 0x38, 0x9f, 0xd9,
	0x74, 0x86, 0x1d, 0xcb, 0xc6, 0x6c, 0xe0, 0x98, 0x67, 0xb0, 0x85, 0x1e, 0x01, 0xfd, 0xda, 0xd4,
	0x8e, 0x4d, 0x2d, 0x7b, 0x8c, 0xe1, 0xd1, 0xf7, 0xf5, 0x98, 0x86, 0x6d, 0xe2, 0x01, 0xee, 0xc2,
	0xf6, 0x8d, 0xee, 0x13, 0xeb, 0x5c, 0x8d, 0x68, 0xd9, 0x14, 0x3e, 0x45, 0xa7, 0xe0, 0xe4, 0xda,
	0xf4, 0x04, 0xf7, 0x8c, 0x21, 0x66, 0x46, 0xf7, 0xed, 0xd8, 0xcd, 0x55, 0x1a, 0xae, 0x6b, 0xf5,
	0xec, 0x21, 0xb6, 0x29, 0x7c, 0x56, 0x36, 0x4e, 0xf1, 0xbc, 0x98, 0x41, 0xd9, 0x1b, 0x67, 0x6c,
	0x53, 0xc3, 0xb2, 0xe1, 0x73, 0xf4, 0x13, 0x78, 0x75, 0x3d, 0x42, 0x67, 0x78, 0xc9, 0x6e, 0x1b,
	0x8f, 0x29, 0x31, 0xb2, 0xe0, 0xe7, 0xde, 0xbf, 0xe8, 0xd4, 0xfb, 0xda, 0x17, 0xed, 0xce, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0xa8, 0x74, 0x8e, 0x01, 0x0b, 0x00, 0x00,
}
