
1.  Why did my Jenkins pipeline fail with `ERROR: Could not find credentials entry with ID 'target-ssh'` right after I pushed my Kubernetes code?
    Jenkins checked out the old `main` branch instead of the new `challenge/deploy-to-kubernetes` branch. Because this is a fresh playground, it didn't have the old SSH keys from previous challenges, causing the crash. You need to manually select the correct branch in the Jenkins UI.

2.  Why did I get a `Forbidden: pod updates may not change fields` error when I tried to add resource limits to my running Pod?
    Kubernetes locks in the physical resource constraints of a pod once it is running, so you cannot hot-swap memory limits on a live pod. You must delete the existing pod using `kubectl delete pod myapp --ignore-not-found=true` before applying the new manifest.

3.  Why did my pipeline suddenly fail with `error: You must be logged in to the server (Unauthorized)` when deploying to Kubernetes?
    The Kubernetes security token generated for the `jenkins-robot` service account expired. Modern Kubernetes tokens are temporary and usually expire after 1 hour. You need to generate a new token and update it in the Jenkins Credentials vault.

4.  If Jenkins can pull the `ttl.sh` image successfully, why does Kubernetes fail with an `ImagePullBackOff`?
    Jenkins and Kubernetes are different computers on potentially different networks. Jenkins pulling the image just proves the CI/CD server has internet access. The Kubernetes worker node might be behind a strict firewall or lack public internet routing, preventing it from reaching `ttl.sh`.

5.  What is the difference between Liveness and Readiness probes, and why aren't they redundant?
    A Liveness Probe checks if the application is frozen; if it fails, Kubernetes kills and restarts the pod. A Readiness Probe checks if the application is ready to receive network traffic (e.g., done loading a database). An app can be "alive" but not "ready", which is why both are needed.

6.  What goes wrong if you don't set memory requests vs. if you don't set limits for a Pod?
    If neither are set, the pod can consume unlimited CPU/memory, potentially crashing the entire worker node (the "noisy neighbor" problem). If only limits are set (without requests), the pod has a maximum budget but no guaranteed baseline; if the server becomes crowded, Kubernetes will randomly kill this pod first to save space.


