# GLDuplicate

The `glduplicate` program allows you to delete duplicate variables. When both prefixed variable and unprefixed variable exists, it deletes the unprefixed variable.

It only modifies the variable export file (`.gitlab.var.json` file).

The application also has a read-only mode, in which only read calls are made: `-dryrun`

## Installation

```
go install github.com/didier13150/glduplicate@latest
```

## Application Usage

```
❯ ./glduplicate -help
Usage: ./glduplicate [options]
  -dryrun
        Run in dry-run mode (read only).
  -prefixenv string
        Var env which value contains prefix (default "*")
  -prefixkey string
        Var key which value contains prefix (default "VAR_PREFIX")
  -prefixsep string
        Separator beztween prefix and real variable name (default "_")
  -varfile string
        File which contains vars. (default ".gitlab-vars.json")
  -verbose
        Make application more talkative.`
```

Debug mode exports the environment and variable tables to the `debug.txt` file.

## File Descriptions

* **variables** file

    ```
    [
      {
        "key": "DEBUG_ENABLED",
        "value": "1",
        "description": null,
        "environment_scope": "*",
        "raw": true,
        "hidden": false,
        "protected": false,
        "masked": false
      }
    ]
    ```
    | Key               | Description                                                         | Value Type      | Default Value | Notes                 |
    | ----------------- | ------------------------------------------------------------------- | --------------- | ------------- | --------------------- |
    | key               | Variable key (unique name per environment)                          | non-null string |               | required              |
    | value             | Variable value                                                      | non-null string |               | required              |
    | description       | Environment description                                             | nullable string | _null_        | optional for creation |
    | environment_scope | Variable scope                                                      | non-null string | __*__         | required              |
    | raw               | Flag indicating that the variable is uninterpretable                | boolean         | false         | required              |
    | hidden            | Flag indicating that the variable should be hidden in the *job* log | boolean         | false         | required              |
    | protected         | Flag indicating that the variable is a protected variable           | boolean         | false         | required              |
    | masked            | Flag indicating that the variable is a masked variable              | boolean         | false         | required              |

    * hidden: Hidden from job logs and can never be revealed in pipelines once the variable is saved.
    * protected: Export the variable to pipelines running only on protected branches and tags.
    * masked: Hidden from job logs, but the value can be revealed in pipelines.


## Runtime

The application can use environment variables to simplify command-line options (it reuses an environment variable from `glcli`)..

| Variable            | Default value               |
| ------------------- | --------------------------- |
| GLCLI_VAR_FILE      | .gitlab-vars.json           |

