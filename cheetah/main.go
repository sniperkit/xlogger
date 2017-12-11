package main

import (
	"farm.e-pedion.com/repo/cheetah/config"
	"farm.e-pedion.com/repo/cheetah/data"
	"farm.e-pedion.com/repo/cheetah/logger"
	"flag"
	"fmt"
	"net/http"
)

type errorHandler func(http.ResponseWriter, *http.Request) error

func (h errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

type chitaHandler struct {
}

func (h chitaHandler) json(w http.ResponseWriter, status int, result interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return nil
}

func (h chitaHandler) result(w http.ResponseWriter, status int) error {
	w.WriteHeader(status)
	return nil
}

func (h chitaHandler) error(w http.ResponseWriter, err error) error {
	w.WriteHeader(http.StatusInternalServerError)
	return err
}

func (h chitaHandler) Insert(w http.ResponseWriter, r *http.Request) error {
	var user data.User
	if decodeErr := user.Unmarshall(r.Body); decodeErr != nil {
		return h.error(w, decodeErr)
	}
	manager := data.NewUserManager()
	if saveErr := manager.Save(user); saveErr != nil {
		return h.error(w, saveErr)
	}
	return h.result(w, http.StatusCreated)
}

func (h chitaHandler) Find(w http.ResponseWriter, r *http.Request) error {
	return h.json(w, http.StatusOK, nil)
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "conf", "./etc/cheetah/cheetah.yaml", "Program configuration file")
	flag.Parse()
	fmt.Printf("ChitaInitializing[Path=%s]\n", configPath)
	if err := config.Setup(configPath); err != nil {
		panic(err)
	}
	configuration := config.Get()

	if err := logger.Setup(configuration.Logger); err != nil {
		panic(err)
	}
	rootLogger := logger.GetLogger()

	handler := new(chitaHandler)
	http.Handle("/insert", errorHandler(handler.Insert))
	http.Handle("/find/", errorHandler(handler.Find))

	fmt.Printf("ChitaStarted[Bind=%v]\n", configuration.Bind)
	rootLogger.Info("ChitaStarted",
		logger.String("bind", configuration.Bind),
	)
	err := http.ListenAndServe(configuration.Bind, nil)
	if err != nil {
		panic(err)
	}
}
