package comment

import (
	"context"
	"github.com/ncostamagna/go_http_utils/meta"
	"github.com/ncostamagna/go_http_utils/response"
)

// Endpoints struct
type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		Create Controller
		GetAll Controller
	}

	GetAllReq struct {
		ID     []string `json:"id"`
		UserID []string `json:"user_id"`
		PostID []string `json:"post_id"`
		Limit  int      `json:"limit"`
		Page   int      `json:"page"`
	}

	CreateReq struct {
		UserID  string `json:"user_id"`
		PostID  string `json:"post_id"`
		Name    string `json:"name"`
		Comment string `json:"comment"`
	}

	Config struct {
		LimPageDef string
	}
)

// MakeEndpoints handler endpoints
func MakeEndpoints(s Service, config Config) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		// Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s, config),
		// Update: makeUpdateEndpoint(s),
		// Delete: makeDeleteEndpoint(s),
	}
}

func makeCreateEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateReq)

		if req.Name == "" {
			return nil, response.BadRequest(ErrNameRequired.Error())
		}

		if req.Comment == "" {
			return nil, response.BadRequest(ErrCommentRequired.Error())
		}

		if req.PostID == "" {
			return nil, response.BadRequest(ErrPostIDRequired.Error())
		}

		if req.UserID == "" {
			return nil, response.BadRequest(ErrUserIDRequired.Error())
		}

		comment, err := s.Create(ctx, req.UserID, req.PostID, req.Name, req.Comment)
		if err != nil {

			return nil, response.InternalServerError(err.Error())
		}

		return response.Created("success", comment, nil), nil
	}
}

func makeGetAllEndpoint(s Service, config Config) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(GetAllReq)

		filters := Filters{
			ID:     req.ID,
			UserID: req.UserID,
			PostID: req.PostID,
		}

		count, err := s.Count(ctx, filters)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		meta, err := meta.New(req.Page, req.Limit, count, config.LimPageDef)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		comments, err := s.GetAll(ctx, filters, meta.Offset(), meta.Limit())
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		return response.OK("success", comments, meta), nil
	}
}
