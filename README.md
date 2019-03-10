go get github.com/ixuzhi/testmode

git tag v1.0.0
git push --tags

git checkout -b v1
git push -u origin v1

go mod init modulename

go build