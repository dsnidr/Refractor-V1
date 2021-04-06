# Refractor

![Latest Version Tag](https://img.shields.io/github/v/tag/sniddunc/refractor?label=version&style=flat-square)
![License](https://img.shields.io/github/license/sniddunc/refractor?style=flat-square)
![Lines of Code](https://img.shields.io/tokei/lines/github/sniddunc/refractor?label=lines%20of%20code&style=flat-square)
![Contributors](https://img.shields.io/github/contributors/sniddunc/refractor?color=%2397CA00&style=flat-square)

A game server community management panel written in Go and React.

## Features

- Easy installation with Docker
- Support for Mordhau and Minecraft
- Real time server player list
- Player infraction logging (warnings, mutes, kicks and bans)
- Player summary lookup
- Player search
- Infraction search
- Real-time chat
- Support for multiple servers and users
- User access control
- and more!

## Motivation

Game server moderation is a hard task. Most moderators are volunteers who contribute their time to keep their respective communities in good shape to provide fun for their players.

Moderators have a tricky task because it often involves keeping an eye on — in many cases — dozens of players while still trying to enjoy the game themselves.

Many moderators work as part of a larger team where communication is key in order to keep rulings consistent between different moderators. Often times, this is done in a group chat or a rudimentary logging system. This can work, but it's not perfect as it becomes difficult for moderators to quickly find a player's history and other key information in the very fast paced environment of online game servers.

With Refractor, all this information is centralized and easy to find with the help of a modern user interface. With Refractor, quickly finding relevant data for player interactions is easy so that moderators can moderate more effectively while still enjoying the game.

## Supported Games

The following games are currently supported:

- Mordhau
- Minecraft

# Installing with Docker

Docker is the recommended installation method. It is by far the easiest method and it takes care of TLS and API proxying for you.

Installation on a fresh machine is recommended, as there will likely be conflicts when installing on a machine which already hosts a web server.

If you decide to install manually, **please make sure you encrypt your traffic properly**. It not only protects you, but also your users and players. It is very important that you always encrypt your traffic.

## 1. Requirements

You require a valid domain name which points to the machine you're installing Refractor on.

You also need to have Docker and Docker Compose installed. You can find installation steps for your specific system at the following two links:

[Install Docker Engine](https://docs.docker.com/engine/install/)

[Install Docker Compose](https://docs.docker.com/compose/install/)

Additioanlly, Refractor needs access to a MySQL database. Depending on who hosts your game server, you might already have one at your disposal. Check your game server's control panel.

If you have a MySQL database you can use, skip to step 2.

### Creating a MySQL Docker container

If you don't yet have a MySQL database, you can install one on the same machine as Refractor. The recommended way to do this is with Docker. With Docker already installed, you can run the following command to set up a MySQL server container:

```zsh
docker run --name mysql --restart unless-stopped --network host \
  -e MYSQL_ROOT_HOST='0.0.0.0'  -e MYSQL_ROOT_PASSWORD='strongpassword' \
  -d mysql/mysql-server:latest
```

**Make sure you replace `strongpassword` with a strong password.** This is the password for your database's root user, so it's very important that it's secure.

Once this command is done executing, run the following command:

```zsh
sudo docker ps
```

You should see your new MySQL container listed there, and it's status should say that it's up.

Now you need to create your database within your new MySQL installation. To open a session with your MySQL server, use the following command:

```zsh
docker exec -it mysql mysql -u root -p
```

The first `mysql` is the name of your Docker container. The second is the mysql command available in the container. `-u` is where you specify your username and `-p` tells the server that you'll provide a password.

It should ask you for your password. Once you enter it, you should be presented with a MySQL prompt and see something like: `mysql>`.

You can now create a database by typing the following command:

```sql
CREATE DATABASE refractor;
```

This will create a new database called refractor.

### Creating an additional user

For security reasons, it's best not to run applications right from the root user. Instead, you should create an additional user. You can do this using the following command from inside the MySQL prompt:

```sql
CREATE USER refractor IDENTIFIED BY 'password';
```

`refractor` is the name of the new user and `password` is the new user's password.

Next, you need to give this user permissions on the database you created previously. You can do this with the following command:

```sql
GRANT ALL PRIVILEGES ON refractor.* TO refractor;
```

The first `refractor` is the name of the database which you're assigning the new user's privileges on. The second `refractor` is the name of your new user. This will grant the user all permissions.

Now, just flush the permissions to make them take effect.

```sql
FLUSH PRIVILEGES;
```

## 2. Cloning the Repository

Once you have Docker installed, it's time to pull the code from this repository. Use git to clone this repository to a location of your choice.

```zsh
git clone https://github.com/Sniddunc/Refractor.git
cd Refractor
```

## 3. Running the setup script

> If you are a on a system with low memory (1gb or less) you may want to create a swapfile. See Troubleshooting below under section "**The installation hangs, freezes up my server or just takes forever**"

Inside the directory you cloned, you should see a file called `setup.sh`.

This file should already be executable, but if it isn't then you can run the following command to make it executable.

```zsh
chmod +x setup.sh
```

Then, you can simply run the `setup.sh` script.

```zsh
./setup.sh
```

> It's a good idea to do a test run before generating a proper SSL certificate. To do this, edit `setup.sh` and change the value of the `staging` variable near the top of the file to 1 before running the script. Once you've run through the install and are confident that everything is working properly, change the value of staging back to 0. This is entirely optional, but highly recommended.

Follow the instructions in the setup script. Provide your domain name, email address and then enter your MySQL database's credentials.

Once this is done, the docker images for the various components of Refractor will be built. You will likely see lots of text appearing on your screen as everything gets installed. This is perfectly normal.

This installation could take a good amount of time to complete, so be patient!

Once the installation is complete, navigate to your domain and you should be presented with the Refractor login screen. If you enter the credentials you provided for the initial user, you should be able to log in.

You're all set. Enjoy Refractor!

## Troubleshooting

### **The installation hangs, freezes up my server or just takes forever**

This issue is often caused by a lack of system memory during the build process of the React client. If you are on a low memory system (1gb or less) then you may consider creating a swapfile of 1gb or more to help along the build process. Research how to create a swapfile for your Linux distribution for more information.

### **Too many certificates already issued for exact set of domains**

Let's Encrypt (the service used to handle our certificates) has rate limiting. If you get this error, you have requested too many certificates in their rate limiting timeframe. More info can be found (here)[https://letsencrypt.org/docs/rate-limits/].

### **Challenge failed / Some challenges have failed**

This is likely because the domain record you're using has not yet been fully propagated. Domain changes can take up to 48 hours to propagate fully. You should re-try the installation in at least a few hours, if not longer.

> If you ran into and resolved any issues during the installation which you think others may run into, feel free to add a section under troubleshooting and submit a PR.
