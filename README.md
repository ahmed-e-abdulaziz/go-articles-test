# go-articles-test

Simple Articles Backend in Go

## What's this project?

This is a simple demonstration of how I write Go code, as I have been requested to do so on multiple occasions.
You can create articles and fetch them, add comments to articles, and fetch comments on articles.
I did it in about 5 hours of work over a few days
Feel free to contact me if you want to discuss this repo.

## How to run?

Use the `make run-local` command for an easy local run, and ensure the correct database host and credentials as in the `local_db_env_vars_init.sh`.

Or you can use `make run` but make sure to expose the same env vars as in `local_db_env_vars_init.sh`

## API

All the endpoints are available under a versioned system. The current version is `v1`

### Fetch All Articles

const currentApiVersionUri = "/v1"
const articlesUri = currentApiVersionUri + "/articles"
const commentsUri = articlesUri + "/:id/comments"

**Endpoint:** `/v1/articles GET`
**Response Body:**

```json
[
 {
    "id": 1,
    "title": "Awesome Go",
    "content": "A curated list of awesome Go frameworks, libraries, and software",
    "creation_timestamp": "2024-12-11T09:02:20.715864Z"
 },
 {
    "id": 2,
    "title": "Awesome Java",
    "content": "A curated list of awesome Java frameworks, libraries, and software",
    "creation_timestamp": "2024-12-11T09:02:31.029818Z"
 },
 {
    "id": 3,
    "title": "Awesome Javascript",
    "content": "A collection of awesome browser-side JavaScript libraries, resources, and shiny things.",
    "creation_timestamp": "2024-12-11T11:01:52.887083Z"
 },
 {
    "id": 4,
    "title": "Awesome Python",
    "content": "A collection of awesome browser-side Python libraries, resources, and shiny things.",
    "creation_timestamp": "2024-12-11T11:02:00.839645Z"
 }
]
```

### Fetch Article By ID

**Endpoint:** `/v1/articles/{id} GET`
**Path Param:** *id*: The id of the requested article
**Response Body:**

```json
{
    "id": 1,
    "title": "Awesome Go",
    "content": "A curated list of awesome Go frameworks, libraries, and software",
    "creation_timestamp": "2024-12-11T09:02:20.715864Z"
}
```

### Add Article

**Endpoint:** `/v1/articles POST`
**Request Body:**

```json
{
    "title": "Awesome Python",
    "content": "A collection of awesome browser-side Python libraries, resources, and shiny things."
}
```

**Response Headers:**

- On Success: HTTP Status = `201`

### Add Comments

**Endpoint:** `/v1/articles/{id}/comments POST`
**Path Param:** *id*: The id of the article that the comment will be added to
**Request Body:**

```json
{
    "author": "Some dude",
    "content": "I like that! 😀"
}
```

**Response Headers:**

- On Success:
  - HTTP Status = `201`
- On Failure:
  - Invalid ID path parm: HTTP Status = `400`
  - Invalid comment structure: HTTP Status = `400`
  - No article exists for the ID: HTTP Status = `400`

### Get Comments For Article

**Endpoint:** `/v1/articles/{id}/comments GET`
**Path Param:** *id*: The id of the article to get the comments for
**Response Body:**

```json
[
 {
        "id": 2,
        "article_id": 1,
        "author": "John Doe",
        "content": "Lovely, thanks a lot for sharing",
        "creation_timestamp": "2024-12-11T13:38:18.628236Z"
 },
 {
        "id": 3,
        "article_id": 1,
        "author": "Some dude",
        "content": "I like that! 😀",
        "creation_timestamp": "2024-12-11T13:39:00.477245Z"
 },
 {
        "id": 1,
        "article_id": 1,
        "author": "Ahmed Ehab",
        "content": "I like the plethora of ideas, the deep trenches of nuances, and the overarching hand of beauty in this article",
        "creation_timestamp": "2024-12-11T13:37:52.031339Z"
 }
]
```

**Response Headers:**

- On Success:
  - HTTP Status = `200`
- On Failure:
  - Invalid ID path parm: HTTP Status = `400`
  - No article exists for the ID provided: HTTP Status = `404`
