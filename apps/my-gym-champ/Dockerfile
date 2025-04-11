FROM node:20-alpine as build

WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy the rest of the application
COPY . .

# Build the application
RUN npx nx build my-gym-champ --configuration=production

# Stage 2: Serve the app with Nginx
FROM nginx:alpine

# Copy the build output
COPY --from=build /app/dist/apps/my-gym-champ /usr/share/nginx/html

# Copy nginx configuration
COPY ./apps/my-gym-champ/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"] 