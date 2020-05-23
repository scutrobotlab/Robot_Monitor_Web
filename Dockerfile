FROM node:lts as build-env
WORKDIR /app
COPY ./frontend/package.json ./frontend/package-lock.json /app/
RUN npm ci

FROM build-env as build
COPY ./frontend /app
RUN npm run build

FROM golang:latest
WORKDIR /build
COPY ./backend ./main.go ./go.sum ./go.mod /build/
RUN mkdir -p frontend/dist \
    && go build -ldflags="-w -s" -o Robot_Monitor_Web_linux
COPY --from=build /app/dist/ /build/frontend/dist

CMD ./Robot_Monitor_Web_linux $PORT
