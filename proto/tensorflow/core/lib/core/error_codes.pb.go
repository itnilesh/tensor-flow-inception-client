// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/lib/core/error_codes.proto

package tensorflow_error

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The canonical error codes for TensorFlow APIs.
//
// Warnings:
//
// -   Do not change any numeric assignments.
// -   Changes to this list should only be made if there is a compelling
//     need that can't be satisfied in another way.  Such changes
//     must be approved by at least two OWNERS.
//
// Sometimes multiple error codes may apply.  Services should return
// the most specific error code that applies.  For example, prefer
// OUT_OF_RANGE over FAILED_PRECONDITION if both codes apply.
// Similarly prefer NOT_FOUND or ALREADY_EXISTS over FAILED_PRECONDITION.
type Code int32

const (
	// Not an error; returned on success
	Code_OK Code = 0
	// The operation was cancelled (typically by the caller).
	Code_CANCELLED Code = 1
	// Unknown error.  An example of where this error may be returned is
	// if a Status value received from another address space belongs to
	// an error-space that is not known in this address space.  Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	Code_UNKNOWN Code = 2
	// Client specified an invalid argument.  Note that this differs
	// from FAILED_PRECONDITION.  INVALID_ARGUMENT indicates arguments
	// that are problematic regardless of the state of the system
	// (e.g., a malformed file name).
	Code_INVALID_ARGUMENT Code = 3
	// Deadline expired before operation could complete.  For operations
	// that change the state of the system, this error may be returned
	// even if the operation has completed successfully.  For example, a
	// successful response from a server could have been delayed long
	// enough for the deadline to expire.
	Code_DEADLINE_EXCEEDED Code = 4
	// Some requested entity (e.g., file or directory) was not found.
	// For privacy reasons, this code *may* be returned when the client
	// does not have the access right to the entity.
	Code_NOT_FOUND Code = 5
	// Some entity that we attempted to create (e.g., file or directory)
	// already exists.
	Code_ALREADY_EXISTS Code = 6
	// The caller does not have permission to execute the specified
	// operation.  PERMISSION_DENIED must not be used for rejections
	// caused by exhausting some resource (use RESOURCE_EXHAUSTED
	// instead for those errors).  PERMISSION_DENIED must not be
	// used if the caller can not be identified (use UNAUTHENTICATED
	// instead for those errors).
	Code_PERMISSION_DENIED Code = 7
	// The request does not have valid authentication credentials for the
	// operation.
	Code_UNAUTHENTICATED Code = 16
	// Some resource has been exhausted, perhaps a per-user quota, or
	// perhaps the entire file system is out of space.
	Code_RESOURCE_EXHAUSTED Code = 8
	// Operation was rejected because the system is not in a state
	// required for the operation's execution.  For example, directory
	// to be deleted may be non-empty, an rmdir operation is applied to
	// a non-directory, etc.
	//
	// A litmus test that may help a service implementor in deciding
	// between FAILED_PRECONDITION, ABORTED, and UNAVAILABLE:
	//  (a) Use UNAVAILABLE if the client can retry just the failing call.
	//  (b) Use ABORTED if the client should retry at a higher-level
	//      (e.g., restarting a read-modify-write sequence).
	//  (c) Use FAILED_PRECONDITION if the client should not retry until
	//      the system state has been explicitly fixed.  E.g., if an "rmdir"
	//      fails because the directory is non-empty, FAILED_PRECONDITION
	//      should be returned since the client should not retry unless
	//      they have first fixed up the directory by deleting files from it.
	//  (d) Use FAILED_PRECONDITION if the client performs conditional
	//      REST Get/Update/Delete on a resource and the resource on the
	//      server does not match the condition. E.g., conflicting
	//      read-modify-write on the same resource.
	Code_FAILED_PRECONDITION Code = 9
	// The operation was aborted, typically due to a concurrency issue
	// like sequencer check failures, transaction aborts, etc.
	//
	// See litmus test above for deciding between FAILED_PRECONDITION,
	// ABORTED, and UNAVAILABLE.
	Code_ABORTED Code = 10
	// Operation tried to iterate past the valid input range.  E.g., seeking or
	// reading past end of file.
	//
	// Unlike INVALID_ARGUMENT, this error indicates a problem that may
	// be fixed if the system state changes. For example, a 32-bit file
	// system will generate INVALID_ARGUMENT if asked to read at an
	// offset that is not in the range [0,2^32-1], but it will generate
	// OUT_OF_RANGE if asked to read from an offset past the current
	// file size.
	//
	// There is a fair bit of overlap between FAILED_PRECONDITION and
	// OUT_OF_RANGE.  We recommend using OUT_OF_RANGE (the more specific
	// error) when it applies so that callers who are iterating through
	// a space can easily look for an OUT_OF_RANGE error to detect when
	// they are done.
	Code_OUT_OF_RANGE Code = 11
	// Operation is not implemented or not supported/enabled in this service.
	Code_UNIMPLEMENTED Code = 12
	// Internal errors.  Means some invariants expected by underlying
	// system has been broken.  If you see one of these errors,
	// something is very broken.
	Code_INTERNAL Code = 13
	// The service is currently unavailable.  This is a most likely a
	// transient condition and may be corrected by retrying with
	// a backoff.
	//
	// See litmus test above for deciding between FAILED_PRECONDITION,
	// ABORTED, and UNAVAILABLE.
	Code_UNAVAILABLE Code = 14
	// Unrecoverable data loss or corruption.
	Code_DATA_LOSS Code = 15
	// An extra enum entry to prevent people from writing code that
	// fails to compile when a new code is added.
	//
	// Nobody should ever reference this enumeration entry. In particular,
	// if you write C++ code that switches on this enumeration, add a default:
	// case instead of a case that mentions this enumeration entry.
	//
	// Nobody should rely on the value (currently 20) listed here.  It
	// may change in the future.
	Code_DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_ Code = 20
)

var Code_name = map[int32]string{
	0:  "OK",
	1:  "CANCELLED",
	2:  "UNKNOWN",
	3:  "INVALID_ARGUMENT",
	4:  "DEADLINE_EXCEEDED",
	5:  "NOT_FOUND",
	6:  "ALREADY_EXISTS",
	7:  "PERMISSION_DENIED",
	16: "UNAUTHENTICATED",
	8:  "RESOURCE_EXHAUSTED",
	9:  "FAILED_PRECONDITION",
	10: "ABORTED",
	11: "OUT_OF_RANGE",
	12: "UNIMPLEMENTED",
	13: "INTERNAL",
	14: "UNAVAILABLE",
	15: "DATA_LOSS",
	20: "DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_",
}
var Code_value = map[string]int32{
	"OK":                  0,
	"CANCELLED":           1,
	"UNKNOWN":             2,
	"INVALID_ARGUMENT":    3,
	"DEADLINE_EXCEEDED":   4,
	"NOT_FOUND":           5,
	"ALREADY_EXISTS":      6,
	"PERMISSION_DENIED":   7,
	"UNAUTHENTICATED":     16,
	"RESOURCE_EXHAUSTED":  8,
	"FAILED_PRECONDITION": 9,
	"ABORTED":             10,
	"OUT_OF_RANGE":        11,
	"UNIMPLEMENTED":       12,
	"INTERNAL":            13,
	"UNAVAILABLE":         14,
	"DATA_LOSS":           15,
	"DO_NOT_USE_RESERVED_FOR_FUTURE_EXPANSION_USE_DEFAULT_IN_SWITCH_INSTEAD_": 20,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}
func (Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_error_codes_cf5b5c09495457a9, []int{0}
}

func init() {
	proto.RegisterEnum("tensorflow.error.Code", Code_name, Code_value)
}

func init() {
	proto.RegisterFile("tensorflow/core/lib/core/error_codes.proto", fileDescriptor_error_codes_cf5b5c09495457a9)
}

var fileDescriptor_error_codes_cf5b5c09495457a9 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xcb, 0x4e, 0x5b, 0x31,
	0x10, 0x86, 0x1b, 0xa0, 0x5c, 0x4c, 0x42, 0x06, 0x43, 0x2f, 0xcf, 0xc0, 0x22, 0xa8, 0xea, 0x13,
	0x4c, 0xce, 0xcc, 0x49, 0xac, 0x38, 0xe3, 0x23, 0x5f, 0x02, 0x5d, 0x59, 0x05, 0x0e, 0x55, 0x55,
	0x5a, 0x57, 0x27, 0x48, 0xbc, 0x40, 0x1f, 0xba, 0xcb, 0xca, 0xe9, 0xa2, 0xdd, 0x59, 0xf6, 0xef,
	0x99, 0xff, 0xd3, 0xa7, 0xae, 0x9e, 0xfb, 0x1f, 0xdb, 0x32, 0x3c, 0x3e, 0x95, 0x97, 0xeb, 0xfb,
	0x32, 0xf4, 0xd7, 0x4f, 0x5f, 0xef, 0xfe, 0x1e, 0xfa, 0x61, 0x28, 0x43, 0xbe, 0x2f, 0x0f, 0xfd,
	0x76, 0xf6, 0x73, 0x28, 0xcf, 0x45, 0xc3, 0xbf, 0xec, 0x6c, 0xf7, 0x7a, 0xf5, 0x6b, 0x5f, 0x1d,
	0x34, 0xe5, 0xa1, 0xd7, 0x87, 0x6a, 0xcf, 0xad, 0xe0, 0x95, 0x9e, 0xa8, 0x93, 0x06, 0xa5, 0x61,
	0x6b, 0x99, 0x60, 0xa4, 0x4f, 0xd5, 0x51, 0x92, 0x95, 0xb8, 0x1b, 0x81, 0x3d, 0x7d, 0xa9, 0xc0,
	0xc8, 0x06, 0xad, 0xa1, 0x8c, 0x7e, 0x91, 0xd6, 0x2c, 0x11, 0xf6, 0xf5, 0x1b, 0x75, 0x4e, 0x8c,
	0x64, 0x8d, 0x70, 0xe6, 0xdb, 0x86, 0x99, 0x98, 0xe0, 0xa0, 0x0e, 0x12, 0x17, 0x73, 0xeb, 0x92,
	0x10, 0xbc, 0xd6, 0x5a, 0x9d, 0xa1, 0xf5, 0x8c, 0xf4, 0x29, 0xf3, 0xad, 0x09, 0x31, 0xc0, 0x61,
	0xfd, 0xd9, 0xb1, 0x5f, 0x9b, 0x10, 0x8c, 0x93, 0x4c, 0x2c, 0x86, 0x09, 0x8e, 0xf4, 0x85, 0x9a,
	0x26, 0xc1, 0x14, 0x97, 0x2c, 0xd1, 0x34, 0x18, 0x99, 0x00, 0xf4, 0x5b, 0xa5, 0x3d, 0x07, 0x97,
	0x7c, 0x53, 0xb7, 0x2c, 0x31, 0x85, 0x7a, 0x7f, 0xac, 0xdf, 0xa9, 0x8b, 0x16, 0x8d, 0x65, 0xca,
	0x9d, 0xe7, 0xc6, 0x09, 0x99, 0x68, 0x9c, 0xc0, 0x49, 0x6d, 0x8e, 0x73, 0xe7, 0x6b, 0x4a, 0x69,
	0x50, 0x63, 0x97, 0x62, 0x76, 0x6d, 0xf6, 0x28, 0x0b, 0x86, 0x53, 0x7d, 0xae, 0x26, 0x49, 0xcc,
	0xba, 0xb3, 0x5c, 0x31, 0x98, 0x60, 0xac, 0xc7, 0xea, 0xd8, 0x48, 0x64, 0x2f, 0x68, 0x61, 0xa2,
	0xa7, 0xea, 0x34, 0x09, 0x6e, 0xd0, 0x58, 0x9c, 0x5b, 0x86, 0xb3, 0x0a, 0x44, 0x18, 0x31, 0x5b,
	0x17, 0x02, 0x4c, 0xf5, 0x4a, 0x2d, 0xc8, 0xe5, 0x8a, 0x98, 0x02, 0x67, 0xcf, 0x81, 0xfd, 0x86,
	0x29, 0xb7, 0xce, 0xe7, 0x36, 0xc5, 0xe4, 0x6b, 0xcd, 0x0e, 0x65, 0x87, 0x56, 0x13, 0xc4, 0x2d,
	0x26, 0x1b, 0xb3, 0x91, 0x1c, 0x6e, 0x4c, 0x6c, 0x96, 0xd9, 0x48, 0x88, 0x8c, 0x94, 0xe1, 0x72,
	0xfe, 0x41, 0xbd, 0x2f, 0xc3, 0x97, 0xd9, 0x7f, 0x7a, 0x1e, 0x87, 0xcf, 0xdf, 0xfb, 0x97, 0x32,
	0x7c, 0x9b, 0x03, 0x57, 0x53, 0x55, 0xd2, 0xb6, 0xab, 0x16, 0xb7, 0xdd, 0xe8, 0xf7, 0x68, 0x74,
	0x77, 0xb8, 0x53, 0xfa, 0xf1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xf7, 0x64, 0xd4, 0x00,
	0x02, 0x00, 0x00,
}
