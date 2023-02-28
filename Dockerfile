FROM storezhang/alpine:3.16.2


LABEL author="storezhang<华寅>"
LABEL email="storezhang@gmail.com"
LABEL qq="160290688"
LABEL wechat="storezhang"
LABEL description="Drone持续集成Node插件，支持测试、依赖管理、编译、打包等常规功能"


# 复制文件
COPY node /bin


# 模块存储目录
ENV MODULE_PATH /var/lib/node
# 修复安装其它模块时报SSL Provider错误
ENV NODE_OPTIONS --openssl-legacy-provider
# Pnpm模块存储路径
ENV XDG_DATA_HOME /var/lib/node


RUN set -ex \
    \
    \
    \
    # 安装依赖库
    && apk update \
    # 安装Node.js主体程序
    && apk --no-cache --update add nodejs \
    # 安装Npm依赖管理
    && apk --no-cache --update add npm \
    # 加速Npm
    && npm config set registry https://registry.npmmirror.com \
    # 安装Pnpm依赖管理
    && npm install --global pnpm \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/node \
    \
    \
    \
    && rm -rf /var/cache/apk/*


# 工作目录
WORKDIR /drone/src


# 执行命令
ENTRYPOINT /bin/node
