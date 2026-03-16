# GitFX

[![](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT "License: MIT")

[English](./README.md) | [繁體中文](./README.zh-tw.md)

## 目的

这个项目的目标是提供最简单、最快速、最无痛的方式来设置自托管的 Git 服务。

GitFX 是 [Gitea](https://github.com/go-gitea/gitea) 的一个分叉。由于它是用 Go 语言编写的，它可以在 Go 支持的所有平台和架构上运行，包括 Linux、macOS 和 Windows 的 x86、amd64、ARM 和 PowerPC 架构。

## 构建

从源代码树的根目录运行：

    TAGS="bindata" make build

如果需要 SQLite 支持：

    TAGS="bindata sqlite sqlite_unlock_notify" make build

`build` 目标分为两个子目标：

- `make backend` 需要 [Go Stable](https://go.dev/dl/)，所需版本在 [go.mod](/go.mod) 中定义。
- `make frontend` 需要 [Node.js LTS](https://nodejs.org/en/download/) 或更高版本以及 [pnpm](https://pnpm.io/installation)。

需要互联网连接来下载 go 和 npm 模块。从包含预构建前端文件的官方源代码压缩包构建时，不会触发 `frontend` 目标，因此可以在没有 Node.js 的情况下构建。

## 使用

构建后，默认情况下会在源代码树的根目录生成一个名为 `gitea` 的二进制文件。要运行它，请使用：

    ./gitea web

> [!注意]
> 如果您对使用我们的 API 感兴趣，我们提供了实验性支持，并附有 [文件](https://docs.gitea.com/api)。

## 贡献

预期的工作流程是：Fork -> Patch -> Push -> Pull Request

> [!注意]
>
> 1. **在开始进行 Pull Request 之前，您必须阅读 [贡献者指南](CONTRIBUTING.md)。**
> 2. 如果您在项目中发现了漏洞，请私下写信给 **security@gitea.io**。谢谢！

## 许可证

这个项目是根据 MIT 许可证授权的。
请参阅 [LICENSE](LICENSE) 文件以获取完整的许可证文本。

## 归属

GitFX 是从 [Gitea](https://github.com/go-gitea/gitea) 分叉而来，Gitea 最初从 [Gogs](https://gogs.io) 分叉。
