// Copyright 2015-2018 trivago N.V.
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

package format

import (
	"testing"

	"github.com/trivago/gollum/core"
	"github.com/trivago/tgo/ttesting"
)

func TestJSON(t *testing.T) {
	expect := ttesting.NewExpect(t)

	config := core.NewPluginConfig("", "format.JSON")
	plugin, err := core.NewPluginWithConfig(config)
	expect.NoError(err)

	formatter, casted := plugin.(*JSON)
	expect.True(casted)

	msg := core.NewMessage(nil, []byte(`{"a":"bar","b":2,"c":{"a":"bar"},"d":["a","b"]}`), nil, core.InvalidStreamID)

	err = formatter.ApplyFormatter(msg)
	expect.NoError(err)

	metadata := msg.GetMetadata()

	expect.MapEqual(metadata, "a", "bar")
	expect.MapEqual(metadata, "b", 2.0)

	c, err := metadata.MarshalMap("c")
	expect.NoError(err)
	expect.MapEqual(c, "a", "bar")

	expect.MapEqual(metadata, "d", []interface{}{"a", "b"})
}
