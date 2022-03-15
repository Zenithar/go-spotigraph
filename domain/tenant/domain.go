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

// DomainOption defines property alteration function applicable to the domain
// object factory.
type DomainOption func(*defaultDomain)

// WithID is used to set the domain object identifier.
func WithID(value ID) DomainOption {
	return func(dd *defaultDomain) {
		dd.id = value
	}
}

// WithLabel is used to set the label domain object property.
func WithLabel(value string) DomainOption {
	return func(dd *defaultDomain) {
		dd.label = value
	}
}

// WithSlug is used to set the slug domain object property.
func WithSlug(value string) DomainOption {
	return func(dd *defaultDomain) {
		dd.slug = value
	}
}

// WithLocked is used to set the locked domain object property.
func WithLocked(value bool) DomainOption {
	return func(dd *defaultDomain) {
		dd.locked = value
	}
}

// New instanciates a domain object.
func New(opts ...DomainOption) Tenant {
	dd := &defaultDomain{}

	// Apply builder options
	for _, o := range opts {
		o(dd)
	}

	return dd
}

// -----------------------------------------------------------------------------

var _ Tenant = (*defaultDomain)(nil)

type defaultDomain struct {
	id     ID
	label  string
	slug   string
	locked bool
}

func (dd *defaultDomain) GetID() ID        { return dd.id }
func (dd *defaultDomain) GetLabel() string { return dd.label }
func (dd *defaultDomain) GetSlug() string  { return dd.slug }
func (dd *defaultDomain) IsLocked() bool   { return dd.locked }
