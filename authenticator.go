// Copyright 2019, 2021 The Alpaca Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/base64"
	"net/http"
)

type authenticator struct {
	username, password string
}

func (a authenticator) do(req *http.Request, rt http.RoundTripper) (*http.Response, error) {
	creds := a.username+":"+a.password
	base64Text := make([]byte, base64.StdEncoding.EncodedLen(len(creds)))
    base64.StdEncoding.Encode(base64Text, []byte(creds))
    encodedCreds:= string(base64Text)

	req.Header.Set("Proxy-Authorization",
		"Basic "+encodedCreds)

	return rt.RoundTrip(req)
}
