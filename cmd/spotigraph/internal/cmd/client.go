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

package cmd

import (
	"github.com/spf13/cobra"

	chapterv1 "go.zenithar.org/spotigraph/api/gen/go/spotigraph/chapter/v1"
	guildv1 "go.zenithar.org/spotigraph/api/gen/go/spotigraph/guild/v1"
	personv1 "go.zenithar.org/spotigraph/api/gen/go/spotigraph/person/v1"
	squadv1 "go.zenithar.org/spotigraph/api/gen/go/spotigraph/squad/v1"
)

// -----------------------------------------------------------------------------

func clientCmd() *cobra.Command {
	c := &cobra.Command{
		Use:     "client",
		Aliases: []string{"c", "cli"},
		Short:   "Query the gRPC server",
	}

	// Add subcommands
	c.AddCommand(
		chapterv1.ChapterAPIClientCommand(),
		personv1.PersonAPIClientCommand(),
		squadv1.SquadAPIClientCommand(),
		guildv1.GuildAPIClientCommand(),
	)

	return c
}
