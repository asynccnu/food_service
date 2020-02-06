## food_service

> 命名规范

- 数据库字段，表名，json，都用下划线。
- Go语言中的变量，驼峰命名。

> 美食服务

### 食堂信息

id从1递增
```
  ("学子", 1),
  ("学子", 2),
  ("东一", 1),
  ("东一", 2),
  ("东二", 1),
  ("东二", 2),
  ("桂香园", 1),
  ("桂香园", 2),
  ("桂香园", 3),
  ("博雅园", 1);
```
### 三个栏目的介绍

你的附近， 店家信息，(此处不显示图片)有图片。店家名称，所在食堂，特色推荐(多个)。(此处不显示)人均消费(适用于在线菜单)

华师必吃， 产品信息，有图片。菜名，店家名称，所在食堂，主要原料，简单介绍。

在线菜单， 八个食堂的名字，食堂各楼的店家(人均消费)，店家详情页(图片，名称，人均，店家介绍，菜单(栏目))

----------
店家信息 图片 店家名称 所在食堂 特色推荐(多个) 人均消费 店家介绍 菜单(名字和价格)

产品信息 图片 店家名称 所在食堂 主要原料  菜名 价格 简单介绍

### 搜索服务

热门搜索

搜索历史

模糊搜索 (美食名称，店家名称，区分两名字)

restaurant
```sql
id          int
picture url string
name        string
location    string
sales_volumn int
introduction string
```

food
```sql
id          int
picture url string
name        string
location    string #不一定存
ingredients string
price       string
introduction string
```