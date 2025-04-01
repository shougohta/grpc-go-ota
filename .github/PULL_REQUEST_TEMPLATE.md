🧩 共通項目
<details> <summary>空文字のは許容するのか？しない場合、考慮はできているか？</summary> <details> <summary>✅ Good</summary> - フロントで入力制限されており、空文字がそもそも送信されない - 空文字の場合はデフォルト値をセットしている </details> <details> <summary>❌ Bad</summary> - バリデーションがなく、空文字で保存されてしまう - 空文字とnullの扱いが混在している </details> </details> <details> <summary>自動生成されたフロントの型定義は正しいか？undefinedはないか？nullableは適切か？</summary> <!-- Good / Badを必要に応じて追加 --> </details> <details> <summary>破壊的変更の場合も、再度に明示的なreturnをしているか？</summary> </details> <!-- 以下、他の共通項目も同様に展開 -->
🐍 Ruby
<details> <summary>N+1になっていないか？preload, eager_loadの使い分けは正しいか？</summary> <details> <summary>✅ Good</summary> - `includes(:association)` でN+1回避されている - `preload`と`eager_load`の違いを理解して選択されている </details> <details> <summary>❌ Bad</summary> - `each`内で `.association` を呼んでいる（クエリが都度発行される） - コメントなしで `eager_load` 使用（不要なJOINでパフォーマンス低下） </details> </details> <!-- 以下、他のRuby項目 -->
💻 Javascript
<details> <summary>配列処理は適切か？（reduce, filter, map）</summary> </details>
🐹 Go
<details> <summary>ログが残っていないか</summary> </details>