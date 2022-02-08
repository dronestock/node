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
  setttings:
    scripts:
      - build # 默认值，可以不用配置
```
