# Android Screen Time Utility

This command-line utility makes your Android phone much less distracting. Imagine a dumbphone - with only basic functionality like calls, messages, music, directions - plus NFC payments and your favorite Play Store apps.

Here's how it works:
1. Remove all non-essential apps from your phone. All you really need are calls, texts, maps, notes, calendars, music, and timers/alarms. Maybe keep some others like banking, tickets/passes, or travel, but remove anything that isn't important. Delete distracting apps such as social media and web browsers.
2. Install this project's executable to your PATH, then run "android" in your terminal with your phone connected to your computer. It will remove the Play Store and Chrome.

That's it! It's effective because once you remove distracting apps and access to the wide internet, there's not much left to do on your phone, and no way to re-install them without connecting your phone back to your computer.

If you realize you need to install an app on your phone:
1. Run "android" in your terminal. This re-installs the Play Store and Chrome.
2. Install and set up the app, as usual
3. Run "android" in your terminal again to uninstall the Play Store and Chrome.

### Installation

You can download a pre-built executable from the "Actions" tab, which you should then add to your PATH.

Alternatively, install it manually by running:

`go build`  
`go install`
