# Garrett Thompson (Mr.Lockpick) - Speech IRSEC 2024
I did little to no changes to the code that where signficiant, so if you want to use this version of the store I recomend using ISTS-24s store base and changing that in case I messed anything up. 

## Found Bugs and Vulns

So, if a blueteamer goes to /admin on the store it completely just lets them use the admin pannel. Some blueteamers found out about this last year so you gotta fix it now the secrets out. 

Also the item "ADMIN ACTION" or in my case "ADMIN ACTION ONLY" is a special store item since its used when you manually change the balances of accounts. So if you change the name of this like I did, you gotta change it in some configs as well or else you get a error along the lines of "en0 item not found". 

Finally you have to delete the database file in database to reset store items but it will clear out all database information including team points, purchase history, and rebuild all the items from scratch. To fix this I would recomend having a way to clear out one table (using a simple shell script or something) so it cleans out the item data but not the team data. Right now we just nuke all which means we can only add items not remove items once the comp starts. 

##  Final notes
I am pushing my active store repo from my personal ritsec gitlab to the IRSEC 2024 gitlab. If you need to see some version history stuff that you cant see on here let me know I can sauce you up. Good luck finding my discord. 

## The Future 
This is probably the last time we are using this code, I want to do it from scratch for the hell of it so next store will look alot different. Good luck out there shoppers. 


# Ash Ketchups notes from the ISTS 2024 Store~~~

This is the store shamelessly stolen from last year's ISTS with a couple of color changes.

## Running the store

To run the store all you have to do in run `rebuild_and_redeploy_docker.sh`.

For development purposes I recommend using `sudo go run main.go`.
To change the port the store is running on, change `Router.Run(":80")` to whatever port you want. The ports in the `docker-compose.yml` file will also need to be changed if you are using that.

If you are using a different box than `store.ists.io`, you will have to change the hostname and revese proxy in `.caddy/Caddyfile`, as well as the hostname in `databases/config.yml`.

## Updating the Discord bot

There are 3 different files you need to update for the Discord bot:
```
databases/config.yml
discord/discord.go
bot/config.go
```
### Setting up the webhook

Go to the channel where you want store purchases to show up, and click on the channel settings. Go to `Integrations` and create a `New Webhook`. Name it whatever you want and make sure it is targeting the right channel. Click `Copy Webhook URL` and paste it in `databases/config.yml` next to `webhook`.

Now go to `Server Settings > Roles` and find the role you want mentioned whenever a purchase is made. Click `Copy Role ID` and go to `discord/discord.go`. Replace `1207896532305846353` with the role ID you copied.

### Setting up the bot

Go to [https://discord.com/developers/applications](https://discord.com/developers/applications) and click `New Application`. Name it whatever you want.

Copy the `APPLICATION ID` under `General Information`. Replace the `AppID` in `bot/config.go` with the application ID you copied.

Now go to `Bot` and click `Reset Token`. `Copy` the token and paste it after `bot` in `databases/config.yml`.

Now go to `OAuth2` and under `Default Authorization Link` select `In-app Authorization`. Under `OAuth2 URL Generator > SCOPES` select bot, then under `OAuth2 URL Generator > BOT PERMISSIONS` select `Send Messages`. Then click `Copy` under `GENERATED URL`. You can have the CA add the bot using that link, or you can add it if you have Admin permissions in the server.

Last things last, left click on the server and click `Copy Server ID`. Replace `GuildID` in `bot/config.go` with the server ID you copied.

The bot should now be good to go! You may have to restart the docker container if you are getting application errors when sending slash commands.

## Configuration

Pretty much all of your other configuration will be done in `databases/config.yml`.

There you can change the store items, users, purple team tokens, KOTH tokens, CTF flags, and sponsor tokens (if applicable).

Once all your changes have been made to the `databases/config.yml` file, you might want to delete the database at `databases/data.sqlite` so you get a fresh new database to work with.
# wicys_ctf_store
# wicys_ctf_store
# wicys_ctf_store
