FROM node:15-alpine
WORKDIR /app

# Install build dependencies
#RUN apk --no-cache --virtual build-dependencies add \
#    python \
#    make \
#    g++ \
#    bash

# Install dependencies
COPY package*.json ./
RUN yarn install --silent
RUN yarn global add react-scripts --silent
RUN yarn global add serve --silent

# Remove build deps
#RUN apk del build-dependencies

# Copy source into container
COPY . .

# Build
RUN yarn run build

# Expose port 80
EXPOSE 80

# Run
CMD ["serve", "-s", "build", "-l", "80"]
