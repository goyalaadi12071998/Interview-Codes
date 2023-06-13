# Prerequisite
1. Docker
2. MySql

# HOW TO RUN PROJECT
0. Clone this repo in your local
1. Install Mysql
2. If you are running the application using docker, then please change the username and password in configs/docker.toml file accoding to your   local mysql username and password.
3. Create DATABASE bitespeed
4. Create a TABLE name contacts

   ```
   CREATE TABLE contacts ( id bigint NOT NULL PRIMARY KEY AUTO_INCREMENT, phone_number varchar(255), email varchar(255), linked_id int, link_preference varchar(255), created_at bigint, updated_at bigint, deleted_at bigint )
   ```
4. RUN docker-compose up --build
5. Project is running



# Note:

1. At lot of places the code is like

```
    if err != nil {
        return err
    }

    return nil
```

which is basically the same thing but i write it as like this because if err != nil then we can add logs of failure

2. There is one test case which is not present in the sheet.

Assuming the current state of database is 

```
+----+--------------+-------------------+-----------+-----------------+---------------+---------------+------------+
| id | phone_number | email             | linked_id | link_preference | created_at    | updated_at    | deleted_at |
+----+--------------+-------------------+-----------+-----------------+---------------+---------------+------------+
| 44 | 111111       | aadi+15@gmail.com |         0 | primary         | 1686633110091 | 1686633110091 |          0 |
| 45 | 222222       | aadi+16@gmail.com |        44 | secondary       | 1686633116922 | 1686633121414 |          0 |
| 46 | 333333       | aadi+17@gmail.com |         0 | primary         | 1686633178602 | 1686633178602 |          0 |
| 47 | 333333       | aadi+18@gmail.com |        46 | secondary       | 1686633191246 | 1686633191246 |          0 |
| 48 | 444444       | aadi+19@gmail.com |         0 | primary         | 1686633253388 | 1686633253388 |          0 |
| 49 | 555555       | aadi+19@gmail.com |        48 | secondary       | 1686633264886 | 1686633264886 |          0 |
+----+--------------+-------------------+-----------+-----------------+---------------+---------------+------------+
```

And some hit a req - 

```
{
    "email":"aadi+18@gmail.com",
    "phoneNumber": "555555"
}
```

Now both the different contacts are already secondary
