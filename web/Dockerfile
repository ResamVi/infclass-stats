FROM node:16.1.0 as build-stage
WORKDIR /app
COPY package*.json ./
ENV VUE_APP_API_URL="wss://inf.resamvi.io"
RUN npm install
COPY . .
RUN npm run build

# production stage
FROM nginx:stable-alpine as production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
