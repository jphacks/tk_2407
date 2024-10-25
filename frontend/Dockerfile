FROM shion1305/pnpm:22-alpine3.20 AS builder
COPY pnpm-lock.yaml /root
WORKDIR /root
RUN pnpm fetch --prod
COPY . /root
RUN pnpm run build
CMD [ "pnpm", "start" ]
EXPOSE 3000
