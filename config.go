package tidy

import "errors"

func Configure() config {
	return config{}
}

type config struct {
	backends []LeveledBackend
}

type toBuilder struct {
	level Level
	cfg   config
}

func (this toBuilder) To(backend BackendBuilder) config {
	cfg := this.cfg
	cfg.backends = append(cfg.backends, NewLeveledBackend(this.level, backend.Build()))
	return cfg
}

func (this config) LogFromLevel(level Level) toBuilder {
	return toBuilder{
		level: level,
		cfg:   this,
	}
}

func (this config) BuildDefault() error {
	if len(this.backends) == 0 {
		return errors.New("no backend found in config, forgot Configure().To() call?")
	}

	// TODO: support multiple backends
	defaulBackend = this.backends[0]
	return nil
}

type BackendBuilder interface {
	Build() Backend
}