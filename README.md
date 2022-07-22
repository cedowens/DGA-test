# DGA-test
simple code to help with DGA nxdomain response testing

1. > go build
2. > ./DGA-TEST

This simple code will read from the included nxdomains.txt text file to make requests for 2002 non-existent domains, which will generate 8008 NXDOMAIN responses (dns r code 3) - from 2 A and 2 AAAA requests per domain. This can be helpful with testing any detections around NXDOMAIN responses over a threshold of time.
