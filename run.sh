docker build . -t nutriguide-backend-image:test && \
docker run --rm --name nutriguide-backend -p 8080:8080 nutriguide-backend-image:test
