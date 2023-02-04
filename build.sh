cd main;
env GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo main.go;
cd ../;
mkdir -p build/bin;
mv main/main build/bin/voting-app;
echo "Built successfully"