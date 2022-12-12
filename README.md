
# Creating gomobile project

Creating new go mod:
```
$ go mod init test
```

## Set GOPATH
Lets create our workspace and the same time lets use the workspace directory is our $GOPATH too.
```
$ mkdir -p /Users/burhon/Documents/testgo
$ cd /Users/burhon/Documents/testgo
$ export GOPATH=`pwd`
```

## Set Android related envs
Lets export the Android SDK and Android NDK
```
$ export ANDROID_NDK_HOME=/Users/burhon/Library/Android/sdk/ndk/19.2.5345600
$ export ANDROID_HOME=/Users/code_joker/Library/Android/sdk
```

## Downloading gobind and gomobile 
gobind is a tool that generates language bindings that make it possible to call Go functions from Java and Objective-C. It is called internally by gomobile which can help us build cross-platform applications. We need this two to build our mobile app library.
Optionally, one can develop an entire mobile application using build command of gomobile. But we wont be doing that. We will, however, develop a library that can be used by mobile applications. And to build this, we will use bind command of gomobile.
Download gobind and gomobile
```
$ go install golang.org/x/mobile/cmd/gomobile@latest
$ go get golang.org/x/mobile/cmd/gobind
```

And export the bin as part of the $PATH and init gomobile
```
$ export PATH=$PATH:'/Users/code_joker/go/bin'
$ gomobile init
```
