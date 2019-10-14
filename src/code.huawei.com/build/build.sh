#!/usr/bin/env bash
set -x

function main() {
    local cur_path=$(pwd)
    local base_path=$(readlink -f "${cur_path}/../../../")
    local org_path=$(readlink -f "${cur_path}/../")
    local main_path=${base_path}/src/code.huawei.com/cmd/demo
    export GOPATH=${base_path}
    echo "go path is ${GOPATH}"
    echo "main path is ${main_path}"
    local bin_dir=${org_path}/bin
    echo "bin dir is ${bin_path}"
    mkdir -p "${bin_dir}"
    local bin_file="${bin_dir}/domomesher"
    [ -f "${bin_file}" ] && rm -rf "${bin_file}"
    go build -o "${bin_file}" "${main_path}/main.go"
}

main