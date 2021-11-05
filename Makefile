bind-data:
	go-bindata -o=bindata/bindata.go -pkg=bindata ./static/... ./views/...
build-linux:
	gox -osarch="linux/amd64" -ldflags "-s -w" -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -output="image-hosting_{{.OS}}_{{.Arch}}"
build-windows:
	gox -osarch="windows/amd64" -ldflags "-s -w" -gcflags="all=-trimpath=${PWD}" -asmflags="all=-trimpath=${PWD}" -output="image-hosting_{{.OS}}_{{.Arch}}"

