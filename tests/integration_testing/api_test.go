package integration_testing

import (
	"bytes"
	"encoding/json"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPost(t *testing.T) {
	// Test POST /v1/parking-lot
	Convey("Subject: Test CreateParkingLot POST API\n", t, func() {
		// Create a request body
		requestBody := map[string]interface{}{
			"name":      "p15",
			"capacity":  50,
			"createdBy": 1,
		}
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("Failed to marshal JSON request body: %v", err)
		}

		req, err := http.NewRequest("POST", "/v1/parking-lot", bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, req)

		// logs.Info("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

		Convey("Subject: Test CreateParkingLot POST API\n", t, func() {
			Convey("Status Code Should Be 1000", func() {
				So(w.Code, ShouldEqual, 1000)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
		})
	})

	//Test  /v1/parking-slot
	Convey("Subject: Test CreateParkingSlot POST API\n", t, func() {
		// Create a request body
		requestBody := map[string]interface{}{
			"parkingLotId": 1,
		}
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("Failed to marshal JSON request body: %v", err)
		}

		req, err := http.NewRequest("POST", "/v1/parking-slot", bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, req)

		// logs.Info("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

		Convey("Subject: Test CreateParkingSlot POST API\n", t, func() {
			Convey("Status Code Should Be 1000", func() {
				So(w.Code, ShouldEqual, 1000)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
		})
	})

	//Test  park/vehicle
	Convey("Subject: Test park vehicle POST API\n", t, func() {
		// Create a request body
		requestBody := map[string]interface{}{
			"parkingLotId": 7,
			"vehicleId":    1,
		}
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("Failed to marshal JSON request body: %v", err)
		}

		req, err := http.NewRequest("POST", "/v1/vehicle/park", bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, req)

		// logs.Info("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

		Convey("Subject: Test park vehicle POST API\n", t, func() {
			Convey("Status Code Should Be 1000", func() {
				So(w.Code, ShouldEqual, 1000)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
		})
	})

	//Test  /v1/vehicle
	Convey("Subject: Test add vehicle POST API\n", t, func() {
		// Create a request body
		requestBody := map[string]interface{}{
			"userId":         3,
			"registrationNo": "gghde77dfdf",
		}
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("Failed to marshal JSON request body: %v", err)
		}

		req, err := http.NewRequest("POST", "/v1/vehicle", bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, req)

		// logs.Info("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

		Convey("Subject: Add park vehicle POST API\n", t, func() {
			Convey("Status Code Should Be 1000", func() {
				So(w.Code, ShouldEqual, 1000)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
		})
	})

}

func TestPut(t *testing.T) {

	Convey("Subject: Test Update Parking lot API\n", t, func() {
		// Create a request body
		requestBody := map[string]interface{}{
			"name":      "lot122",
			"capacity":  5,
			"createdBy": 1,
			"active":    true,
		}
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("Failed to marshal JSON request body: %v", err)
		}

		req, err := http.NewRequest("PUT", "/v1/parking-lot/13", bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, req)

		// logs.Info("testing", "TestPost", "Code[%d]\n%s", w.Code, w.Body.String())

		Convey("Subject: Test UpdateParkingLot POST API\n", t, func() {
			Convey("Status Code Should Be 1000", func() {
				So(w.Code, ShouldEqual, 1000)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
		})
	})

	Convey("Subject: Test Update Parking slot API\n", t, func() {
		// Create a request body
		requestBody := map[string]interface{}{
			"parkingLotId":  1,
			"isMaintenance": false,
			"isOccupied":    false,
		}
		jsonBody, err := json.Marshal(requestBody)
		if err != nil {
			t.Fatalf("Failed to marshal JSON request body: %v", err)
		}

		req, err := http.NewRequest("PUT", "/v1/parking-slot/13", bytes.NewBuffer(jsonBody))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		beego.BeeApp.Handlers.ServeHTTP(w, req)

		// logs.Info("testing", "TestPost", "Code[%d]\n%s", w.Code, w.Body.String())

		Convey("Subject: Test UpdateParkingSlot POST API\n", t, func() {
			Convey("Status Code Should Be 1000", func() {
				So(w.Code, ShouldEqual, 1000)
			})
			Convey("The Result Should Not Be Empty", func() {
				So(w.Body.Len(), ShouldBeGreaterThan, 0)
			})
		})
	})

}
