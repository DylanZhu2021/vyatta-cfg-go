package get

// GetEnvName ${vyos_libexec_dir} ==> vyos_libexec_dir
func GetEnvName(env string) string {
	return string(env[2 : len(env)-1])
}
