# GitFX

[![](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT "License: MIT")

[English](./README.md) | [简体中文](./README.zh-cn.md)

## 目的

這個項目的目標是提供最簡單、最快速、最無痛的方式來設置自託管的 Git 服務。

GitFX 是 [Gitea](https://github.com/go-gitea/gitea) 的一個分叉。由於它是用 Go 語言編寫的，它可以在 Go 支援的所有平台和架構上運行，包括 Linux、macOS 和 Windows 的 x86、amd64、ARM 和 PowerPC 架構。

## 構建

從源代碼樹的根目錄運行：

    TAGS="bindata" make build

如果需要 SQLite 支援：

    TAGS="bindata sqlite sqlite_unlock_notify" make build

`build` 目標分為兩個子目標：

- `make backend` 需要 [Go Stable](https://go.dev/dl/)，所需版本在 [go.mod](/go.mod) 中定義。
- `make frontend` 需要 [Node.js LTS](https://nodejs.org/en/download/) 或更高版本以及 [pnpm](https://pnpm.io/installation)。

需要互聯網連接來下載 go 和 npm 模塊。從包含預構建前端文件的官方源代碼壓縮包構建時，不會觸發 `frontend` 目標，因此可以在沒有 Node.js 的情況下構建。

## 使用

構建後，默認情況下會在源代碼樹的根目錄生成一個名為 `gitea` 的二進制文件。要運行它，請使用：

    ./gitea web

> [!注意]
> 如果您對使用我們的 API 感興趣，我們提供了實驗性支援，並附有 [文件](https://docs.gitea.com/api)。

## 貢獻

預期的工作流程是：Fork -> Patch -> Push -> Pull Request

> [!注意]
>
> 1. **在開始進行 Pull Request 之前，您必須閱讀 [貢獻者指南](CONTRIBUTING.md)。**
> 2. 如果您在項目中發現了漏洞，請私下寫信給 **security@gitea.io**。謝謝！

## 許可證

這個項目是根據 MIT 許可證授權的。
請參閱 [LICENSE](LICENSE) 文件以獲取完整的許可證文本。

## 歸屬

GitFX 是從 [Gitea](https://github.com/go-gitea/gitea) 分叉而來，Gitea 最初從 [Gogs](https://gogs.io) 分叉。
