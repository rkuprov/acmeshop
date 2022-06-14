package server

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	"AcmeShop/pkg/db"
	"AcmeShop/pkg/db/bq"
	"github.com/gorilla/mux"
)

func registerRoutes(h *mux.Router) {
	h.NewRoute().Methods(http.MethodGet).Path("/products/").HandlerFunc(getProductByID)
	h.NewRoute().Methods(http.MethodGet).Path("/products/bulk/").HandlerFunc(getProductBulk)
}

// getProductByID is a handler for the GET /products/{id} endpoint
func getProductByID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	q, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ac := db.NewQuery(db.WithTable(db.AcmeTableProducts), db.WithID(q.Get("id"))).String()
	qr, err := bq.BQUtil.Query(ac)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.Write(qr)
}

// getProductBulk is a handler for the GET /products/bulk endpoint that unmarshalls the body into a product struct
func getProductBulk(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	// get body from request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// unmarshal body into product struct
	// NOTE: I decided that the request body must come `\n` delimited.
	reqs := bytes.Split(body, []byte("\n"))
	q := db.NewQuery(db.WithAction("SELECT"), db.WithConjunction("FROM"), db.WithTable(db.AcmeTableProducts))
	for i := range reqs {
		if len(reqs[i]) == 0 {
			continue
		}
		q.BuildQuery(db.WithID(string(reqs[i])))
	}

	out, err := bq.BQUtil.Query(q.String())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(out)
}
