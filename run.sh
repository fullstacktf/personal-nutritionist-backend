docker build . -t nutriguide-backend-image:test && \
docker run --rm --name nutriguide-backend -p 8080:80 nutriguide-backend-image:test
