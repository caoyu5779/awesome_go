package restoreIpAddresses

import (
	"testing"
	"reflect"
)

func TestRestoreIpAddresses(t *testing.T) {
	t.Run("test restore ip address", func(t *testing.T) {
		ips := "24525511135"

		got := RestoreIpAddresses(ips)

		want := []string{
			"245.255.11.135",
			"245.255.111.35",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v ; want : %v", got, want)
		}
	})
}
