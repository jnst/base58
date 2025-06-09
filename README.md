# Base58

Go言語で実装されたBase58エンコーディングライブラリとCLIツールです。

## 特徴

- Bitcoin標準のBase58文字セットを使用
- エンコード/デコード機能
- CLIアプリケーション
- ファイル入力対応
- 標準入力/出力対応
- 包括的なテストスイート

## インストール

```bash
go get github.com/jnst/base58
```

CLIツールをビルド:

```bash
git clone https://github.com/jnst/base58.git
cd base58
go build -o base58 ./cmd
```

## ライブラリ使用例

```go
package main

import (
    "fmt"
    "github.com/jnst/base58"
)

func main() {
    // エンコード
    data := []byte("Hello World")
    encoded := base58.Encode(data)
    fmt.Println(encoded) // JxF12TrwUP45BMd

    // デコード
    decoded, err := base58.Decode(encoded)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(decoded)) // Hello World
}
```

## CLI使用例

### エンコード

```bash
# 引数から
./base58 encode "Hello World"

# 標準入力から
echo "Hello World" | ./base58 encode

# ファイルから
./base58 encode -f input.txt
```

### デコード

```bash
# 引数から
./base58 decode JxF12TrwUP45BMd

# 標準入力から
echo "JxF12TrwUP45BMd" | ./base58 decode

# ファイルから
./base58 decode -f encoded.txt
```

### ヘルプ

```bash
./base58 help
./base58 -h
./base58 --help
```

## API

### Encode

```go
func Encode(data []byte) string
```

バイト配列をBase58文字列にエンコードします。

### Decode

```go
func Decode(s string) ([]byte, error)
```

Base58文字列をバイト配列にデコードします。無効な文字が含まれる場合はエラーを返します。

## 文字セット

Bitcoin標準のBase58文字セットを使用:
```
123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz
```

紛らわしい文字 (0, O, I, l) は除外されています。

## 開発

### 必要なツール

```bash
# 開発環境のセットアップ
make setup
```

### テスト

```bash
# 全テスト実行
make test

# ライブラリテストのみ
go test -v

# CLIテストのみ
cd cmd && go test -v
```

### コード品質チェック

```bash
# フォーマット + リント + テスト
make check

# リントのみ
make lint

# フォーマットのみ
make fmt
```

### ビルド

```bash
# CLI実行ファイルをビルド
make build

# クリーンアップ
make clean
```

## ライセンス

MIT License