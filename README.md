# Global Entry Slot Notifier

Global Entry Slot Notifier is a command-line tool that checks for available Global Entry appointment slots at specified 
locations and sends notifications via an [app](https://ntfy.sh/) or system notification.

## Features

- Periodically checks for available Global Entry appointment slots.
- Sends notifications via an app (using [ntfy.sh](https://ntfy.sh/) download on Android/iPhone/Mac) or system notification 
  (using the beeep package).
- Configurable via command-line flags.

## Usage

```bash
./global-entry-slot-notifier --location <location_id> --notifier <notifier_type> [--topic <ntfy_topic>] [--interval <duration>]
```

### Flags
* `--location` (required): Specify the location ID for the Global Entry appointment.
* `--notifier` (required): Specify the notifier type (app or system).
* `--topic` (required if notifier is app): Specify the ntfy.sh [topic](https://docs.ntfy.sh/) to send notifications to.
* `--interval` (optional): Specify the interval (in seconds, e.g. 30s) at which to check for available appointments. Default is 60s.

### Examples

1. System Notification

```bash
./global-entry-slot-notifier --location 5001 --notifier system --interval 90s
```

2. App Notification

```bash
./global-entry-slot-notifier --location 5001 --notifier app --topic my-ntfy-topic
```

### License
This project is licensed under the MIT License. See the LICENSE file for details.

### Contributing
Contributions are welcome! Please open an issue or submit a pull request for any changes.

### Contact
For questions or feedback, please contact arun@gopalpuri.com