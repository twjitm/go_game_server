cd ~/go/src/go_game_server/proto
rm -rf *.go
protoc -I  ./ *.proto --go_out=plugins=grpc:./
easyjson -all *.go