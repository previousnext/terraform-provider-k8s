FROM golang:1.19-alpine3.17 as build

RUN apk add make gcc musl-dev

WORKDIR /go/src/github.com/previousnext/terraform-provider-k8s
COPY . /go/src/github.com/previousnext/terraform-provider-k8s

RUN go build -ldflags "-linkmode external -extldflags -static" -o terraform-provider-k8s

FROM hashicorp/terraform:0.14.8 as run

RUN apk add bash

RUN mkdir -p /root/.terraform.d/plugins

COPY --from=build /go/src/github.com/previousnext/terraform-provider-k8s/terraform-provider-k8s /root/.terraform.d/plugins/registry.terraform.io/previousnext/k8s/99.0.0/linux_amd64/terraform-provider-k8s_v99.0.0
RUN chmod +x /root/.terraform.d/plugins/registry.terraform.io/*/*/*/linux_amd64/terraform-provider-*
