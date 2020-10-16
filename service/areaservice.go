package areaservice

import (
	"encoding/json"
	"futaapibooking/dbconnect"
	"net/http"
)

func GetAreaByKey(w http.ResponseWriter, r *http.Request) error {
	search_key := r.FormValue("search")
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}
	if search_key != "" {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"name": search_key,
				},
			},
		}
	}

	data, err := dbconnect.GetAreaSearch(query)
	if err != nil {
		return err
	}
	e := json.NewEncoder(w)
	//e.SetIndent("", "  ")
	e.Encode(data)
	return nil
}
