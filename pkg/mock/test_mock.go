package pkg

import (
	"github.com/arya237/foodPilot/pkg"
	"testing"
)

func TestSomething(t *testing.T) {
	mock := NewMockRequiredFunctions()

	// Setup mock behavior
	mock.SetGetAccessToken(func(studentNumber, password string) (string, error) {
		if studentNumber == "12345" && password == "pass" {
			return "valid-token", nil
		}
		return "", pkg.ErorrInvalidToken // note: typo in your var name
	})

	mock.SetReserveFood(func(token string, meal pkg.ReserveModel) (string, error) {
		if token != "valid-token" {
			return "", pkg.ErorrInvalidToken
		}
		return "reservation-success", nil
	})

	// Use mock in your code under test
	token, err := mock.GetAccessToken("12345", "pass")
	if err != nil {
		t.Fatal(err)
	}

	meal := pkg.ReserveModel{
		ProgramId:  "prog1",
		FoodTypeId: "1",
		MealTypeId: "lunch",
	}

	result, err := mock.ReserveFood(token, meal)
	if err != nil {
		t.Fatal(err)
	}

	if result != "reservation-success" {
		t.Errorf("expected reservation-success, got %s", result)
	}

	// Verify call counts
	getAcc, _, reserve := mock.GetCallCounters()
	if getAcc != 1 || reserve != 1 {
		t.Errorf("unexpected call counts: GetAccessToken=%d, ReserveFood=%d", getAcc, reserve)
	}
}
