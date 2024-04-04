## asdgo

> [!CAUTION]
> This repository is heavily under development and is not ready for production use. It's untested, it uses global vars, it's not documented and I'm new Go. I'm currently using it to build my projects and sometimes I decide to rewrite stuff, thus breaking things. I use it for almost all my personal Go web projects. I will update/remove this message when it's ready for public use and is going into a stable release.

Asdgo is an opinionated and simple Go web "framework", or rather a set of components which glues together some packages and logic that I don't want to write everytime I start a new project.

This project uses the following amazing packages (please go check them out)
- [Echo framework](https://echo.labstack.com/docs) (for routing)
- [GORM](https://gorm.io/) (for database)
- [Templ](https://templ.guide/) (for templating)

### Components

Every component is prefixed with the letter `a`, so it doesn't conflict with the naming of other packages (or internal packages). So `template` becomes `atemplate`, `mail` becomes `amail`, etc, etc.

#### `acontext` 
(For retrieving things from the context)

#### `adatabase` 
(For default user model)

#### `ahash` 
(For hash)

#### `amail` 
(For sending emails)

#### `amiddleware`

#### `aqueue` 
(A simple queue)

#### `aschedule` 
(A simple schedule)

#### `asession`

#### `atemplate`

#### `avalidate`

### Example

```go
asd := asdgo.New(asdgo.Config{})

asd.GET("/", func (c echo.Context) error {
    return c.SendString(http.StatusOK, "Hello, world!")
})
```

### Feedback

If you have any ideas on how to improve things, please open a issue or pull request.
