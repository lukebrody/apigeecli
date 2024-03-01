## apigeecli securityprofiles detach

Detach a security profile from an environment

### Synopsis

Detach a security profile from an environment

```
apigeecli securityprofiles detach [flags]
```

### Options

```
  -e, --env string    Apigee environment name
  -h, --help          help for detach
  -n, --name string   Name of the security profile
```

### Options inherited from parent commands

```
  -a, --account string   Path Service Account private key in JSON
      --default-token    Use Google default application credentials access token
      --disable-check    Disable check for newer versions
      --metadata-token   Metadata OAuth2 access token
      --no-output        Disable printing all statements to stdout
  -o, --org string       Apigee organization name
      --print-output     Control printing of info log statements (default true)
  -t, --token string     Google OAuth Token
```

### SEE ALSO

* [apigeecli securityprofiles](apigeecli_securityprofiles.md)	 - Manage Adv API Security Profiles

###### Auto generated by spf13/cobra on 19-Dec-2023