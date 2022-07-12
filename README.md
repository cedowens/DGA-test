# DGA-test
simple code to help with DGA nxdomain response testing

1. > go build
2. > ./DGA-TEST

This simple code will read from the included dg-test-domains text file to make requests for non-existent domains, which will generate NXDOMAIN responses (dns r code 3). This can be helpful with testing any detections around NXDOMAIN responses over a threshold of time.
