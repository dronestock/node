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
# TODO 暂时解决Node版本过高导致的无法编译的问题
ENV NODE_OPTIONS --openssl-legacy-provider


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
    # 安装Yarn依赖管理
    && npm install --global yarn \
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
