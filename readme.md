# Ech0 V4 Comment Migrator

本工具用于将Ech0 V3的评论（twikoo）迁移到Ech0 V4（原生评论）。

## 迁移准备

1. 导出twikoo的评论数据。
2. 备份v4数据库（非必须，但强烈建议）。迁移程序将**直接操作数据库**，请确认你明白自己在做什么。
3. 由于v3到v4重构了echo的id从数字到uuid，你需要提供一份新旧id的对照表，格式参考template.csv。只需要提供有评论的echo的id。

## 使用方法

```bash
# linux
e4cm -csv template.csv -json twikoo.json -db ech0.db

# windows
e4cm.exe -csv template.csv -json twikoo.json -db ech0.db
```

## 参数说明

|参数|必填|说明|
|---|---|---|
|`-csv`|是|旧id与新id的对照表文件|
|`-json`|是|twikoo的评论数据文件|
|`-db`|是|v4数据库文件|
