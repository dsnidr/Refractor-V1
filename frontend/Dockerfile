FROM node:15-alpine
WORKDIR /app

# Install git
RUN apk --no-cache add git

# Copy git info into container
COPY ./.git ./

# Install dependencies
COPY ./frontend/package*.json ./
RUN yarn install --silent
RUN yarn global add react-scripts --silent
RUN yarn global add serve --silent

# Remove build deps
#RUN apk del build-dependencies

# Copy source into container
COPY ./frontend/. .

# Build
RUN yarn run build

# Expose port 80
EXPOSE 80

# Run
CMD ["serve", "-s", "build", "-l", "80"]
