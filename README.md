# go-saas
go framework for saas(multi-tenancy). `go-saas` targets to provide saas solution for go

# Overview

install

```
go get github.com/goxiaoy/go-saas
```

* Different database architecture
  * [x] Single-tenancy:  Each database stores data from only one tenant.
  * [x] Multi-tenancy:  Each database stores data from multiple separate tenants (with mechanisms to protect data privacy).
  * [x] Hybrid tenancy models are also available. Some tenants share one database while others own individual databases
* Domain driven design (DDD)
* Support multiple web framework
    * [x] gin
    * [ ] iris
* Support multiple orms
    * [x] [gorm](https://github.com/go-gorm/gorm)
    * [ ] [ent](https://github.com/ent/ent)
* Customizable tenant resolver
    * [x] Query String
    * [x] Form parameters
    * [x] Header
    * [x] Cookie
    * [ ] Route
    * [x] Domain format
  
* Features
    * [x] Auto set tenant value before operation
    * [x] Temporarily disable tenant filter( Only works in database `Multi-tenancy` mode)
    
    
# Sample Project

 * [gin-vue-admin](https://github.com/Goxiaoy/gin-vue-admin) This is a fork repo from original one, which uses `go-saas` to build `saas` pattern
 
# Documentation
 Refer to [wiki](https://github.com/Goxiaoy/go-saas/wiki)


#Referene

https://docs.microsoft.com/en-us/azure/azure-sql/database/saas-tenancy-app-design-patterns
