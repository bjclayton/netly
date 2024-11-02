package main

import "net/http"

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	//data := map[string]string{
	//	"status": "ok",
	//	"env":    app.config.env,
	//	//"version": version,
	//}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Ok"))

	//if err := app.jsonResponse(w, http.StatusOK, data); err != nil {
	//	app.internalServerError(w, r, err)
	//}
}
