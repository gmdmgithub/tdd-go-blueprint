package respond

import (
	"encoding/json"
	"net/http"
)

// Success - ...
func PerformRequest(w http.ResponseWriter, r *http.Request, status int, body interface{}) {

	// json.NewEncoder(w).Encode(body)

	bytesBody, err := json.Marshal(body)
	if err != nil {
		// TODO :handle error
	}
	if _, err := w.Write(bytesBody); err != nil {
		// TODO :handle error
	}

}
