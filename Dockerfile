FROM node:18 as assets

WORKDIR /app
COPY . .
WORKDIR /app/assets

RUN npm install

WORKDIR /app

RUN npx tailwindcss -c ./assets/tailwind.config.js -i ./assets/src/css/app.css -o ./assets/dist/css/app.css

# TODO: actual asset build step (optimize images, minify js, etc...)
RUN rm -rf ./assets/dist/images ./assets/dist/js 2> /dev/null || true
RUN cp -rf ./assets/src/images ./assets/dist/images
RUN cp -rf ./assets/src/js ./assets/dist/js

FROM golang:1.21

WORKDIR /app
EXPOSE 8000

RUN go install github.com/a-h/templ/cmd/templ@v0.2.408

COPY . .
COPY --from=assets /app/assets /app/assets

RUN templ generate
RUN go build -o /app/main /app/main.go

RUN chmod +x /app/main

CMD ["/app/main"]
