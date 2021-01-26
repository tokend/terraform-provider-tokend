go get -v -u github.com/alecthomas/gometalinter
gometalinter --install &> /dev/null
go install -v
gometalinter --vendor ./v3/...
