// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc.proto

package communication

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

type Player struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IsFaction            bool     `protobuf:"varint,2,opt,name=isFaction,proto3" json:"isFaction,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	NumberOfBuilding     int32    `protobuf:"varint,4,opt,name=numberOfBuilding,proto3" json:"numberOfBuilding,omitempty"`
	NumberOfEntities     int32    `protobuf:"varint,5,opt,name=numberOfEntities,proto3" json:"numberOfEntities,omitempty"`
	Stone                int32    `protobuf:"varint,6,opt,name=stone,proto3" json:"stone,omitempty"`
	Wood                 int32    `protobuf:"varint,7,opt,name=wood,proto3" json:"wood,omitempty"`
	Food                 int32    `protobuf:"varint,8,opt,name=food,proto3" json:"food,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{0}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Player) GetIsFaction() bool {
	if m != nil {
		return m.IsFaction
	}
	return false
}

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Player) GetNumberOfBuilding() int32 {
	if m != nil {
		return m.NumberOfBuilding
	}
	return 0
}

func (m *Player) GetNumberOfEntities() int32 {
	if m != nil {
		return m.NumberOfEntities
	}
	return 0
}

func (m *Player) GetStone() int32 {
	if m != nil {
		return m.Stone
	}
	return 0
}

func (m *Player) GetWood() int32 {
	if m != nil {
		return m.Wood
	}
	return 0
}

func (m *Player) GetFood() int32 {
	if m != nil {
		return m.Food
	}
	return 0
}

// Messages liés à getMap
type GetMapRequest struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Width                int32    `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMapRequest) Reset()         { *m = GetMapRequest{} }
func (m *GetMapRequest) String() string { return proto.CompactTextString(m) }
func (*GetMapRequest) ProtoMessage()    {}
func (*GetMapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{1}
}

func (m *GetMapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMapRequest.Unmarshal(m, b)
}
func (m *GetMapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMapRequest.Marshal(b, m, deterministic)
}
func (m *GetMapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMapRequest.Merge(m, src)
}
func (m *GetMapRequest) XXX_Size() int {
	return xxx_messageInfo_GetMapRequest.Size(m)
}
func (m *GetMapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMapRequest proto.InternalMessageInfo

func (m *GetMapRequest) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *GetMapRequest) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *GetMapRequest) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *GetMapRequest) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetMapReply struct {
	Map                  []*Zone  `protobuf:"bytes,1,rep,name=map,proto3" json:"map,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMapReply) Reset()         { *m = GetMapReply{} }
func (m *GetMapReply) String() string { return proto.CompactTextString(m) }
func (*GetMapReply) ProtoMessage()    {}
func (*GetMapReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{2}
}

func (m *GetMapReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMapReply.Unmarshal(m, b)
}
func (m *GetMapReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMapReply.Marshal(b, m, deterministic)
}
func (m *GetMapReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMapReply.Merge(m, src)
}
func (m *GetMapReply) XXX_Size() int {
	return xxx_messageInfo_GetMapReply.Size(m)
}
func (m *GetMapReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMapReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetMapReply proto.InternalMessageInfo

func (m *GetMapReply) GetMap() []*Zone {
	if m != nil {
		return m.Map
	}
	return nil
}

// Messages liés à setMap
type SetMapRequest struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Width                int32    `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	Map                  []*Zone  `protobuf:"bytes,5,rep,name=map,proto3" json:"map,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetMapRequest) Reset()         { *m = SetMapRequest{} }
func (m *SetMapRequest) String() string { return proto.CompactTextString(m) }
func (*SetMapRequest) ProtoMessage()    {}
func (*SetMapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{3}
}

func (m *SetMapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetMapRequest.Unmarshal(m, b)
}
func (m *SetMapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetMapRequest.Marshal(b, m, deterministic)
}
func (m *SetMapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMapRequest.Merge(m, src)
}
func (m *SetMapRequest) XXX_Size() int {
	return xxx_messageInfo_SetMapRequest.Size(m)
}
func (m *SetMapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetMapRequest proto.InternalMessageInfo

func (m *SetMapRequest) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *SetMapRequest) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *SetMapRequest) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *SetMapRequest) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *SetMapRequest) GetMap() []*Zone {
	if m != nil {
		return m.Map
	}
	return nil
}

type SetMapReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetMapReply) Reset()         { *m = SetMapReply{} }
func (m *SetMapReply) String() string { return proto.CompactTextString(m) }
func (*SetMapReply) ProtoMessage()    {}
func (*SetMapReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{4}
}

func (m *SetMapReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetMapReply.Unmarshal(m, b)
}
func (m *SetMapReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetMapReply.Marshal(b, m, deterministic)
}
func (m *SetMapReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetMapReply.Merge(m, src)
}
func (m *SetMapReply) XXX_Size() int {
	return xxx_messageInfo_SetMapReply.Size(m)
}
func (m *SetMapReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SetMapReply.DiscardUnknown(m)
}

var xxx_messageInfo_SetMapReply proto.InternalMessageInfo

// Messages liés à updateMap
type UpdateMapRequest struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	Width                int32    `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	Map                  []*Zone  `protobuf:"bytes,5,rep,name=map,proto3" json:"map,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMapRequest) Reset()         { *m = UpdateMapRequest{} }
func (m *UpdateMapRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateMapRequest) ProtoMessage()    {}
func (*UpdateMapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{5}
}

func (m *UpdateMapRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMapRequest.Unmarshal(m, b)
}
func (m *UpdateMapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMapRequest.Marshal(b, m, deterministic)
}
func (m *UpdateMapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMapRequest.Merge(m, src)
}
func (m *UpdateMapRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateMapRequest.Size(m)
}
func (m *UpdateMapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMapRequest proto.InternalMessageInfo

func (m *UpdateMapRequest) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *UpdateMapRequest) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *UpdateMapRequest) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *UpdateMapRequest) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *UpdateMapRequest) GetMap() []*Zone {
	if m != nil {
		return m.Map
	}
	return nil
}

type UpdateMapReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateMapReply) Reset()         { *m = UpdateMapReply{} }
func (m *UpdateMapReply) String() string { return proto.CompactTextString(m) }
func (*UpdateMapReply) ProtoMessage()    {}
func (*UpdateMapReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{6}
}

func (m *UpdateMapReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMapReply.Unmarshal(m, b)
}
func (m *UpdateMapReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMapReply.Marshal(b, m, deterministic)
}
func (m *UpdateMapReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMapReply.Merge(m, src)
}
func (m *UpdateMapReply) XXX_Size() int {
	return xxx_messageInfo_UpdateMapReply.Size(m)
}
func (m *UpdateMapReply) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMapReply.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMapReply proto.InternalMessageInfo

// Message de structure d'organisation de la carte
type Zone struct {
	Type                 int32     `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Resource             *Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	Building             *Building `protobuf:"bytes,3,opt,name=building,proto3" json:"building,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Zone) Reset()         { *m = Zone{} }
func (m *Zone) String() string { return proto.CompactTextString(m) }
func (*Zone) ProtoMessage()    {}
func (*Zone) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{7}
}

func (m *Zone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Zone.Unmarshal(m, b)
}
func (m *Zone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Zone.Marshal(b, m, deterministic)
}
func (m *Zone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Zone.Merge(m, src)
}
func (m *Zone) XXX_Size() int {
	return xxx_messageInfo_Zone.Size(m)
}
func (m *Zone) XXX_DiscardUnknown() {
	xxx_messageInfo_Zone.DiscardUnknown(m)
}

var xxx_messageInfo_Zone proto.InternalMessageInfo

func (m *Zone) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Zone) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Zone) GetBuilding() *Building {
	if m != nil {
		return m.Building
	}
	return nil
}

type Resource struct {
	Ph                   int32    `protobuf:"varint,1,opt,name=ph,proto3" json:"ph,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{8}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetPh() int32 {
	if m != nil {
		return m.Ph
	}
	return 0
}

func (m *Resource) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type Building struct {
	Ph                   int32    `protobuf:"varint,1,opt,name=ph,proto3" json:"ph,omitempty"`
	Type                 int32    `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Width                int32    `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	Height               int32    `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Building) Reset()         { *m = Building{} }
func (m *Building) String() string { return proto.CompactTextString(m) }
func (*Building) ProtoMessage()    {}
func (*Building) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{9}
}

func (m *Building) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Building.Unmarshal(m, b)
}
func (m *Building) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Building.Marshal(b, m, deterministic)
}
func (m *Building) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Building.Merge(m, src)
}
func (m *Building) XXX_Size() int {
	return xxx_messageInfo_Building.Size(m)
}
func (m *Building) XXX_DiscardUnknown() {
	xxx_messageInfo_Building.DiscardUnknown(m)
}

var xxx_messageInfo_Building proto.InternalMessageInfo

func (m *Building) GetPh() int32 {
	if m != nil {
		return m.Ph
	}
	return 0
}

func (m *Building) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Building) GetWidth() int32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Building) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

// Messages liés à click
type ClickRequest struct {
	Coordinates          *Coordinates `protobuf:"bytes,1,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ClickRequest) Reset()         { *m = ClickRequest{} }
func (m *ClickRequest) String() string { return proto.CompactTextString(m) }
func (*ClickRequest) ProtoMessage()    {}
func (*ClickRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{10}
}

func (m *ClickRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickRequest.Unmarshal(m, b)
}
func (m *ClickRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickRequest.Marshal(b, m, deterministic)
}
func (m *ClickRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickRequest.Merge(m, src)
}
func (m *ClickRequest) XXX_Size() int {
	return xxx_messageInfo_ClickRequest.Size(m)
}
func (m *ClickRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClickRequest proto.InternalMessageInfo

func (m *ClickRequest) GetCoordinates() *Coordinates {
	if m != nil {
		return m.Coordinates
	}
	return nil
}

type ClickReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClickReply) Reset()         { *m = ClickReply{} }
func (m *ClickReply) String() string { return proto.CompactTextString(m) }
func (*ClickReply) ProtoMessage()    {}
func (*ClickReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{11}
}

func (m *ClickReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClickReply.Unmarshal(m, b)
}
func (m *ClickReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClickReply.Marshal(b, m, deterministic)
}
func (m *ClickReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClickReply.Merge(m, src)
}
func (m *ClickReply) XXX_Size() int {
	return xxx_messageInfo_ClickReply.Size(m)
}
func (m *ClickReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ClickReply.DiscardUnknown(m)
}

var xxx_messageInfo_ClickReply proto.InternalMessageInfo

// Message liés à moveTo
type MoveToRequest struct {
	Actuel               *Coordinates `protobuf:"bytes,1,opt,name=actuel,proto3" json:"actuel,omitempty"`
	Destionation         *Coordinates `protobuf:"bytes,2,opt,name=destionation,proto3" json:"destionation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MoveToRequest) Reset()         { *m = MoveToRequest{} }
func (m *MoveToRequest) String() string { return proto.CompactTextString(m) }
func (*MoveToRequest) ProtoMessage()    {}
func (*MoveToRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{12}
}

func (m *MoveToRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveToRequest.Unmarshal(m, b)
}
func (m *MoveToRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveToRequest.Marshal(b, m, deterministic)
}
func (m *MoveToRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveToRequest.Merge(m, src)
}
func (m *MoveToRequest) XXX_Size() int {
	return xxx_messageInfo_MoveToRequest.Size(m)
}
func (m *MoveToRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveToRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MoveToRequest proto.InternalMessageInfo

func (m *MoveToRequest) GetActuel() *Coordinates {
	if m != nil {
		return m.Actuel
	}
	return nil
}

func (m *MoveToRequest) GetDestionation() *Coordinates {
	if m != nil {
		return m.Destionation
	}
	return nil
}

type MoveToReply struct {
	Chemin               []*Coordinates `protobuf:"bytes,1,rep,name=chemin,proto3" json:"chemin,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MoveToReply) Reset()         { *m = MoveToReply{} }
func (m *MoveToReply) String() string { return proto.CompactTextString(m) }
func (*MoveToReply) ProtoMessage()    {}
func (*MoveToReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{13}
}

func (m *MoveToReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveToReply.Unmarshal(m, b)
}
func (m *MoveToReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveToReply.Marshal(b, m, deterministic)
}
func (m *MoveToReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveToReply.Merge(m, src)
}
func (m *MoveToReply) XXX_Size() int {
	return xxx_messageInfo_MoveToReply.Size(m)
}
func (m *MoveToReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveToReply.DiscardUnknown(m)
}

var xxx_messageInfo_MoveToReply proto.InternalMessageInfo

func (m *MoveToReply) GetChemin() []*Coordinates {
	if m != nil {
		return m.Chemin
	}
	return nil
}

type Coordinates struct {
	X                    int32    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int32    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coordinates) Reset()         { *m = Coordinates{} }
func (m *Coordinates) String() string { return proto.CompactTextString(m) }
func (*Coordinates) ProtoMessage()    {}
func (*Coordinates) Descriptor() ([]byte, []int) {
	return fileDescriptor_77a6da22d6a3feb1, []int{14}
}

func (m *Coordinates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinates.Unmarshal(m, b)
}
func (m *Coordinates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinates.Marshal(b, m, deterministic)
}
func (m *Coordinates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinates.Merge(m, src)
}
func (m *Coordinates) XXX_Size() int {
	return xxx_messageInfo_Coordinates.Size(m)
}
func (m *Coordinates) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinates.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinates proto.InternalMessageInfo

func (m *Coordinates) GetX() int32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Coordinates) GetY() int32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*Player)(nil), "communication.Player")
	proto.RegisterType((*GetMapRequest)(nil), "communication.getMapRequest")
	proto.RegisterType((*GetMapReply)(nil), "communication.getMapReply")
	proto.RegisterType((*SetMapRequest)(nil), "communication.setMapRequest")
	proto.RegisterType((*SetMapReply)(nil), "communication.setMapReply")
	proto.RegisterType((*UpdateMapRequest)(nil), "communication.updateMapRequest")
	proto.RegisterType((*UpdateMapReply)(nil), "communication.updateMapReply")
	proto.RegisterType((*Zone)(nil), "communication.Zone")
	proto.RegisterType((*Resource)(nil), "communication.Resource")
	proto.RegisterType((*Building)(nil), "communication.Building")
	proto.RegisterType((*ClickRequest)(nil), "communication.clickRequest")
	proto.RegisterType((*ClickReply)(nil), "communication.clickReply")
	proto.RegisterType((*MoveToRequest)(nil), "communication.moveToRequest")
	proto.RegisterType((*MoveToReply)(nil), "communication.moveToReply")
	proto.RegisterType((*Coordinates)(nil), "communication.Coordinates")
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor_77a6da22d6a3feb1) }

var fileDescriptor_77a6da22d6a3feb1 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x9e, 0xdb, 0x25, 0xb4, 0x2f, 0xed, 0x34, 0x19, 0x04, 0xa1, 0x0c, 0x51, 0x45, 0x42, 0x2a,
	0x1c, 0x7a, 0xc8, 0x38, 0x22, 0xa4, 0x31, 0x40, 0x42, 0x62, 0x02, 0x59, 0x5c, 0x98, 0xb8, 0x64,
	0x89, 0xb7, 0x5a, 0x24, 0x76, 0x88, 0x1d, 0xb6, 0xde, 0x90, 0xe0, 0x7f, 0xe0, 0xdf, 0xdb, 0x9f,
	0x82, 0xec, 0xfc, 0x68, 0x92, 0x75, 0x65, 0x1c, 0x10, 0x37, 0xfb, 0xbd, 0xcf, 0x9f, 0xbf, 0x97,
	0xf7, 0x3d, 0x07, 0x86, 0x59, 0x1a, 0xce, 0xd3, 0x4c, 0x28, 0x81, 0xc7, 0xa1, 0x48, 0x92, 0x9c,
	0xb3, 0x30, 0x50, 0x4c, 0x70, 0xef, 0x12, 0x81, 0xfd, 0x21, 0x0e, 0x96, 0x34, 0xc3, 0x3b, 0xd0,
	0x63, 0x91, 0x8b, 0xa6, 0x68, 0x36, 0x22, 0x3d, 0x16, 0xe1, 0x3d, 0x18, 0x32, 0xf9, 0x26, 0x08,
	0x35, 0xce, 0xed, 0x4d, 0xd1, 0x6c, 0x40, 0x56, 0x01, 0x8c, 0x61, 0x9b, 0x07, 0x09, 0x75, 0xfb,
	0x53, 0x34, 0x1b, 0x12, 0xb3, 0xc6, 0x4f, 0x61, 0x97, 0xe7, 0xc9, 0x09, 0xcd, 0xde, 0x9f, 0xbe,
	0xcc, 0x59, 0x1c, 0x31, 0x7e, 0xe6, 0x6e, 0x4f, 0xd1, 0xcc, 0x22, 0x57, 0xe2, 0x4d, 0xec, 0x6b,
	0xae, 0x98, 0x62, 0x54, 0xba, 0x56, 0x1b, 0x5b, 0xc5, 0xf1, 0x1d, 0xb0, 0xa4, 0x12, 0x9c, 0xba,
	0xb6, 0x01, 0x14, 0x1b, 0xad, 0xe0, 0x5c, 0x88, 0xc8, 0xbd, 0x65, 0x82, 0x66, 0xad, 0x63, 0xa7,
	0x3a, 0x36, 0x28, 0x62, 0x7a, 0xed, 0x7d, 0x82, 0xf1, 0x19, 0x55, 0x47, 0x41, 0x4a, 0xe8, 0xd7,
	0x9c, 0x4a, 0x85, 0x47, 0x80, 0x2e, 0x4c, 0x9d, 0x16, 0x41, 0x17, 0x7a, 0xb7, 0x34, 0xe5, 0x59,
	0x04, 0x2d, 0xf5, 0x55, 0xe7, 0x2c, 0x52, 0x0b, 0x53, 0x97, 0x45, 0x8a, 0x0d, 0xbe, 0x0b, 0xf6,
	0x82, 0xb2, 0xb3, 0x85, 0x2a, 0xcb, 0x29, 0x77, 0xde, 0x33, 0x70, 0x2a, 0xea, 0x34, 0x5e, 0xe2,
	0xc7, 0xd0, 0x4f, 0x82, 0xd4, 0x45, 0xd3, 0xfe, 0xcc, 0xf1, 0x6f, 0xcf, 0x5b, 0x5f, 0x7a, 0x7e,
	0x2c, 0x38, 0x25, 0x3a, 0xef, 0x7d, 0x47, 0x30, 0x96, 0xff, 0x46, 0x51, 0x25, 0xc1, 0xfa, 0x83,
	0x84, 0x31, 0x38, 0x72, 0x25, 0xdc, 0xfb, 0x81, 0x60, 0x37, 0x4f, 0xa3, 0x40, 0xd1, 0xff, 0x29,
	0x6a, 0x17, 0x76, 0x1a, 0x22, 0xb4, 0xae, 0x9f, 0x08, 0xb6, 0x8f, 0xcb, 0x5e, 0xab, 0x65, 0x4a,
	0x4b, 0x39, 0x66, 0x8d, 0xf7, 0x61, 0x90, 0x51, 0x29, 0xf2, 0x2c, 0xa4, 0x46, 0x98, 0xe3, 0xdf,
	0xeb, 0x50, 0x93, 0x32, 0x4d, 0x6a, 0xa0, 0x3e, 0x74, 0x52, 0x59, 0xb3, 0xbf, 0xf6, 0x50, 0xe5,
	0x50, 0x52, 0x03, 0xbd, 0x39, 0x0c, 0x2a, 0x2a, 0x3d, 0x25, 0xe9, 0xa2, 0xd4, 0xd1, 0x4b, 0x17,
	0xb5, 0xb2, 0xde, 0x4a, 0x99, 0xf7, 0x19, 0x06, 0xb5, 0xcf, 0x6f, 0x80, 0xff, 0x4b, 0xd3, 0xbd,
	0x83, 0x51, 0x18, 0xb3, 0xf0, 0x4b, 0xd5, 0xa7, 0xe7, 0xe0, 0x84, 0x42, 0x64, 0x11, 0xe3, 0x81,
	0xa2, 0xd2, 0x5c, 0xe5, 0xf8, 0x93, 0x4e, 0x55, 0x87, 0x2b, 0x04, 0x69, 0xc2, 0xbd, 0x11, 0x40,
	0xc9, 0x56, 0x1a, 0x61, 0x9c, 0x88, 0x6f, 0xf4, 0xa3, 0xa8, 0xd8, 0x7d, 0xb0, 0x83, 0x50, 0xe5,
	0x34, 0xbe, 0x01, 0x71, 0x89, 0xc4, 0x2f, 0x60, 0x14, 0x51, 0xa9, 0xd3, 0x41, 0xfd, 0x78, 0x6c,
	0x3e, 0xd9, 0xc2, 0x7b, 0x07, 0xe0, 0x54, 0x22, 0xf4, 0x58, 0xf9, 0x60, 0x87, 0x0b, 0x9a, 0x30,
	0x5e, 0x4e, 0xd6, 0x46, 0x09, 0x05, 0xd2, 0x7b, 0x02, 0x4e, 0x23, 0xbc, 0xc9, 0xcb, 0xfe, 0x00,
	0xec, 0xc3, 0x98, 0x51, 0xae, 0xfc, 0x4b, 0x04, 0xfd, 0xa3, 0x20, 0xc5, 0xaf, 0xc0, 0x2e, 0xc6,
	0x1a, 0xef, 0x75, 0xae, 0x6a, 0x3d, 0x24, 0x93, 0xc9, 0x35, 0x59, 0xfd, 0x25, 0xb7, 0x34, 0x8b,
	0x5c, 0xcf, 0x22, 0x37, 0xb2, 0xc8, 0x16, 0xcb, 0x11, 0x0c, 0xeb, 0xa1, 0xc0, 0x8f, 0x3a, 0xd0,
	0xee, 0xcc, 0x4e, 0x1e, 0x5e, 0x0f, 0x30, 0x74, 0xfe, 0x2f, 0x04, 0xa3, 0xb7, 0x5c, 0xd1, 0xac,
	0x78, 0xc6, 0x25, 0x3e, 0x00, 0xcb, 0xf4, 0x1f, 0x3f, 0xe8, 0x1c, 0x6d, 0x7a, 0x6c, 0x72, 0x7f,
	0x7d, 0xb2, 0x2e, 0xb4, 0x68, 0xd7, 0x95, 0x42, 0x5b, 0x56, 0xba, 0x52, 0x68, 0xa3, 0xc7, 0xde,
	0xd6, 0x89, 0x6d, 0xfe, 0x4f, 0xfb, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x19, 0xeb, 0x82,
	0xac, 0x06, 0x00, 0x00,
}