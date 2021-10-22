package detector

import (
	"testing"

	"github.com/kubeshop/testkube/pkg/api/v1/client"
	"github.com/stretchr/testify/assert"
)

const (
	exampleValidContent       = `{ "info": { "_postman_id": "3d9a6be2-bd3e-4cf7-89ca-354103aab4a7", "name": "TestKube", "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json" }, "item": [ { "name": "Health", "event": [ { "listen": "test", "script": { "exec": [ "pm.test(\"Status code is 200\", function () {", "    pm.response.to.have.status(200);", "});" ], "type": "text/javascript" } } ], "request": { "method": "GET", "header": [], "url": { "raw": "{{URI}}/health", "host": [ "{{URI}}" ], "path": [ "health" ] } }, "response": [] } ] } `
	exampleInvalidContent     = `{"some":"json content"}`
	exampleInvalidJSONContent = `some non json content`
)

func TestPostmanCollectionAdapter(t *testing.T) {

	t.Run("Is return true when valid content", func(t *testing.T) {
		detector := PostmanCollectionAdapter{}
		is, name := detector.Is(client.CreateScriptOptions{
			Content: exampleValidContent,
		})

		assert.Equal(t, "postman/collection", name)
		assert.True(t, is, "content should be of postman type")
	})

	t.Run("Is return false in case of invalid JSON content", func(t *testing.T) {
		detector := PostmanCollectionAdapter{}
		is, name := detector.Is(client.CreateScriptOptions{
			Content: exampleInvalidContent,
		})

		assert.Empty(t, name)
		assert.False(t, is, "content should not be of postman type")

	})

	t.Run("Is return false in case of content which is not JSON ", func(t *testing.T) {
		detector := PostmanCollectionAdapter{}
		is, name := detector.Is(client.CreateScriptOptions{
			Content: exampleInvalidJSONContent,
		})

		assert.Empty(t, name)
		assert.False(t, is, "content should not be of postman type")
	})
}
