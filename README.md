# GOプログラミング実践入門
- https://github.com/mushahiroyuki/gowebprog/blob/master/ch02/chitchat/data/data.go

## メソッド
特定の型に関連付けられた関数

## ユーティリティ関数
関数の単なる集まり

## 空インターフェイス
どんな型でも良い？？
- Goには静的型付けプログラミング言語であることを回避する巧妙な方法があり、インターフェイスを用いることで異なる型を受け付ける柔軟性を持っている。

## zshrcを更新する
```
source .zshrc
```
[参考リンク](http://norizo333.hatenablog.com/entry/20090921/1253523438)

# postgresDBの作成
[参考記事](https://teratail.com/questions/243773)
`psql -U postgres`
`postgres=# CREATE ROLE {username} LOGIN CREATEROLE PASSWORD {password}`