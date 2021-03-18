package main

type panicWriter struct {
}

func (e panicWriter) Write(p []byte) (int, error) {
	log.Error(string(p))
	return len(p), nil
}

type accessWriter struct {
}

func (e accessWriter) Write(p []byte) (int, error) {
	log.Info(string(p))
	return len(p), nil
}
