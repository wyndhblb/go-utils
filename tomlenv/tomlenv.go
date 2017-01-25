/*
Normal TOML does not support ENV variables in the config files.

This gets "around" that by basically replacing

$ENV{VAR_NAME:default}

w/ the VAR_NAME from the env if present or the default

a simple strings replacer

*/

package tomlenv

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"regexp"
)

var envReg *regexp.Regexp

func init() {
	envReg = regexp.MustCompile(`\$ENV\{(.*?)\}`)
}

func DecodeFile(filename string, cfg interface{}) (meta toml.MetaData, err error) {

	//slurp in the file as we need to do some replacing
	bits, err := ioutil.ReadFile(filename)
	if err != nil {
		return meta, err
	}
	return DecodeBytes(bits, cfg)

}

func doReplace(inbys []byte) []byte {

	// ye-old-replace
	for _, mtch := range envReg.FindAllSubmatch(inbys, -1) {
		if len(mtch) != 2 {
			continue
		}
		if len(mtch[0]) == 0 {
			continue
		}

		inbtween := bytes.Split(mtch[1], []byte(":"))
		envvar := string(inbtween[0])
		def := []byte("")
		if len(inbtween) >= 2 {
			def = inbtween[1]
		}
		env := os.Getenv(envvar)
		if len(env) > 0 {
			inbys = bytes.Replace(inbys, mtch[0], []byte(env), -1)
		} else {
			inbys = bytes.Replace(inbys, mtch[0], def, -1)
		}
	}
	return inbys
}

func DecodeBytes(inbys []byte, cfg interface{}) (toml.MetaData, error) {
	inbys = doReplace(inbys)
	return toml.Decode(string(inbys), cfg)
}

func Decode(instr string, cfg interface{}) (toml.MetaData, error) {
	return DecodeBytes([]byte(instr), cfg)
}
