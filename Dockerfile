FROM scratch
ADD main /
CMD ["/maind"]

#script used to build statically linked bin
#CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#Using a scratch container allows the images to be much much smaller
#got the idea from watching a talk on kubernties by Kelsey Hightower
#super funny and informative
#https://www.youtube.com/watch?v=XPC-hFL-4lU