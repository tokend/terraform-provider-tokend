FROM golang:1.18 as builder

WORKDIR /go/src/github.com/tokend/terraform-provider-tokend
COPY . .
ARG CI_GILAB_TOKEN
ENV CI_GILAB_TOKEN=${CI_GILAB_TOKEN}
RUN git config --global url."https://gitlab-ci-token:$CI_GILAB_TOKEN@gitlab.com".insteadOf https://gitlab.com
RUN go env -w GOPRIVATE=gitlab.com/*
RUN go mod download all
RUN CGO_ENABLED=0 GOOS=linux go build -o /terraform-provider-tokend -v github.com/tokend/terraform-provider-tokend


FROM hashicorp/terraform:1.2.9
WORKDIR /opt
RUN mkdir -p  /root/.terraform.d/plugins/registry.terraform.io/hashicorp/tokend/1.0.0/linux_amd64
COPY --from=builder /terraform-provider-tokend /root/.terraform.d/plugins/registry.terraform.io/hashicorp/tokend/1.0.0/linux_amd64/terraform-provider-tokend_v1.0.0
