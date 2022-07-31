# MiraiGo-module-bili-followup

ID: `com.aimerneige.bili.followup`

Module for [MiraiGo-Template](https://github.com/Logiase/MiraiGo-Template)

## 鸣谢

感谢 [bilibili-API-collect](https://github.com/SocialSisterYi/bilibili-API-collect) 项目提供的 API 接口。

- [bilibili-API-collect#用户空间相关#投稿](https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/user/space.md#%E6%8A%95%E7%A8%BF)

## 功能

- 定时检查指定 b 站视频 up 主是否更新视频，检测到新视频发布后第一时间在指定群内通知更新信息。

## 使用方法

在适当位置引用本包

```go
package example

imports (
    // ...

    _ "github.com/yukichan-bot-module/MiraiGo-module-bili-followup"

    // ...
)

// ...
```

在全局配置文件中填入查询间隔（单位分钟）以及配置文件路径（默认路径为 `./followup.yaml`）。

当配置文件中需要查询的 up 较多时，请适当减少查询间隔以防止被 b 站封禁 ip。

```yaml
aimerneige:
  followup:
    sleep: 1
    path: "./followup.yaml"
```

在 `followup.yaml` 中填入 up 主及被通知群的配置：

```yaml
2143739: # up 主 id
  - 678429920 # 群号
  - 731500560
413848483:
  - 328521977
  - 306979321
```
