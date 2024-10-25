FROM shion1305/pnpm:22-alpine3.20 AS builder
WORKDIR /root
COPY . .
RUN pnpm install
COPY . /root
RUN pnpm run build
CMD [ "pnpm", "start" ]
EXPOSE 3000
