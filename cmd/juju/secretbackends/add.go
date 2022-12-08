// Copyright 2022 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package secretbackends

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/juju/cmd/v3"
	"github.com/juju/errors"
	"github.com/juju/gnuflag"
	"github.com/juju/utils/v3/keyvalues"
	"gopkg.in/yaml.v2"

	"github.com/juju/juju/api/client/secretbackends"
	jujucmd "github.com/juju/juju/cmd"
	"github.com/juju/juju/cmd/modelcmd"
	"github.com/juju/juju/secrets/provider"
	_ "github.com/juju/juju/secrets/provider/all"
)

type addSecretBackendCommand struct {
	modelcmd.ControllerCommandBase
	out cmd.Output

	AddSecretBackendsAPIFunc func() (AddSecretBackendsAPI, error)

	Name        string
	BackendType string

	// Attributes from a file.
	ConfigFile cmd.FileVar
	// Attributes from key value args.
	KeyValueAttrs map[string]interface{}
}

var addSecretBackendsDoc = `
Adds a new secret backend for storing secret content.

You must specify a name for the backend and its type,
followed by any necessary backend specific config values.
Config may be specified as key values ot read from a file.
Any key values override file content if both are specified.

To rotate the backend access credential/token (if specified), use
the "token-rotate" config and supply a duration.

Examples:
    juju add-secret-backend myvault vault --config /path/to/cfg.yaml
    juju add-secret-backend myvault vault token-rotate=10m --config /path/to/cfg.yaml
    juju add-secret-backend myvault vault endpoint=https://vault.io:8200 token=s.1wshwhw
`

// AddSecretBackendsAPI is the secrets client API.
type AddSecretBackendsAPI interface {
	AddSecretBackend(secretbackends.SecretBackend) error
	Close() error
}

// NewAddSecretBackendCommand returns a command to list secrets backends.
func NewAddSecretBackendCommand() cmd.Command {
	c := &addSecretBackendCommand{}
	c.AddSecretBackendsAPIFunc = c.secretBackendsAPI

	return modelcmd.WrapController(c)
}

func (c *addSecretBackendCommand) secretBackendsAPI() (AddSecretBackendsAPI, error) {
	root, err := c.NewAPIRoot()
	if err != nil {
		return nil, errors.Trace(err)
	}
	return secretbackends.NewClient(root), nil

}

// Info implements cmd.Info.
func (c *addSecretBackendCommand) Info() *cmd.Info {
	return jujucmd.Info(&cmd.Info{
		Name:    "add-secret-backend",
		Purpose: "Add a new secret backend to the controller.",
		Doc:     addSecretBackendsDoc,
	})
}

// SetFlags implements cmd.SetFlags.
func (c *addSecretBackendCommand) SetFlags(f *gnuflag.FlagSet) {
	f.Var(&c.ConfigFile, "config", "path to yaml-formatted configuration file")
}

func (c *addSecretBackendCommand) Init(args []string) error {
	if len(args) < 2 {
		return errors.New("must specify backend name and type")
	}
	c.Name = args[0]
	c.BackendType = args[1]
	args = args[2:]
	// The remaining arguments are divided into keys to set.
	c.KeyValueAttrs = make(map[string]interface{})
	for _, arg := range args {
		splitArg := strings.SplitN(arg, "=", 2)
		if len(splitArg) != 2 {
			return errors.NotValidf("key value %q", arg)
		}
		key := splitArg[0]
		if len(key) == 0 {
			return errors.Errorf(`expected "key=value", got %q`, arg)
		}
		if _, exists := c.KeyValueAttrs[key]; exists {
			return keyvalues.DuplicateError(
				fmt.Sprintf("key %q specified more than once", key))
		}
		c.KeyValueAttrs[key] = splitArg[1]
	}

	if len(c.KeyValueAttrs) == 0 && c.ConfigFile.Path == "" {
		return errors.New("must specify a config file or key values")
	}
	return nil
}

func (c *addSecretBackendCommand) readFile(ctx *cmd.Context) (map[string]interface{}, error) {
	attrs := make(map[string]interface{})
	if c.ConfigFile.Path == "" {
		return attrs, nil
	}
	var (
		data []byte
		err  error
	)
	if c.ConfigFile.Path == "-" {
		// Read from stdin
		data, err = io.ReadAll(ctx.Stdin)
	} else {
		// Read from file
		data, err = c.ConfigFile.Read(ctx)
	}
	if err != nil {
		return nil, errors.Trace(err)
	}

	if err := yaml.Unmarshal(data, &attrs); err != nil {
		return nil, errors.Trace(err)
	}
	return attrs, nil
}

// Run implements cmd.Run.
func (c *addSecretBackendCommand) Run(ctxt *cmd.Context) error {
	attrs, err := c.readFile(ctxt)
	if err != nil {
		return errors.Trace(err)
	}
	for k, v := range c.KeyValueAttrs {
		attrs[k] = v
	}

	const tokenRotate = "token-rotate"
	var tokenRotateInterval *time.Duration
	tokenRotateStr, ok := attrs[tokenRotate]
	if ok {
		delete(attrs, tokenRotate)
		rotateInterval, err := time.ParseDuration(fmt.Sprintf("%s", tokenRotateStr))
		if err != nil || (rotateInterval/time.Second) < 60 {
			return errors.NotValidf("token rotate interval %q", tokenRotateStr)
		}
		tokenRotateInterval = &rotateInterval
	}

	p, err := provider.Provider(c.BackendType)
	if err != nil {
		return errors.Annotatef(err, "invalid secret backend %q", c.BackendType)
	}
	configValidator, ok := p.(provider.ProviderConfig)
	if ok {
		err = configValidator.ValidateConfig(nil, attrs)
		if err != nil {
			return errors.Annotate(err, "invalid provider config")
		}
	}

	backend := secretbackends.SecretBackend{
		Name:                c.Name,
		BackendType:         c.BackendType,
		TokenRotateInterval: tokenRotateInterval,
		Config:              attrs,
	}
	api, err := c.AddSecretBackendsAPIFunc()
	if err != nil {
		return errors.Trace(err)
	}
	defer api.Close()

	err = api.AddSecretBackend(backend)
	return errors.Trace(err)
}