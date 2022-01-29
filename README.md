# go-cp
## description
I want to use unix cp command with windows.  
Just in time, I made my own to study Go os and io pacakge as well.  
Now, I implement only function that copy over than one file to other dir.  
When I want to use other cp's function, I implement it.  
## usage
### install
you need to Go1.16+(because I don't use ioutil package)
``` go install github.com/kimuson13/go-cp@latest ```
### command 
```go-cp [want to copy file's paths] [want to paste dir]```
## Future Prospect
- [ ] paste file with difference name than before
- [ ] some flags such as overwrite, create backup and so on...
