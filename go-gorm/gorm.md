# active record / gorm 比較

## create
### AR
```ruby
user = User.create(name: 'hoge')

user = User.new(name: 'hoge')
user.save
```

### gorm
```go
user  := User{Name: "hoge"}
db.Create(&user)
```

## update
### AR
```ruby
user.update(name: 'hoge')
```

### gorm
```go
user.Name = "hoge"
db.Save(&user)

# 指定したカラムだけ
db.Model(&user).Update("name": "hoge")
db.Model(&user).Updates(User{Name: "hoge"})
```

## delete
### AR
```ruby
user.delete
user.destroy
```

### gorm
```go
db.Delete(&user)

# 物理削除
db.Unscoped().Delete(&user)
```

## 検索
### 最初の1件
```ruby
User.first
```

```go
db.First(&user)
```

### 最後の1件
```ruby
User.last
```

```go
db.Last(&user)
```

### 条件検索
```ruby
User.where(name: 'hoge')

User.where(name: 'hoge').where(text: 'fuga')
```

```go
db.Where("name = ?", "hoge").Find(&users)

db.Where("name = ?", "hoge").Where("text = ?").Find(&users)
```
