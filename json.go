// json project
// Copyright (C) 2017  geosoft1  geosoft1@gmail.com
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package json implements encoding and decoding of JSON as defined in RFC 4627
// as high level wrapper over standard json package.
package json

import (
	"encoding/json"
	"io"
	"log"
)

// Decode a JSON source into given structure.
func Decode(r io.Reader, v interface{}) {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		log.Println(err.Error())
		return
	}
}

// Encode a given structure into indented JSON string.
func Encode(w io.Writer, v interface{}) {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	if err := enc.Encode(v); err != nil {
		log.Println(err.Error())
		return
	}
}
