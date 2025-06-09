# Base58 Library and CLI Specification

## Overview
Base58エンコーディングライブラリとCLIツールの実装仕様書です。

## Requirements

### Functional Requirements

#### Library (base58パッケージ)
1. **エンコード機能**
   - バイト配列を入力として受け取り、Base58文字列を返す
   - 空のバイト配列に対しては空の文字列を返す
   - 先頭の0バイトは'1'文字として表現する

2. **デコード機能**
   - Base58文字列を入力として受け取り、バイト配列を返す
   - 無効な文字が含まれる場合はエラーを返す
   - 空の文字列に対しては空のバイト配列を返す

3. **文字セット**
   - Bitcoin標準のBase58文字セットを使用: `123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz`
   - 紛らわしい文字(0, O, I, l)は除外

#### CLI Application
1. **エンコードコマンド**
   - 標準入力または引数からデータを読み取り、Base58エンコードして出力
   - ファイル入力のサポート

2. **デコードコマンド**
   - 標準入力または引数からBase58文字列を読み取り、デコードして出力
   - ファイル入力のサポート

3. **ヘルプ機能**
   - 使用方法の表示

### Non-Functional Requirements

1. **パフォーマンス**
   - 大きなデータに対しても効率的に動作する
   - メモリ使用量を最小限に抑える

2. **エラーハンドリング**
   - 適切なエラーメッセージの提供
   - 不正な入力に対する堅牢な処理

3. **テスト**
   - 単体テストによる機能の検証
   - エッジケースのテスト

## API Design

### Library API

```go
package base58

// Encode encodes the given data as a base58 string
func Encode(data []byte) string

// Decode decodes the given base58 string and returns the decoded data
func Decode(s string) ([]byte, error)
```

### CLI Interface

```bash
# エンコード
base58 encode [data]
base58 encode -f <file>
echo "data" | base58 encode

# デコード
base58 decode [base58_string]
base58 decode -f <file>
echo "base58_string" | base58 decode

# ヘルプ
base58 help
base58 -h
base58 --help
```

## Implementation Notes

1. **Algorithm**
   - Big Integer演算を使用した除算ベースの実装
   - 先頭の0バイトの適切な処理

2. **Dependencies**
   - 標準ライブラリのみを使用
   - 外部依存なし

3. **File Structure**
   ```
   base58/
   ├── go.mod
   ├── base58.go       # ライブラリ実装
   ├── base58_test.go  # テスト
   ├── cmd/
   │   └── main.go     # CLIアプリケーション
   ├── SPEC.md         # この仕様書
   └── README.md       # 使用方法
   ```

## Test Cases

### Encoding Tests
- 空データ: `[]byte{}` → `""`
- 単一バイト: `[0x00]` → `"1"`
- 通常データ: `"Hello World"` → 期待値
- バイナリデータ: 各種バイト値の組み合わせ

### Decoding Tests
- 空文字列: `""` → `[]byte{}`
- 単一文字: `"1"` → `[0x00]`
- 無効文字: エラーケース
- 通常文字列: エンコード結果の逆変換

### CLI Tests
- 標準入力/出力の動作確認
- ファイル入力の動作確認
- エラーケースの動作確認