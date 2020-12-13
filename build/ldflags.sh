version=$(git describe --tags --dirty)
gitCommit=$(git rev-parse --short HEAD)
buildTime=$(date +%Y-%m-%d_%H:%M:%S%z)
echo "-X github.com/mt-inside/badpod/pkg/data.Version=${version} -X github.com/mt-inside/badpod/pkg/data.GitCommit=${gitCommit} -X github.com/mt-inside/badpod/pkg/data.BuildTime=${buildTime}"
