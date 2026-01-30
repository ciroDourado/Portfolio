FROM node:20-alpine

WORKDIR /app

ARG uid=1000
ARG user=node

RUN apk add --no-cache libc6-compat

COPY package*.json ./

RUN chown -R $uid:$uid /app

USER $uid

RUN npm install

COPY --chown=$uid:$uid . .

EXPOSE 3000

CMD ["npm", "run", "dev"]
