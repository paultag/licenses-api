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

package license

import (
	"encoding/json"
	"os"
)

//
type Identifier struct {
	Identifier string
	Scheme     string
}

//
type Link struct {
	Note string
	URL  string
}

//
type OtherName struct {
	Name string
	Note string
}

//
type Text struct {
	ContentType string
	Name        string
	URL         string
}

//
type License struct {
	Id           string `json:"id"`
	Identifiers  []Identifier
	Links        []Link
	Name         string
	OtherNames   []OtherName
	SupersededBy string
	Tags         []string
	Texts        []Text
}

func LoadLicensesFiles(path string) ([]License, error) {
	ret := []License{}
	fh, err := os.Open(path)
	if err != nil {
		return []License{}, err
	}
	if err := json.NewDecoder(fh).Decode(&ret); err != nil {
		return []License{}, err
	}
	return ret, nil
}

// vim: foldmethod=marker
