# Example project

combination of `go-saas`,`gin`,`gorm(sqlite)`

```shell
go run github.com/goxiaoy/go-saas/examples
```
---
Host side ( use shared database):

Open `http://localhost:8888/posts`

---
Multi-tenancy ( use shared database):

Open `http://localhost:8888/posts?__tenant=1`

Open `http://localhost:8888/posts?__tenant=2`

---
Single-tenancy ( use separate database):

Open `http://localhost:8888/posts?__tenant=3`