 
Debugging Scenario: ImagePullBackOff

 1. Hypotheses

 	A: The Kubernetes worker nodes are running in a private network without public internet access, so they cannot reach the `ttl.sh` website even though the Jenkins server can.

 	B: The image on `ttl.sh` expired. Since it only lives for 2 hours, it might have been deleted between the time Jenkins pushed it and Kubernetes tried to pull it.

 2. Verification Steps

 To verify the network: Run kubectl run test-ping --image=alpine -- sh -c "ping -c 3 ttl.sh" to see if the cluster can actually connect to the website.

 To verify expiration: Run kubectl describe pod myapp  to read the exact error message from the Kubernetes engine.


 3. The Fix

Instead of putting the image on a public website, upload it to a private storage area that Kubernetes servers already have permission to access. 


 4. Lesson

Just because Jenkins can download the image doesn't mean Kubernetes can. Jenkins has its own internet connection, but the Kubernetes servers are completely separate computers—they need their own network rules and permissions to talk to that same registry. 

