#!/bin/bash

# Use tput to determine the right sequences for bold/colored text
bold=$(tput bold)
reset=$(tput sgr0)

# Make sure docker-compose is installed
if ! [ -x "$(command -v docker-compose)" ]; then
  echo "Error: docker-compose is not installed">&2
  exit 1
fi

echo ""
# Check if the files modified by this script were modified perviously
if ! [ $(cmp -s ./default/docker-compose.yaml ./docker-compose.yaml) ] || \
   ! [ $(cmp -s ./default/nginx.conf ./nginx/nginx.conf) ]; then
  echo "${bold}Your config files were previously modified. If you continue running this script, your changes will be lost.${reset}"
  read -p "Would you like to continue? (y/N) " decision
  if [ "$decision" != "Y" ] && [ "$decision" != "y" ]; then
    exit
  fi

  # Overwrite changes
  rm ./nginx/nginx.conf
  cp ./default/nginx.conf ./nginx/nginx.conf

  rm ./docker-compose.yaml
  cp ./default/docker-compose.yaml ./docker-compose.yaml
  echo ""
fi

echo "${bold}DOMAIN SETUP${reset}"
#domain="USER INPUT"
read -p "Enter your domain: " domain

#email="USER INPUT"
read -p "Enter your email: " email

echo ""
echo "${bold}DATABASE SETUP${reset}"
#db_uri="USER INPUT"
echo "Refractor currently supports MySQL, meaning that you need a SQL database already set up before you can continue with the installation of Refractor."
echo "Tips:"
echo "  1. If you do not already have a MySQL database, you can easily install one with Docker! Google \"Docker MySQL installation\" for guidance."
echo ""
echo "  2. If your MySQL server is hosted in a Docker container, I suggest running it in ${bold}host${reset} networked mode."
echo "     While running in host mode, you use the address of the ${bold}docker0${reset} network adapter as the DB host."
echo ""

read -p "${bold}Do you have a MySQL database and it's connection details ready? (y\N)${reset} " decision
if [ "$decision" != "Y" ] && [ "$decision" != "y" ]; then
  echo ""
  echo "Please run this installer again once you have your connection details ready!"
  exit
fi

echo ""
read -p "> DB Host: " db_host
read -p "> DB Port: " db_port
read -p "> DB Name: " db_name
read -p "> DB User: " db_user
read -p "> DB User Password: " db_password

addr_part="@"

if [ "$db_host" != "" ]; then
  addr_part+="tcp($db_host"
fi

if [ "$addr_part" != "@" ] && [ "$db_port" != "" ]; then
  addr_part+=":$db_port)"
fi

db_uri="${db_user}:${db_password}${addr_part}\/${db_name}\?charset=utf8mb4\&collation=utf8mb4_unicode_ci"

echo "DB URI: $db_uri"

echo ""
echo ""
echo "${bold}INITIAL USER SETUP${reset}"
#initial_username="USER INPUT"
read -p "> Username: " initial_username
read -p "> Email: " initial_email
echo -n "> Password: " ; read -s initial_password ; echo ""
echo ""

# Generate a random 32 byte string for the JWT secret
jwt_secret=$(tr -dc A-Za-z0-9 </dev/urandom | head -c 32 ; echo "")

# Write variables out to file placeholders
sed -ri "s/\{\{DOMAIN\}\}/${domain}/"                     ./nginx/nginx.conf ./docker-compose.yaml
sed -ri "s/\{\{EMAIL\}\}/${email}/"                       ./docker-compose.yaml
sed -ri "s/\{\{INITIAL_USERNAME\}\}/${initial_username}/" ./docker-compose.yaml
sed -ri "s/\{\{INITIAL_EMAIL\}\}/${initial_email}/"       ./docker-compose.yaml
sed -ri "s/\{\{INITIAL_PASSWORD\}\}/${initial_password}/" ./docker-compose.yaml
sed -ri "s/\{\{JWT_SECRET\}\}/${jwt_secret}/"             ./docker-compose.yaml
sed -ri "s/\{\{DB_URI\}\}/${db_uri}/"                     ./docker-compose.yaml

# Write domain out to .env file in the frontend
echo "REACT_APP_DOMAIN=$domain" >> ./frontend/.env

rsa_key_size=4096
data_path="./data/certbot"
staging=0 # set to 1 if you're testing your setup to avoid rate limits

# Check for existing data
if [ -d "$data_path" ]; then
  echo "${bold}Existing certificate data found for $domain. If you continue, this data will be overwritten.${reset}"
  read -p "Would you like to continue? (y/N) " decision
  if [ "$decision" != "Y" ] && [ "$decision" != "y" ]; then
    exit
  fi
fi

# Get recommended TLS params
if [ ! -e "$data_path/conf/options-ssl-nginx.conf" ] || [ ! -e "$data_path/conf/ssl-dhparams.pem" ]; then
  echo "Fetching recommended TLS parameters..."
  mkdir -p "$data_path/conf"
  curl -s https://raw.githubusercontent.com/certbot/certbot/master/certbot-nginx/certbot_nginx/_internal/tls_configs/options-ssl-nginx.conf > "$data_path/conf/options-ssl-nginx.conf"
  curl -s https://raw.githubusercontent.com/certbot/certbot/master/certbot/certbot/ssl-dhparams.pem > "$data_path/conf/ssl-dhparams.pem"
  echo ""
fi

# Create dummy certificate so that nginx can start up.
# If we didn't do this, nginx would not be able to start properly with our config which would make it so we cant generate a certificate.
# Classic chicken or the egg first situation.
echo "Generating a dummy certificate for $domain..."
path="/etc/letsencrypt/live/$domain"
mkdir -p "$data_path/conf/live/$domain"
docker-compose run --rm --entrypoint "\
  openssl req -x509 -nodes -newkey rsa:$rsa_key_size -days 1\
    -keyout '$path/privkey.pem' \
    -out '$path/fullchain.pem' \
    -subj '/CN=localhost'" certbot

# Start nginx
echo "Starting nginx..."
docker-compose up --force-recreate -d proxy

# Now that nginx started, we delete the dummy certificate
echo "Deleting the dummy certificate for $domain..."
docker-compose run --rm --entrypoint "\
  rm -Rf /etc/letsencrypt/live/$domain && \
  rm -Rf /etc/letsencrypt/archive/$domain && \
  rm -Rf /etc/letsencrypt/renewal/$domain.conf" certbot
echo

# Enable staging mode if necessary
if [ $staging != "0" ]; then staging_arg="--staging"; fi

# Request valid certificate from Let's Encrypt

docker-compose run --rm --entrypoint "\
  certbot certonly --webroot -w /var/www/certbot \
    $staging_arg \
    --email $email \
    -d $domain \
    --rsa-key-size $rsa_key_size \
    --agree-tos \
    --force-renewal" certbot
echo

# Reload nginx to use the new certificate
echo "Reloading nginx..."
docker-compose exec proxy nginx -s reload

echo "Done!"


echo "If you received a message from nginx mentioning an unknown or disconnected host ${bold}refractor-backend${reset}, something prevented the backend from starting."
echo "You should run ${bold}docker logs refractor-backend${reset} to see what the issue is."
echo "Once the issue was fixed and the backend is functioning normally, you must restart the proxy using ${bold}docker restart proxy${reset}"
echo ""
echo "Enjoy Refractor!"
echo ""
