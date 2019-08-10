package fetch

import (
	"have-you-ordered/internal/app/orderserver"
	"testing"
)

func TestFetch(t *testing.T) {
	orderserver.Fetch("20190808")
}

func TestFetchAll(t *testing.T) {
	orderserver.StartFetchingMealRecord("20181214", "20190809")
}
