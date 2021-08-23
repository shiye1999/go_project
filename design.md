数据库中表格式如下

```
CREATE TABLE user_info
(
    uid serial NOT NULL,
    username character varying(100) NOT NULL,
    created date,
    CONSTRAINT user_pkey PRIMARY KEY (uid)
);

CREATE TABLE comment
(
    id serial NOT NULL,
    uid integer,
    created date,
    intro character varying(1000) NOT NULL,
    CONSTRAINT comment_pkey PRIMARY KEY (id)
);
```

![image-20210823163921294](C:\Users\25328\AppData\Roaming\Typora\typora-user-images\image-20210823163921294.png)

![image-20210823165642974](C:\Users\25328\AppData\Roaming\Typora\typora-user-images\image-20210823165642974.png)