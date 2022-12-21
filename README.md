
# Creating gomobile project

### Downloading gomobile and init
```
$ go install golang.org/x/mobile/cmd/gomobile@latest
$ export PATH=$PATH:'/Users/code_joker/go/bin'
$ gomobile init
```

### Creating new go mod:
```
$ go mod init test
```

## Set GOPATH
Lets create our workspace and the same time lets use the workspace directory is our $GOPATH too.
```
$ mkdir -p /Users/code_joker/Documents/testgo
$ cd /Users/code_joker/Documents/testgo
$ export GOPATH=`pwd`
```

## Set Android related envs
Lets export the Android SDK and Android NDK
```
$ export ANDROID_NDK_HOME=/Users/code_joker/Library/Android/sdk/ndk/19.2.5345600
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
$ touch $GOPATH/src/not.expert/test/test.go
```

Put this as contents of the test.go
```
package test

import (
"fmt"
)

func Greet (name string) string {
message := fmt.Sprintf("Hello World, &s", name) 
return message
}
```

## Build the library
Using gomobile, we can then build the libraries respective to the traget platform, in our case ios and android.

Build framework for ios
```
$ gomobile bind -target=ios -tags nowatchdog -o $GOPATH/build/test.xcframework -v $GOPATH/src/not.expert/test
```

And build aar for android
```
$ gomobile bind -target=android -o $GOPATH/build/test.aar -v $GOPATH/src/not.expert/test
```
-----------------------------------------------------------------------------

## Create app in Xcode(ios)
Lets create native ios app in Xcode

Open Xcode and create a new Xcode app

![Create a new Xcode project](https://miro.medium.com/max/720/1*6H5euen0mZ7MQYQ_QteElw.webp)

Click **Next** button 

Then: 
+ Product name: **testgo**
+ organization identifier: **not.expert**
+ Interface: **Storyboard**

Click **Next**, your project is created.

### Importing a Gomobile framework
Now we are going to import our **test.xcframework** so that we can interact with it in Swift. Open Finder, and navigate to the frameworks folder. Then drag the **test.xcframework** file into Xcode at the bottom of the Project navigator. A dialog will open to confirm the action, be sure to check **Copy items if needed**. And you can follow the [link](https://denbeke.be/blog/programming/go-mobile-example-running-caddy-ios/#:~:text=Building%20an%20iOS%20app%20with%20the%20framework)

And compile your code for ios:
```
$ gomobile build -target=ios -tags nowatchdog -v $GOPATH/src/not.expert/test
```

And run your Xcode project

-------------------------------------------------


# Create app in Android Studio(android)
Lets create native android app in Android Studio

Open Android Studio and create **New project**
And choose **Empty activity**, click **Next**
And set your project name **Name: test**, click **Finish**.
Your native android app is created.

## Import our go project to Android Studio Project
Let's import our go project which compile(.aar & .jar) to Android Studio Project

+ Copy your go project test.aar and test-sources.jar
+ Paste to **test/app/libs** your native android app which create 
+ Add **implementation fileTree(dir: 'libs', include: ['*.jar', '*.aar'])** to test/app/build.gradle in dependencies: {}
+ And Synchronize and rebuild the project

## And inport your go project into MainActivity
```
package not.expert.test;

import androidx.appcompat.app.AppCompatActivity;
import test.Test;
import android.os.Bundle;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        
        Test.greeting("android");
    }
}
```

And Run your android project




