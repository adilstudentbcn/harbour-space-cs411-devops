DEBUG.md
Challenge 5 
1. Ranked Hypotheses
Hypothesis 1 : The AWS Security Group is acting like a closed gate at the edge of the network. It isn't allowing outside traffic to reach port 4444, so the request gets "dropped" and never arrives, causing the laptop to wait forever.

Hypothesis 2: The application itself is only listening to its own internal "loopback" address and can’t hear any requests coming from the outside world.

2. Verification Steps
For Hypothesis 1: we can  check the "Inbound Rules" in the AWS Security Group settings for the EC2 instance. If there isn't a specific rule that allows TCP traffic on port 4444 from the outside world (0.0.0.0/0), that’s the problem.


For Hypothesis 2:  I will run netstat -tulpn | grep 4444. If the output shows 127.0.0.1:4444, it confirms the app is only listening locally and ignoring the internet.


3. The Fix
Fix for Hypothesis 1: Add an "Inbound Rule" in the AWS Security Group to allow port 4444 traffic from anywhere.
Fix for Hypothesis 2: Change the application's configuration or code so that it binds to 0.0.0.0 instead of just 127.0.0.1 .


4. The Underlying Lesson
A "Connection Refused" means your request reached the server and was rejected . A "Hang/Timeout" means your request disappeared because the Security Group silently dropped it before it could ever reach the destination.

