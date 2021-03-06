// Copyright 2019 Samaritan Authors
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

package utils

import (
	"io"
	"net"
	"testing"
)

func TestTestServer(t *testing.T) {
	srv := StartTestServer(t, "", func(conn *net.TCPConn) {
		io.Copy(conn, conn)
		conn.Close()
	})
	c, err := net.Dial("tcp", srv.Addr().String())
	if err != nil {
		t.Fatal("Failed to connect to test server: ", srv.Addr())
	}
	c.Close()
	srv.Close()
}
