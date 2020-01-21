package singleton

import "testing"

func TestSingleton(t *testing.T) {

	instance1 := GetGovernment()
	instance2 := GetGovernment()

	if instance1 != instance2 {
		t.Error("Objects are not equal!\n")
	}
}
