// Copyright 2015-2017 HenryLee. All Rights Reserved.
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

package teleport

import (
	"time"
)

// Config peer config
type Config struct {
	Id               string
	ReadTimeout      time.Duration // readdeadline for underlying net.Conn
	WriteTimeout     time.Duration // writedeadline for underlying net.Conn
	TlsCertFile      string
	TlsKeyFile       string
	SlowApiDuration  time.Duration // ns,µs,ms,s ...
	DefaultCodec     string
	DefaultGzipLevel int32

	DialTimeout time.Duration // for client role
	ListenAddrs []string      // for server role
}
