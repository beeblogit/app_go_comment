package comment

import (
	"context"
	//"errors"

	"github.com/ncostamagna/go_lib_response/response"
	//"github.com/ncostamagna/gocourse_meta/meta"
)

//Endpoints struct
type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		Create Controller
	}

	CreateReq struct {
		UserID string `json:"user_id"`
		PostID string `json:"post_id"`
		Name string `json:"name"`
		Comment string `json:"comment"`
	}

	Config struct {
		LimPageDef string
	}
)

//MakeEndpoints handler endpoints
func MakeEndpoints(s Service, config Config) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		// Get:    makeGetEndpoint(s),
		// GetAll: makeGetAllEndpoint(s, config),
		// Update: makeUpdateEndpoint(s),
		// Delete: makeDeleteEndpoint(s),
	}
}

func makeCreateEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateReq)

		if req.Name == "" {
			return nil, response.BadRequest("")
		}

		if req.Comment == "" {
			return nil, response.BadRequest("")
		}

		if req.PostID == "" {
			return nil, response.BadRequest("")
		}

		comment, err := s.Create(ctx, req.UserID, req.PostID, req.Name, req.Comment)
		if err != nil {

			/*if err == ErrEndLesserStart ||
				err == ErrInvalidStartDate ||
				err == ErrInvalidEndDate {
				return nil, response.BadRequest(err.Error())
			}*/

			return nil, response.InternalServerError(err.Error())
		}

		return response.Created("success", comment, nil), nil
	}
}