# Golang - Fiber API begin

## API Reference - test localhost:3000

#### Sign Up

```
  POST /users/signup
```

| Parameter   | Type     | Description                 |
| :---------- | :------- | :-------------------------- |
| `full_name` | `string` |                             |
| `email`     | `string` | **Required**. Your email    |
| `password`  | `string` | **Required**. Your password |

#### Sign In

```
  POST /users/signin
```

| Parameter  | Type     | Description                 |
| :--------- | :------- | :-------------------------- |
| `email`    | `string` | **Required**. Your email    |
| `password` | `string` | **Required**. Your password |

#### Create-Books

```
  POST /api/create_books
```

| Parameter   | Type     | Description                       |
| :---------- | :------- | :-------------------------------- |
| `author`    | `string` | **Required**. User_id in database |
| `title`     | `string` | **Required**. Custom              |
| `publisher` | `string` | **Required**. Custom              |

#### Delete-Book-By-ID

```
  DELETE /api/delete_book/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Get-Book-By-ID

```
  GET /api/get_books/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### Get-All-Book

```
  GET /api/books
```
