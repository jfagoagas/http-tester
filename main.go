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
	"flag"

	"github.com/rs/zerolog/log"
)

var (
	url  = flag.String("u", "", "URL to test responses HTTP status code")
	auth = flag.String("a", "", "Authorization Bearer Header")
)

func main() {
	flag.Usage = usage
	flag.Parse()
	if *url == "" {
		usage()
		log.Fatal().Msg("No input parameters.")
	}
	httpVerbTest(url)
}
