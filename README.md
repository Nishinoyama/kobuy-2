

## 設計ルール

+ Handler
  + Httpリクエストを受け取る
  + レスポンスを返す
+ Controller
  + レスポンスを作成する
+ Service
  + エンティティを作成する、DBから引き寄せる

## go entの罠とその回避

### デフォルトでstruct tagが `omitempty`

+ `0`や`nil`などに意味があっても、そのkey-valueを消されてしまうので厄介

#### 回避策

+ schema内のFieldに``StructTag(`json:"field"`)``を追加