# network-fanout
A TCP / UDP packet proxy fanout

```shell
APP:
network-fanout

COMMAND:
server


AVAILABLE SUBCOMMANDS:
help : Print this help message

PARSING ORDER: (set values will override in this order)
CLI Flag > Toml Config > JSON Config > Environment

VARIABLES:
+--------------------+-----------+----------+------------------+--------------------------------+
|        FLAG        |  DEFAULT  | REQUIRED |     ENV NAME     |          DESCRIPTION           |
+--------------------+-----------+----------+------------------+--------------------------------+
| --targets          |        -- | Required | TARGETS          |                                |
| --log-output       | text      | No       | LOG_OUTPUT       | Choose between text and json   |
| --log              | info      | No       | LOG              | You can set from more logs to  |
|                    |           |          |                  | less logs: debug, info, warn,  |
|                    |           |          |                  | error or fatal                 |
| --mode             | tcp       | No       | MODE             | You can use tcp or udp         |
| --source-host      | localhost | No       | SOURCE_HOST      | Hostname to use as source host |
| --port             |      8083 | No       | PORT             |                                |
| --read-buffer-size |      1024 | No       | READ_BUFFER_SIZE |                                |
+--------------------+-----------+----------+------------------+--------------------------------+

```