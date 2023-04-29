package puver

import "testing"

func TestHashData(t *testing.T) {
	str := `24dcaaf939545bf4a7e231b0f3634e1288532911d838cfc71c14f3bcec7a56e5  terraform-provider-test_darwin_arm64.zip
	72728b914f2326887edb3cd31a02180b6c77453eb3dbd3640d9c7df164465180  terraform-provider-test_linux_arm64.zip
	745efc964c0da50f0e29c9a848d0ac4d123e71970c7d2b6a67b57e452d7e6d45  terraform-provider-test_linux_amd64.zip
	bdaf390a31da5c1966a53f912d17c85ecd6a1d61f2f3f1733debcce3f1295cc1  terraform-provider-test_darwin_amd64.zip`

	hash := HashData(str, "darwin", "amd64")

	if "bdaf390a31da5c1966a53f912d17c85ecd6a1d61f2f3f1733debcce3f1295cc1" != hash {
		t.Errorf("failed HashData: %s\n", hash)
	}
}
