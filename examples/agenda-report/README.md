# agenda-report

In this simple case I just send list of my pending tasks on my localhost's
mailbox for every morning.

To achieve this:

1. Build `agenda-report.go`. You can edit mail address to send.

2. Move `agenda-report.service` and `agenda-report.timer` in
`~/.config/systemd/user/`. You should specify path to builded binary in
`ExecStart` section of `agenda-report.service`.

3. Use following command to set up [systemd timers](https://www.freedesktop.org/software/systemd/man/systemd.timer.html):
```
systemctl --user enable agenda-report.timer
systemctl --user start agenda-report.timer
```

