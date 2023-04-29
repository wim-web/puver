# puver

puver publish terraform provider.

## requirements

- SHA256SUMS file
- SHA256SUMS.sig file 
- binary file

These files can be easily generated using [go releaser](https://goreleaser.com/ci/actions/#signing).

## installation

```
go install github.com/wim-web/puver@latest
```

## deploy

```
puver deploy -t faketoken \ 
    -o org -n test \
    --pubkey-path asset/pub.key \
    --shasum-path asset/terraform-provider-test_0.1.1_checksums.txt \
    --shasum-sig-path asset/terraform-provider-test_0.1.1_checksums.txt.sig \
    -v 0.1.1 \
    --os darwin --arch arm64 \
    --binary-path asset/terraform-provider-test_darwin_arm64.zip
```
