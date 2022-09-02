# maven

Drone持续集成Node插件，功能

- 内置支持
  - Npm
  - Yarn
- 依赖管理
- 脚本执行

## 使用

非常简单，只需要在`.drone.yml`里增加配置

```yaml
- name: 编译
  image: dronestock/node
  settings:
    folder: . # 默认值
    scripts:
      - build # 默认值
```

## 感谢Jetbrains

本项目通过`Jetbrains开源许可IDE`编写源代码，特此感谢
[![Jetbrains图标](https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png)](https://www.jetbrains.com/?from=dronestock/node)
