package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/didier13150/gitlablib"
)

func main() {

	prefixKey := "VAR_PREFIX"
	prefixEnv := "*"
	prefixSep := "_"
	varsFile := ".gitlab-vars.json"
	verboseMode := false
	dryrunMode := false

	if len(os.Getenv("GLCLI_VAR_FILE")) > 0 {
		varsFile = os.Getenv("GLCLI_VAR_FILE")
	}

	var varsFileOpt = flag.String("varfile", varsFile, "File which contains vars.")
	var prefixKeyOpt = flag.String("prefixkey", prefixKey, "Var key which value contains prefix")
	var prefixEnvOpt = flag.String("prefixenv", prefixEnv, "Var env which value contains prefix")
	var prefixSepOpt = flag.String("prefixsep", prefixSep, "Separator between prefix and real variable name")
	var verboseOpt = flag.Bool("verbose", false, "Make application more talkative.")
	var dryrunOpt = flag.Bool("dryrun", false, "Run in dry-run mode (read only).")

	flag.Usage = func() {
		fmt.Print("Remove duplicate variables (keep the one with the prefix) from glcli export file\n\n")
		fmt.Printf("Usage: " + os.Args[0] + " [options]\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *verboseOpt {
		log.Print("Verbose mode is active")
		verboseMode = true
	}
	if *dryrunOpt {
		log.Print("Dry run mode is active")
		dryrunMode = true
	}
	if varsFileOpt != nil {
		varsFile = *varsFileOpt
	}
	if prefixKeyOpt != nil {
		prefixKey = *prefixKeyOpt
	}
	if prefixEnvOpt != nil {
		prefixEnv = *prefixEnvOpt
	}
	if prefixSepOpt != nil {
		prefixSep = *prefixSepOpt
	}

	glvar := gitlablib.NewGitlabVar("", "", verboseMode)
	if dryrunMode {
		glvar.DryrunMode = dryrunMode
	}
	if verboseMode {
		log.Printf("Importing variable from %s file\n", varsFile)
	}
	glvar.ImportVars(varsFile)

	varPrefix := getValue(&glvar, prefixKey, prefixEnv) + prefixSep
	if verboseMode {
		log.Printf("Var prefix is defined to \"%s\"\n", varPrefix)
	}

	var prefixedVars []gitlablib.GitlabVarData
	for _, varItem := range glvar.FileData {
		if len(varItem.Key) > len(varPrefix) {
			if varItem.Key[0:len(varPrefix)] == varPrefix {
				prefixedVars = append(prefixedVars, varItem)
			}
		}
	}

	for _, varItem := range prefixedVars {
		if len(varItem.Key) <= len(varPrefix) {
			continue
		}
		duplicateKey := varItem.Key[len(varPrefix):len(varItem.Key)]

		if getValue(&glvar, duplicateKey, varItem.Env) != "" {
			if verboseMode {
				log.Printf("Found duplicate for prefixed var %s[%s] : %s", varItem.Key, varItem.Env, duplicateKey)
			}
			if dryrunMode {
				log.Printf("%s[%s] should be deleted but dry run mode is active", duplicateKey, varItem.Env)
			} else {
				for i, item := range glvar.FileData {
					if item.Key == duplicateKey && item.Env == varItem.Env {
						log.Printf("Deleting %s[%s] from variable file", duplicateKey, varItem.Env)
						glvar.FileData = append(glvar.FileData[:i], glvar.FileData[i+1:]...)
					}
				}
			}
		}
	}
	if !dryrunMode {
		if verboseMode {
			log.Printf("Exporting variable to %s file\n", varsFile)
		}
		glvar.GitlabData = glvar.FileData
		glvar.ExportVars(varsFile)
	}
}

func getValue(glvar *gitlablib.GitlabVar, key string, env string) string {
	for _, varItem := range glvar.FileData {
		if varItem.Key == key && varItem.Env == env {
			return varItem.Value
		}
	}
	return ""
}
