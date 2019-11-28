prep:
	go install -v github.com/rakyll/statik

get: prep
	git clone https://github.com/looterz/reaper.git || :
	cd reaper && rm -rf .git

build: get
	go generate -v
	go build -v -ldflags '-s'
	upx -9 reaperd

clean:
	go clean -v
	rm -rf reaper reaperd statik
