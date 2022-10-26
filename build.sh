echo "BUILDING.."
go build ./
mv ./main ./build/main
cp -r ./res/* ./build/res
echo "BUILT"