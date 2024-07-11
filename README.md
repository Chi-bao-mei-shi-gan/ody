# Ody

`Ody` 是一個用於從遠端存儲庫克隆代碼並替換指定名稱的小工具。

## 特性

- 從遠端 Git 存儲庫克隆代碼
- 遞歸地遍歷所有文件並替換指定名稱
- 刪除 `.git` 目錄並重命名目錄

## 安裝

確保已經設置 `GOPRIVATE` 和 `GOINSECURE` 環境變量，以允許從自架的 GitLab 伺服器下載模組：

```sh
export GOPRIVATE=https://github.com/Chi-bao-mei-shi-gan
export GOINSECURE=https://github.com/Chi-bao-mei-shi-gan

# 添加到 .bashrc 或 .zshrc 以便每次打開終端時都生效
echo 'export GOPRIVATE=https://github.com/Chi-bao-mei-shi-gan' >> ~/.bashrc
echo 'export GOINSECURE=https://github.com/Chi-bao-mei-shi-gan' >> ~/.bashrc
source ~/.bashrc
```

使用 go install 安裝 Ody：

```sh
go install github.com/Chi-bao-mei-shi-gan/ody@latest
go install github.com/Chi-bao-mei-shi-gan/ody@3996df2
```

## 使用

安裝完成後，可以使用 `ody` 命令從遠端存儲庫克隆代碼並替換指定名稱：

```sh
ody -n <newName>
```

要顯示幫助信息，可以使用 -h 標誌：

```sh
ody -h
```
