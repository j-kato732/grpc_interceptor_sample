// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/aimo.proto

package proto

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetPeriodRequest) Validate() error {
	return nil
}
func (this *GetPeriodResponse) Validate() error {
	if this.Response != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Response); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Response", err)
		}
	}
	if this.Result != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Result); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Result", err)
		}
	}
	return nil
}
func (this *DefaultResponse) Validate() error {
	return nil
}
func (this *Result) Validate() error {
	for _, item := range this.Period {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Period", err)
			}
		}
	}
	return nil
}
func (this *Period) Validate() error {
	return nil
}
func (this *GetUserInfoRequest) Validate() error {
	if this.Period == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Period", fmt.Errorf(`value '%v' must not be an empty string`, this.Period))
	}
	return nil
}
func (this *GetUserInfoResponse) Validate() error {
	if this.Response != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Response); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Response", err)
		}
	}
	if this.Result != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Result); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Result", err)
		}
	}
	return nil
}
func (this *GetUserInfoResult) Validate() error {
	if this.UserInfo != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UserInfo); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UserInfo", err)
		}
	}
	return nil
}
func (this *UserInfo) Validate() error {
	return nil
}