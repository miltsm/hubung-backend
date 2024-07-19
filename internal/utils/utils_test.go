package utils

import "testing"

func TestUUIDGeneration(t *testing.T) {
	uuid := GenerateUUID()
	t.Log(uuid)
	if len(uuid) > 0 {
		t.Fatal("empty uuid")
	}
}
