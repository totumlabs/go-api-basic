package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gilcrest/go-api-basic/domain/user"

	"github.com/rs/zerolog/hlog"

	"github.com/gilcrest/go-api-basic/domain/errs"
	"github.com/gilcrest/go-api-basic/service"
)

// CreateMovie is a HandlerFunc used to create a Movie
func (s *Server) handleMovieCreate(w http.ResponseWriter, r *http.Request) {
	logger := *hlog.FromRequest(r)

	u, err := user.FromRequest(r)
	if err != nil {
		errs.HTTPErrorResponse(w, logger, err)
		return
	}

	// Declare request body (rb) as an instance of service.MovieRequest
	rb := new(service.CreateMovieRequest)

	// Decode JSON HTTP request body into a Decoder type
	// and unmarshal that into the MovieRequest struct in the
	// AddMovieHandler
	err = json.NewDecoder(r.Body).Decode(&rb)
	defer r.Body.Close()
	// Call decoderErr to determine if body is nil, json is malformed
	// or any other error
	err = decoderErr(err)
	if err != nil {
		errs.HTTPErrorResponse(w, logger, err)
		return
	}

	response, err := s.CreateMovieService.Create(r.Context(), rb, u)
	if err != nil {
		errs.HTTPErrorResponse(w, logger, err)
		return
	}

	// Encode response struct to JSON for the response body
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, logger, errs.E(errs.Internal, err))
		return
	}
}

// handleMovieUpdate handles PUT requests for the /movies/{id} endpoint
// and updates the given movie
func (s *Server) handleMovieUpdate(w http.ResponseWriter, r *http.Request) {

	logger := *hlog.FromRequest(r)

	u, err := user.FromRequest(r)
	if err != nil {
		errs.HTTPErrorResponse(w, logger, err)
		return
	}

	// gorilla mux Vars function returns the route variables for the
	// current request, if any. id is the external id given for the
	// movie
	vars := mux.Vars(r)
	extlid := vars["extlID"]

	// Declare request body (rb) as an instance of service.MovieRequest
	rb := new(service.UpdateMovieRequest)

	// Decode JSON HTTP request body into a Decoder type
	// and unmarshal that into requestData
	err = json.NewDecoder(r.Body).Decode(&rb)
	defer r.Body.Close()
	// Call DecoderErr to determine if body is nil, json is malformed
	// or any other error
	err = decoderErr(err)
	if err != nil {
		errs.HTTPErrorResponse(w, logger, err)
		return
	}

	// External ID is from path variable, need to set separate
	// from decoding response body
	rb.ExternalID = extlid

	response, err := s.UpdateMovieService.Update(r.Context(), rb, u)
	if err != nil {
		errs.HTTPErrorResponse(w, logger, err)
		return
	}

	// Encode response struct to JSON for the response body
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, logger, errs.E(errs.Internal, err))
		return
	}
}

// handleMovieDelete handles DELETE requests for the /movies/{id} endpoint
// and updates the given movie
func (s *Server) handleMovieDelete(w http.ResponseWriter, r *http.Request) {

	logger := *hlog.FromRequest(r)

	// gorilla mux Vars function returns the route variables for the
	// current request, if any. id is the external id given for the
	// movie
	vars := mux.Vars(r)
	extlID := vars["extlID"]

	response, err := s.DeleteMovieService.Delete(r.Context(), extlID)
	if err != nil {
		errs.HTTPErrorResponse(w, logger, err)
		return
	}

	// Encode response struct to JSON for the response body
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, logger, errs.E(errs.Internal, err))
		return
	}
}

// handleFindMovieByID handles GET requests for the /movies/{id} endpoint
// and finds a movie by its ID
func (s *Server) handleFindMovieByID(w http.ResponseWriter, r *http.Request) {

	logger := *hlog.FromRequest(r)

	// gorilla mux Vars function returns the route variables for the
	// current request, if any. id is the external id given for the
	// movie
	vars := mux.Vars(r)
	extlID := vars["extlID"]

	response, err := s.FindMovieService.FindMovieByID(r.Context(), extlID)
	if err != nil {
		errs.HTTPErrorResponse(w, logger, err)
		return
	}

	// Encode response struct to JSON for the response body
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, logger, errs.E(errs.Internal, err))
		return
	}
}

// handleFindAllMovies handles GET requests for the /movies endpoint and finds
// all movies
func (s *Server) handleFindAllMovies(w http.ResponseWriter, r *http.Request) {

	logger := *hlog.FromRequest(r)

	response, err := s.FindMovieService.FindAllMovies(r.Context())
	if err != nil {
		errs.HTTPErrorResponse(w, logger, err)
		return
	}

	// Encode response struct to JSON for the response body
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, logger, errs.E(errs.Internal, err))
		return
	}
}

// handleLoggerRead handles GET requests for the /logger endpoint
func (s *Server) handleLoggerRead(w http.ResponseWriter, r *http.Request) {
	lgr := *hlog.FromRequest(r)

	response := s.LoggerService.Read()

	// Encode response struct to JSON for the response body
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.E(errs.Internal, err))
		return
	}
}

// handleLoggerUpdate handles PUT requests for the /logger endpoint
// and updates the logger globals
func (s *Server) handleLoggerUpdate(w http.ResponseWriter, r *http.Request) {
	lgr := *hlog.FromRequest(r)

	// Declare rb as an instance of service.LoggerRequest
	rb := new(service.LoggerRequest)

	// Decode JSON HTTP request body into a json.Decoder type
	// and unmarshal that into rb
	err := json.NewDecoder(r.Body).Decode(&rb)
	defer r.Body.Close()
	// Call DecoderErr to determine if body is nil, json is malformed
	// or any other error
	err = decoderErr(err)
	if err != nil {
		errs.HTTPErrorResponse(w, lgr, err)
		return
	}

	response, err := s.LoggerService.Update(rb)
	if err != nil {
		errs.HTTPErrorResponse(w, lgr, err)
		return
	}

	// Encode response struct to JSON for the response body
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, lgr, errs.E(errs.Internal, err))
		return
	}
}

// Ping handles GET requests for the /ping endpoint
func (s *Server) handlePing(w http.ResponseWriter, r *http.Request) {
	// pull logger from request context
	logger := *hlog.FromRequest(r)

	// pull the context from the http request
	ctx := r.Context()

	response := s.PingService.Ping(ctx, logger)

	// Encode response struct to JSON for the response body
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		errs.HTTPErrorResponse(w, logger, errs.E(errs.Internal, err))
		return
	}
}
