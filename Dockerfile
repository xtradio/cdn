FROM nginx:1.15.4-alpine

COPY nginx.conf /etc/nginx/conf.d/cdn.conf

