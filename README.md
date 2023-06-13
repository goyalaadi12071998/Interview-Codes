# HOW TO RUN PROJECT

#Prerequisite
1. Docker
2. MySql

1. Install Mysql
2. Create DATABASE bitespeed
3. Create a TABLE name contacts

   ```
   CREATE TABLE contacts ( id bigint NOT NULL PRIMARY KEY AUTO_INCREMENT, phone_number varchar(255), email varchar(255), linked_id int, link_preference varchar(255), created_at bigint, updated_at bigint, deleted_at bigint )
   ```
4. RUN docker-compose up --build
5. Project is running



# Note:

At lot of places the code is like

```
    if err != nil {
        return err
    }

    return nil
```

which is basically the same thing but i write it as like this because if err != nil then we can add logs of failure