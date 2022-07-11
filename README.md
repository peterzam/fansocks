# Fansocks

### How it works:

![Diagram](https://i.imgur.com/pogx78g.png)

---

### Executable Binary Usage :
```bash

$ ./fansocks
  -bind string
        Bind Address for SOCKS5 (default "0.0.0.0:1080")
  -csv string
        SOCKS5 server list csv file (default "socks.csv")
```

---

### Container Build
```
$ podman build -t peterzam/fansocks .
```

### Container Run:

```bash
$ podman run -d -p 1080:1080 -v ./socks.csv:/socks.csv fansocks
```

---

## LICENSE

Fansocks is licensed under [GNU Affero General Public License](https://www.gnu.org/licenses/agpl-3.0.en.html), AGPL.

---