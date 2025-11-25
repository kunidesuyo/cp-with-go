# メモ

## 標準入出力

### 入力

* 1単語ずつ

```go
sc := bufio.NewScanner(os.Stdin)
sc.Split(bufio.ScanWords)
sc.Scan()
```
