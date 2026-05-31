DEBUG.md

Two ranked hypotheses :

1. The SSH logout kills the app: When Jenkins logs in, it opens a temporary SSH session. It starts the app inside this session. When Jenkins logs out, Linux destroys that session and sends a kill signal to everything inside it, including the app.

2. The app is hiding from the outside: The app might actually be running in the background, but it is only listening to internal traffic.


One verification step per hypothesis :

 To test the SSH logout theory: I would log into the server, start the app using `nohup ./main &` (which tells the app to ignore the kill signal), and then log out. If I can reach the app from my laptop after logging out, it proves the SSH logout was the killer.

 To test the hiding theory: I would log into the server and type the command ss -tlnp | grep 4444.. If the screen shows the numbers 127.0.0.1, it means the app is locked in "internal only" mode. It is only answering calls from inside the server and completely ignoring outside connections from my laptop. 


 The Fix
To fix this the right way, we use systemd. I created a file at `/etc/systemd/system/main.service`. This acts like a permanent "night watchman keeping the app running safely in the background even after users log out.


 The Lesson
Real server apps need a "process supervisor" (like systemd) to run independently in the background so they can survive without a human logged in.
a
