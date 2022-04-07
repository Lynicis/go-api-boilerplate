package healthhandlermodel

// HealthEndpoint health check endpoint body
type HealthEndpoint struct {
	Status string `json:"status"`
}
