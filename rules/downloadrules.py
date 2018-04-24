import requests
import shutil

RULE_LISTS = {
    "http://sysctl.org/cameleon/hosts": "cameleon_hosts.txt",
    "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts": "stevenblack_hosts.txt",
    "https://hosts-file.net/ad_servers.txt": "hosts_filead_servers_hosts.txt",

    "https://mirror1.malwaredomains.com/files/justdomains": "malwaredomains.txt",
    "https://s3.amazonaws.com/lists.disconnect.me/simple_ad.txt": "disconnectme_ad.txt",
    "https://s3.amazonaws.com/lists.disconnect.me/simple_tracking.txt": "disconnectme_tracking.txt",
    "https://zeustracker.abuse.ch/blocklist.php?download=domainblocklist": "zeustracker.txt"
    
}

def main():
    for k in RULE_LISTS:
        v = RULE_LISTS[k]

        print("Downloading {0}".format(v))
        with open(v, "wb+") as f:
            r = requests.get(k)
            f.write(r.content)


if __name__ == "__main__":
    main()