package unit_testing

import (
	"ParkMe/services"
	"testing"
)

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestCalculatePayment(t *testing.T) {

	t.Run("Test with non frictional duration", func(t *testing.T) {
		got := int(services.CalculatePayment(2))
		want := 20
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test with non frictional duration", func(t *testing.T) {
		got := int(services.CalculatePayment(2.5))
		want := 25
		assertCorrectMessage(t, got, want)
	})

	//Convey("Subject: Test payment calculation function POST API\n", t, func() {
	//	Convey("Test with non frictional duration", func() {
	//		So(services.CalculatePayment(2.5), ShouldEqual, 25)
	//	})
	//	Convey("Test with non frictional duration", func() {
	//		So(services.CalculatePayment(2), ShouldEqual, 25)
	//	})
	//})

}
