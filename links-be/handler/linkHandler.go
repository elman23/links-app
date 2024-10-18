package handler

import (
	"encoding/json"
	"fmt"
	"links-be/dto"
	"links-be/service"
	"net/http"

	"github.com/gorilla/mux"
)

type LinkHandler struct {
	service service.ILinkService
}

func (lh LinkHandler) CreateLink(w http.ResponseWriter, r *http.Request) {
	u := dto.Link{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		fmt.Fprintf(w, "Error in reading post body %s", err)
		return
	}

	link, err := lh.service.CreateLink(u)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error in Creating link %s", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(*link); err != nil {
		panic(err)
	}
}

func (lh LinkHandler) FindLink(w http.ResponseWriter, r *http.Request) {
	linkId := mux.Vars(r)["link_id"]

	link, err := lh.service.FindLink(linkId)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error in Finding link %v", err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(*link); err != nil {
		panic(err)
	}
}

func (lh LinkHandler) FindLinks(w http.ResponseWriter, r *http.Request) {

	links, err := lh.service.FindLinks()
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error in Finding link %v", err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(links); err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Cannot parse %v", err))
	}
}

func (lh LinkHandler) DeleteLink(w http.ResponseWriter, r *http.Request) {
	linkId := mux.Vars(r)["link_id"]

	resultString, err := lh.service.DeleteLink(linkId)
	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("Error in Deleting link %v", err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, *resultString)
}

func NewLinkHandler(service service.ILinkService) LinkHandler {
	return LinkHandler{
		service: service,
	}
}
