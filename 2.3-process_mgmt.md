# Day 3: Process Management (ps, top, kill)

## Introduction
Process management is a fundamental concept in Linux. It involves monitoring and controlling running processes to optimize system performance and resource usage. In this section, we will explore three essential commands: `ps`, `top`, and `kill`.

---

## 1. Viewing Processes with `ps`
The `ps` command displays information about active processes. Unlike `top`, which updates dynamically, `ps` provides a static snapshot.

### Common `ps` Commands:
- `ps` – Shows processes running in the current terminal.
- `ps aux` – Lists all running processes with details like user, CPU, and memory usage.
- `ps -ef` – Displays all processes in a different format, including parent-child relationships.
- `ps -u <username>` – Lists processes belonging to a specific user.
- `ps -p <PID>` – Displays information about a specific process.

### Example Output:
```bash
PID    USER  %CPU %MEM COMMAND
1234   root   0.1  0.5  nginx
5678   user1  0.3  1.2  firefox
```

---

## 2. Monitoring Processes with `top`
The `top` command provides a real-time, continuously updated view of system processes.

### Key Features:
- Displays CPU and memory usage per process.
- Allows sorting by CPU, memory, or process ID.
- Interactive: Press `h` for help, `q` to quit, and `k` to kill a process.

### Useful Options:
- `top -u <username>` – Shows processes for a specific user.
- `top -o %MEM` – Sorts processes by memory usage.

### Example Output:
```
PID   USER   %CPU  %MEM  COMMAND
1234  root   10.0   2.5  nginx
5678  user1  15.3   4.0  firefox
```

---

## 3. Managing Processes with `kill`
The `kill` command terminates processes based on their PID (Process ID).

### Common `kill` Commands:
- `kill <PID>` – Sends the default `TERM` signal to terminate a process gracefully.
- `kill -9 <PID>` – Forces a process to stop immediately using the `KILL` signal.
- `killall <process_name>` – Terminates all instances of a process.
- `pkill <name>` – Kills processes by name pattern.

### Example:
```bash
kill 1234      # Attempts to terminate process with PID 1234
kill -9 5678   # Forces termination of process with PID 5678
killall firefox  # Stops all running Firefox processes
```

---

## Conclusion
Understanding process management in Linux helps maintain system stability and optimize performance. Commands like `ps`, `top`, and `kill` allow users to monitor and control processes efficiently.

### Additional Tips:
- Use `htop` for an enhanced `top` alternative with better UI.
- Use `nice` and `renice` to manage process priority.
- Check logs for terminated processes using `dmesg` or `journalctl`.

Stay tuned for **Day 4: Job Control (bg, fg, jobs, nohup)**!

