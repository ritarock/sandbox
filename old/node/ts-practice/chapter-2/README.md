# tsconfig.json のオプション
生成コマンド
```bash
$ npx tsc --init
```

オプション |説明
-|-
include | TSC が TypeScript ファイルを見つけるためにどのフォルダを探すか
lib | コード実行する環境にどの API が存在していると TSC が想定しておくか
module | TSC がどのモジュールシステムにコンパイルするか
outDir | 生成する JS コードをどのディレクトリに格納するか
strict | 厳格にチェックを行うか
target | TSC がどの JS バージョンにコンパイルするか
