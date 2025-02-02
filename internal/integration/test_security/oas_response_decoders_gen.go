// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/ogen-go/ogen/validate"
)

func decodeDisjointSecurityResponse(resp *http.Response) (res *DisjointSecurityOK, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &DisjointSecurityOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeIntersectSecurityResponse(resp *http.Response) (res *IntersectSecurityOK, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &IntersectSecurityOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}

func decodeOptionalSecurityResponse(resp *http.Response) (res *OptionalSecurityOK, _ error) {
	switch resp.StatusCode {
	case 200:
		// Code 200.
		return &OptionalSecurityOK{}, nil
	}
	return res, validate.UnexpectedStatusCode(resp.StatusCode)
}
