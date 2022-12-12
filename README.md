
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
Gobind is a tool that generates language bindings that make it possible to call Go functions from Java and Objective-C. It is called internally by gomobile which can help us build cross-platform applications. We need this two to build our mobile app library.

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

## The sample App
Lets create the go project structure for our sample app. Let's sc directory which contains our Go package not.expert and finally our package test
```
$ mkdir -p $GOPATH/src
$ mkdir -p $GOPATH/src/not.expert
$ mkdir -p $GOPATH/src/not.expert/test
```

Create the test.go file under test package
```
$ touch $GOPATH/src/not.expert/helloworld/test.go
```

Put this as contents of the test.go
```ruby
package 'test'
import (
"fmt"
func Greet (name string) string {
message := fmt.Sprintf("Hello World, &s", name) 
return message
}
```

## Build the library
Using gomobile, we can then build the libraries respective to the traget platform, in our case ios and android.

Build framework for ios
```
$ gomobile bind -target=ios -o $GOPATH/build/test.xcframework -v $GOPATH/src/not.expert/test
```

And build aar for android
```
$ gomobile bind -target=android -o $GOPATH/build/test.aar -v $GOPATH/src/not.expert/test
```

## Create app in Xcode(ios)
Lets create native ios app in Xcode

Open Xcode and create a new Xcode app

![Create a new Xcode project]([./assets/Screenshot 2022-12-12 at 14.10.01.png](https://github.com/burhon97/gomobile-test/blob/master/assets/Screenshot%202022-12-12%20at%2014.10.01.png))

Click **Next** button 

Then: 
+ Product name: **testgo**
+ organization identifier: **not.expert**
+ Interface: **Storyboard**

Click **Next**, your project is created.

### Importing a Gomobile framework
Now we are going to import our **test.xcframework** so that we can interact with it in Swift. Open Finder, and navigate to the frameworks folder. Then drag the **test.xcframework** file into Xcode at the bottom of the Project navigator. A dialog will open to confirm the action, be sure to check Copy items if needed. And you can follow the [link](https://denbeke.be/blog/programming/go-mobile-example-running-caddy-ios/#:~:text=Building%20an%20iOS%20app%20with%20the%20framework)




