# go-tutorial

### wsl:ubuntu-LTS mysql起動
```
$ sudo service mysql start
```

### mysqlログイン
```
$ sudo mysql -u USER -p
Enter password: ****************
```

### データベース一覧表示
```
mysql> show databases;
```

### データベース作成
```
mysql> create database DATABASE_NAME;
```

### データベース選択
```
mysql> use DATABASE_NAME;
```

### テーブル作成
```
CREATE TABLE `userinfo` (
    `uid` INT(10) NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(64) NULL DEFAULT NULL,
    `departname` VARCHAR(64) NULL DEFAULT NULL,
    `created` DATE NULL DEFAULT NULL,
    PRIMARY KEY (`uid`)
);

CREATE TABLE `userdetail` (
    `uid` INT(10) NOT NULL DEFAULT '0',
    `intro` TEXT NULL,
    `profile` TEXT NULL,
    PRIMARY KEY (`uid`)
);

mysql> create table `members` (
    `id`        int not null auto_increment primary key,
    `name`      varchar(100) not null,
    `email`     varchar(100) not null,
    `password`  varchar(100) not null
) engine=InnoDB default charset=utf8;
```
