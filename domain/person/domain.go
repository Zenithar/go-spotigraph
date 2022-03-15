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

package person

// DomainOption defines property alteration function applicable to the domain
// object factory.
type DomainOption func(*defaultDomain)

// WithID is used to set the domain object identifier.
func WithID(value ID) DomainOption {
	return func(dd *defaultDomain) {
		dd.id = value
	}
}

// WithPrincipal is used to set the principal domain object property.
func WithPrincipal(value string) DomainOption {
	return func(dd *defaultDomain) {
		dd.principal = value
	}
}

// New instanciates a domain object.
func New(opts ...DomainOption) Person {
	dd := &defaultDomain{}

	// Apply builder options
	for _, o := range opts {
		o(dd)
	}

	return dd
}

// -----------------------------------------------------------------------------

var _ Person = (*defaultDomain)(nil)

type defaultDomain struct {
	id        ID
	principal string
	locked    bool
}

func (dd *defaultDomain) GetID() ID            { return dd.id }
func (dd *defaultDomain) GetPrincipal() string { return dd.principal }
func (dd *defaultDomain) IsLocked() bool       { return dd.locked }
