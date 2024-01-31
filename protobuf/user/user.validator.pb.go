// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: protobuf/user/user.proto

package user

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Address) Validate() error {
	return nil
}
func (this *User) Validate() error {
	if this.Address != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Address); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Address", err)
		}
	}
	return nil
}
func (this *CreateUserRequest) Validate() error {
	if this.User != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.User); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("User", err)
		}
	}
	return nil
}