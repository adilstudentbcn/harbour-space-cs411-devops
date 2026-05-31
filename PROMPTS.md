# Session Log: Q&A

**Why was the deployment failing with "Permission denied" at the beginning?**
The deployment failed because the Jenkins service, which runs the pipeline, did not have the correct SSH identity to connect to the Docker VM. Even though you had access as the 'laborant' user, Jenkins needed its own authorized access, which required us to manually add the SSH private key to the Jenkins global credentials vault and configure the pipeline to use it during the deployment step.

**Why did the pipeline freeze when we first tried to deploy to the docker VM?**
The pipeline froze because we used the agent { label 'docker' } directive, which instructed Jenkins to wait for a specific worker node named "docker" to become available. However, that node was not registered or active in the Jenkins system. Because the pipeline was programmed to wait indefinitely for that specific machine to report for duty, it simply sat there in a "pending" state forever.

**What did we forget in the Jenkinsfile that caused the deployment to error out?**
We forgot to specify the correct SSH user in the deployment command. Initially, Jenkins tried to connect as the user jenkins@docker, which did not have the necessary permissions. Once we updated the command to use laborant@docker, Jenkins was able to successfully authenticate and execute the commands on the VM.

**Why did we decide to skip the HEALTHCHECK task to make things easier?**
We decided to skip the HEALTHCHECK task because it was unnecessarily complex for this environment. Since we were using a scratch base image (which is completely empty and lacks standard tools like curl or wget), adding a health check would have required writing and compiling extra Go code just to monitor the container, which would have increased the build complexity and maintenance effort without significantly adding to the core project goal.

**What was the main reason our build failed on the server but worked on my laptop?**
The main reason was an architecture mismatch known as an "exec format error." Your laptop uses an Apple Silicon (ARM64) processor, while the Docker VM uses an Intel/AMD (x86_64) processor. When you built the binary, it was compiled with the specific machine-language instructions for your laptop's CPU, which the VM could not understand. By adding GOOS=linux and GOARCH=amd64 to your build command, you ensured the binary was compiled specifically for the server's CPU architecture instead.
