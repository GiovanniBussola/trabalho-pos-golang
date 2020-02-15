package handlers

import "encoding/json"

//esta função gera um erro no formato esperado
//pelo http.ResponseWriter e em um formato
//json
func formatJSONError(internalmessage string, message string) []byte {
	appError := struct {
		Message string `json:"message"`
		InternalMessage string `json:"internalmessage"`
	} {
		message,
		internalmessage,
	}
	response, err := json.Marshal(appError)
	if err != nil {
		return []byte(err.Error())
	}
	return response
}
