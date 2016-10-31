
all:
	GOOS=linux GOARCH=386 go build -o gonarr-linux-386
	GOOS=linux GOARCH=amd64 go build -o gonarr-linux-amd64
	GOOS=windows GOARCH=386 go build -o gonarr-windows-386.exe
	GOOS=windows GOARCH=amd64 go build -o gonarr-windows-amd64.exe

upload:
	aws s3 cp gonarr-linux-amd64       s3://kgi-gonarr/ --acl public-read
	aws s3 cp gonarr-linux-386         s3://kgi-gonarr/ --acl public-read
	aws s3 cp gonarr-windows-amd64.exe s3://kgi-gonarr/ --acl public-read
	aws s3 cp gonarr-windows-386.exe   s3://kgi-gonarr/ --acl public-read
