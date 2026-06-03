Challenge 5  prompts 

What is the difference between a "Hang" and "Connection Refused"?
 A "Connection Refused" means your data packet successfully reached the server, but the computer explicitly told you the door was locked (the port was closed). A "Hang" (or timeout) means your packet was silently dropped by a firewall, like a Security Group, before it ever reached the destination, so you never received an answer at all.


Why is 0.0.0.0/0 on Port 22 a security risk? 
Setting Port 22 to 0.0.0.0/0 exposes your SSH service to the entire internet. This makes your instance a constant target for automated brute-force attacks where hackers try to guess your credentials or exploit the SSH service.


What is the benefit of using Instance Connect? 
Instance Connect provides a secure way to access your instance directly through the AWS Console. This removes the "administrative bottleneck" of having to constantly update your Security Group inbound rules every time your home or school network IP address changes.


What is the difference between "Stopping" and "Terminating" an instance? 
"Stopping" an instance is like powering off your laptop; the machine shuts down, but your data remains on the disk so you can "Start" it again later. "Terminating" is a permanent delete action that wipes the instance and its storage entirely from your account.


Why must the app bind to 0.0.0.0 instead of 127.0.0.1? 
If an app binds to 127.0.0.1 (localhost), it is only listening to itself, like an office worker wearing noise-canceling headphones. Binding to 0.0.0.0 allows the app to "hear" and accept traffic coming from any network interface, which is required for it to be accessible from the outside world.


