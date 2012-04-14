// Code generated by protoc-gen-go from "datastore.proto"
// DO NOT EDIT!

package datastore

import proto "goprotobuf.googlecode.com/hg/proto"
import "math"
import "os"


// Reference proto, math & os imports to suppress error if they are not otherwise used.
var _ = proto.GetString
var _ = math.Inf
var _ os.Error


type Property_Meaning int32

const (
	Property_BLOB             = 14
	Property_TEXT             = 15
	Property_BYTESTRING       = 16
	Property_ATOM_CATEGORY    = 1
	Property_ATOM_LINK        = 2
	Property_ATOM_TITLE       = 3
	Property_ATOM_CONTENT     = 4
	Property_ATOM_SUMMARY     = 5
	Property_ATOM_AUTHOR      = 6
	Property_GD_WHEN          = 7
	Property_GD_EMAIL         = 8
	Property_GEORSS_POINT     = 9
	Property_GD_IM            = 10
	Property_GD_PHONENUMBER   = 11
	Property_GD_POSTALADDRESS = 12
	Property_GD_RATING        = 13
	Property_BLOBKEY          = 17
)

var Property_Meaning_name = map[int32]string{
	14: "BLOB",
	15: "TEXT",
	16: "BYTESTRING",
	1:  "ATOM_CATEGORY",
	2:  "ATOM_LINK",
	3:  "ATOM_TITLE",
	4:  "ATOM_CONTENT",
	5:  "ATOM_SUMMARY",
	6:  "ATOM_AUTHOR",
	7:  "GD_WHEN",
	8:  "GD_EMAIL",
	9:  "GEORSS_POINT",
	10: "GD_IM",
	11: "GD_PHONENUMBER",
	12: "GD_POSTALADDRESS",
	13: "GD_RATING",
	17: "BLOBKEY",
}
var Property_Meaning_value = map[string]int32{
	"BLOB":             14,
	"TEXT":             15,
	"BYTESTRING":       16,
	"ATOM_CATEGORY":    1,
	"ATOM_LINK":        2,
	"ATOM_TITLE":       3,
	"ATOM_CONTENT":     4,
	"ATOM_SUMMARY":     5,
	"ATOM_AUTHOR":      6,
	"GD_WHEN":          7,
	"GD_EMAIL":         8,
	"GEORSS_POINT":     9,
	"GD_IM":            10,
	"GD_PHONENUMBER":   11,
	"GD_POSTALADDRESS": 12,
	"GD_RATING":        13,
	"BLOBKEY":          17,
}

func NewProperty_Meaning(x int32) *Property_Meaning {
	e := Property_Meaning(x)
	return &e
}
func (x Property_Meaning) String() string {
	return proto.EnumName(Property_Meaning_name, int32(x))
}

type EntityProto_Kind int32

const (
	EntityProto_GD_CONTACT = 1
	EntityProto_GD_EVENT   = 2
	EntityProto_GD_MESSAGE = 3
)

var EntityProto_Kind_name = map[int32]string{
	1: "GD_CONTACT",
	2: "GD_EVENT",
	3: "GD_MESSAGE",
}
var EntityProto_Kind_value = map[string]int32{
	"GD_CONTACT": 1,
	"GD_EVENT":   2,
	"GD_MESSAGE": 3,
}

func NewEntityProto_Kind(x int32) *EntityProto_Kind {
	e := EntityProto_Kind(x)
	return &e
}
func (x EntityProto_Kind) String() string {
	return proto.EnumName(EntityProto_Kind_name, int32(x))
}

type Index_Property_Direction int32

const (
	Index_Property_ASCENDING  = 1
	Index_Property_DESCENDING = 2
)

var Index_Property_Direction_name = map[int32]string{
	1: "ASCENDING",
	2: "DESCENDING",
}
var Index_Property_Direction_value = map[string]int32{
	"ASCENDING":  1,
	"DESCENDING": 2,
}

func NewIndex_Property_Direction(x int32) *Index_Property_Direction {
	e := Index_Property_Direction(x)
	return &e
}
func (x Index_Property_Direction) String() string {
	return proto.EnumName(Index_Property_Direction_name, int32(x))
}

type CompositeIndex_State int32

const (
	CompositeIndex_WRITE_ONLY = 1
	CompositeIndex_READ_WRITE = 2
	CompositeIndex_DELETED    = 3
	CompositeIndex_ERROR      = 4
)

var CompositeIndex_State_name = map[int32]string{
	1: "WRITE_ONLY",
	2: "READ_WRITE",
	3: "DELETED",
	4: "ERROR",
}
var CompositeIndex_State_value = map[string]int32{
	"WRITE_ONLY": 1,
	"READ_WRITE": 2,
	"DELETED":    3,
	"ERROR":      4,
}

func NewCompositeIndex_State(x int32) *CompositeIndex_State {
	e := CompositeIndex_State(x)
	return &e
}
func (x CompositeIndex_State) String() string {
	return proto.EnumName(CompositeIndex_State_name, int32(x))
}

type Query_Hint int32

const (
	Query_ORDER_FIRST    = 1
	Query_ANCESTOR_FIRST = 2
	Query_FILTER_FIRST   = 3
)

var Query_Hint_name = map[int32]string{
	1: "ORDER_FIRST",
	2: "ANCESTOR_FIRST",
	3: "FILTER_FIRST",
}
var Query_Hint_value = map[string]int32{
	"ORDER_FIRST":    1,
	"ANCESTOR_FIRST": 2,
	"FILTER_FIRST":   3,
}

func NewQuery_Hint(x int32) *Query_Hint {
	e := Query_Hint(x)
	return &e
}
func (x Query_Hint) String() string {
	return proto.EnumName(Query_Hint_name, int32(x))
}

type Query_Filter_Operator int32

const (
	Query_Filter_LESS_THAN             = 1
	Query_Filter_LESS_THAN_OR_EQUAL    = 2
	Query_Filter_GREATER_THAN          = 3
	Query_Filter_GREATER_THAN_OR_EQUAL = 4
	Query_Filter_EQUAL                 = 5
	Query_Filter_IN                    = 6
	Query_Filter_EXISTS                = 7
)

var Query_Filter_Operator_name = map[int32]string{
	1: "LESS_THAN",
	2: "LESS_THAN_OR_EQUAL",
	3: "GREATER_THAN",
	4: "GREATER_THAN_OR_EQUAL",
	5: "EQUAL",
	6: "IN",
	7: "EXISTS",
}
var Query_Filter_Operator_value = map[string]int32{
	"LESS_THAN":             1,
	"LESS_THAN_OR_EQUAL":    2,
	"GREATER_THAN":          3,
	"GREATER_THAN_OR_EQUAL": 4,
	"EQUAL":                 5,
	"IN":                    6,
	"EXISTS":                7,
}

func NewQuery_Filter_Operator(x int32) *Query_Filter_Operator {
	e := Query_Filter_Operator(x)
	return &e
}
func (x Query_Filter_Operator) String() string {
	return proto.EnumName(Query_Filter_Operator_name, int32(x))
}

type Query_Order_Direction int32

const (
	Query_Order_ASCENDING  = 1
	Query_Order_DESCENDING = 2
)

var Query_Order_Direction_name = map[int32]string{
	1: "ASCENDING",
	2: "DESCENDING",
}
var Query_Order_Direction_value = map[string]int32{
	"ASCENDING":  1,
	"DESCENDING": 2,
}

func NewQuery_Order_Direction(x int32) *Query_Order_Direction {
	e := Query_Order_Direction(x)
	return &e
}
func (x Query_Order_Direction) String() string {
	return proto.EnumName(Query_Order_Direction_name, int32(x))
}

type Error_ErrorCode int32

const (
	Error_BAD_REQUEST                  = 1
	Error_CONCURRENT_TRANSACTION       = 2
	Error_INTERNAL_ERROR               = 3
	Error_NEED_INDEX                   = 4
	Error_TIMEOUT                      = 5
	Error_PERMISSION_DENIED            = 6
	Error_BIGTABLE_ERROR               = 7
	Error_COMMITTED_BUT_STILL_APPLYING = 8
	Error_CAPABILITY_DISABLED          = 9
	Error_TRY_ALTERNATE_BACKEND        = 10
)

var Error_ErrorCode_name = map[int32]string{
	1:  "BAD_REQUEST",
	2:  "CONCURRENT_TRANSACTION",
	3:  "INTERNAL_ERROR",
	4:  "NEED_INDEX",
	5:  "TIMEOUT",
	6:  "PERMISSION_DENIED",
	7:  "BIGTABLE_ERROR",
	8:  "COMMITTED_BUT_STILL_APPLYING",
	9:  "CAPABILITY_DISABLED",
	10: "TRY_ALTERNATE_BACKEND",
}
var Error_ErrorCode_value = map[string]int32{
	"BAD_REQUEST":                  1,
	"CONCURRENT_TRANSACTION":       2,
	"INTERNAL_ERROR":               3,
	"NEED_INDEX":                   4,
	"TIMEOUT":                      5,
	"PERMISSION_DENIED":            6,
	"BIGTABLE_ERROR":               7,
	"COMMITTED_BUT_STILL_APPLYING": 8,
	"CAPABILITY_DISABLED":          9,
	"TRY_ALTERNATE_BACKEND":        10,
}

func NewError_ErrorCode(x int32) *Error_ErrorCode {
	e := Error_ErrorCode(x)
	return &e
}
func (x Error_ErrorCode) String() string {
	return proto.EnumName(Error_ErrorCode_name, int32(x))
}

type Action struct {
	XXX_unrecognized []byte
}

func (this *Action) Reset() {
	*this = Action{}
}

type StringProto struct {
	Value            *string "PB(bytes,1,req,name=value)"
	XXX_unrecognized []byte
}

func (this *StringProto) Reset() {
	*this = StringProto{}
}

type Integer32Proto struct {
	Value            *int32 "PB(varint,1,req,name=value)"
	XXX_unrecognized []byte
}

func (this *Integer32Proto) Reset() {
	*this = Integer32Proto{}
}

type Integer64Proto struct {
	Value            *int64 "PB(varint,1,req,name=value)"
	XXX_unrecognized []byte
}

func (this *Integer64Proto) Reset() {
	*this = Integer64Proto{}
}

type BoolProto struct {
	Value            *bool "PB(varint,1,req,name=value)"
	XXX_unrecognized []byte
}

func (this *BoolProto) Reset() {
	*this = BoolProto{}
}

type DoubleProto struct {
	Value            *float64 "PB(fixed64,1,req,name=value)"
	XXX_unrecognized []byte
}

func (this *DoubleProto) Reset() {
	*this = DoubleProto{}
}

type BytesProto struct {
	Value            []byte "PB(bytes,1,req,name=value)"
	XXX_unrecognized []byte
}

func (this *BytesProto) Reset() {
	*this = BytesProto{}
}

type VoidProto struct {
	XXX_unrecognized []byte
}

func (this *VoidProto) Reset() {
	*this = VoidProto{}
}

type PropertyValue struct {
	Int64Value       *int64                        "PB(varint,1,opt,name=int64Value)"
	BooleanValue     *bool                         "PB(varint,2,opt,name=booleanValue)"
	StringValue      *string                       "PB(bytes,3,opt,name=stringValue)"
	DoubleValue      *float64                      "PB(fixed64,4,opt,name=doubleValue)"
	Pointvalue       *PropertyValue_PointValue     "PB(group,5,opt,name=PointValue)"
	Uservalue        *PropertyValue_UserValue      "PB(group,8,opt,name=UserValue)"
	Referencevalue   *PropertyValue_ReferenceValue "PB(group,12,opt,name=ReferenceValue)"
	XXX_unrecognized []byte
}

func (this *PropertyValue) Reset() {
	*this = PropertyValue{}
}

type PropertyValue_PointValue struct {
	X                *float64 "PB(fixed64,6,req,name=x)"
	Y                *float64 "PB(fixed64,7,req,name=y)"
	XXX_unrecognized []byte
}

func (this *PropertyValue_PointValue) Reset() {
	*this = PropertyValue_PointValue{}
}

type PropertyValue_UserValue struct {
	Email             *string "PB(bytes,9,req,name=email)"
	AuthDomain        *string "PB(bytes,10,req,name=auth_domain)"
	Nickname          *string "PB(bytes,11,opt,name=nickname)"
	Gaiaid            *int64  "PB(varint,18,req,name=gaiaid)"
	ObfuscatedGaiaid  *string "PB(bytes,19,opt,name=obfuscated_gaiaid)"
	FederatedIdentity *string "PB(bytes,21,opt,name=federated_identity)"
	FederatedProvider *string "PB(bytes,22,opt,name=federated_provider)"
	XXX_unrecognized  []byte
}

func (this *PropertyValue_UserValue) Reset() {
	*this = PropertyValue_UserValue{}
}

type PropertyValue_ReferenceValue struct {
	App              *string                                     "PB(bytes,13,req,name=app)"
	NameSpace        *string                                     "PB(bytes,20,opt,name=name_space)"
	Pathelement      []*PropertyValue_ReferenceValue_PathElement "PB(group,14,rep,name=PathElement)"
	XXX_unrecognized []byte
}

func (this *PropertyValue_ReferenceValue) Reset() {
	*this = PropertyValue_ReferenceValue{}
}

type PropertyValue_ReferenceValue_PathElement struct {
	Type             *string "PB(bytes,15,req,name=type)"
	Id               *int64  "PB(varint,16,opt,name=id)"
	Name             *string "PB(bytes,17,opt,name=name)"
	XXX_unrecognized []byte
}

func (this *PropertyValue_ReferenceValue_PathElement) Reset() {
	*this = PropertyValue_ReferenceValue_PathElement{}
}

type Property struct {
	Meaning          *Property_Meaning "PB(varint,1,opt,name=meaning,enum=datastore.Property_Meaning)"
	MeaningUri       *string           "PB(bytes,2,opt,name=meaning_uri)"
	Name             *string           "PB(bytes,3,req,name=name)"
	Value            *PropertyValue    "PB(bytes,5,req,name=value)"
	Multiple         *bool             "PB(varint,4,req,name=multiple)"
	XXX_unrecognized []byte
}

func (this *Property) Reset() {
	*this = Property{}
}

type Path struct {
	Element          []*Path_Element "PB(group,1,rep,name=Element)"
	XXX_unrecognized []byte
}

func (this *Path) Reset() {
	*this = Path{}
}

type Path_Element struct {
	Type             *string "PB(bytes,2,req,name=type)"
	Id               *int64  "PB(varint,3,opt,name=id)"
	Name             *string "PB(bytes,4,opt,name=name)"
	XXX_unrecognized []byte
}

func (this *Path_Element) Reset() {
	*this = Path_Element{}
}

type Reference struct {
	App              *string "PB(bytes,13,req,name=app)"
	NameSpace        *string "PB(bytes,20,opt,name=name_space)"
	Path             *Path   "PB(bytes,14,req,name=path)"
	XXX_unrecognized []byte
}

func (this *Reference) Reset() {
	*this = Reference{}
}

type User struct {
	Email             *string "PB(bytes,1,req,name=email)"
	AuthDomain        *string "PB(bytes,2,req,name=auth_domain)"
	Nickname          *string "PB(bytes,3,opt,name=nickname)"
	Gaiaid            *int64  "PB(varint,4,req,name=gaiaid)"
	ObfuscatedGaiaid  *string "PB(bytes,5,opt,name=obfuscated_gaiaid)"
	FederatedIdentity *string "PB(bytes,6,opt,name=federated_identity)"
	FederatedProvider *string "PB(bytes,7,opt,name=federated_provider)"
	XXX_unrecognized  []byte
}

func (this *User) Reset() {
	*this = User{}
}

type EntityProto struct {
	Key              *Reference        "PB(bytes,13,req,name=key)"
	EntityGroup      *Path             "PB(bytes,16,req,name=entity_group)"
	Owner            *User             "PB(bytes,17,opt,name=owner)"
	Kind             *EntityProto_Kind "PB(varint,4,opt,name=kind,enum=datastore.EntityProto_Kind)"
	KindUri          *string           "PB(bytes,5,opt,name=kind_uri)"
	Property         []*Property       "PB(bytes,14,rep,name=property)"
	RawProperty      []*Property       "PB(bytes,15,rep,name=raw_property)"
	XXX_unrecognized []byte
}

func (this *EntityProto) Reset() {
	*this = EntityProto{}
}

type CompositeProperty struct {
	IndexId          *int64   "PB(varint,1,req,name=index_id)"
	Value            []string "PB(bytes,2,rep,name=value)"
	XXX_unrecognized []byte
}

func (this *CompositeProperty) Reset() {
	*this = CompositeProperty{}
}

type Index struct {
	EntityType       *string           "PB(bytes,1,req,name=entity_type)"
	Ancestor         *bool             "PB(varint,5,req,name=ancestor)"
	Property         []*Index_Property "PB(group,2,rep,name=Property)"
	XXX_unrecognized []byte
}

func (this *Index) Reset() {
	*this = Index{}
}

type Index_Property struct {
	Name             *string                   "PB(bytes,3,req,name=name)"
	Direction        *Index_Property_Direction "PB(varint,4,opt,name=direction,enum=datastore.Index_Property_Direction,def=1)"
	XXX_unrecognized []byte
}

func (this *Index_Property) Reset() {
	*this = Index_Property{}
}

const Default_Index_Property_Direction Index_Property_Direction = Index_Property_ASCENDING

type CompositeIndex struct {
	AppId            *string               "PB(bytes,1,req,name=app_id)"
	Id               *int64                "PB(varint,2,req,name=id)"
	Definition       *Index                "PB(bytes,3,req,name=definition)"
	State            *CompositeIndex_State "PB(varint,4,req,name=state,enum=datastore.CompositeIndex_State)"
	XXX_unrecognized []byte
}

func (this *CompositeIndex) Reset() {
	*this = CompositeIndex{}
}

type Transaction struct {
	Handle           *uint64 "PB(fixed64,1,req,name=handle)"
	App              *string "PB(bytes,2,req,name=app)"
	MarkChanges      *bool   "PB(varint,3,opt,name=mark_changes,def=0)"
	XXX_unrecognized []byte
}

func (this *Transaction) Reset() {
	*this = Transaction{}
}

const Default_Transaction_MarkChanges bool = false

type Query struct {
	App                *string           "PB(bytes,1,req,name=app)"
	NameSpace          *string           "PB(bytes,29,opt,name=name_space)"
	Kind               *string           "PB(bytes,3,opt,name=kind)"
	Ancestor           *Reference        "PB(bytes,17,opt,name=ancestor)"
	Filter             []*Query_Filter   "PB(group,4,rep,name=Filter)"
	SearchQuery        *string           "PB(bytes,8,opt,name=search_query)"
	Order              []*Query_Order    "PB(group,9,rep,name=Order)"
	Hint               *Query_Hint       "PB(varint,18,opt,name=hint,enum=datastore.Query_Hint)"
	Count              *int32            "PB(varint,23,opt,name=count)"
	Offset             *int32            "PB(varint,12,opt,name=offset,def=0)"
	Limit              *int32            "PB(varint,16,opt,name=limit)"
	CompiledCursor     *CompiledCursor   "PB(bytes,30,opt,name=compiled_cursor)"
	EndCompiledCursor  *CompiledCursor   "PB(bytes,31,opt,name=end_compiled_cursor)"
	CompositeIndex     []*CompositeIndex "PB(bytes,19,rep,name=composite_index)"
	RequirePerfectPlan *bool             "PB(varint,20,opt,name=require_perfect_plan,def=0)"
	KeysOnly           *bool             "PB(varint,21,opt,name=keys_only,def=0)"
	Transaction        *Transaction      "PB(bytes,22,opt,name=transaction)"
	Distinct           *bool             "PB(varint,24,opt,name=distinct)"
	Compile            *bool             "PB(varint,25,opt,name=compile,def=0)"
	FailoverMs         *int64            "PB(varint,26,opt,name=failover_ms)"
	Strong             *bool             "PB(varint,32,opt,name=strong)"
	XXX_unrecognized   []byte
}

func (this *Query) Reset() {
	*this = Query{}
}

const Default_Query_Offset int32 = 0
const Default_Query_RequirePerfectPlan bool = false
const Default_Query_KeysOnly bool = false
const Default_Query_Compile bool = false

type Query_Filter struct {
	Op               *Query_Filter_Operator "PB(varint,6,req,name=op,enum=datastore.Query_Filter_Operator)"
	Property         []*Property            "PB(bytes,14,rep,name=property)"
	XXX_unrecognized []byte
}

func (this *Query_Filter) Reset() {
	*this = Query_Filter{}
}

type Query_Order struct {
	Property         *string                "PB(bytes,10,req,name=property)"
	Direction        *Query_Order_Direction "PB(varint,11,opt,name=direction,enum=datastore.Query_Order_Direction,def=1)"
	XXX_unrecognized []byte
}

func (this *Query_Order) Reset() {
	*this = Query_Order{}
}

const Default_Query_Order_Direction Query_Order_Direction = Query_Order_ASCENDING

type CompiledQuery struct {
	Primaryscan      *CompiledQuery_PrimaryScan     "PB(group,1,req,name=PrimaryScan)"
	Mergejoinscan    []*CompiledQuery_MergeJoinScan "PB(group,7,rep,name=MergeJoinScan)"
	IndexDef         *Index                         "PB(bytes,21,opt,name=index_def)"
	Offset           *int32                         "PB(varint,10,opt,name=offset,def=0)"
	Limit            *int32                         "PB(varint,11,opt,name=limit)"
	KeysOnly         *bool                          "PB(varint,12,req,name=keys_only)"
	Entityfilter     *CompiledQuery_EntityFilter    "PB(group,13,opt,name=EntityFilter)"
	XXX_unrecognized []byte
}

func (this *CompiledQuery) Reset() {
	*this = CompiledQuery{}
}

const Default_CompiledQuery_Offset int32 = 0

type CompiledQuery_PrimaryScan struct {
	IndexName                  *string  "PB(bytes,2,opt,name=index_name)"
	StartKey                   *string  "PB(bytes,3,opt,name=start_key)"
	StartInclusive             *bool    "PB(varint,4,opt,name=start_inclusive)"
	EndKey                     *string  "PB(bytes,5,opt,name=end_key)"
	EndInclusive               *bool    "PB(varint,6,opt,name=end_inclusive)"
	StartPostfixValue          []string "PB(bytes,22,rep,name=start_postfix_value)"
	EndPostfixValue            []string "PB(bytes,23,rep,name=end_postfix_value)"
	EndUnappliedLogTimestampUs *int64   "PB(varint,19,opt,name=end_unapplied_log_timestamp_us)"
	XXX_unrecognized           []byte
}

func (this *CompiledQuery_PrimaryScan) Reset() {
	*this = CompiledQuery_PrimaryScan{}
}

type CompiledQuery_MergeJoinScan struct {
	IndexName        *string  "PB(bytes,8,req,name=index_name)"
	PrefixValue      []string "PB(bytes,9,rep,name=prefix_value)"
	ValuePrefix      *bool    "PB(varint,20,opt,name=value_prefix,def=0)"
	XXX_unrecognized []byte
}

func (this *CompiledQuery_MergeJoinScan) Reset() {
	*this = CompiledQuery_MergeJoinScan{}
}

const Default_CompiledQuery_MergeJoinScan_ValuePrefix bool = false

type CompiledQuery_EntityFilter struct {
	Distinct         *bool      "PB(varint,14,opt,name=distinct,def=0)"
	Kind             *string    "PB(bytes,17,opt,name=kind)"
	Ancestor         *Reference "PB(bytes,18,opt,name=ancestor)"
	XXX_unrecognized []byte
}

func (this *CompiledQuery_EntityFilter) Reset() {
	*this = CompiledQuery_EntityFilter{}
}

const Default_CompiledQuery_EntityFilter_Distinct bool = false

type CompiledCursor struct {
	MultiqueryIndex  *int32                     "PB(varint,1,opt,name=multiquery_index)"
	Position         []*CompiledCursor_Position "PB(group,2,rep,name=Position)"
	XXX_unrecognized []byte
}

func (this *CompiledCursor) Reset() {
	*this = CompiledCursor{}
}

type CompiledCursor_Position struct {
	StartKey         *string                               "PB(bytes,27,opt,name=start_key)"
	Indexvalue       []*CompiledCursor_Position_IndexValue "PB(group,29,rep,name=IndexValue)"
	Key              *Reference                            "PB(bytes,32,opt,name=key)"
	StartInclusive   *bool                                 "PB(varint,28,opt,name=start_inclusive,def=1)"
	XXX_unrecognized []byte
}

func (this *CompiledCursor_Position) Reset() {
	*this = CompiledCursor_Position{}
}

const Default_CompiledCursor_Position_StartInclusive bool = true

type CompiledCursor_Position_IndexValue struct {
	Property         *string        "PB(bytes,30,opt,name=property)"
	Value            *PropertyValue "PB(bytes,31,req,name=value)"
	XXX_unrecognized []byte
}

func (this *CompiledCursor_Position_IndexValue) Reset() {
	*this = CompiledCursor_Position_IndexValue{}
}

type RunCompiledQueryRequest struct {
	App              *string        "PB(bytes,5,req,name=app)"
	NameSpace        *string        "PB(bytes,6,opt,name=name_space)"
	CompiledQuery    *CompiledQuery "PB(bytes,1,req,name=compiled_query)"
	OriginalQuery    *Query         "PB(bytes,2,opt,name=original_query)"
	Count            *int32         "PB(varint,3,opt,name=count)"
	FailoverMs       *int64         "PB(varint,4,opt,name=failover_ms)"
	XXX_unrecognized []byte
}

func (this *RunCompiledQueryRequest) Reset() {
	*this = RunCompiledQueryRequest{}
}

type Cursor struct {
	Cursor           *uint64 "PB(fixed64,1,req,name=cursor)"
	App              *string "PB(bytes,2,opt,name=app)"
	XXX_unrecognized []byte
}

func (this *Cursor) Reset() {
	*this = Cursor{}
}

type Error struct {
	XXX_unrecognized []byte
}

func (this *Error) Reset() {
	*this = Error{}
}

type Cost struct {
	IndexWrites      *int32 "PB(varint,1,opt,name=index_writes)"
	IndexWriteBytes  *int32 "PB(varint,2,opt,name=index_write_bytes)"
	EntityWrites     *int32 "PB(varint,3,opt,name=entity_writes)"
	EntityWriteBytes *int32 "PB(varint,4,opt,name=entity_write_bytes)"
	XXX_unrecognized []byte
}

func (this *Cost) Reset() {
	*this = Cost{}
}

type GetRequest struct {
	Key              []*Reference "PB(bytes,1,rep,name=key)"
	Transaction      *Transaction "PB(bytes,2,opt,name=transaction)"
	FailoverMs       *int64       "PB(varint,3,opt,name=failover_ms)"
	Strong           *bool        "PB(varint,4,opt,name=strong)"
	XXX_unrecognized []byte
}

func (this *GetRequest) Reset() {
	*this = GetRequest{}
}

type GetResponse struct {
	Entity           []*GetResponse_Entity "PB(group,1,rep,name=Entity)"
	XXX_unrecognized []byte
}

func (this *GetResponse) Reset() {
	*this = GetResponse{}
}

type GetResponse_Entity struct {
	Entity           *EntityProto "PB(bytes,2,opt,name=entity)"
	XXX_unrecognized []byte
}

func (this *GetResponse_Entity) Reset() {
	*this = GetResponse_Entity{}
}

type PutRequest struct {
	Entity           []*EntityProto    "PB(bytes,1,rep,name=entity)"
	Transaction      *Transaction      "PB(bytes,2,opt,name=transaction)"
	CompositeIndex   []*CompositeIndex "PB(bytes,3,rep,name=composite_index)"
	Trusted          *bool             "PB(varint,4,opt,name=trusted,def=0)"
	Force            *bool             "PB(varint,7,opt,name=force,def=0)"
	MarkChanges      *bool             "PB(varint,8,opt,name=mark_changes,def=0)"
	XXX_unrecognized []byte
}

func (this *PutRequest) Reset() {
	*this = PutRequest{}
}

const Default_PutRequest_Trusted bool = false
const Default_PutRequest_Force bool = false
const Default_PutRequest_MarkChanges bool = false

type PutResponse struct {
	Key              []*Reference "PB(bytes,1,rep,name=key)"
	Cost             *Cost        "PB(bytes,2,opt,name=cost)"
	XXX_unrecognized []byte
}

func (this *PutResponse) Reset() {
	*this = PutResponse{}
}

type TouchRequest struct {
	Key              []*Reference      "PB(bytes,1,rep,name=key)"
	CompositeIndex   []*CompositeIndex "PB(bytes,2,rep,name=composite_index)"
	XXX_unrecognized []byte
}

func (this *TouchRequest) Reset() {
	*this = TouchRequest{}
}

type TouchResponse struct {
	Cost             *Cost "PB(bytes,1,opt,name=cost)"
	XXX_unrecognized []byte
}

func (this *TouchResponse) Reset() {
	*this = TouchResponse{}
}

type DeleteRequest struct {
	Key              []*Reference "PB(bytes,6,rep,name=key)"
	Transaction      *Transaction "PB(bytes,5,opt,name=transaction)"
	Trusted          *bool        "PB(varint,4,opt,name=trusted,def=0)"
	Force            *bool        "PB(varint,7,opt,name=force,def=0)"
	MarkChanges      *bool        "PB(varint,8,opt,name=mark_changes,def=0)"
	XXX_unrecognized []byte
}

func (this *DeleteRequest) Reset() {
	*this = DeleteRequest{}
}

const Default_DeleteRequest_Trusted bool = false
const Default_DeleteRequest_Force bool = false
const Default_DeleteRequest_MarkChanges bool = false

type DeleteResponse struct {
	Cost             *Cost "PB(bytes,1,opt,name=cost)"
	XXX_unrecognized []byte
}

func (this *DeleteResponse) Reset() {
	*this = DeleteResponse{}
}

type NextRequest struct {
	Cursor           *Cursor "PB(bytes,1,req,name=cursor)"
	Count            *int32  "PB(varint,2,opt,name=count)"
	Offset           *int32  "PB(varint,4,opt,name=offset,def=0)"
	Compile          *bool   "PB(varint,3,opt,name=compile,def=0)"
	XXX_unrecognized []byte
}

func (this *NextRequest) Reset() {
	*this = NextRequest{}
}

const Default_NextRequest_Offset int32 = 0
const Default_NextRequest_Compile bool = false

type QueryResult struct {
	Cursor           *Cursor         "PB(bytes,1,opt,name=cursor)"
	Result           []*EntityProto  "PB(bytes,2,rep,name=result)"
	SkippedResults   *int32          "PB(varint,7,opt,name=skipped_results)"
	MoreResults      *bool           "PB(varint,3,req,name=more_results)"
	KeysOnly         *bool           "PB(varint,4,opt,name=keys_only)"
	CompiledQuery    *CompiledQuery  "PB(bytes,5,opt,name=compiled_query)"
	CompiledCursor   *CompiledCursor "PB(bytes,6,opt,name=compiled_cursor)"
	XXX_unrecognized []byte
}

func (this *QueryResult) Reset() {
	*this = QueryResult{}
}

type GetSchemaRequest struct {
	App              *string "PB(bytes,1,req,name=app)"
	NameSpace        *string "PB(bytes,5,opt,name=name_space)"
	StartKind        *string "PB(bytes,2,opt,name=start_kind)"
	EndKind          *string "PB(bytes,3,opt,name=end_kind)"
	Properties       *bool   "PB(varint,4,opt,name=properties,def=1)"
	XXX_unrecognized []byte
}

func (this *GetSchemaRequest) Reset() {
	*this = GetSchemaRequest{}
}

const Default_GetSchemaRequest_Properties bool = true

type Schema struct {
	Kind             []*EntityProto "PB(bytes,1,rep,name=kind)"
	MoreResults      *bool          "PB(varint,2,opt,name=more_results,def=0)"
	XXX_unrecognized []byte
}

func (this *Schema) Reset() {
	*this = Schema{}
}

const Default_Schema_MoreResults bool = false

type GetNamespacesRequest struct {
	App              *string "PB(bytes,1,req,name=app)"
	StartNamespace   *string "PB(bytes,2,opt,name=start_namespace)"
	EndNamespace     *string "PB(bytes,3,opt,name=end_namespace)"
	XXX_unrecognized []byte
}

func (this *GetNamespacesRequest) Reset() {
	*this = GetNamespacesRequest{}
}

type GetNamespacesResponse struct {
	Namespace        []string "PB(bytes,1,rep,name=namespace)"
	MoreResults      *bool    "PB(varint,2,opt,name=more_results,def=0)"
	XXX_unrecognized []byte
}

func (this *GetNamespacesResponse) Reset() {
	*this = GetNamespacesResponse{}
}

const Default_GetNamespacesResponse_MoreResults bool = false

type AllocateIdsRequest struct {
	ModelKey         *Reference "PB(bytes,1,req,name=model_key)"
	Size             *int64     "PB(varint,2,opt,name=size)"
	Max              *int64     "PB(varint,3,opt,name=max)"
	XXX_unrecognized []byte
}

func (this *AllocateIdsRequest) Reset() {
	*this = AllocateIdsRequest{}
}

type AllocateIdsResponse struct {
	Start            *int64 "PB(varint,1,req,name=start)"
	End              *int64 "PB(varint,2,req,name=end)"
	XXX_unrecognized []byte
}

func (this *AllocateIdsResponse) Reset() {
	*this = AllocateIdsResponse{}
}

type CompositeIndices struct {
	Index            []*CompositeIndex "PB(bytes,1,rep,name=index)"
	XXX_unrecognized []byte
}

func (this *CompositeIndices) Reset() {
	*this = CompositeIndices{}
}

type AddActionsRequest struct {
	Transaction      *Transaction "PB(bytes,1,req,name=transaction)"
	Action           []*Action    "PB(bytes,2,rep,name=action)"
	XXX_unrecognized []byte
}

func (this *AddActionsRequest) Reset() {
	*this = AddActionsRequest{}
}

type AddActionsResponse struct {
	XXX_unrecognized []byte
}

func (this *AddActionsResponse) Reset() {
	*this = AddActionsResponse{}
}

type BeginTransactionRequest struct {
	App              *string "PB(bytes,1,req,name=app)"
	XXX_unrecognized []byte
}

func (this *BeginTransactionRequest) Reset() {
	*this = BeginTransactionRequest{}
}

type CommitResponse struct {
	Cost             *Cost "PB(bytes,1,opt,name=cost)"
	XXX_unrecognized []byte
}

func (this *CommitResponse) Reset() {
	*this = CommitResponse{}
}

func init() {
	proto.RegisterEnum("datastore.Property_Meaning", Property_Meaning_name, Property_Meaning_value)
	proto.RegisterEnum("datastore.EntityProto_Kind", EntityProto_Kind_name, EntityProto_Kind_value)
	proto.RegisterEnum("datastore.Index_Property_Direction", Index_Property_Direction_name, Index_Property_Direction_value)
	proto.RegisterEnum("datastore.CompositeIndex_State", CompositeIndex_State_name, CompositeIndex_State_value)
	proto.RegisterEnum("datastore.Query_Hint", Query_Hint_name, Query_Hint_value)
	proto.RegisterEnum("datastore.Query_Filter_Operator", Query_Filter_Operator_name, Query_Filter_Operator_value)
	proto.RegisterEnum("datastore.Query_Order_Direction", Query_Order_Direction_name, Query_Order_Direction_value)
	proto.RegisterEnum("datastore.Error_ErrorCode", Error_ErrorCode_name, Error_ErrorCode_value)
}
