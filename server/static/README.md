# readme
Notice: I couldn't get the server to build (see ##isues encountered section). So I added a small
go server that emulates what the Java server (appears to be doing) to stay inside the 3 hour window.

In theory my Java code should work (if I can get it to build), I'd be interested to have some feedback as to how this would be solved from one of the eBay engineers :-) 

Though I wrote a really simple Go application in `../server/` that should do exactly what is asked so that the frontend builds too. Though you may be able to get the java project to run also.

## building
You will need Go, preferably the latest version:

    cd server
    go build
    ./server

The server will sit on localhost:8080.

    Open up: localhost:8080/static/

## isues encountered
I had trouble building this on my Macbook. There seems to be a few bugs with Spring and the
newer versions of Java (11)

    https://github.com/spring-projects/spring-framework/issues/22814
    https://stackoverflow.com/questions/46230413/jdk9-an-illegal-reflective-access-operation-has-occurred-org-python-core-pysys
    https://stackoverflow.com/questions/43447798/springboot-unable-to-start-embedded-container

Most of these I tried, but couldn't get to work. Setting vm flags, trying to use new dependencies, my Gradle version is the latest, etc.
