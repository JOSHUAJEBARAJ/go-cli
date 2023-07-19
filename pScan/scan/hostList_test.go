package scan_test

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"joshua.com/pScan/scan"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		host      string
		exepctLen int
		expectErr error
	}{
		{"Addnew", "host2", 2, nil},
		{"AddExisting", "host1", 1, scan.ErrExists},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			h1 := &scan.HostList{}
			if err := h1.Add("host1"); err != nil {
				t.Fatal(err)
			}
			err := h1.Add(tc.host)
			if tc.expectErr != nil {
				if err == nil {
					t.Fatalf("Expected eroor, got nil instead")
				}
				if !errors.Is(err, tc.expectErr) {
					t.Errorf("expecter error %q, got %q instead", tc.expectErr, err)
				}
				return
			}

		})
	}

}

func TestRemove(t *testing.T) {
	testCases := []struct {
		name      string
		host      string
		expectLen int
		expectErr error
	}{
		{"Remove existing", "host1", 1, nil},
		{"RemoveNotfound", "host3", 1, scan.ErrNotExists},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			h1 := &scan.HostList{}
			for _, h := range []string{"host1", "host2"} {
				if err := h1.Add(h); err != nil {
					t.Fatal(err)
				}
			}
			err := h1.Remove(tc.host)

			if tc.expectErr != nil {
				if err == nil {
					t.Fatalf("Expected eroor, got nil instead")
				}
				if !errors.Is(err, tc.expectErr) {
					t.Errorf("expecter error %q, got %q instead", tc.expectErr, err)
				}
				return
			}
		})
	}
}

// save load test

func TestSaveLoad(t *testing.T) {
	hl1 := scan.HostList{}
	h2 := scan.HostList{}
	hl1.Add("host1")
	tf, err := ioutil.TempFile("", "")

	if err != nil {
		t.Fatalf("Error creating temp file")
	}
	defer os.Remove(tf.Name())

	if err := hl1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}
	if err := h2.Load(tf.Name()); err != nil {
		t.Fatalf("Erro reading from file: %s", err)

		if hl1.Hosts[0] != h2.Hosts[0] {
			t.Errorf("Host  %q should match %q", hl1.Hosts[0], h2.Hosts[0])
		}
	}
}

// test case for not load the file which doesn't exist

func TestLoadNoFile(t *testing.T) {
	tf, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatalf("Error creating temp file %s", err)
	}
	if err := os.Remove(tf.Name()); err != nil {
		t.Fatalf("Error deleting temp file %s", err)
	}

}
