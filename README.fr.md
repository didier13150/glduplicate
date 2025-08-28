# GLDuplicate

Le programme `glduplicate` permet la suppression des variables dupliquées. Lorsqu'il existe une variable préfixée et un doublon sans préfixe, il supprime la variable sans préfixe.

Il ne fait que modifier le fichier d'export des variables (fichier `.gitlab.var.json`).

L'application possède également un mode lecture seule, qui liste les modifications sans les appliquer: `-dryrun`

## Usage de l'application

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

## Description des fichiers

* Fichier concernant **les variables**

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
    
    | Clé               | Description                                                                   | Type de valeur                           | Valeur par défaut | Remarques                        |
    | ----------------- | ----------------------------------------------------------------------------- | ---------------------------------------- | ----------------- | -------------------------------- |
    | key               | Clé de la variable (nom unique par environnement)                             | chaîne de caractères non nulle           |                   | obligatoire                      |
    | value             | Valeur de la variable                                                         | chaîne de caractères non nulle           |                   | obligatoire                      |
    | description       | Description de l'environnement                                                | chaîne de caractères qui peut être nulle | _null_            | facultatif pour la création      |
    | environment_scope | Portée de la variable                                                         | chaîne de caractères non nulle           | __*__             | obligatoire                      |
    | raw               | Drapeau indiquant que la variable est une variable non interprétable          | boolean                                  | false             | obligatoire                      |
    | hidden            | Drapeau indiquant que la variable doit être cachée dans le journal des *jobs* | boolean                                  | false             | obligatoire                      |
    | protected         | Drapeau indiquant que la variable est une variable protégée                   | boolean                                  | false             | obligatoire                      |
    | masked            | Drapeau indiquant que la variable est une variable masquée                    | boolean                                  | false             | obligatoire                      |
    

    * hidden: Masqué dans les journaux des *jobs* et ne peut jamais être révélé dans les pipelines une fois la variable enregistrée.
    * protected: Exporter la variable vers les pipelines exécutés uniquement sur des branches et des *tags* protégés.
    * masked: Masqué dans les journaux des *jobs*, mais la valeur peut être révélée dans les pipelines.

## Utilisation

L'application peut utiliser des variables d'environnement afin de simplifier les options de la ligne de commande (il ré-utilise une variable d'environnement de `glcli`).

| Variable           | valeur par défaut           |
| ------------------ | --------------------------- |
| GLCLI_VAR_FILE     | .gitlab-vars.json           |


