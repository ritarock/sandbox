# USASE
マイグレーションファイルを作成。
```bash
$ docker-compose -f docker-compose.dev.yml run generation { arg }
```

実行。
```bash
$ docker-compose up
```

# API
## users
- `GET    /users` - ユーザ一覧を取得
- `POST   /users` - ユーザを新規作成
- `GET    /users/:user_id` - user_id を持つユーザを取得
- `PUT    /users/:user_id` - user_id を持つユーザを更新
- `DELETE /users/:user_id` - user_id を持つユーザを削除

## tasks
- `GET    /users/:user_id/tasks` - user_id を持つユーザのタスク一覧を取得
- `POST   /users/:user_id/tasks` - user_id を持つユーザにタスクを追加
- `GET    /users/:user_id/tasks/:task_id` - user_id を持つユーザの task_id を取得
- `PUT    /users/:user_id/tasks/:task_id` - user_id を持つユーザの task_id を更新
- `DELETE /users/:user_id/tasks/:task_id` - user_id を持つユーザの task_id を削除

## filter
- `GET    /users/:user_id/tasks`
  - `/users/:user_id/tasks?state=new` - 未着手のタスクを取得
  - `/users/:user_id/tasks?state=working` - 進行中のタスクを取得
  - `/users/:user_id/tasks?state=completed` - 完了済のタスクを取得
  - `/users/:user_id/tasks?state=pending` - 保留中のタスクを取得
