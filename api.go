/* {{{ Copyright (c) Paul R. Tagliamonte <paultag@opensource.org>, 2015
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>. }}} */

package main

import (
	"encoding/json"
	"net/http"

	"opensource.org/api/license"
)

func writeJSON(w http.ResponseWriter, data interface{}, code int) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return err
	}
	return nil
}

func writeError(w http.ResponseWriter, message string, code int) error {
	return writeJSON(w, map[string][]map[string]string{
		"errors": []map[string]string{
			map[string]string{"message": message},
		},
	}, code)
}

func main() {
	mux := http.NewServeMux()
	licenses, err := license.LoadLicensesFiles("/home/paultag/licenses.json")
	if err != nil {
		panic(err)
	}

	licenseIds := licenses.GetIds()
	licenseIdMap := licenses.GetIdMap()

	mux.HandleFunc("/licenses/", func(w http.ResponseWriter, req *http.Request) {
		writeJSON(w, licenseIds, 200)
	})

	licenseEndpoint := "/license/"
	mux.HandleFunc(licenseEndpoint, func(w http.ResponseWriter, req *http.Request) {
		path := req.URL.Path[len(licenseEndpoint):]
		if license, ok := licenseIdMap[path]; ok {
			writeJSON(w, license, 200)
			return
		}
		writeError(w, "Unknown license", 404)
	})

	http.ListenAndServe(":8000", mux)
}

// vim: foldmethod=marker
