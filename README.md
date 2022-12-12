
# Creating gomobile project

Creating new go mod:
```
$ go mod init test
```

##Set GOPATH
Lets create our workspace and the same time lets use the workspace directory is our $GOPATH too.
```
$ mkdir -p /Users/burhon/Documents/testgo
$ cd /Users/burhon/Documents/testgo
$ export GOPATH=`pwd`
```

##Set Android related envs
Lets export the Android SDK and Android NDK
```
$ export ANDROID_NDK_HOME=/Users/burhon/Library/Android/sdk/ndk/19.2.5345600
$ export ANDROID_HOME=/Users/code_joker/Library/Android/sdk
```
