package incident

import (
	"context"
	"fmt"
)

// IncidentsService handles communication with the incident related
// methods of the Incident.io API.
//
// API docs: https://api-docs.incident.io/#tag/Incidents
type IncidentsService service

// List list all incidents for an organisation.
//
// API docs: https://api-docs.incident.io/#operation/Incidents_List
func (s *IncidentsService) List(ctx context.Context, opts *IncidentsListOptions) (*IncidentsList, *Response, error) {
	u := "incidents"
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	v := &IncidentsList{}
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// Get returns a single incident.
//
// id represents the unique identifier for the incident
//
// API docs: https://api-docs.incident.io/#operation/Incidents_Show
func (s *IncidentsService) Get(ctx context.Context, id string) (*IncidentResponse, *Response, error) {
	u := fmt.Sprintf("incidents/%s", id)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	// TODO Should we return the incident directly? Would be more userfriendly - Maybe talking to the Incident.io folks?
	v := &IncidentResponse{}
	resp, err := s.client.Do(ctx, req, v)
	if err != nil {
		return nil, resp, err
	}

	return v, resp, nil
}

// API docs: https://api-docs.incident.io/tag/Incidents-V2#operation/Incidents%20V2_Create
func (s *IncidentsService) Create(ctx context.Context, opts *IncidentCreateOptions) (*IncidentResponse, *Response, error) {
	body, err := createBody(opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", "incidents", body)
	if err != nil {
		return nil, nil, err
	}

	incident := IncidentResponse{}
	resp, err := s.client.Do(ctx, req, &incident)
	if err != nil {
		return nil, resp, err
	}

	return &incident, resp, nil
}
