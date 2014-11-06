docker run -h `hostname` -p 8180:8180  -e HOST_IP="$(curl icanhazip.com)" -e SOURCE_PATH="/gopath/src/app/" app
