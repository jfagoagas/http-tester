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
	"fmt"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

var (
	httpVerbs = []string{
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"DELETE",
		"CONNECT",
		"OPTIONS",
		"TRACE",
		"PATCH",
	}

	securityHeaders = []string{
		"Strict-Transport-Security",
		"Content-Security-Policy",
		"X-Frame-Options",
		"X-Content-Type-Options",
		"Referrer-Policy",
		"Permissions-Policy",
	}
)

func banner() {
	fmt.Printf("Tool to test HTTP verbs and headers response for an URL.\n")
}

func usage() {
	banner()
	fmt.Printf("./http-verbs-test -u <URL> \n")
	fmt.Printf("If you need an Authorization header with a Bearer JWT,\nyou must input in a config.yml like 'token: <JWT>'\n")
	log.Fatal()
}

func (config *Config) readConfiguration() {
	err := yaml.Unmarshal(yamlConfig, config)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
