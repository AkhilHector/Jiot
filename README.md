### Javascript and IOT

> `NOTE 1`: This is project `JIOT` and this is in the experimental stages of development but the support would be rolled out for :
- `Chrome (33.0)`
- `Firefox (44.0)`

> `NOTE 2`: This is just a presentation mockup and brief scripts which are part of my talk/presentation for `FOSSASIA 2016`.

> `NOTE 3`: THIS IS NOT AN NPM MODULE`.

### Okay What is this really About:

IOT (Internet of Things) just simply means the combination of different physical devices be it an embedded circuit or a sensor etc. So consider it to be a culmination of different devices which are connected to the Internet.
`PS` : Google the definition and go to Wikipedia if you are really interested to know what is IOT all about.

Javascript is a fully prototypal Inheritance based programming language which offers us the ability to program both on the client side and Server side. This although will involve a lot of concepts and good practices about writing media servers with Node.Js

## Slide - 1

### Principle Behind Connected Devices

- IOT simply refers to the physical devices as in embedded circuits, physical sensors, etc.
- We establish a network connecting the physical devices to the internet.
- Examples of physical devices or the embedded circuits may be :
    + Heart Beat Sensors,
    + Gyro/Accelerometers IMU (ADXL335, MPU6050)
    + Ultrasonic Sensors

> `COMMENT` :

- So having a physical device/s being connected to the internet just doesn't ensure that it comes into the segment of IOT. Most of the developers fail to notice the effective usage of communicating between the connected devices. Most importantly the security of these devices is in its infancy, so there must be care taken for when deploying IOT based web applications.
- If the project/application is all about just taking data, then we have to ensure a front-end or a web based interface is made for the client end in order to monitor the data.
- If the project/application is all about maintaing/manipulating different physical objects and stuff like Home automation, bluetooth controlled applications etc, then you should ensure a method writing modular API endpoints for performing all the operations instead of just hard coding the values upon each analog/digital value.
- Lastly it always is about how optimized the application is, so if your code involves dealing with lot of scripts then ensure you write `cron` jobs in order to add up to the beauty of the application.

## Slide - 2

### Media Servers

- Media Servers are just a help or give support to the ability of streaming/downloading the content which is present in a SSD/HDD or flash Drives etc.
- So this means the application must involve an API endpoint which covers the :
    + GET
    + POST
    + OPTIONS
    + DELETE
- Apart from serving the content and responding to the API calls we should also give support from the application's point of view for different mime-type(media-server).

> COMMENT :

- The point of involving Javascript here is how effectively we are able to use the concept of media servers | This talk/presentation deals mostly with writing media servers and providing the scope for integrating it with IOT.
- Developing this would feature us in being involved with the job of connecting the hardware and software.
- The required hardware requirements would be :
    + Raspberry Pi
    + SD Card
    + HDD/SSD, flashdrives etc.
    + Wifi adaptor 2.4 GHz (Raspberry PI supproted)
- The required software requirements would be :
    + Raspberry Pi Supported OS [preferred - Raspbian OS]
    + Node [> 4.0 && < 4.2.6]
    + NPM [Tagged along with Node install]
    + git
- The next step would be creating a bootable SD card by creating an image on the SD card. It could be achieved using the command :
```
sudo dd bs=8M if= /to/the/dir/where/the/file/is of= /to/the/SD/card/mountpoint
```
- If git is installed you can traverse this codebase and run the script `install.sh` in order to install the required software dependencies for our purpose, else download git first using the command `sudo apt-get install git` and then follow the perform what is mentioned in the first part of this point.
- Lastly writing an media server API which has GET, POST, DELETE, OPTIONS endpoints and is talking to the physical devices is needed.
- Effectively cache and content and oversee the capability of the server to use the HTTP 304.
- Support different mime-types and fix the URI in such a way that when an unknown mime-type is being opened the media file should not automatically download.

## Slide - 3

### Technology Javascript (Node.js)

- With the help of the server side development language Node.Js we can really explore the poosibilty of creating this.
- For Node we require these core packages :
    + http
    + net
    + fs
    + varnish
    Although it is very much agreeable to the fact of writing code my making the most of the node's core modules varnish is a very good option for providing a caching mechanism.
- Apart from varnish all the above mentioned modules are core modules of the language [Node.Js]

> COMMENT :

- It is to be noted that we can have the possibility of writing a HTTP middleware which will just suffice the purpose of our application and instead of using a third party routing and HTTP middleware like `express` we can actually pre-define the set of routes that will have a GET/POST end point and accordingly we can reply with a response.
- We have the possibility of matching the entered URL by regular expression, so if an unknown route is being requested then we can prompt a 404.
- By using the core modules we can simplify the application with respect to its dependence on third party modules and since the usage of core modules can help in building cleaner and faster looking applications.

## Slide - 4

### Serving the Content

- For designing a router we can simplify the job by writing a function handler, so the serving mechanism would be a function handler that takes request and response as a parameter.
- Define the file/media which is to be addressed at a particular route and properly write the end point linked to the route.
- If routes are to be dynamically addressed write a regular expression match to the function handler in order to match the file name captured by the end point and serve it.

> COMMENT :

- Although this is more about code conventions and things to lookout for when writing/developing the application there are more insights into the specifics which are mentioned above.
- The first phase of the project is to deliver API calls which are normal, i.e a request is sent to the API and depending upon the request the file/media is being served.
- In order to write regular expression match for the request made to the API router, we have to use the module `net` which gives the application an ability to match the requests based on a regex pattern.

## Slide - 5

### The Gotchas

- Defining the headers explicitly is actually a very good notion whilst designing the media server router.
- Note the point that you should not count on node either for processing unknown mime-types or for defining headers.
- So the primary problem which might occur will be when an unknown mime-type is being opened, automatically the file is being prompted to download. Since the project is mostly dependent on the browser and it deals with media/files there is a way we can address to solve the problem of auto download.
- Have a `.csv` file which contains all the supported set of mime-types which you wish to offer and depending on the request that is made to the API strip the mime-type of each request and check whether it is supported by your router, if yes then go display else, make sure you prompt a 404 or a similar Sorry page addressing the client.


