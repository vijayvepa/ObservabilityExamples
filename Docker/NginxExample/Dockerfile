FROM nginx:1.27.2
RUN apt-get update 
COPY index.html /usr/share/nginx/html/
CMD ["nginx","-g", "daemon off;"]
EXPOSE 80