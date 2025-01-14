// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/openconfig/lemming/dataplane"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/local"
	"google.golang.org/grpc/reflection"

	log "github.com/golang/glog"
	dpb "github.com/openconfig/lemming/proto/dataplane"
)

var (
	port = flag.Int("port", dataplane.Port, "Port to listen on")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf("localhost:%d", *port)
	list, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %q: %v", addr, err)
	}
	srv := grpc.NewServer(grpc.Creds(local.NewCredentials()))

	data := &dataplane.Server{}
	dpb.RegisterHALServer(srv, data)
	reflection.Register(srv)

	if err := srv.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
