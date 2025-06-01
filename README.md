Go organises code into units called modules, and each module needs to be in a separate folder (you can’t have modules within modules). Each module needs a `go.mod` file
which tells Go that this folder contains a module

Creating a go module is straightforward and should be the first step by executing the command

`go init <module_name>`

Note that this will create a `go.mod` in your current folder. Since every module needs to be in its own folder
