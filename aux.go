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
)

var httpVerbs = []string{
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

func banner() {
	fmt.Printf("Tool to test HTTP verbs response for an URL.\n")
}

func usage() {
	banner()
	fmt.Printf("./http-verbs-test -u <URL> -h <Bearer Token>\n")
	log.Fatal()
}
