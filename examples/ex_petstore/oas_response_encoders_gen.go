// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeCreatePetsResponse(response CreatePetsRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *CreatePetsCreated:
		w.WriteHeader(201)
		span.SetStatus(codes.Ok, http.StatusText(201))
		return nil

	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("/pets"+`: unexpected response type: %T`, response)
	}
}
func encodeListPetsResponse(response ListPetsRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Pets:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("/pets"+`: unexpected response type: %T`, response)
	}
}
func encodeShowPetByIdResponse(response ShowPetByIdRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *Pet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	case *ErrorStatusCode:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(response.StatusCode)
		st := http.StatusText(response.StatusCode)
		if response.StatusCode >= http.StatusBadRequest {
			span.SetStatus(codes.Error, st)
		} else {
			span.SetStatus(codes.Ok, st)
		}
		e := jx.GetEncoder()
		defer jx.PutEncoder(e)

		response.Response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}
		return nil

	default:
		return errors.Errorf("/pets/{petId}"+`: unexpected response type: %T`, response)
	}
}