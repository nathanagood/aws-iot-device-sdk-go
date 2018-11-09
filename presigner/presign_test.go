// Copyright 2018 SEQSENSE, Inc.
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

package presigner

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
)

func ExamplePresigner_PresignWss() {
	os.Clearenv()
	os.Setenv("AWS_ACCESS_KEY_ID", "AKAAAAAAAAAAAAAAAAAA")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "1111111111111111111111111111111111111111")
	os.Setenv("AWS_SESSION_TOKEN", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	os.Setenv("AWS_REGION", "world-1")

	sess := session.Must(session.NewSession())
	ps := New(sess)
	wssURL, err := ps.PresignWss("test.iot.world-1.amazonaws.com", time.Hour*24, time.Unix(0, 0))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", wssURL)

	// Output:
	// wss://test.iot.world-1.amazonaws.com/mqtt?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKAAAAAAAAAAAAAAAAAA%2F19700101%2Fworld-1%2Fiotdevicegateway%2Faws4_request&X-Amz-Date=19700101T000000Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=4cfbc8acc899f7aac3153cd17c94204d6989f86d8cb1173e46143512270c89c2&X-Amz-Security-Token=aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
}

func ExamplePresigner_PresignWss_withoutSessionToken() {
	os.Clearenv()
	os.Setenv("AWS_ACCESS_KEY_ID", "AKAAAAAAAAAAAAAAAAAA")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "1111111111111111111111111111111111111111")
	os.Setenv("AWS_REGION", "world-1")

	sess := session.Must(session.NewSession())
	ps := New(sess)
	wssURL, err := ps.PresignWss("test.iot.world-1.amazonaws.com", time.Hour*24, time.Unix(0, 0))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", wssURL)

	// Output:
	// wss://test.iot.world-1.amazonaws.com/mqtt?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKAAAAAAAAAAAAAAAAAA%2F19700101%2Fworld-1%2Fiotdevicegateway%2Faws4_request&X-Amz-Date=19700101T000000Z&X-Amz-Expires=86400&X-Amz-SignedHeaders=host&X-Amz-Signature=4cfbc8acc899f7aac3153cd17c94204d6989f86d8cb1173e46143512270c89c2
}

func ExamplePresigner_PresignWssNow() {
	sess := session.Must(session.NewSession())
	ps := New(sess)
	wssURL, err := ps.PresignWssNow("test.iot.world-1.amazonaws.com")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", wssURL)
}
