# go-cp
## 概要
windowsでunixのcpコマンドを使いたいのでせっかくだし自作する。  
今は1つ以上のファイルを他のdirにコピーする機能だけ必要なので実装している。  
今後必要なら増やしていく。 
## 使い方
下記のコマンドでインストールする  
``` go install github.com/kimuson13/go-cp@latest ```
そして  
```go-cp [コピーしたいファイルのパス(1つ以上)] [コピー先のdir]```  
で使うことができる。
## 今後の展開
- [ ] copy先をdir以外にもできるようにする。
- [ ] フラッグの実装
