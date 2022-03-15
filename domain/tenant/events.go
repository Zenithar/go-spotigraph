// Licensed to Thibault Normand under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Thibault Normand licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
package tenant

import (
	"github.com/rs/xid"
	systemv1 "zntr.io/spotigraph/api/gen/go/spotigraph/system/v1"
)

func EventCreated(id ID, label, slug string) *systemv1.Event {
	return &systemv1.Event{
		EventType: systemv1.EventType_EVENT_TYPE_TENANT_CREATED,
		EventId:   xid.New().String(),
		Payload: &systemv1.Event_TenantCreated{
			TenantCreated: &systemv1.TenantCreated{
				Id:    string(id),
				Label: label,
				Slug:  slug,
			},
		},
	}
}
