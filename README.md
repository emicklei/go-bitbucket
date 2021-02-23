## go-bitbucket

client package for the bitbucket.org [api 2.0](https://confluence.atlassian.com/display/BITBUCKET/Version+2)

MIT License, (c) 2015, http://ernestmicklei.com

### example

```
  func main() {
    c := bitbucket.New("your-user-name", "its-password")
    err := c.RepositoriesDo(func(each bitbucket.Repository) error {
      log.Println(each.Name)
      return nil
    })
    log.Println(err)
  }
```
