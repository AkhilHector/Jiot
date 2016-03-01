### Javascript and IOT

> NOTE 1: This is project `JIOT` and this is in the experimental stages of development but the support would be rolled out for :
- Chrome (33.0)
- Firefox (44.0)

> NOTE 2: This is just a presentation mockup and brief scripts which are part of my talk/presentation for FOSSASIA 2016.

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

