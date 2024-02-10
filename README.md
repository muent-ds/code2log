# アプリ名: 未定
## 概要
開発におけるドキュメント化・ブログ化を支援するプロダクト

## 技術構成
- Frontend
  - Language : TypeScript
  - Framework : Next.js v14
  - Package Manager : npm
- Backend
  - Language : Python
  - Framework : FastAPI

## ブランチ戦略
**Merge** : {issue番号}-xxx　▶︎　develop　▶︎　main
- main : 本番用ブランチ
- develop : 開発用ブランチ
- {issue番号}-xxx : 各Issueに対応する特定の開発用ブランチ
  - 例 : `12-replace_to_hamburger_menu`

## 開発手順

### Issue駆動開発
[Issue駆動開発について](https://gist.github.com/Enchan1207/0ea2c7a7d6a3c16aea5683435d1972f8)

#### Workflow
- Issue
  - 新規機能の追加やバグの修正を行うにあたって，まずはIssueを作成する．
  - Labelで内容や優先度を明確にする．

- Brunch
  - Issueに基づいてブランチを作成する.
  - ブランチ名 : `{issue番号}-概要`
    - 例 : 12-replace_to_hamburger_menu

- Commit
  - commit message : `[{commit種別}-#{issue番号}]変更内容の要約`
    - 例 : `[update-12]replace menu button with hamburger menu`
  - `#{issue番号}`を含めることでissueと相互参照可能になる．

- Pull Request
  - 変更が済んだら作成したブランチから`develop`ブランチへのPRを作成する．
  - PRのサイドバーにある`Development`をクリックし，`Link an issue from this repository`からIssueをリンクする

#### Label
- TYPE
  - enhancement : 新規機能の追加
  - bug fix : バグの修正
  - improvement : 既存機能の改良
  - modification : 既存機能の修正
  - security fix : セキュリティに関する問題の修正
- STATUS
  - available : 待機中
  - in progress : 作業中
  - done : 完了
  - canceled : 中止
  - pending : 保留中
  - unnecessary : 着手不要
  - duplicate : 重複した内容
- PRIORITY
  - asap : 最優先
  - high : 高
  - medium : 中
  - low : 低
  - safe : 余裕があれば取り組む
 
#### commit種別
- add : 新規ファイル・新規機能の追加
- fix : バグの修正
- update : 既存機能の変更
- remove : ファイル等の削除
- refacter : リファクタリング
- upgrade : バージョンの更新
- revert : 変更の取り消し
