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

Fill the environment variables with your information:

`autoexec.cfg`
```
ec_bindaddr <myip>
ec_port <myport>
ec_password <mypassword>
ec_output_level 2
```

`docker-compose.yml` should match the above
```
- SERVER_IP=<myip>
- ECON_PORT=<myport>
- ECON_PASSWORD=<mypassword>
```

`web/Dockerfile`
```
ENV VUE_APP_API_URL="wss://inf.resamvi.io:8001"
```

I'm using [Plausible](https://plausible.io/).  
If you host this yourself you may want to stop the app from sending user stats:

`web/public/index.html`
```diff
    ...
    <title>InfClass Statistics</title>

-   <script async defer data-domain="stats.resamvi.io" src="https://pls.resamvi.io/js/pls.js"></script>
</head>
```


# How to make sure your Infclass mod is compatible with infclass-stats

TODO: Copy in what message logs are required 'Protocol'