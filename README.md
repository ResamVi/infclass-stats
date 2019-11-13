![alt](https://i.imgur.com/NwFf7Pn.png)

`infclass-stats` parses econ statements to gain statistics of the [Teeworlds InfClass mod](https://github.com/yavl/teeworlds-infclassR).  
A web-based interface will evaluate and display those insights in real-time.

# Features

## Tracks Currently online/Most active/Most kills

![players](https://i.imgur.com/YbpBttx.png)

## Counts Human/Zombie wins & Roles picked

![count](https://i.imgur.com/j4Nl0Ol.png)

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
