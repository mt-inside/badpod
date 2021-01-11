version=$(git describe --tags --dirty)
buildTime=$(date +"%F %T%Z")
echo -X "'"github.com/mt-inside/badpod/pkg/data.Version=${version}"'" -X "'"github.com/mt-inside/badpod/pkg/data.BuildTime=${buildTime}"'"
