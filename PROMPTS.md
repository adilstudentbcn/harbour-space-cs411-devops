

I want to complete the systemd Stretch Task. How do I configure my unit file so that it runs as a non-root user and automatically restarts on failure?

To fulfill the Stretch Task requirements, you need to apply two specific directives inside the [Service] block of your systemd unit file.
Here is what that configuration looks like:
Ini, TOML
[Unit]
Description=CS411 Application
After=network.target

[Service]
User=laborant
ExecStart=/home/laborant/main
Restart=on-failure

[Install]
WantedBy=multi-user.target


How this fulfills the requirements:
User=laborant: This explicitly tells systemd to run the application as the standard laborant user rather than root. This is a critical security best practice known as the "Principle of Least Privilege." If the application is ever compromised, the attacker only gains access to that restricted user account, not the entire server.
Restart=on-failure: This instructs systemd's process supervisor to monitor the application. If the application exits with a non-zero exit code (meaning it crashed or encountered a fatal error), systemd will automatically restart it to ensure high availability.





—--------------------------------------------------------------------------
Can you explain the Debug Scenario in simple terms? Why does my app die when I close the SSH session, and why is systemd the best fix for this?

Here is what happens when you run an application manually during an SSH session:
Why the app dies
When you log into a server via SSH, Linux opens a temporary workspace (your shell session). Any program you start directly in that terminal is considered a "child" of your session. When you disconnect, Linux cleans up by destroying that workspace. It sends a "hangup" signal (SIGHUP) to all child programs, which instantly kills your application.
Why systemd is the best fix
While you could use a command like nohup (which simply tells the app to put on "earplugs" and ignore the SIGHUP signal), it is just a temporary band-aid. If your app crashes an hour later because of a code error, it will stay dead.
systemd is the superior solution because it provides process supervision. When you configure systemd:
It completely detaches the application from your temporary SSH session.
It hands control over to the server's core initialization daemon, running it safely in the background.
It monitors the application continuously. If the app crashes at 3:00 AM, systemd automatically restarts it without a human needing to log in.

—-----------------------------------


My pipeline ran successfully, but Ansible said changed=0 and the auto-grader still says 1/2. How do I configure Jenkins to read from my new branch instead of main?

This is a very common scenario in CI/CD pipelines.
Why this happens: When you initially create a Jenkins pipeline linked to a Git repository, the default configuration tells Jenkins to only monitor and pull code from the main (or master) branch. Because you pushed your new systemd and playbook fixes to your specific challenge branch (challenge/first-deployment-pipeline), Jenkins essentially ignored them. It woke up, pulled the old code from main, and deployed that instead. Since the old code on the main branch hadn't changed, Ansible correctly reported changed=0.
How to fix it: To instruct Jenkins to read from your new branch, you need to update its configuration:
Go to your Jenkins dashboard and select your job.
Click on Configure in the left-hand menu.
Scroll down to the Source Code Management section (or the Pipeline section, if configured as a Pipeline script from SCM).
Look for the Branches to build setting (specifically the Branch Specifier field).
Change the value from */main to exactly */challenge/first-deployment-pipeline.
Click Save at the bottom of the page.
The next time you click Build Now, Jenkins will pull the updated code from your branch, and Ansible will successfully apply the new changes to the target server.





—------------------------------------------------
The auto-checker is still failing because it says the service isn't enabled. What is the difference between 'started' and 'enabled' in Linux, and how do I add this to my Ansible playbook?


That is an excellent question and a very common gotcha in Linux system administration.
Here is the fundamental difference between the two concepts:
Started (state: started): This turns the application on immediately. However, it is not persistent. If the server reboots or loses power, the application will remain off when the server comes back online.
Enabled (enabled: true): This creates a symbolic link in the system's startup directories. It essentially tells the Linux operating system, "Every time you boot up, start this application automatically."
The auto-checker is looking for the enabled state because a true production application must be able to survive a server reboot without a human needing to log in and turn it back on.
How to add this to your Ansible playbook
To fix this, you just need to add enabled: true to the systemd task in your playbook.yml. Here is exactly what that block should look like:
YAML
   - name: Ensure the service is running
      become: True
      ansible.builtin.systemd:
        name: myapp.service
        daemon_reload: true
        enabled: true
        state: started

By having both enabled: true and state: started, Ansible ensures the app is running right now and will automatically start if the server ever reboots.

