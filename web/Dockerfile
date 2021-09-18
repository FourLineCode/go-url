FROM node:16-alpine AS development

WORKDIR /web

RUN apk --no-cache add curl

RUN curl -f https://get.pnpm.io/v6.14.js | node - add --global pnpm
RUN pnpm add -g pnpm

COPY . .

EXPOSE 3000

CMD [ "pnpm", "dev" ]