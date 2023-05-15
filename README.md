# Crda CLI 1.5

This project is still in development mode.
For analysis, currently, only Java's Maven ecosystem is implemented.

- Staging version (Pre-release) can be found [here](https://github.com/RHEcosystemAppEng/crda-cli/releases/tag/staging).
- Sample projects can be found in the [crda-usage-examples repo](https://github.com/RHEcosystemAppEng/crda-usage-examples)

## Run using executable

Download the binary based on your OS from the [Releases tab](https://github.com/RHEcosystemAppEng/crda-cli/releases).

The _Html_ report is saved as a local file in the OS's temporary folder.

```shell
$ crda analyse /path/to/maven/project/pom.xml

Summary Report for Dependency Analysis:

Total Scanned Dependencies:  10
Total Scanned Transitive Dependencies:  193
Direct Vulnerable Dependencies:  4
Total Vulnerabilities:  14
Critical Vulnerabilities:  0
High Vulnerabilities:  3
Medium Vulnerabilities:  8
Low Vulnerabilities:  3

Full Report:  file:///tmp/crda/stack-analysis-maven-1684149652.html
```

## Run using image

Running using and image will **not** create a _Html_ report.
Instead, it will print the _Json_ version of the report.<br/>
This behavioural is also achievable manually using the `--json` flag:

```shell
$ crda analyse /path/to/maven/project/pom.xml --json

{
        "dependencies": {
                "scanned": 10,
                "transitive": 193
        },
        "vulnerabilities": {
                "critical": 0,
                "direct": 4,
                "high": 3,
                "low": 3,
                "medium": 8,
                "total": 14
        }
}
```

### Java

From your project path (replace PWD with path if needed):

```shell
podman run --rm -it \
    -v $HOME/.m2:/opt/app-root/src/.m2 \
    -v $HOME/.crda:/opt/app-root/src/.crda \
    -v $PWD:/app \
    quay.io/ecosystem-appeng/crda-cli:staging analyse pom.xml
```

## Help

```shell
$ crda help

Use this tool for CodeReady Dependency Analytics reports

Usage:
  crda [command]

Available Commands:
  analyse     Preform dependency analysis report
  auth        Link crda user with snyk
  completion  Generate a completions script
  config      Manage crda config
  help        Help about any command
  version     Get binary version

Flags:
  -m, --client string   The invoking client for telemetry (default "terminal")
  -d, --debug           Set DEBUG log level
  -c, --no-color        Toggle colors in output.

Use "crda [command] --help" for more information about a command.
```
