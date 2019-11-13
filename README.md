![alt](https://raw.githubusercontent.com/ResamVi/infclass-stats/master/web/src/assets/bg.png)

`infclass-stats` parses econ statements to gain insights of player performances of the [Teeworlds InfClass mod](https://github.com/yavl/teeworlds-infclassR).  
A web-based interface will evaluate and display those insights in real-time.

Visit [https://stats.resamvi.io](https://stats.resamvi.io) to see yourself.

# Features

## Evaluate the Week's best Players 

![best players](https://user-images.githubusercontent.com/6261556/68749753-095b5000-05ff-11ea-9241-48daef0b3dc9.PNG)


## Track current Activity

![activity](https://user-images.githubusercontent.com/6261556/68749901-458eb080-05ff-11ea-807f-8924013400c0.PNG)

![Unbenannt](https://user-images.githubusercontent.com/6261556/68750012-7c64c680-05ff-11ea-9a29-bbb4c50a2808.PNG)


## Count Kills and Survivals

![count](https://user-images.githubusercontent.com/6261556/68749970-68b96000-05ff-11ea-95d3-6045f1aa58c7.PNG)

## Compare Classes

![classes](https://user-images.githubusercontent.com/6261556/68750201-c8b00680-05ff-11ea-91fa-a84f349f6172.PNG)
![alive](https://user-images.githubusercontent.com/6261556/68750336-fe54ef80-05ff-11ea-8f73-5e6232817f3d.PNG)


## Display best performing Players on their Classes

![score](https://user-images.githubusercontent.com/6261556/68750444-28a6ad00-0600-11ea-85ab-5686396ec8d5.PNG)
![top5](https://user-images.githubusercontent.com/6261556/68750481-38be8c80-0600-11ea-90c4-81b7584a7a80.PNG)

## Show Map Statistics

![maps](https://user-images.githubusercontent.com/6261556/68750567-5e4b9600-0600-11ea-88b5-693e988caa84.PNG)


# Setup

infclass-stats runs on Port `8000` so make sure it is free.

Fill the <..> with your information:

`autoexec.cfg` of the teeworlds server
```
ec_bindaddr <ip>
ec_port <port>
ec_password <password>
ec_output_level 2
```

Set up your `config/config.go` of infclass-stats
```
SERVER_IP = "<ip>"

ECON_PORT = "<port>"

ECON_PASSWORD = "<password>"

...

MYSQL_USER = "<mysql user>"

MYSQL_PASSWORD = "<mysql password>"
```

Setup a database named "infclass"

Replace the IP in `web/src/App.vue`
```
const ws = new WebSocket('ws://<ip>:8000/');
```

Then run

```
go run main.go
```

or

```
./start.sh
```
which does the above repeatedly (helps recovering, when server crashes)


# Config

# SSL

This implementation uses SSL. If you want to skip (easier testing) this step then:

Change `wss` to `ws` in `App.vue`

```
const ws = new WebSocket('ws://localhost:8000/subscribe');
```

If you plan to use SSL create a symlink to the project root

```
ln -s /etc/letsencrypt/live/<url>/fullchain.pem ~/infclass-stats/fullchain.pem
ln -s /etc/letsencrypt/live/<url>/privkey.pem ~/infclass-stats/privkey.pem
```

# How to make sure your Infclass mod is compatible with infclass-stats

TODO

Copy in what message logs are required 'Protocol'

# TODO

- Remove logs and fmts where possible

# Website

1. Install dependencies

```
yarn install
```

2. Build artifacts`

```
yarn run build
```

3. Move `dist` to `/var/www/` or appropriate place
