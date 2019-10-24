
// Copyright (c) 2012-2019 Red Hat, Inc.
// This program and the accompanying materials are made
// available under the terms of the Eclipse Public License 2.0
// which is available at https://www.eclipse.org/legal/epl-2.0/
//
// SPDX-License-Identifier: EPL-2.0
//
// Contributors:
//   Red Hat, Inc. - initial API and implementation
 
 package main

import (
    "flag"
    "fmt"
    "log"
    "net/http"
)

var (
    message string
    address string
)

func init() {
    flag.StringVar(&message, "message", "there", "a message to print")
    flag.StringVar(&address, "address", "0.0.0.0:8080", "address/port to listen on")
}

func main() {
    flag.Parse()
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s\n", message)
    })
    log.Printf("listening on %s...", address)
    log.Fatal(http.ListenAndServe(address, nil))
}