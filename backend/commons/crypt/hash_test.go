package crypt

import (
	"testing"
)

func TestShouldHashCorrectly(t *testing.T) {
	expected_hash := "FAOSJzKzuqjeiDZD1ZRsBWIq/OkUaI+Y0aArdgbXLx3y/KXbd9ZutSmsh5b7gz8htTc4P6dS+yj3nhMLlBjum9vqmjlh2Mhuue1Se7puTJ9AsXqwMds4t8yT+udx8bnswACdvogCtvHJiLikK8gEjaze3Qy9VY4dPFFVBsR8e8M="
	salt, err := StringToByte("YiQ9Vcgvl21WBX8R6jF2TgMcmPVA5V1GCem+Z+GLj+AhCJ0vflKkNrAywzKIY2VNnfiHJteXimDENs63MJ0kFg==")
	if err != nil {
		t.Errorf("Error encoutered during Base64 decoding : %s", err.Error())
	}

	if len(salt) != 64 {
		t.Logf("Salt length is wrong : %v bits", len(salt))
		t.FailNow()
	}

	value := HashPassword([]byte("TestPassword!"), salt)
	value_str := ByteToString(value)
	if value_str != expected_hash {
		t.Logf("Hash has not the intended value :\ngiven : %s\nexpected : %s", value_str, expected_hash)
		t.FailNow()
	}
}
