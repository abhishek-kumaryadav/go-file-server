
# Project Title

A simple golang server to upload, view, search and download files


## API Reference

#### Get all files

```http
  GET /file-sharing
```

#### Get a file

```http
  GET /file-sharing/file/{file-name}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `file-name`      | `string` | **Required**. Name of the file to fetch |

#### Search for a file using partial name

```http
  GET /file-sharing/search/{file-name}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `file-name`      | `string` | **Required**. Partial name of the file to search |

#### Upload a file

```http
  POST /file-sharing
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `file`      | `file` | **Required**. Attach the file to be uploaded |



## Authors

- [@abhishek-kumaryadav](https://github.com/abhishek-kumaryadav)

