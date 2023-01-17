package main

import (
	"errors"
	"net/http"
	"useLaborartory-backend/models"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json: "status"`
		Message string `json: "message"`
		Version string `json: "version"`
	}{
		Status:  "active",
		Message: "useLaboratory api is running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) registerLab(w http.ResponseWriter, r *http.Request) {
	var requestPayload models.RegisterLabType

	err := app.readJSON(w, r, &requestPayload)

	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	if len(requestPayload.LabNumber) < 1 {
		app.errorJSON(w, errors.New("랩실을 선택해 주세요"), http.StatusBadRequest)
	} else if len(requestPayload.Applicant) < 1 {
		app.errorJSON(w, errors.New("신청자 이름을 작성해 주세요"), http.StatusBadRequest)
	} else if len(requestPayload.StudentId) != 4 {
		app.errorJSON(w, errors.New("학번을 다시 확인해 주세요"), http.StatusBadRequest)
	} else if len(requestPayload.Password) < 4 {
		app.errorJSON(w, errors.New("비밀번호는 4자 이상이어야 합니다"), http.StatusBadRequest)
	}

	success, err := app.DB.Registration(requestPayload)

	if err != nil || !success {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	app.writeJSON(w, http.StatusAccepted, "랩실 신청이 완료되었습니다")

}
