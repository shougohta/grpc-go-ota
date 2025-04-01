## ✅ セルフレビュー観点表


### 🗂 共通
- [ ] 空文字のは許容するのか？しない場合、考慮はできているか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] 自動生成されたフロントの型定義は正しいか？undifinedはないか？nullableは適切か？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] 破壊的変更の場合も、再度に明示的なreturnをしているか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] IDなどを不用意にハードコードしてないか？（環境によって変動可能性がある）
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] schemaファイルやクエリ系のインデントは綺麗かつ揃っているか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] 既存コードのコピペ時に、適切な変数名に変えているか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] PRコメントはリロードしてから見る（たまにコメントが一部反映されない）
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] nullと0の区別はされているか？（そもそも区別が必要か検討する）
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] コミットメッセージは意味のあるまとまりになっているか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] 不要な変更＝ゴミが残っていないか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] メソッド、変数、キーの名称は適切か？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] コメントが分かりやすく書かれているか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] スキップすべきパターンは網羅されているか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] nil対応は適切か？(?、compact,&.,早期リターン)
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] エラーハンドリングは入っているか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] クラスの責務は単一か？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>


### 🗂 Ruby
- [ ] N＋1になっていないか？preload, eager_loadの使い分けは正しいか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] 不要なDBアクセスを行っていないか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] 削除処理において、子レコードの削除対応が漏れていないか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] コントローラーのエンドポイントのコメントは正しいか？コピペになってないか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] service層でcreate,update,deleteする場合、usecase層でトランザクションを張っているか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] formクラスでのバリデーションは.filled(:string)のように空欄除去かつ一意データ型か？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] メインリソースの取得はfindにしているか？※リソースをfind取ってる場合は自動で404になる
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] メインリソース以外の取得でfindを使っていないか？※find_byにしないと、見つからない時例外発生する
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] specはshared_context,shared_exampleでまとめているか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] レコードが増えた際にパフォーマンスが劣化しないか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] 配列処理は適切か？（ruby：each,find_each,map）
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] モデルのscope処理は適切か？pluckを使っていないか？（配列になる＝ActiveRecordRelationインスタンスで無くなるので、遅延評価でなくなり、複数回クエリ発行される）
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] blank?とpresent?になっているか？不用意にnil?やempty?にしてないか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] 関連付けは適切か？(has_one,has_many,dependant)
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>


### 🗂 Javascript
- [ ] 配列処理は適切か？（javasciript：reduce,filter,map）
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>


### 🗂 Go
- [ ] パッケージ内のあるべきファイルに関数を定義しているか（同じパッケージであれば、頭文字を大文字にするだけで関数を呼べる）
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] ログが残っていないか
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] 不用意にメソッド名が長くなってないか？workspaceなど、自明な単語を含んでいないか？
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] 不用意に型変換していないか？（本来型変換不要なので元のschemaを修正する）
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

- [ ] 不用意にnull許容をしていないか？（例：null.String）
  <details>
    <summary>Good / Bad 例</summary>

    ✅ **Good**  
    - 例1  
    - 例2  

    ❌ **Bad**  
    - 例1  
    - 例2  
  </details>

