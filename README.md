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

# 最適化版の比較ベンチマーク
make bench-compare

# 最適化版のみのベンチマーク
make bench-optimized

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