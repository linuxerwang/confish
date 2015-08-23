package confish

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

type wirelessSettings struct {
	Name     string         `cfg-attr:"name"`
	SSID     string         `cfg-attr:"ssid"`
	Channels []string       `cfg-attr:"channels"`
	Speed    int            `cfg-attr:"speed"`
	Enabled  bool           `cfg-attr:"enabled"`
	Strength float32        `cfg-attr:"strength"`
	Qos      map[string]int `cfg-attr:"qos"`
}

type networkSettings struct {
	Owner    string              `cfg-attr:"owner"`
	Wireless []*wirelessSettings `cfg-attr:"wireless"`
}

func TestWrite(t *testing.T) {
	ns := &networkSettings{
		"Library",
		[]*wirelessSettings{
			{
				"DLink", "dlink-1234-5678", []string{"2", "5", "10"}, 54, false, 0.5,
				map[string]int{
					"public":  100,
					"private": 200,
					"corp":    500,
				},
			},
			{"Linksys", "linksys-5678-1234", []string{"6", "7", "8", "9", "10"}, 180, true, 0.8,
				map[string]int{
					"public":  50,
					"private": 50,
					"corp":    120,
				},
			},
		},
	}

	var b bytes.Buffer
	Write(&b, ns, "network")

	restored := &networkSettings{}
	err := Parse(strings.NewReader(b.String()), restored)
	if err != nil {
		t.Fatalf("failed to parse confish file")
	}

	if !reflect.DeepEqual(restored, ns) {
		t.Fatalf("got network settings %+v, want %+v", restored, ns)
	}
}
