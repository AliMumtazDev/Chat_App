FROM node:18.16.0-alpine
WORKDIR /app
COPY chat-app/package*.json ./
RUN npm install
COPY ./chat-app .
RUN npm run build
EXPOSE 3000

CMD ["npm", "start"]