package grpcsrv

import (
	"context"
	"errors"

	"github.com/Tiksy1/otus_hw-test/hw12_13_14_15_calendar/internal/app"
	"github.com/Tiksy1/otus_hw-test/hw12_13_14_15_calendar/internal/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type API struct {
	application *app.App
	UnimplementedEventServiceServer
}

func NewAPI(application *app.App) *API {
	return &API{application: application, UnimplementedEventServiceServer: UnimplementedEventServiceServer{}}
}

func (a *API) CreateEvent(ctx context.Context, event *Event) (*CreateEventResponse, error) {
	err := a.application.CreateEvent(ctx, toAppEvent(event))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &CreateEventResponse{}, nil
}

func (a *API) UpdateEvent(ctx context.Context, event *Event) (*UpdateEventResponse, error) {
	err := a.application.UpdateEvent(ctx, toAppEvent(event))
	if err != nil {
		if errors.Is(err, storage.ErrEventDoesNotExist) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &UpdateEventResponse{}, nil
}

func (a *API) RemoveEvent(ctx context.Context, eventID *EventID) (*RemoveEventResponse, error) {
	err := a.application.RemoveEvent(ctx, eventID.Id)
	if err != nil {
		if errors.Is(err, storage.ErrEventDoesNotExist) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &RemoveEventResponse{}, nil
}

func (a *API) Events(ctx context.Context, query *EventsQuery) (*EventsValues, error) {
	events, err := a.application.Events(ctx, query.From, query.To)
	if err != nil {
		if errors.Is(err, storage.ErrNoEvents) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbEvents := make([]*Event, len(events))
	for i, e := range events {
		pbEvents[i] = toPBEvent(e)
	}

	return &EventsValues{Events: pbEvents}, nil
}

func toAppEvent(event *Event) app.Event {
	return app.Event{
		ID:          event.Id,
		Title:       event.Title,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		Description: event.Description,
		OwnerID:     event.OwnerId,
		RemindIn:    event.RemindIn,
	}
}

func toPBEvent(event app.Event) *Event {
	return &Event{
		Id:          event.ID,
		Title:       event.Title,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		Description: event.Description,
		OwnerId:     event.OwnerID,
		RemindIn:    event.RemindIn,
	}
}
