package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDefaultConfig(t *testing.T) {
	got, err := loadDefaultConfig()
	if err != nil {
		t.Fatal(err)
	}
	if got.Api == nil {
		t.Fatal("expected a valid api config got nil")
	}
	if got.Redis == nil {
		t.Fatal("expected a valid redis config got nil")
	}
	if got.Etcd == nil {
		t.Fatal("expected a valid etcd config got nil")
	}
	if got.Loki == nil {
		t.Fatal("expected a valid loki config got nil")
	}
}

func TestLoadConfigFromFile(t *testing.T) {
	fname := "test.yaml"
	if err := os.WriteFile(filepath.Join(".", fname),
		[]byte(defaultConfig), os.FileMode(0777)); err != nil {
		t.Fatal(err)
	}
	got, err := loadConfigFromFile(filepath.Join(".", fname))
	if err != nil {
		t.Fatal(err)
	}

	if got.Api == nil {
		t.Fatal("expected a valid api config got nil")
	}
	if got.Redis == nil {
		t.Fatal("expected a valid redis config got nil")
	}
	if got.Etcd == nil {
		t.Fatal("expected a valid etcd config got nil")
	}
	if got.Loki == nil {
		t.Fatal("expected a valid loki config got nil")
	}
	t.Cleanup(func() {
		if err := os.Remove(filepath.Join(".", fname)); err != nil {
			t.Fatal(err)
		}
	})

}
