FROM migrate/migrate

WORKDIR /app

RUN apk add --no-cache bash

COPY ./wait-for-it.sh ./wait-for-it.sh
COPY ./migrator/entrypoint.sh ./entrypoint.sh

RUN chmod +x ./wait-for-it.sh
RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
