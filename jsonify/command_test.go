package jsonify_test

import (
	"fmt"
	"github.com/mattJamesBoyle/jsonifyDir/jsonify"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDo(t *testing.T) {
	t.Run("Should return an error if the directory can't be found", func(t *testing.T) {
		_, err := jsonify.Do("oo")
		assert.NotNil(t, err)
	})

	t.Run("Should return a valid result if the directory is valid", func(t *testing.T) {
		result, _ := jsonify.Do("")
		fmt.Println(result)
		assert.NotNil(t, result)
	})
}
