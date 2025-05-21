FROM node:22-alpine

WORKDIR /app

COPY node/package*.json ./

RUN npm install

COPY node/src/ ./src/

COPY proto/ ./proto/

EXPOSE ${NODE_PORT}

CMD ["npm", "start"]