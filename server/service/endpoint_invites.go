package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/fleetdm/fleet/server/kolide"
)

type createInviteRequest struct {
	payload kolide.InvitePayload
}

type createInviteResponse struct {
	Invite *kolide.Invite `json:"invite,omitempty"`
	Err    error          `json:"error,omitempty"`
}

func (r createInviteResponse) error() error { return r.Err }

func makeCreateInviteEndpoint(svc kolide.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createInviteRequest)
		invite, err := svc.InviteNewUser(ctx, req.payload)
		if err != nil {
			return createInviteResponse{Err: err}, nil
		}
		return createInviteResponse{invite, nil}, nil
	}
}

type listInvitesRequest struct {
	ListOptions kolide.ListOptions
}

type listInvitesResponse struct {
	Invites []kolide.Invite `json:"invites"`
	Err     error           `json:"error,omitempty"`
}

func (r listInvitesResponse) error() error { return r.Err }

func makeListInvitesEndpoint(svc kolide.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(listInvitesRequest)
		invites, err := svc.ListInvites(ctx, req.ListOptions)
		if err != nil {
			return listInvitesResponse{Err: err}, nil
		}

		resp := listInvitesResponse{Invites: []kolide.Invite{}}
		for _, invite := range invites {
			resp.Invites = append(resp.Invites, *invite)
		}
		return resp, nil
	}
}

type deleteInviteRequest struct {
	ID uint
}

type deleteInviteResponse struct {
	Err error `json:"error,omitempty"`
}

func (r deleteInviteResponse) error() error { return r.Err }

func makeDeleteInviteEndpoint(svc kolide.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteInviteRequest)
		err := svc.DeleteInvite(ctx, req.ID)
		if err != nil {
			return deleteInviteResponse{Err: err}, nil
		}
		return deleteInviteResponse{}, nil
	}
}

type verifyInviteRequest struct {
	Token string
}

type verifyInviteResponse struct {
	Invite *kolide.Invite `json:"invite"`
	Err    error          `json:"error,omitempty"`
}

func (r verifyInviteResponse) error() error { return r.Err }

func makeVerifyInviteEndpoint(svc kolide.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(verifyInviteRequest)
		invite, err := svc.VerifyInvite(ctx, req.Token)
		if err != nil {
			return verifyInviteResponse{Err: err}, nil
		}
		return verifyInviteResponse{Invite: invite}, nil
	}
}
