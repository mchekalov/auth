package converter

import (
	"github.com/mchekalov/auth/internal/model"
	desc "github.com/mchekalov/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// FromAPIToServUserInfo converts a CreateRequest object from the API to a model.UserInfo entity.
func FromAPIToServUserInfo(req *desc.CreateRequest) *model.UserInfo {
	return &model.UserInfo{
		Name:  req.Info.Name,
		Email: req.Info.Email,
		Role:  int64(req.Info.Role),
	}
}

// FromServToAPIID converts an ID from service layer to a CreateResponse object for the API.
func FromServToAPIID(in *model.UserID) *desc.CreateResponse {
	return &desc.CreateResponse{
		Id: in.UserID,
	}
}

// FromServToAPIUser converts an user from service layer to a GetResponse object for the API.
func FromServToAPIUser(in *model.User) *desc.GetResponse {
	var updatedAt *timestamppb.Timestamp
	if in.UpdatedAt.Valid {
		updatedAt = timestamppb.New(in.UpdatedAt.Time)
	}

	return &desc.GetResponse{
		User: &desc.User{
			Id: in.ID,
			Info: &desc.UserInfo{
				Name:            in.Name,
				Email:           in.Email,
				Password:        in.Password,
				PasswordConfirm: in.PasswordConfirm,
				Role:            desc.Role(in.Role),
			},
			CreatedAt: timestamppb.New(in.CreatedAt),
			UpdatedAt: updatedAt,
		},
	}
}

// FromAPIToServID conver from to
func FromAPIToServID(in *desc.DeleteRequest) *model.UserID {
	return &model.UserID{
		UserID: in.GetId(),
	}
}

// FromAPIToServIDGet conver from to
func FromAPIToServIDGet(in *desc.GetRequest) *model.UserID {
	return &model.UserID{
		UserID: in.GetId(),
	}
}

// FromAPIToServUpdate conver from to
func FromAPIToServUpdate(in *desc.UpdateRequest) *model.UpdateInfo {
	return &model.UpdateInfo{
		ID:    in.Wrap.Id,
		Name:  in.Wrap.Name.Value,
		Email: in.Wrap.Email.Value,
		Role:  int64(in.Wrap.Role),
	}
}
