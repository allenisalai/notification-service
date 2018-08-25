package controller

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/allenisalai/notification-service/internal/model"
	"github.com/allenisalai/notification-service/internal"
	"encoding/json"
	"strconv"
	"github.com/allenisalai/notification-service/internal/response"
	"fmt"
)

func NotificationTypeCGet() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		nts := new([]model.NotificationType)
		database.GetDb().Find(&nts)

		obj, err := json.Marshal(nts)
		if err != nil {
			response.WriteErrorResponse(w, 400, "json-encode", "Failed to marshal results")
			return
		}

		w.Write(obj)
	}
}

func NotificationTypeGet() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])

		if err != nil {
			response.WriteErrorResponse(w, 400, "bad-id", fmt.Sprintf("%v could not be parsed into an integer", vars["id"]))
			return
		}

		nt := model.NotificationType{}
		err = nt.GetById(database.GetDb(), id)
		if err != nil {
			response.WriteErrorResponse(w, 404, "missing-object", "Object Not found")
			return
		}

		obj, _ := json.Marshal(nt)
		w.Write(obj)
	}
}

func NotificationTypePost() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		nt := model.NotificationType{}
		err := decoder.Decode(&nt)

		if err != nil {
			w.WriteHeader(500)
			log.Fatalf(err.Error())
		}
		log.Print(nt)

		errs := database.GetDb().Create(&nt).GetErrors()

		if len(errs) > 0 {
			for _, e := range errs {
				log.Fatalf(e.Error())
			}
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(201)

		ret, err := json.Marshal(nt)

		if err != nil {
			log.Fatalf(err.Error())
			w.WriteHeader(500)
		}
		w.Write(ret)
	}
}

func NotificationTypePatch() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		nt := model.NotificationType{}
		id, err := strconv.Atoi(vars["id"])

		if err != nil {
			w.WriteHeader(400)
			log.Fatalf("%v could not be parsed into an integer", vars["id"])
		}

		err = nt.GetById(database.GetDb(), id)
		if err != nil {
			w.WriteHeader(404)
			return
		}

		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&nt)
		if err != nil {
			w.WriteHeader(500)
			log.Fatalf(err.Error())
			return
		}

		nt.Save(database.GetDb())

		ret, err := json.Marshal(&nt)
		if err != nil {
			log.Fatalf(err.Error())
			w.WriteHeader(500)
		}
		w.Write(ret)
	}
}

func NotificationTypeDelete() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		nt := model.NotificationType{}
		id, err := strconv.Atoi(vars["id"])

		if err != nil {
			w.WriteHeader(400)
			log.Fatalf("%v could not be parsed into an integer", vars["id"])
		}

		err = nt.GetById(database.GetDb(), id)
		if err != nil {
			w.WriteHeader(404)
			return
		}

		database.GetDb().Delete(&nt)

		w.WriteHeader(202)
	}
}
