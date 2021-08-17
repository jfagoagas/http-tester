/*
Copyright (c) 2021 José Fagoaga jfagoagas@gmail.com

This program is free software: you can redistribute it and/or modify it under
the terms of the GNU General Public License as published by the Free Software
Foundation, version 3.

This program is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with
this program. If not, see <https://www.gnu.org/licenses/>.

SPDX-License-Identifier: GPL-3.0-only
*/
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func createHTTPClient() (client *http.Client) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client = &http.Client{Transport: transport, Timeout: time.Second * 10}
	return client
}

func createRequest(verb string, url *string, body io.Reader) (request *http.Request) {
	request, err := http.NewRequest(verb, *url, nil)
	if err != nil {
		log.Error().Err(err)
	}
	if *auth != "" {
		request.Header.Add("Authorization", "Bearer "+*auth)
	}
	return request
}

func doHTTPRequest(client *http.Client, request *http.Request) (response *http.Response) {
	response, err := client.Do(request)
	if err != nil {
		log.Error().Err(err)
	}
	if response == nil {
		fmt.Println(response)
		log.Fatal().Msg("Can not parse response")
	}
	return response
}

func httpVerbTest(url *string) {
	client := createHTTPClient()
	fmt.Printf("URL: %s\n\n", *url)
	for _, verb := range httpVerbs {
		request := createRequest(verb, url, nil)
		response := doHTTPRequest(client, request)
		fmt.Printf("%s\t%d\n", verb, response.StatusCode)
		if verb == httpVerbs[len(httpVerbs)-1] {
			fmt.Printf("\nResponse headers:\n")
			for header := range response.Header {
				content := ""
				for _, value := range response.Header[header] {
					content += value
					fmt.Printf("  - %s: %s\n", header, content)
				}
			}
		}
	}
}
