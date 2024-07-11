Steps to reproduce the bug

Install make and protobuf-compiler
```bash
sudo apt update
sudo install make
sudo install protobuf-compiler=3.21.12-8.2build1
````

Check installed version via
`protoc --version`
I've reproduced the bug with `libprotoc 3.21.12` on `Ubuntu 24.04` running `WSL2` at Windows 11

The example compiles a protobuf and runs a non-TLS grpc server on port `8080`
If you want to change the port you can override it via env `export GRPC_PORT=<your-port>`

You can build and run the server app via make command
```bash
export GRPC_POST=8080
make
```

I'm using Postman v11.3.1 to check the issue
<img src="assets/postman.png">
