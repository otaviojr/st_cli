# st_cli
[![Donate](/docs/donation.png?raw=true)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=65XBWNBZ69ZP4&currency_code=USD&source=url)

This is a SmartThings command line client written in golang

# Motivation

When trying to make an app in python to run on a raspberry pi zero, I noticed
that it takes 8s only to start running the script. Now imagine waiting 15s to
unlock your door. Now, imagine that if you need to go to the bathroom! :-)

Nodejs do not make it much better.

But, golang can create native applications which make things much faster

# Installing

```
git clone git@github.com:otaviojr/st_cli.git .
cd st_cli
go build -o st_cli ./src/*
```

# Using it

```
./st_cli help
```

# Token

In order to allow the client to have access to your devices using Smartthings API you must
provide a **token**

You can get a Smartthings Token here:
[https://account.smartthings.com/tokens](https://account.smartthings.com/tokens)

# Using it by examples

if you want to list your devices:

```
./st_cli device list --token <<smartthing token>>
```

you can use jq (json command line parser) to get only information you need

```
./st_cli device list --token <<smartthing token>> | jq --raw-output '.items[]|.label,.deviceId'
```

you can filter by SmartThings capabilities.

```
./st_cli device list --token  <<smartthing token>> --capability switchLevel | jq --raw-output '.items[]|.label,.deviceId'
```

you can send a command to change the device state:

```
./st_cli device command --token  <<smartthing token>> --device  <<device_id>> --command setLevel --arguments 50 --capability switchLevel
```

you want to know if your door is locked?

```
./st_cli device status --token <<smartthing token>> --device <<device_id>> | jq --raw-output '.components.main.lock.lock.value'
```

you want to unlock your door?

```
./st_cli device command --token <<smartthing token>> --device <<device_id>> --command unlock --capability lock
```

list all your scenes?

```
./st_cli scenes list --token <<smartthing token>>
```

execute a scene?

```
./st_cli scenes execute --token <<smartthing token>> --scene <<scene_id>>
```

# Sending token as environment variable

if you don't want to send the token all the time, you can set the SMARTTHINGS_TOKEN environment variable.

# Donation

And... if this helps you to save time and money. Pay me a coffee. :-)

[![Donate](/docs/donation.png?raw=true)](https://www.paypal.com/cgi-bin/webscr?cmd=_donations&business=65XBWNBZ69ZP4&currency_code=USD&source=url)
