package model

import "testing"

func TestIdGenerator(t *testing.T) {

	t.Run("id_generator", func(t *testing.T) {

		t.Logf("testing id:%d", int64(GetNextId()))
	})
}
