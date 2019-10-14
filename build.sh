#!/usr/bin/env bash

set -e

function init() {
    cur_dir=$(pwd)
}

function build() {
    local main_dir="${cur_dir}"/src/code.huawei.com/cmd/demo
    unset GOPATH
    export GOPATH="${cur_dir}"
    echo "go path is ${GOPATH}"
    local bin_file="${cur_dir}/demomesher"

    [ -f "${bin_file}" ] && rm -rf "${bin_file}"
    [ -f "${cur_dir}/Dockerfile" ] && rm -rf "${cur_dir}"/Dockerfile
    [ -f "${cur_dir}/start.sh" ] && rm -rf "${cur_dir}/start.sh"

    CGO_ENABLED=0 go build -o "${bin_file}" "${main_dir}/main.go"

    cp "${cur_dir}/src/code.huawei.com/deployment/Dockerfile" "${cur_dir}"
    cp "${cur_dir}/src/code.huawei.com/script/start.sh" "${cur_dir}"
}

function main() {
  init
  build
}

main