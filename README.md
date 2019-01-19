# Battery Notifier :battery:
This is a simple application to show current battery status (percentage and charging status). It can also watch the battery continuously and send periodic desktop notifications when the battery goes below a threshold percentage.


## Options :construction_worker:
* `-v`: show application version.
* `-t`: battery percentage threshold, below which the battery will be condiered as *low* and the user will start getting desktop notifications about low battery.

* `-l`: battery check interval during low (< threshold) battery.

* `-n`: battery check interval during good/normal (> threshold) battery.

* `-w`: continuously watch battery level at preset interval. The interval depends on values of '-n' and '-l'.

* `-h`: get help message and default values of flags/options.

## Creating a *Systemd Service* to keep watching battery even across reboots
* Copy the executable to /usr/local/bin/

* Create a systemd service unit file like below:
```
$ cat /lib/systemd/system/battery-notifier.service 
[Unit]
Description=System Battery Monitor and Notifier
After=multi-user.target

[Service]
User=<normal system username but not root>
ExecStart=/usr/local/bin/battery_notifier -w
Restart=on-failure
RestartSec=30s

[Install]
WantedBy=multi-user.target
```

* Save the above file and perform the following to always start the battery-notifier service:
```
$ sudo systemctl daemon-reload
$ sudo systemctl start battery-notifier
$ sudo systemctl enable battery-notifier
```
