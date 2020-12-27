# panenga-api

## 起動
`docker-compose up -d --build`

## コンテナの削除
`docker-compose down --volumes`

## 命名規則
### ファイル名：snake_case
"物（単数形か複数形かも考慮）" + "アンダーバー " + "メソッド名" .go
例）skills_get.go
### 関数名・構造体名：UpperCamelCase
例）GetSkills()
### 変数名：lowerCamelCase
例）skillData 

## 各ディレクトリの記述内容（MVC）

- Modelについて
	Modelにはデータベースに変更を加える処理を書き込む.
- Controller(Handler)について
	ControllerはModelやViewからメソッドを呼び出し，JSONをフロントに返す.
- Viewについて
    Viewにはレスポンスを成形する処理を書く.

## ブランチ名について
- master
プロダクトとしてリリースするためのブランチ.
- develop
開発ブランチ. リリース前はこのブランチが最新バージョンとなる.
- feature
機能の追加をするためのブランチ.
feat/#{対応するissue番号}/{エンドポイント名など内容の詳細}
- fix
変更・修正用のブランチ.
fix/#{対応するissue番号}/{内容の詳細}

## コミットメッセージについて

以下のフォーマットとする.
```
1行目：変更内容の要約や概要）
2行目 ：空行
3行目以降：変更した理由や内容の詳細
```

##### 1行目のフォーマット
- [fix]：バグなどの修正

- [update]：機能修正

- [wip]：作業中（WIP：Work In Progress）

- [clean]：整理
