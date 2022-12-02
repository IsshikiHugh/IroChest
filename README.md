# IroChest

## Introduction

**IroChest** is a lightweight quick note tool that can store things with different categories and tags.

Here provides a server and a command line client. You can also connect it to other clients with the api provided.



## Server Deploy

### Direct

Just make sure you have go environment. And my dev environment is `go1.19.2 darwin/arm64`.

After that, go to the server directory and create `config.yaml` (or copy from `config/config_sample.yaml`), and modified it.

```yaml
access: "Original Access Key" 	# This will allow you to keep your notes secrect.
port: 5252						# Allocate the port.
```

After that, run the following instructions to run the server. 

```shell
go mod tidy
go build
./IroChest
```

You can use several ways to realize persistent deployment, such as `nohup`, `screen`, `tmux` and so on.

### Docker

> Unsupported now!

## Client Start

> 咕咕咕咕……

## TODO

- [ ] The simplist client.
- [ ] Implement the authority management mechanism.
- [ ] Write API docs.
- [ ] Support docker.
- [ ] Better client.
- [ ] Maybe website.
- [ ] ...
