package km

//go test for keyManager
import (
	"testing"
)

func TestReadFiles(t *testing.T) {
	t.Run("Launch", func(t *testing.T) {
		t.Helper()
		Run()
	})
}
