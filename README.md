# Base58

## Base58とは

Base58は、人間にとって読みやすくタイプしやすいよう設計されたエンコーディング方式です。Base64エンコーディングから以下の紛らわしい文字を除外したものです：

- `0` (数字のゼロ) - 大文字のOと混同しやすい
- `O` (大文字のオー) - 数字の0と混同しやすい
- `I` (大文字のアイ) - 小文字のlと混同しやすい
- `l` (小文字のエル) - 大文字のIと混同しやすい

**主な用途：**
- Bitcoin アドレス
- 暗号通貨のウォレットアドレス
- 短縮URL
- ユーザーフレンドリーな識別子

## 概要

このライブラリは、Go言語で実装された高性能なBase58エンコーディング/デコーディングライブラリとCLIツールです。Bitcoin標準の文字セットを使用し、オブジェクトプールによる最適化でメモリアロケーションを大幅に削減しています。

## 特徴

- **Bitcoin標準**: Bitcoin標準のBase58文字セット `123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz` を使用
- **高性能**: オブジェクトプール最適化により最大99.9%のメモリアロケーション削減
- **ライブラリ + CLI**: ライブラリとしても、コマンドラインツールとしても使用可能
- **柔軟な入力**: 引数、標準入力、ファイル入力に対応
- **信頼性**: 包括的なテストスイート、ファズテスト、ベンチマーク
- **依存関係なし**: 標準ライブラリのみを使用

## 要件

- Go 1.18 以降（`any`型を使用するため）

## インストール

### ライブラリとして使用
```bash
go get github.com/jnst/base58
```

### CLIツールのインストール
```bash
go install github.com/jnst/base58/cmd/base58@latest
```

### ソースからビルド
```bash
git clone https://github.com/jnst/base58.git
cd base58
make build
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

### 標準版

#### Encode

```go
func Encode(data []byte) string
```

バイト配列をBase58文字列にエンコードします。

#### Decode

```go
func Decode(s string) ([]byte, error)
```

Base58文字列をバイト配列にデコードします。無効な文字が含まれる場合はエラーを返します。

### パフォーマンス

高性能実装により、メモリアロケーションを大幅に削減：

| データサイズ | アロケーション数 | メモリ使用量 |
|-------------|-----------------|-------------|
| 32B         | 2 allocs | 96 B/op |
| 1KB         | 2 allocs | 2,818 B/op |
| 4KB         | 2 allocs | 12,338 B/op |

詳細は [OPTIMIZATION_RESULTS.md](OPTIMIZATION_RESULTS.md) を参照してください。

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
go test -run=TestCLI

# ファズテストによる正確性テスト
go test -run=TestFuzz
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

### ベンチマーク

```bash
# 全ベンチマーク実行
make bench

# 複数回実行で安定性確認
make bench-stability

# ベンチマーク結果をファイルに保存
make bench-save
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