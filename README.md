## Final Project Virtual Intern BTPNS

### Initiating Go Backend Application
1. Lakukan konfigurasi file ```.env``` sesuai file ```.env.example```
2. Pastikan sudah melakukan instalasi go dengan menjalankan command ```go version```
3. Jalankan command ```go run main.go```

### Tech Stack
1. Go
2. Gin
3. Gorm
4. JWT
5. MySQL

<hr>

### Documentation

#### List Entities 
1. [User](#User)
2. [Photo](#Photo)

#### User
1. ##### Register User ```POST METHOD```
    endpoint : ```/auth/register```    
    json request body : 
    ```
    {
    "username": "taufiqoo",
    "password": "taufiq123",
    "email": "taufiq@gmail.com"
    }
    ```
    json response : 
    ```
    {
    "message": "successfully register user",
    "status": 201
    }
    ```
    json response missing value : 
    ```
    {
    "error": "UserAuth.password: non zero value required;UserAuth.username: non zero value required;email: non zero value required",
    "message": "bad body request",
    "status": 400
    }
    ```
    json response validation password : 
    ```
    {
    "error": "UserAuth.password: pass1 does not validate as stringlength(6|255)",
    "message": "bad body request",
    "status": 400
    }
    ```
    json response validation email : 
    ```
    {
    "error": "email: xxx.yyy does not validate as email",
    "message": "bad body request",
    "status": 400
    }
    ```
    json response validation conflict : 
    ```
    {
    "message": "username or email is already used",
    "status": 409
    }
    ```
2. ##### Login User ```GET METHOD```
    endpoint : ```/auth/login```    
    json request body : 
    ```
    {
    "username": "taufiqoo",
    "password": "taufiq123",
    }
    ```
    json response : 
    ```
    {
    "message": "user successfully login",
    "status": 202
    }
    ```
3. ##### Logout User ```GET METHOD```
    endpoint : ```/auth/logout```   
    json response : 
    ```
    {
    "message": "successfully logout",
    "status": 200,
    }
    ```
4. ##### Update User ```PUT METHOD```
    endpoint : ```/users```    
    json request body : 
    ```
    {
    "username": "taufiqoo",
    "password": "taufiq321",
    "email": "taufiqoo@gmail.com"
    }
    ```
    json response : 
    ```
    {
    "message": "successfully update data",
    "status": 200
    }
    ```
5. ##### Delete User ```DELETE METHOD```
    endpoint : ```/users```    
    json response : 
    ```
    {
    "message": "successfully delete data",
    "status": 200
    }
    ```

#### Photo
1. ##### Create New Photo ```POST METHOD```
    endpoint : ```/photos```    
    json request body :
    ```
    {
    "title": "backend developer golang",
    "caption": "What is a Golang Developer and What Do they Do?",
    "photo_url": "https://fullscale.io/wp-content/uploads/2022/04/golang-developer.png"
    }
    ```
    json response : 
    ```
    {
    "data_record": {
        "ID": 1,
        "CreatedAt": "2023-09-07T22:21:55.804+07:00",
        "UpdatedAt": "2023-09-07T22:21:55.804+07:00",
        "DeletedAt": null,
        "user_id": 1,
        "title": "backend developer golang",
        "caption": "What is a Golang Developer and What Do they Do?",
        "photo_url": "https://fullscale.io/wp-content/uploads/2022/04/golang-developer.png"
    },
    "message": "successfully create photo",
    "status": 201
    }
    ```
    json response if fields missing :
    ```
    {
    "error": "caption: non zero value required;photo_url: non zero value required;title: non zero value required",
    "message": "bad body request",
    "status": 400
    }
    ```
2. ##### Get All Photos ```GET METHOD```
    endpoint : ```/photos```   
    json response : 
    ``` 
    {
    "data": [
        {
            "ID": 1,
            "CreatedAt": "2023-09-07T15:21:55.804Z",
            "UpdatedAt": "2023-09-07T15:21:55.804Z",
            "DeletedAt": null,
            "user_id": 1,
            "title": "backend developer golang",
            "caption": "What is a Golang Developer and What Do they Do?",
            "photo_url": "https://fullscale.io/wp-content/uploads/2022/04/golang-developer.png"
        },
        {
            "ID": 2,
            "CreatedAt": "2023-09-07T15:25:34.296Z",
            "UpdatedAt": "2023-09-07T15:25:34.296Z",
            "DeletedAt": null,
            "user_id": 1,
            "title": "becmoing a software engineer",
            "caption": "Are you interested in a career as a software engineer but are not sure how to start preparing for it? We have you covered!",
            "photo_url": "https://i3.ytimg.com/vi/9wF-gv-isOY/maxresdefault.jpg"
        }
    ],
    "message": "successfully fetch photos",
    "status": 200
    }
    ```

3. ##### Get Photos By User ID ```GET METHOD```
    endpoint : ```/photos?user={number}```    
    json response : 
    ``` 
    {
    "data": [
        {
            "ID": 1,
            "CreatedAt": "2023-09-07T15:21:55.804Z",
            "UpdatedAt": "2023-09-07T15:21:55.804Z",
            "DeletedAt": null,
            "user_id": 1,
            "title": "backend developer golang",
            "caption": "What is a Golang Developer and What Do they Do?",
            "photo_url": "https://fullscale.io/wp-content/uploads/2022/04/golang-developer.png"
        },
        {
            "ID": 2,
            "CreatedAt": "2023-09-07T15:25:34.296Z",
            "UpdatedAt": "2023-09-07T15:25:34.296Z",
            "DeletedAt": null,
            "user_id": 1,
            "title": "becmoing a software engineer",
            "caption": "Are you interested in a career as a software engineer but are not sure how to start preparing for it? We have you covered!",
            "photo_url": "https://i3.ytimg.com/vi/9wF-gv-isOY/maxresdefault.jpg"
        }
    ],
    "message": "successfully fetch photos",
    "status": 200
    }
    ```

4. ##### Get Photo By ID ```POST METHOD```
    endpoint : ```/photos/:id```    
    json response : 
    ```
    {
    "data": {
        "ID": 1,
        "CreatedAt": "2023-09-07T15:21:55.804Z",
        "UpdatedAt": "2023-09-07T15:21:55.804Z",
        "DeletedAt": null,
        "user_id": 1,
        "title": "backend developer golang",
        "caption": "What is a Golang Developer and What Do they Do?",
        "photo_url": "https://fullscale.io/wp-content/uploads/2022/04/golang-developer.png"
    },
    "message": "successfully fetch photo",
    "status": 200
    }
    ```

5. ##### Update Photo ```PUT METHOD```
    endpoint : ```photos/1```    
    json response : 
    ```
    {
    "data_record": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "2023-09-07T22:37:05.791+07:00",
        "DeletedAt": null,
        "user_id": 0,
        "title": "becmoing a software engineer in the world",
        "caption": "Are you interested in a career as a software engineer but are not sure how to start preparing for it? We have you covered!",
        "photo_url": "https://i3.ytimg.com/vi/9wF-gv-isOY/maxresdefault.jpg"
    },
    "message": "successfully update data",
    "status": 200
    }
    ```
    json response failed update : 
    ```
    {
    "message": "failed to update data",
    "status": 400
    }
    ```

6. ##### Delete Photo ```DELETE METHOD```
    endpoint : ```/photos/:id```    
    json response : 
    ```
    {
    "message": "successfully delete data",
    "status": 200
    }
    ```
    json response if fail or try to delete another use photo : 
    ```
    {
    "message": "failed to delete data",
    "status": 400
    }
    ```
