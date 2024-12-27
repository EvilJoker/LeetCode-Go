current_path=$(pwd)

cd ../

ln -s $current_path tmp

cd tmp

go test -v ./...

cd ../

rm -rf tmp