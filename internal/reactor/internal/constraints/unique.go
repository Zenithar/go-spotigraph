// Copyright 2019 Thibault NORMAND
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

package constraints

import (
	"context"
	"fmt"

	"go.zenithar.org/pkg/db"
	"go.zenithar.org/spotigraph/internal/helpers"
	"go.zenithar.org/spotigraph/internal/repositories"
)

// mustBeUnique specification checks if the given name already exists
func mustBeUnique(finder EntityRetrieverFunc, attribute string) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		// Retrieve object from repository
		object, err := finder(ctx)
		if err != nil && err != db.ErrNoResult {
			return err
		}
		if !isNil(object) {
			return fmt.Errorf("%s is already used", attribute)
		}

		return nil
	}
}

// PersonPrincipalMustBeUnique returns specification for user principal uniqueness
func PersonPrincipalMustBeUnique(users repositories.PersonRetriever, principal string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return users.FindByPrincipal(ctx, helpers.PrincipalHashFunc(principal))
		}, "Person principal")
}

// SquadLabelMustBeUnique returns specification for squad name uniqueness
func SquadLabelMustBeUnique(squads repositories.SquadRetriever, name string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return squads.FindByLabel(ctx, name)
		}, "Squad name")
}

// GuildLabelMustBeUnique returns specification for chapter name uniqueness
func GuildLabelMustBeUnique(guilds repositories.GuildRetriever, name string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return guilds.FindByLabel(ctx, name)
		}, "Guild name")
}

// ChapterLabelMustBeUnique returns specification for chapter name uniqueness
func ChapterLabelMustBeUnique(chapters repositories.ChapterRetriever, name string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return chapters.FindByLabel(ctx, name)
		}, "Chapter name")
}

// TribeLabelMustBeUnique returns specification for tribe name uniqueness
func TribeLabelMustBeUnique(tribes repositories.TribeRetriever, name string) func(ctx context.Context) error {
	return mustBeUnique(
		func(ctx context.Context) (interface{}, error) {
			return tribes.FindByLabel(ctx, name)
		}, "Tribe name")
}
