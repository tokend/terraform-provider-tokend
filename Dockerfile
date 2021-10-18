FROM golang:1.15 as builder

WORKDIR /go/src/github.com/tokend/terraform-provider-tokend
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /terraform-provider-tokend -v github.com/tokend/terraform-provider-tokend


FROM zenika/terraform-aws-cli:release-5.1_terraform-0.11.15_awscli-1.20.41
WORKDIR /opt
COPY ./run.sh /bin/run.sh
COPY --from=builder /terraform-provider-tokend /bin/terraform-provider-tokend
