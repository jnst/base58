# Base58 Performance Optimization Plan

現在のベンチマーク結果に基づいて、パフォーマンス改善案を分析・提案します。

## 現在の性能問題

### 1. メモリアロケーション過多
- **32Bデータ**: 46 allocs/op（エンコード）
- **1KBデータ**: 1,377 allocs/op（エンコード）
- **16KBデータ**: 22,008 allocs/op（エンコード）
- **原因**: big.Int演算での頻繁なメモリ確保

### 2. 非線形的な性能劣化
- **時間計算量**: 理論的にはO(n)だが、実際はO(n²)に近い
- **16KBデータ**: 129ms（期待値の10倍以上）
- **原因**: big.Int演算のコストがデータサイズに対して非線形

### 3. エンコード vs デコードの性能差
- **エンコード**: より多くのメモリとアロケーション
- **デコード**: 相対的に効率的
- **原因**: エンコード時の除算処理とバッファ管理

## 改善案

### 優先度A: メモリ最適化

#### 1. big.Intプールの導入
```go
var bigIntPool = sync.Pool{
    New: func() interface{} {
        return new(big.Int)
    },
}
```
- **効果**: アロケーション回数を70-80%削減
- **実装コスト**: 低
- **リスク**: 低

#### 2. バッファサイズの最適化
```go
// より正確なサイズ計算
func calculateBufferSize(dataLen int) int {
    // log(256)/log(58) = 1.3658...
    return (dataLen * 1366) / 1000 + 2
}
```
- **効果**: メモリ使用量を10-15%削減
- **実装コスト**: 低
- **リスク**: 低

### 優先度B: アルゴリズム改善

#### 3. 先頭ゼロバイトの特別処理
```go
// 先頭ゼロを除外してからbig.Int演算
func optimizedEncode(data []byte) string {
    leading := countLeadingZeros(data)
    if leading == len(data) {
        return strings.Repeat("1", leading)
    }
    // 実際のデータのみでbig.Int演算
    result := bigIntEncode(data[leading:])
    return strings.Repeat("1", leading) + result
}
```
- **効果**: 大量のゼロバイトで大幅改善
- **実装コスト**: 中
- **リスク**: 低

#### 4. チャンク処理の導入
```go
const maxChunkSize = 1024 // 1KB

func chunkedEncode(data []byte) string {
    if len(data) <= maxChunkSize {
        return currentEncode(data)
    }
    // 大きなデータを分割処理
}
```
- **効果**: 大きなデータで線形性能を維持
- **実装コスト**: 高
- **リスク**: 中（正確性の検証が必要）

### 優先度C: 詳細最適化

#### 5. 文字列操作の最適化
```go
func optimizedStringBuild(size int) *strings.Builder {
    var builder strings.Builder
    builder.Grow(size) // 事前にサイズを確保
    return &builder
}
```
- **効果**: 文字列連結で5-10%改善
- **実装コスト**: 低
- **リスク**: 低

#### 6. ルックアップテーブルの最適化
```go
// 配列ベースのより高速な実装
var alphabetLookup [256]int

func init() {
    for i := range alphabetLookup {
        alphabetLookup[i] = -1 // 無効値
    }
    for i, char := range []byte(alphabet) {
        alphabetLookup[char] = i
    }
}
```
- **効果**: デコードで5-10%改善
- **実装コスト**: 低
- **リスク**: 低

## 実装順序

### フェーズ1: 即効性のある改善
1. big.Intプール（週内）
2. バッファサイズ最適化（週内）
3. 文字列操作最適化（週内）

### フェーズ2: アルゴリズム改善
1. 先頭ゼロバイト特別処理（1-2週間）
2. ルックアップテーブル最適化（1週間）

### フェーズ3: 大規模改善
1. チャンク処理の導入（2-3週間）
2. 並列処理の検討（必要に応じて）

## 期待効果

### 小さなデータ（32B）
- **現在**: 1.6μs, 504B, 46 allocs
- **予想**: 0.8μs, 200B, 15 allocs
- **改善率**: 50%高速化、60%メモリ削減

### 中サイズデータ（1KB）
- **現在**: 534μs, 15KB, 1,377 allocs
- **予想**: 200μs, 8KB, 400 allocs
- **改善率**: 60%高速化、50%メモリ削減

### 大きなデータ（16KB）
- **現在**: 129ms, 244KB, 22,008 allocs
- **予想**: 20ms, 100KB, 5,000 allocs
- **改善率**: 85%高速化、60%メモリ削減

## 測定・検証計画

### ベンチマーク比較
```bash
# 現在の実装をベースラインとして保存
go test -bench=. -benchmem > baseline.txt

# 各改善後に比較
go test -bench=. -benchmem > optimized_v1.txt
benchcmp baseline.txt optimized_v1.txt
```

### 正確性検証
```bash
# ランダムデータでのラウンドトリップテスト
go test -run=TestOptimizedRoundTrip -count=1000
```

### パフォーマンス監視
- 各フェーズ後にベンチマーク実行
- リグレッション検出
- 継続的な性能監視

## リスク管理

### 正確性リスク
- **対策**: 包括的なテストスイート
- **検証**: 既存テストを全て通すことを確認
- **追加**: ランダムデータでのfuzz testing

### 互換性リスク
- **対策**: 公開APIの変更なし
- **検証**: 既存のクライアントコードが動作することを確認

### 複雑性リスク
- **対策**: 段階的な実装
- **検証**: コードレビューと十分なテスト
- **回避**: 問題があれば前の実装に戻せる構造