#!/bin/bash

red='\033[0;31m'
green='\033[0;32m'
yellow='\033[0;33m'
plain='\033[0m'

cur_dir=$(pwd)

# check root
[[ $EUID -ne 0 ]] && echo -e "${red}Fatal error:${plain}please run this script with root privilege\n" && exit 1

# check os
if [[ -f /etc/redhat-release ]]; then
    release="centos"
elif cat /etc/issue | grep -Eqi "debian"; then
    release="debian"
elif cat /etc/issue | grep -Eqi "ubuntu"; then
    release="ubuntu"
elif cat /etc/issue | grep -Eqi "centos|red hat|redhat"; then
    release="centos"
elif cat /proc/version | grep -Eqi "debian"; then
    release="debian"
elif cat /proc/version | grep -Eqi "ubuntu"; then
    release="ubuntu"
elif cat /proc/version | grep -Eqi "centos|red hat|redhat"; then
    release="centos"
else
    echo -e "${red}check system os failed,please contact with author!${plain}\n" && exit 1
fi

arch=$(arch)

if [[ $arch == "x86_64" || $arch == "x64" || $arch == "amd64" ]]; then
    arch="amd64"
elif [[ $arch == "aarch64" || $arch == "arm64" ]]; then
    arch="arm64"
else
    arch="amd64"
    echo -e "${red}fail to check system arch,will use default arch here: ${arch}${plain}"
fi

echo "Architecture: ${arch}"

if [ $(getconf WORD_BIT) != '32' ] && [ $(getconf LONG_BIT) != '64' ]; then
    echo "patch dosen't support 32bit(x86) system,please use 64 bit operating system(x86_64) instead,if there is something wrong,plz let me know"
    exit -1
fi

os_version=""

# os version
if [[ -f /etc/os-release ]]; then
    os_version=$(awk -F'[= ."]' '/VERSION_ID/{print $3}' /etc/os-release)
fi
if [[ -z "$os_version" && -f /etc/lsb-release ]]; then
    os_version=$(awk -F'[= ."]+' '/DISTRIB_RELEASE/{print $2}' /etc/lsb-release)
fi

if [[ x"${release}" == x"centos" ]]; then
    if [[ ${os_version} -le 6 ]]; then
        echo -e "${red}please use CentOS 7 or higher version${plain}\n" && exit 1
    fi
elif [[ x"${release}" == x"ubuntu" ]]; then
    if [[ ${os_version} -lt 16 ]]; then
        echo -e "${red}please use Ubuntu 16 or higher version${plain}\n" && exit 1
    fi
elif [[ x"${release}" == x"debian" ]]; then
    if [[ ${os_version} -lt 8 ]]; then
        echo -e "${red}please use Debian 8 or higher version${plain}\n" && exit 1
    fi
fi

install_base() {
    if [[ x"${release}" == x"centos" ]]; then
        yum install wget curl tar -y
    else
        apt install wget curl tar -y
    fi
}


install_patch() {
    systemctl stop patch
    cd /usr/local/ || exit

    if [ $# == 0 ]; then
        last_version=$(curl -Ls "https://api.github.com/repos/mazafard/patch/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
        if [[ ! -n "$last_version" ]]; then
            echo -e "${red}refresh patch version failed,it may due to Github API restriction,please try it later${plain}"
            exit 1
        fi
        echo -e "get patch latest version succeed:${last_version},begin to install..."
        wget -N --no-check-certificate -O /usr/local/patch-linux-${arch}.tar.gz https://github.com/mazafard/patch/releases/download/${last_version}/patch-linux-${arch}.tar.gz
        if [[ $? -ne 0 ]]; then
            echo -e "${red}dowanload patch failed,please be sure that your server can access Github{plain}"
            exit 1
        fi
    else
        last_version=$1
        url="https://github.com/mazafard/patch/releases/download/${last_version}/patch-linux-${arch}.tar.gz"
        echo -e "begin to install patch v$1 ..."
        wget -N --no-check-certificate -O /usr/local/patch-linux-${arch}.tar.gz ${url}
        if [[ $? -ne 0 ]]; then
            echo -e "${red}dowanload patch v$1 failed,please check the verison exists${plain}"
            exit 1
        fi
    fi

    if [[ -e /usr/local/patch/ ]]; then
        rm /usr/local/patch/ -rf
    fi

    tar zxvf patch-linux-${arch}.tar.gz
    rm patch-linux-${arch}.tar.gz -f
    cd patch || exit
    cp -f patch.service /etc/systemd/system/
    wget --no-check-certificate -O /usr/bin/patch https://raw.githubusercontent.com/mazafard/patch/master/patch.sh
    chmod +x /usr/local/patch/patch.sh
    chmod +x /usr/bin/patch
    systemctl daemon-reload
    systemctl enable patch
    systemctl start patch
    echo -e "${green}patch v${last_version}${plain} install finished,it is working now..."
    echo -e ""
    echo -e "patch control menu usages: "
    echo -e "----------------------------------------------"
    echo -e "patch              - Enter     control menu"
    echo -e "patch start        - Start     patch "
    echo -e "patch stop         - Stop      patch "
    echo -e "patch restart      - Restart   patch "
    echo -e "patch status       - Show      patch status"
    echo -e "patch enable       - Enable    patch on system startup"
    echo -e "patch disable      - Disable   patch on system startup"
    echo -e "patch log          - Check     patch logs"
    echo -e "patch update       - Update    patch "
    echo -e "patch install      - Install   patch "
    echo -e "patch uninstall    - Uninstall patch "
    echo -e "patch geo          - Update    geo  data"
    echo -e "----------------------------------------------"
}

echo -e "${green}excuting...${plain}"
install_base
install_patch $1