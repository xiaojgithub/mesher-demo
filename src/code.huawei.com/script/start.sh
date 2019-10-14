#!/usr/bin/env sh

init() {
    cur_dir=$(pwd)
    base_dir=$(readlink -f "${cur_dir}/../")
    bin_dir="${base_dir}/bin"
    bin_file_name="demomesher"
}

# just export 127.0.0.1,listen 127.0.0.1
exportIP() {
  local ethName="eth0"

  local ip_add="127.0.0.1"
  which ifconfig
  if [ $? -eq 0 ];then
    ip_add=$(ifconfig "${ethName}" | grep "inet" | awk '{print $2}')
  else
    which ip
    if [ $? -eq 0 ]; then
      local data=$(ip a | grep "inet" | grep "eth0" | awk '{print $2}')
      # shellcheck disable=SC2116
      ip_add=$(echo "${data%%/*}")
    fi
  fi
  ip_add="127.0.0.1"
  export LISTEN_IP=${ip_add}
}

run() {
    "${bin_dir}"/$bin_file_name
}

main() {
    init
    exportIP
    run
}

main